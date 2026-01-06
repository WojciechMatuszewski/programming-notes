# AI Hero – Personal Assistant

Learnings from [this course](https://www.aihero.dev/cohorts/build-your-own-ai-personal-assistant-in-typescript)

- [Project repository](https://github.com/ai-hero-dev/cohort-002-project)
- [Skills repository](https://github.com/ai-hero-dev/cohort-002-skill-building)

## Learnings

- BM25 is a _keyword search_ algorithm that is already quite battle-tested.

  - It is quite simple, but **effective for exact keyword matches**.

  - Note that it does not handle the _semantics_ of words.

- Sometimes, you _do not need_ to extract the user prompt from the messages to call the LLM.

  There was an exercise for asking the LLM to generate keywords for BM25 search based on the user prompt. Extracting the user prompt from the `Array<UIMessage>` (or `Array<ModelMessage>`) is quite a hassle.

  Instead, you can send the messages to the LLM (provided the array of messages is not that large).

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

  - **You can combine two searching algorithms via the Reciprocal Rank Fusion algorithm** to create a "hybrid search" algorithm.

    It is actually a pretty straightforward algorithm. You can express it as a sum of all `1/(k + rank(x))` over the results. [Here is a good video on this topic](https://www.youtube.com/watch?v=2uBcjEecr38).

    The `k` is a constant, usually `60`, and it controls how much we _favor_ the lower-ranking results. The bigger the `k`, the less we favor the lower rankings.

- Matt argues that whenever you are building some kind of semantic search functionality, you most likely need to add a _query rewriting_ step before passing the query to the semantic search algorithm.

  - I concur. We can't expect users to formulate _good_ search queries. They might be vague, incorrect, and full of grammar mistakes. Ideally, our system would handle those gracefully.

- TIL that **you can generate embeddings with the `ai` package**.

  - I mean, it makes sense, but I was unaware of this functionality. In the past, I've used `langchain` to do this, which means adding an unnecessary dependency to the project.

    - The `ai` package has `embed` and `embedMany` functions that you can use.

- **You can't forget about _chunking_ when using embeddings**.

  - Chunking means splitting the content you want to generate embeddings for. This is beneficial because now you can search through smaller embedding vectors, which increases _retrieval accuracy_.

    - In addition, you will end up passing smaller chunks back to the LLM. This means less context bloat.

  - There are many ways you can chunk the content.

    - **Fixed-size chunking** is when you set a static number for the chunk size and the overlap. This is the easiest but also the most "naive" way to do this.

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
         * The separators are applied in the order they appear in the array.
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

    - Of course, the main drawback here is that you **introduce additional latency, and a potential point of failure, to the workflow**. Calling _yet another LLM_ is not free either.

    - There are external services you could use here as well. [Example](https://cohere.com/rerank).

- Keep in mind that you can create a _function that returns a tool_.

  - This is quite powerful. It allows you to create _dynamic_ tools.

    We used this technique when working on the re-ranker. In addition to passing the search query and chunks, we also included message history (only `user` and `assistant` messages) in the LLM context.

    Doing this allows the LLM to pick the relevant chunks better.

    ```ts
    // Instead of this
    const searchTool = tool({});

    // You can do this
    const createSearchTool = (props) => {
      return tool({
        // use props here.
      });
    };
    ```

- I really like how the `ai-sdk/react` package works with the backend.

  - You can infer the types of tools the LLM has access to.

    ```ts
    export type MyMessage = UIMessage<
      never,
      {
        "frontend-action": "refresh-sidebar";
      },
      InferUITools<ReturnType<typeof getTools>>
    >;

    const getTools = (messages: UIMessage[]) => ({
      search: searchTool(messages),
      filterEmails: filterEmailsTool(),
    });
    ```

  - Then you can use those types on the frontend when handling messages.

    ```tsx
    messages.map((message) => {
      return message.parts.map((part) => {
        switch (part.type) {
          // I believe the `tool-${}` prefix is automatically added by the package.
          case "tool-search": {
          }
        }
      });
    });
    ```

- In the course, we've changed our filter and search emails tools to return a _truncated_ email body and we've added another "retrieve" tool to get the full email.

  - I'm a bit conflicted about this implementation. On one hand, it makes sense to me—emails can be long, and we don't want to include all that text in the context if it's not relevant.

    On the other hand, wouldn't this cause many more `getEmail` tool calls, which could bloat the context even more?

    Matt called this the **metadata-first retrieval pattern**, but all I could find online regarding this involves appending LLM-generated metadata to chunks before embedding them. Perhaps we'll do that in future lessons?

    I remember learning about the _progressive disclosure_ pattern as [described in this blog post](https://www.anthropic.com/engineering/code-execution-with-mcp). This feels similar, but it's actually quite different. Why is that?

    Also, I'm unsure about the example we used in the course to test this. The query was: "Give me 10 email bodies." This resulted in two LLM calls: first to filter the emails, then to fetch data for those 10 emails. Wouldn't we only need one tool call (the `filterEmails`) if we didn't truncate their bodies?

    As I explore this implementation further, I'm starting to understand its significance. **We can instruct the LLM to skip fetching the whole email if the metadata already contains the information it needs to answer the user query.** In addition, the LLM can fetch the full content of only _some_ emails returned by our search tools. All of this will save us tokens in the long run.

- When doing the "memory system" exercises, we looked at two approaches: (1) always call the LLM in the `onFinish` callback to _force_ the LLM to consider whether it needs to manage memories, and (2) expose the memory management as a tool for the main agent.

  Consider the pros and cons of both approaches.

  For the first approach (the manual, forced call), we get a lot of determinism, and we **ensure that the system prompt of the main agent has no influence over the memory management concern**. However, we also increase system latency and costs, since you will _always_ call the LLM.

  For the second approach, we lose determinism and have to "pollute" the main agent system prompt with heuristics for managing memories, but we make the system cost less and reduce latency.

  **This is a typical situation where you have to consider whether to give the LLM more control or not**. Some would argue that LLMs will get better and better, so betting on giving the LLM more control is the way to go. I tentatively subscribe to that view.

- Any memory system will grow as the user interacts more with the LLM. **To make the memories useful at any size, we can deploy a similar RAG approach as before**.

  1. When creating and updating memories, make sure to generate embeddings for each memory.

  2. Upon receiving a message from the user, **use an LLM to semantically search through the memories**. You can use the same technique we did before: generate keywords and the query, use both bm25 and semantic search, and combine the results.

  3. Push only the most relevant memories to the "main" LLM context.

  This solution is quite nice, but it's not ideal. You have more "knobs" to turn in your system (like the topK value for memories), which means it is more complex to maintain.

- When adding memories to the _project_, we decided to convert the _entire_ message history to text in order to semantically search against all memories.

  - I'm unsure about this approach. The larger the conversation, the more noise there is in that "conversation text" embedding.

- TIL about various `ChatTransport` options you can specify when using the `useChat` hook. [You can read the documentation here](https://ai-sdk.dev/docs/ai-sdk-ui/transport).

  - We overwrote the `DefaultChatTransport` to **send only one message to the backend**. Why? We are already persisting all the messages, so there is no need to send everything again.

  - All this work nicely aligned with our task of only providing the LLM with N recent messages and using semantic search (static, not LLM-driven) to extract relevant older messages.

    ```ts
    const messages = validatedMessagesResult.data;

    const recentMessages = messages.slice(-MESSAGE_HISTORY_LENGTH);

    const olderMessages = messages.slice(0, messages.length - recentMessages.length);

    const allRelevantOlderMessages = await searchMessages({
      recentMessages,
      olderMessages,
    });

    const relevantOlderMessages = allRelevantOlderMessages.slice(0, OLD_MESSAGES_TO_USE);
    ```

    The `searchMessages` function creates a single blob of text out of `recentMessages` and, via cosine similarity, checks if there are any other messages that might be relevant.

    Again, I'm unsure how I feel about this approach. On one hand, it's quite easy to do. On the other, that one blob of text from the `recentMessages` might be quite large—it contains a lot of "noise".

    I've asked Claude about possible alternatives, and they either include (1) using the LLM to rewrite the query (which we currently create by extracting text from `recentMessages`), or (2) creating embeddings for each of the `recentMessages`.

- In addition to feeding the LLM `recentMessages` and the `memories`, we also implemented adding the _most relevant_ chats.

  Our implementation was quite _naive_, because alongside LLM-generated metadata for the chat—things like `tags` or `whatToAvoid`—we also provided all the messages. This implementation could bloat our context window quite a lot.

  Having said that, if you optimize this, perhaps by chunking or being more selective, you can achieve pretty neat results. Consider what we have in place:

  1. The **working memory**, which is the **current conversation (the most recent and relevant older messages)**.
  2. The **episodic memory**, which is historical experience and its takeaways—**this is our _most relevant chats_ and the `whatToAvoid` or `whatWentWell` properties**.
  3. The **semantic memory**, which is the **memory system and the RAG implementation for the emails**.
  4. The **"rules" and "skills" for interaction**, which are tools and the system prompt.

  All of this implements the so-called **_cognitive architecture for language agents_** [which you can read more about here](https://github.com/ALucek/agentic-memory?tab=readme-ov-file).

- **Apart from ensuring your system is working as expected, evaluations are a very good way to run A/B tests for various models**.

  - If you have a robust suite of evaluations, you can easily run them against different models to compare results.

    - **Models get better and cheaper with time. You might not need to use the most powerful model for your use-case**.

- Seeing the results of evaluations for different models, especially when comparing smaller vs. larger models, is quite fascinating.

  - For example, when asking for weather information `gemini-2.5-flash-lite` assumed the units, but the `gemini-2.5-pro`, instead of calling a tool, asked what units it should use.

- **When writing evals, we tend to fixate on the _happy path_, but adding "adversarial inputs" into your test harness is just as valuable**.

  - The real world is messy. We can't expect the users to provide the LLM with coherent prompts all the time. **You must make sure that your system gracefully handles inputs that do not make sense**.

    - Again, this is quite powerful when using multiple models. You can see which model is good at handling the sad and happy paths.

- From my empirical experience, I've noticed that the **system prompt carries much more "weight" than any description you can attach at the tool level**.

- When working on evaluations for memory extraction, we first implemented the LLM-as-a-Judge, and then replaced it with cosine similarity-based matcher.

  - Both approaches has their tradeoffs. For the LLM-as-a-Judge, you get precision, but that can backfire. In addition, you have to _align_ your judge with human labels.

  The the cosine similarity approach is easier to work with, but is definitely less precise. For example, "Working at Google" and "Working at Facebook" might produce quite high scores, but they are very different in terms of facts.

- The way you structure your code has implications on how _easy_ it will be for you to write evaluations.

  - You most likely **want to isolate the `streamText` and `generateObject` calls in separate functions, so you can export them later on**.

    You might also want to **consider using [the `Agent` abstraction](https://ai-sdk.dev/docs/agents/overview) to encapsulate your agents**.

- Implementing HITL (human in the loop) was easier than I though.

  - The `ai` package exposes all the necessary abstractions to do that, and in the `v6` it will be even easier!

- TIL that `@ai-sdk/mcp` package exists. It allows you to programmatically adds tools from other MCP servers to your agent.

  - This is quite powerful, as it **allows you to "pre-process" the tools before you add them to `tools` property to your agent**.

    Let us say you have an MCP server which you really want to use, but you only need one tool it exposes. Nothing stops you from removing all the other tool definitions that `@ai-sdk/mcp` returned, and passing only that single one to the `tools` property of your agent.

    ```ts
    import { experimental_createMCPClient } from "@ai-sdk/mcp";
    import type { ToolSet } from "ai";

    export const getMCPTools = async (): Promise<ToolSet> => {
      const client = await experimental_createMCPClient({
        transport: {
          type: "http",
          url: process.env.MCP_URL!,
        },
      });

      // Check out the parameters of the `tools` method.
      // You can pass ZOD schema here for type-safety.
      const tools = await client.tools();

      console.log(Object.keys(tools));

      // You can do that for other tools.
      if ("add_tools" in tools) {
        delete tools.add_tools;
      }

      if ("edit_tools" in tools) {
        delete tools.edit_tools;
      }

      console.log(Object.keys(tools));

      return tools as ToolSet;
    };
    ```

- TIL that you can **send arbitrary data in the `parts` array, as long as you provide mapping for that data on the backend**.

  - The `convertToModelMessages` allows you to map parts, which makes all of this much easier.

  - We leveraged "custom" `parts` to implement HITL for external MCP tools. You can propagate which tools user has enabled on the frontend, and then wrap their calls with HITL code on the backend. Pretty powerful!

## Wrapping up

This course was great! I really liked the split between the _project_ and _skills_.
