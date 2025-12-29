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

- TIL that **you can generate embeddings with the `ai` package**.

  - I mean, it makes sense, but I was unaware of this functionality. In the past, I've used `langchain` to do this, which means adding unnecessary dependency to the project.

    - The `ai` package has `embed` and `embed` many functions that you can use.

- **You can't forget about _chunking_ when using embeddings**.

  - Chunking means splitting the content you want to generate embeddings for. This is beneficial, because now you can search through smaller embedding vectors which increases _retrieval accuracy_.

    - In addition, you will end up passing smaller chunks back to the LLM. This means less context bloat.

  - There are many ways you can chunk the content.

    - **Fixed-size chunking** is when you set a static number for the the chunk size and the overlap. This is the easiest but also the most "naive" way to do this.

      ```ts
      const splitter = new TextSplitter({
        chunkSize: 1000,
        chunkOverlap: 100,
      });
      ```

    - **Recursive chunking** is when you take the fixed-size chunks, but you apply the splitting recursively on various separators. [I found this explanation](https://dev.to/eteimz/understanding-langchains-recursivecharactertextsplitter-2846) to be very informative.

      ````ts
      const splitter = new RecursiveCharacterTextSplitter({
        chunkSize: 1000,
        chunkOverlap: 100,
        /**
         * The separators are applied in order they appear in the array.
         * Apply -> check if chunk is less than `chunkSize`.
         *  If not -> apply next.
         *  If yes -> check if we can merge with others.
         * Repeat.
         */
        separators: ["\n## ", "\n### ", "\n#### ", "\n#### ", "\n##### ", "```\n\n", `\n--- CHAPTER ---\n`, "\n\n"],
      });
      ````

- **Another thing you can try is to re-rank your chunks**.

  - First, you generate the chunks. Consider using the recursive text splitter and playing around with the `chunkSize` and `chunkOverlap` parameters.

  - Then, you pass the topK chunks to a re-ranker LLM which decides which chunks are _really_ relevant given the context. The context might be the conversation history, or something else. It's up to you.

    - Doing this **has a chance of greatly increasing the quality of the chunks you pass to the "main" LLM while "taking the pressure off" your retrieval algorithm**.

    - Of course, the main drawback here is that you **introduce additional latency, and potential point of failure, to the workflow**. Calling _yet another LLM_ is not free either.

    - There are external external services you could use here as well. [Example](https://cohere.com/rerank).
