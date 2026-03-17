# Blueprint Rebrand Design

Full rebrand of SDD-OS from "SDD — Spec-Driven Development" to **Blueprint**. Includes renaming all user-facing and internal references, remapping commands and phases to a construction metaphor, and applying a blueprint-inspired visual theme to all terminal UI.

## Brand Identity

- **Name:** Blueprint (no tagline)
- **Core metaphor:** Software as construction. You draft blueprints, an architect plans the build, builders execute, inspectors review.
- **Lifecycle:** DABI (Draft, Architect, Build, Inspect) — replaces BPER

## Command Renaming

| Current | Proposed |
|---|---|
| `/sdd:brainstorm` | `/bp:draft` |
| `/sdd:plan` | `/bp:architect` |
| `/sdd:execute` | `/bp:build` |
| `/sdd:review` | `/bp:inspect` |
| `/sdd:merge` | `/bp:merge` |
| `/sdd:progress` | `/bp:progress` |
| `/sdd:gap-analysis` | `/bp:gap-analysis` |
| `/sdd:back-propagate` | `/bp:revise` |
| `/sdd:help` | `/bp:help` |

## Agent Renaming

| Current File | New File | Agent Name |
|---|---|---|
| `agents/spec-writer.md` | `agents/drafter.md` | Drafter |
| `agents/plan-architect.md` | `agents/architect.md` | Architect |
| `agents/implementer.md` | `agents/builder.md` | Builder |
| `agents/gap-analyzer.md` | `agents/surveyor.md` | Surveyor |
| `agents/adversarial-reviewer.md` | `agents/inspector.md` | Inspector |
| `agents/convergence-monitor.md` | `agents/convergence-monitor.md` | (unchanged) |

## Script Renaming

| Current | Proposed |
|---|---|
| `scripts/sdd` | `scripts/blueprint` |
| `scripts/sdd-launch-session.sh` | `scripts/blueprint-launch-session.sh` |
| `scripts/sdd-analytics.sh` | `scripts/blueprint-analytics.sh` |
| `scripts/sdd-status-poller.sh` | `scripts/blueprint-status-poller.sh` |
| `scripts/sdd-picker.ts` | `scripts/blueprint-picker.ts` |
| `scripts/setup-execute.sh` | `scripts/setup-build.sh` |
| `scripts/dashboard-progress.sh` | `scripts/dashboard-progress.sh` (unchanged) |
| `scripts/dashboard-activity.sh` | `scripts/dashboard-activity.sh` (unchanged) |

## Branch & Worktree Naming

- Branch prefix: `sdd/feature-name` becomes `blueprint/feature-name`
- Worktree naming: `project-sdd-feature` becomes `project-blueprint-feature`

## Terminology Mapping

Global string replacements across all docs, skills, agents, commands, and references:

| Find | Replace |
|---|---|
| `SDD` (as brand) | `Blueprint` |
| `Spec-Driven Development` | `Blueprint` |
| `spec` (as document noun) | `blueprint` |
| `specification` (as document noun) | `blueprint` |
| `BPER` | `DABI` |
| `Brainstorm, Plan, Execute, Review` | `Draft, Architect, Build, Inspect` |
| `frontier` | `site` |
| `feature frontier` | `build site` |
| `spec-writer` | `drafter` |
| `implementer` | `builder` |
| `gap-analyzer` | `surveyor` |
| `back-propagate` / `backpropagation` | `revise` / `revision` |

### Exclusions (do NOT replace)

- "spec" referring to external concepts (OpenAPI spec, test spec, npm spec)
- "framework" referring to external frameworks (React, pytest, etc.)
- Convergence monitoring terminology (no natural construction equivalent)
- Generic terms: dead ends, iterations, cycles

## Visual Theme — Blueprint Aesthetic

### Color Palette

| Role | Current | New |
|---|---|---|
| Headers/brand | Cyan `\033[36m` | Blue `\033[34m` |
| Secondary/accents | Cyan | Bold blue `\033[1;34m` |
| Success/done | Green `\033[32m` | Green (keep) |
| Warning/WIP | Yellow `\033[33m` | Yellow (keep) |
| Error/blocked | Red `\033[31m` | Red (keep) |
| Dividers | Dim `─` | Blue dim `─` (grid lines) |
| Progress bar (done) | Green bg `\033[42m` | Blue bg `\033[44m` |

### ASCII Header

```
  ┌──────────────────────────┐
  │  B L U E P R I N T       │
  └──────────────────────────┘
```

Technical drawing style — thin box-drawing characters, spaced lettering.

### Symbol Updates

| Current | New | Meaning |
|---|---|---|
| `→` | `▸` | Directional indicator |
| `✓` | `■` | Completed (built) |
| `·` | `○` | Pending (not yet built) |
| `⟳` | `⟳` | Iteration (keep) |

### Dividers

Blue-tinted `─` lines evoking grid lines on blueprint paper.

## Scope — Files to Touch

### Must rename (files)
- `scripts/sdd` → `scripts/blueprint`
- `scripts/sdd-launch-session.sh` → `scripts/blueprint-launch-session.sh`
- `scripts/sdd-analytics.sh` → `scripts/blueprint-analytics.sh`
- `scripts/sdd-status-poller.sh` → `scripts/blueprint-status-poller.sh`
- `scripts/sdd-picker.ts` → `scripts/blueprint-picker.ts`
- `scripts/setup-execute.sh` → `scripts/setup-build.sh`
- `agents/spec-writer.md` → `agents/drafter.md`
- `agents/plan-architect.md` → `agents/architect.md`
- `agents/implementer.md` → `agents/builder.md`
- `agents/gap-analyzer.md` → `agents/surveyor.md`
- `agents/adversarial-reviewer.md` → `agents/inspector.md`

### Must rename (commands)
- `commands/brainstorm.md` → `commands/draft.md`
- `commands/plan.md` → `commands/architect.md`
- `commands/execute.md` → `commands/build.md`
- `commands/review.md` → `commands/inspect.md`
- `commands/back-propagate.md` → `commands/revise.md`
- `commands/gap-analysis.md` — keep filename, update frontmatter `name: sdd-gap-analysis` → `name: blueprint-gap-analysis`
- `commands/help.md` — keep filename, update frontmatter `name: sdd-help` → `name: blueprint-help`, update all `/sdd:*` references to `/bp:*`
- `commands/merge.md` — keep filename, update frontmatter `name: sdd-merge` → `name: blueprint-merge`, update `sdd/*` branch references
- `commands/progress.md` — keep filename, update frontmatter `name: sdd-progress` → `name: blueprint-progress`

### Must rename (skills directories)
- `skills/sdd-methodology/` → `skills/methodology/`
- `skills/spec-writing/` → `skills/blueprint-writing/`
- `skills/backpropagation/` → `skills/revision/`
- Update `name:` field in each skill's `SKILL.md` frontmatter accordingly

### Must rename (references)
- `references/bper-phases.md` → `references/dabi-phases.md`

### Must update — command frontmatter (critical)

Every command's YAML frontmatter `name:` field must change from `sdd-*` to `blueprint-*`. This is how Claude Code discovers slash commands — missing any will break command registration.

### Must update — environment variables (critical)

These env vars are cross-process contracts — both sides must be renamed atomically:
- `SDD_WORKTREE_PATH` → `BLUEPRINT_WORKTREE_PATH` (set in `setup-build.sh`, read in `commands/build.md`)
- `SDD_PICKER_OUTFILE` → `BLUEPRINT_PICKER_OUTFILE` (set in `scripts/blueprint`, read in `scripts/blueprint-picker.ts`)

### Must update — install paths and plugin registration

- Install directory: `~/.sdd` → `~/.blueprint`
- `MARKETPLACE_NAME="sdd-os"` → `MARKETPLACE_NAME="blueprint"` (or new repo name)
- Plugin settings key: `"sdd@${MARKETPLACE_NAME}"` → `"blueprint@${MARKETPLACE_NAME}"` in `install.sh`
- Note: existing installations will need the old key removed from `~/.claude/settings.json`

### Must update — misc files
- `scripts/package.json`: `"name": "sdd-scripts"` → `"name": "blueprint-scripts"`
- `COMPLETION_PROMISE="SPEC COMPLETE"` in `setup-build.sh` → `"BLUEPRINT COMPLETE"`
- Temp file patterns: `/tmp/sdd-picker-*` → `/tmp/blueprint-picker-*`, `/tmp/sdd-launch-*` → `/tmp/blueprint-launch-*`, `/tmp/sdd-tiers-*` → `/tmp/blueprint-tiers-*`

### Must update — user-facing file conventions
- `context/specs/` directory convention → `context/blueprints/`
- `spec-{domain}.md` filename pattern → `blueprint-{domain}.md`
- `spec-overview.md` → `blueprint-overview.md`
- `feature-frontier-*.md` filename pattern → `build-site-*.md`
- Regex patterns in `derive_name()` functions updated to match new prefix

### Must update content (all files)
- `plugin.json` — name, description, all command/agent references
- `install.sh` — branding, symlink names, install paths
- `README.md` — full rebrand
- `example.md` — update example conversations (fix `/sdd brainstorm` → `/bp:draft` etc.)
- All files in `references/` — terminology swaps
- All files in `skills/` — terminology swaps + SKILL.md frontmatter
- All files in `agents/` — terminology swaps within prompts
- All files in `commands/` — terminology swaps within prompts + YAML frontmatter
- All files in `scripts/` — color palette, symbols, headers, internal variable references

### Must update (internal strings in scripts)
- Session names: `sdd` → `blueprint`
- Worktree patterns: `*-sdd-*` → `*-blueprint-*`
- Branch prefixes: `sdd/` → `blueprint/`
- Dashboard headers: `SDD PROGRESS` → `BLUEPRINT PROGRESS`, etc.
- CLI help text and usage strings
- Color variable values (cyan → blue)
- Symbol characters (→ to ▸, ✓ to ■, · to ○)

### Exclusions — explicit decisions

- **Ralph Loop**: Keep name as-is. It's a named concept, not SDD branding. The `ralph-loop.local.md` state file stays unchanged.
- **Convergence monitoring**: Keep all terminology as-is.
- **Skill plugin prefix**: After rename, skills referenced as `bp:methodology` (not `blueprint:blueprint-methodology`) — avoid redundancy in skill names.
