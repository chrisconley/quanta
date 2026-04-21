# Analysis: liabru/matter-js

- **Repo**: liabru/matter-js (~18.2k stars)
- **Type**: JavaScript 2D rigid body physics engine for the web (Canvas/browser + Node.js).
- **Length**: ~7.5 KB. Medium-short; heavy on outbound links, light on inline code.

## Sections

1. **Logo + tagline + horizontal nav** — SVG logo image, one-line blockquote tagline ("*Matter.js* is a JavaScript 2D rigid body physics engine for the web"), a bold brm.io URL, then a single-line dotted nav (`Demos ・ Gallery ・ Features ・ Plugins ・ Install ・ Usage ・ Examples ・ Docs ・ Wiki ・ References ・ License`) using the Japanese `・` separator.
2. **Demos** — a three-column HTML `<table>` of `<ul>` lists with ~40 demo links pointing to hash anchors on brm.io/matter-js/demo (e.g., `#newtonsCradle`, `#wreckingBall`, `#slingshot`, `#cloth`, `#raycasting`).
3. **Gallery** — bulleted list of ~12 real-world sites using Matter.js (Patrick Heng, USELESS, Google's "Game of The Year", Phaser, Pablo The Flamingo, Glyphfinder, etc.), each formatted as `[Site Name](url) by Author/Studio`, ending with a `[more...]` wiki link.
4. **Features** — 24-item flat bullet list mixing physics capabilities ("Rigid bodies," "Restitution," "Sleeping and static bodies"), tooling ("MatterTools," "World state serialisation"), and a closing meta-bullet ("An original JavaScript physics implementation (not a port)").
5. **Install** — npm one-liner and indented `<script src="matter.js">` tag.
6. **Performance with other tools (e.g. Webpack, Vue etc.)** — a caveat section flagging that Webpack sourcemaps and Vue watchers degrade real-time perf, pointing to issue #1001.
7. **Usage** — no inline example; defers to the `Getting started`, `Running`, and `Rendering` wiki pages.
8. **Tutorials / Examples / Plugins** — each is 1–2 lines of outbound links (wiki, `examples/` dir, codepen collection, `matter-plugin-boilerplate`).
9. **Documentation** — single line pointing to `brm.io/matter-js/docs/` and the wiki.
10. **Building and Contributing** — `npm install` + `npm run dev` in indented code blocks, link to `CONTRIBUTING.md`.
11. **Changelog / References / License** — three trailing one-liner sections; License restates "absolutely no warranty" in plain prose under the MIT link.

## Tone & style

- **Register**: Neutral-reference and slightly proud — understated, but willing to boast where earned ("An original JavaScript physics implementation (not a port)", "Cross-browser and Node.js support (Chrome, Firefox, Safari, IE8+)").
- **Voice**: Third-person descriptive with occasional second-person imperative in install/build ("You can install using...", "To build you must first install node.js").
- **Formatting**: Raw HTML `<table>` for multi-column demo links, indented (four-space) code blocks rather than fenced, `###` headings throughout (no `##`), and the decorative `・` in the nav bar.

## Notable conventions

- **Gallery as social proof** — the "See how others are using matter.js physics" list of credited production sites is doing the job that badges, stars, or testimonials do in other READMEs; it signals "this has shipped in the wild" without any metrics.
- **Three-column HTML table for demos** — instead of a long single list or a grid of screenshots, the author uses an actual `<table>` element to balance ~40 demo names across columns; reads as hand-curated and ranked rather than alphabetized.
- **"Performance with other tools" as a top-level section** — unusual to dedicate a named section in the README to toolchain gotchas (Webpack sourcemaps, Vue watchers); signals that real-time perf is the library's identity and that tooling footguns are a recurring support burden.
- **Wiki-first documentation posture** — Usage, Tutorials, References, Plugins, Getting Started, Running, and Rendering all route to the GitHub wiki rather than being inlined; the README is a directory, not a manual.
- **Indented four-space code blocks instead of fenced** — consistent with older-generation JS READMEs; combined with `###` headings and the HTML `<table>` it reads as a stylistically early-2010s document that has aged without a rewrite.
- **License section ends with a prose warranty disclaimer** — `"As stated in the license, absolutely no warranty is provided."` is redundant with the MIT link but reads as deliberate emphasis, not filler.

## Takeaway pattern

This is the "portal README" archetype: the README itself does almost no teaching — instead it acts as a hub routing visitors to the demo site, wiki, API docs, examples directory, codepen collection, and gallery. It works because the product is visual (physics simulation) and the author has built a rich companion site at brm.io; the README's job is to frame the project, prove traction via the gallery, and hand off to better-suited surfaces. Lesson: if your library has a dedicated docs site and a compelling demo surface, the README should be a one-screen orientation (tagline, nav, features, install) plus a proof-of-use gallery — resist the pull to re-teach what the wiki teaches better.
