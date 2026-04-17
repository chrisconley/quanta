package quanta_test

import (
	"fmt"

	"github.com/chrisconley/quanta"
)

// ExampleMeasure_Quantize demonstrates the calculate-then-quantize workflow.
// All arithmetic happens at full quanta in Measure. Rounding happens once
// when crossing the boundary to Quantized — the posted fact for storage or invoicing.
func ExampleMeasure_Quantize() {
	usd := quanta.UnitSpec{Code: "USD", Quantum: "0.01"}

	subtotal := quanta.MustNewMeasure(quanta.MeasureSpec{Unit: usd, Quantity: "19.99"})
	taxRate := quanta.MustNewDecimal("0.08875") // 8.875% tax

	tax, _ := subtotal.Mul(taxRate)                // full quanta: 1.7741125
	total, _ := subtotal.Add(tax)                  // full quanta: 21.7641125
	result := total.Quantize(quanta.RoundHalfEven) // snap to $0.01

	fmt.Println(total.Quantity())            // working value (unrounded)
	fmt.Println(result.Value.Decimal())      // posted fact (rounded once)
	fmt.Println(result.Remainder.Quantity()) // dust from rounding
	// Output:
	// 21.7641125
	// 21.76
	// 0.0041125
}

// ExampleLargestRemainderStrategy_Allocate demonstrates dust-free splitting.
// $10.00 split three ways: naive division gives $3.33 × 3 = $9.99 (lost penny).
// Largest remainder distributes the extra penny to the part with the largest
// fractional remainder, guaranteeing the parts sum to exactly the total.
func ExampleLargestRemainderStrategy_Allocate() {
	usd := quanta.UnitSpec{Code: "USD", Quantum: "0.01"}
	total := quanta.MustNewMeasure(quanta.MeasureSpec{Unit: usd, Quantity: "10.00"})

	strategy := quanta.LargestRemainderStrategy{}
	result, _ := strategy.Allocate(total, []int64{1, 1, 1}, quanta.RoundHalfEven)

	for i, part := range result.Parts {
		fmt.Printf("Part %d: %s\n", i+1, part.Decimal())
	}
	// Output:
	// Part 1: 3.34
	// Part 2: 3.33
	// Part 3: 3.33
}

// ExampleMeasure_Add demonstrates unit compatibility enforcement.
// Measures with the same unit code can be added; different codes are rejected.
func ExampleMeasure_Add() {
	usd := quanta.UnitSpec{Code: "USD", Quantum: "0.01"}
	eur := quanta.UnitSpec{Code: "EUR", Quantum: "0.01"}

	dollars := quanta.MustNewMeasure(quanta.MeasureSpec{Unit: usd, Quantity: "100.00"})
	euros := quanta.MustNewMeasure(quanta.MeasureSpec{Unit: eur, Quantity: "90.00"})
	moreDollars := quanta.MustNewMeasure(quanta.MeasureSpec{Unit: usd, Quantity: "50.00"})

	sum, _ := dollars.Add(moreDollars)
	fmt.Println(sum.Quantity())

	_, err := dollars.Add(euros)
	fmt.Println(err)
	// Output:
	// 150.00
	// cannot add measures with incompatible specs: USD[q=0.01] vs EUR[q=0.01]
}
