# quanta

Billing-grade decimal arithmetic for Go. Calculate in high precision, round once at the boundary.

[![Go Reference](https://pkg.go.dev/badge/github.com/chrisconley/quanta.svg)](https://pkg.go.dev/github.com/chrisconley/quanta)

> **Status:** Pre-1.0. The core types and pipeline are stable and exercised by the test suite, but the public API may still change before `v1.0.0`.

## The problem

If you've built a billing, metering, or ledger system in Go, you've probably hit at least one of these:

- **Invoice totals that don't match the sum of line items.** Tax got rounded before it was added, so the stored total is a cent off from `sum(lines)`.
- **Splits that lose a penny.** `$10.00 / 3` rounds to `$3.33` three times, and the customer's statement is short `$0.01`.
- **Dollars silently added to euros.** Both are `float64` (or `decimal.Decimal`), so the compiler can't stop you.
- **Exact-equality checks that fail on posted values.** `"12.50"` and `"12.5"` should be the same invoice amount, but `==` on the decimal strings disagrees.
- **Rounding logic scattered across the codebase.** Every service re-implements "snap to cents" and they don't all agree on how to break ties.

`quanta` is the arithmetic layer for systems that care about these things. It is not a billing platform, not a double-entry ledger, and not a drop-in `decimal` replacement. It is a small set of domain types that make the right thing easy and the wrong thing a compile or runtime error.

It's aimed at Go teams building **billing, metering, subscription, or payments systems** — places where the bugs above are career risks, not annoyances. If you just need precise decimal math without unit, rounding, and allocation domain types, reach for [`cockroachdb/apd`](https://github.com/cockroachdb/apd) (which `quanta` is built on) or [`shopspring/decimal`](https://github.com/shopspring/decimal) instead.

## Scope

`quanta` is the **arithmetic layer**: values in, values out, with explicit rounding at the boundary.

**In scope:** exact decimal arithmetic, unit-safe `Measure` and `Quantized` types, rounding modes, allocation strategies with explicit dust handling.

**Out of scope:** currency conversion / FX rates, tax engines, invoicing and billing orchestration, double-entry bookkeeping, event-stream metering, persistence and serialization format opinions, multi-currency arithmetic on a single `Measure`.

The boundary is intentional. Pricing, ledgers, and metering pipelines all need exact math — but each has its own rules about history, equality, and audit. `quanta` gives you a typed value you can hand off to those systems and never wants it back.

## Install

```sh
go get github.com/chrisconley/quanta
```

Requires Go 1.23 or later.

## Hello world

Calculate a line total with tax at full precision, then snap to cents exactly once for the invoice:

```go
package main

import (
    "fmt"

    "github.com/chrisconley/quanta"
)

func main() {
    usd := quanta.UnitSpec{Code: "USD", Quantum: "0.01"}

    subtotal := quanta.MustNewMeasure(quanta.MeasureSpec{Unit: usd, Quantity: "19.99"})
    taxRate := quanta.MustNewDecimal("0.08875") // 8.875%

    tax, _ := subtotal.Mul(taxRate)                // 1.7741125 (full precision)
    total, _ := subtotal.Add(tax)                  // 21.7641125 (full precision)
    result := total.Quantize(quanta.RoundHalfEven) // snap to $0.01, once

    fmt.Println("working:", total.Quantity())          // 21.7641125
    fmt.Println("posted: ", result.Value.Decimal())    // 21.76
    fmt.Println("dust:   ", result.Remainder.Quantity()) // 0.0041125
}
```

## Core idea: two types, one boundary

```
   Measure  ──── calc, calc, calc ────▶  Measure  ──▶  Quantize ──▶  Quantized
   (working)                              (working)    (once)        (posted fact)
```

- **`Measure`** is a **working value**. High precision. Supports `Add`, `Sub`, `Mul`, `Div`. Use it inside calculations.
- **`Quantized`** is a **posted fact**. Stored as an integer multiplier of the unit's quantum. No arithmetic — if you want to compute with it, you're in the wrong type. Use it at boundaries: database, API response, invoice row.
- **`Quantize`** is the one-way bridge. Rounding happens here and nowhere else. It returns both the posted value **and** the remainder (the "dust") so you can audit it or reinject it later.

This separation is the whole point. You can't accidentally round an intermediate result, because intermediate results are always `Measure`. And you can't accidentally keep computing with a posted value, because `Quantized` has no math.

## Why this design

**Round once, at the boundary.** Intermediate calculations stay at full precision. Rounding happens only in `Quantize`, which converts `Measure` into `Quantized`. You can't accidentally round an intermediate because `Measure` has no rounding method — there's no way to call it.

**Working and posted are different types.** `Measure` is for computation; `Quantized` is a posted fact. `Quantized` has no arithmetic — so if you want to keep computing with a posted value, you're already in the wrong type. This eliminates a class of bugs where "the stored value" and "the recomputed value" silently diverge.

**Unit compatibility is checked on every operation.** Every `Measure` carries its `Unit` (code + quantum). `USD + EUR` fails before it reaches your invoice. `USD@0.01` and `USD@0.001` are different units — same currency, different precision, deliberately not interchangeable.

**Dust is explicit.** `Quantize` returns both the posted value and the remainder. Allocation strategies surface the trade-off: `LargestRemainderStrategy` distributes dust so parts sum exactly; `ProRataTruncateStrategy` truncates and returns the remainder so you can route it to a house account. `quanta` does not silently absorb rounding error.

**Exact equality on posted values.** `Quantized` is stored as an integer multiplier of the unit's quantum. `"12.50"` and `"12.5"` resolve to the same integer — they're equal, they serialize identically, and sums of `Quantized` values have no floating-point drift.

## Progressive examples

### 1. Enforced unit compatibility

```go
usd := quanta.UnitSpec{Code: "USD", Quantum: "0.01"}
eur := quanta.UnitSpec{Code: "EUR", Quantum: "0.01"}

dollars := quanta.MustNewMeasure(quanta.MeasureSpec{Unit: usd, Quantity: "100.00"})
euros := quanta.MustNewMeasure(quanta.MeasureSpec{Unit: eur, Quantity: "90.00"})

_, err := dollars.Add(euros)
// err: cannot add measures with incompatible specs: USD[q=0.01] vs EUR[q=0.01]
```

The unit code travels with the value. `USD` and `EUR` are different types at runtime even though they share a storage representation.

### 2. Splitting a total without losing pennies

`$10.00 / 3` is the canonical hard case: naive division gives `$3.33 × 3 = $9.99` and loses a penny. The **largest remainder** method distributes the dust so the parts sum to exactly the quantized total:

```go
usd := quanta.UnitSpec{Code: "USD", Quantum: "0.01"}
total := quanta.MustNewMeasure(quanta.MeasureSpec{Unit: usd, Quantity: "10.00"})

strategy := quanta.LargestRemainderStrategy{}
result, _ := strategy.Allocate(total, []int64{1, 1, 1}, quanta.RoundHalfEven)

for i, part := range result.Parts {
    fmt.Printf("Part %d: %s\n", i+1, part.Decimal())
}
// Part 1: 3.34
// Part 2: 3.33
// Part 3: 3.33
```

Guarantees: `sum(Parts) == Quantized total` and `Dust == 0`.

If you'd rather track dust explicitly (for example, to route it to a house account), use `ProRataTruncateStrategy`:

```go
strategy := quanta.ProRataTruncateStrategy{}
result, _ := strategy.Allocate(total, []int64{1, 1, 1}, quanta.RoundHalfEven)
// Parts: [3.33, 3.33, 3.33], Dust: 0.01
```

### 3. Invoice pipeline: many lines, one quantization pass

```go
usd := quanta.MustNewUnit("USD", "0.01")

productA := quanta.NewMeasureFrom(usd, quanta.MustNewDecimal("19.99"))
productB := quanta.NewMeasureFrom(usd, quanta.MustNewDecimal("34.50"))

subtotal, _ := productA.Add(productB)                         // full precision
discount, _ := subtotal.Mul(quanta.MustNewDecimal("0.10"))    // full precision
net, _     := subtotal.Sub(discount)                          // full precision
tax, _     := net.Mul(quanta.MustNewDecimal("0.085"))         // full precision
total, _   := net.Add(tax)                                    // full precision

// Quantize at the boundary. Rounding mode can vary per field
// (e.g. round tax up to favor the business).
invoiceTax   := tax.Quantize(quanta.RoundUp).Value
invoiceTotal := total.Quantize(quanta.RoundHalfEven).Value
```

Because `Quantized` stores an integer multiplier of the quantum, you can also sum line totals exactly by summing multipliers — no floating-point drift, no trailing-zero normalization.

### 4. Metering at sub-penny precision

Nothing about `quanta` is money-specific. Define a unit for anything countable:

```go
tokens := quanta.UnitSpec{Code: "tokens", Quantum: "0.001"}
calls  := quanta.UnitSpec{Code: "api-calls", Quantum: "1"}

usage := quanta.MustNewMeasure(quanta.MeasureSpec{Unit: tokens, Quantity: "1250.4375"})
rate  := quanta.MustNewDecimal("0.00002")       // $/token
cost  := usage.Mul(rate)                        // compute at token precision
```

## Features

- **Decimal arithmetic** backed by [`cockroachdb/apd`](https://github.com/cockroachdb/apd) (IEEE 754 decimal128, 34-digit precision, banker's rounding default).
- **`Unit` = code + quantum**: one concept, not two. `USD@0.01` and `USD@0.001` are different units.
- **`Measure`**: high-precision working values with `Add`/`Sub`/`Mul`/`Div` and unit-compatibility checks.
- **`Quantized`**: posted facts as integer multiples of a quantum. Exact equality, safe to store and compare.
- **Seven rounding modes**: `HalfEven` (default), `HalfUp`, `HalfDown`, `Up`, `Down`, `Ceiling`, `Floor`.
- **Allocation strategies**:
  - `LargestRemainderStrategy` — dust-free, parts always sum to the quantized total.
  - `ProRataTruncateStrategy` — truncates each share and returns the remainder explicitly.
- **Zero-dependency public API.** `apd` is an internal implementation detail; consumers never touch it.
- **Comprehensive tests**: unit, integration, property-based (fuzz), and benchmarks.

## How it compares

| You need… | Use |
|---|---|
| Raw decimal arithmetic, no domain concepts | [`cockroachdb/apd`](https://github.com/cockroachdb/apd), [`shopspring/decimal`](https://github.com/shopspring/decimal) |
| Money-typed amounts with currency metadata | [`Rhymond/go-money`](https://github.com/Rhymond/go-money) |
| **Working-vs-posted separation, unit safety, allocation** | **`quanta`** |
| Double-entry bookkeeping | A ledger library (e.g. [`formancehq/ledger`](https://github.com/formancehq/ledger)) on top of `quanta` |
| A full billing platform (metering, invoicing, dunning) | Lago, Kill Bill, OpenMeter |

## API reference

Full API documentation: [pkg.go.dev/github.com/chrisconley/quanta](https://pkg.go.dev/github.com/chrisconley/quanta).

Main types:

| Type | Role |
|---|---|
| `Decimal` | Immutable high-precision decimal. Thin wrapper over `apd.Decimal`. |
| `Quantum` | Smallest meaningful increment for a unit (e.g. `0.01` USD, `1` JPY). |
| `Unit` | `code + quantum`. Unit compatibility is checked on every operation. |
| `Measure` | Working value. Mutable-free arithmetic, high precision, unit-checked. |
| `Quantized` | Posted fact. Integer multiplier of quantum. No arithmetic. |
| `QuantizationResult` | `{ Value Quantized; Remainder Measure }` returned from `Measure.Quantize`. |
| `Rounding` | One of seven rounding modes. |
| `AllocationStrategy` | Interface implemented by `LargestRemainderStrategy` and `ProRataTruncateStrategy`. |

## Development

```sh
go test ./...              # unit + integration tests
go test -run=Example ./... # runnable documentation examples
go test -fuzz=Fuzz         # property-based tests
go test -bench=. ./...     # benchmarks
```

## Contributing

Issues and pull requests are welcome. Before opening a PR, please run `go test ./...` and make sure the examples under `example_test.go` still match their expected output.

## Credits

`quanta` is built on [`cockroachdb/apd`](https://github.com/cockroachdb/apd) (Apache 2.0), which does the heavy lifting of decimal arithmetic. Thanks to its maintainers for a rock-solid foundation.

## License

MIT. See [LICENSE](LICENSE).
