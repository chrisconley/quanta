# Analysis: elizaOS/eliza

- **Repo**: elizaOS/eliza (~18.2k stars)
- **Type**: TypeScript framework/platform for building multi-agent AI applications (CLI + server + web UI + plugins).
- **Length**: ~9.8 KB. Medium-length; marketing-forward landing page with a condensed quick start.

## Sections

1. **Centered HTML header block** — project name, one-line positioning ("The Open-Source Framework for Multi-Agent AI Development"), and a subtitle elaborating on build/deploy/manage.
2. **Badge stacks** — five separate centered `<div>` rows: (a) Trendshift promo badge, (b) npm downloads / GitHub release / arXiv paper / DeepWiki, (c) "for-the-badge" GitHub stars / forks / last-commit, (d) MIT license / npm version / contributors, (e) Documentation / Twitter / Discord.
3. **What is Eliza?** — two-paragraph elevator pitch naming chatbots, business process automation, and game NPCs as use cases; links to hosted docs at docs.elizaos.ai.
4. **Key Features** — seven emoji-prefixed bullets (Rich Connectivity, Model Agnostic, Modern Web UI, Multi-Agent Architecture, Document Ingestion/RAG, Highly Extensible, "It Just Works"), plus a blockquote pointer to the external plugin registry at `elizaOS-plugins/registry`.
5. **Getting Started (5-Minute Quick Start)** — fork in the road: CLI path for beginners vs. monorepo path for contributors. Prerequisites (Node v23+, bun, WSL2 note for Windows). Four numbered steps: Install the CLI, Create Your Project, Configure Your API Key, Start Your Agent. Recommends pglite + openai + "project" type.
6. **Collapsible `<details>` block: Advanced CLI Commands & Usage** — development workflow, agent/environment management, debugging (with `LOG_LEVEL=debug`).
7. **Running elizaOS Core Standalone** — git clone + two `bun run examples/...` invocations for chat.ts and standalone.ts.
8. **Architecture Overview** — ASCII tree of the `packages/` monorepo (server, client, cli, core, app Tauri desktop, plugin-sql), with one-line descriptions of each of the four main packages.
9. **How to Contribute** — points to `CONTRIBUTING.md` and issue templates for bug reports and feature requests; instructs PR authors to open an issue first.
10. **License** — MIT.
11. **Citation** — BibTeX entry for "Eliza: A Web3 friendly AI Agent Operating System" (arXiv:2501.06781, Walters et al., 2025).
12. **Contributors** — contrib.rocks image.
13. **Star History** — star-history.com chart.

## Tone & style

- **Register**: Promotional / pragmatic hybrid — marketing landing-page phrasing ("all-in-one, extensible platform," "It Just Works") layered over step-by-step developer onboarding.
- **Voice**: Second-person imperative for instructions ("Get your first AI agent running in just a few commands"); first-person plural for project voice ("We welcome contributions").
- **Formatting**: Heavy HTML in the header (`<div align="center">`, width-pinned badges), emoji-prefixed section headings throughout, one collapsible `<details>` block, ASCII directory tree, BibTeX fenced block, horizontal-rule separators (`---`) dividing marketing from instructions.

## Notable conventions

- **arXiv paper badge alongside npm/download badges** — treats the project as both a published research artifact and a shipped package; the Citation section with BibTeX reinforces academic positioning.
- **Trendshift badge at the very top**, above functional badges — signals attention to growth/discovery rankings as a marketing surface.
- **Explicit "two paths" fork** in Getting Started separating CLI users from contributors, with the contributor path deferred rather than inlined — keeps the quick start genuinely quick.
- **DeepWiki "Ask DeepWiki" badge** — nods to AI-assisted doc browsing as a supported entry point.
- **Repo name (elizaOS), npm package names (@elizaos/core, @elizaos/cli), and prose name ("Eliza") aren't unified** — the readme uses all three interchangeably.
- **Citation's paper subtitle "A Web3 friendly AI Agent Operating System" reveals crypto-adjacent origin** (Discord invite is `discord.gg/ai16z`) that the main pitch avoids — a deliberate repositioning from crypto-native to general-purpose AI framework.

## Takeaway pattern

This is the "startup-rebranded framework" README: a project with crypto/Web3 roots (visible in the Discord URL and paper subtitle) repackaged as a general-purpose multi-agent platform, with heavy marketing layering (five badge rows, emoji headings, "It Just Works") covering a competent-but-ordinary CLI quick start. It works as a top-of-funnel conversion page — the `<details>` collapse keeps the quick start short, and the BibTeX citation buys academic credibility — but the tension between "all-in-one platform" promises and "7-bullet feature list plus generic architecture diagram" delivery is palpable. Lesson: when rebranding an existing project for a new audience, the old audience's fingerprints (Discord slugs, coauthor lists, paper titles) leak through and readers notice.
