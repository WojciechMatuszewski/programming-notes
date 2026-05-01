# AI Agents

## Multi-Agent Collaboration in Production

> Reading and taking notes from [this blog post](https://www.taskade.com/blog/multi-agent-production)

### On Agent Memory

The article claims that the **agent is only as good as its memory system**.

Given the importance the _context_ and _tools_ have, I agree – memory, after all, is part of "context engineering".

What's fascinating to me (and after reading about it sounds quite obvious), is the fact that **you might want to deploy multiple "types" of memory for your agents to consume**.

- The "core" memory which includes things like the agents' personality and identity.

- The "learning" memory which is computed over-time and includes things like user preferences and patterns.

- The "working" memory which is the _current conversation_. Keep in mind that you will need to summarize the past messages at some point to prevent overflowing the agent context.

- The "navigation" memory which tells the model _where_ in the product they are. I think this is quite situational. To me, this should answer the question of "what resource I'm working on, and what resources I was working on in this conversation".

- The "reference" memory which is _external_ knowledge. Basically RAG.

**The article says that, by using this taxonomy, they greatly improved the output and functionality of their agents**.

### On Model Selection

Running all tasks on the most expensive model will, most likely, yield best results, but it is not sustainable.

You have to start thinking about "routing" a given task to an appropriate model (note about the UX here. We most likely want the user to have an ability to choose the model as well!).

The article talks about "model credit system" which detects which model is the best given the users current credit "score" and the request complexity.

#### Switching Models

**The article states that switching models mid-task produces worse results than either model alone**. I would like to see some more concrete data on this. I wonder if it's true when "switching up", so from a less powerful to more powerful model.

### Multi-Agent Team Chat

Three patterns:

- The _fan-out_ when you require breath. A good example here would be a research task. You can spawn multiple agents that research different "domains" of the problem.

- The _chain_ when you require sequential processing – when output of agent A depends on agent B. Imagine one agent pulling some data, and then the other creating dashboard based on this data. **Notice that agents have different roles here**.

- The _debate_ when you need a consensus over something. Ideally you pair a "skeptic" with "proponent" so you have balanced arguments on each side. You will need another agent to present you with results (or for you to come up with conclusion yourself).

### Guardrails and Loops

I've experienced this myself first-hand. The agent was executing the same set of tools over and over again unaware of what is going on.

The article suggest following mitigations:

- Injecting a corrective instruction into the conversation. This works well, but you have to have the ability to do so in the first place. Depending on how you design the system, you might need to interrupt your agent programmatically to do it.

- Force a summary exit. You stop the agent manually and instruct it to summarize what it done so far. Partial summary is better than no summary. **If you report this failure case transparently, you will build trust with the user**. There is nothing worse than user seeing an agent spinning wheels and burning _their_ money!

This is quite critical to nail down right. I would argue that you need a very extensive test coverage here.

Continue on "Context Window Management".
