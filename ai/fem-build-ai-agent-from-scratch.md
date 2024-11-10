# Build an AI Agent from Scratch

> Notes from [this workshop](https://frontendmasters.com/workshops/build-ai-agent/).

## LLM Primer

> [Course notes](https://clumsy-humor-894.notion.site/1-LLM-Primer-13554fed51a380c391c3fbeaab456711).

- Currently, LLMs are pretty good at predicting what the _next_ word in a sentence could be.

  - It takes into the account _everything_ you typed so far up to that point.

  - Context is the king.

- The T in `ChatGPT` is for _transformers_.

  - _Transformers_ allow the model to keep track of the context.

- **A _transformer_ consist of multiple layers of logic**.

  - It is not a "single thing".

- The more parameters model has, the bigger it is. The bigger the model, the more accurate it is with the answers.

  - The **parameters determine how input data is transformed**.

    - They play critical role in the _transformers_ architecture as they dictate what weights are assigned to words.

- **An _AI Agent_ vs. _AI Assistant_**.

  - **Agent** can run a "loop of thought". It can interact with various tools and asses whether the output is valid or not.

  - **Assistant** is a dumbed-down version of _agent_. **Performs tasks as directed**. It does not use any "feedback" mechanism.

- LangChain used to be very valuable, especially for Python developers.

  - Nowadays, the ChatGPT SDK overlaps with LangChain a lot.

## Hello Chat

- By default, if you use the `chat.completions` OpenAI endpoint, the response will be _transactional_.

  - This means there is **no memory** or recollection of previous messages.

- The main trick with chat-based AI apps is to **effectively manage tokens and the context**.

  - Sending the whole conversation alongside a new message works in theory, but in practice, it will cost you a lot to keep this kind of architecture running.

    - In addition to cost, you must consider token limits.

## Memory

- There are various _types_ (the `role` property) of messages you can send to ChatGPT client.

  - The `system` is for setting the AI behavior.

    - Consider **making the AI write prompt for you**. This is a bit meta, but it works!

    - You might want to change this prompt dynamically. For example, injecting the current time or any other information that changes might help you get better response.

  - The `user` is for user input.

  - The `assistant` is for AI responses.

  - The `tool` is for responses from tools.

- To implement database for messages Scott opted for `lowdb/node` package.

  - Pretty interesting package!

- **When building a chat LLM, having a good message summaries it a key to a great UX**.

  - If this piece is lacking, your users will not get good answers.

## What is an Agent

- An LLM that has the capability to run various _tools_ and keeps the memory of previous inputs.

  - These **allow the agent to make decisions** based on the "current state of the world".

- The OpenAI SDK integrates with `zod` to ensure the response from a _tool_ matches the schema. [You can read more here](https://platform.openai.com/docs/guides/structured-outputs).

  - This makes using the tool easier as you no longer have to validate responses yourself.

- **Scott showcases a neat technique to force the AI to come up with a good reason to call a given tool**.

  ```js
  const weatherTool = {
    name: "get_weather",
    description: "Gets the weather",
    parameters: z.object({
      // The `reason` parameter is the key!
      reason: z.string().describe("Why did you pick this tool?"),
    }),
  };
  ```

  Now, you can log the `reason` parameter value.

  I've played around with this approach and it changes the output of the LLM.

  Before, the query `What is the weather` would prompt the LLM to call this function.

  After, the same query, would make the LLM to ask about the location.

Start part 4 next -> https://frontendmasters.com/workshops/build-ai-agent/
