# Analysis: SnapKit/Masonry

- **Repo**: SnapKit/Masonry (~18.2k stars)
- **Type**: Objective-C DSL for Apple Auto Layout (iOS + macOS); legacy predecessor to SnapKit (the Swift successor from the same org).
- **Length**: ~16.2 KB. Long-form; effectively a mini-tutorial with running code examples.

## Sections

1. **H1 title with inline badges** — Travis CI, Coveralls, Carthage-compatible, CocoaPods version, all on the same line as the `# Masonry` heading.
2. **Maintenance/successor bold callout** — single bolded paragraph stating Masonry is still maintained but recommending SnapKit for Swift projects.
3. **Intro paragraph** — defines Masonry as a light-weight layout framework wrapping AutoLayout, notes iOS and macOS support, points to the Masonry iOS Examples workspace.
4. **What's wrong with NSLayoutConstraints?** — motivation section: a ~40-line Objective-C code block showing verbose `NSLayoutConstraint constraintWithItem:...` boilerplate for a simple "fill superview with padding" case, plus a paragraph dismissing Visual Format Language as an alternative.
5. **Prepare to meet your Maker!** — the Masonry rewrite of the same example using `MASConstraintMaker`; includes a shorter one-line `make.edges.equalTo(superview).with.insets(padding)` variant. Notes that Masonry auto-adds constraints and sets `translatesAutoresizingMaskIntoConstraints = NO;`.
6. **Not all things are created equal** — the three equality operators (`.equalTo`, `.lessThanOrEqualTo`, `.greaterThanOrEqualTo`) with four argument-type subsections: (1) MASViewAttribute — full markdown table mapping `view.mas_left` → `NSLayoutAttributeLeft` across 11 attributes; (2) UIView/NSView; (3) NSNumber with primitive autoboxing via `mas_equalTo` and `MAS_SHORTHAND_GLOBALS`; (4) NSArray.
7. **Learn to prioritize** — `.priority`, `.priorityHigh`, `.priorityMedium`, `.priorityLow` with blockquote definitions and chain examples.
8. **Composition, composition, composition** — MASCompositeConstraints subsections for `edges`, `size`, `center`, plus chaining (`make.left.right.and.bottom.equalTo(superview)`).
9. **Hold on for dear life** — three numbered strategies for updating constraints: (1) References via stored `MASConstraint *` properties and `uninstall`; (2) `mas_updateConstraints`; (3) `mas_remakeConstraints`.
10. **When the ^&\*!@ hits the fan!** — debugging aid; shows the ugly default "Unable to simultaneously satisfy constraints" console output, then the pretty Masonry version with named constraints like `<MASLayoutConstraint:ConstantConstraint UILabel:messageLabel.height >= 5000>`.
11. **Where should I create my constraints?** — canonical `DIYCustomView` skeleton showing `init`, `+requiresConstraintBasedLayout`, `-updateConstraints`, `-didTapButton:` pattern.
12. **Installation** — CocoaPods one-liner, `MAS_SHORTHAND` macro tip, `#import "Masonry.h"`.
13. **Code Snippets** — Xcode code-snippet templates for `mas_make`, `mas_update`, `mas_remake` with path `~/Library/Developer/Xcode/UserData/CodeSnippets`.
14. **Features** — five one-line bullets positioning against alternatives ("Not limited to subset of Auto Layout," "No crazy macro magic," "Not string or dictionary based").
15. **TODO** — three-item list ("Eye candy," "Mac example project," "More tests and examples").

## Tone & style

- **Register**: Pragmatic with deliberate wisecracking — playful chapter titles ("Prepare to meet your Maker!", "Hold on for dear life", "When the ^&\*!@ hits the fan!") anchor otherwise technical content.
- **Voice**: Mix of first-person plural ("we are committed to fixing bugs") and second-person imperative ("if you want view.left to be greater than…").
- **Formatting**: Heavy use of `obj-c` fenced blocks, one markdown table for MASViewAttribute mapping, blockquotes used for operator definitions (`> .equalTo equivalent to NSLayoutRelationEqual`), inline linked word "orsome" pointing to a YouTube joke (`youtube.com/watch?v=YaIZF8uUTtk`).

## Notable conventions

- **Opens with an explicit "use SnapKit instead" redirect** — rare and honest for a still-maintained project to prominently demote itself; also rare to see a project recommend a different repo from the same org in its first paragraph.
- **Teaches-by-contrast structure** — "What's wrong with X?" before the solution, with a long bad-code block followed by a short good-code block. The NSLayoutConstraint monstrosity is genuinely 40 lines of real Objective-C to sell the DSL.
- **No table of contents, no badges for community** — this is a pre-emoji-era README (`# Masonry` + inline shields is the whole header), and the informal chapter titles double as the nav.
- **"TODO" section left at the bottom** — a public promise list that includes "Eye candy" and "More tests and examples," signaling the project no longer attracts roadmap work (consistent with the Swift successor recommendation up top).
- **Code-snippet installation instructions** — the `~/Library/Developer/Xcode/UserData/CodeSnippets` tip is specific enough that it assumes a committed Xcode user and would read as foreign to anyone who hasn't customized Xcode snippets.
- **Debugging section showcases Objective-C-specific category hacking** — overriding `-[NSLayoutConstraint description]` via a category is the kind of hack readers either find charming or alarming; the README presents it without apology.

## Takeaway pattern

This is the "retiring-but-dignified elder library" README: a mature Objective-C DSL that opens by pointing readers to its Swift successor, then spends ~15 KB teaching the DSL properly for anyone still on the legacy path. It works because the tone is consistent (playful chapter titles, running code-by-contrast examples) and because the maintenance statement up front sets honest expectations. Lesson: when a project is in a "maintained but surpassed" state, lead with that fact rather than burying it — a "we recommend [successor]" paragraph in the first 10 lines preserves the readme's usefulness to the library's remaining audience without misleading new arrivals.
