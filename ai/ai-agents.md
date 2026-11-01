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

The article suggest following mitigation:

- Injecting a corrective instruction into the conversation. This works well, but you have to have the ability to do so in the first place. Depending on how you design the system, you might need to interrupt your agent programmatically to do it.

- Force a summary exit. You stop the agent manually and instruct it to summarize what it done so far. Partial summary is better than no summary. **If you report this failure case transparently, you will build trust with the user**. There is nothing worse than user seeing an agent spinning wheels and burning _their_ money!

This is quite critical to nail down right. I would argue that you need a very extensive test coverage here.

### Context Window Management

Here, the article touches on various ways you can "trim" the context, but still provide enough information to the LLM

- The **simplest** way is to **remove old messages in the conversation**. The "message trimming technique" is quite basic, but if the old messages are no longer relevant, there is no point in keeping them in the context window. The question, of course, becomes _how_ do you decide whether the message is relevant or not.

- You can **summarize existing messages**. The "compaction technique" **has a latency cost since you have to call another LLM to do the summarization**. Also, keep in mind that compacting messages will always be a lossy process – some nuance might not be preserved.

- Another solution is **RAG**, or the "selective reference loading" as the article calls it. Basically retrieve only the relevant parts of the thing LLM is asking for from the _reference memory_.

- Finally, and I believe this is often overlooked, **you can trim-down the tool call results**. Depending on the situation, this might be quite tricky – how do you ensure only relevant data is returned to the LLM or quite easy – does the LLM need to know all the metadata that's pulled from the database?

---

Such a great article. Even though there is clearly an incentive on this blog to route users to their "taskade" products, they mostly avoided it.

For me, the most critical takeaways are:

1. **Context engineering is, most likely, more important than the model itself**. The article claims that a "powerful" model with lackluster context will perform worse than a "less powerful" model with very good context.

2. **Scoping and guardrails are not optional**. I'm a huge believer in the fact that LLMs will get better and better. This belief makes me put a lot of faith into the LLMs and skip some of the "scoping" the article mentioned – only allowing the LLM to do specific things, instead of making it more of a "general" agent. I still believe that, in the future, LLMs will be able to handle "general" workflows, but as it is today, we need to make sure we build one agent that does one thing and does it really well.

## Using many agents does not necessarily mean getting better results

> Based on [this article](https://claude.com/resources/articles/building-agents-when-to-fold-subagents-into-your-main-agent-loop)

This is something I've noticed personally as well.

It goes like this:

1. You notice that performing some chunk of work with a single agent is "slow".

2. You attempt to parallelize that chunk of work by introducing multiple sub-agents.

3. You _feel_ like the work is done faster and better.

4. BUT when you actually look at the _metrics_ (which you should have, even for your own harness), you notice that **the coordination overhead is costing you lots of tokens and time**.

Using multiple agents is _very_ good practice for SOME tasks, but I found most of my projects end up overly complex. In such cases, you might want to _fold_ sub-agents back into a single agent.

### The goldmine of context

Let's say you have a video to edit. You decide to split the video into logical chunks and hand off editing to multiple sub-agents. You _think_ that the work will be done faster and with better quality.

What ends up happening is quite the opposite (in most cases at least). Instead of a nice, cohesive video, you are presented with disjointed parts that kind of work, but not really.

What happened here is that **each sub-agent is missing the context of the other sub-agents**. If you notice this happening, **the work you are trying to parallelize might not be a good candidate for parallelization**.

Sub-agents shine where each piece of work is completely isolated and perhaps independent of the others. Otherwise, you might be better off using a single agent. While the work might take a bit longer, you will NOT waste time trying to "merge" all those results.

### Split by context, not "slice of work"

Have you ever attempted to split a piece of work based on the "stack"? For example, you might be tempted to split a feature between a _backend_ and a _frontend_ agent.

During implementation, agents will most likely encounter something that is not explicitly within the specification. That might be true for both the _backend_ and the _frontend_ agents.

Since they are completely separate, and there is no sharing of context between them, the decision each agent makes is _isolated_ to that one agent. **This makes it really easy for them to create misalignment between _frontend_ and _backend_**.

**Instead, make it so a single agent handles a _single context_ of the work**. Different contexts are much more likely to be completely separate. This means that they are much more likely to not depend on each other that much.
