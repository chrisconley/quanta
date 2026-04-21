# Analysis: wasp-lang/wasp

- **Repo**: wasp-lang/wasp (~18.3k stars)
- **Type**: DSL-driven full-stack framework for React + Node.js + Prisma ("a Rails-like framework" per the README); compiler written in Haskell.
- **Length**: ~8.7 KB. Medium-length; pitch + code sample + community / sponsor real estate.

## Sections

1. **Centered HTML header block** — wasp logo image, tagline "The fastest way to develop full-stack web apps with React & Node.js."
2. **Badge row** — three badges (license, latest release, Discord).
3. **Link rows** — two rows of centered inline links: (Website | Docs) and (Discord | Twitter | Youtube), then a third standalone link: "Deployed? Get swag! 👕" to a Typeform.
4. **Horizontal rule separator**.
5. **What is Wasp?** — etymology ("**W**eb **A**pplication **Sp**ecification"), Rails comparison, "Build your app in a day and deploy it with a single CLI command!"
6. **Why is Wasp awesome** (subsection) — four emoji-prefixed bullets: Quick start, No boilerplate, No lock-in ("you have complete control over the code (and can actually check it out in `.wasp/` directory"), and "Perfect for AI by design" (explicit AI-agent pitch).
7. **Features** (subsection) — six feature links to hosted docs (Full-stack Auth, RPC, Deployment, Jobs, Email Sending, Full-stack Type Safety) with a trailing "..." bullet.
8. **Code example** (subsection) — a `main.wasp` config snippet showing `app TodoApp`, `route`, `page`, and `query` declarations, plus a companion `schema.prisma` snippet with a `Task` model. Closed with a TodoApp tutorial link flanked by 👉 / 👈 emoji.
9. **How it works** (subsection) — architecture diagram image explaining that the Wasp compiler takes the `.wasp` file + source and emits a full web app.
10. **Get started** — `npm i -g @wasp.sh/wasp-cli@latest` install one-liner and a pointer to the quick-start docs page.
11. **Have a Wasp app deployed? - we will send you swag!** — repeat of the Typeform swag offer, this time with a paragraph.
12. **AI Agent Plugins** — pitches the Wasp Agent Plugins doc page for Cursor / Claude Code integration.
13. **Project status** — "Currently, Wasp is in beta," admits ongoing breaking changes, points to development roadmap (GitHub Projects). Notes the currently supported stack is React + TanStack Query, Node.js + Express.js, Prisma.
14. **Contributing** — directs contributors to `waspc/` for compiler details, mentions core is Haskell but there are non-Haskell parts, includes non-code contribution suggestions (star, email list, roadmap).
15. **Careers** — link to Notion careers page.
16. **Sponsors** — six current sponsor entries, each with GitHub avatar thumbnail + first-person thank-you note ("Our first sponsor ever! Thanks so much, Michel ❤️, from the whole Wasp Team, for bravely going where nobody has been before :)!").
17. **Past sponsors** — three avatar thumbnails without thank-you text.

## Tone & style

- **Register**: Pragmatic + warm-startup — Rails references, "bravely going where nobody has been before :)," smiley faces in prose, emphasis on "Team" voice.
- **Voice**: First-person plural (strongly: "we are currently focusing," "we'd love to send some swag your way," "we are thankful for your support"); second-person imperative for instructions.
- **Formatting**: Centered HTML `<div align=center>` header, emoji in feature bullets and in prose, no collapsibles, two code blocks (`.wasp` DSL + Prisma schema), architecture diagram image, avatar-thumbnail image grid for sponsors.

## Notable conventions

- **"Perfect for AI by design" as a first-class feature bullet** and a dedicated "AI Agent Plugins" section — unusual for a pre-AI-era framework to retrofit this as a top-line selling point; signals Wasp has read the room on coding-agent workflows.
- **Swag pitch repeated twice** — once as a link in the header and once as a full section with a Typeform link; the "Deployed? Get swag! 👕" inline link is a clever deployment-funnel capture disguised as a perk.
- **Sponsor section is narrative, not mechanical** — each sponsor gets a personalized thank-you sentence (not just an avatar grid). This signals a small community where the maintainers can remember who supported them, and it's a deliberate warmth signal.
- **"Careers" section in the README** — unusual; most OSS projects don't recruit from the README. Confirms Wasp is a venture-backed company (Wasp Team / Notion careers page) and that contributor-to-hire funnel is an explicit goal.
- **Transparent about Haskell core** — "The core of Wasp is built in Haskell" is surfaced in the Contributing section, potentially gating contributor interest but also honest about the stack mismatch between Wasp-the-product (JS/TS users) and Wasp-the-compiler (Haskell authors).
- **"Project status" section admits beta and ongoing breakage** — honest rather than minimized ("you can expect numerous changes and improvements in the future"), despite the prominent "single CLI command deployment" pitch above.

## Takeaway pattern

This is the "funded-OSS startup front page" README: balances a tight product pitch (etymology → Rails comparison → code example → architecture diagram) with community-company fusion signals (narrative sponsors, Careers link, personalized swag funnel, named "Wasp Team"). It works because the technical story is substantive (the `.wasp` DSL snippet does the pitching job that five feature bullets can't) and the warmth is earned (named sponsors, not anonymized tiers). Lesson: venture-backed OSS can be upfront about being a company without losing community credibility if the sponsor/careers/swag scaffolding is humanized rather than corporate-tiered, and if the project status section is honest about beta state.
