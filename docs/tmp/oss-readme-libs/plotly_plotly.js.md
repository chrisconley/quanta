# Analysis: plotly/plotly.js

- **Repo**: plotly/plotly.js (~18.2k stars)
- **Type**: Standalone JavaScript data visualization library that also powers Python `plotly` and R `plotly` bindings.
- **Length**: ~8 KB. Medium; most depth is in the contributor table, not in API content.

## Sections

1. **Logo + 3 shield badges + promo art** — centered `plotlyjs-logo@2x.png`, then npm/CI/MIT badges, then a long intro paragraph framing plotly.js as the shared engine behind Plotly.py and Plotly.R, then a centered promo PNG (`plotly_2017.png`), a "[Contact us] for Plotly.js consulting" line, and a second centered "Maintained by Plotly" banner image linking to `dash.plotly.com/project-maintenance`.
2. **Table of contents** — 9 bulleted links to the sections below, preceded by the H2 `## Table of contents`.
3. **Load as a node module** — `npm i --save plotly.js-dist-min` block plus an ES6 `import` and a CommonJS `require` snippet, with a callout pointing to `plotly.js-dist` for unminified.
4. **Load via script tag** — H3 "The script HTML element" with a blockquote tutorial aside ("basic knowledge of HTML and JSON syntax is enough to get started, i.e., with or without JavaScript!"), then a complete `<head>`/`<body>` runnable HTML example with `cdn.plot.ly/plotly-3.5.0.min.js` and a `div#gd` + `Plotly.newPlot` script block. Follow-on H3s cover ES6 module imports, un-minified CDN URLs with charset guidance, a v1/v2 "plotly-latest" deprecation note, **MathJax** loading (v2 and v3 variants), and **Need to have several WebGL graphs on a page?** which tells you to preload `virtual-webgl@1.0.6`.
5. **Bundles** — two-item numbered list distinguishing official `dist` bundles from user-built custom bundles, pointing to `dist/README.md` and `CUSTOM_BUNDLE.md`.
6. **Alternative ways to load and build plotly.js** — single paragraph redirecting advanced consumers to `BUILDING.md`, mentioning `lib/index.js` and `lib/index-basic.js` by path.
7. **Documentation** — two-sentence section; notes docs are generated from `graphing-library-docs` via Jekyll on GitHub Pages.
8. **Bugs and feature requests** — one-paragraph pointer to issue templates and `CONTRIBUTING.md`.
9. **Contributing** — one-paragraph pointer to `CONTRIBUTING.md`.
10. **Notable contributors** — intro paragraph + a 23-row markdown table with columns `Contributor | GitHub | Status`, where Status is one of `Active, Maintainer`, `Active, Community Contributor`, or `Hall of Fame`.
11. **Copyright and license** — "Code and documentation copyright 2025 Plotly, Inc." + MIT link + a **Versioning** subsection citing SemVer and linking the GitHub Releases page.
12. **Community** — three-bullet list pointing at X/LinkedIn, community forum (tagged `plotly-js`), Stack Overflow (tagged `plotly.js`), and an npm-keyword convention (`plotly`) for downstream packages.

## Tone & style

- **Register**: Corporate-OSS — the README reads like a company landing page ("Contact us for Plotly.js consulting," "Maintained by Plotly" banner, explicit 2025 copyright on Plotly, Inc.).
- **Voice**: Third-person descriptive throughout, with occasional second-person instructional in Install ("You may also consider using...", "Please note that as of v2...").
- **Formatting**: Heavy use of `---` horizontal rules between every top-level section, centered `<p align="center">` / `<div align="center">` blocks for logo and maintenance banner, H3 question-style subheadings inside "Load via script tag" (`### Need to have several WebGL graphs on a page?`), and a long `<script src="https://cdn.plot.ly/plotly-3.5.0.min.js">`-style example that doubles as copy-paste starter HTML.

## Notable conventions

- **"Maintained by Plotly" banner as social proof** — an image-sized commitment statement linking to a dedicated `project-maintenance` page, not just a text line. Frames the project as vendor-backed rather than community-drift OSS.
- **Contact-us-for-consulting line above the fold** — the monetization path is named in the first 20 lines ("Contact us for Plotly.js consulting, dashboard development, application integration, and feature additions"); more honest than most company-owned OSS which bury commercial services below the fold.
- **Notable contributors table with "Hall of Fame" status** — a tiered recognition system (`Active, Maintainer` / `Active, Community Contributor` / `Hall of Fame`) running 23 rows deep. Preserves credit for alumni without implying current availability — a thoughtful solution to the "is this person still on the project?" question that most contributor lists leave ambiguous.
- **Section separators as `---` horizontal rules** — visually chunks the README into distinct pages; combined with the H2s this produces a "brochure" rhythm that rewards scanning over reading.
- **Explicit CDN version pinning guidance** — the "plotly-latest will no longer be updated past v1.58.5" blockquote is a breaking-change notice baked into the install instructions themselves, not relegated to a changelog.
- **MathJax v2-vs-v3 install instructions as a dedicated H3** — signals that scientific/equation rendering is a first-class use case the maintainers expect you to hit; most visualization libraries wouldn't call this out.
- **"Developers should use the keyword `plotly` on packages... distributing through npm"** — ecosystem-management via npm keyword convention, rare to see spelled out.

## Takeaway pattern

This is the "company-OSS landing page" archetype: the README markets a venture-backed project to three distinct audiences (script-tag hobbyists, npm consumers, enterprise buyers) in the same document, using centered banner imagery, monetization callouts, and a tiered contributor table to signal corporate backing and community depth simultaneously. It works because plotly.js genuinely serves all three — a solo dev copy-pasting a script tag and a Fortune 500 procurement team evaluating Dash both land on this page. Lesson: when your OSS is commercially backed, don't hide it — name the consulting offering near the top, dedicate a maintenance banner to it, and use a tiered contributor/alumni table to show continuity; readers would rather see the commercial context up front than discover it after integrating.
