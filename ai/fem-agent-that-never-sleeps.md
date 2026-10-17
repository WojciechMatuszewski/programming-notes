# Build AI Agents that Never Sleeps

Based on [this course](https://master.dev/workshops/background-agents/)

## Welcome to background agents

When you interact with agents synchronously, you are bound by the "prompt speed". Yes, the agent might work for hours on end, but it's you who has to start it. The work happens only when you explicitly say so.

Background agents flip this relationship. The agent is now autonomous and can react not only to your events (the input prompt), but also to things that happen in the "outside world", which is usually some kind of system.

## Give the agent a goal

In this section, Scott talked about how _giving the agent a view into the current state of the "world"_ is a requirement for building a good background agent. The agent has to know what is going on so it can make proper decisions based on that state.

Since we control the harness, we can inject the "current state of the world" into the prompt. Remember: the prompt might not come from us typing it. It might be composed based on an event that triggered a certain code path.

I've noticed that Scott, in pretty much every agent-related workshop they lead, uses some kind of _durable execution_ layer. In this case, it's [Inngest](ingest). It's quite interesting to me, because I have not had a chance (or felt the need) to use this technology at my work. Perhaps I'm missing something?

In the app, we have one spot where we ask the LLM to derive the next action we ought to take based on the current state. Scott implemented this flow by calling a regular text model. The nice thing with text models, as opposed to models like Jev, is that they can _elaborate_ on their choice. You can force the model to output the `reason` field **which you can't really get from a classifier model, like Jev**.

## Make progress durable

That is what using "durable execution layer" gives us – we can _resume_ the agentic loop after bad things happen, like a server crash.

But this level of robustness does not come without a price. **Every "step" in the workflow needs to be idempotent** and **the workflow itself needs to be deterministic**.

- **Deterministic** means that the same input always produces the same output.

  - The `2+3` is deterministic. The `Date.now()` is not.

- **Idempotent** means that running a given function once produces the same results as running it multiple times.

  - Those function calls _can_ have side-effects, but the side-effect must not change the state of the world after it runs for the first time.

In addition to what we talked about so far, you will also need to reason a bit more about the code. The question "should this be it's own step" will be asked often.

A good heuristic seem to be put every `async` operation within the step. You _might_ want to put multiple `async` operations within the step, but only if you are okay with _both_ of those operations running together when step retries.

LLM calls are also a good example of something that you most likely want _inside_ the `step` function. LLM calls are not deterministic, even with a strict schema.

Start "Human in the Loop"
