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

- I wish there were an easy way to _stop_ Claude after it completes a TODO.

  - Claude Code has the `TodoWrite` built-in tool, but I could not configure a hook to stop Claude when it uses it. I believe the `auto-accept edits` mode does not allow for it?

  - **Manual `todo.md` files are where it's at**. The `TodoWrite` tool is a bit buggy, and it's quite "hidden" from you. I like to have a view into each todo in my editor.

- **Using AI tools is not only about writing code**. One use-case that I really like is asking the LLM to list things that we ought to consider refactoring and order them by priority.

- **You can ask the LLM to ask you questions based on the plan to refine it**.

  - Remember, the plan is quite important. The more refined it is, the better.

## Working Incrementally with AI Todos and Git

- Forcing the LLM to provide you with "user test" at the end of the implementation step is a good way to ensure you are aligned with what the LLM wants to implement.

  - If the "user test" is off, it means that, despite correct TODOs, the LLM's understanding of the work might be different than yours.

  - You can also ask the LLM to come up with different testing plans at this stage. Perhaps listing what test cases it will write and so on.

## Optional Part 2 Homework – Live Practice Session

This lecture focused on working with MCP servers.

1. Adding MCP servers.
2. Using MCP servers.
3. Developing MCP servers.

A good introduction if you are unfamiliar with those.

## Defining Good Code for LLM Context

You will most likely want to include some instructions about code quality in the `CLAUDE.md` or `AGENTS.md` files.

If you are on a team that already defines these in some form of shared documentation, you are in luck. You can probably paste that document into the relevant agent file and call it a day.

Otherwise, **consider creating a draft document and asking the LLM what it would improve about that document, or what questions it needs answered**. This way, the LLM becomes a co-creator of the document.

In this lecture, Issac presents a document containing a set of questions related to _abstractions_, _code duplication_, and so on. Answering those questions yourself, and asking the LLM if it needs clarification, is a good way to create a "code quality" document for the LLM.

Things Issac mentioned that I agree with:

1. Duplication is most likely okay. **Consider allowing the agent to duplicate the code, and then refactor it yourself**. This way, you are in charge of writing the most important pieces of code—abstractions that will be used across the codebase.

2. Aim for "deep" functions that hide a lot of functionality. Again, you can refactor them later yourself.

3. **Make sure the agent does not try to add defensive error handling**. Error handling is _very critical_ in any production application. It is best that _you_ write the logic for it so you understand how errors propagate.

4. No mocks or placeholders. I've noticed this myself as well. Agents often create mocks for mocks and placeholders everywhere in test code. That's not good. This practice breeds brittle tests.

There are more things to think about. This is not an exhaustive list.

## Using an Agent to Identify Tech Debt

In this session, Issac showcased asking the LLM to find opportunities for refactoring and to rank them based on (subjective) priority. You are most likely familiar with the tech debt in your codebase, so the LLM's output should not surprise you.

This way, you have an overview of things you _might_ want to tackle. **I find this technique particularly effective for checking if there are any quick wins that, when addressed, would improve code quality**.

The LLM helps you translate the "knowledge in your head" into a document that you can share with others.

## Wrapping up

This course was okay at the begging, but started to become a bit too shallow for me at the end.

If you are starting your AI journey, take this course. If you have been working with AI for a while, consider watching first couple of practical lessons, especially the ones about creating a good `plan.md` file.
