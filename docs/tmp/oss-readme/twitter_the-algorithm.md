# Analysis: twitter/the-algorithm

- **Repo**: twitter/the-algorithm (~73k stars)
- **Type**: Open-source dump of X/Twitter's recommendation algorithm
- **Length**: ~7.3 KB

## Sections

1. H1 title + two-paragraph framing that links to the engineering blog post
2. Architecture — top-level table categorizing components as Data / Model / Software framework, each a row with a link and description
3. For You Timeline subsection with a system-diagram image and a per-component table (Candidate Source / Ranking / Post mixing)
4. Recommended Notifications subsection (second component table)
5. Build and test code — a short disclaimer that BUILD files exist but there is no top-level WORKSPACE
6. Contributing (HackerOne for security, blog link)

No install, no examples, no license (at the README level), no badges.

## Tone & style

- **Register**: corporate-engineering, detached. No hype; also no warmth.
- **Voice**: third-person; describes the system in architectural terms.
- **Formatting**: dominated by tables. Almost no prose.

## Notable conventions

- README treats the repo as reference architecture rather than something to clone and run. This is honest — the build scaffolding is incomplete.
- Components have inline percentages (e.g., "~50% of posts come from this candidate source") — unusually quantitative for a README.
- Security routed to a public bug-bounty program.

## Takeaway pattern

The "transparency artifact" README: code is shared for inspection, not execution. The README's job is to orient readers in the system, not to help them build it. Tables + diagram + component descriptions is the right structure for that use case.
