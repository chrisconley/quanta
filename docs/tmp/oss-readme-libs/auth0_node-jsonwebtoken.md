# Analysis: auth0/node-jsonwebtoken

- **Repo**: auth0/node-jsonwebtoken (~18.2k stars)
- **Type**: Node.js implementation of JSON Web Tokens (RFC 7519); the de facto JWT library for Express/Node ecosystems.
- **Length**: ~15 KB. Long — essentially a complete API reference inlined into the README.

## Sections

1. **H1 + two-cell status table** — `# jsonwebtoken` followed by a two-column markdown table with Travis CI and david-dm Dependency badges (both now-dead services, unrevised).
2. **One-paragraph intro** — "An implementation of JSON Web Tokens" linking RFC 7519, noting `draft-ietf-oauth-json-web-token-08` as the developed-against draft and `node-jws` as the underlying dependency.
3. **Install** — single `$ npm install jsonwebtoken` fenced block.
4. **Migration notes** — two bullets linking wiki migration pages (v7→v8, v8→v9); no inline breaking-change summary.
5. **Usage / jwt.sign** — full signature `jwt.sign(payload, secretOrPrivateKey, [options, callback])`, split into sync vs async behavior, then each parameter explained in prose, then an `options:` bullet list with ~12 entries (`algorithm`, `expiresIn`, `notBefore`, `audience`, `issuer`, `jwtid`, `subject`, `noTimestamp`, `header`, `keyid`, `mutatePayload`, `allowInsecureKeySizes`, `allowInvalidAsymmetricKeyTypes`), multiple blockquote caveats, and 4 running JS examples (HMAC sync, RSA sync, RSA async, backdating).
6. **Token Expiration (exp claim)** — H4 aside quoting RFC's NumericDate definition verbatim in a long blockquote, then three equivalent code blocks showing manual `exp` vs `expiresIn: 60*60` vs `expiresIn: '1h'`.
7. **jwt.verify** — parallel structure to `sign`: full signature, sync/async explanation, parameter prose, ~12-item `options` list (`algorithms`, `audience`, `complete`, `issuer`, `jwtid`, `ignoreExpiration`, `subject`, `clockTolerance`, `maxAge`, `clockTimestamp`, `nonce`, `allowInvalidAsymmetricKeyTypes`), then a single ~50-line fenced block showing 9 back-to-back verify scenarios (symmetric sync/async, invalid token, asymmetric, audience/issuer/jwtid/subject mismatches, alg mismatch, JWKS via `jwks-rsa`).
8. **jwt.decode (inside `<details>`)** — the unsafe peek API is hidden behind a collapsible `<details><summary>Need to peek into a JWT without verifying it? (Click to expand)</summary>` with prominent "__Warning:__ This will __not__ verify" callouts.
9. **Errors & Codes** — three named error classes (`TokenExpiredError`, `JsonWebTokenError`, `NotBeforeError`) each with an object-shape bullet list (`name`, `message`, extra fields) and a sample callback block; the `JsonWebTokenError` message bullet nests 8 possible error strings.
10. **Algorithms supported** — 13-row markdown table mapping alg parameter value (HS256…ES512, plus `none`) to digital-signature algorithm description, with parenthetical node version gates for PS*.
11. **Refreshing JWTs** — editorial caveat section explicitly declining to include refresh as a feature ("We are not comfortable including this as part of the library") and linking a gist + issue + PR for users who still want it.
12. **TODO / Issue Reporting / Author / License** — four terminal one-liners; TODO is a single bullet ("X.509 certificate chain is not checked"), Issue Reporting directs security reports to Auth0's Responsible Disclosure Program.

## Tone & style

- **Register**: RFC-adjacent and security-conscious — the blockquoted NumericDate definition, verbatim draft reference, and repeated `__Warning:__` callouts position this as a spec-implementing library rather than a convenience helper.
- **Voice**: Second-person instructional inside options (`"if you want to check audience..."`), first-person plural for editorial stances (`"We are not comfortable including this"`, `"we recommend you to think carefully"`).
- **Formatting**: Dense bullet lists for options with nested blockquotes for caveats/examples, `__bold__` underscore style (not `**`) for warnings, `<details>`/`<summary>` to gate the unsafe `decode` API, and markdown tables for both the top badges and the algorithms reference.

## Notable conventions

- **Hides `jwt.decode` in a `<details>` collapsible** — a deliberate UX nudge: `decode` is footgun-shaped (skips signature verification), so the README makes you click past a warning-laden summary to reach it. Almost no other JWT docs treat decode this way.
- **Per-option caveats as nested blockquotes** — each option that accepts a time-span repeats the identical blockquote about vercel/ms parsing (`> Eg: 60, "2 days"...`). The duplication is intentional — each option is a potential lookup target, and the caveat must live where the reader lands.
- **"Refreshing JWTs" as a non-feature section** — most libraries either ship refresh or stay silent; this README explicitly documents why it refuses to, pointing at a gist and an issue thread. The stance is itself the documentation.
- **Error messages enumerated as a bullet list inside the `message` field** — the `JsonWebTokenError.message` bullet has 8 nested sub-bullets of exact string values (`'jwt malformed'`, `'invalid signature'`, etc.); this doubles as a grep target for users matching on error text.
- **Dead-service badges left in place** — the top table still shows Travis CI and david-dm; they're both defunct, but the README hasn't been cleaned up, which is mildly revealing about maintenance pace.
- **"developed against `draft-ietf-oauth-json-web-token-08`"** — naming the exact IETF draft revision is a posture move: this is a library that cares about spec provenance, not just "JWT support."

## Takeaway pattern

This is the "API reference inlined as README" archetype: no separate docs site, no external reference page — every option, every error class, every algorithm is documented in-repo in README.md. It works for node-jsonwebtoken because the API surface is just three functions (`sign`, `verify`, `decode`) and because users typically arrive via npm, not via a marketing site. Lesson: when your library is a pure primitive with a small number of entry points and a security-critical failure mode, collapsing the docs into the README — with prominent warnings at the exact lookup points (options, error names, unsafe APIs) — serves readers better than a separate docs site they'd skip.
