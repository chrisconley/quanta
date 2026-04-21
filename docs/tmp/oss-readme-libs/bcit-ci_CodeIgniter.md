# Analysis: bcit-ci/CodeIgniter

- **Repo**: bcit-ci/CodeIgniter (~18.3k stars)
- **Type**: PHP Application Development Framework (MVC toolkit) — specifically the legacy **CodeIgniter 3** branch in maintenance mode.
- **Length**: ~2.3 KB. Very short; reStructuredText-formatted rather than markdown.

## Sections

1. **What is CodeIgniter** — three-sentence value-proposition paragraph using the framed `###################` reST heading style; describes CI as "an Application Development Framework - a toolkit - for people who build web sites using PHP," promises faster delivery than coding from scratch via "a rich set of libraries for commonly needed tasks."
2. **CodeIgniter 3** — explicit legacy-branch notice: points at `CodeIgniter4` as the latest version, states that CI3 targets PHP 5.6+, and declares the branch "in maintenance, receiving mostly just security updates."
3. **Release Information** — one sentence saying the repo contains in-development code, with a link to `codeigniter.com/download` for stable releases.
4. **Changelog and New Features** — one sentence linking the user-guide changelog file at `user_guide_src/source/changelog.rst` on the `develop` branch.
5. **Server Requirements** — "PHP version 5.6 or newer is recommended" plus a paragraph warning against using the compatibility floor ("It should work on 5.4.8 as well, but we strongly advise you NOT to run such old versions of PHP").
6. **Installation** — one-line pointer to the user guide's installation section; zero inline install steps.
7. **License** — one-line pointer to the license file at `user_guide_src/source/license.rst`.
8. **Resources** — 6-bullet external link list: User Guide, Contributing Guide, Language File Translations, Community Forums, Community Wiki, Slack channel; followed by a two-line security-reporting notice pointing at `security@codeigniter.com` and a HackerOne page.
9. **Acknowledgement** — three-line closer thanking EllisLab (the original corporate owner before the BCIT handoff), all contributors, and "you, the CodeIgniter user."

## Tone & style

- **Register**: Formal-institutional and oddly warm — reads like a university lab's project README (BCIT = British Columbia Institute of Technology), which it is.
- **Voice**: Third-person descriptive ("CodeIgniter is..."), first-person plural only in the security warning and the closing Acknowledgement ("we strongly advise you NOT to run...", "The CodeIgniter team would like to thank...").
- **Formatting**: reStructuredText rather than markdown — heading underlines use `###`, `***`, and `===`, links are `` `text <url>`_ `` form, and the whole file would render differently on GitHub than a `.md`. No code blocks, no badges, no images.

## Notable conventions

- **reStructuredText in a README named `README.md`** (implied by the source) — the choice signals an older PHP-era convention where reST was used for Sphinx-compatible user guides; the framework ships its user guide in reST, and the README matches.
- **Legacy-branch disclaimer in the second section, not the first** — the "What is CodeIgniter" paragraph sells the framework as if it were current; only the next section reveals this is CI3 (legacy) and you probably want CI4. Compare to Masonry which puts the "use SnapKit instead" redirect in the very first paragraph.
- **Zero inline install or quickstart** — Installation is a single link to the user guide; the README refuses to duplicate docs. Combined with the lack of code examples, the document reads as a pointer page, not a tutorial.
- **Security reporting listed as a sub-paragraph under Resources** — the email + HackerOne link are presented as a follow-up sentence after the community links bullet list, not as a standalone "Security Policy" section. Reads as older-style before `SECURITY.md` became GitHub convention.
- **Acknowledgement to EllisLab** — the original commercial owner (EllisLab transferred CI to BCIT in 2014) is still credited in the closing section years later; a deliberate historical record, not a dated artifact.
- **"PHP 5.6+" as the baseline in a document still current** — surfaces the project's lifecycle more honestly than the "legacy" label alone; a maintainer who advises against PHP 5.4 while accepting PHP 5.6 is running a security-maintenance-only shop.

## Takeaway pattern

This is the "institutional maintenance notice" archetype: a README whose primary job is to state what version lives in this repo, what PHP versions it targets, where to find the real docs, and whom to email about security — no marketing, no examples, no badges. It works because CodeIgniter's audience arrives already knowing the framework (it's 18k stars because of its 2010s heyday, not current adoption) and just needs to confirm which branch this is. Lesson: when your project is in long-tail maintenance mode and documentation lives elsewhere, a short reST-or-markdown pointer page plus an explicit "this is the legacy version, see X for the current one" notice serves drive-by readers better than a pretty README that pretends the project is still on the growth curve — but put the legacy notice first, not second.
