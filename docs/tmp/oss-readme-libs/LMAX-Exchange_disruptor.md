# Analysis: LMAX-Exchange/disruptor

- **Repo**: LMAX-Exchange/disruptor (~17.7k stars)
- **Type**: Java high-performance inter-thread messaging / ring-buffer library (the canonical LMAX Disruptor pattern implementation).
- **Length**: ~0.8 KB. Minuscule — arguably the shortest README in this cohort.

## Sections

1. **H1 title** — `= LMAX Disruptor` in AsciiDoc syntax (not Markdown).
2. **Badge row** — three `image:` AsciiDoc macros: Java CI with Gradle, CodeQL, and License.
3. **One-line tagline** — "A High Performance Inter-Thread Messaging Library."
4. **Maintainer** — single line: "LMAX Development Team."
5. **Support** — two bullets: GitHub issue tracker and a `groups.google.com/group/lmax-disruptor` Google Group.
6. **Documentation** — five bulleted links to the GitHub Pages site: Overview, User Guide, API Documentation (module-summary Javadoc), Developer Guide, and a wiki FAQ.

That is the entire file.

## Tone & style

- **Register**: Institutional-terse. No marketing language, no feature claims beyond the one-line tagline, no emoji, no screenshots.
- **Voice**: Impersonal. No "we," no "you" — just labels and links.
- **Formatting**: AsciiDoc (`.md` extension notwithstanding) — `=` for H1, `==` for H2, `image:url[alt,link=url]` for badges, space-prefixed bullets under Support. GitHub renders AsciiDoc natively so the file displays correctly despite the filename.

## Notable conventions

- **File is `.md` but content is AsciiDoc** — a mismatch that suggests either an early AsciiDoc migration that forgot the rename or deliberate use of GitHub's AsciiDoc renderer under a Markdown filename. Either way, copy-paste readers may be confused by `= Title` and `image:...[]`.
- **No "What is it?" beyond the tagline** — the README does not explain the ring-buffer pattern, does not show a code sample, does not name competitors or use cases. It assumes readers already know the Disruptor pattern (and most visitors do — the pattern is famous from the 2011 LMAX paper).
- **"Maintainer: LMAX Development Team"** — naming the sponsoring company rather than individuals is unusual for OSS; signals corporate stewardship and that individual maintainers aren't the contact surface.
- **No installation instructions** — no Maven/Gradle coordinate snippet, no "add this dependency" block. Readers must click through to the user guide to learn the groupId/artifactId, which is surprising for a project this popular.
- **No code of conduct, no CONTRIBUTING link, no license section** — beyond the license badge in the header. License text lives in `LICENCE.txt` (British spelling, unlinked from the README body).
- **No quickstart, no architecture diagram** — everything is deferred to external docs. The README is pure link-rail.
- **Google Group over Slack/Discord** — signals a pre-2015 community-tooling era that hasn't been updated; the project is in low-churn maintenance mode and its established user base finds it through search, not onboarding UX.

## Takeaway pattern

This is the "famous-enough-to-not-bother" README archetype: a project whose reputation precedes it so thoroughly that the maintainers feel no need to pitch, teach, or demo in-repo. It works, minimally, because (a) anyone typing "disruptor" into GitHub already knows what the LMAX Disruptor is from the 2011 paper / InfoQ talk, and (b) the linked User Guide and Developer Guide are genuinely well-produced. It fails as a top-of-funnel page for anyone who arrived by accident or by pattern-matching "inter-thread messaging" — no code sample, no Maven coordinates, no "why would I use this over `java.util.concurrent`?" Lesson: a README this spare only works when the project's fame does the pitching for it; if you're tempted to write this minimally, check whether your project has a Martin Fowler blog post explaining it first.
