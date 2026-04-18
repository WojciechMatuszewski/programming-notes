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

## The Chat Experience

- OpenAI API response shapes seems to be a standard. Most providers have OpenAI-compatible API.

- On the frontend, the `useAIChat` and similar hooks are doing a lot of work for you, but you still have to handle various states of tools and `parts`.

  - If you are not careful, you will litter your FE code with multiple ternaries that are hard to read.

## The Eval Discipline

- The **pass@k** is the **probability that the system produces a correct answer in at least one of k attempts**.

  - Used in "best of k" workflows. Basically you run "the thing" multiple times and see how many of those runs were "successful".

- The **pass^k** is the **probability that all k attempts are correct**.

  - Consider having 20 test cases. How many times in the row can you run all those test cases and have **all of them** succeed?

- In this lesson, we used the 1-5 judging scale. **I'm not a fan**, but at least we have a _scoring rubric_ we can fall back onto.

  - [Read this to learn more about why having Likert scales is not helpful](https://hamel.dev/blog/posts/evals-faq/#q-why-do-you-recommend-binary-passfail-evaluations-instead-of-1-5-ratings-likert-scales).

    - One thing to note: **using _binary_ scorers make it harder to see "the agent is improving even if it's not perfect yet" signal**.

- **Node.js allows you to programmatically (not via CLI flag) load env vars**.

  - [See this documentation page](https://nodejs.org/api/environment_variables.html#dotenv). You can use the `process.loadEnvFile` or `util.parseEnv` functions!

- Scott recommends using a framework to present evaluation results.

- I'm unsure if jumping to LLM-As-Judge implementation so early is a good idea (will be done in part 5?).

## Automated Scorers

- In this lesson, we changed our `golden.json` data structure.

  - To properly accommodate for complex flows, we have to "seed" the conversation with `user` and `assistant` message types.

- Given that we want to re-use the agent we have in Cloudflare worker and the eval suite, I wonder why we did not use the [`Agents` abstraction](https://ai-sdk.dev/docs/agents/overview).

- We first started with **code-base scorers**. That's quite good. I like that Scott kept is as simple as possible.

- A couple of scorers we created:

  - The _schema scorer_ which checks if the agent produced a valid Excalidraw output. TBH, I'm unsure if this scorer is needed at all. We could offload that validation to the tool `outputSchema`.

    - But, if we did that, would we know that it failed? Would we track those failures over time?

- "The goal of the scorer is to surface what we could improve later on". I like this analogy, but I also believe scorers are good for making sure we do not regress in functionality.

- Adding "human review" scorers in Braintrust is quite powerful.

  - It allows you to include SME's (subject matter experts) into the process. **When SME's add comments, they are included in the dataset**. This allows you to create a very nice feedback loop.

  - When creating this scorer, **use pre-defined options instead of sliders or similar**. The less choice you give someone, the faster they will be able to grade the output.

## Context Engineering

- Do not try to make the agent good at everything. **Make sure it is really good in a given niche**.

- Scott decided to include a brief mention what each tool does in the system prompt. I'm a bit conflicted here.

  - On the one hand, it makes sense. We want the LLM to be fully aware of the tools.

  - On the other, this approach does not scale well. Do you put the same description for a given tool here as in the tool definition? Do we do something different? What about drift in the descriptions? Also, where there are multiple tools, especially those discovered "lazily" (via progressive disclosure), how would that work?

- In the new system prompt, we introduced a concept of a "negative prompts". For example

  ```
  DO NOT place two elements at the same coordinates
  ```

  **You can use "negative prompts" to steer the LLM away from it's "general" knowledge**.

  Consider other diagraming tools (relevant to what we are building). Most of them put text _inside_ the shapes by default. That is not the case for Excalidraw API.

  We can use "negative prompts" to help the model "notice" this difference.

  **Scott warns that we should not overdo it**. If you put too many "negative prompts" the LLM will start to ignore it.

  **"Negative prompts" work really well for common failure cases in your niche**.

- In our application, to improve agent quality, we decided to embed the _canvas state_ into the system prompt. **We did that by transforming `json` data into [`TOON`](https://toonformat.dev/guide/getting-started.html) format**.

  - I like how Scott put it: `json` is for machine-to-machine communication, `yaml` is for configuration and `TOON` (or any other similar format) is for LLM data injection.

  - **We want to transform `json` into any other more token-efficient format** that is also somewhat readable.

- There are at least two ways of sending data back to the agent:

  - via the `body` property you can define on the `useAIChatAgent` hook options.

    - You only get the latest snapshot.

    - Clear separation between additional data and LLM messages.

    - Send with every request. Might not work well for "situational data".

  - via a custom `data-x` message part. This requires wrapping the `sendMessage` function or (I have not tested it), defining a custom `transport`.

    - By default, does not propagate to the parts LLM sees (the `data-x` prefix is dropped by AI-SDK before it reaches the LLM).

    - Good for "situational data". Think sending a piece of additional data upon a certain user interaction.

    - You can transform this part to text so LLM sees this.

    - Gives you the ability to see a "history" of data. Each message captures a snapshot of how this data looked when we send that message.

  **Regardless of the approach you choose** consider how sending data from FE to BE works in the context of _tool continuation_.

  By _tool continuation_, I mean LLM stopping so you can execute a client tool (perhaps HITL workflow). If, before sending "continue" message to the LLM, the canvas state changed, the LLM might "see" a stale state.

  This boils down to how you send that "continue" message. In our case, we would use `addToolOutput` which does not trigger the callback you set in the `body`.

Start part 7
