# Scale Migration Plan — Smallest-Unblock Cadence

Migration plan for introducing `Scale` into quanta and migrating
meters-credits to use it. Every quanta commit ships the minimum that
enables exactly one meters-credits migration, then meters-credits
migrates, then quanta ships the next minimum.

Companion to `2026-04-21-api-redesign-scale.md` (design decisions and
rationale). This file is the execution sequence.

## Initiative 1: Fix linearRate's input-unit bug

**Quanta commit 1** — `scale.go` + `scale_test.go`. Ship only:

- `type Scale struct{…}`
- `NewLinearScale(from, to Unit, factor string) (Scale, error)` + `MustNewLinearScale`
- `Apply(Measure) (Measure, error)` — input-unit-checked
- `From() Unit`, `To() Unit`

No `Then`, no `Identity`, no `RatioOf`, no `FactorString`, no `String` —
add them when a caller asks. Tag, publish.

**Meters-credits commit A** — `rating/rate_impl.go`:

- Replace `linearRate{factor, unit}` with `linearRate{scale quanta.Scale}`
- `convert` becomes `return r.scale.Apply(u)`
- `outputUnit()` returns `r.scale.To()`
- Factory in `newRateFromSpec` constructs via `quanta.NewLinearScale`

Done. Input-unit bug fixed. No other rate types touched yet.

## Initiative 2: graduatedRate

**Quanta commit 2** — `Measure.Slice(bounds []Measure) ([]Measure, error)`
only. Commit boundary semantics in docstring + tests. Tag.

**Meters-credits commit B** — `graduatedRate.convert`:

- Replace the tier-walk bookkeeping with `u.Slice(r.bounds)` + per-slice
  `Scale.Apply` + `Measure.Add` accumulator (`quanta.MustNewMeasure(r.outputUnit, "0")`
  as zero — don't need a `Zero(unit)` helper yet)
- Delete `prevBound`, zero accumulator, raw-Decimal re-wrap

## Initiative 3: volumeRate

**Quanta commit 3** — `Measure.Bucket(bounds []Measure) (int, error)`
only. Tag.

**Meters-credits commit C** — `volumeRate.convert`:

- `i, err := u.Bucket(r.bounds); return r.scales[i].Apply(u)`

## Initiative 4: remove Measure.Mul/Div

**Meters-credits commits D1–D9** — one per `.Mul()/.Div()` call site
(9 total across 5 files). Each site gets either:

- Migrated to `quanta.Scale.Apply` (if Scale-shaped)
- Made explicit: `quanta.MustNewMeasure(u, m.Quantity().Mul(d).String())`
  (ugly on purpose)

**Quanta commit 4** — delete `Measure.Mul(Decimal)`, `Measure.Div(Decimal)`,
and their tests. Tag.

## Everything else is demand-driven

Ship each of these only when a specific meters-credits commit proves the
need:

- `Scale.Then(Scale) Scale` — when a chain use case appears in rating
- `IdentityScale(unit) Scale` — when a default/no-op is awkward to express
- `Measure.RatioOf(Measure) Scale` — when a call site builds a ratio by hand
- `Zero(unit) Measure` — when `MustNewMeasure(u, "0")` becomes noise
- `Scale.FactorString() string`, `Scale.String()` — when logging/display
  needs it

Each is a one-method quanta commit + immediate consumer migration.

## Initiative 5: Decimal hiding

Still a decision gate, not a plan. Revisit after Initiative 4 merges; the
remaining Decimal footprint in meters-credits is what decides Option A
(full hide, months of migration) vs Option B (add string constructors +
query methods, leave Decimal exported as power-user escape hatch).

## Cadence summary

4 mandatory quanta releases (Scale, Slice, Bucket, Mul/Div removal) + N
demand-driven micro-releases. Each mandatory release is a single feature.
Each feature unblocks exactly one meters-credits file change. No bundled
quanta PRs, no "prep for future consumer" code.
