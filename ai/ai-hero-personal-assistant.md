# AI Hero – Personal Assistant

Learnings from [this course](https://www.aihero.dev/cohorts/build-your-own-ai-personal-assistant-in-typescript)

- [Project repository](https://github.com/ai-hero-dev/cohort-002-project)
- [Skills repository](https://github.com/ai-hero-dev/cohort-002-skill-building)

## Learnings

- BM25 is a _keyword search_ algorithm that is already quite battle-tested.

  - It is quite simple, but **effective for exact keyword matches**.

  - Note that it does not handle _semantics_ of the word.

- Sometimes, you _do not need_ to extract the user prompt from the messages to call the LLM.

  There was an exercise for asking LLM to generate keywords for BM25 search based on the user prompt. Extracting the user prompt from the `Array<UIMessage>` (or `Array<ModelMessage>`) is quite a hassle.

  Instead, you can send the messages to the LLM (granted the array of messages is not that large).

  ```ts
  async function generateKeywords(messages: Array<UIMessage>) {
    const {
      object: { keywords },
    } = await generateObject({
      system: KEYWORD_GENERATOR_SYSTEM_PROMPT,
      messages: convertToModelMessages(messages),
      model: google("gemini-2.5-flash"),
      schema: z.object({
        keywords: z.array(z.string()).min(1).max(5),
      }),
    });

    return keywords;
  }

  // As an alternative, you could extract the user prompt like so:

  import { convertToModelMessages } from "ai";

  function extractUserPrompt(messages: Array<UIMessage>) {
    const modelMessages = convertToModelMessages(messages);

    const lastUserMessage = modelMessages.findLast((msg) => msg.role === "user");

    return typeof lastUserMessage?.content === "string" ? lastUserMessage.content : null;
  }
  ```

- BM25 is very good for _exact keyword_ matching, and using _embeddings_ is quite nice for semantic search. Can we combine those two approaches?

  - **You can combine two searching algorithms via Reciprocal Rank Fusion algorithm** to create "hybrid search" algorithm.

    It is actually a pretty straightforward algorithm. You can express it as a sum of all `1/(k + rank(x))` over the results. [Here is a good video on this topic](https://www.youtube.com/watch?v=2uBcjEecr38).

    The `k` is a constant, usually `60`, and it controls how much we _favor_ the lower-ranking results. The bigger the `k`, the less we favor the lower rankings.

- Matt argues that whenever you are building some kind of semantic search functionality, you most likely need to add _query rewriting_ functionality step before passing the query to the semantic search algorithm.

  - I concur. We can't expect users to formulate _good_ search queries. They might be vague, incorrect, and full of grammar mistakes. Ideally, our system would handle those gracefully.

Start lesson 20 - adding BM search
