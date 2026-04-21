# Analysis: dropzone/dropzone

- **Repo**: dropzone/dropzone (~18.2k stars)
- **Type**: Vanilla-JS drag-and-drop file upload library that turns any HTML element into a styled upload target with previews, progress, and XHR handling.
- **Length**: ~2.6 KB. Very short; landing-page-lite with a quickstart and a features list.

## Sections

1. **Logo image** — top-of-file SVG banner served from the repo's `assets` branch; no centering wrapper (sits flush left).
2. **Single CI badge** — "Test and Release" GitHub Actions badge; no npm version, no license, no downloads badges.
3. **Elevator paragraph** — two short paragraphs: "Dropzone is a JavaScript library that turns any HTML element into a dropzone…" plus "It's fully configurable, can be styled according to your needs and is trusted by thousands."
4. **Centered screenshot** — user-uploaded GitHub asset showing the rendered dropzone UI.
5. **Quickstart** — `npm install --save dropzone` / `yarn add dropzone` install commands, then ES6 and CommonJS usage snippets showing the `new Dropzone("div#myId", { url: "/file/post" })` two-line minimum viable example.
6. **"Checkout our example implementations"** — bullet-style emoji link (👉) to the `dropzone-examples` sibling repo.
7. **"Not using a package manager or bundler?"** — standalone-tag instructions using unpkg CDN URLs for `dropzone.min.js` and `dropzone.min.css`, then an inline `<div class="my-dropzone">` plus initialization script.
8. **Horizontal rule, then two link bullets** — `📚 Full documentation` (docs.dropzone.dev) and `⚙️ src/options.js` link for "all available options."
9. **Blockquote IE/6.0.0-beta notice** — ⚠️ warning that IE support is being dropped in the next major; advises opting into `6.0.0-beta.1` with pinned version.
10. **Community** — directs support traffic to GitHub Discussions or StackOverflow (`dropzone.js` tag), explicitly **not** to the issues tracker unless it's a bug. Second blockquote reminder to read CONTRIBUTING.md.
11. **Main features ✅** — 10 bulleted features including chunked uploads, S3 multipart, browser image resizing, high-DPI support, "Beautiful by default," "Well tested."
12. **# MIT License** — one-line link to the LICENSE file, rendered as an H1 (inconsistent with prior H2 section levels).

## Tone & style

- **Register**: Plain, matter-of-fact, faintly modest ("trusted by thousands," "Well tested"); no marketing hype, no emoji headers except the features-checkmark and sparing inline emoji.
- **Voice**: Descriptive third-person for the library pitch; second-person imperative in the quickstart.
- **Formatting**: Logo + single badge (unusually sparse), two code-block styles (npm + JS for bundler users, HTML + script for vanilla users), two ⚠️ blockquote notices, horizontal-rule dividers breaking the "install → docs → warning" flow.

## Notable conventions

- **Only one badge** — just the CI status. No npm version, no license shield, no download count, no bundle size; a deliberate austerity choice for a library with 18k stars that could easily sport a dozen.
- **H1 for "MIT License" at the end** disrupts the H2-based heading hierarchy — a legacy formatting quirk preserved rather than normalized.
- **Support-channel triage is bolded in-prose** — "use the discussions section or stackoverflow ... and **not** the GitHub issues tracker" — rare explicitness about routing users away from issues, with a second bolded contributing-guide reminder in the same section.
- **Two target audiences split into separate sections** — bundler users (Quickstart with ES6/CJS) and script-tag users (unpkg CDN); many modern READMEs drop the second audience entirely, but Dropzone still serves the "paste into an HTML file" crowd.
- **Beta opt-in buried in a blockquote** rather than promoted to a header — "you can already opt into `6.0.0-beta.1`" is almost apologetic; the warning is framed as a heads-up, not a migration push.
- **Link to `src/options.js` as API reference** — treats the source file as canonical docs surface, not a secondary artifact.

## Takeaway pattern

This is the "mature maintenance-mode library" README: the project is old (pre-ES6-module era, IE still being phased out in 2026), widely deployed, and the README's job is to get integrators copy-pasting within 30 seconds and route everyone else to Discussions so the maintainer isn't drowning in issues. Sparse badging and a single CI shield signal confidence ("you already know what this is"); the two installation paths (bundler + script tag) preserve backward compatibility with the older audience the project acquired years ago. Lesson: legacy-but-loved libraries earn trust through quiet austerity — no hype, firm support-channel triage, and an honest "trusted by thousands" without inflating the pitch. The implicit message is that you're joining a crowd that already settled on this choice.
