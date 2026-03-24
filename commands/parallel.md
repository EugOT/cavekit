---
name: bp-parallel
description: "Execute build site tasks in parallel — dispatches subagents per frontier wave"
argument-hint: "[--filter PATTERN] [--max-concurrent 3] [--max-waves 10]"
allowed-tools: ["Bash(${CLAUDE_PLUGIN_ROOT}/scripts/setup-build.sh:*)", "Bash(cd *)", "Bash(git *)"]
---

# Blueprint Parallel — Tier-Parallel Task Execution

This command executes build site tasks in parallel by dispatching independent subagents for each ready task in the current frontier.

## Step 1: Setup Worktree

Run the setup script to create an isolated worktree (same as `/bp:build`):

```!
"${CLAUDE_PLUGIN_ROOT}/scripts/setup-build.sh" $ARGUMENTS
```

Look for `BLUEPRINT_WORKTREE_PATH=` in the output. If present, `cd` into that path:

```
cd <BLUEPRINT_WORKTREE_PATH value>
```

## Step 2: Parse Arguments

Defaults:
- `--max-concurrent 3` — max subagents running at once
- `--max-waves 10` — max frontier waves before stopping
- `--filter` — passed through to setup-build.sh

## Step 3: Execution Loop

Repeat until all tasks are done or max-waves reached:

### 3a. Read State
1. Read the build site from `context/sites/` (the one selected by setup-build.sh)
2. Read ALL `context/impl/impl-*.md` files to build task status map
3. Read `context/impl/dead-ends.md` if it exists

### 3b. Compute Frontier
Identify **ready tasks**: tasks that are NOT done AND whose `blockedBy` dependencies are ALL done.

- Parse each task row in the build site
- Check its status in impl tracking
- Check if all its `blockedBy` task IDs have status DONE
- Collect all ready tasks — this is the **frontier**

### 3c. Report Wave
Print the frontier:
```
═══ Wave {N} ═══
{count} task(s) ready for parallel execution:
  {task_id}: {title} (tier {N}, deps: {deps})
  ...
```

### 3d. Dispatch

**If 0 ready tasks**: All tasks are either done or blocked. Check if everything is done — if so, report completion. If tasks remain but none are ready, report the blockage and stop.

**If 1 ready task**: Implement it directly yourself (no subagent overhead). Follow the builder workflow:
1. Read blueprint requirement + acceptance criteria
2. Implement the task
3. Run validation gates (build → tests)
4. Commit: `T-{ID}: {what was done}`
5. Update impl tracking

**If 2+ ready tasks**: Dispatch in parallel:

1. **You take the first task** and implement it directly (same as single-task flow above)
2. **Dispatch subagents for the remaining tasks** (up to max-concurrent - 1), each via the Agent tool:

```
Agent(
  subagent_type: "bp:builder",
  isolation: "worktree",
  prompt: "You are executing a single task for Blueprint parallel build.

TASK: {task_id} — {title}
SPEC: {spec_name}
REQUIREMENT: {requirement_id}

BUILD SITE: {path to build site}
BLUEPRINTS: {paths to relevant blueprint files}

ACCEPTANCE CRITERIA (from blueprint):
{paste the acceptance criteria for this task's requirement}

DEAD ENDS TO AVOID:
{paste relevant dead ends, or 'None'}

INSTRUCTIONS:
1. Read the build site entry for {task_id} to understand the full task
2. Read the blueprint requirement {requirement_id} for context
3. Implement the task following the plan's steps
4. Write tests as needed
5. Run validation: build must pass, tests must pass
6. CRITICAL: Commit with message '{task_id}: {description}' BEFORE reporting — your worktree is auto-deleted when you finish. Uncommitted work is lost.
7. Report your result as:
   TASK RESULT:
   - Task: {id} — {title}
   - Status: COMPLETE | PARTIAL | BLOCKED
   - Files created: {list}
   - Files modified: {list}
   - Issues: {any}
"
)
```

3. **Wait for all subagents to complete** (they run concurrently)
4. **Merge results**: Claude Code automatically cleans up subagent worktrees on completion, but the **branches and commits persist** in git. For each subagent that made changes:
   - The Agent tool result includes the branch name (e.g., `worktree/agent-abc123`)
   - Fetch the branch: `git fetch . <branch>` (may not be needed if local)
   - Merge each branch into your current branch: `git merge <branch> --no-edit`
   - If merge conflicts occur, resolve them using blueprint context (understand what each task intended)
   - After successful merge, clean up the branch: `git branch -D <branch>`
   - If the Agent result says "no changes were made", skip — the worktree was already cleaned up with nothing to merge

### 3e. Update Tracking
After all tasks in the wave complete (both your own and subagents'):

1. Update `context/impl/impl-*.md` with status for each completed task:
   ```
   | {task_id} | DONE | {brief description of what was implemented} |
   ```
2. If any task failed or was partial, record it appropriately
3. Record dead ends in `context/impl/dead-ends.md` if any were encountered

### 3f. Next Wave
Increment wave counter and go back to step 3a.

## Step 4: Completion

When all tasks in the build site are done:

```
═══ PARALLEL BUILD COMPLETE ═══
Waves executed: {N}
Tasks completed: {done}/{total}
```

## Circuit Breakers

- **3 consecutive test failures on the same task** → mark task as BLOCKED, document in dead-ends.md, skip to next task
- **Merge conflict unresolvable** → stop the wave, report which branches conflict, let the user decide
- **All remaining tasks blocked** → report the dependency chain and stop

## Critical Rules

- ONE task per subagent — never batch tasks
- Subagents MUST use `isolation: "worktree"` — never let parallel agents share a working directory
- Merge after EVERY wave — do not accumulate unmerged branches
- Update impl tracking after EVERY wave — the next wave reads it to compute the frontier
- If a subagent reports PARTIAL or BLOCKED, record it and continue — do not retry in the same wave
- Never push to remote unless explicitly asked
- Commit your own work before merging subagent branches (clean working tree required for merge)
