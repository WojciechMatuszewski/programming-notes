# Build AI Agents that Never Sleeps

Based on [this course](https://master.dev/workshops/background-agents/)

## Welcome to background agents

When you interact with agents synchronously, you are bound by the "prompt speed". Yes, the agent might work for hours on end, but it's you who has to start it. The work happens only when you explicitly say so.

Background agents flip this relationship. The agent is now autonomous and can react not only to your events (the input prompt), but also to things that happen in the "outside world", which is usually some kind of system.

## Give the agent a goal

In this section, Scott talked about how _giving the agent a view into the current state of the "world"_ is a requirement for building a good background agent. The agent has to know what is going on so it can make proper decisions based on that state.

Since we control the harness, we can inject the "current state of the world" into the prompt. Remember: the prompt might not come from us typing it. It might be composed based on an event that triggered a certain code path.

I've noticed that Scott, in pretty much every agent-related workshop they lead, uses some kind of _durable execution_ layer. In this case, it's [Inngest](ingest). It's quite interesting to me, because I have not had a chance (or felt the need) to use this technology at my work. Perhaps I'm missing something?

Stopped on part 2, start on the challenges and adding new stuff.
