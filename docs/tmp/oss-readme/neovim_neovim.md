# Analysis: neovim/neovim

- **Repo**: neovim/neovim (~98.9k stars)
- **Type**: Hyperextensible Vim-based text editor
- **Length**: ~5.3 KB

## Sections

1. Centered H1 with logo + Docs / Chat nav links
2. Badges (Coverity, Repology packages, Debian CI, downloads)
3. Mission paragraph: bullets describing why Neovim exists (simplify maintenance, split work, enable UIs, maximize extensibility)
4. Features — bulleted list with links to GUIs, API clients, terminal emulator, etc.
5. Install from package (releases + managed package links for every major distro)
6. Install from source (CMake + make)
7. Transitioning from Vim (single link)
8. Project layout (ASCII tree with per-subdirectory descriptions)
9. License (Apache 2.0 since a specific commit, with `vim-patch` contributions under Vim license)

## Tone & style

- **Register**: technical, project-veteran. No hype.
- **Voice**: third-person, neutral.
- **Formatting**: Setext h2 underlines (`--------`); heavy use of link references at the bottom of the file.

## Notable conventions

- Mission before features — the project defines *why* before *what*.
- Explicit package-manager link list rather than a one-liner install — respects users on many platforms.
- ASCII project-layout tree in README (instead of relegated to CONTRIBUTING.md) — signals "newcomers, here's the map."
- License note carefully names the commit after which the dual-licensing applies.

## Takeaway pattern

A README for a fork/rewrite project: it must *justify its existence* against the parent (Vim), which it does with the bulleted mission statement. Otherwise very pragmatic. The project-layout ASCII tree is a distinguishing habit in C projects.
