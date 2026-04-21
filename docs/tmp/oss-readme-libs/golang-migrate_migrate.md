# Analysis: golang-migrate/migrate

- **Repo**: golang-migrate/migrate (~17.9k stars)
- **Type**: Go database migration CLI and library; reads ordered up/down SQL files and applies them across many databases.
- **Length**: ~5.6 KB. Short, dense, and structured as a reference index into the repo's subdirectories.

## Sections

1. **Badge row (no header image)** — eight shields inline at the top: GitHub Actions status, GoDoc, Coveralls coverage, packagecloud.io debs, Docker Hub pulls, supported Go versions (1.24/1.25), GitHub release, Go Report Card.
2. **# migrate** — H1, then a bold one-liner: "Database migrations written in Go. Use as CLI or import as library." Followed by three bullets stating design philosophy ("Drivers are 'dumb'," "Database drivers don't assume things or try to correct user input. When in doubt, fail.").
3. **Fork attribution** — one line: "Forked from mattes/migrate."
4. **Databases** — 20+ bulleted database driver links, each pointing to a `database/<name>/` subdirectory. Includes Postgres variants (PGX v4/v5, Redshift), enterprise DBs (MS SQL, Spanner, CockroachDB), NoSQL (MongoDB, Neo4j, Cassandra), and even `Shell`. Three entries are flagged `[todo #165]` etc., linking to the original mattes repo's issues.
5. **Database URLs** — URL syntax primer with an explicit list of reserved characters that must be percent-encoded, plus two Python one-liners (python2 and python3) for URL-encoding a sample password.
6. **Migration Sources** — 10 bulleted source drivers (filesystem, io/fs, go-bindata, pkger, GitHub/GitHub EE, Bitbucket, Gitlab, S3, GCS).
7. **CLI usage** — three-bullet design statement, including "No config search paths, no config files, no magic ENV var injections." One basic usage example and a Docker example.
8. **Use in your Go project** — API-stability bullets ("API is stable and frozen," "Thread-safe and no goroutine leaks," "Bring your own logger"). Two Go code samples: `migrate.New` with a GitHub source, and `NewWithDatabaseInstance` wrapping an existing `*sql.DB`.
9. **Getting started** — single link to `GETTING_STARTED.md`.
10. **Tutorials** — two links (CockroachDB, PostgreSQL) plus "(more tutorials to come)".
11. **Migration files** — filename convention example (`1481574547_create_users_table.up.sql`) with a "Why two files?" link to the FAQ.
12. **Coming from another db migration tool?** — pointer to third-party `migradaptor` with an explicit "not affiliated or supported" disclaimer.
13. **Versions table** — markdown table with check/cross emoji for master, v4 (supported), and v3 ("DO NOT USE").
14. **Development and Contributing** — short paragraph, Makefile and CONTRIBUTING.md links, FAQ link.
15. **Horizontal rule + awesome-go pointer** — "Looking for alternatives?" linking to awesome-go's database section.

## Tone & style

- **Register**: Engineering-terse; stated design principles ("When in doubt, fail.", "No magic ENV var injections.") treated as selling points.
- **Voice**: Third-person descriptive for design principles; imperative for instructions.
- **Formatting**: All-shields header with no logo or tagline image; bulleted driver catalogs rather than prose; two minimal Go code blocks; one versions table with `:white_check_mark:` / `:x:` emoji; zero emoji otherwise.

## Notable conventions

- **Design philosophy as bullet points, not a "Why migrate?" paragraph** — "Drivers are 'dumb'," "Fail when in doubt," and "No magic ENV var injections" function as anti-features, directly contrasting with competing migration tools.
- **Fork attribution surfaced at the top**, not buried in credits — and `[todo #XXX]` references point back to the upstream `mattes/migrate` issue tracker, acknowledging unfinished inherited work.
- **Python URL-encoding snippets for both python2 and python3** — the python2 line is an anachronism by 2026, but preserved rather than cleaned up; suggests a "stable-release-only" editing discipline.
- **Explicit "DO NOT USE" row for v3** with bolded shout and two alternative import paths shown — unusually blunt version-support signaling.
- **Third-party adapter disclaimer** — "migradaptor is not affiliated or supported by this project" is an uncommonly explicit warning for a linked sibling tool; reflects the project's firm support-boundary posture.
- **Exit link to awesome-go alternatives** at the very bottom — confident enough in their product to send visitors to shop the competition.

## Takeaway pattern

This is the "Go stdlib-style reference README": a terse, bulleted, tabular directory of drivers and invariants, expecting readers already know what a migration tool is and just need to confirm driver coverage, import path, and version support. It works because the domain is mature and the audience is self-selecting — nobody lands on a Go DB migration README curious; they land ready to integrate. The tone borrows from Go's own documentation culture (small surface, explicit failure modes, no magic) and the README doubles as a support-scope contract: these databases, these sources, v4 only, shell out to awesome-go if that's not enough. Lesson: for infrastructure libraries, a catalog README with explicit non-features ("no config files, no magic ENV injections") telegraphs competence better than a marketing pitch — users want to know what it won't do as much as what it will.
