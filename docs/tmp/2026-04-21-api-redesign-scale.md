# API Redesign — Scale & Minimal Surface

Scratch notes from the session that scoped quanta's API toward a minimal
public surface and introduced `Scale` as a first-class primitive. Everything
here is **decision + rationale**; the migration plan is in the last section.

## Where we landed

### Public API (target shape)

Six concepts, no Decimal in signatures (likely), no Quantum as a user type:

- `Unit` — code + quantum (constructed via string quantum)
- `Measure` — high-precision working value
- `Quantized` — posted fact, int64 multiplier
- `Scale` — unit-checked scalar transform (new)
- `Rounding` — enum
- `AllocationStrategy` + `LargestRemainderStrategy` + `ProRataTruncateStrategy`

Plus result types: `QuantizationResult`, `AllocationResult`.

Cut from public surface:
- `Measure.Mul(Decimal)` / `Measure.Div(Decimal)` — footgun shortcuts; every
  real use case is Rate-shaped (tax/FX/markup) or allocation-shaped.
  Averaging-over-count is the only genuine dimensionless case and is better
  expressed as sugar over Scale when demand warrants.
- `Quantum` as a user-facing type — already only 2 refs in meters-credits.
  Stays internal; Unit construction takes a quantum *string*.

Deferred (gate after Initiative 1 merges):
- Whether `Decimal` stays exported. 267 `quanta.Decimal`/`NewDecimal` refs
  and 249 `.Quantity()`/`.Decimal()` accessor calls in meters-credits —
  hiding is a 3–6 month migration. Option B (add string constructors + query
  methods, leave Decimal as power-user escape hatch) likely permanent.

### Scale, not Rate / Transform / Conversion

Debate covered four names; Scale won on scope-precision:

- **Conversion** — ruled out by same-unit case ("apply 10% tip to $100 → $100"
  doesn't convert anything).
- **Rate** — strong candidate (finance-familiar, composes via "effective
  rate"), but "rate" is already rating's domain word in meters-credits.
  More importantly: "Rate" invites scope creep (curves, lookup tables).
- **Transform** — too generic. "Transform is idiomatic Go" claim overstated
  (Go uses it for byte/text streams, not value→value functions).
- **Scale** — wins on (a) identity naturalness ("scale by 1" ≠ awkward),
  (b) semantic precision (signals multiplicative-only scope), (c) no domain
  collision, (d) dimension-neutral on same-unit and cross-unit.

Tradeoff: Scale would need to retire if quanta ever grew affine transforms
(`a*x + b`, e.g. flat-fee-plus-per-unit). Rate could have stretched. Betting
affine isn't coming.

### Piecewise rates: Option C (split primitives down, compositions up)

Considered four placements for graduated/volume rates:

- A: everything in quanta (overreach — quanta becomes a billing engine)
- B: only Linear in quanta (leaks re-wrap pattern back into rating)
- **C: primitives down (Linear, Slice, Bucket), compositions up**
- D: A + policy tree (overreach squared)

**Landed on C.** Graduated (sum-slices) and volume (pick-bucket) are two
algorithms over one data structure. Choosing *which* is pricing policy,
which belongs in rating. Quanta provides the dimensional primitives:

- `Measure.Slice(bounds []Measure) ([]Measure, error)` — partition
- `Measure.Bucket(bounds []Measure) (int, error)` — index lookup
- `Scale.Apply` — unit-checked scalar multiply + retag

`graduatedRate.convert` and `volumeRate.convert` shrink to ~10 lines each
against these primitives. `tier`, `rateConverter`, manual re-wraps all
collapse.

### Input-unit bug to fix

`meters-credits/rating/rate_impl.go:linearRate` holds `{factor, outputUnit}`
only — no input-unit check. `convert` can silently accept a Measure in the
wrong unit and re-tag. `Scale.Apply` enforces `m.Unit() == s.From()`, which
closes the bug by construction.

## Open question 1: Affine transforms

Scale is multiplicative-only by design. If a two-part-tariff use case
(flat fee + per-unit rate) appears, Scale can't express it. Options:

1. Keep Scale multiplicative; compose flat-fee separately (as a Measure
   addition in rating).
2. Extend Scale to affine (`NewAffineScale(from, to, factor, offset)`).
3. Introduce a second primitive (`AffineTransform` or similar).

Current bet: (1) — flat fees are pricing policy and belong in rating, not
quanta. Revisit if a consumer emerges that needs affine at the dimensional
layer.

## Open question 2: Slice boundary semantics

`Measure.Slice(bounds)` needs a committed policy on:

- Inclusive-lower / exclusive-upper? (standard for tax brackets)
- Sorted bounds required, or self-sort?
- Overflow: extra tail slice when input exceeds last bound, or error?
- Negative input: error, or valid (mirrored)?

These are dimensional-adjacent but not purely dimensional — e.g., tax-bracket
edge semantics at exactly $100,000 are a policy. Recommendation: commit
to inclusive-lower/exclusive-upper + sorted-required + extra-tail, document
rigidly, and let rating layer any alternate semantics on top. Pin with tests
so the choice is visible.

## Open question 3: Demand-driven sugar

Deferred until a consumer asks:

- `Scale.Then(Scale) (Scale, error)` — composition with construction-time
  dimension check
- `IdentityScale(unit) Scale` — same-unit, factor 1
- `Measure.RatioOf(Measure) (Scale, error)` — returns `Scale(b.Unit→a.Unit)`
- `Measure.Mean([]Measure) Measure` — sugar over Scale
- `Zero(unit) Measure` — sugar for `MustNewMeasure(u, "0")`
- `Scale.FactorString()` / `Scale.String()` — display

Each is a one-method quanta commit paired with an immediate consumer
migration. No speculative "obvious" additions.

## Open question 4: Graduated-vs-volume elevation to quanta

Would push Option A (full piecewise in quanta) if:

- A second consumer in the codebase needs piecewise-linear conversion
  (taxes, multi-currency spreads, energy brackets).
- The graduated/volume distinction turns out to be not-two-algorithms-
  but-a-lattice (stepped-volume, capped-graduated, blended).

Not there yet. Option C holds.

## Migration plan (summary)

Every quanta commit ships the minimum that unblocks one meters-credits
migration. Each mandatory quanta release is a single feature.

**Initiative 1 — Scale + linearRate**
- Quanta: `Scale` type with `NewLinearScale`, `MustNewLinearScale`, `Apply`,
  `From`, `To` (nothing else yet).
- meters-credits: migrate `linearRate` to embed `Scale`. Input-unit bug fixed.

**Initiative 2 — Measure.Slice + graduatedRate**
- Quanta: `Measure.Slice`.
- meters-credits: migrate `graduatedRate.convert`.

**Initiative 3 — Measure.Bucket + volumeRate**
- Quanta: `Measure.Bucket`.
- meters-credits: migrate `volumeRate.convert`.

**Initiative 4 — cut Mul/Div footguns**
- meters-credits: migrate 9 `.Mul()/.Div()` call sites (5 files).
- Quanta: remove `Measure.Mul(Decimal)`, `Measure.Div(Decimal)`.

**Initiative 5 — decision gate on hiding Decimal**
- Not a plan. Revisit after Initiatives 1–4 merge; remaining Decimal
  footprint in meters-credits dictates Option A vs Option B.

Demand-driven additions (Then, Identity, RatioOf, Mean, Zero, display) land
as one-commit micro-releases during/after the above, only when a consumer
commit actually needs them.

## Out-of-scope surfaced during this session

- **Naming collision in meters-credits.** Its `rating.Rate` type stays;
  quanta uses `Scale` specifically to avoid forcing a downstream rename.
- **`linearRate` may vanish entirely.** Once migrated to `Scale`, the wrapper
  struct is near-empty. Whether to delete it or keep as a named domain type
  is a meters-credits-internal decision — not blocking quanta.
- **`quanta.Zero` ambiguity.** Current `Zero()` returns `Decimal`. If a
  `Zero(unit Unit) Measure` overload lands, the naming collision needs
  resolving (rename the Decimal one to `ZeroDecimal`, or accept the overload
  asymmetry). Surface when demand arrives.
- **AllocationStrategy + Scale.** Allocation weights are currently `[]int64`.
  A Scale-based weight API is conceivable but out of scope — allocation
  semantics (dust, largest-remainder tiebreaks) are orthogonal to Scale.
