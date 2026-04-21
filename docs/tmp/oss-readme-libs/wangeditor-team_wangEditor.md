# Analysis: wangeditor-team/wangEditor

- **Repo**: wangeditor-team/wangEditor (~18.2k stars)
- **Type**: Open-source web rich-text editor with JS / Vue / React bindings (Chinese-origin project; primary docs in Chinese).
- **Length**: ~0.3 KB. Tiny — a stub README that exists mostly to point elsewhere.

## Sections

1. **H1 title** — "wangEditor 5" (version in the title).
2. **Language switch** — single-line Markdown link `[English](./README-en.md)` at the top.
3. **介绍 (Introduction)** — one line of Simplified Chinese describing the project as "开源 Web 富文本编辑器，开箱即用，配置简单。支持 JS Vue React" (open-source web rich-text editor, out-of-the-box, simple to configure; supports JS, Vue, React), plus two links to documentation and demo on `wangeditor.com` and an editor screenshot `![](./docs/images/editor.png)`.
4. **交流 (Discussion)** — single bullet pointing to GitHub Issues for "讨论问题和建议" (discussing questions and suggestions).
5. **捐赠 (Donations)** — one line pointing to `opencollective.com/wangeditor`.

## Tone & style

- **Register**: Utility-stub. No salesmanship, no feature list, no badges of any kind. Reads more like a placeholder index than a README.
- **Voice**: None — the README has no first- or second-person voice; everything is nominal labels and links.
- **Formatting**: Plain Markdown. Four-character section headers in Chinese ("介绍", "交流", "捐赠") rather than translated ("Introduction", "Contact", "Donate"); one screenshot with no alt text (`![](./docs/images/editor.png)`).

## Notable conventions

- **Chinese-first with English as a side door** — the entire README is Chinese with a single `[English](./README-en.md)` link at the top; this inverts the more common "English default with `README.zh-CN.md` alongside" pattern and signals where the primary audience is.
- **Version in the H1** — "wangEditor 5" rather than "wangEditor"; the major version is part of the branding, which tracks for a project whose v4 and v5 are known to be API-incompatible.
- **No badges** — no CI, no version, no license, no download count. For a project with ~18k stars and a v5 rewrite, the absence is deliberate rather than accidental.
- **Donations section without a sponsorship narrative** — a standalone "捐赠" section pointing to Open Collective, with no explanation of what donations fund. Reads as an understated ask.
- **Three sections, all single-line** — the README's entire Chinese-language payload is three sentences. All real content lives at `wangeditor.com` (documentation, demo) or on GitHub Issues (community).
- **Screenshot as the only visual** — no logo, no diagram, just an editor screenshot from `docs/images/editor.png` rendered with empty alt text.
- **No license reference** — the word "license" / "协议" does not appear; readers have to open the `LICENSE` file blindly.
- **Missing "Installation" / "Quick Start"** — no `npm install`, no `<script src=...>` tag, no `import` example. For a drop-in editor, this is a striking omission and pushes all first-run concerns to the docs site.

## Takeaway pattern

This is the "docs-site-is-the-real-README" archetype, in its minimal non-English form: the GitHub README is a marketing-free navigation stub that delegates onboarding, installation, API, and examples entirely to `wangeditor.com`. It works as a signpost when readers arrive expecting a Chinese-language editor with a Chinese docs site, and the single English link handles spillover. It fails any drive-by discovery use case — a reader who stumbles into the repo via "rich text editor React" gets no version info, no install command, no license signal, and no feature list to compare against TinyMCE / Quill / TipTap. Lesson: when your true documentation lives on a branded external site and your primary audience reads your language, the README can shrink to a three-line stub — but recognize that you've traded discoverability for that compression, and decide whether that's the trade you want.
