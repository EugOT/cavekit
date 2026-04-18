---
name: peer-review
description: >
  Patterns for using a second AI agent or model to challenge the primary builder agent's work.
  Covers six review modes (Diff Critique, Design Challenge, Threaded Debate, Delegated Scrutiny,
  Deciding Vote, Coverage Audit), how to set up peer review with any model via MCP server,
  peer review iteration loops that alternate builder and reviewer prompts, the Codex Loop Mode
  (Cavekit + Ralph Loop + Codex as reviewer via CLI or MCP fallback), and prompt templates for
  each strategy. The peer reviewer's job is to find what the builder missed, not to agree.
  Triggers: "peer review", "peer review agent", "use another model to review",
  "second opinion on code", "cross-model review", "peer review loop", "ralph loop with codex",
  "cavekit ralph", "cross-model loop", "codex peer reviewer".
---

# Peer Review

Use a second AI agent to challenge the first agent's work. The most effective quality gate beyond automated tests.

## Core Principle

> **The peer reviewer's job is to find what the builder missed, not to agree.**

A review that says "looks good" is wasted. Instruct the peer reviewer to be critical, challenge assumptions, and look for what is *not* there.

## Why Peer Review Works

Every LLM has blind spots — patterns it over-relies on, edge cases it misses, implicit architectural assumptions. A second model (or same model with a different role prompt) catches a different set of issues.

**What peer review catches that tests miss:**
- Architectural over/under-engineering
- Missing error handling
- Security vulnerabilities the builder didn't consider
- Cavekit requirements technically met but poorly implemented
- Dead code, unused imports, unnecessary complexity
- Performance pitfalls at scale
- Edge cases not covered by the cavekit

---

## Review Modes

| Mode | Timing | Mechanism |
|------|--------|-----------|
| **Diff Critique** | After implementation | Second model inspects changeset with fault-finding prompt; builder incorporates valid fixes |
| **Design Challenge** | During planning | Second model proposes alternatives; builder evaluates both against spec |
| **Threaded Debate** | Complex trade-offs | Multi-turn conversation on a persistent thread |
| **Delegated Scrutiny** | Substantial reviews | Dedicated teammate owns the peer review interaction, reports to lead |
| **Deciding Vote** | Two approaches conflict | Lead presents both to peer model; it analyzes trade-offs and recommends |
| **Coverage Audit** | Validation phase | Coverage data + gap analysis fed to peer model for independent assessment |

### Choosing the Right Mode

```
Need peer review
├─ Reviewing completed code?
│   ├─ < 500 lines → Diff Critique
│   └─ Large/full feature → Delegated Scrutiny
├─ Designing architecture?
│   ├─ Single decision → Deciding Vote
│   └─ Full system → Design Challenge
├─ Debating trade-offs?
│   ├─ Extended back-and-forth → Threaded Debate
│   └─ Decisive answer → Deciding Vote
└─ Validating tests? → Coverage Audit
```

---

## Setting Up Peer Review via MCP Server

Any model exposing an MCP server can be a peer reviewer. Pattern is model-agnostic.

### Generic MCP Configuration

Add the peer review model to `.mcp.json`:

```json
{
  "mcpServers": {
    "peer reviewer": {
      "command": "{ADVERSARY_CLI}",
      "args": ["mcp-server"],
      "env": { "API_KEY": "{ADVERSARY_API_KEY}" }
    }
  }
}
```

### Two Core MCP Tools

1. **Start session** — new conversation. Params: prompt, approval policy, sandbox, model. Returns thread/session ID.
2. **Reply to session** — continue. Params: thread ID, follow-up message. Returns response.

The thread ID is critical — it enables multi-turn context.

### Example: Start

```
Tool: peer reviewer.start_session
Parameters:
  prompt: "Review these code changes for bugs, security issues,
           missing edge cases, and spec compliance. Be critical —
           your job is to find problems. Changes: {DIFF_CONTENT}"
  model: "{ADVERSARY_MODEL}"
```

### Example: Follow-Up

```
Tool: peer reviewer.reply_to_session
Parameters:
  thread_id: "{THREAD_ID}"
  message: "Good findings. Now focus on error handling paths.
            For each function that can fail, verify explicit
            error handling and correct error propagation."
```

---

## Strategy Details

### 1. Diff Critique

**When:** After builder completes a feature or fix.

**Process:**
1. Builder implements and commits
2. Generate diff: `git diff {BASE_BRANCH}...HEAD`
3. Send diff to peer reviewer with review prompt
4. Parse findings into actionable items
5. Builder applies fixes for valid findings
6. Optional: re-review after fixes

**Review Prompt Template:**
```markdown
You are a senior code reviewer. Review critically.

## What to look for:
- Bugs, logic errors, off-by-one
- Security vulnerabilities (injection, auth bypass, data exposure)
- Missing error handling and edge cases
- Performance issues (N+1, unnecessary allocations, blocking calls)
- Cavekit compliance: does this match the requirements?
- Code quality: naming, structure, unnecessary complexity

## What NOT to do:
- Do not say "looks good" unless you genuinely found zero issues
- Do not suggest stylistic changes unless they affect readability significantly
- Do not rewrite — describe the problem and where it is

## Cavekit requirements:
{CAVEKIT_REQUIREMENTS}

## Changes:
{DIFF_CONTENT}

## Output format:
For each finding:
- **Severity:** CRITICAL / HIGH / MEDIUM / LOW
- **File:** path and line range
- **Issue:** what is wrong
- **Why:** why this matters
- **Suggestion:** how to fix
```

### 2. Design Challenge

**When:** Planning phase, before implementation.

**Process:**
1. Builder drafts plan/architecture
2. Send plan + kits to peer reviewer
3. Peer reviewer proposes alternatives or critiques
4. Builder validates both against kits
5. Human decides on genuine trade-offs

**Architecture Review Prompt Template:**
```markdown
You are a systems architect reviewing a proposed design. Find
weaknesses, over-engineering, missing considerations, and better
alternatives.

## Kits (what must be built):
{CAVEKIT_CONTENT}

## Proposed architecture:
{PLAN_CONTENT}

## Evaluate:
1. Does it satisfy all cavekit requirements?
2. Is it over-engineered for the scope?
3. Are there simpler alternatives meeting the same requirements?
4. What failure modes exist? How does it recover?
5. What are the scaling bottlenecks?
6. What dependencies introduce risk?
```

### 3. Threaded Debate

**When:** Complex design discussions needing extended back-and-forth.

**Process:**
1. Start session presenting the problem
2. Use reply-to-session across multiple turns
3. Maintain thread ID throughout
4. Summarize when discussion converges

**Key:** Threads accumulate context. Keep on a single topic to avoid dilution.

### 4. Delegated Scrutiny

**When:** Large tasks where review itself is substantial.

**Process:**
1. Team lead spawns teammate for peer review coordination
2. Teammate owns peer reviewer MCP interaction
3. Manages multi-turn review sessions
4. Summarizes findings for lead
5. Lead assigns fixes to appropriate builders

**Why delegate:** Peer review consumes significant context. A dedicated teammate preserves the lead's context for coordination.

### 5. Deciding Vote

**When:** Builder and human (or two agents) disagree.

**Process:**
1. Present both perspectives to peer reviewer
2. Ask for trade-off evaluation
3. Ask for a recommendation with reasoning
4. Human has final say

**Tie-Breaking Prompt Template:**
```markdown
Two approaches proposed. Evaluate both critically and recommend one.

## Context:
{PROBLEM_DESCRIPTION}

## Approach A:
{APPROACH_A}

## Approach B:
{APPROACH_B}

## Evaluation criteria:
- Correctness: which is more likely correct?
- Simplicity: which is easier to understand/maintain?
- Performance: which performs better for expected use?
- Risk: which has fewer failure modes?

## Recommendation:
Pick one and explain why. If neither is clearly better, say so
and explain what info would break the tie.
```

### 6. Coverage Audit

**When:** During validation, after tests generated and run.

**Process:**
1. Run coverage analysis
2. Generate coverage report
3. Send report + kits to peer reviewer
4. Peer reviewer identifies: untested edge cases, missing integration tests, cavekit requirements without tests
5. Builder adds missing tests

---

## Peer Review Iteration (Convergence Loop with Review)

Alternating convergence loops — each iteration alternates between building and reviewing.

### Pattern

```
Iter 1: Builder runs against spec → code
Iter 2: Reviewer runs against code + spec → findings
Iter 3: Builder runs against spec + findings → fixed code
Iter 4: Reviewer runs against updated code + spec → new findings
...repeat until findings converge to zero (or trivial)
```

### Implementation with Separate Prompts

**`prompts/build.md`:**
```markdown
Implement requirements in the cavekit. Read impl tracking for
context. Read review findings and address them.

Input: kits/, plans/, impl/, review-findings.md (if exists)
Output: source code, updated impl tracking
Exit: all cavekit requirements implemented, all findings addressed
```

**`prompts/review.md`:**
```markdown
Review current implementation against cavekit. Be critical. Find
bugs, missing requirements, security issues, quality problems.

Input: kits/, plans/, source code, impl/
Output: review-findings.md
Exit: all source files reviewed against all cavekit requirements
```

### Running

```bash
# Terminal 1: Builder loop
{LOOP_TOOL} prompts/build.md -n 5 -t 2h

# Terminal 2: Reviewer loop (staggered 30 min)
{LOOP_TOOL} prompts/review.md -n 5 -t 2h -d 30m
```

Both share the same git repo. Reviewer reads latest committed code; builder reads latest `review-findings.md`. They converge through git.

### Convergence Signal

- Reviewer findings drop to zero or only LOW severity
- Builder diffs between iterations are minimal
- All cavekit requirements reviewed and confirmed met

---

## Anti-Patterns

### 1. Peer reviewer as Yes-Man
**Problem:** Says "looks good" without real issues.
**Fix:** Instruct: "If you find zero issues, explain what areas you checked and why they're correct. An empty review is suspicious."

### 2. Peer reviewer Rewrites Everything
**Problem:** Full rewrites instead of issue identification.
**Fix:** "Your output is a list of findings, not a pull request."

### 3. Builder Ignores Findings
**Problem:** Dismisses findings without addressing.
**Fix:** "For each finding, either fix it and explain the fix, or explain why it's invalid. You may not skip any finding."

### 4. Infinite Disagreement Loop
**Problem:** Builder/reviewer keep going without converging.
**Fix:** Set max iteration count. After N, escalate to human. Persistent disagreement usually signals an ambiguous spec.

### 5. Same Model Reviewing Itself
**Problem:** Same model + same prompt for both roles.
**Fix:** At minimum, different prompts with different roles. Ideally, different model or version.

---

## Prompt Templates Quick Reference

| Mode | Key Prompt Instruction |
|------|----------------------|
| Diff Critique | "Find bugs, security issues, missing edge cases. Do not say 'looks good'." |
| Design Challenge | "Find weaknesses and simpler alternatives. Evaluate failure modes." |
| Threaded Debate | "Continue the discussion. Build on previous context." |
| Delegated Scrutiny | "Own the peer reviewer interaction. Summarize findings for the lead." |
| Deciding Vote | "Evaluate both approaches. Recommend one with explicit reasoning." |
| Coverage Audit | "Identify untested edge cases and spec requirements without tests." |

---

## Integration with Cavekit Lifecycle

| Hunt Phase | Peer Review Role |
|-------------|-----------------|
| **Draft** | Review kits for completeness, ambiguity, missing edge cases |
| **Architect** | Challenge the plan before implementation begins |
| **Build** | Review implementation against kits after each feature |
| **Inspect** | Alternate build/review convergence loop |
| **Monitor** | Validate monitoring covers all failure modes |

Most impactful point: **Inspect** — catches issues that neither tests nor single-agent loops find.

---

## Cross-References

- **convergence-monitoring** — detecting when iterations converged
- **validation-first** — peer review is Gate 6 (human/agent review)
- **prompt-pipeline** — structuring builder/reviewer prompts
- **revision** — when peer reviewer finds a cavekit gap, revise into kits
- **impl-tracking** — record findings in impl tracking

---

## Codex Loop Mode — Cavekit + Ralph Loop + Codex Peer Reviewer

Most rigorous automated QA available: Cavekit cavekit through a Ralph Loop where Claude builds and Codex adversarially reviews every few iterations. Different model = different training data, biases, and blind spots.

### Why This Works

| Factor | Single-Model Loop | Codex Loop Mode |
|--------|-------------------|-----------------|
| Blind spots | Same model every iteration | Two models catch different issue classes |
| Cavekit drift | Builder may silently deviate | Peer reviewer checks compliance explicitly |
| Quality floor | Converges to "good enough for one model" | Converges to "survives cross-examination" |
| Dead ends | May retry failed approaches | Peer flags repeated patterns |

### Architecture

```
┌─────────────────────────────────────────────────────┐
│                   Ralph Loop                         │
│  (Stop hook feeds same prompt each iteration)        │
│                                                      │
│  ┌──────────┐    ┌──────────────┐    ┌────────────┐ │
│  │  Claude   │───▶│ Build from   │───▶│  Commit    │ │
│  │  (Build)  │    │ cavekit      │    │  changes   │ │
│  └──────────┘    └──────────────┘    └──────┬─────┘ │
│       ▲                                      │       │
│       │                                      ▼       │
│  ┌──────────┐    ┌──────────────┐    ┌────────────┐ │
│  │  Fix      │◀──│ Parse        │◀──│  Codex CLI │ │
│  │  findings │    │ findings     │    │  (Review)  │ │
│  └──────────┘    └──────────────┘    └────────────┘ │
│                                                      │
│  Completion: all cavekit requirements met +         │
│              no CRITICAL/HIGH findings               │
└─────────────────────────────────────────────────────┘
```

### Review Invocation: Codex CLI (primary) vs MCP (legacy)

1. **Codex CLI delegation (primary)** — `scripts/codex-review.sh` calls `codex` in `--approval-mode full-auto`. Faster, no MCP overhead. Findings go to `context/impl/impl-review-findings.md`.
2. **MCP server (legacy fallback)** — Codex as MCP server in `.mcp.json`. Used only when CLI unavailable.

`setup-build.sh` auto-detects: if `codex-review.sh` + `codex` CLI present → CLI delegation. Else → MCP fallback.

### Activation

```bash
/ck:make --peer-review                       # Codex Loop (every 2nd iteration default)
/ck:make --peer-review --review-interval 1   # review every iteration
/ck:make --peer-review --codex-model gpt-5.4-mini   # cheaper reviewer
```

### What `--peer-review` Does

1. **Validates** Codex CLI installed (or MCP fallback configured).
2. **Configures** Codex as MCP server in `.mcp.json` if CLI unavailable.
3. **Builds** Ralph Loop prompt embedding:
   - Cavekit path + plan/impl files
   - Build/review iteration alternation
   - Peer review prompt template for Codex
   - Completion criteria tied to acceptance criteria
4. **Starts** Ralph Loop via stop hook.

### Codex CLI Invocation

```bash
source scripts/codex-review.sh
bp_codex_review --base main
```

Produces structured findings with severity (P0–P3). Graceful fallback if Codex unavailable.

### MCP Fallback Config

```json
{
  "mcpServers": {
    "codex-reviewer": {
      "command": "codex",
      "args": ["mcp-server", "-c", "model=\"gpt-5.4\""]
    }
  }
}
```

### Iteration Pattern

```
Iter 1: BUILD  — Read cavekit, implement first requirement
Iter 2: REVIEW — Call Codex CLI, get findings, fix CRITICAL/HIGH
Iter 3: BUILD  — Continue, address remaining findings
Iter 4: REVIEW — Codex again, new findings on new code
...
Iter N: BUILD  — All requirements met, all findings fixed
             → outputs <promise>CAVEKIT COMPLETE</promise>
```

Default: review every 2nd iteration. `--review-interval 1` = every iteration.

### Peer Review Findings File

Findings tracked in `context/impl/impl-review-findings.md`:

```markdown
# Peer Review Findings

## Latest Review: Iteration 4 — 2026-03-14T10:30:00Z
### Reviewer: Codex (gpt-5.4)

| # | Severity | File | Issue | Status |
|---|----------|------|-------|--------|
| 1 | CRITICAL | src/auth.ts:L42 | Missing input validation on token | FIXED |
| 2 | HIGH | src/auth.ts:L67 | Race condition in session refresh | FIXED |
| 3 | MEDIUM | src/auth.ts:L15 | Unused import | NEW |
| 4 | LOW | src/auth.ts:L3 | Comment typo | WONTFIX |

## History
### Iteration 2
| # | Severity | File | Issue | Status |
|---|----------|------|-------|--------|
| 1 | CRITICAL | src/auth.ts:L20 | SQL injection in login query | FIXED |
```

### Completion Criteria (Codex Loop Mode)

Loop exits on completion promise. Claude must only output it when ALL are true:

- All cavekit requirements (R-numbers) implemented
- All acceptance criteria pass
- No CRITICAL/HIGH findings unfixed
- Build passes
- Tests pass
- At least one review iteration completed with no new CRITICAL/HIGH findings

### Review-Only Mode

For reviewing existing code without building:

```bash
/ck:review --codex     # single Codex-only review
```

Each iteration calls Codex to review code against cavekit, then fixes issues.

### Prerequisites (Codex Loop Mode)

1. Codex CLI: `npm install -g @openai/codex`
2. OpenAI API key (via `codex login` or env var)
3. Cavekit file exists at given path

### Convergence Signals (Codex Loop Mode)

- Codex findings drop to zero or LOW/MEDIUM only
- Iteration diffs minimal
- All requirements confirmed met by both Claude and Codex

If max iterations without converging:
- Check `context/impl/impl-review-findings.md` for persistent issues
- Consider whether cavekit needs clarification
- Run `/ck:revise --trace` to trace issues back to kits
