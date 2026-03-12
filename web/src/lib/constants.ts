export const VERSION_ORDER = [
<<<<<<< HEAD
  "v0_mini", "v0", "v1", "v2", "v3", "v4", "v5", "v6", "v7", "v8a", "v8b", "v8c", "v9"
] as const;

// Only show these in the learning path (skip v0_mini)
export const LEARNING_PATH = [
  "v0", "v1", "v2", "v3", "v4", "v5", "v6", "v7", "v8a", "v8b", "v8c", "v9"
] as const;
=======
  "s01", "s02", "s03", "s04", "s05", "s06", "s07", "s08", "s09", "s10", "s11", "s12"
] as const;

export const LEARNING_PATH = VERSION_ORDER;
>>>>>>> upstream/main

export type VersionId = typeof LEARNING_PATH[number];

export const VERSION_META: Record<string, {
  title: string;
  subtitle: string;
  coreAddition: string;
  keyInsight: string;
  layer: "tools" | "planning" | "memory" | "concurrency" | "collaboration";
  prevVersion: string | null;
}> = {
<<<<<<< HEAD
  v0: { title: "Bash Agent", subtitle: "Bash is All You Need", coreAddition: "Single-tool agent loop", keyInsight: "One tool (bash) is enough to be useful", layer: "tools", prevVersion: null },
  v1: { title: "Basic Agent", subtitle: "The Model IS the Agent", coreAddition: "Multi-tool + dispatcher", keyInsight: "4 tools beat 1: read, write, edit, bash", layer: "tools", prevVersion: "v0" },
  v2: { title: "Todo Agent", subtitle: "Make Plans Visible", coreAddition: "TodoManager for planning", keyInsight: "Visible plans improve task completion", layer: "planning", prevVersion: "v1" },
  v3: { title: "Subagent", subtitle: "Divide and Conquer", coreAddition: "Agent registry + Task tool", keyInsight: "Context isolation prevents confusion", layer: "planning", prevVersion: "v2" },
  v4: { title: "Skills Agent", subtitle: "Knowledge Externalization", coreAddition: "SkillLoader + dynamic injection", keyInsight: "Skills inject via tool_result, not system prompt", layer: "planning", prevVersion: "v3" },
  v5: { title: "Compression Agent", subtitle: "Strategic Forgetting", coreAddition: "3-layer context compression", keyInsight: "Forgetting old results enables infinite work", layer: "memory", prevVersion: "v4" },
  v6: { title: "Tasks Agent", subtitle: "Shared Task Board", coreAddition: "TaskManager with CRUD + deps", keyInsight: "File-based persistence outlives process memory", layer: "planning", prevVersion: "v5" },
  v7: { title: "Background Agent", subtitle: "Fire and Forget", coreAddition: "BackgroundManager + notifications", keyInsight: "Non-blocking execution via threads + queue", layer: "concurrency", prevVersion: "v6" },
  v8a: { title: "Team Foundation", subtitle: "From Commands to Collaboration", coreAddition: "TeammateManager + team identity", keyInsight: "Persistent teammates vs one-shot subagents", layer: "collaboration", prevVersion: "v7" },
  v8b: { title: "Team Messaging", subtitle: "Inbox-Based Communication", coreAddition: "File-based inbox + 5 message types", keyInsight: "Async JSONL inboxes decouple communication", layer: "collaboration", prevVersion: "v8a" },
  v8c: { title: "Team Coordination", subtitle: "Shared Board + Protocol", coreAddition: "Shutdown protocol + plan approval", keyInsight: "Dependency graph prevents duplicate work", layer: "collaboration", prevVersion: "v8b" },
  v9: { title: "Autonomous Agent", subtitle: "Teammates That Think", coreAddition: "Idle cycle + auto-claiming", keyInsight: "Polling + timeout makes teammates autonomous", layer: "collaboration", prevVersion: "v8c" },
};

export const LAYERS = [
  { id: "tools" as const, label: "Tools & Execution", color: "#3B82F6", versions: ["v0", "v1"] },
  { id: "planning" as const, label: "Planning & Coordination", color: "#10B981", versions: ["v2", "v3", "v4", "v6"] },
  { id: "memory" as const, label: "Memory Management", color: "#8B5CF6", versions: ["v5"] },
  { id: "concurrency" as const, label: "Concurrency", color: "#F59E0B", versions: ["v7"] },
  { id: "collaboration" as const, label: "Collaboration", color: "#EF4444", versions: ["v8a", "v8b", "v8c", "v9"] },
=======
  s01: { title: "The Agent Loop", subtitle: "Bash is All You Need", coreAddition: "Single-tool agent loop", keyInsight: "The minimal agent kernel is a while loop + one tool", layer: "tools", prevVersion: null },
  s02: { title: "Tools", subtitle: "One Handler Per Tool", coreAddition: "Tool dispatch map", keyInsight: "The loop stays the same; new tools register into the dispatch map", layer: "tools", prevVersion: "s01" },
  s03: { title: "TodoWrite", subtitle: "Plan Before You Act", coreAddition: "TodoManager + nag reminder", keyInsight: "An agent without a plan drifts; list the steps first, then execute", layer: "planning", prevVersion: "s02" },
  s04: { title: "Subagents", subtitle: "Clean Context Per Subtask", coreAddition: "Subagent spawn with isolated messages[]", keyInsight: "Subagents use independent messages[], keeping the main conversation clean", layer: "planning", prevVersion: "s03" },
  s05: { title: "Skills", subtitle: "Load on Demand", coreAddition: "SkillLoader + two-layer injection", keyInsight: "Inject knowledge via tool_result when needed, not upfront in the system prompt", layer: "planning", prevVersion: "s04" },
  s06: { title: "Compact", subtitle: "Three-Layer Compression", coreAddition: "micro-compact + auto-compact + archival", keyInsight: "Context will fill up; three-layer compression strategy enables infinite sessions", layer: "memory", prevVersion: "s05" },
  s07: { title: "Tasks", subtitle: "Task Graph + Dependencies", coreAddition: "TaskManager with file-based state + dependency graph", keyInsight: "A file-based task graph with ordering, parallelism, and dependencies -- the coordination backbone for multi-agent work", layer: "planning", prevVersion: "s06" },
  s08: { title: "Background Tasks", subtitle: "Background Threads + Notifications", coreAddition: "BackgroundManager + notification queue", keyInsight: "Run slow operations in the background; the agent keeps thinking ahead", layer: "concurrency", prevVersion: "s07" },
  s09: { title: "Agent Teams", subtitle: "Teammates + Mailboxes", coreAddition: "TeammateManager + file-based mailbox", keyInsight: "When one agent can't finish, delegate to persistent teammates via async mailboxes", layer: "collaboration", prevVersion: "s08" },
  s10: { title: "Team Protocols", subtitle: "Shared Communication Rules", coreAddition: "request_id correlation for two protocols", keyInsight: "One request-response pattern drives all team negotiation", layer: "collaboration", prevVersion: "s09" },
  s11: { title: "Autonomous Agents", subtitle: "Scan Board, Claim Tasks", coreAddition: "Task board polling + timeout-based self-governance", keyInsight: "Teammates scan the board and claim tasks themselves; no need for the lead to assign each one", layer: "collaboration", prevVersion: "s10" },
  s12: { title: "Worktree + Task Isolation", subtitle: "Isolate by Directory", coreAddition: "Composable worktree lifecycle + event stream over a shared task board", keyInsight: "Each works in its own directory; tasks manage goals, worktrees manage directories, bound by ID", layer: "collaboration", prevVersion: "s11" },
};

export const LAYERS = [
  { id: "tools" as const, label: "Tools & Execution", color: "#3B82F6", versions: ["s01", "s02"] },
  { id: "planning" as const, label: "Planning & Coordination", color: "#10B981", versions: ["s03", "s04", "s05", "s07"] },
  { id: "memory" as const, label: "Memory Management", color: "#8B5CF6", versions: ["s06"] },
  { id: "concurrency" as const, label: "Concurrency", color: "#F59E0B", versions: ["s08"] },
  { id: "collaboration" as const, label: "Collaboration", color: "#EF4444", versions: ["s09", "s10", "s11", "s12"] },
>>>>>>> upstream/main
] as const;
