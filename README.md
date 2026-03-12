<<<<<<< HEAD
# Learn Claude Code - Bash is all you & agent need

<p align="center">
  <img src="./assets/cover.webp" alt="Learn Claude Code" width="800">
</p>

[![Python 3.10+](https://img.shields.io/badge/python-3.10+-blue.svg)](https://www.python.org/downloads/)
[![Tests](https://github.com/shareAI-lab/learn-claude-code/actions/workflows/test.yml/badge.svg)](https://github.com/shareAI-lab/learn-claude-code/actions)
[![License: MIT](https://img.shields.io/badge/License-MIT-green.svg)](./LICENSE)

> **Disclaimer**: This is an independent educational project by [shareAI Lab](https://github.com/shareAI-lab). It is not affiliated with, endorsed by, or sponsored by Anthropic. "Claude Code" is a trademark of Anthropic.

**Learn how modern AI agents work by building one from scratch.**

[Chinese / 中文](./README_zh.md) | [Japanese / 日本語](./README_ja.md)

---

## Why This Repository?

We created this repository out of admiration for Claude Code - **what we believe to be the most capable AI coding agent in the world**. Initially, we attempted to reverse-engineer its design through behavioral observation and speculation. The analysis we published was riddled with inaccuracies, unfounded guesses, and technical errors. We deeply apologize to the Claude Code team and anyone who was misled by that content.

Over the past six months, through building and iterating on real agent systems, our understanding of **"what makes a true AI agent"** has been fundamentally reshaped. We'd like to share these insights with you. All previous speculative content has been removed and replaced with original educational material.

---

> Works with **[Kode CLI](https://github.com/shareAI-lab/Kode)**, **Claude Code**, **Cursor**, and any agent supporting the [Agent Skills Spec](https://agentskills.io/specification).

<img height="400" alt="demo" src="https://github.com/user-attachments/assets/0e1e31f8-064f-4908-92ce-121e2eb8d453" />

## What You'll Learn

After completing this tutorial, you will understand:

- **The Agent Loop** - The surprisingly simple pattern behind all AI coding agents
- **Tool Design** - How to give AI models the ability to interact with the real world
- **Explicit Planning** - Using constraints to make AI behavior predictable
- **Context Management** - Keeping agent memory clean through subagent isolation
- **Knowledge Injection** - Loading domain expertise on-demand without retraining
- **Context Compression** - How agents work beyond their context window limits
- **Task Systems** - From personal notes to team project boards
- **Parallel Execution** - Background tasks and notification-driven workflows
- **Team Messaging** - Persistent teammates communicating through inboxes
- **Autonomous Teams** - Self-organizing agents that find and claim their own work
=======
[English](./README.md) | [中文](./README-zh.md) | [日本語](./README-ja.md)  
# Learn Claude Code -- A nano Claude Code-like agent, built from 0 to 1

```
                    THE AGENT PATTERN
                    =================

    User --> messages[] --> LLM --> response
                                      |
                            stop_reason == "tool_use"?
                           /                          \
                         yes                           no
                          |                             |
                    execute tools                    return text
                    append results
                    loop back -----------------> messages[]


    That's the minimal loop. Every AI coding agent needs this loop.
    Production agents add policy, permissions, and lifecycle layers.
```

**12 progressive sessions, from a simple loop to isolated autonomous execution.**
**Each session adds one mechanism. Each mechanism has one motto.**

> **s01** &nbsp; *"One loop & Bash is all you need"* &mdash; one tool + one loop = an agent
>
> **s02** &nbsp; *"Adding a tool means adding one handler"* &mdash; the loop stays the same; new tools register into the dispatch map
>
> **s03** &nbsp; *"An agent without a plan drifts"* &mdash; list the steps first, then execute; completion doubles
>
> **s04** &nbsp; *"Break big tasks down; each subtask gets a clean context"* &mdash; subagents use independent messages[], keeping the main conversation clean
>
> **s05** &nbsp; *"Load knowledge when you need it, not upfront"* &mdash; inject via tool_result, not the system prompt
>
> **s06** &nbsp; *"Context will fill up; you need a way to make room"* &mdash; three-layer compression strategy for infinite sessions
>
> **s07** &nbsp; *"Break big goals into small tasks, order them, persist to disk"* &mdash; a file-based task graph with dependencies, laying the foundation for multi-agent collaboration
>
> **s08** &nbsp; *"Run slow operations in the background; the agent keeps thinking"* &mdash; daemon threads run commands, inject notifications on completion
>
> **s09** &nbsp; *"When the task is too big for one, delegate to teammates"* &mdash; persistent teammates + async mailboxes
>
> **s10** &nbsp; *"Teammates need shared communication rules"* &mdash; one request-response pattern drives all negotiation
>
> **s11** &nbsp; *"Teammates scan the board and claim tasks themselves"* &mdash; no need for the lead to assign each one
>
> **s12** &nbsp; *"Each works in its own directory, no interference"* &mdash; tasks manage goals, worktrees manage directories, bound by ID

---

## The Core Pattern

```python
def agent_loop(messages):
    while True:
        response = client.messages.create(
            model=MODEL, system=SYSTEM,
            messages=messages, tools=TOOLS,
        )
        messages.append({"role": "assistant",
                         "content": response.content})

        if response.stop_reason != "tool_use":
            return

        results = []
        for block in response.content:
            if block.type == "tool_use":
                output = TOOL_HANDLERS[block.name](**block.input)
                results.append({
                    "type": "tool_result",
                    "tool_use_id": block.id,
                    "content": output,
                })
        messages.append({"role": "user", "content": results})
```

Every session layers one mechanism on top of this loop -- without changing the loop itself.

## Scope (Important)

This repository is a 0->1 learning project for building a nano Claude Code-like agent.
It intentionally simplifies or omits several production mechanisms:

- Full event/hook buses (for example PreToolUse, SessionStart/End, ConfigChange).  
  s12 includes only a minimal append-only lifecycle event stream for teaching.
- Rule-based permission governance and trust workflows
- Session lifecycle controls (resume/fork) and advanced worktree lifecycle controls
- Full MCP runtime details (transport/OAuth/resource subscribe/polling)

Treat the team JSONL mailbox protocol in this repo as a teaching implementation, not a claim about any specific production internals.

## Quick Start

```sh
git clone https://github.com/shareAI-lab/learn-claude-code
cd learn-claude-code
pip install -r requirements.txt
cp .env.example .env   # Edit .env with your ANTHROPIC_API_KEY

python agents/s01_agent_loop.py       # Start here
python agents/s12_worktree_task_isolation.py  # Full progression endpoint
python agents/s_full.py               # Capstone: all mechanisms combined
```

### Web Platform

Interactive visualizations, step-through diagrams, source viewer, and documentation.

```sh
cd web && npm install && npm run dev   # http://localhost:3000
```
>>>>>>> upstream/main

## Learning Path

```
<<<<<<< HEAD
Start Here
    |
    v
[v0: Bash Agent] ----------> "One tool is enough"
    |                         16-196 lines
    v
[v1: Basic Agent] ----------> "The complete agent pattern"
    |                          4 tools, ~417 lines
    v
[v2: Todo Agent] -----------> "Make plans explicit"
    |                          +TodoManager, ~531 lines
    v
[v3: Subagent] -------------> "Divide and conquer"
    |                          +Task tool, ~623 lines
    v
[v4: Skills Agent] ----------> "Domain expertise on-demand"
    |                           +Skill tool, ~783 lines
    v
[v5: Compression Agent] ----> "Never forget, work forever"
    |                          +ContextManager, ~896 lines
    v
[v6: Tasks Agent] ----------> "From sticky notes to kanban"
    |                          +TaskManager, ~1075 lines
    v
[v7: Background Agent] -----> "Don't wait, keep working"
    |                          +BackgroundManager, ~1142 lines
    v
[v8a: Team Foundation] --> "Creating and managing teammates"
    |                       +TeammateManager, ~1395 lines
    v
[v8b: Messaging] --------> "Teammates that communicate"
    |                       +SendMessage/Inbox, ~1557 lines
    v
[v8c: Coordination] -----> "Shared tasks and shutdown protocol"
    |                       +SharedBoard/Protocol, ~1612 lines
    v
[v9: Autonomous Agent] -----> "A self-organizing team"
                               +Idle cycle, ~1683 lines
```

**Recommended approach:**
1. Read and run v0 first - understand the core loop
2. Compare v0 and v1 - see how tools evolve
3. Study v2 for planning patterns
4. Explore v3 for complex task decomposition
5. Master v4 for building extensible agents
6. Study v5 for context management and compression
7. Explore v6 for persistent task tracking
8. Understand v7 for parallel background execution
9. Study v8a/v8b/v8c for team lifecycle and messaging
      a. v8a - TeammateManager (creation, deletion, config, tool scoping)
      b. v8b - Message protocol (5 types, JSONL inbox, inbox routing)
      c. v8c - Shared task board, shutdown protocol, plan approval
      d. Trace a full lifecycle across v8a -> v8b -> v8c
10. Master v9 for autonomous multi-agent collaboration

**Note:** v8 is split into three progressive sub-versions (v8a -> v8b -> v8c), each adding one concept. This makes the jump from v7 more gradual.

## Learning Progression

```
v0(196) -> v1(417) -> v2(531) -> v3(623) -> v4(783)
   |          |          |          |          |
 Bash      4 Tools    Planning   Subagent   Skills

-> v5(896) -> v6(1075) -> v7(1142) -> v8a(1395) -> v8b(1557) -> v8c(1612) -> v9(1683)
     |           |            |            |            |            |            |
 Compress     Tasks      Background   Foundation   Messaging   Coordination  Autonomous
```

## Quick Start

```bash
# Clone the repository
git clone https://github.com/shareAI-lab/learn-claude-code
cd learn-claude-code

# Install dependencies
pip install -r requirements.txt

# Configure API key
cp .env.example .env
# Edit .env with your ANTHROPIC_API_KEY

# Run any version
python v0_bash_agent.py         # Minimal (start here!)
python v1_basic_agent.py        # Core agent loop
python v2_todo_agent.py         # + Todo planning
python v3_subagent.py           # + Subagents
python v4_skills_agent.py       # + Skills
python v5_compression_agent.py  # + Context compression
python v6_tasks_agent.py        # + Task system
python v7_background_agent.py   # + Background tasks
python v8a_team_foundation.py  # + Team foundation
python v8b_messaging.py        # + Team messaging
python v8c_coordination.py     # + Team coordination
python v9_autonomous_agent.py  # + Autonomous teams
```

## Running Tests

```bash
# Run full test suite
python tests/run_all.py

# Run unit tests only
python tests/test_unit.py

# Run tests for a specific version
python -m pytest tests/test_v8a.py tests/test_v8b.py tests/test_v8c.py -v
```

## The Core Pattern

Every coding agent is just this loop:

```python
while True:
    response = model(messages, tools)
    if response.stop_reason != "tool_use":
        return response.text
    results = execute(response.tool_calls)
    messages.append(results)
```

That's it. The model calls tools until done. Everything else is refinement.

## Version Comparison

| Version | Lines | Tools | Core Addition | Key Insight |
|---------|-------|-------|---------------|-------------|
| [v0](./v0_bash_agent.py) | ~196 | bash | Recursive subagents | One tool is enough |
| [v1](./v1_basic_agent.py) | ~417 | bash, read, write, edit | Core loop | Model as Agent |
| [v2](./v2_todo_agent.py) | ~531 | +TodoWrite | Explicit planning | Constraints enable complexity |
| [v3](./v3_subagent.py) | ~623 | +Task | Context isolation | Clean context = better results |
| [v4](./v4_skills_agent.py) | ~783 | +Skill | Knowledge loading | Expertise without retraining |
| [v5](./v5_compression_agent.py) | ~896 | +ContextManager | 3-layer compression | Forgetting enables infinite work |
| [v6](./v6_tasks_agent.py) | ~1075 | +TaskCreate/Get/Update/List | Persistent tasks | Sticky notes to kanban |
| [v7](./v7_background_agent.py) | ~1142 | +TaskOutput/TaskStop | Background execution | Serial to parallel |
| [v8a](./v8a_team_foundation.py) | ~1395 | +TeamCreate/TeamDelete | Team foundation | Building the team structure |
| [v8b](./v8b_messaging.py) | ~1557 | +SendMessage/Inbox | Team messaging | Communication channels |
| [v8c](./v8c_coordination.py) | ~1612 | +SharedBoard/Protocol | Team coordination | Command to collaboration |
| [v9](./v9_autonomous_agent.py) | ~1683 | +Idle cycle/auto-claim | Autonomous teams | Collaboration to self-organization |

## Sub-Mechanism Guide

Each version introduces one core class, but the real learning is in the sub-mechanisms. This map helps you find specific concepts:

| Sub-Mechanism | Version | Key Code | What to Look For |
|---------------|---------|----------|------------------|
| **Agent loop** | v0-v1 | `agent_loop()` | The `while tool_use` loop pattern |
| **Tool dispatch** | v1 | `process_tool_call()` | How tool_use blocks map to functions |
| **Explicit planning** | v2 | `TodoManager` | Single `in_progress` constraint, system reminders |
| **Context isolation** | v3 | `run_subagent()` | Fresh message list per subagent |
| **Tool filtering** | v3 | `AGENT_TYPES` | Explore agents get read-only tools |
| **Skill injection** | v4 | `SkillLoader` | Content prepended to system prompt |
| **Microcompact** | v5 | `ContextManager.microcompact()` | Old tool outputs replaced with placeholders |
| **Auto-compact** | v5 | `ContextManager.auto_compact()` | 85.3% threshold (formula-based) triggers API summarization |
| **Large output handling** | v5 | `ContextManager.handle_large_output()` | >40K tokens saved to disk, preview returned |
| **Transcript persistence** | v5 | `ContextManager.save_transcript()` | Full history appended to `.jsonl` |
| **Task CRUD** | v6 | `TaskManager` | create/get/update/list with JSON persistence |
| **Dependency graph** | v6 | `addBlocks/addBlockedBy` | Completion auto-unblocks dependents |
| **Background execution** | v7 | `BackgroundManager.run_in_background()` | Thread-based, immediate task_id return |
| **ID prefix convention** | v7 | `_PREFIXES` | `b`=bash, `a`=agent (v8 adds `t`=teammate) |
| **Notification bus** | v7 | `drain_notifications()` | Queue drained before each API call |
| **Notification injection** | v7 | attachment-based notification | Injected into last user message |
| **Teammate lifecycle** | v8a | `_teammate_loop()` | Work -> check inbox -> exit pattern |
| **Tool scoping** | v8a | `TEAMMATE_TOOLS` | Teammates get 9 tools (no TeamCreate/Delete/Task/Skill) |
| **File-based inbox** | v8b | `send_message()/check_inbox()` | JSONL format, per-teammate files |
| **Message protocol** | v8b | `MESSAGE_TYPES` | 5 types: message, broadcast, shutdown_req/resp, plan_approval |
| **Shutdown protocol** | v8c | `SharedBoard/Protocol` | Graceful shutdown and plan approval |
| **Idle cycle** | v9 | `_teammate_loop()` | active -> idle -> poll inbox -> wake -> active |
| **Task claiming** | v9 | `_teammate_loop()` | Idle teammates auto-claim unclaimed tasks |
| **Identity preservation** | v9 | `auto_compact` + identity | Teammate name/role re-injected after compression |

## File Structure

```
learn-claude-code/
|-- v0_bash_agent.py          # ~196 lines: 1 tool, recursive subagents
|-- v0_bash_agent_mini.py     # ~16 lines: extreme compression
|-- v1_basic_agent.py         # ~417 lines: 4 tools, core loop
|-- v2_todo_agent.py          # ~531 lines: + TodoManager
|-- v3_subagent.py            # ~623 lines: + Task tool, agent registry
|-- v4_skills_agent.py        # ~783 lines: + Skill tool, SkillLoader
|-- v5_compression_agent.py   # ~896 lines: + ContextManager, 3-layer compression
|-- v6_tasks_agent.py         # ~1075 lines: + TaskManager, CRUD with dependencies
|-- v7_background_agent.py    # ~1142 lines: + BackgroundManager, parallel execution
|-- v8a_team_foundation.py  # ~1395 lines: + TeammateManager, team lifecycle
|-- v8b_messaging.py        # ~1557 lines: + SendMessage, inbox, message protocol
|-- v8c_coordination.py     # ~1612 lines: + Shared task board, shutdown protocol
|-- v9_autonomous_agent.py    # ~1683 lines: + Idle cycle, auto-claim, identity preservation
|-- skills/                   # Example skills (pdf, code-review, mcp-builder, agent-builder)
|-- docs/                     # Technical documentation (EN + ZH + JA)
|-- articles/                 # Blog-style articles (ZH)
+-- tests/                    # Unit, feature, and integration tests
=======
Phase 1: THE LOOP                    Phase 2: PLANNING & KNOWLEDGE
==================                   ==============================
s01  The Agent Loop          [1]     s03  TodoWrite               [5]
     while + stop_reason                  TodoManager + nag reminder
     |                                    |
     +-> s02  Tool Use            [4]     s04  Subagents            [5]
              dispatch map: name->handler     fresh messages[] per child
                                              |
                                         s05  Skills               [5]
                                              SKILL.md via tool_result
                                              |
                                         s06  Context Compact      [5]
                                              3-layer compression

Phase 3: PERSISTENCE                 Phase 4: TEAMS
==================                   =====================
s07  Tasks                   [8]     s09  Agent Teams             [9]
     file-based CRUD + deps graph         teammates + JSONL mailboxes
     |                                    |
s08  Background Tasks        [6]     s10  Team Protocols          [12]
     daemon threads + notify queue        shutdown + plan approval FSM
                                          |
                                     s11  Autonomous Agents       [14]
                                          idle cycle + auto-claim
                                     |
                                     s12  Worktree Isolation      [16]
                                          task coordination + optional isolated execution lanes

                                     [N] = number of tools
```

## Architecture

```
learn-claude-code/
|
|-- agents/                        # Python reference implementations (s01-s12 + s_full capstone)
|-- docs/{en,zh,ja}/               # Mental-model-first documentation (3 languages)
|-- web/                           # Interactive learning platform (Next.js)
|-- skills/                        # Skill files for s05
+-- .github/workflows/ci.yml      # CI: typecheck + build
>>>>>>> upstream/main
```

## Documentation

<<<<<<< HEAD
### Technical Tutorials (docs/)

- [v0: Bash is All You Need](./docs/v0-bash-is-all-you-need.md)
- [v1: Model as Agent](./docs/v1-model-as-agent.md)
- [v2: Structured Planning](./docs/v2-structured-planning.md)
- [v3: Subagent Mechanism](./docs/v3-subagent-mechanism.md)
- [v4: Skills Mechanism](./docs/v4-skills-mechanism.md)
- [v5: Context Compression](./docs/v5-context-compression.md)
- [v6: Tasks System](./docs/v6-tasks-system.md)
- [v7: Background Tasks](./docs/v7-background-tasks.md)
- [v8: Team Messaging](./docs/v8-team-messaging.md)
- [v9: Autonomous Teams](./docs/v9-autonomous-teams.md)

### Articles

See [articles/](./articles/) for blog-style explanations.

## Using the Skills System

### Example Skills Included

| Skill | Purpose |
|-------|---------|
| [agent-builder](./skills/agent-builder/) | Meta-skill: how to build agents |
| [code-review](./skills/code-review/) | Systematic code review methodology |
| [pdf](./skills/pdf/) | PDF manipulation patterns |
| [mcp-builder](./skills/mcp-builder/) | MCP server development |

### Scaffold a New Agent

```bash
# Use the agent-builder skill to create a new project
python skills/agent-builder/scripts/init_agent.py my-agent

# Specify complexity level
python skills/agent-builder/scripts/init_agent.py my-agent --level 0  # Minimal
python skills/agent-builder/scripts/init_agent.py my-agent --level 1  # 4 tools
```

### Install Skills for Production

```bash
# Kode CLI (recommended)
kode plugins install https://github.com/shareAI-lab/shareAI-skills

# Claude Code
claude plugins install https://github.com/shareAI-lab/shareAI-skills
```

## Configuration

```bash
# .env file options
ANTHROPIC_API_KEY=sk-ant-xxx      # Required: Your API key
ANTHROPIC_BASE_URL=https://...    # Optional: For API proxies
MODEL_ID=claude-sonnet-4-5-20250929  # Optional: Model selection
```

## Related Projects

| Repository | Description |
|------------|-------------|
| [Kode](https://github.com/shareAI-lab/Kode) | Production-ready open source agent CLI |
| [shareAI-skills](https://github.com/shareAI-lab/shareAI-skills) | Production skills collection |
| [Agent Skills Spec](https://agentskills.io/specification) | Official specification |

## Philosophy

> **The model is 80%. Code is 20%.**

Modern agents like Kode and Claude Code work not because of clever engineering, but because the model is trained to be an agent. Our job is to give it tools and stay out of the way.

## Contributing

Contributions are welcome! Please feel free to submit issues and pull requests.

- Add new example skills in `skills/`
- Improve documentation in `docs/`
- Report bugs or suggest features via [Issues](https://github.com/shareAI-lab/learn-claude-code/issues)
=======
Mental-model-first: problem, solution, ASCII diagram, minimal code.
Available in [English](./docs/en/) | [中文](./docs/zh/) | [日本語](./docs/ja/).

| Session | Topic | Motto |
|---------|-------|-------|
| [s01](./docs/en/s01-the-agent-loop.md) | The Agent Loop | *One loop & Bash is all you need* |
| [s02](./docs/en/s02-tool-use.md) | Tool Use | *Adding a tool means adding one handler* |
| [s03](./docs/en/s03-todo-write.md) | TodoWrite | *An agent without a plan drifts* |
| [s04](./docs/en/s04-subagent.md) | Subagents | *Break big tasks down; each subtask gets a clean context* |
| [s05](./docs/en/s05-skill-loading.md) | Skills | *Load knowledge when you need it, not upfront* |
| [s06](./docs/en/s06-context-compact.md) | Context Compact | *Context will fill up; you need a way to make room* |
| [s07](./docs/en/s07-task-system.md) | Tasks | *Break big goals into small tasks, order them, persist to disk* |
| [s08](./docs/en/s08-background-tasks.md) | Background Tasks | *Run slow operations in the background; the agent keeps thinking* |
| [s09](./docs/en/s09-agent-teams.md) | Agent Teams | *When the task is too big for one, delegate to teammates* |
| [s10](./docs/en/s10-team-protocols.md) | Team Protocols | *Teammates need shared communication rules* |
| [s11](./docs/en/s11-autonomous-agents.md) | Autonomous Agents | *Teammates scan the board and claim tasks themselves* |
| [s12](./docs/en/s12-worktree-task-isolation.md) | Worktree + Task Isolation | *Each works in its own directory, no interference* |

## What's Next -- from understanding to shipping

After the 12 sessions you understand how an agent works inside out. Two ways to put that knowledge to work:

### Kode Agent CLI -- Open-Source Coding Agent CLI

> `npm i -g @shareai-lab/kode`

Skill & LSP support, Windows-ready, pluggable with GLM / MiniMax / DeepSeek and other open models. Install and go.

GitHub: **[shareAI-lab/Kode-cli](https://github.com/shareAI-lab/Kode-cli)**

### Kode Agent SDK -- Embed Agent Capabilities in Your App

The official Claude Code Agent SDK communicates with a full CLI process under the hood -- each concurrent user means a separate terminal process. Kode SDK is a standalone library with no per-user process overhead, embeddable in backends, browser extensions, embedded devices, or any runtime.

GitHub: **[shareAI-lab/Kode-agent-sdk](https://github.com/shareAI-lab/Kode-agent-sdk)**

---

## Sister Repo: from *on-demand sessions* to *always-on assistant*

The agent this repo teaches is **use-and-discard** -- open a terminal, give it a task, close when done, next session starts blank. That is the Claude Code model.

[OpenClaw](https://github.com/openclaw/openclaw) proved another possibility: on top of the same agent core, two mechanisms turn the agent from "poke it to make it move" into "it wakes up every 30 seconds to look for work":

- **Heartbeat** -- every 30s the system sends the agent a message to check if there is anything to do. Nothing? Go back to sleep. Something? Act immediately.
- **Cron** -- the agent can schedule its own future tasks, executed automatically when the time comes.

Add multi-channel IM routing (WhatsApp / Telegram / Slack / Discord, 13+ platforms), persistent context memory, and a Soul personality system, and the agent goes from a disposable tool to an always-on personal AI assistant.

**[claw0](https://github.com/shareAI-lab/claw0)** is our companion teaching repo that deconstructs these mechanisms from scratch:

```
claw agent = agent core + heartbeat + cron + IM chat + memory + soul
```

```
learn-claude-code                   claw0
(agent runtime core:                (proactive always-on assistant:
 loop, tools, planning,              heartbeat, cron, IM channels,
 teams, worktree isolation)          memory, soul personality)
```

## About
<img width="260" src="https://github.com/user-attachments/assets/fe8b852b-97da-4061-a467-9694906b5edf" /><br>

Scan with Wechat to fellow us,  
or fellow on X: [shareAI-Lab](https://x.com/baicai003)  
>>>>>>> upstream/main

## License

MIT

---

<<<<<<< HEAD
**Model as Agent. That's the whole secret.**

[@baicai003](https://x.com/baicai003) | [shareAI Lab](https://github.com/shareAI-lab)
=======
**The model is the agent. Our job is to give it tools and stay out of the way.**
>>>>>>> upstream/main
