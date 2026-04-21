# Analysis: neoclide/coc.nvim

- **Repo**: neoclide/coc.nvim (~25.1k stars)
- **Type**: LSP client / intellisense engine for Vim/Neovim
- **Length**: ~112.5 KB — by far the longest in the sample

## Sections

1. Centered logo + tagline ("Make your Vim/Neovim as smart as VS Code")
2. Badges (Anti-996 license, CI, Codecov, docs, DeepWiki)
3. Custom popup screenshot
4. Why? — four emoji-led bullets (🚀 Fast, 💎 Reliable, 🌟 Featured, ❤️ Flexible)
5. Quick Start — Vim/Neovim version requirements, Node install, vim-plug install, extension install, language-server config
6. Wiki pointer list
7. Example Vim configuration — an extremely long reference config (~hundreds of lines of Vimscript) covering completion, diagnostics, code actions, refactoring, formatting, rename, selection, document symbols, etc.
8. (File continues well past 100 KB with feature docs, FAQs, troubleshooting, etc.)

## Tone & style

- **Register**: power-user focused, technically dense.
- **Voice**: second-person imperative during config steps.
- **Formatting**: lots of fenced `vim` blocks; emoji markers on headline bullets; emphatic `❗️` markers on warnings.

## Notable conventions

- README effectively embeds a reference configuration — discouraged in most projects but here justified because Vim users *expect* to paste-and-adapt config.
- Uses the Anti-996 license badge (a non-standard political license). A distinctive governance signal.
- DeepWiki badge ("Ask DeepWiki") shows up — evidence of LLM-generated documentation becoming a first-class README element.
- The very length of the README is its own signal: this is a dense, deep tool and the maintainers expect you to live inside the file.

## Takeaway pattern

The "paste-this-config" README: standard advice is to keep configs in a separate doc, but for Vim-ecosystem projects, putting the full example config in the README reduces friction more than it clutters. An outlier that works because its audience is very specific.
