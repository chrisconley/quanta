# Analysis: motion-canvas/motion-canvas

- **Repo**: motion-canvas/motion-canvas (~17.9k stars)
- **Type**: TypeScript library + editor for programming informative vector animations using generators, intended for voice-over-synced technical explainers.
- **Length**: ~4.2 KB. Short; contributor-facing more than user-facing.

## Sections

1. **Centered header** — dark-mode-aware `<picture>` element with `prefers-color-scheme: dark` source swap for the logo (180px, motioncanvas.github.io-hosted SVGs).
2. **Four-badge row** — "published with lerna", "powered by vite", npm version of `@motion-canvas/core`, Discord shield.
3. **# Motion Canvas** — H1, then a two-bullet definition: "Motion Canvas is two things: a TypeScript library that uses generators to program animations. An editor providing a real-time preview of said animations." Plus a follow-up paragraph characterizing it as "a specialized tool designed to create informative vector animations and synchronize them with voice-overs."
4. **Using Motion Canvas** — a single sentence redirecting users to the hosted getting-started docs; no inline quickstart.
5. **Developing Motion Canvas locally** — monorepo package table (name + description) listing eleven packages: `2d`, `core`, `create`, `docs`, `e2e`, `examples`, `internal`, `player`, `template`, `ui`, `vite-plugin`. Then installation instructions: `npm install` and `npx lerna run build`.
6. **Developing Editor** (subsection) — `npm run template:dev` command explanation ("starts a vite server that watches the `core`, `2d`, `ui`, and `vite-plugin` packages").
7. **Developing Player** (subsection) — two-step sequence: `npm run template:build` then `npm run player:dev`.
8. **Installing a local version of Motion Canvas in a project** — detailed walkthrough for linking the monorepo into a sibling project: ASCII tree of the expected directory layout, diff-style `package.json` snippet showing the `file:` protocol swap (`"@motion-canvas/core": "file:../motion-canvas/packages/core"`), and a `vite.config.ts` snippet with `server.fs.strict: false`. Explains the `/@fs/` prefix requirement for external file loading.
9. **Contributing** — one-sentence pointer to `CONTRIBUTING.md`.
10. **Link definitions** — markdown reference-style link definitions for `authenticate`, `template`, `discord`, `docs` at the bottom.

## Tone & style

- **Register**: Contributor-technical, quiet, confident; no marketing-speak, no emoji in headings, no "awesome features" list.
- **Voice**: Third-person descriptive for the project pitch ("Motion Canvas is two things"); second-person imperative for setup ("After cloning the repo, run..."); first-person-plural briefly in "Check out our getting started guide."
- **Formatting**: `<picture>` dark-mode logo swap, monospace package name/description table with pipe-delimited columns, one diff-fenced code block (`diff` language) for `package.json` changes, one TypeScript snippet for `vite.config.ts`, reference-style link definitions collected at the end.

## Notable conventions

- **README is primarily a contributor manual, not a user landing page** — the sole user-facing section is one sentence redirecting to external docs; everything else concerns monorepo layout, dev scripts, and linking local forks.
- **Dark-mode `<picture>` tag with `prefers-color-scheme: dark`** — a small but deliberate polish signal; many large OSS READMEs still ship a single light-mode logo.
- **`file:` protocol linking walkthrough** — unusually thorough instructions for running a local fork inside a downstream animation project, including the specific vite `server.fs.strict: false` workaround and an explanation of *why* it's required (editor styles loaded via `/@fs/` prefix). Anticipates a real friction point ("when you want to use your own fork with some custom-made features").
- **Monorepo package table with one-line descriptions** — replaces the usual "architecture overview" ASCII art; lerna + vite badges at the top reinforce that this is the technology signal, not a feature pitch.
- **Reference-style link definitions at the bottom** (`[docs]: https://...`, `[discord]: https://...`) — keeps the prose clean; most OSS READMEs inline every URL.
- **No license section, no contributors grid, no stargazer chart, no sponsors** — strikingly minimal "trust" scaffolding; the project assumes readers arrived from the docs site or a YouTube video and already know what it is.
- **"Motion Canvas is two things"** is an unusually confident opening — binary decomposition rather than a single elevator pitch.

## Takeaway pattern

This is the "docs site does the selling, README does the plumbing" archetype. The project's marketing surface lives entirely at motioncanvas.io — the product is highly visual (vector animations) and a README screenshot would trivialize it — so the README abandons the landing-page role entirely and becomes a monorepo operator's manual. The most substantive section is the `file:`-linking walkthrough, which is the exact kind of guidance you'd expect to land on after a frustrated fork attempt; treating it as first-class README content respects the maintainer's likely assessment of which questions they get asked most. Lesson: when your product is inherently visual and lives on an external doc site, the README can shed the landing-page conventions entirely and optimize for contributor/fork workflows instead — the 18k stargazers clearly aren't arriving through the README, so don't duplicate the docs site there.
