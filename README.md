<h1 align="center">cavekit</h1>

<p align="center">
  <strong>compressed spec-driven development for claude code</strong><br/>
  <sub>one file · three commands · zero sub-agents</sub>
</p>

---

## what this is

Plan-then-execute forgets. SDD remembers — but most SDD frameworks bury that
value under agent swarms, dashboards, and ceremony that costs more tokens
than it saves.

Cavekit v2 keeps only what earns its place:

- **durable spec** — `SPEC.md` at repo root survives context resets.
- **caveman encoding** — ~75% fewer tokens than prose. Symbols, fragments,
  pipe tables for repeating records.
- **backprop reflex** — every test failure becomes a `§B` entry; classes of
  bug become `§V` invariants the spec never forgets.

That's the whole pitch.

## commands

| cmd | job |
|---|---|
| `/ck:spec` | create / amend / backprop `SPEC.md`. Sole mutator. |
| `/ck:build` | native plan → execute against spec. Auto-backprops on failure. |
| `/ck:check` | read-only drift report. Lists §V / §I / §T violations. |

## install

Add to your Claude Code marketplace sources:

```bash
/plugin marketplace add juliusbrussee/cavekit
/plugin install ck@cavekit
```

Or clone into `~/.claude/plugins/` directly. Commands become available as
`/ck:spec`, `/ck:build`, `/ck:check`.

## format

See [`FORMAT.md`](./FORMAT.md). Fixed sections: §G goal, §C constraints,
§I interfaces, §V invariants, §T tasks (pipe table), §B bugs (pipe table).

## files

```
FORMAT.md             spec schema + caveman encoding rules
commands/             three markdown command prompts
skills/caveman        encoding skill
skills/backprop       bug → spec protocol (six steps)
```

## non-goals

- no sub-agents. Main Claude does the work.
- no dashboards. `cat SPEC.md` is the dashboard.
- no parallel workers. One thread, one spec, one diff.
- no JSON / YAML spec bodies. Markdown + pipe tables.
- no hooks, no orchestration binaries, no TypeScript helpers.

---

## cavekit v1

v1 is **not deprecated** — it is frozen. It remains a fully working plugin
for users who prefer its scope.

**What v1 is**:

> Spec-driven AI development with an autonomous execution loop. Four-command
> Hunt lifecycle (`/ck:sketch` → `/ck:map` → `/ck:make` → `/ck:check`), plus
> `/ck:ship`, `/ck:review`, `/ck:revise`, `/ck:status`, `/ck:design`,
> `/ck:research`, `/ck:init`, `/ck:config`, `/ck:resume`, `/ck:help` — 16
> slash commands total. 12 named sub-agents. Per-task token budgets,
> stop-hook state machine, model-tier routing, auto-backpropagation from
> test failures, tool-result caching, Codex peer review, Karpathy
> behavioral guardrails, caveman token compression, knowledge-graph
> integration, and design-system enforcement. Parallel wave execution
> and team mode.

**Pick v1 if** you want the full autonomous loop, parallel agents, peer
review, or design-system workflow. **Pick v2 if** you want the distilled
core — one spec, three commands, no orchestration.

### install v1

Marketplace:

```bash
/plugin marketplace add juliusbrussee/cavekit@v1.3.1-final
/plugin install ck@cavekit
```

Git:

```bash
git clone -b v1.3.1-final https://github.com/juliusbrussee/cavekit.git
```

v1 docs live at that tag — `git checkout v1.3.1-final` and read the README
there for full command reference, skill catalog, and the Hunt lifecycle
guide.

### choosing, or moving

See [`UPGRADE.md`](./UPGRADE.md). Honest framing:
- Stay on v1 if your project has active `context/kits/` investment.
- Move to v2 if you want fewer moving parts and smaller token bills.
- It is a **two-way door** — `SPEC.md` is plain markdown; nothing traps
  you in either direction.

## philosophy

> The spec is the only artifact that earns its tokens. Everything else that
> costs tokens must either save more tokens later or the user's attention,
> or it gets cut.

See [`CHANGELOG.md`](./CHANGELOG.md) for the v1 → v2 break.

## license

MIT.
