# Analysis: openai/evals

- **Repo**: openai/evals (~18.2k stars)
- **Type**: Python framework + registry for evaluating LLMs (OpenAI-published, model-graded and code-graded evals).
- **Length**: ~6.3 KB. Compact; setup-and-pointer-heavy rather than reference-heavy.

## Sections

1. **H1 + blockquote callout** — `# OpenAI Evals` followed immediately by `> You can now configure and run Evals directly in the OpenAI Dashboard. [Get started →]` — redirecting casual users to the hosted product.
2. **Intro paragraphs** — defines what Evals is (framework + registry), positions private evals for proprietary data, and argues for eval importance via a quoted endorsement from "OpenAI's President Greg Brockman" (linked tweet).
3. **Embedded Brockman tweet image** — inlined screenshot of the referenced tweet (`<img width="596" alt="…" src="github.com/openai/evals/assets/…">`).
4. **Setup** — OpenAI API key env var (`OPENAI_API_KEY`), cost warning, Weights & Biases alternative, "Minimum Required Version: Python 3.9".
5. **Downloading evals** (subsection of Setup) — Git-LFS required; `git lfs fetch --all` / `git lfs pull`, and an `--include=evals/registry/data/${your eval}` pattern for fetching a single eval.
6. **Making evals** (subsection of Setup) — clone-and-editable-install (`pip install -e .`), optional `[formatters]` extra, `pre-commit install` workflow.
7. **Running evals** — `pip install evals` for users who only want to run evals; pointers to `docs/run-evals.md` and `docs/eval-templates.md`; mentions Completion Function Protocol at `docs/completion-fns.md`; optional Snowflake logging with `SNOWFLAKE_ACCOUNT`/`SNOWFLAKE_DATABASE`/`SNOWFLAKE_USERNAME`/`SNOWFLAKE_PASSWORD` env vars.
8. **Writing evals** — four bulleted pointers (build-eval.md, custom-eval.md, completion-fns.md, external cookbook link); PR submission policy; **explicit note that custom-code evals are not currently accepted**, only model-graded YAML evals.
9. **FAQ** — four question/answer pairs in a question-as-heading style: start-to-finish example (`examples/` folder), same-eval-multiple-ways (`coqa.yaml`), hang-at-end known issue ("you should be able to interrupt it safely"), and a playful "I am a world-class prompt engineer. I choose not to code" entry explaining YAML-only contribution.
10. **Disclaimer** — MIT licensing of contributions, rights assertion for uploaded data, statement that **"OpenAI reserves the right to use this data in future service improvements to our product,"** and a link to usage policies.

## Tone & style

- **Register**: Institutional / pragmatic — reads like internal documentation that was published; very little marketing language.
- **Voice**: Mostly first-person plural ("We offer an existing registry," "We suggest getting started by…") with occasional second-person direct address in the FAQ.
- **Formatting**: Very few badges (none at top), one embedded tweet screenshot, inline `sh`/`env` fenced blocks, blockquote used only for the Dashboard redirect at top. No tables, no collapsibles, no emoji section markers.

## Notable conventions

- **Top-of-readme redirect to the hosted product** — the first thing after the H1 is "you can now do this in the Dashboard," actively steering non-contributors off the repo.
- **Executive endorsement as justification** — the inline Brockman tweet screenshot is used as rhetorical weight to answer "why should I care about evals" rather than a technical argument.
- **Explicit restriction on accepted contribution types** — "we are currently not accepting evals with custom code!" is a policy statement rare in inviting-contributions READMEs; reframes the repo as a curated registry rather than an open framework.
- **Disclaimer retains rights over contributed data** — "OpenAI reserves the right to use this data in future service improvements" puts a clear commercial-use flag on PRs; most OSS READMEs don't surface this at all.
- **FAQ uses raw question text as headings** (including conversational framings like "I am a world-class prompt engineer. I choose not to code.") — unusual voice for a corporate-authored README.
- **Git-LFS requirement is framed as normal** — first technical instruction after setup is `git lfs fetch --all`, which filters users who aren't comfortable with LFS.

## Takeaway pattern

This is the "curated registry masquerading as a framework" README: OpenAI publishes the tooling but treats the repo as a governed dataset (LFS-gated, custom-code PRs refused, contributor-data rights reserved) and gently redirects casual users to the hosted Dashboard. It works for its actual purpose — soliciting YAML-and-JSON eval contributions that feed internal model evaluation — while being honest about the restrictions. Lesson: when a repo's primary purpose is collecting contributions rather than enabling downstream use, the README should say so plainly (contribution policy, data rights, hosted alternative), even if that shrinks the contributor funnel.
