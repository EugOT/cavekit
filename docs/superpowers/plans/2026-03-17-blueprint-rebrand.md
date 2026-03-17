# Blueprint Rebrand Implementation Plan

> **For agentic workers:** REQUIRED: Use superpowers:subagent-driven-development (if subagents available) or superpowers:executing-plans to implement this plan. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Rebrand "SDD — Spec-Driven Development" to "Blueprint" across the entire project — file names, content, commands, agents, visual theme.

**Architecture:** Pure rename and theming — no structural changes. File renames via `git mv`, content updates via find-and-replace with manual review for context-sensitive terms (e.g., "spec" that refers to external specs should stay). Visual theme applies blueprint aesthetic (blue tones, box-drawing headers, construction symbols).

**Tech Stack:** Bash scripts, TypeScript (picker), Claude Code plugin manifest (JSON), Markdown docs

---

## Chunk 1: Core Infrastructure (plugin.json, install.sh, CLI, package.json)

### Task 1: Rename and update plugin.json

**Files:**
- Modify: `plugin.json`

- [ ] **Step 1: Update plugin.json**

```json
{
  "name": "blueprint",
  "description": "Blueprint — a methodology for building software with AI coding agents where blueprints drive the development process. Includes DABI lifecycle, prompt pipelines, validation gates, revision, convergence monitoring, and agent team coordination patterns.",
  "version": "2.0.0"
}
```

- [ ] **Step 2: Commit**

```bash
git add plugin.json
git commit -m "rebrand: rename plugin from sdd to blueprint"
```

### Task 2: Rename scripts/sdd to scripts/blueprint and update content

**Files:**
- Rename: `scripts/sdd` → `scripts/blueprint`
- Modify: `scripts/blueprint` (content updates)

- [ ] **Step 1: Rename the file**

```bash
git mv scripts/sdd scripts/blueprint
```

- [ ] **Step 2: Update all content in scripts/blueprint**

Replace all branding and references:
- Header comment: `sdd — Global CLI for Spec-Driven Development` → `blueprint — Global CLI for Blueprint`
- Usage block: all `sdd` → `blueprint`
- `usage()` function: all `sdd` references → `blueprint`
- Worktree patterns: `${PROJECT_NAME}-sdd-` → `${PROJECT_NAME}-blueprint-`
- `sed "s/^${PROJECT_NAME}-sdd-//"` → `sed "s/^${PROJECT_NAME}-blueprint-//"`
- Session name references: `"sdd"` → `"blueprint"`
- Branch patterns: `'sdd/*'` → `'blueprint/*'`
- Picker temp file: `/tmp/sdd-picker-` → `/tmp/blueprint-picker-`
- Env var: `SDD_PICKER_OUTFILE` → `BLUEPRINT_PICKER_OUTFILE`
- Status output: `"SDD Agent Status:"` → `"Blueprint Agent Status:"`
- `"No active SDD worktrees found."` → `"No active Blueprint worktrees found."`
- Merge message: `"Run /sdd:merge"` → `"Run /bp:merge"`
- Picker script reference: `sdd-picker.ts` → `blueprint-picker.ts`
- Launch script reference: `sdd-launch-session.sh` → `blueprint-launch-session.sh`
- Analytics script reference: `sdd-analytics.sh` → `blueprint-analytics.sh`
- Color: `CY=$'\033[36m'` stays in this file (no color vars here — uses none)

- [ ] **Step 3: Commit**

```bash
git add scripts/blueprint
git commit -m "rebrand: rename sdd CLI to blueprint with updated references"
```

### Task 3: Rename and update scripts/sdd-launch-session.sh

**Files:**
- Rename: `scripts/sdd-launch-session.sh` → `scripts/blueprint-launch-session.sh`
- Modify: content updates

- [ ] **Step 1: Rename**

```bash
git mv scripts/sdd-launch-session.sh scripts/blueprint-launch-session.sh
```

- [ ] **Step 2: Update content**

- Header comment: `sdd-launch-session` → `blueprint-launch-session`
- `SESSION_NAME="sdd"` → `SESSION_NAME="blueprint"`
- Worktree path: `${PROJECT_NAME}-sdd-${name}` → `${PROJECT_NAME}-blueprint-${name}`
- Branch name: `sdd/${name}` → `blueprint/${name}`
- Derive name sed: remove `feature-frontier-` → remove `build-site-` (update regex)
- Launcher echo: `"SDD Agent: $name"` → `"Blueprint Agent: $name"`
- Temp files: `/tmp/sdd-launch-` → `/tmp/blueprint-launch-`
- Window name: `"sdd-agents"` → `"blueprint-agents"`
- Send keys: `/sdd:execute` → `/bp:build`
- Status poller reference: `sdd-status-poller.sh` → `blueprint-status-poller.sh`
- Dashboard pane titles: `"sdd-progress"` → `"blueprint-progress"`, `"sdd-activity"` → `"blueprint-activity"`
- Output: `"Launched ${#FRONTIERS[@]} SDD agents:"` → `"Launched ${#FRONTIERS[@]} Blueprint agents:"`
- Kill instruction: `"Kill all: sdd --kill"` → `"Kill all: blueprint --kill"`

- [ ] **Step 3: Commit**

```bash
git add scripts/blueprint-launch-session.sh
git commit -m "rebrand: rename launch session script to blueprint"
```

### Task 4: Rename and update scripts/sdd-analytics.sh

**Files:**
- Rename: `scripts/sdd-analytics.sh` → `scripts/blueprint-analytics.sh`
- Modify: content + visual theme

- [ ] **Step 1: Rename**

```bash
git mv scripts/sdd-analytics.sh scripts/blueprint-analytics.sh
```

- [ ] **Step 2: Update content and apply blueprint theme**

- Header: `sdd-analytics` → `blueprint-analytics`
- Color: `CY=$'\033[36m'` → `BL=$'\033[34m'` and use `BL` where `CY` was used for headers
- Title: `"SDD Analytics"` → Blueprint ASCII header:
  ```
  ┌──────────────────────────┐
  │  B L U E P R I N T       │
  └──────────────────────────┘
  ```
  Then `"Analytics"` subtitle below
- Divider: use `BL` (blue) instead of `D` (dim) for dividers
- Worktree patterns: `${PROJECT_NAME}-sdd-` → `${PROJECT_NAME}-blueprint-`
- Temp file: `/tmp/sdd-tiers-` → `/tmp/blueprint-tiers-`
- Error message: `"No loop logs found. Run /sdd:execute first."` → `"No loop logs found. Run /bp:build first."`
- Symbol updates: `✕` stays (already used for dead ends)

- [ ] **Step 3: Commit**

```bash
git add scripts/blueprint-analytics.sh
git commit -m "rebrand: rename analytics script with blueprint theme"
```

### Task 5: Rename and update scripts/sdd-status-poller.sh

**Files:**
- Rename: `scripts/sdd-status-poller.sh` → `scripts/blueprint-status-poller.sh`
- Modify: content

- [ ] **Step 1: Rename**

```bash
git mv scripts/sdd-status-poller.sh scripts/blueprint-status-poller.sh
```

- [ ] **Step 2: Update content**

- Header: `sdd-status-poller` → `blueprint-status-poller`
- `SESSION_NAME="sdd"` → `SESSION_NAME="blueprint"`
- Worktree pattern: `${PROJECT_NAME}-sdd-` → `${PROJECT_NAME}-blueprint-`
- Sed pattern: `s/^${PROJECT_NAME}-sdd-//` → `s/^${PROJECT_NAME}-blueprint-//`
- Symbol updates: `✓` → `■`, `·` → `○`

- [ ] **Step 3: Commit**

```bash
git add scripts/blueprint-status-poller.sh
git commit -m "rebrand: rename status poller to blueprint"
```

### Task 6: Rename and update scripts/sdd-picker.ts

**Files:**
- Rename: `scripts/sdd-picker.ts` → `scripts/blueprint-picker.ts`
- Modify: content

- [ ] **Step 1: Rename**

```bash
git mv scripts/sdd-picker.ts scripts/blueprint-picker.ts
```

- [ ] **Step 2: Update content**

- `deriveName` regex: `feature-frontier-` → `build-site-` in the replace pattern
- Also handle legacy `feature-frontier-` pattern for backwards compat during transition
- Worktree path: `${projectName}-sdd-` → `${projectName}-blueprint-`
- Env var: `SDD_PICKER_OUTFILE` → `BLUEPRINT_PICKER_OUTFILE`
- Error messages: `"Run /sdd:plan"` → `"Run /bp:architect"`
- Picker prompt: `"Select frontiers to launch"` → `"Select build sites to launch"`

- [ ] **Step 3: Commit**

```bash
git add scripts/blueprint-picker.ts
git commit -m "rebrand: rename picker to blueprint with updated terminology"
```

### Task 7: Rename and update scripts/setup-execute.sh

**Files:**
- Rename: `scripts/setup-execute.sh` → `scripts/setup-build.sh`
- Modify: content

- [ ] **Step 1: Rename**

```bash
git mv scripts/setup-execute.sh scripts/setup-build.sh
```

- [ ] **Step 2: Update content**

- Header: `SDD Execute Setup Script` → `Blueprint Build Setup Script`
- Help text: all `sdd` → `blueprint`, `execute` → `build`
- `COMPLETION_PROMISE="SPEC COMPLETE"` → `COMPLETION_PROMISE="BLUEPRINT COMPLETE"`
- Worktree: `${PROJECT_NAME}-sdd-` → `${PROJECT_NAME}-blueprint-`
- Branch: `sdd/${WT_NAME}` → `blueprint/${WT_NAME}`
- Env var output: `SDD_WORKTREE_PATH=` → `BLUEPRINT_WORKTREE_PATH=`
- Error messages: `"Run /sdd plan"` → `"Run /bp:architect"`
- Ralph prompt title: `"# SDD Execute"` → `"# Blueprint Build"`
- All `spec` references (as document noun) → `blueprint`
- All `frontier` → `site`
- All `feature frontier` → `build site`
- Output header: `"SDD Execute — Loop activated!"` → `"Blueprint Build — Loop activated!"`
- Spec dir: `context/specs` → `context/blueprints`
- Frontier dir: stays `context/frontiers` → update to `context/sites` (or keep — decide based on scope)

Note: The `context/` directory structure convention (`context/specs/`, `context/frontiers/`) is user-facing. For this task, update the script to reference `context/blueprints/` and `context/sites/` as the new convention. However, also check for old paths as fallback for backwards compatibility during transition.

- [ ] **Step 3: Commit**

```bash
git add scripts/setup-build.sh
git commit -m "rebrand: rename setup-execute to setup-build with blueprint terminology"
```

### Task 8: Update scripts/dashboard-progress.sh with blueprint theme

**Files:**
- Modify: `scripts/dashboard-progress.sh`

- [ ] **Step 1: Apply blueprint visual theme**

- Color: `CY=$'\033[36m'` → `BL=$'\033[34m'`, use `BL` for headers
- Title: `"SDD PROGRESS"` → Blueprint ASCII header (compact version for sidebar):
  ```
  BLUEPRINT
  ```
  (bold blue, spaced lettering)
- `BGG=$'\033[42m'` → `BGB=$'\033[44m'` (blue background for progress bar)
- Worktree pattern: `${PROJECT_NAME}-sdd-` → `${PROJECT_NAME}-blueprint-`
- Frontier search: `*frontier*` → `*site*` (also keep `*frontier*` as fallback)
- Symbol: `✓` keeps (used contextually, fine to keep for tiers)
- Dividers: apply blue dim styling

- [ ] **Step 2: Commit**

```bash
git add scripts/dashboard-progress.sh
git commit -m "rebrand: apply blueprint theme to progress dashboard"
```

### Task 9: Update scripts/dashboard-activity.sh with blueprint theme

**Files:**
- Modify: `scripts/dashboard-activity.sh`

- [ ] **Step 1: Apply blueprint visual theme**

- Color: `CY=$'\033[36m'` → `BL=$'\033[34m'`, use `BL` for headers
- Title stays `"ACTIVITY"` (already generic, just change color to blue)
- Dividers: blue dim

- [ ] **Step 2: Commit**

```bash
git add scripts/dashboard-activity.sh
git commit -m "rebrand: apply blueprint theme to activity dashboard"
```

### Task 10: Update scripts/package.json

**Files:**
- Modify: `scripts/package.json`

- [ ] **Step 1: Update name**

```json
{
  "name": "blueprint-scripts",
  "private": true,
  "type": "module",
  "dependencies": {
    "@inquirer/prompts": "^7.0.0"
  },
  "devDependencies": {
    "tsx": "^4.0.0"
  }
}
```

- [ ] **Step 2: Commit**

```bash
git add scripts/package.json
git commit -m "rebrand: rename package to blueprint-scripts"
```

### Task 11: Update install.sh

**Files:**
- Modify: `install.sh`

- [ ] **Step 1: Update all branding and references**

- Clone URL comment: `~/.sdd` → `~/.blueprint`
- Color: `CY=$'\033[36m'` → `BL=$'\033[34m'`
- Header: `"SDD — Spec-Driven Development"` → Blueprint ASCII header
- `MARKETPLACE_NAME="sdd-os"` → `MARKETPLACE_NAME="blueprint"`
- Plugin key: `"sdd@${MARKETPLACE_NAME}"` → `"blueprint@${MARKETPLACE_NAME}"`
- Preflight warning: `"/sdd:..."` → `"/bp:..."`
- Script references: `scripts/sdd` → `scripts/blueprint`, etc. for all renamed scripts
- Symlink: `ln -sf ... "$BIN_DIR/sdd"` → `ln -sf ... "$BIN_DIR/blueprint"`
- Info messages: `"Installing sdd command..."` → `"Installing blueprint command..."`
- Usage output: all `sdd` → `blueprint`, all command names updated
- `/sdd:brainstorm` → `/bp:draft`
- `/sdd:plan` → `/bp:architect`
- `/sdd:execute` → `/bp:build`
- `/sdd:review` → `/bp:inspect`
- `/sdd:merge` → `/bp:merge`
- Use `BL` color for headers instead of `CY`

- [ ] **Step 2: Commit**

```bash
git add install.sh
git commit -m "rebrand: update installer for blueprint"
```

---

## Chunk 2: Commands (rename files + update frontmatter + content)

### Task 12: Rename and update command files

**Files:**
- Rename: `commands/brainstorm.md` → `commands/draft.md`
- Rename: `commands/plan.md` → `commands/architect.md`
- Rename: `commands/execute.md` → `commands/build.md`
- Rename: `commands/review.md` → `commands/inspect.md`
- Rename: `commands/back-propagate.md` → `commands/revise.md`
- Modify: `commands/gap-analysis.md` (frontmatter + content)
- Modify: `commands/help.md` (frontmatter + content)
- Modify: `commands/merge.md` (frontmatter + content)
- Modify: `commands/progress.md` (frontmatter + content)

- [ ] **Step 1: Rename command files**

```bash
git mv commands/brainstorm.md commands/draft.md
git mv commands/plan.md commands/architect.md
git mv commands/execute.md commands/build.md
git mv commands/review.md commands/inspect.md
git mv commands/back-propagate.md commands/revise.md
```

- [ ] **Step 2: Update frontmatter in all renamed commands**

For each renamed file, update the `name:` field:
- `commands/draft.md`: `name: sdd-brainstorm` → `name: blueprint-draft`
- `commands/architect.md`: `name: sdd-plan` → `name: blueprint-architect`
- `commands/build.md`: `name: sdd-execute` → `name: blueprint-build`
- `commands/inspect.md`: `name: sdd-review` → `name: blueprint-inspect`
- `commands/revise.md`: `name: sdd-back-propagate` → `name: blueprint-revise`

Also update description text — replace "spec" (document noun) → "blueprint", "SDD" → "Blueprint", "frontier" → "site", etc.

- [ ] **Step 3: Update frontmatter in non-renamed commands**

- `commands/gap-analysis.md`: `name: sdd-gap-analysis` → `name: blueprint-gap-analysis`
- `commands/help.md`: `name: sdd-help` → `name: blueprint-help`
- `commands/merge.md`: `name: sdd-merge` → `name: blueprint-merge`
- `commands/progress.md`: `name: sdd-progress` → `name: blueprint-progress`

- [ ] **Step 4: Update content in all command files**

Apply terminology mapping across all 9 command files:
- `SDD` → `Blueprint`
- `Spec-Driven Development` → `Blueprint`
- `spec` (document noun) → `blueprint`
- `specification` (document noun) → `blueprint`
- `BPER` → `DABI`
- `frontier` → `site`
- `feature frontier` → `build site`
- `spec-writer` → `drafter`
- `implementer` → `builder`
- `gap-analyzer` → `surveyor`
- `back-propagate` / `backpropagation` → `revise` / `revision`
- `/sdd:brainstorm` → `/bp:draft`
- `/sdd:plan` → `/bp:architect`
- `/sdd:execute` → `/bp:build`
- `/sdd:review` → `/bp:inspect`
- `/sdd:back-propagate` → `/bp:revise`
- `/sdd:merge` → `/bp:merge`
- `/sdd:progress` → `/bp:progress`
- `/sdd:gap-analysis` → `/bp:gap-analysis`
- `/sdd:help` → `/bp:help`
- `context/specs/` → `context/blueprints/`
- `spec-overview.md` → `blueprint-overview.md`
- `spec-{domain}.md` → `blueprint-{domain}.md`
- `feature-frontier-*.md` → `build-site-*.md`

Special attention for `commands/build.md` (formerly execute.md):
- `allowed-tools` references `setup-execute.sh` → `setup-build.sh`
- Env var: `SDD_WORKTREE_PATH` → `BLUEPRINT_WORKTREE_PATH`

Special attention for `commands/help.md`:
- This file has ~38 references to `/sdd:*` — update ALL of them
- Legacy commands section: update or remove old `/sdd init`, `/sdd spec-from-refs` etc.
- Skill references: `sdd:sdd-methodology` → `bp:methodology`, `sdd:spec-writing` → `bp:blueprint-writing`, etc.

- [ ] **Step 5: Commit**

```bash
git add commands/
git commit -m "rebrand: rename all commands from sdd to blueprint"
```

---

## Chunk 3: Agents (rename files + update content)

### Task 13: Rename and update agent files

**Files:**
- Rename: `agents/spec-writer.md` → `agents/drafter.md`
- Rename: `agents/plan-architect.md` → `agents/architect.md`
- Rename: `agents/implementer.md` → `agents/builder.md`
- Rename: `agents/gap-analyzer.md` → `agents/surveyor.md`
- Rename: `agents/adversarial-reviewer.md` → `agents/inspector.md`
- Modify: `agents/convergence-monitor.md` (content only)

- [ ] **Step 1: Rename agent files**

```bash
git mv agents/spec-writer.md agents/drafter.md
git mv agents/plan-architect.md agents/architect.md
git mv agents/implementer.md agents/builder.md
git mv agents/gap-analyzer.md agents/surveyor.md
git mv agents/adversarial-reviewer.md agents/inspector.md
```

- [ ] **Step 2: Update content in all agent files**

Apply terminology mapping in each file. Key replacements:
- Agent role descriptions: use new names (Drafter, Architect, Builder, Surveyor, Inspector)
- `spec` (document noun) → `blueprint`
- `SDD` → `Blueprint`
- `frontier` → `site`
- `BPER` → `DABI`
- Keep "spec" where it refers to external specs (OpenAPI, test specs)

- [ ] **Step 3: Commit**

```bash
git add agents/
git commit -m "rebrand: rename all agents from sdd to blueprint"
```

---

## Chunk 4: Skills (rename directories + update SKILL.md frontmatter + content)

### Task 14: Rename and update skill directories

**Files:**
- Rename dir: `skills/sdd-methodology/` → `skills/methodology/`
- Rename dir: `skills/spec-writing/` → `skills/blueprint-writing/`
- Rename dir: `skills/backpropagation/` → `skills/revision/`
- Modify: all `skills/*/SKILL.md` frontmatter and content

- [ ] **Step 1: Rename skill directories**

```bash
git mv skills/sdd-methodology skills/methodology
git mv skills/spec-writing skills/blueprint-writing
git mv skills/backpropagation skills/revision
```

- [ ] **Step 2: Update SKILL.md frontmatter in renamed dirs**

- `skills/methodology/SKILL.md`: `name: sdd-methodology` → `name: methodology`
- `skills/blueprint-writing/SKILL.md`: `name: spec-writing` → `name: blueprint-writing`
- `skills/revision/SKILL.md`: `name: backpropagation` → `name: revision`

Update descriptions: `SDD` → `Blueprint`, `spec` → `blueprint`, `BPER` → `DABI`, `backpropagation` → `revision`

- [ ] **Step 3: Update content in ALL skill SKILL.md files**

Apply terminology mapping across all 13 skill directories. The most heavily impacted:
- `skills/methodology/SKILL.md` — core methodology doc, heaviest SDD branding
- `skills/blueprint-writing/SKILL.md` — spec → blueprint throughout
- `skills/revision/SKILL.md` — backpropagation → revision
- `skills/prompt-pipeline/SKILL.md` — BPER → DABI
- `skills/validation-first/SKILL.md` — BPER → DABI
- `skills/brownfield-adoption/SKILL.md` — SDD → Blueprint, spec → blueprint
- `skills/adversarial-loop/SKILL.md` — SDD → Blueprint, spec → blueprint
- Others: lighter touch, mainly SDD → Blueprint

- [ ] **Step 4: Commit**

```bash
git add skills/
git commit -m "rebrand: rename and update all skills for blueprint"
```

---

## Chunk 5: References (rename + update content)

### Task 15: Rename and update reference docs

**Files:**
- Rename: `references/bper-phases.md` → `references/dabi-phases.md`
- Modify: all files in `references/`

- [ ] **Step 1: Rename**

```bash
git mv references/bper-phases.md references/dabi-phases.md
```

- [ ] **Step 2: Update content in all reference files**

Apply terminology mapping across all reference files:
- `references/dabi-phases.md` — heaviest: BPER → DABI, all phase names, spec → blueprint, frontier → site
- `references/validation-gates.md` — spec → blueprint (careful: "test framework" stays, "pytest" stays)
- `references/prompt-engineering.md` — `{FRAMEWORK}` variable stays (refers to external framework), SDD → Blueprint
- `references/convergence-patterns.md` — SDD → Blueprint
- `references/git-as-memory.md` — SDD → Blueprint, spec → blueprint
- `references/agent-team-patterns.md` — agent names updated, SDD → Blueprint
- `references/session-feedback-protocol.md` — SDD → Blueprint
- `references/multi-repo-strategy.md` — SDD → Blueprint, spec → blueprint

- [ ] **Step 3: Commit**

```bash
git add references/
git commit -m "rebrand: rename and update all references for blueprint"
```

---

## Chunk 6: README, example.md, and final cleanup

### Task 16: Rewrite README.md

**Files:**
- Modify: `README.md`

- [ ] **Step 1: Full rebrand of README**

Key changes:
- Title: `# SDD — Spec-Driven Development` → `# Blueprint`
- Description: update to reference Blueprint
- Install: `~/.sdd` → `~/.blueprint`, `sdd` → `blueprint`
- CLI section: all `sdd` → `blueprint`
- Commands section: all `/sdd:*` → `/bp:*` with new names
- Worktree: `sdd/<frontier-name>` → `blueprint/<site-name>`
- `/sdd:execute` → `/bp:build`
- Branch references: `sdd/*` → `blueprint/*`
- File structure: `context/specs/` → `context/blueprints/`, `spec-overview.md` → `blueprint-overview.md`, `spec-{domain}.md` → `blueprint-{domain}.md`, `feature-frontier-*.md` → `build-site-*.md`
- Command table: update all names
- CLI table: update all names
- Terminology: "spec" → "blueprint", "frontier" → "site", BPER descriptions → DABI

- [ ] **Step 2: Commit**

```bash
git add README.md
git commit -m "rebrand: rewrite README for blueprint"
```

### Task 17: Update example.md

**Files:**
- Modify: `example.md`

- [ ] **Step 1: Update example conversations**

- Fix existing inconsistency: `/sdd brainstorm` (space) → `/bp:draft` (colon)
- All command references updated
- All terminology updated (spec → blueprint, frontier → site, etc.)

- [ ] **Step 2: Commit**

```bash
git add example.md
git commit -m "rebrand: update example conversations for blueprint"
```

### Task 18: Final verification

- [ ] **Step 1: Search for remaining "sdd" references**

```bash
grep -ri "sdd" --include="*.md" --include="*.sh" --include="*.ts" --include="*.json" . | grep -v node_modules | grep -v ".git/" | grep -v "docs/superpowers/"
```

Any remaining hits should be reviewed — they're either:
- Legitimate (part of a URL, historical reference) → leave
- Missed during rename → fix

- [ ] **Step 2: Search for remaining "spec" references that should be "blueprint"**

```bash
grep -ri "\bspec\b" --include="*.md" . | grep -v node_modules | grep -v ".git/" | grep -v "docs/superpowers/" | grep -v "OpenAPI" | grep -v "test spec" | grep -v "npm"
```

Review each hit — external spec references stay, document noun references should be "blueprint".

- [ ] **Step 3: Verify all command frontmatter names are correct**

```bash
grep "^name:" commands/*.md
```

Expected output: all `blueprint-*` prefixed names.

- [ ] **Step 4: Verify all skill frontmatter names are correct**

```bash
grep "^name:" skills/*/SKILL.md
```

- [ ] **Step 5: Test install.sh runs without errors**

```bash
bash -n install.sh
```

- [ ] **Step 6: Test scripts/blueprint runs without errors**

```bash
bash -n scripts/blueprint
```

- [ ] **Step 7: Final commit if any fixes needed**

```bash
git add -A
git commit -m "rebrand: final cleanup and verification"
```
