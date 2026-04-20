# CHANGELOG

## v2.0.0 — the rewrite

Full rewrite. Not backward compatible. Different shape, same name.

### philosophy

Kept only what earned its tokens:

- `SPEC.md` — durable, addressable, caveman-encoded
- three commands — `/ck:spec`, `/ck:build`, `/ck:check`
- two skills — `caveman` encoding, `backprop` protocol

### added

- single `SPEC.md` format with six addressable sections (§G §C §I §V §T §B)
- pipe-table encoding for §T (tasks) and §B (bugs)
- caveman symbol set (→ ∴ ∀ ∃ ! ? ⊥ ≠ ∈ ∉ ≤ ≥ & |) as default for spec writes
- bug → §B → §V backprop reflex wired into `/ck:build` failure path
- `/ck:spec from-code` — distill spec from existing codebase
- `/ck:check` — read-only drift report (replaces five v1 review flavors)

### removed

- 13 of 16 v1 commands (sketch/map/make/ship/review/revise/status/init/config/resume/help/design/research/team/make-parallel)
- all 12 named sub-agents
- 19 of 21 skills
- Go binary and source (`cmd/`, `internal/`, `bin/`, `cavekit` executable)
- shell hooks (`hooks/`, `scripts/cavekit-launch-session.sh`, stop-hook state machine)
- TS tooling (`scripts/cavekit-picker.ts`, `scripts/cavekit-router.cjs`)
- Codex peer-review bridge (`.codex-plugin/`)
- `context/kits/`, `context/plans/`, `context/impl/`, `context/refs/` directories
- autonomous loop, per-task budgets, model-tier routing
- design-system `DESIGN.md` workflow
- knowledge-graph `graphify-out/` integration
- parallel wave execution and team mode
- `install.sh` (618 lines → 0)

### changed

- caveman was opt-in for inter-agent chatter in v1; default for spec writes in v2
- version: 3.1.0 → 2.0.0 (deliberate reset — v2 is a new project under the old name)
- README, plugin metadata, marketplace entry

### migration

See [`UPGRADE.md`](./UPGRADE.md). No automated migrator — the v1 kit shape
does not map cleanly to v2's single file. Recommended path: run
`/ck:spec from-code` on your existing v1 project to distill a v2 spec
from your built code.

### v1 reachability

v1 is frozen at tag `v1.3.1-final`. Stays installable and documented.
Fixes only for critical bugs; no new features.

---

## v1.3.1 and prior

See git log before the `v2.0.0` commit, or check out `v1.3.1-final`:

```bash
git checkout v1.3.1-final
```
