# AI Hero DeepSearch course

> Based on [this course](https://www.aihero.dev/cohorts/build-deepsearch-in-typescript).

## Learnings

- It is very interesting to see the teacher using an LLM to solve the exercises.

  - Sadly, I'm unsure if this is a good way to learn. When learning, YOU are supposed to do the exercises, otherwise you will not remember what you learned.

- **TIL** about [serper](https://serper.dev). It's a quite nice API for building prototypes related to searching data.

- The `ai` package does a lot for us in terms of message formatting.

  - The `UIMessage` type and the `parts` property are pretty awesome to use. They allow me to interleave the UI for tool calls with the messages returned by the LLM.

- **TIL** about _search grounding_.

  - It's a technique to "ground" the LLM's responses with real, current data. It's most likely implemented via tools.

  - **Some models, like the ones from Google, have this built-in**.

    - This makes sense given that they come from Google.

  - It is somewhat similar to RAG, but the difference is that you are using the _latest data from the internet_.
    - RAG is usually scoped to a small domain and uses pre-built documents.

- _Drizzle_ is very powerful and has great DX.

  - Having said that, I still sometimes miss the flexibility of DynamoDB.
    - I wanted to create a counter that fails the increment if the `amount` is greater than some number (should be an `upsert` operation). I was unable to do all of this in one operation.

- The `ai` SDK can "re-create" all the messages using only `parts`.

  - I understand why, but it's a bit awkward to have to specify an empty content when passing initial messages containing only the parts.

- In the workshop, we use the `use-stick-to-bottom` package to ensure the scroll stays at the bottom when new messages from the LLM come in.

  - We could have used the `flex-direction: row-reverse` "trick" alongside the [`reading-flow`](https://developer.chrome.com/blog/reading-flow), but we would still need to write some logic to handle "scroll pinning".

  - The library, in my opinion, does not handle SSR that well. Regardless of whether you choose to use animations or not, there is a delay between the logic firing and the content appearing on the screen.

- **`drizzle-kit` does not have a built-in way to reset the database**.

  - This is weird if you ask me.

- **For tracing and observability, we used `langfuse`**.

  - I like the UI. In my opinion, it is much better than the one in braintrust.

- The **`ai` package and the `useChat` hook do not seem to handle switching threads mid-streaming very well**.

  1. Type something in one thread.
  2. Immediately switch to another thread.
  3. The messages from the first thread start to appear in the second thread.

  To fix it, one would most likely need to [implement resuming a stream](https://ai-sdk.dev/docs/ai-sdk-ui/chatbot-message-persistence#resuming-ongoing-streams).

- While attempting to re-implement the `use-stick-to-bottom` hook from [this package](https://github.com/stackblitz-labs/use-stick-to-bottom), I learned a neat pattern for managing client-side only constructs that work on DOM nodes.

Consider the following code:

```tsx
const [observer] = useState(() => {
  return new IntersectionObserver();
});
```

The goal here is to create an intersection observer once, and then use it with refs.

The **problem here is that the `useState` also runs on the server, and the `IntersectionObserver` is not a thing in Node.js**.

So, you might reach for `useEffect` to fix this issue.

```tsx
useEffect(() => {
  const observer = new IntersectionObserver();
  // More code
}, []);
```

The **problem here is that we might be re-creating the observer multiple times** depending on what you put into the dependency array. This might lead to race condition where something happened, but the observer was not yet attached.

The answer to all those issues is the _callback ref_ pattern.

```tsx
const lastObserver = useRef();

const elementRef = useCallback((node) => {
  lastObserver.current.disconnect();
  if (!node) {
    return;
  }

  lastObserver.current = new IntersectionObserver();
  // code
}, []);

<div ref={elementRef} />;
```

This way, we create the observer only when necessary, and, since the callback is invoked only on the client, we are safe to use any DOM-specific APIs in that callback.

The library also implements a handy helper called [`useRefCallback`](https://github.com/stackblitz-labs/use-stick-to-bottom/blob/3262b1110d600d6d8baac676dc2822d1d6dcb6b9/src/useStickToBottom.ts#L607) which allows you to use the variable that holds the callback as a regular ref!

- Be mindful of the function signature when using `fn.name`, especially in middleware-like functions.

  - If you are using anonymous callbacks, like `map(() => {})` the `fn.name` will be `undefined`!

- While we are not using that library in the course, I've taken a look into [`autoevals`](https://github.com/braintrustdata/autoevals) package.

  - It seems like most evaluators are wrappers around calling an LLM with specific system prompt.
  - There are some other evaluators that depend on doing math on vectors.

- In the course, we created a "global rate limiter" that is responsible for ensuring we do not call our "/chat" endpoint too often.

  - The implementation _waits_ till the next window if the requests in the current window are exhausted. I'm unsure if I like this pattern.
    - My main problem is that we _wait_ without doing anything in the background – we pay for idle! IMO it would be much better for us to handle the `429` on the client and wait there.

- **When writing _factuality_ eval**, we asked the LLM to grade the output by **selecting one of the options that we then mapped to scores**.

  - Supposedly, **this is because LLMs have trouble picking numerical scores**.
  - It also gives us an ability to calibrate scoring. When the LLM choose B, or C, we can grade it ourselves, and we know how to, since the option clearly describes the "quality" of the answer.

- One of the evals we wrote is a _relevancy_ eval. I find its inner-workings quite fascinating. [Taken from here](https://github.com/mastra-ai/mastra/blob/2d81790b9a9ec25b952ad3556ed5a92d03248751/packages/evals/src/metrics/llm/answer-relevancy/prompts.ts#L1).

  - First, we ask an LLM to split the answer to _statements_.
  - Then, we ask the LLM to grade the statements as they relate to the input (so the user prompt/question).
  - Then we average out the scores. Again, we **do not** ask the LLM to provide numerical scores, but rather `yes`, `no` and `unsure` answers we map to scores.

- Interestingly, [Anthropic's prompt best practices guide](https://docs.anthropic.com/en/docs/build-with-claude/prompt-engineering/use-xml-tags) mentions using XML tags.

- In one of the exercises, **we refactored the built-in agent loop of `maxSteps:10` into separate steps**. This was done to **improve the accuracy of results and tool choices**.

  - The **main problem** is that we are **increasing latency**. Instead of making one LLM call, we are making multiple LLM calls.

- There is the `experimental_transform` setting on the `streamText` which you can use to control chunking and how long to wait for a chunk to stream.

  - See [documentation](https://ai-sdk.dev/docs/reference/ai-sdk-core/smooth-stream) to learn more.

- We used another LLM to generate the chat title.

  - The main thing here is to ensure you are not blocking your main stream to do this. Make that request to the LLM "in the background" and await only in the `onFinish`.
    - As an improvement, you can _stream_ back the result via `dataStream.writeData` if the promise resolves _before_ `onFinish` finishes.

- In the workshop, we talked about _agents_ vs. _workflows_ and how "agentic" workflows are much less deterministic.

  - I like the metaphor of "agentic" dial that you can use to _steer_ your system to be more or less "agentic".

- The "agentic loop" evolved a lot throughout the workshop.

  - First, it was a single LLM call that also invoked tools.

    - This worked, but it was really not deterministic.

  - Then, we split the LLM call into two – one for gathering information and one for answering the question.

  - Then, we split the loop even more using three LLM calls.
    - One for answering the question, one for deciding what next action should be, and one for "evaluating" the whole workflow.
      - This makes our system quite deterministic, and **enables us to write evals for each separate "part" of that system**.

  **The more LLMs calls we added, the slower our workflow became**. This is quite problematic, and I do not have a good answer to this yet. It's not about the _latency_ but rather the _feeling_ that the workflow is slow. Perhaps adding some visual feedback would help?

- In the very end, we've added two more LLM calls:
  - One for assessing whether the query is "safe". Think _guardrails_.
  - One for checking if LLM needs clarification to proceed.
    - This one is very good for UX because it feels like LLM is "trying" to help you by asking follow-up questions.
