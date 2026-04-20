# UPGRADE v1 → v2

Honest answer: v2 is not a minor version of v1. It is a different shape with
the same name. This doc helps you decide whether to move, and if so, how.

## SHOULD YOU UPGRADE?

**Stay on v1 if**:
- Your project has a large `context/kits/` investment you actively iterate on
- You rely on the autonomous loop, parallel wave execution, or peer review
- Your team has shared muscle memory on `/ck:sketch → /ck:map → /ck:make`
- Your hooks / scripts integrate with the v1 state machine

**Move to v2 if**:
- You want fewer moving parts
- You find yourself fighting the framework more than you use it
- Token cost of invoking v1 commands outweighs the value
- You start a fresh project and want the distilled version

Either is a valid answer. v1 is not abandoned — it is frozen. Frozen code
does not rot as fast as it looks like it does.

## WHAT CHANGED

| v1 | v2 |
|---|---|
| 16 slash commands | 3 (`/ck:spec`, `/ck:build`, `/ck:check`) |
| 12 named sub-agents | 0 — main Claude does the work |
| 21 skills | 2 (`caveman`, `backprop`) |
| `context/kits/` directory | single `SPEC.md` at repo root |
| Hunt lifecycle (sketch/map/make/check) | flat spec → build → check |
| Go binary, shell hooks, TS picker | none |
| Autonomous loop with stop-hook | native Claude Code plan-then-execute |
| Design system, knowledge graph, Codex review | cut |
| Parallel wave execution | single-thread |
| Caveman opt-in for internal chatter | caveman default for spec writes |

## MIGRATION PATH

There is no automated migrator. `kits/` structure does not map cleanly to
`SPEC.md` — the point of v2 is that caveman + pipe tables replace that
tree. A script would produce lossy nonsense.

**Recommended path** for an in-flight v1 project that has working code:

1. Check out a fresh branch off your current v1 branch: `git checkout -b v2-migration`.
2. Install cavekit v2 as a separate plugin or branch of the repo.
3. Run `/ck:spec from-code`. v2 will walk your built code and produce
   a `SPEC.md`. **The code is the source of truth**, not your old kits.
4. Review the generated spec. Amend with `/ck:spec amend §X`.
5. Your old `context/kits/` stays in git history. If you ever need the
   original reasoning, `git log -- context/kits/`.
6. Delete the old directory on the v2-migration branch once you trust the
   spec. Commit. Merge when ready.

**If you have not built anything yet** (still in sketch phase): you have the
easiest migration. Scrap the kit, start with `/ck:spec <your idea>`.

## WHAT YOU LOSE

- **Autonomous loop**: v2 has no stop-hook state machine. Each `/ck:build`
  invocation is one plan-then-execute. If you liked "leave it running for
  an hour," v2 does not do that. Use a shell loop or stay on v1.
- **Parallel execution**: v2 is deliberately single-thread. Big projects
  take linear wall-clock time. This was a considered trade.
- **Peer review via Codex**: cut. If you want a second model on a diff,
  run it manually or install a peer-review skill separately.
- **Design system, knowledge graph, team mode**: cut. Separate tools if
  you need them.
- **Dashboards**: cut. `cat SPEC.md | grep §T` replaces them.

## WHAT YOU GAIN

- Drastically smaller context footprint on every invocation
- A spec you can read in 30 seconds
- No more "which agent should I invoke?"
- No more orphaned state files
- Backprop as a reflex in every build, not an opt-in

## v1 REACHABILITY

v1 is frozen at tag `v1.3.1-final`. Always installable:

```bash
git clone -b v1.3.1-final https://github.com/juliusbrussee/cavekit.git
```

No v1 code in v2 history is destroyed — `git log main` on an older clone
still shows every v1 commit. v1 documentation stays at that tag.

## ONE-WAY DOOR?

No. You can switch back. `SPEC.md` is plain markdown — nothing stops you
from re-exporting it into `context/kits/*.md` if you decide v1 was right
for your project. The work is not trapped.

## QUESTIONS

Open an issue on GitHub. Label `v1` for v1 bugs (fixes only for critical
issues), `v2` for v2 questions.
