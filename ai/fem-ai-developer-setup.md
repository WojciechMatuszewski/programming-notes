# FrontendMasters AI Developer Setup

## Learnings

- AI can generate all kinds of code query quickly. **This includes technical debt**.

  - This starts to be a problem when people do not take time to read the LLM output and refactor it.

- If the LLM does not have context, it is guessing.

  - If it's not in the context, the LLM is very likely to hallucinate.

- **Writing your plan first will save you lots of time.**

- Commits and Git are even more important.

  - You might loose a lot of progress without "checkpoints".

- Using LLM to explain the code is a very powerful way to catch up on codebases.

  - While it might not know about all the decisions and circumstances that went into writing a given piece of code, it can explain to you what the code does and which areas of the codebase are more important than others.

- Most of the AI coding tools support _rules_.

  - Think of rules as _instructions_ for the LLM that is applicable to the codebase you are working on. There are also global rules that apply to every project.

- There are a lot of options for _background agents_ – agents that will pull your repo and work on it in the _background_.

  - You have to turn off privacy settings and allow the bot to pull your repo. That means giving a lot of permissions to a 3rd party. I'm not a fan.

- Claude Code has a notion of _subagents_.

  - **A _subagent_ context window is separate of the "main" agent**.

  - Running multiple _subagents_ in parallel allows you to work on things faster.

    - Think having a "junior-engineer" _subagent_ whose role is to ask questions and then having a "architect" _subagent_ whose role is to answer the questions and amend the design if necessary.

- With AI generated code, having a very strict eslint rules is even more important than ever.

  - LLMs can ignore the rules you specified in any "configuration" files like `CLAUDE.md`.
