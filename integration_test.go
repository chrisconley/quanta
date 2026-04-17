package quanta

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestInvoiceCalculationPipeline demonstrates the full Measure->Quantized pipeline
// for a realistic billing scenario: calculate subtotal, tax, discounts, then
// quantize for invoice line items.
func TestInvoiceCalculationPipeline(t *testing.T) {
	usdSpec := MustNewUnit("USD", "0.01")

	// Line item 1: Product A
	productA := NewMeasureFrom(usdSpec, MustNewDecimal("19.99"))

	// Line item 2: Product B
	productB := NewMeasureFrom(usdSpec, MustNewDecimal("34.50"))

	// Subtotal (high precision)
	subtotal, err := productA.Add(productB)
	assert.NoError(t, err)
	assert.True(t, subtotal.Quantity().Equal(MustNewDecimal("54.49")))

	// Apply 10% discount (high precision)
	discountRate := MustNewDecimal("0.10")
	discount, err := subtotal.Mul(discountRate)
	assert.NoError(t, err)
	assert.True(t, discount.Quantity().Equal(MustNewDecimal("5.449"))) // High precision

	// Discounted subtotal
	discountedSubtotal, err := subtotal.Sub(discount)
	assert.NoError(t, err)
	assert.True(t, discountedSubtotal.Quantity().Equal(MustNewDecimal("49.041")))

	// Apply 8.5% tax (high precision)
	taxRate := MustNewDecimal("0.085")
	tax, err := discountedSubtotal.Mul(taxRate)
	assert.NoError(t, err)
	// Tax is 4.16848... (high precision)

	// Total before quantization
	total, err := discountedSubtotal.Add(tax)
	assert.NoError(t, err)

	// Quantize for invoice (round tax up to favor business)
	quantizedProductA := productA.Quantize(RoundHalfEven).Value
	quantizedProductB := productB.Quantize(RoundHalfEven).Value
	quantizedDiscount := discount.Quantize(RoundHalfEven).Value
	quantizedTax := tax.Quantize(RoundUp).Value // Round tax up
	quantizedTotal := total.Quantize(RoundHalfEven).Value

	// Verify quantized values
	assert.Equal(t, int64(1999), quantizedProductA.Multiplier())
	assert.True(t, quantizedProductA.Decimal().Equal(MustNewDecimal("19.99")))

	assert.Equal(t, int64(3450), quantizedProductB.Multiplier())
	assert.True(t, quantizedProductB.Decimal().Equal(MustNewDecimal("34.50")))

	assert.Equal(t, int64(545), quantizedDiscount.Multiplier())
	assert.True(t, quantizedDiscount.Decimal().Equal(MustNewDecimal("5.45")))

	assert.Equal(t, int64(417), quantizedTax.Multiplier())
	assert.True(t, quantizedTax.Decimal().Equal(MustNewDecimal("4.17")))

	assert.Equal(t, int64(5321), quantizedTotal.Multiplier())
	assert.True(t, quantizedTotal.Decimal().Equal(MustNewDecimal("53.21")))

	// Verify invoice line items can be aggregated exactly by summing multipliers
	invoiceTotal := quantizedProductA.Multiplier() + quantizedProductB.Multiplier() - quantizedDiscount.Multiplier() + quantizedTax.Multiplier()
	assert.Equal(t, int64(5321), invoiceTotal)
	assert.Equal(t, quantizedTotal.Multiplier(), invoiceTotal)
}

// TestMeterAggregationPipeline demonstrates the full pipeline for metering:
// collect usage records, aggregate with high precision, then quantize for storage.
func TestMeterAggregationPipeline(t *testing.T) {
	tokensSpec := MustNewUnit("tokens", "0.001")

	// Simulate 5 API requests with varying token usage
	request1 := NewMeasureFrom(tokensSpec, MustNewDecimal("12.3456"))
	request2 := NewMeasureFrom(tokensSpec, MustNewDecimal("8.7654"))
	request3 := NewMeasureFrom(tokensSpec, MustNewDecimal("15.2345"))
	request4 := NewMeasureFrom(tokensSpec, MustNewDecimal("3.9876"))
	request5 := NewMeasureFrom(tokensSpec, MustNewDecimal("9.1234"))

	// Aggregate (high precision)
	sum1, err := request1.Add(request2)
	assert.NoError(t, err)
	sum2, err := sum1.Add(request3)
	assert.NoError(t, err)
	sum3, err := sum2.Add(request4)
	assert.NoError(t, err)
	total, err := sum3.Add(request5)
	assert.NoError(t, err)

	// Verify high-precision total
	expected := MustNewDecimal("49.4565")
	assert.True(t, total.Quantity().Equal(expected))

	// Quantize for storage (neutral rounding)
	quantizedTotal := total.Quantize(RoundHalfEven).Value

	assert.Equal(t, int64(49456), quantizedTotal.Multiplier())
	assert.True(t, quantizedTotal.Decimal().Equal(MustNewDecimal("49.456")))

	// Simulate storing to database (store multiplier)
	storedMultiplier := quantizedTotal.Multiplier()

	// Simulate reading from database and reconstructing
	reconstructed := DeprecatedNewQuantizedFromMultiplier(tokensSpec, storedMultiplier)

	assert.True(t, reconstructed.Equal(quantizedTotal))
	assert.True(t, reconstructed.Decimal().Equal(MustNewDecimal("49.456")))
}

// TestRatingPipeline demonstrates usage-based pricing:
// usage (tokens) x rate ($/token) -> cost, then quantize for billing.
func TestRatingPipeline(t *testing.T) {
	// Usage spec
	tokensSpec := MustNewUnit("tokens", "0.0001")

	// Money spec
	usdSpec := MustNewUnit("USD", "0.01")

	// Customer used 12,345.6789 tokens
	usage := NewMeasureFrom(tokensSpec, MustNewDecimal("12345.6789"))

	// Rate: $0.0025 per token
	// Note: We use scalar multiplication since we can't multiply Measure x Measure
	ratePerToken := MustNewDecimal("0.0025")

	// Calculate cost (high precision)
	// In a real system, this would be Measure x Rate -> Measure
	// For this test, we multiply the usage value by the rate
	costValue, err := usage.Quantity().Mul(ratePerToken)
	assert.NoError(t, err)

	cost := NewMeasureFrom(usdSpec, costValue)

	// Verify high-precision cost
	expected := MustNewDecimal("30.86419725")
	assert.True(t, cost.Quantity().Equal(expected))

	// Quantize for billing (round half-even)
	quantizedCost := cost.Quantize(RoundHalfEven).Value

	assert.Equal(t, int64(3086), quantizedCost.Multiplier())
	assert.True(t, quantizedCost.Decimal().Equal(MustNewDecimal("30.86")))
}

// TestMultiCurrencyInvoice demonstrates type safety: cannot mix different currencies.
func TestMultiCurrencyInvoice(t *testing.T) {
	usdSpec := MustNewUnit("USD", "0.01")
	eurSpec := MustNewUnit("EUR", "0.01")

	usdAmount := NewMeasureFrom(usdSpec, MustNewDecimal("100.00"))
	eurAmount := NewMeasureFrom(eurSpec, MustNewDecimal("85.00"))

	// Attempting to add USD + EUR should fail
	_, err := usdAmount.Add(eurAmount)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "incompatible specs")
	assert.Contains(t, err.Error(), "USD")
	assert.Contains(t, err.Error(), "EUR")
}

// TestQuantumMismatchCalculation demonstrates that measures with different quanta
// (but same code) can be calculated together, but quantized values must match exactly.
func TestQuantumMismatchCalculation(t *testing.T) {
	// Two USD specs with different precision
	highPrecision := MustNewUnit("USD", "0.0001")
	lowPrecision := MustNewUnit("USD", "0.01")

	// Measures with different quanta can be added (CompatibleForCalc)
	amount1 := NewMeasureFrom(highPrecision, MustNewDecimal("10.5678"))
	amount2 := NewMeasureFrom(lowPrecision, MustNewDecimal("5.25"))

	total, err := amount1.Add(amount2)
	assert.NoError(t, err)
	assert.True(t, total.Quantity().Equal(MustNewDecimal("15.8178")))

	// Quantized values with different quanta are NOT equal (different specs)
	quantized1 := amount1.Quantize(RoundHalfEven).Value
	quantized2 := amount2.Quantize(RoundHalfEven).Value

	assert.False(t, quantized1.Unit().Equal(quantized2.Unit()))
	assert.False(t, quantized1.Equal(quantized2)) // Even if we quantized the same value, specs differ
}

// TestTieredPricingCalculation demonstrates a tiered pricing scenario:
// different rates for different usage tiers, with final quantization.
func TestTieredPricingCalculation(t *testing.T) {
	usdSpec := MustNewUnit("USD", "0.01")

	// Tier 1: 0-100 units @ $0.10/unit = $10.00
	tier1Units := MustNewDecimal("100")
	tier1Rate := MustNewDecimal("0.10")
	tier1Cost, err := tier1Units.Mul(tier1Rate)
	assert.NoError(t, err)

	// Tier 2: 101-500 units (400 units) @ $0.08/unit = $32.00
	tier2Units := MustNewDecimal("400")
	tier2Rate := MustNewDecimal("0.08")
	tier2Cost, err := tier2Units.Mul(tier2Rate)
	assert.NoError(t, err)

	// Tier 3: 501-1000 units (234 units) @ $0.05/unit = $11.70
	tier3Units := MustNewDecimal("234")
	tier3Rate := MustNewDecimal("0.05")
	tier3Cost, err := tier3Units.Mul(tier3Rate)
	assert.NoError(t, err)

	// Total usage: 734 units, Total cost calculation
	tier1Measure := NewMeasureFrom(usdSpec, tier1Cost)
	tier2Measure := NewMeasureFrom(usdSpec, tier2Cost)
	tier3Measure := NewMeasureFrom(usdSpec, tier3Cost)

	sum1, err := tier1Measure.Add(tier2Measure)
	assert.NoError(t, err)
	totalCost, err := sum1.Add(tier3Measure)
	assert.NoError(t, err)

	// Verify total: $53.70
	assert.True(t, totalCost.Quantity().Equal(MustNewDecimal("53.70")))

	// Quantize for invoice
	quantizedTotal := totalCost.Quantize(RoundHalfEven).Value

	assert.Equal(t, int64(5370), quantizedTotal.Multiplier())
	assert.True(t, quantizedTotal.Decimal().Equal(MustNewDecimal("53.70")))
}

// TestAllocationRoundingError demonstrates the importance of quantization:
// splitting a total into parts may introduce rounding errors that must be handled.
func TestAllocationRoundingError(t *testing.T) {
	usdSpec := MustNewUnit("USD", "0.01")

	// Total amount to split: $10.00
	total := NewMeasureFrom(usdSpec, MustNewDecimal("10.00"))

	// Split equally among 3 parties (high precision: $3.333...)
	parties := MustNewDecimal("3")
	shareValue, err := total.Quantity().Div(parties)
	assert.NoError(t, err)
	share := NewMeasureFrom(usdSpec, shareValue)

	// Each share is $3.333333... (precision is 34 total digits, so 33 after decimal)
	expected := MustNewDecimal("3.333333333333333333333333333333333")
	assert.True(t, share.Quantity().Equal(expected))

	// Quantize each share
	quantizedShare := share.Quantize(RoundHalfEven).Value

	assert.Equal(t, int64(333), quantizedShare.Multiplier())
	assert.True(t, quantizedShare.Decimal().Equal(MustNewDecimal("3.33")))

	// If we give $3.33 to each of 3 parties, we've only distributed $9.99
	distributed := quantizedShare.Multiplier() * 3
	assert.Equal(t, int64(999), distributed)

	// There's $0.01 of "dust" remaining ($10.00 - $9.99)
	// This is where allocation strategies (like largest-remainder) would be needed
	// to ensure sum(parts) == whole
	totalMultiplier := int64(1000) // $10.00
	dust := totalMultiplier - distributed
	assert.Equal(t, int64(1), dust) // $0.01 dust

	// In a real system, this dust would be allocated to one of the parties
	// using an allocation strategy (to be implemented in future commits)
}

// TestNegativeValues demonstrates handling of credits, refunds, and negative amounts.
func TestNegativeValues(t *testing.T) {
	usdSpec := MustNewUnit("USD", "0.01")

	// Original charge: $100.00
	charge := NewMeasureFrom(usdSpec, MustNewDecimal("100.00"))

	// Partial refund: -$25.50
	refund := NewMeasureFrom(usdSpec, MustNewDecimal("-25.50"))

	// Net amount
	net, err := charge.Add(refund)
	assert.NoError(t, err)
	assert.True(t, net.Quantity().Equal(MustNewDecimal("74.50")))

	// Quantize
	quantizedCharge := charge.Quantize(RoundHalfEven).Value
	quantizedRefund := refund.Quantize(RoundHalfEven).Value
	quantizedNet := net.Quantize(RoundHalfEven).Value

	assert.Equal(t, int64(10000), quantizedCharge.Multiplier())
	assert.Equal(t, int64(-2550), quantizedRefund.Multiplier())
	assert.Equal(t, int64(7450), quantizedNet.Multiplier())

	// Verify arithmetic: charge + refund = net
	calculatedNet := quantizedCharge.Multiplier() + quantizedRefund.Multiplier()
	assert.Equal(t, quantizedNet.Multiplier(), calculatedNet)
}

// TestDoubleRoundingDivergence proves that rounding intermediates produces
// a different (wrong) answer than rounding once at the end.
//
// The calculation: $10.00 × 1/3 × 1/3 × 9 = $10.00 exactly.
// Rounding after each step accumulates error: $10.00 → $3.33 → $1.11 → $9.99.
// Keeping full precision through all steps and rounding once: $10.00.
func TestDoubleRoundingDivergence(t *testing.T) {
	usd := MustNewUnit("USD", "0.01")
	one := MustNewDecimal("1")
	three := MustNewDecimal("3")
	nine := MustNewDecimal("9")
	third, err := one.Div(three)
	assert.NoError(t, err)

	start := NewMeasureFrom(usd, MustNewDecimal("10.00"))

	// WRONG: round each intermediate
	naive1, err := start.Mul(third)
	assert.NoError(t, err)
	naiveQ1 := naive1.Quantize(RoundHalfEven) // 3.33

	naive2, err := NewMeasureFrom(usd, naiveQ1.Value.Decimal()).Mul(third)
	assert.NoError(t, err)
	naiveQ2 := naive2.Quantize(RoundHalfEven) // 1.11

	naive3, err := NewMeasureFrom(usd, naiveQ2.Value.Decimal()).Mul(nine)
	assert.NoError(t, err)
	naiveResult := naive3.Quantize(RoundHalfEven) // 9.99 ← wrong

	// CORRECT: round once at the end
	correct1, err := start.Mul(third) // 3.333...
	assert.NoError(t, err)
	correct2, err := correct1.Mul(third) // 1.111...
	assert.NoError(t, err)
	correct3, err := correct2.Mul(nine) // 10.00
	assert.NoError(t, err)
	correctResult := correct3.Quantize(RoundHalfEven) // 10.00 ← correct

	assert.Equal(t, int64(999), naiveResult.Value.Multiplier(),
		"rounding intermediates: $10 × 1/3 × 1/3 × 9 drifts to $9.99")
	assert.Equal(t, int64(1000), correctResult.Value.Multiplier(),
		"rounding once: $10 × 1/3 × 1/3 × 9 = $10.00")
}
