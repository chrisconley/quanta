# Analysis: knadh/listmonk

- **Repo**: knadh/listmonk (~19.6k stars)
- **Type**: Self-hosted newsletter / mailing list manager
- **Length**: ~2.1 KB. Compact.

## Sections

1. Sponsor banner (Zerodha) — right-floated HTML image
2. Logo (clickable)
3. One-paragraph pitch: "standalone, self-hosted … single binary … PostgreSQL"
4. Product screenshot
5. Live demo link
6. Installation
   - Docker subsection (compose file)
   - Binary subsection (new-config, install, upgrade flags)
7. Developers (one paragraph — stack and contribution pointer)
8. License

## Tone & style

- **Register**: pragmatic, neutral, sysadmin-friendly.
- **Voice**: third-person with active, imperative install instructions.
- **Formatting**: h2 sections; horizontal `__________________` separators between install methods (an unusual touch that visually reinforces "pick one").

## Notable conventions

- Leads with the product screenshot, which is important for a self-hosted web app — users decide on visual quality.
- Puts the Zerodha sponsor first (via right-floated image), keeping provenance visible without interrupting the text flow.
- Gives only a minimum viable install; everything else lives at `listmonk.app/docs`.

## Takeaway pattern

Classic self-hosted-OSS style: screenshot + two-path install (Docker / binary) + docs link. Short enough that a visitor knows within 10 seconds whether the project fits. Stack mentioned briefly (Go + Vue + Buefy) for would-be contributors.
