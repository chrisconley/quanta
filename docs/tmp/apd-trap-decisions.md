# apd Trap Set & Special-Value Inputs — Open Questions

Scratch notes from the session that consolidated `newCalcContext` and added
`TestDecimal_Traps`. Everything here is a **deferred decision** — the current
behavior is pinned by tests; the questions are whether to change it.

## Where we landed

- Consolidated four internal `apd.Context` constructions into one helper:
  `newCalcContext(rounding apd.Rounder)` in `decimal.go:16`.
  Built on `apd.BaseContext.WithPrecision(34)` (idiomatic) with Rounding and
  Traps overridden to preserve the historic narrow set:
  `InvalidOperation | DivisionByZero | Overflow`.
- `TestDecimal_Traps` pins behavior for every reachable apd Condition
  (8 subtests — 5 errors, 3 silent).
- `TestMeasure_TrapsPropagate` smoke-tests the wrapper error chain.
- Benchmark (`benchmark_test.go`) stays inline on purpose (external package;
  exposing the helper would leak apd in the public API) but mirrors the new
  pattern for honest overhead comparison.

The three probes in `/tmp/apd_probe` used during investigation were deleted;
if re-needed, the apd source is at
`~/go/pkg/mod/github.com/cockroachdb/apd/v3@v3.2.3/` — `condition.go` has the
taxonomy and `GoError` logic (the hard-coded `SystemOverflow|SystemUnderflow`
branch that always errors regardless of traps is the key find).

## Open question 1: Narrow traps vs `apd.DefaultTraps`

**Current:** `InvalidOperation | DivisionByZero | Overflow` (3 conditions).

**`apd.DefaultTraps`:** `SystemOverflow | SystemUnderflow | Overflow |
Underflow | Subnormal | DivisionUndefined | DivisionByZero |
DivisionImpossible | InvalidOperation` (9 conditions).

### What flips if we adopt `DefaultTraps`

The only subtest in `TestDecimal_Traps` that would change behavior under
`DefaultTraps` in a way reachable through quanta's public API is:

- **`DivisionUndefined silently returns NaN on 0/0`** → would become an
  error case. That's the known gap. Fix is a one-line test update plus the
  trap-set change.

Other conditions in `DefaultTraps` that are NOT currently reachable through
quanta's public API:
- `Underflow` (standalone, not `SystemUnderflow`) — at our magnitudes this
  always co-fires with `SystemUnderflow`, which errors regardless of traps.
  So trapping standalone `Underflow` adds no observable behavior.
- `Subnormal` — same; absorbed into `SystemUnderflow` at our magnitudes.
- `DivisionImpossible` — only fires via `QuoInteger`, which quanta doesn't
  expose.

### What flips if we *narrow further*

Removing `InvalidOperation` from the trap set would change:
- sNaN arithmetic → silent NaN (currently errors)
- `Infinity - Infinity` → silent NaN (currently errors)

Removing `Overflow` would not change observable behavior — `SystemOverflow`
is hard-coded to error on the same inputs. So `Overflow` in our trap set is
defensive for a hypothetical future apd version that changes that coupling,
not load-bearing today.

### Arguments for `DefaultTraps`

- Only one surface change through the public API (`0/0`), and that's a real
  footgun (silent NaN from what looks like a normal division).
- Matches the apd maintainer's recommended default for well-behaved
  numerical code.
- The "if I don't understand the condition, I want to know" posture aligns
  with a billing-grade library. Underflow/Subnormal being defensive-only
  today doesn't mean they'll stay that way if we add new operations (e.g.
  sqrt, pow, exp).

### Arguments for the narrow set

- Existing tests (including the 8 characterization subtests) currently
  pass. Changing the trap set is a behavior change, however small.
- `0/0` returning `NaN` is arguably IEEE-correct and matches float64 Go
  behavior. Some callers may rely on the NaN-propagation semantics.

### Recommendation

Adopt `DefaultTraps`. One test update, no observable breakage through the
current API, genuine footgun closed. But this is a judgment call; either
decision is defensible now that the behavior is pinned.

## Open question 2: Should `NewDecimal` accept special-value strings?

`quanta.NewDecimal` currently accepts (verified):
- `"NaN"` → quiet NaN (propagates silently through arithmetic)
- `"sNaN"` → signaling NaN (triggers InvalidOperation trap)
- `"Infinity"`, `"-Infinity"`, `"Inf"` → infinity values
- Likely also `"snan"`, case variants, etc. (apd's SetString is permissive)

### Why this matters

- Quiet NaN is the ONE condition that silently corrupts results without
  any trap catching it. Widening traps doesn't close it. The only way to
  prevent `MustNewDecimal("NaN").Add(x)` from producing silent garbage
  is to reject `"NaN"` at construction time.
- `Infinity - Infinity` returning an InvalidOperation error is arguably
  weird: the caller typed `"Infinity"` into a *billing* library, which
  suggests they made a mistake much earlier. Hard-to-compose stacktrace
  locations.
- This is an input-validation question, orthogonal to the trap set.

### Options

1. **Status quo.** Accept everything apd accepts. Document the footgun.
2. **Reject non-finite inputs in `NewDecimal`.** Check `apd.Decimal.Form`
   after `SetString`; return an error if it's not `apd.Finite`. This is
   one-liner and has zero effect on legitimate callers — no billing system
   should ever be constructing Decimals from the string `"NaN"`.
3. **Reject in `NewDecimal` but expose a `NewSpecialDecimal` escape hatch**
   for callers who genuinely need non-finite values (no known use case
   today, so skip unless someone asks).

### Recommendation

Option 2. `NewDecimal` is the entry point for a billing-grade library —
accepting NaN from string input is a default that serves no one. Making it
an error closes the silent-NaN hole that no trap set can close. Before
doing this: grep the repo + workspace for known callers, confirm no test
or production code intentionally passes non-finite strings.

## Open question 3: Error-text assertions are brittle

`TestDecimal_Traps` asserts on apd's error strings:
- `"division by zero"`, `"exponent out of range"`, `"invalid operation"`

These are stable in apd v3.2.3 but not part of apd's public API contract.
An apd upgrade could silently change them and we'd get test failures that
look scary but aren't actual behavior changes.

### Options

1. **Status quo.** Pin the strings; accept that apd upgrades may require
   a mechanical test update.
2. **Assert on `apd.Condition` bits instead of error strings.** Would
   require `Decimal.Add/Sub/Mul/Div` to return a typed error that preserves
   the Condition — a non-trivial API change to solve a theoretical problem.
3. **Use `errors.Is` against sentinel error values.** apd doesn't currently
   export its trap errors as sentinels; would need upstream contribution
   or wrapping in quanta.

### Recommendation

Status quo. apd's error text has been stable across versions and the
cost of a mechanical update during upgrade is low. Not worth an API change.

## Open question 4: `Quantized.Decimal` panics on trap

`Quantized.Decimal()` at `quantized.go:143` calls `Decimal.Mul(multiplier,
quantum)` and panics on error. Unreachable through normal construction
(`DeprecatedNewQuantizedFromMultiplier` takes int64 × quantum, both always
in-range), so the panic is defensive.

But if someone constructed a pathological `Quantized` (e.g. `multiplier =
math.MaxInt64`, `quantum = "1E99999"`), `.Decimal()` would panic. The
`Deprecated` prefix suggests this constructor is slated for removal, at
which point the panic path becomes truly unreachable and the defensive
`panic` can be simplified away.

**Action:** nothing until the `Deprecated*` removal happens. Note here so
it's on the cleanup list.

## Out-of-scope follow-ups surfaced during this work

- **README concurrency claim.** Was the original reason this thread
  started. Now defended by `concurrency_test.go` and `TestGotcha_Shared
  ApdDecimalAcrossGoroutines`. README still doesn't explicitly claim
  concurrency safety — worth adding a short line once the trap-set
  decision lands (since that affects what "safe" means at the edges).
- **Other README improvements** deferred from the earlier "what else
  should we address?" turn. The list wasn't exhausted; concurrency was
  item #3 of several. Re-open when ready.
