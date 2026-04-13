# FEM AI Engineering Fundamentals

- [Link to the repository](https://github.com/Hendrixer/ai-engineering-fundamentals).

- [Link to the workshop](https://frontendmasters.com/workshops/ai-engineering/).

## What is an AI Engineer?

I really like the definition Scott presents in the workshop.

> An AI Engineer is a system builder. You take foundation models and **turn them into dependable product features**.

It is _very_ easy to ship a POC nowadays, but does it stand against the scrutiny of real-world production usage?

## Your First Cloudflare Agent

A couple of things related to Agents on Cloudflare:

1. Integrated with _Durable Objects_, so the state is handled for you.

2. Supports WebSocket and SSE. We opted to use WebSockets instead of SSE. Reasons for doing that might include:

   - The CDN proxy might disconnect HTTP connections after X seconds. That does not apply to WebSockets, but might be problematic for SSE.

   - Cloudflare has a "hibernation" feature that allows it to evict _Durable Objects_ from memory whilst maintaining the WebSocket connection. [You can read more about this feature here](https://thomasgauvin.com/writing/how-cloudflare-durable-objects-websocket-hibernation-works/).

- **When writing tool descriptions (or `SKILL.md` descriptions)** you might want to consider structuring them in the following way:

  `<what it does> <when to use it> <when NOT to use it> <what it returns (not applicable to SKILL.md)> <parameter details (not applicable to SKILL.md)>`

  - The `<when NOT to use it>` does not literally mean listing all cases when it does not make sense to use a given tool. **It's about giving positive and negative _signals_ to the LLM**. For example: "Prefer this tool over X when you have to do Y".

  - Different providers have different guides for writing effective tool descriptions. Before you run something on production, make sure that you understand what those are!

  - **At the end of the day, you should evaluate those tool descriptions to ensure the LLM calls the tool it should call**.

- Scott mentioned that it is vital to start with something very simple first. Be "naive", do not be afraid of simplicity.

  - I second this approach so much. Yes, if we are aware of some best practices, we can front-load those, but it is usually better to build something simple first and see how it works in a "real world".

  - Also, consider people you are working with. If you build something really complex, what are the chances they will onboard quickly to the project at the start? **IMO the complexity should raise gradually**.

TODO: Feelings about Cloudflare. Start part 3
TODO: different techniques Scott mentioned (prompt engineering guide website)

## The Chat Experience

Start part 3
