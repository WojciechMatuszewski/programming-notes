# Mega Dev

## Mechanics Behind Generate AI

- During training, the labs are using "markers" (so far I've only read about HTML-like tags, for example `<|im_start|>`) which are used to denote where a given message start, ends and _who_ wrote it.

  - Think of those as punctuation. Without those special "tags", the model would have a really hard time knowing where a given sentence starts and ends. In addition, it would have really hard time to distinguish between the _system prompt_ and the _user input_.

- Tools, which you know very well, allow us to _extend_ the model capability that go beyond it's original training data set.

- The _agent_ is "just" calling the LLM in a loop and appending results of the tool calls into the history.

## Kickoff

- When introducing agents to a codebase, **consider running a "can you efficiently work in this repo" sweep**. You would be surprised how many things the agent can find. **This is an excellent way of finding all the inefficiencies in the repository**, especially as it relates to agents.

- Agents can help you improve your harness. Agents can improve your agents. Agents can improve themselves.

  - **The core principle to remember is that YOU CAN AUTOMATE A LOT OF IMPROVEMENTS**. Let agents run themselves. Let them find issues and fix those!

- Agents flourish in tedious. Previously, we steered away from tedious because we had better things to do, but those tedious things are also valuable.

  - Now you have a superpower: you can offload the tedious to agents and focus on the most important parts.

## Steering Models Through Context

- Context is _very_ important.

  - Every compaction is lossy. Keep that in mind for long-lived agents.

- Your `CLAUDE.md` or `AGENTS.md` should, most likely **not contain any code paths**.

  - The models are quite capable and can _discover_ the code-paths themselves.

  - **What you want to put there: things that are not obvious**, things that you see agents often miss.

- The `description` field in the front-matter for skill files is **to tell the model what the skill does and WHEN the skill should be used**. All of this should be written in a short format. It's imperative that the descriptions are short as those are get pulled into the model context.

## Building Knowledge Systems

- **Your "knowledge system" should just be your codebase**.

- The best thing you can do is to make your codebase simpler and easier to understand for _you_ and _your agent_.

  - If you have consistent patterns in your codebase, your agent will follow them.

  - If you have a good file structure established, your agent will follow that structure.

  The point here is that agents have a hard time noticing that the pattern is _not_ optimal (or that the code is repeated) than you.

- **Make it explicit that agents can "ask dumb questions"**. While reading the "thinking blocks" of LLMs, I've noticed that they can _imply_ a lot of things. Things that they should really ask about. Make it easy for them to _ask_!
