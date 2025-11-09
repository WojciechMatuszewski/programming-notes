# Elite AI-Assisted Coding

> Taking notes based on [this course](https://elite-ai-assisted-coding.dev/p/course)

## Introduction to AI-Assisted Software Development

This lesson was a very brief and quite shallow introduction to the course contents. Nothing major to note.

## Context Engineering & Frictionless Setup

- TIL about `jina.ai` and how easily you can transform any page into markdown.

- Symlinks are making a comeback! Since we have so many configuration tools like `CLAUDE.md`, `AGENTS.md`, `.mcp.json`, or `.cursor/mcp.json`, using symlinks to ensure they have the same contents is quite nice.

  - You can also use a tool like [`ruler`](https://github.com/intellectronica/ruler) for this.

- Asking the agent to plan features is great, but **you can also ask the agent to plan for refactoring work**.

  - Force it to provide you with a cost/benefit analysis. While I do not personally trust it, it will most likely make the output better.

## HW1

- The `plan` mode in Claude Code is nice.

  - **Make sure to _really_ read the plan**. I noticed that **the plan Claude proposes is often overly ambitious**.

    - Just like you would reduce scope of a feature, you should reduce the scope of the plan.

- I wish there was an easy way to _stop_ Claude after it completes a TODO.

  - Claude Code has the `TodoWrite` built-in tool, but I could not configure a hook to stop Claude when it uses it. I believe the `auto-accept edits` mode does not allow for it?

  - **Manual `todo.md` files are where it's at**. The `TodoWrite` tool is a bit buggy, and it's quite "hidden" from you. I like to have a view into each todo in my editor.

- **Using AI tools is not only about writing code**. One use-case that I really like is asking the LLM to list things that we ought to consider refactoring and order them by priority.
