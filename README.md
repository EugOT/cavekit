# Cavekit v2

Compressed spec-driven development for Claude Code.

One file. Three commands. Zero sub-agents.

## WHY

Plan-then-execute forgets. SDD remembers — but most SDD frameworks bury that
value under agent swarms, dashboards, and ceremony that wastes tokens.

Cavekit v2 keeps only what earns its place:

- **Durable spec** — `SPEC.md` at repo root survives context resets.
- **Caveman encoding** — ~75% fewer tokens than prose. Symbols, fragments,
  pipe tables for repeating records.
- **Backprop reflex** — every test failure becomes a `§B` entry; classes of
  bug become `§V` invariants the spec will never forget.

That's the whole pitch.

## COMMANDS

| cmd | job |
|---|---|
| `/ck:spec` | create / amend / backprop `SPEC.md`. Sole mutator. |
| `/ck:build` | native plan → execute against spec. Auto-backprops on failure. |
| `/ck:check` | read-only drift report. Lists §V / §I / §T violations, suggests remedy. |

## FORMAT

See `FORMAT.md`. Fixed sections: §G goal, §C constraints, §I interfaces,
§V invariants, §T tasks (pipe table), §B bugs (pipe table). Caveman prose
plus two pipe tables. Nothing else.

## FILES

```
FORMAT.md          spec schema + caveman encoding rules
commands/          three markdown command prompts
skills/caveman     encoding skill
skills/backprop    bug → spec protocol (six steps)
```

## NON-GOALS

- No sub-agents. Main Claude does the work.
- No dashboards. `cat SPEC.md` is the dashboard.
- No parallel workers. One thread, one spec, one diff.
- No JSON / YAML spec bodies. Markdown + pipe tables.
- No hooks, no orchestration binaries, no TypeScript helpers.

## INSTALL

Add as a Claude Code plugin via marketplace, or clone into
`~/.claude/plugins/`. Commands become available as `/ck:spec`,
`/ck:build`, `/ck:check`.

## HISTORY

v1 (`main` branch) had 16 commands, 12 agents, 21 skills, a Go binary,
shell hooks, and a dashboard. v2 replaces all of it with this.

MIT.
