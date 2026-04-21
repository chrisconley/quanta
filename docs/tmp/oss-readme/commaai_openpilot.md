# Analysis: commaai/openpilot

- **Repo**: commaai/openpilot (~60.7k stars)
- **Type**: Driver-assistance / robotics OS
- **Length**: ~7.2 KB

## Sections

1. Centered hero: H1, bolded tagline ("openpilot is an operating system for robotics.")
2. Centered nav links (Docs · Roadmap · Contribute · Community · Shop)
3. One-line quick-start (`bash <(curl …)`)
4. Badges (CI, license, Twitter, Discord)
5. Video grid (3 embedded YouTube thumbnails as images)
6. "Using openpilot in a car" (numbered requirements list)
7. Branches (table of release/nightly branches)
8. "To start developing openpilot" (developer onboarding bullets; includes a hiring/bounties plug)
9. Safety and Testing (bulleted list citing ISO26262, SIL tests, HIL tests, panda repo)
10. `<details>` MIT license + indemnity + ALPHA-QUALITY warning
11. `<details>` user data policy + privacy disclosure

## Tone & style

- **Register**: engineering-serious but confident; mixes marketing and regulatory caution.
- **Voice**: third-person, with strategically bold lines ("YOU ARE RESPONSIBLE FOR COMPLYING WITH LOCAL LAWS").
- **Formatting**: HTML for the hero and media grid; Markdown for the body; collapsible `<details>` sections to hide legalese without removing it.

## Notable conventions

- Safety section reads like a product-quality claim (ISO26262, HIL, 10-device regression closet) — essentially a trust-building artifact.
- Explicitly labels the software "ALPHA QUALITY FOR RESEARCH" in capitalized bold, inside a collapsed section.
- Commercial product links (shop, jobs, bounties) are intermingled with open-source onboarding — signals dual identity.

## Takeaway pattern

A README for a high-risk open-source project: must balance pitch (videos, shop link), onboarding (docs, branches), and disclaimers (safety claims, indemnity). The collapsible legal sections are a good pattern for keeping the readable body short while still making the legalese linkable.
