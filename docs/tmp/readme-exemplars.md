# README exemplars for quanta / metron

Research across 5 angles: finance/billing, spec+reference-impl, narrow-scope discipline, exceptional READMEs, and typed-domain-value libraries. 40+ exemplars, deduplicated below.

## A. Finance / billing / ledger / decimal (most directly adjacent)

1. **shopspring/decimal** (Go) — https://github.com/shopspring/decimal — "Alternative libraries" section honestly sends users away; "Known Issues" section is a maturity signal.
2. **cockroachdb/apd** (Go) — https://github.com/cockroachdb/apd — Standards-anchoring (IEEE 754 / GDA spec); condition-flags/traps model reports *how* a computation happened.
3. **paupino/rust-decimal** (Rust) — https://github.com/paupino/rust-decimal — Bit-layout diagram; feature-flag matrix as scope fence.
4. **Rhymond/go-money** (Go) — https://github.com/Rhymond/go-money — Floating-point horror-story opener; `ErrCurrencyMismatch` pattern.
5. **dinerojs/dinero.js** (TS) — https://github.com/dinerojs/dinero.js — "When should I use Dinero?" decision section with pros/cons.
6. **JodaOrg/joda-money** (Java) — https://github.com/JodaOrg/joda-money — Aggressive brevity + explicit non-goals paragraph.
7. **beancount/beancount** (Python) — https://github.com/beancount/beancount — Shows the file format in the README itself.
8. **simonmichael/hledger** (Haskell) — https://github.com/simonmichael/hledger — Comparison matrix vs ledger-cli and beancount.
9. **formancehq/ledger** (Go) — https://github.com/formancehq/ledger — Concrete use-case enumeration (marketplaces, wallets, split payments).
10. **getlago/lago** (Go/Ruby) — https://github.com/getlago/lago — Counter-example: platform-README template to *avoid*.
11. **killbill/killbill** (Java) — https://github.com/killbill/killbill — "Production users" section as trust signal.
12. **govalues/decimal** (Go) — https://github.com/govalues/decimal — Head-to-head benchmark tables vs shopspring/apd.
13. **moov-io/ach** (Go) — https://github.com/moov-io/ach — Supported-features matrix as scope fence.
14. **tigerbeetle/tigerbeetle** (Zig) — https://github.com/tigerbeetle/tigerbeetle — Opinionated design non-negotiables section.

## B. Spec + reference-implementation pattern (relevant to metron)

15. **CloudEvents** — https://github.com/cloudevents/spec — Version compatibility matrix across SDKs.
16. **JSON Schema + santhosh-tekuri/jsonschema** — Conformance statement ("passes official test suite, draft X–Y").
17. **CommonMark + cmark** — https://github.com/commonmark/cmark — Uses the literal phrase "reference implementation" in sentence one; lists peer implementations.
18. **SPDX + spdx/tools-golang** — "Intended to make it easier for programs to work with X files" framing.
19. **OpenLineage** — https://github.com/OpenLineage/OpenLineage — Explicit "what (spec) vs how (integrations)" dichotomy.
20. **SLSA + slsa-verifier** — https://github.com/slsa-framework/slsa-verifier — "Limitations" section naming what's out of scope.
21. **in-toto/attestation** — https://github.com/in-toto/attestation — Separate-repo structure; spec repo contains only schemas.
22. **OCI Image Spec + go-containerregistry** — Table-of-contents-style spec README (each link = one concept).
23. **TOML + BurntSushi/toml** (Go) — https://github.com/BurntSushi/toml — Version-badge-first style declaring spec compatibility high in README.

## C. Narrow-scope discipline / "not for you, use X"

24. **dtolnay/semver** (Rust) — https://github.com/dtolnay/semver — Header literally is "Scope of this crate"; states the domain, names adjacent domains, redirects without apology.
25. **etcd-io/bbolt** (Go) — https://github.com/etcd-io/bbolt — Prose-per-competitor comparison (not just a table); "Caveats & Limitations" with ~15 specific failure modes.
26. **sirupsen/logrus** (Go) — https://github.com/sirupsen/logrus — "Logrus is in maintenance-mode… check out Zerolog, Zap, Apex" — the ultimate scope-honesty move.
27. **chronotope/chrono** (Rust) — https://github.com/chronotope/chrono — Numbered "Limitations" as first-class README section; points to Chrono-TZ for tz data.
28. **fabian-hiller/valibot** (TS) — https://github.com/fabian-hiller/valibot — Names the dominant alternative (Zod) and quantifies the difference; includes "Coming from Zod?" migration guide.
29. **ai/nanoid** (JS) — https://github.com/ai/nanoid — Opens with "Comparison with UUID" *before* features; quantified benchmarks vs named competitors.
30. **ulid/spec** — https://github.com/ulid/spec — Four-bullet "why existing approaches fall short" opener.
31. **golang-migrate/migrate** (Go) — https://github.com/golang-migrate/migrate — Explicit anti-feature bullets ("no config search paths, no magic ENV vars").
32. **python-attrs/attrs** (Python) — https://github.com/python-attrs/attrs — Names the stdlib alternative (`dataclasses`) and links to a separate COMPARISON.md.

## D. Exceptional README craft (small/medium libs)

33. **BurntSushi/ripgrep** (Rust) — https://github.com/BurntSushi/ripgrep — Reproducible benchmark tables with exact commands; "performance cliffs" table showing where rg *loses*.
34. **charmbracelet/bubbletea** (Go) — https://github.com/charmbracelet/bubbletea — Narrative tutorial README — you build a working app top-to-bottom.
35. **sindresorhus/*** (Node) — Rigid minimal template: description, Install, Usage, typed-option API block, Related. The discipline is the value.
36. **pocketbase/pocketbase** (Go) — https://github.com/pocketbase/pocketbase — `> [!WARNING]` pre-1.0 callout as GitHub alert block.
37. **hashicorp/raft** (Go) — https://github.com/hashicorp/raft — Deep-links to the authoritative paper; respects audience expertise.
38. **tokio-rs/tokio** (Rust) — https://github.com/tokio-rs/tokio — "Related Projects" ecosystem map (hyper, tonic, tracing, mio).
39. **zellij-org/zellij** (Rust) — https://github.com/zellij-org/zellij — Horizontal nav bar above prose; answers "what does this feel like?" first.
40. **antonmedv/fx** (Go) — https://github.com/antonmedv/fx — Radical minimalism: README is a GIF + one link; all docs offsite.
41. **miniflux/v2** (Go) — https://github.com/miniflux/v2 — Intent-based feature groupings, not flat lists.
42. **pydantic/pydantic** (Python) — https://github.com/pydantic/pydantic — Runnable 15-line example demonstrating value prop; visible error messages.

## E. Typed-domain-value libraries (Measure/Quantized-style)

43. **iliekturtles/uom** (Rust) — https://github.com/iliekturtles/uom — Mars Climate Orbiter hook; shows the compile error inline (`// error[E0308]: mismatched types`).
44. **Noda Time** (C#) — https://github.com/nodatime/nodatime — "Helps you think about your data more clearly" — positions types as a *thinking tool*.
45. **hgrecco/pint** (Python) — https://github.com/hgrecco/pint — Tagline inversion: "makes units easy" (not "catches unit errors").
46. **colinhacks/zod** (TS) — https://github.com/colinhacks/zod — "Parse, don't validate" as the frame.
47. **Effect-TS / @effect/schema** (TS) — https://github.com/Effect-TS/effect — Branded types idiom: a validated value can only be built through a parser.

---

## Cross-cutting patterns to consider cherry-picking

1. **Name the disaster.** uom has Mars Climate Orbiter. go-money has the floating-point penny bug. quanta's hook could be a named incident (Knight Capital, a public invoicing-rounding lawsuit, the Vancouver Stock Exchange index).
2. **Show the failure inline in examples.** Put `// err: ErrUnitMismatch…` in the code sample, not only in prose (uom's technique).
3. **"Scope of this library" as a real H2** (dtolnay/semver). Not a buried paragraph — a top-level heading.
4. **Prose-per-competitor, not a sparse table** (bbolt, shopspring). 2–4 sentences per alternative, each ending "pick X if…". quanta's current table is a starting point; consider prose form.
5. **Name the dominant alternatives by name** (valibot→Zod, attrs→dataclasses, logrus→Zerolog). Vague "other libraries" reads as defensive.
6. **"Why not just X?" FAQ.** Anticipate: why not `shopspring/decimal` alone? why not `int64` cents? why not `go-money`?
7. **Anti-feature bullets** (golang-migrate). "quanta does NOT: auto-round, infer locale, silently truncate, expose apd publicly, …" — trust signal in billing contexts.
8. **Pre-1.0 callout as `> [!WARNING]`** (PocketBase). GitHub renders it distinctively; the current inline italic is easy to miss.
9. **Conformance statement** (jsonschema, TOML). "Tested against X golden cases / passes suite Y."
10. **"Reference implementation" phrase, spelled plainly** (CommonMark). For metron especially.
11. **Standards anchor** (apd → IEEE 754). quanta already does this subtly; could be more prominent.
12. **"Think about your data more clearly"** framing (Noda Time). Measure-vs-Quantized as a *clarifying* split, not just a correctness guard.
13. **Parse-don't-validate framing for Quantize.** The `Measure → Quantize → Quantized` pipeline *is* a parsing boundary. Say so explicitly.
14. **"Related / See also" with honest redirects** (tokio). A section that sends readers elsewhere.
15. **Visible error messages in hello-world** (pydantic). Show what a failure looks like.
16. **Drop sections aggressively** (antonmedv/fx). Audit: does each section belong here or in GoDoc/CHANGELOG?

## Closest structural analogues to emulate

- **For quanta:** cockroachdb/apd (pre-1.0 Go decimal, spec-anchored, serious tone, explicit differences from shopspring).
- **For metron:** in-toto/attestation (spec-only repo + separate reference impl) or OpenLineage (explicit "what vs how" split with small client libs).
