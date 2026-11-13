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

  The point here is that agents have a harder time noticing that a pattern is _not_ optimal (or that code is repeated) than you do.

- **Make it explicit that agents can "ask dumb questions"**. While reading the "thinking blocks" of LLMs, I've noticed that they can _imply_ a lot of things. Things that they should really ask about. Make it easy for them to _ask_!

## Working with AI Harnesses

- The harness is not the UI or the MCPs; it's the agentic loop and its capabilities.

  - You _should_ extend it with the things that are important to your workflows, like various MCPs and skills.

- Harness lock-in is a thing, especially when we talk about big AI labs. Try moving from Claude Code to Codex when you have a robust setup in each. Good luck!

- The harness you use should allow some degree of customization.

  - Customizing things very extensively is a double-edged sword. Your harness might be great, but what if you need to switch? I personally do not buy this argument, but I can see it being valid for some.

## Building Workflows, Not Tool Stacks

In this lesson, Theo talked about the markdown files being pretty much "the code" at this point.

By "workflows," they mean skills that you can compose with each other.

Some good points:

1. Audit and see how agents are using the skills you already have. Perhaps there is something to fix?

2. Doing something with the agent (talking back-and-forth) and then writing a skill file as the artifact from that back-and-forth is a good way to create a high-quality skill.

---

So far, we talked about the basics. That's good, as those are always needed, and it's nice to remind ourselves about how things work on the most basic level. Theo's content focused on the _why_ and not the _how_. Interesting, and I would say a valid choice, given that nowadays the _how_ could be offloaded.

The most important things we talked about:

1. The "agentic loop" is nothing but a glorified model invocation call in a loop.

2. Tools give the model lots of useful capabilities that greatly increase its usefulness. Remember that MCP servers also expose tools.

3. The _harness_ is the layer around the model **and something that gets pretty much all the attention nowadays**.

4. **The better your codebase, the better the agent output**. Engineering principles still apply.

---

## Agentic Primitives

- You have the _agent loop_, the _instructions_ you give to the agent, the _context and the memory_ the agent has, and the _tools_ available to the agent.

- **Do not treat the agent as a black box**. Instrumentation is critical. Without traces, it will be very hard for you to understand what happened and why it happened.

- In recent months, a new concept has emerged: the so-called _context engineering_.

  - It's nothing more than a proper acknowledgement that **what data you push into your agent's context, and how much of it, is very important and greatly influences how the agent works**. This is not only about the system prompt. Keep in mind that all the tool calls, and their output, also land in the context.

  The more stuff in the context, the bigger the payload, and thus the more expensive and slower the response is (that is why we have [message threads](https://platform.claude.com/docs/en/build-with-claude/message-threads) in some APIs).

- Context is not infinite. The size of the context has gotten bigger, but that does not mean you can cram literally everything in there. That is why you might want to give the agent a "memory" functionality. Doing so opens a whole can of worms related to extracting only the relevant data from the memory, and so on, so be careful here!

## Building an Agent

- Anything you _can_ do in code, you _should_.

  - For example, the agent should not decide whether the user _can_ or _can't_ access some data, all of this should be handled in code!

Overall, I'm unsure if this section of the workshop is valuable at all. It's a big code walkthrough, which does not provide enough substance to be useful.

## Q&A
