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

2. Supports WebSocket and SEE. We opted to use WebSockets instead of SSE. Reasons for doing that might include:

   - The CDN proxy might disconnect HTTP connections after X seconds. That does not apply to WebSockets.

   - Cloudflare has a "hibernation" feature that allows it to evict _Durable Objects_ from memory whilst maintaining the WebSocket connection. [You can read more about this feature here](https://thomasgauvin.com/writing/how-cloudflare-durable-objects-websocket-hibernation-works/).

Finished part 2 1:43
