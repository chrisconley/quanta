# Analysis: plotly/plotly.py

- **Repo**: plotly/plotly.py (~18.1k stars)
- **Type**: Interactive browser-based graphing library for Python, wrapping plotly.js; used in Jupyter, Dash, marimo, and standalone HTML.
- **Length**: ~4 KB. Short; the README defers to plotly.com/python for examples and to sibling repos (plotly.js, Kaleido, plotly-geo) for depth.

## Sections

1. **H1 project name** — plain `# plotly.py`.
2. **HTML badge `<table>`** — four rows (Latest Release / User forum / PyPI Downloads / License) presented as a 2-column `<tr>/<td>` table with anchor-wrapped image tags rather than the conventional inline badge row. Note the anchor tags use self-closing `<a ... />` syntax, which is technically invalid HTML but renders fine on GitHub.
3. **"Maintained by Plotly" centered banner** — a `<div align="center">` with a 400px-wide `maintained-by-plotly.png` linking to `dash.plotly.com/project-maintenance`. Doubles as a commercial-affiliation disclosure.
4. **Quickstart** — three lines: `pip install plotly`, a four-line `plotly.express` bar-chart snippet (`px.bar(x=["a","b","c"], y=[1,3,2])` → `fig.show()`), and a link to the Python documentation for more examples.
5. **Overview** — prose pitch: describes plotly.py as "interactive, open-source, and browser-based" with a `:sparkles:` emoji shortcode, highlights its plotly.js foundation and "over 30 chart types" (scientific, 3D, statistical, SVG maps, financial), notes MIT licensing, enumerates render targets (Jupyter, marimo, HTML files, Dash), and ends with "Contact us for consulting, dashboard development, application integration, and feature additions" linking to `plotly.com/consulting-and-oem`.
6. **Centered chart-gallery image** — `plot_images/add_r_img/plotly_2017.png` (a 2017-era chart collage) via `<p align="center">` with a target="_blank" anchor to `plotly.com/python/`.
7. **Horizontal-rule-bracketed link block** — five bullets: Online Documentation, Contributing, Changelog, Code of Conduct, Community forum. Framed by `---` rules above and below.
8. **Installation** — pip and conda-forge one-liners.
9. **Jupyter Widget Support** — install `jupyter` and `anywidget` via pip or conda.
10. **Static Image Export** — points at Kaleido (recommended as of v4.9) vs orca (legacy as of v4.9), with pip and conda install blocks for Kaleido.
11. **Extended Geo Support** — notes that large geographic shape files are split into a pinned `plotly-geo==1.0.0` package (pip and conda), with a link to the `plotly/plotly-geo` sibling repo.
12. **Copyright and Licenses** — code copyright 2019 Plotly, Inc., code MIT, docs under CC BY 4.0.

## Tone & style

- **Register**: Commercial-OSS measured — brand-first ("Maintained by Plotly" banner, "Contact us for consulting"), but without marketing breathlessness. No emoji headings, no exclamation points outside the `:sparkles:` inline.
- **Voice**: Third-person declarative for the pitch ("`plotly.py` is an interactive, open-source..."); first-person plural only in the consulting CTA ("Contact us").
- **Formatting**: HTML `<table>` for the badge row (unusual), centered `<div>`/`<p align="center">` for the hero images, horizontal rules bracketing the link block, backticked package names (`plotly.py`, `kaleido`, `plotly-geo`), GitHub emoji shortcodes (`:sparkles:`) rather than literal emoji.

## Notable conventions

- **HTML `<table>` for the badge row** instead of the now-standard inline badge-shield run — a pattern from the early-2010s README tradition that predates the shields.io convention. Each cell is also malformed: the `<a>` tags are written self-closing (`<a href="..."/>`) so the `<img>` technically isn't wrapped by the anchor, but GitHub's renderer forgives it.
- **"Maintained by Plotly" badge as commercial disclosure** — links to `dash.plotly.com/project-maintenance`, functioning as both a trust signal and a funnel into Plotly's commercial Dash product.
- **Dual package-manager instructions for every install step** (pip and conda, with `-c conda-forge` or `-c plotly` channels) — reflects the scientific-Python audience's split between PyPI and conda ecosystems; this is done for the main install, Jupyter widget, Kaleido, and plotly-geo.
- **Kaleido vs. orca note with version-gated deprecation** — "recommended, supported as of `plotly` version 4.9" vs "legacy as of `plotly` version 4.9" is a concise way to communicate a migration without writing a migration guide.
- **Pinned `plotly-geo==1.0.0`** — the README instructs readers to install a specific geo-shapefile package at an exact version, telegraphing that this sibling project is frozen rather than evolving.
- **Hero image dated 2017 (`plotly_2017.png`)** hosted under an unrelated path `cldougl/plot_images/add_r_img/` — a neglected asset that hasn't been refreshed as the chart-type count has grown, and whose path (`add_r_img`) suggests it was originally added for a different plotly client library (R).
- **"Contact us for consulting" inline in the Overview** — places the monetization CTA above the fold before the chart image and before installation, a deliberate priority the README's minimalism actually highlights.
- **Docs licensed CC BY 4.0 while code is MIT** — called out explicitly in the final section; uncommon dual-license disclosure that matters for downstream doc reuse.
- **`:sparkles:` emoji shortcode** rather than the literal character — a GitHub-specific Markdown flavor that renders as ✨ on GitHub but not elsewhere.

## Takeaway pattern

This is the "funded-scientific-OSS front door" README: short by design, delegates all depth to `plotly.com/python/` and to sibling repos (`plotly/plotly.js`, `plotly/Kaleido`, `plotly/plotly-geo`), and spends its limited real estate on (a) a quickstart that fits in five lines, (b) an Overview paragraph that lists consumer surfaces and render targets rather than features, and (c) explicit commercial-vendor signaling ("Maintained by Plotly" banner, consulting CTA, Dash cross-link). Lesson: when a company maintains a heavily-used scientific library as an ecosystem on-ramp for a commercial product, the README optimizes for "install works, link to docs, meet the vendor" and skips everything else — the legacy HTML `<table>` badges and the aging 2017 chart image survive because nobody is rewriting the README; they're iterating the product. The dual pip/conda instructions are the one concession to audience-specificity.
