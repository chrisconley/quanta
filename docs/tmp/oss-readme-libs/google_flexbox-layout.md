# Analysis: google/flexbox-layout

- **Repo**: google/flexbox-layout (~18.2k stars)
- **Type**: Android library porting CSS Flexible Box layout capabilities to Android ViewGroups and RecyclerView (`FlexboxLayout` and `FlexboxLayoutManager`).
- **Length**: ~14 KB. Long; comprehensive attribute reference plus a feature-comparison table and migration notes.

## Sections

1. **H1 title + Circle CI badge** — `# FlexboxLayout` with a single CircleCI shield.
2. **Elevator pitch** — two-line definition linking to the W3C CSS Flexible Box Module spec.
3. **Installation** — Gradle `implementation 'com.google.android.flexbox:flexbox:3.0.0'` block; bolded migration notice about the 3.0.0 groupId change from `com.google.android` to `com.google.android.flexbox`; a second paragraph flagging the 2.0.0 default-value change for `alignItems` / `alignContent` (stretch → flex_start) as a breaking behavioral change; a third paragraph about AndroidX migration starting 1.1.0.
4. **Usage** — intro line: "two ways of using Flexbox in your layout."
5. **FlexboxLayout subsection** — XML example with `<com.google.android.flexbox.FlexboxLayout ... app:flexWrap="wrap" app:alignItems="stretch">` and three TextView children using `layout_flexBasisPercent`, `layout_alignSelf="center"`, `layout_alignSelf="flex_end"`; followed by a Java equivalent using `setFlexDirection(FlexDirection.ROW)`, `setOrder(-1)`, `setFlexGrow(2)`.
6. **FlexboxLayoutManager (within RecyclerView) subsection** — Java example constructing `FlexboxLayoutManager(context)`, setting `FlexDirection.COLUMN` and `JustifyContent.FLEX_END`; second snippet showing how to cast `ViewGroup.LayoutParams` to `FlexboxLayoutManager.LayoutParams`; rationale paragraph about view recycling; animated GIF `flexbox-layoutmanager.gif`.
7. **Supported attributes/features comparison** — 15-row markdown table with `FlexboxLayout` vs `FlexboxLayoutManager (RecyclerView)` columns; check marks rendered as `check_green_small.png` images; `alignContent`, `layout_order`, and View recycling are the headline asymmetries; Scrolling footnoted as `*1` with caveats.
8. **Supported attributes — FlexboxLayout parent** — six attributes with prose definitions, possible-values bullet lists, and an explainer GIF for each (flex-direction.gif, flex-wrap.gif, justify-content.gif, align-items.gif, align-content.gif), plus the divider attribute trio (`showDividerHorizontal`/`showDividerVertical`/`showDivider` + `dividerDrawableHorizontal`/`Vertical`/`Drawable`) with a full example layout + `divider.xml` shape + rendered PNG.
9. **Supported attributes — children** — eight per-child attributes (`layout_order`, `layout_flexGrow`, `layout_flexShrink`, `layout_alignSelf`, `layout_flexBasisPercent`, `layout_minWidth/Height`, `layout_maxWidth/Height`, `layout_wrapBefore`), each with default values, a prose definition, and an animated GIF.
10. **Others → Known differences from the original CSS specification** — numbered list (1)-(5) enumerating intentional deviations: no `flex-flow`, no `flex` shorthand, `layout_flexBasisPercent` in place of `flex-basis`, `layout_wrapBefore` introduced, and default `alignItems`/`alignContent` set to `flex_start` instead of `stretch` (with a performance rationale: "stretch ... is expensive because the children ... are measured more than twice").
11. **Xamarin Binding** — one-line credit to `@btripp` with a NuGet link.
12. **Demo apps** — `demo-playground` and `demo-cat-gallery` modules, each with its `./gradlew <module>:installDebug` command; the cat-gallery description explicitly compares to Google Photos and flags the OutOfMemoryError risk avoided by `FlexboxLayoutManager`.
13. **How to make contributions** — one line pointing to CONTRIBUTING.md.
14. **License** — one line pointing to LICENSE.

## Tone & style

- **Register**: Technical-reference, Android-engineering flavor. Patient, explain-every-default voice; the spec-comparison section is especially methodical.
- **Voice**: Third-person project voice ("This library tries to achieve...", "A child view won't shrink less than..."), with second-person instructions in the usage section ("You can specify the attributes from a layout XML like:").
- **Formatting**: Bolded callouts with `**...**` for breaking-change notices in Installation; `__attribute__` (double-underscore) used for attribute names in the reference sections, an unusual-in-Markdown emphasis choice that renders bold but reads as "code-ish"; heavy reliance on animated GIFs under `/assets/` (one per attribute) and image-based checkmarks (`check_green_small.png`) in the comparison table instead of Unicode checks; XML fenced blocks for layout resources and Java fenced blocks for imperative setters.

## Notable conventions

- **Image checkmarks in a feature-comparison table** — instead of `✅` or `✓`, the comparison table uses `![Check](/assets/pngs/check_green_small.png)`; the asymmetric gaps (` - ` for unsupported) become visually prominent. A pre-emoji-era convention preserved into the present.
- **Per-attribute GIF policy** — every attribute gets a dedicated looping demo GIF in `/assets/`, even for simple ones like `layout_minWidth`. This is closer to a product-marketing approach than a typical Android library README.
- **Breaking-change callouts inline in Installation** — three consecutive paragraphs flagging different migration-relevant changes (3.0.0 groupId, 2.0.0 default-value flip, 1.1.0 AndroidX). The typo "expeced" in the AndroidX note ("starting from 1.1.0, the library is expeced to use with AndroidX") has survived.
- **"Known differences from the original CSS specification" section** — a dedicated, numbered audit of where the Android port intentionally diverges from `https://www.w3.org/TR/css-flexbox-1`. Explicitly addresses the reader who already knows CSS Flexbox and wants to know what's missing — the fastest onboarding path for that audience. The section also offers rationale (e.g., shorthand attributes are "not practical" in Android XML; `stretch` is "expensive because the children ... are measured more than twice").
- **Double-underscore attribute emphasis** — `__flexDirection__` instead of `**flexDirection**` or `` `flexDirection` `` renders as bold but carries an odd typographic flavor; consistent throughout the attribute reference, so probably deliberate-but-old.
- **Two demo modules named explicitly** — `demo-playground` (attribute tryout) and `demo-cat-gallery` (production-style RecyclerView showcase), each with a ready-to-run Gradle task. Unusual to find two working demo apps called out by name.
- **Xamarin binding credited to an outside contributor in its own section** — `@btripp`'s NuGet package gets a mini-section rather than a buried acknowledgement, a generous pattern for a Google-owned repo.
- **No badges beyond CircleCI** — no Maven Central version, no license, no star/download metrics. Leaves ambient credibility ("google/") to do that work.

## Takeaway pattern

This is the "Google-engineered reference manual" README archetype: a single-page, completeness-oriented document that tries to be the attribute reference, the migration guide, the CSS-vs-Android-flexbox diff, and the playground index all at once. It works because the surface area of CSS Flexbox is small and finite — ~14 attributes — so exhaustive coverage is tractable, and the one-GIF-per-attribute policy makes visual learners productive without leaving the page. It wobbles where maintenance has slipped (the "expeced" typo, the ~2019 AndroidX phrasing, setext-era double-underscore emphasis) — the README has clearly been in low-churn mode for years. Lesson: when your library is a port of a well-specified external standard, the highest-leverage section you can write is "Known differences from the original specification" — it converts every reader who already knows the parent standard into a productive user in one scroll.
