# Analysis: Prinzhorn/skrollr

- **Repo**: Prinzhorn/skrollr (~18.9k stars)
- **Type**: Stand-alone JavaScript parallax/scroll-keyframe animation library; ~12 KB minified, no dependencies.
- **Length**: ~28 KB. Very long; this README is effectively the manual — no separate docs site.

## Sections

1. **Travis CI badge** — single Travis badge (old-school, no shields.io).
2. **"Please note:" abandonment disclaimer** — bolded statement that "skrollr hasn't been under active development since about September 2014," points at the contributors graph as evidence, warns about mobile browser drift, and ends with an editorial aside: "mobile support always sucked (because mobile browsers are hard) and you shouldn't compromise UX for some fancy UI effects. Ever."
3. **skrollr 0.6.30 header** — underline-style `====` h1, tagline "Stand-alone **parallax scrolling** JavaScript library", "Designer friendly. No JavaScript skills needed." Followed by a self-aware italicized paragraph admitting "I wanted to sound hip and use some buzz-words."
4. **Resources / Plugins** — Official (skrollr-menu, skrollr-ie, skrollr-stylesheets) and Third party (skrollr-colors, skrollr-decks) bullet lists.
5. **Resources / In the wild** — points at the wiki for a showcase and a separate "Agencies and freelancers" wiki page for paid skrollr support.
6. **Documentation / Abstract** — explains the core premise ("animate any CSS property of any element depending on the horizontal scrollbar position"), argues against JS-driven animation libraries, and recommends ScrollMagic + jQuery + GSAP as the alternative if you want JS-defined animations. Genuinely pitches a competitor.
7. **Let's get serious** — `<script>` include example, require.js variant, and four incremental HTML examples (background-color interpolation, "barrel roll" transform, bouncing easing, relative-mode `data-top`), each followed by a "View in browser" link and a `##### Lessons learned` bullet list.
8. **Mobile support** — a "The Problem with mobile and the solution" essay explaining why mobile browsers throttle JS during scroll and how skrollr fakes it with CSS transforms on `#skrollr-body`.
9. **AMD** — half-sentence mention with code.
10. **Absolute vs relative mode** — syntax reference `data-[offset]-[anchor]` and `data-[offset]-(viewport-anchor)-[element-anchor]`, with bulleted examples plus a PDF infographic link.
11. **Percentage offsets, Hash navigation, Working with constants, Dynamic constants** — small reference subsections.
12. **CSS classes** — documents `skrollr`, `no-skrollr`, `skrollr-desktop`, `skrollr-mobile`, `skrollable`, `skrollable-before/between/after`.
13. **Animating attributes, Filling missing values, Preventing interpolation** — quirk documentation with example SVG polygon animation and the `!url(...)` escape syntax.
14. **Limitations** — five-bullet enumeration of what skrollr can't animate (unit mismatches, differing transform functions, named/hex colors, etc.).
15. **JavaScript / skrollr.init([options])** — the long options reference: `smoothScrolling`, `smoothScrollingDuration`, `constants`, `scale`, `forceHeight`, `mobileCheck`, `mobileDeceleration`, `skrollrBody`, `edgeStrategy` (with worked example), `beforerender`, `render`, `keyframe` (flagged experimental), `easing` (with built-in list and a Flash-era polynomial-coefficient generator link).
16. **Public API** — `refresh`, `relativeToAbsolute`, `getScrollTop`, `getMaxScrollTop`, `setScrollTop`, `isMobile`, `animateTo`, `stopAnimateTo`, `isAnimatingTo`, `on`/`off`, `destroy`.
17. **Changelog** — single-line pointer to HISTORY.md.

## Tone & style

- **Register**: Conversational, jokey, Web-2.0-blog-era — self-deprecating asides, quips like "-moz-relax and get yourself a cup of -webkit-coffee," and parenthetical "(trolololol)" inside a code comment for a random-number constant.
- **Voice**: First-person singular ("I was lying to you," "I wanted to sound hip," "Originally I wanted to emit the events right on the element... but IE."); addresses the reader directly as "you."
- **Formatting**: Setext-style `====` / `-----` headings (not `#`), nested `#####` "Lessons learned" bullets, old Travis PNG badge, no emoji outside occasional `;-)`, PDF infographic link.

## Notable conventions

- **Abandonment notice above the project title** — the very first substantive section is "**skrollr hasn't been under active development since about September 2014**" with a reasoned opinion ("you shouldn't compromise UX for some fancy UI effects. Ever."). Honest maintenance status communicated before the pitch.
- **"I was lying to you"** section opener for Working with constants — a syntax-retraction device (the earlier "the syntax is `data-[offset]-[anchor]`" was a simplification) that treats the reader as a friend reading a blog post rather than a docs consumer.
- **Explicit competitor recommendation** — "If you prefer to use JavaScript to define your animations make sure to take a look at [ScrollMagic]" in the Abstract, right where a typical README would assert its own superiority.
- **Setext-style headings** (`====` / `-----`) rather than ATX (`#`) throughout — a stylistic tell of a README that predates the modern markdown convention.
- **"But IE." as a terminating explanation** for why events can't be attached to elements — a one-word aside that the audience of that era understood completely.
- **PDF infographic (not SVG/PNG) for the anchor-position guide** — a Flash-adjacent design artifact; same era as the linked `timotheegroleau.com/Flash/experiments/easing_function_generator.htm`.
- **Agencies-and-freelancers wiki link** — an explicit commercial-support board for paid skrollr work, a pre-Stripe-sponsorship way to route monetization.

## Takeaway pattern

This is the "mature one-author JS library from the jQuery era" README: a single-voice manual that doubles as the complete documentation (no separate docs site), written as a friendly blog post with retractions and jokes, and prefaced with an adult-tone abandonment disclaimer. Lesson: when a long-tail library stops being actively developed, the most honest README update is a visible pinned notice above the fold, a recommended alternative (here, ScrollMagic), and a reasoned editorial opinion about why the problem space is harder than buzzword-hungry consumers assume — all of which skrollr does in its first 400 words. It also illustrates how "README as book" was the norm before the docs-site era, and how a conversational voice ("I was lying to you") can carry a 28 KB spec where formal tone would grind.
