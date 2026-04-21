# Analysis: Tencent/APIJSON

- **Repo**: Tencent/APIJSON (~17.9k stars)
- **Type**: JSON-based communication protocol + ORM library that auto-generates backend CRUD APIs and lets frontend shape response JSON without server-side coding.
- **Length**: ~22 KB (exceptionally long for a README). Translated-from-Chinese feel with sprawling galleries of badges, screenshots, user logos, and a parody song.

## Sections

1. **Tencent copyright header** — four-line ASCII block: "Tencent is pleased to support the open source community…", copyright 2020, Apache 2.0 notice (plain text, not a blockquote).
2. **Centered `<h1>`** — "APIJSON" in plain HTML.
3. **Tagline with emoji** — "🏆 Real-Time no-code, powerful and secure ORM 🚀" plus descriptor.
4. **Top nav link row** — Chinese README, Document, Video (Bilibili search), Test (apijson.cn/api), "Ask AI" (DeepWiki).
5. **Three massive badge stacks** — (a) ~30 database version shields (MySQL, PostgreSQL, SQL Server, Oracle, DB2, MariaDB, TiDB, CockroachDB, openGauss, Dameng, Kingbase, Milvus, DuckDB, SurrealDB, ClickHouse, PostGIS, Elasticsearch, Manticore, Presto, Trino, TDSQL, TencentDB, Redis, Kafka, Snowflake, Databricks, Hive, Hadoop, MongoDB, Cassandra, InfluxDB, TDengine, TimescaleDB, IoTDB, DataBend), (b) 9 language-port badges (Java, Go ×2 forks, C#, PHP, Node, Python, Rust, Lua), (c) Spring/SpringBoot/JFinal/Nutz/ShardingSphere badges, (d) Android/iOS/JavaScript client badges.
6. **Hero diagram image** — centered oschina-hosted PNG of the architecture.
7. **Seven-item numbered TOC** — About, Backend usage, Frontend usage, Contributing, Releases, Creator, Donating (links 1–7 via `<h2 id="N">` anchors).
8. **1. About** — what it is (JSON protocol + ORM) plus "Features" broken into "For getting data" and "For API design" subsections.
9. **Katy Perry "Firework" parody lyric** — 20-line rewritten song ("Do you ever feel like a backend slave / Repeating CRUD…") interspersed with line-break `<br />` tags.
10. **Marketing hook paragraph** — bolded "Tired with endless arguments about HTTP API dev or use?" + "Unfold the Power(In Your Soul) with ⭐Star & Clone."
11. **APIJSON Show gallery** — five screenshots/GIFs of Postman, APIAuto UI, generated code download, auto regression tests, and basic-features demos; each with a centered caption paragraph.
12. **2. Backend usage / 3. Frontend usage** — two near-empty sections pointing elsewhere (GitHub links to Java/Android/iOS/JavaScript sub-READMEs); includes `.apk` download links hosted on cnblogs.
13. **4. Contributing** — "Please also ⭐Star the project!" tacked onto standard fork/PR instructions.
14. **5. Releases / 6. Creator** — single-link sections; Creator section includes an image of the creator Tommy Lemon plus his email (mailto:tommylemon@qq.com).
15. **Users of APIJSON** — a `<div style="float:left">` gallery of 45+ corporate logos (Transsion, xmfish, juhu, aupup, YTO, etc.) sized at 75px high.
16. **Contributers of APIJSON** (sic, "Contributers" misspelled twice) — paragraphs listing employer affiliations ("6 Tencent engineers, 1 Microsoft engineer, 1 Zhihu architect, 1 Bytedance(TikTok) engineer…") with two screenshot images dated 2026-04-18.
17. **Statistics** — stargazer chart via starchart.cc plus two growth screenshots; brags that "Hundreds of employees from big famous companies(Tencent, Google, Apple, Microsoft, Amazon, Huawei, Alibaba…) starred".

## Tone & style

- **Register**: Grand, maximalist, and self-promotional with a strong translated-from-Chinese voice ("Unfold the Power(In Your Soul)," "you just gotta depend and configure / And let it init"); heavy use of `<br />` line breaks instead of paragraph flow.
- **Voice**: Second-person exhortations ("Please also ⭐Star the project!"), third-person product pitch, first-person parody-song narrator, and quantitative boasts ("Hundreds of employees from big famous companies…").
- **Formatting**: Inline HTML throughout (`<h2 id="1">…<h2/>` with malformed closing tags on every section header), `<p align="center">` blocks, `<br />` line breaks every sentence, `<div style="float:left">` for the user logo wall, badge overload, dated screenshot filenames (`Screenshot 2026-04-18 at 05 28 48`).

## Notable conventions

- **Malformed closing H2 tags on every section** (`<h2/>` instead of `</h2>`) — GitHub's renderer tolerates it but the pattern repeats 7+ times, suggesting a copy-paste template nobody has corrected in years.
- **Katy Perry "Firework" parody embedded mid-README** — an absolutely distinctive choice; rewrites a pop song to pitch the library ("'Cause there's a powerful tool / You just gotta depend and configure / And let it init"). No other top-starred project does this.
- **30+ database badges in one row** — many point to the same demo repo; the visual weight implies "every data store works" more than any prose claim could.
- **Creator section with creator's personal `qq.com` email and photo** — uncommon directness; most Tencent-owned OSS hides the maintainer behind a corporate address.
- **Corporate-logo wall of users (45+ logos)** — treats this as a trophy case; `<div style="float:left">` rather than centered grids, giving a ragged left-aligned visual.
- **Explicit employer-counting of contributors** — "6 Tencent engineers, 1 Microsoft engineer, 1 Zhihu architect, 1 Bytedance(TikTok) engineer, 1 NetEase engineer, 1 Zoom engineer, 1 YTO Express engineer…" — social proof through FAANG-and-friends headcount rather than GitHub handles.
- **Backend/Frontend sections are one sentence each**, pointing to sub-READMEs — the main README is a billboard, not a manual.
- **"Ask AI" DeepWiki link in the header** — mirrors the elizaOS README convention; emerging 2026 standard for delegating Q&A to AI doc browsers.
- **"Contributers" misspelled twice** — the error persists despite 18k stars; emblematic of the "never edit, only append" maintenance style.

## Takeaway pattern

This is the "trophy-case README" archetype, cranked to its maximum: every piece of social proof (databases supported, language ports, corporate users, famous-company stargazers, contributor employers, creator's personal email) is pulled forward onto the front page, with the actual technical story ("1. About") occupying maybe 5% of the scroll. The Katy Perry parody and the "Unfold the Power(In Your Soul)" copy give it an unmistakable translated-pop-culture signature that you will never see on an American-originated OSS project. It works as a top-of-funnel conversion page for the Chinese enterprise market — the logo wall and employer count are the pitch — but it actively fails Western OSS norms around brevity, skimmability, and doc/marketing separation. Lesson: READMEs are culturally encoded. What reads as maximalist or unprofessional in a Western lens reads as credibility-building and community-celebrating in a Chinese OSS lens; the "never remove, only append" editorial style produces artifacts that function as portfolios of the project's accumulated reputation.
