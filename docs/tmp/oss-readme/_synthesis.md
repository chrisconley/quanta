# README patterns across 25 top GitHub repos — synthesis

Sampled 25 repos at random from the top 2000 GitHub projects by stars (range 18k–99k). Full list of repos and per-README analyses sit alongside this file; raw READMEs in `readmes/`.

## Sample at a glance

| Size bucket | Count | Characterization |
|---|---|---|
| ≤ 3 KB | 7 | Signposts to external docs; runnable one-liners |
| 3–10 KB | 9 | Balanced; features + install + example + links |
| 10–30 KB | 7 | Long-form; extensive feature matrices, team info, migration guides |
| > 30 KB | 2 | README-as-product (directories, full reference configs) |

The 25 repos cover web frameworks (tornado, beego, fastify), system/dev tools (neovim, coc.nvim, curl, openpilot, raylib), libraries (mybatis, fullcalendar, GSYVideoPlayer, listmonk, LosslessCut), AI/ML (fastmcp, Chinese-LLaMA-Alpaca, fchollet/deep-learning-with-python-notebooks, learn-claude-code), docs/tutorials (anthropics/courses, 3d-game-shaders-for-beginners, deeplearningbook-chinese, gold-miner), desktop apps (lencx/ChatGPT, Win11Debloat), transparency dumps (twitter/the-algorithm), and directories (TelegramGroup).

## Section patterns — what shows up and how often

Out of the 25 READMEs:

- **Title + one-line pitch** — 25/25 (universal)
- **Badges** (CI, version, license, community) — 20/25; absent in pure docs/tutorial repos
- **Hero image / logo / screenshot / GIF** — 18/25
- **Runnable code example** — 13/25 (near-universal in library READMEs; absent in directories and transparency dumps)
- **Installation** — 18/25
- **Features list** — 17/25
- **Table of contents** — 9/25 (more common as size grows past ~10 KB)
- **License** — 19/25
- **Contributing pointer** — 18/25
- **Community links (Discord, Slack, Reddit, etc.)** — 16/25
- **Roadmap / changelog** — 8/25 (often linked rather than inlined)
- **Team / maintainers listing** — 5/25 (governance-forward projects only)
- **Benchmarks** — 2/25 (fastify, raylib lists platforms — only fastify runs numbers)
- **Explicit security-reporting section** — 4/25 (curl, fastify, twitter, openpilot)

## Recurring macro-structures

Four main archetypes showed up repeatedly:

### 1. "Signpost" — mature/legacy infra
- Examples: curl, tornado, fullcalendar, anthropics/courses
- Traits: 1–3 KB; minimal prose; website carries the weight; few or no badges; license and contact tend to be surfaced over features.
- Works because the audience self-qualifies; fails as an on-ramp for newcomers.

### 2. "Pitch + Path" — modern developer library
- Examples: fastify, fastmcp, raylib, beego, listmonk, mifi/lossless-cut
- Traits: hero banner, badges row, tagline, runnable quick-start under 30 seconds, features list, docs link, contribution/license at the bottom.
- Strong "try it in 30 seconds" orientation. Feature lists are bulleted, not tabular.

### 3. "Reference + Manifesto" — projects with a strong opinion
- Examples: learn-claude-code, deeplearningbook-chinese, openpilot, lencx/ChatGPT
- Traits: long prose up top, often polemic or mission-driven; the README is the content, not just the door to content.
- Tradeoff: sets positioning well, but readers looking for code have to work for it.

### 4. "Directory / Index" — curation as product
- Examples: TelegramGroup, gold-miner, twitter/the-algorithm, anthropics/courses
- Traits: README dominated by tables or link lists; per-item descriptions; the repo's value is the index itself.
- Editorial labor (scam warnings, translator credits, per-row descriptions) becomes the product.

## Tone axes

- **Personal vs. institutional**: Solo-maintained consumer tools (lossless-cut, Win11Debloat, raylib, lencx/ChatGPT) lean personal — first-person, donation links, "Made with ❤️ in…". Foundation-governed projects (fastify, neovim, curl, mybatis) read institutional — third-person, team listings, LTS policy.
- **Confident vs. deferential**: Some READMEs *declare* value (fastmcp: "FastMCP is the standard framework"; learn-claude-code: "Claude Code is the most elegant and fully-realized agent harness"). Others hedge (deeplearningbook-chinese: "we can't eliminate variance, we need your help"). No correlation with star count.
- **Polemic vs. neutral**: Only learn-claude-code and deeplearningbook-chinese take a clear "side" of a debate in their README.

## Formatting habits worth noting

- **GitHub alerts** (`> [!NOTE]`, `> [!Warning]`, `> [!Tip]`) showed up in Win11Debloat, fastify, lencx/ChatGPT — a recent convention rapidly becoming standard.
- **Collapsible `<details>` blocks** were used to hide legalese / optional install paths (openpilot, Win11Debloat). Good pattern for reducing clutter while keeping content linkable.
- **`<picture>` with `prefers-color-scheme`** (fastmcp) handles GitHub's dark/light themes — expect this to spread.
- **Setext (`===` / `---`) vs ATX (`#`)** — older projects (mybatis, tornado, deeplearningbook-chinese, neovim, raylib) prefer Setext; newer projects prefer ATX. Good proxy for project age.
- **`llms.txt` / `llms-full.txt`** (fastmcp) — explicit LLM-ingestible doc format. Early indicator.
- **DeepWiki badges** (coc.nvim) — LLM-generated external docs are starting to be first-class README entries.
- **Reference links at the bottom** (neovim, mybatis) keep inline prose shorter — aesthetic, not functional, but a marker of careful authorship.

## Non-obvious conventions specific to certain communities

- **Chinese-market repos** (Chinese-LLaMA-Alpaca, GSYVideoPlayer, TelegramGroup, gold-miner, deeplearningbook-chinese): multiple mirrors for downloads (HuggingFace + ModelScope + Baidu / GitCode + Gitee / JitPack + Aliyun Maven), alternative blog platforms (Juejin, Zhihu, CSDN, Jianshu), QQ groups beside Discord/Slack, WeChat QR images. These aren't cosmetic — they reflect actual accessibility constraints.
- **AI model releases** (Chinese-LLaMA-Alpaca): date-stamped "News" sections doubling as changelog; successor-project banners; multi-mirror download tables per model.
- **Tutorial repos** (lettier/3d-game-shaders-for-beginners, fchollet notebooks): README acts as a clickable ToC; content lives in per-topic files; often directly-runnable via Colab or per-section markdown.
- **Transparency dumps** (twitter/the-algorithm): no install, no examples — the repo's purpose is *reading*, not running. README is a map rather than a manual.

## Common mistakes / weaknesses observed

- **No license in README** (anthropics/courses, TelegramGroup, ymcui/Chinese-LLaMA-Alpaca, fchollet notebooks): for content-heavy repos this is forgivable, but still confusing to downstream users.
- **Embedded secrets** (GSYVideoPlayer includes a real GitHub PAT for "user convenience"). Security anti-pattern worth avoiding.
- **README-as-full-docs** (coc.nvim at 112 KB): makes onboarding hard for users who just want to know if the tool fits.
- **Legacy banners left in place long after successor exists** (Chinese-LLaMA-Alpaca, lencx/ChatGPT) — these handle the transition well, but the README still looks active unless read carefully.

## Recommendations distilled for a new project README

1. **Open with a one-sentence pitch that names the *category* and the *differentiator*.** Most strong READMEs do this in the first paragraph.
2. **Prove it runs within a screenful** — install + minimum viable snippet. Skip it only if your audience is pre-qualified (curl, tornado).
3. **Put "why" above "features" for opinionated projects**, "features" above "why" for utility tools.
4. **Use GitHub alerts and `<details>` to reduce noise without deleting information.**
5. **Link-heavy documentation beats prose-heavy documentation** — almost every README in the sample gets shorter by linking out to docs pages for depth.
6. **Badges should say something scannable**: version, license, build. Skip the social-media vanity row unless the community truly lives there.
7. **For tutorial/book/model repos, ToC-style READMEs work well**; don't try to cram content into the front page.
8. **Governance signals (team, sponsors, LTS) matter for enterprise adoption** even though they bore casual readers — put them at the bottom.
9. **Localize thoughtfully** — sibling language READMEs (`README_EN.md`, etc.) beat a single bilingual mashup; link them at the top.
10. **README length should match audience expectations, not project size.** A 100 KB kernel can ship a 2 KB README (curl); a 10 KB toolkit can reasonably ship a 30 KB README if its audience expects to live in it (coc.nvim).

## Selected list

| Repo | Stars | README size |
|---|---|---|
| neovim/neovim | 98,886 | 5.3 KB |
| twitter/the-algorithm | 73,025 | 7.3 KB |
| commaai/openpilot | 60,673 | 7.2 KB |
| shareAI-lab/learn-claude-code | 55,119 | 24.4 KB |
| lencx/ChatGPT | 54,369 | 1.8 KB |
| Raphire/Win11Debloat | 44,966 | 9.7 KB |
| curl/curl | 41,353 | 2.1 KB |
| mifi/lossless-cut | 39,909 | 12.2 KB |
| exacity/deeplearningbook-chinese | 37,242 | 9.8 KB |
| fastify/fastify | 36,088 | 16.5 KB |
| xitu/gold-miner | 34,319 | 13.4 KB |
| beego/beego | 32,416 | 3.4 KB |
| raysan5/raylib | 32,127 | 10.5 KB |
| neoclide/coc.nvim | 25,148 | 112.5 KB |
| PrefectHQ/fastmcp | 24,686 | 5.9 KB |
| tornadoweb/tornado | 22,259 | 1.6 KB |
| CarGuo/GSYVideoPlayer | 21,428 | 23.9 KB |
| AZeC4/TelegramGroup | 21,179 | 249.7 KB |
| anthropics/courses | 20,725 | 1.4 KB |
| fullcalendar/fullcalendar | 20,435 | 1.8 KB |
| mybatis/mybatis-3 | 20,422 | 3.3 KB |
| fchollet/deep-learning-with-python-notebooks | 20,041 | 6.1 KB |
| knadh/listmonk | 19,575 | 2.1 KB |
| lettier/3d-game-shaders-for-beginners | 19,539 | 2.7 KB |
| ymcui/Chinese-LLaMA-Alpaca | 18,946 | 31.3 KB |
