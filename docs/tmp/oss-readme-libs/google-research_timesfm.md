# Analysis: google-research/timesfm

- **Repo**: google-research/timesfm (~18.2k stars)
- **Type**: Pretrained decoder-only foundation model for time-series forecasting (Google Research, open weights on Hugging Face).
- **Length**: ~3.7 KB. Short; front-loaded with dated changelog entries, minimal API surface.

## Sections

1. **H1 + tagline + paper/checkpoint/blog/product links** — `# TimesFM` followed by a one-sentence definition ("TimesFM (Time Series Foundation Model) is a pretrained time-series foundation model developed by Google Research for time-series forecasting"), then a four-item bullet list: the ICML 2024 arXiv paper (`2310.10688`), a Hugging Face collection link, a Google Research blog post, and a nested **"TimesFM in Google 1P Products"** sub-list naming BigQuery ML, Google Sheets (via Connected Sheets), and Vertex Model Garden.
2. **Unsupported-product disclaimer** — single sentence in normal prose: "This open version is not an officially supported Google product." No blockquote, no callout box.
3. **Latest Model Version** — single bold line (`**Latest Model Version:** TimesFM 2.5`).
4. **Archived Model Versions** — one bullet explaining 1.0 and 2.0 live in the `v1/` subdirectory and giving a `pip install timesfm==1.3.0` escape hatch.
5. **Update - Apr. 9, 2026** — dated changelog block announcing a HuggingFace Transformers + PEFT (LoRA) fine-tuning example at `timesfm-forecasting/examples/finetuning/`, plus unit tests and community fixes; credits `@kashif` and `@darkpowerxo` by handle.
6. **Update - Mar. 19, 2026** — dated update crediting `@borealBytes` for adding `AGENTS.md` support and announcing `SKILL.md` in `timesfm-forecasting/`.
7. **Update - Oct. 29, 2025** — single sentence: "Added back the covariate support through XReg for TimesFM 2.5."
8. **Update - Sept. 15, 2025** — the launch announcement for 2.5, with a 5-bullet diff vs 2.0 (200M vs 500M parameters, 16k vs 2048 context, optional 30M quantile head, no more frequency indicator, new forecasting flags), followed by a numbered checklist of post-launch improvements with green check emoji (✅) marking each as completed: Flax version, XReg covariate support, docs/examples/agent skill, LoRA fine-tuning example, unit tests.
9. **Install** — three numbered steps: `git clone`, then `uv venv` + `uv pip install -e .[torch]` with `.[flax]` and `.[xreg]` alternatives shown as comments, then an optional "install your preferred `torch` / `jax` backend" step linking pytorch.org and jax.dev.
10. **Code Example** — ~30-line Python fenced block: import, `torch.set_float32_matmul_precision("high")`, `TimesFM_2p5_200M_torch.from_pretrained("google/timesfm-2.5-200m-pytorch")`, a `model.compile(ForecastConfig(...))` with 7 named flags (`max_context=1024`, `max_horizon=256`, `normalize_inputs=True`, `use_continuous_quantile_head=True`, `force_flip_invariance=True`, `infer_is_positive=True`, `fix_quantile_crossing=True`), and a `model.forecast` call with two synthetic inputs plus shape-annotation comments (`# (2, 12)`, `# (2, 12, 10): mean, then 10th to 90th quantiles`).

## Tone & style

- **Register**: Research-engineering, faintly terse — parameter counts and context-length numbers are cited without narrative ("uses 200M parameters, down from 500M. supports up to 16k context length, up from 2048.").
- **Voice**: Third-person for the product description, first-person plural briefly in community shoutouts ("Huge shoutout to..."), otherwise imperative in the install/code sections.
- **Formatting**: Reverse-chronological "Update - Month. Day, Year" H2 section headers (`## Update - Apr. 9, 2026`), bold inline labels (`**Latest Model Version:**`, `**Archived Model Versions:**`), ✅ emoji checklist in the Sept. 15 update, and `uv`-first install rather than plain `pip`.

## Notable conventions

- **Dated update blocks instead of a CHANGELOG** — four consecutive H2 sections titled `## Update - <Date>` stack above the Install section, so a returning reader sees "what's new" in reverse-chronological order without clicking through. The README *is* the changelog for the recent past.
- **Product-of-Google cross-sell sub-list** — the "TimesFM in Google 1P Products" nested bullet naming BigQuery ML, Google Sheets, and Vertex Model Garden positions the OSS repo as feeder/reference implementation for paid Google surfaces; unusually explicit about the commercial relationship.
- **"This open version is not an officially supported Google product" in a single unbold sentence** — Google Research's boilerplate disclaimer is understated here (no blockquote, no caps) but loadbearing; it's the license-adjacent fine print warning enterprise users that this repo is not on Google Cloud's SLA.
- **Community shoutouts by GitHub handle** — `@kashif`, `@darkpowerxo`, `@borealBytes` are named in the update blocks with direct profile links, framing this as an active open-collab repo rather than a code-drop from Google.
- **`uv`-first install instead of plain `pip`** — specifying `uv venv` and `uv pip install -e .[torch]` as the default path is a modern (2024+) Python-tooling choice that signals the maintainers expect a current ML dev environment; plain `pip` is not shown.
- **Checklist with ✅ marking completed roadmap items** — the Sept. 15 update's numbered list is a public roadmap retrospective; each completed item links to the section/directory where the feature lives. Works as both changelog and tour.
- **Shape-annotation comments in the code example** — `point_forecast.shape  # (2, 12)` and `quantile_forecast.shape  # (2, 12, 10): mean, then 10th to 90th quantiles` teach the output schema inside the example rather than in separate API prose.
- **Split-package via extras** — `[torch]`, `[flax]`, `[xreg]` extras-install variants are shown side by side in one code block, letting the reader pick a backend without reading separate install sections.

## Takeaway pattern

This is the "research-release README" archetype: a short document organized around (1) paper + checkpoint links for academic credibility, (2) a stacked reverse-chronological "Update - Date" log replacing a CHANGELOG, and (3) a single copy-pasteable code example that encodes the entire quickstart. It works because ML practitioners land here from arXiv or Hugging Face already knowing what the model does — they need the pretrained-weights link, the install command, and one working `from_pretrained + forecast` snippet, not a tutorial. Lesson: for actively-evolving research code, inlined "Update - <Date>" blocks beat a separate CHANGELOG (returning readers see the deltas without clicking), and a single shape-annotated code example teaches the output schema more efficiently than prose API docs.
