# README patterns across 25 lower-star top-2000 libraries — synthesis

Sampled 25 libraries from the bottom of the top-2000 GitHub projects by stars (a narrow star band, ~17.5k–18.9k). These are all code libraries — no courses, directories, tutorials, transparency dumps, or consumer desktop apps — which is a deliberate contrast with the prior random-25 sample. Full list of repos and per-README analyses sit alongside this file; raw READMEs in `readmes/`.

## Sample at a glance

| Size bucket | Count | Characterization |
|---|---|---|
| ≤ 3 KB | 4 | Stub READMEs; most content lives on hosted docs site or in git history |
| 3–10 KB | 12 | Landing-page-lite: pitch + badge row + one code block + docs link |
| 10–30 KB | 8 | Inline reference / cookbook / feature-comparison; README is the manual |
| > 30 KB | 1 | skrollr — README-as-complete-book (jQuery-era single-author manual) |

All 25 are libraries, frameworks, or toolkits — no directories, content repos, or consumer apps. That filter tightens the pattern space considerably versus the random sample. The mix spans: JS/TS UI and animation (matter-js, plotly.js, motion-canvas, wangEditor, naive-ui, vant-weapp, Shopify/draggable, skrollr, dropzone, eliza), front-stack frameworks (wasp-lang/wasp), PHP (CodeIgniter), Android/iOS (flexbox-layout, Masonry), .NET (Dapper), Java (LMAX-Exchange/disruptor), ML/AI (LightGBM, timesfm, openai/evals), scientific Python (plotly.py, akshare), infrastructure (golang-migrate, auth0/node-jsonwebtoken), developer tooling (conventional-changelog/commitlint), and China-origin backend (Tencent/APIJSON).

## Section patterns — what shows up and how often

Out of the 25 READMEs:

- **Title + one-line pitch** — 25/25 (universal; wangEditor's is Chinese-only)
- **Badges** — 21/25 (absent in CodeIgniter, timesfm, openai/evals, wangEditor — the projects that either predate shields.io or actively refuse them)
- **Hero image / logo / screenshot / GIF** — 20/25
- **Runnable code example** — 17/25 (higher than random-25's 13/25; libraries with clear entry points default to "show the first call")
- **Installation** — 20/25
- **Features list** — 14/25
- **Table of contents** — 4/25 (lower than random-25's 9/25; these READMEs are shorter on average and skip the TOC)
- **License section** — 15/25
- **Contributing pointer** — 15/25
- **Community links (Discord / Slack / forum / DingTalk)** — 12/25
- **Roadmap / changelog (inline or linked)** — 10/25
- **Team / maintainers listing** — 7/25
- **Benchmarks** — 1/25 (just Dapper, and the table honestly shows Hand-Coded SQL beating Dapper)
- **Explicit security-reporting section** — 2/25 (jsonwebtoken, CodeIgniter)
- **Retirement / legacy / successor notice** — 5/25 (Masonry, CodeIgniter, Shopify/draggable, skrollr, wangEditor-by-implication) — substantially higher than random-25
- **Commercial / consulting / sponsor-product funnel** — 10/25 (plotly.js, plotly.py, Dapper, wasp, APIJSON, akshare, skrollr, eliza, timesfm, commitlint)

## Recurring macro-structures

Four archetypes dominate this cohort; two of them also appeared in the random-25 sample, two are new or amplified here.

### 1. "Hub / signpost" — README as a router to the real docs site
- Examples: matter-js, LightGBM, naive-ui, vant-weapp, wangEditor, LMAX-Exchange/disruptor, motion-canvas, plotly.js, plotly.py
- Traits: short (often ≤ 10 KB), badge row, one-paragraph pitch, one install line (sometimes none), and a link bar or doc-site pointer doing all the heavy lifting. Code examples are sparse or absent — the docs site has them.
- Works when the hosted docs site is genuinely better than any README could be. Fails as drive-by discovery: wangEditor at 0.4 KB doesn't even tell you it's a React/Vue editor without clicking through.
- Over-represented versus random-25, because library projects with mature doc sites outgrow inline tutorialization by the time they cross ~15k stars.

### 2. "Inline reference / cookbook" — README as the complete API
- Examples: auth0/node-jsonwebtoken, DapperLib/Dapper, google/flexbox-layout, SnapKit/Masonry, skrollr
- Traits: 15–30+ KB, one worked example per feature, `<details>` collapsibles for unsafe APIs (jsonwebtoken hides `decode` behind a warning), attribute tables, before/after migration notes. Typically setext headings (`===` / `---`) and older formatting habits.
- Works for small, stable API surfaces (jsonwebtoken has three functions; Dapper has three extension methods). The README becomes the canonical lookup surface because users read it via `npm view`, search engines, or package-manager previews rather than a separate docs site.
- No equivalent archetype at this density in the random-25 — there, only `coc.nvim` at 112 KB went this way.

### 3. "Retirement / legacy / successor" — README as a honest handoff
- Examples: Shopify/draggable ("no longer worked on by anyone at Shopify"), SnapKit/Masonry ("we recommend SnapKit for Swift"), skrollr ("no active development since September 2014"), bcit-ci/CodeIgniter ("CodeIgniter 3, in maintenance mode, use CodeIgniter 4")
- Traits: the first substantive paragraph (or immediately below the title) tells you the project is in sunset mode, often with a named successor and a rationale. Some place the notice under a generic heading ("Development" in Shopify/draggable) to preserve the rest of the README structure; others rewrite the top paragraphs wholesale.
- 4/25 in this cohort versus ~0 in the random-25 sample. The lower-star band captures projects at the long-tail end of their lifecycle: widely starred in their era, now on the downslope.
- When done well (skrollr, CodeIgniter, Masonry), these READMEs are among the most honest in the sample.

### 4. "Trophy-case" — social proof as README content
- Examples: Tencent/APIJSON, plotly.js, akfamily/akshare, lightgbm-org/LightGBM, wasp (narrative sponsors), matter-js (gallery)
- Traits: galleries of consumer logos, tiered contributor tables ("Hall of Fame"), exhaustive ecosystem lists (LightGBM's ~35 "External (Unofficial) Repositories"), citation papers, stargazer-prestige blurbs ("Hundreds of employees from Tencent, Google, Apple, Microsoft, Amazon, Huawei, Alibaba…"), paper/citation BibTeX.
- APIJSON takes this to an extreme with a Katy Perry "Firework" parody, a 45-logo user wall, and employer counts for contributors. LightGBM is the restrained institutional version: research papers cited, ecosystem enumerated, no prose boasting.
- "Ecosystem as README content" is a distinctive form of this — LightGBM and APIJSON both list every known third-party integration as a way to demonstrate centrality without claiming it.

## Tone axes

- **Vendor-forward vs community-first**: Corporate-owned OSS here (plotly.js, plotly.py, google/flexbox-layout, google-research/timesfm, openai/evals, Tencent/APIJSON, auth0/node-jsonwebtoken) consistently surfaces commercial context — consulting banners, sibling-product cross-sells, "1P Products" sub-lists, contributor-data-rights disclaimers. Community-first READMEs (matter-js, dropzone, motion-canvas, LMAX/disruptor) surface none of that; the tone is neutral or personal.
- **Institutional vs personal voice**: Solo-authored libraries leak authorial voice even at 18k stars — naive-ui's "Kinda Interesting" tagline and "I try to make it not rather slow," skrollr's "I was lying to you" retraction, Masonry's "Prepare to meet your Maker!" chapter titles. Foundation-style READMEs (LightGBM, commitlint, migrate) are impersonal throughout.
- **Confident vs deferential**: A surprising share in this cohort lean deferential ("I try to make it not rather slow," "Not affiliated or supported," "you shouldn't compromise UX for some fancy UI effects. Ever.") — often paired with legacy or maintenance status. Opposite extreme: APIJSON's "🏆 Real-Time no-code, powerful and secure ORM 🚀" and the "Unfold the Power(In Your Soul)" marketing copy.
- **Cultural register**: Chinese-origin projects (wangEditor, vant-weapp, APIJSON, akshare) show consistent patterns — mirrored doc sites, DingTalk/WeChat community channels, QR codes as demo surface, maximalist "never remove, only append" editorial style, translated-from-Chinese tone on English surfaces.

## Formatting habits worth noting

- **Setext headings (`===` / `---`) are common here** — Dapper, LightGBM, skrollr, sometimes with modern content beneath. Good proxy for project age; five of the 25 use setext at the H1 or H2 level. The random-25 sample showed the same correlation.
- **GitHub `[!NOTE]` alerts** — LightGBM used one for a repo-migration announcement. Rarer in this cohort than random-25, reflecting that most of these projects have not been editorially refreshed since the alerts landed.
- **Collapsible `<details>` blocks** — jsonwebtoken hides `jwt.decode` behind one (because the API is unsafe); eliza hides advanced CLI commands. Underused relative to its value for reducing clutter.
- **`<picture>` with `prefers-color-scheme: dark`** — motion-canvas is the sole example in this cohort. Still rare.
- **Reference-style link footnotes at the bottom** — motion-canvas, commitlint, skrollr. Correlates with careful authorship, not project age specifically.
- **AsciiDoc content in `.md` files** — LMAX-Exchange/disruptor uses AsciiDoc syntax (`=` title, `image:url[...]`) in a file GitHub renders as Markdown; disrupts copy-paste and IDE rendering. CodeIgniter uses reStructuredText similarly.
- **HTML `<table>` for badges** — plotly.py uses a 4-row HTML table for its badge wall instead of the standard inline shields run; a pre-shields.io convention preserved.
- **`:sparkles:` and other GitHub emoji shortcodes** — plotly.py uses `:sparkles:` (renders on GitHub only). Emoji-shortcode vs literal-emoji is a small authorial signal.
- **Malformed-but-tolerated HTML** — APIJSON uses `<h2/>` self-closing tags for every section heading, plotly.py wraps image anchors as `<a href=".../>`; GitHub's renderer tolerates both, but they're editorial tells of a "never rewrite, only append" style.
- **DeepWiki / "Ask AI" links** — eliza, APIJSON surface AI-assisted doc browsing as first-class entry points. Expect this to spread rapidly.

## Non-obvious conventions specific to certain communities

- **Chinese-ecosystem READMEs** (wangEditor, vant-weapp, APIJSON, akshare): dual doc-site mirrors (`国内` vs. GitHub Pages), DingTalk/WeChat group IDs with "member limit reached" flags, QR codes as the primary demo surface, Aliyun package mirrors with HTTP-and-trust-bypass for PyPI in China, corporate logo walls as trophy cases, "Contributers" (sic) counts broken down by FAANG employer affiliation, and bilingual asymmetries where Chinese prose runs above English.
- **Corporate-OSS with commercial cross-sell** (plotly.js, plotly.py, timesfm, openai/evals, auth0/node-jsonwebtoken, Dapper): consulting banners, "Maintained by X" images, sibling-product funnels (Dash for Plotly, Dapper Plus for Dapper, BigQuery ML for TimesFM), and contribution-data-rights disclaimers ("OpenAI reserves the right to use this data in future service improvements").
- **Research releases** (timesfm, LightGBM, eliza): arXiv paper badges, BibTeX citations, "not an officially supported product" disclaimers, and inline "Update - <Date>" changelog blocks replacing a separate CHANGELOG.
- **Long-maintenance legacy libraries** (Masonry, CodeIgniter, skrollr, Shopify/draggable): successor links, retirement dates, honest "don't compromise UX for UI effects" editorial asides, and TODO lists preserved decades after they stopped being tracked.

## Common weaknesses observed

- **Dead-service badges left in place** — auth0/node-jsonwebtoken still shows Travis CI and david-dm shields (both defunct); many others have similar unrefreshed badges. A near-universal editorial-maintenance failure in this band.
- **Typos and malformed markup left in place** — "expeced" (flexbox-layout), "Contributers" x2 (APIJSON), "anonyomous" (Dapper). A 17k-star README apparently gets no copy-editor.
- **Renamed-project artifacts** — akshare still shows `jindaxiang` in three shield URLs despite having moved to `akfamily/akshare`; LightGBM still shows Microsoft Code of Conduct and `Microsoft.LightGBM` Winget badge despite moving to `lightgbm-org`. Partial migrations are common.
- **Assumes audience already knows the project** — LMAX/disruptor has no quickstart, no Maven coordinates, no "why would I use this over `java.util.concurrent`"; Shopify/draggable skips from install to CDN table without ever showing `new Draggable(...)`. Works for famous projects; fails drive-by readers.
- **"Install straight to CDN" skipping the first-code snippet** — draggable, some plotly-ecosystem READMEs. Common in the lower-star band where authors forget newcomers exist.
- **README-as-monument** — APIJSON at 22 KB with parody song lyrics and 45 corporate logos; the editorial posture is "never remove, only append," producing a document that reads as a portfolio of accumulated attention rather than a how-to.
- **No license section** — akshare, motion-canvas, matter-js (has one but it's a one-liner), APIJSON (Apache in a text-plain header). License is always present in repo metadata, but absent from the README a surprising amount of the time.

## Contrast with the random-25 sample

- **Narrower star band, more uniform archetypes**: all 25 sit in ~17.5k–18.9k. The distribution of archetypes is flatter — no 99k-star anomalies like neovim or curl, and no content-directory or tutorial repos. So the patterns compress.
- **Libraries only, no content repos**: the random-25 included book directories, Telegram group curations, deep-learning-book translations, course repos, and transparency dumps. This cohort is pure library/framework/toolkit, which shifts the README format toward "install + example + docs link."
- **Retirement / legacy is much more common here** (5/25 vs ~1/25 in random). The bottom-of-top-2000 band captures projects that starred big in their era (2012–2018) and are now on the downslope; the high-star band captures projects still on the growth curve.
- **Inline API-reference READMEs are more common here** (jsonwebtoken, Dapper, flexbox-layout, Masonry, skrollr). A longtail library with a small API surface tends to fold docs into the README rather than maintain a separate site. The high-star cohort had this only at the coc.nvim extreme.
- **Code examples are more frequent** here (17/25 vs 13/25). Libraries benefit from a first-run snippet more than directories or transparency repos do.
- **TOCs are less frequent** here (4/25 vs 9/25) — shorter README average, less need for navigation anchors.
- **Commercial funnels are more visible** — 10/25 have an explicit monetization signal (Dapper Plus banner, Maintained-by-Plotly, Careers page, consulting CTA, sibling-product cross-sell). The random-25 had fewer, partly because foundation-style projects (neovim, tornado, curl) dominated its upper half.
- **Non-Western patterns are amplified** here — four clear China-origin READMEs (APIJSON, akshare, wangEditor, vant-weapp) versus two in the random-25 (Chinese-LLaMA-Alpaca, GSYVideoPlayer); plus the translated-from-Chinese tone in APIJSON is the most distinctive cultural tell in the sample.
- **Governance / maintainers / security signals are equally rare** in both cohorts — ~2 of 25 in each have an explicit security-reporting section, which is lower than good practice would suggest for libraries this deeply installed.
- **`[!NOTE]` alerts, `<picture>` dark-mode logos, `llms.txt`-style LLM pointers** — the newer formatting conventions that were emerging in the random-25 cohort are rarer here, consistent with these READMEs not having been editorially refreshed in the last two years.

## Recommendations distilled for a library-specific README

1. **If your library is famous enough to not need pitching, the README still has to answer "what is this?" for drive-by readers.** LMAX/disruptor's zero-explanation README and Shopify/draggable's install-to-CDN-table skip fail this test.
2. **When the project is in maintenance, sunset, or succeeded state, say so in the first paragraph, with a named successor.** Masonry, CodeIgniter, and skrollr all do this well; Shopify/draggable wedges it under a misleading "Development" heading and confuses readers.
3. **Inline the API as README content when the surface is small enough to fit.** Three-function libraries (jsonwebtoken, Dapper) end up serving readers better with a 15–25 KB inline reference than a separate docs site they'd skip.
4. **When the docs site is genuinely better, shrink the README to a router** — logo, tagline, one install line, one doc link. naive-ui does this in 2.7 KB with a distinct authorial voice; wangEditor does it in 0.4 KB and leaves drive-by readers nothing.
5. **Be upfront about commercial context.** Plotly and Dapper name their paid products above the fold; OpenAI/evals names the Dashboard redirect in its first blockquote. Readers prefer this to discovering it after integration.
6. **Don't leave dead-service badges, typos, and renamed-org artifacts in place.** Every long-tail README in this sample has at least one. A five-minute editorial pass once a year catches them.
7. **Add a "Known differences from the original specification" section when your library ports a well-known standard** — google/flexbox-layout's section is the fastest on-ramp for a reader who already knows CSS Flexbox, and demonstrates maturity more clearly than any feature bullet.
8. **For research code, replace CHANGELOG with inline "Update - <Date>" blocks**; returning readers see the deltas in order and don't click through.
9. **If your audience lives in a walled ecosystem (WeChat, GFW-restricted China, internal enterprise), the README earns its keep by routing around it** — mirror links, QR codes, localized license pointers are features, not quirks.
10. **Keep the voice human.** The strongest READMEs in this sample (naive-ui, skrollr, Masonry, commitlint's "We're not a sponsored OSS project") carry a recognizable authorial tone that corporate-template READMEs can't replicate — and the tone helps readers calibrate expectations faster than any feature bullet can.

## Selected list

| Repo | README size |
|---|---|
| Prinzhorn/skrollr | 31.2 KB |
| DapperLib/Dapper | 27.3 KB |
| Tencent/APIJSON | 23.4 KB |
| google/flexbox-layout | 17.8 KB |
| auth0/node-jsonwebtoken | 17.3 KB |
| SnapKit/Masonry | 15.8 KB |
| lightgbm-org/LightGBM | 12.6 KB |
| akfamily/akshare | 10.7 KB |
| plotly/plotly.js | 10.5 KB |
| elizaOS/eliza | 9.5 KB |
| liabru/matter-js | 8.8 KB |
| wasp-lang/wasp | 8.5 KB |
| conventional-changelog/commitlint | 8.5 KB |
| golang-migrate/migrate | 8.1 KB |
| Shopify/draggable | 7.4 KB |
| openai/evals | 6.3 KB |
| youzan/vant-weapp | 6.0 KB |
| motion-canvas/motion-canvas | 5.1 KB |
| google-research/timesfm | 4.5 KB |
| plotly/plotly.py | 4.2 KB |
| dropzone/dropzone | 3.6 KB |
| bcit-ci/CodeIgniter | 2.7 KB |
| tusen-ai/naive-ui | 2.6 KB |
| LMAX-Exchange/disruptor | 1.2 KB |
| wangeditor-team/wangEditor | 0.4 KB |
