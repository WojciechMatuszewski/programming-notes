# Build an AI Agent from Scratch v2

> [Based on this workshop](https://frontendmasters.com/workshops/ai-agents-v2/)

> [Course notes](https://publish.obsidian.md/agents-v2/course)

## 01 Intro to Agents

- There are multiple definitions for what _agent_ mean. IMO there is no point in discussing this that much.

  - Know that, an **Agent is an LLM that can take actions in a loop until the task is complete**.

- Agents are good for tasks that are not very deterministic. For example, writing code is non-deterministic, linting the code is.

- **Agents do not know what they do not know**.

  - Better models with ask questions. Smaller models with straight-up lie to you.

## 02 Tool Calling

- You can give the LLM _tools_. Those are functions LLM can tell you to call (and provide back results) when working on a task.

  - Without tools, the LLM would be constrained to the knowledge it was trained on, which is already outdated when it's released to the public.

- When learning about tools, you might come across something called _human-in-the-loop_ or HITL.

  - This technique is for introducing a human input into the agent loop. It adds a bit of determinism to the whole flow.

    - Imagine a `delete.file` tool. You most likely want to have some kind of confirmation dialog that user approves before proceeding.

- TIL about the so-called "Zero Data Retention" policy your admin can apply on the OpenAI organization.

  - In essence, every tool call has ID associated with it. OpenAI uses those IDs for abuse monitoring. When you have "Zero Data Retention" set, you will need to disable the "store" option for OpenAI provider.

    ```ts
    const { text } = await generateText({
      model: openai(MODEL_NAME),
      messages: allMessages,
      system: SYSTEM_PROMPT,
      tools,
      stopWhen: stepCountIs(5),
      providerOptions: {
        openai: {
          store: false, // This one.
        },
      },
    });
    ```

- In the workshop, we decided to build our _own_ tool-calling loop.

  - It's implemented in the `ai` package, but I still think it's worth knowing how to, albeit in a simple matter, to implement it.

    - I have to say, it took me more time to write this loop that I would like to admit 😅.

- When playing around with `ai` package and tools, I've noticed that one can pass the `required` option for the `toolChoice` prop.

  - According to the Claude, there is nothing magical about this. **The "toolChoice" `required` is implemented on the provider side, where the API call honours the "required" parameter**. The `ai` package does not implement any specific logic to make this functionality work.

## 03 Single Turn Evals

- LLMs are inherently non-deterministic (even if you set `temperature` to `0`).

- In most cases, you can't _test_ LLMs the same way you would test a piece of code.

  - You take the output of the LLM and you _measure_ it against your objectives.

- In the course, we talked about _offline_ and _online_ evaluations.

  - The _offline_ evaluations use fixed dataset. You run them before the deployment.

  - The _online_ evaluations use real user traffic data and are run in production.

- Scott mentioned a term **hill climbing**. This is a fancy way of saying that we want to gradually improve our agent.

  > You can read more about this approach [here](https://addyosmani.com/blog/ai-evals/).

  - You run the evals "as-is". Without changing anything in your code yet. This run is your baseline.

  - You make a change that _could_ influence the score.

  - You run the evaluations again.

  - If you improved the scores, you keep the change.

  - Repeat.

  **To me, this sounds like a very easy way to over-fit your test data**. If you iterate on it too much, you most likely will, by accident, start embedding test data into your prompt as "examples".

  One has to be mindful here.

- I like how Scott categorizes entries in a eval data sets.

  - We have a property called `category` which can have multiple values: `golden`, `negative`, `secondary` (or any other).

  Be wary of introducing too many labels. We have a tendency to "over categorize" sometimes. The more labels you have, the less clear the taxonomy of the data set is.

- No matter what evaluation framework you choose, some things are "constant". You can bet there is some kind of:

  - _Scorer_ which is responsible for taking the LLM output and grading it.

  - _Executor_ which is responsible for calling the LLM.

  - _Dataset_ which is a collection of "input" and "expected output".

  - _Experiment_ which invokes the _executor_ and, given then output of the _scorer_ gives you nice visual charts.

## 04 The Agent Loop

- This section was all about actually _writing_ the Agent loop.

- Technically, you get the loop for free when using the `ai` SDK, but it is still worth having some understanding how it works.

  - Nowadays, the real value is actually _making a good use_ of the loop to provide value.

## 05 Multi-turn Evals

- In the previous lessons, we focused on a single input-output scenario. Now, it's time to focus on the conversation as a whole.

- **Since the "output" is an array of messages, which is completely non-deterministic, we have to use either another human or LLM to judge the output**.

  - Writing a good LLM judge is _hard_. You most likely want to involve a _subject matter expert_ (SME) in the process.

- Scott touches on some important topics:

  - Constraining the output of the LLM judge. Free-form answers are introducing too much ambiguity in the process.

    - Consider forcing the LLM to output the _reason_ for the score. I've noticed that the LLMs score is more accurate that way (plus you have some information on the score).

  - Using a "stronger" model than the one you are using for running the evals.

  - You will most likely need to mock tool results to "complete" the conversation.

- **What I do not like** is the **reliance on 1-10 scoring system**.

  - I strongly believe it should be either pass (1) or fail (0). What's the difference between 7 and 8?

- You should instrument the LLM judge calls just like you instrument your "main" LLM agent call.

  - It's really helpful to see the calls of the judge as well!

## 06 File System Tools

- Tools are _similar_ to regular functions, but **errors and what you return to the agent is VERY important**.

  - In regular applications, you might display "Oops, something went wrong" to the user and call it a day. That _should not_ be the case for tools.

    - What you return should help the agent to fulfil the request. This also applies to any errors!

- The more granular the tools are, the more steps the agent has to take. **The more steps agents has to take, the bigger probability of an error**.

  - There is a **delicate balance between providing too many or too few tools**.

## 07 Web Search Context Management

- There are at least two ways you can solve this:

  1. Use the `webSearch` tool provided by services like Gemini or OpenAI.
  2. Use a third-party service, like `exa`, to search the web.

  Both of these approaches work as _tools_. The former provides you with less flexibility but is easier to work with, while the latter gives you a lot of flexibility but requires integration with a specific provider.

- In this section, Scott places a lot of emphasis on _context management_ and how having too many tokens in context negatively affects LLM performance.

  - The "search web" results can add a LOT of tokens to the context.

- Here are some **strategies for managing context "overflow"**:

  - _Compact_ the conversation (extract salient facts), but this is not a silver bullet. Any compaction (usually implemented by asking another LLM to summarize the conversation) is lossy.

  - Keep a _sliding window_ of messages. There is no need to compact anything, but the LLM will not "remember" what the user said in earlier messages. Is that acceptable?

  - Leverage _subagents_. Each _subagent_ has its own context window. The _subagents_ report back to the main agent with their findings. Here, you are performing _compaction_, but at the subagent level. This might produce better results than regular _compaction_.

  - Use `RAG` instead of placing lots of data into the LLM context. This is a very broad and deep topic. Make sure you need it before going down this route.

  - Start a fresh conversation when the context is about to be filled. This might be the easiest to implement, but it could create a jarring user experience.

- Regarding compaction, Scott mentioned a technique where, while processing the user prompt, you have another LLM extract any salient facts from the prompt.

  - This is helpful because you are only looking at one message, which results in a faster and more accurate response.

  - Then, you can add the extracted facts to a list to be used for compaction later on.

- **You need to use another model (API call) to calculate the accurate number of tokens you have consumed so far**.

  - There are libraries like [`tiktoken`](https://github.com/openai/tiktoken) that will give you an _estimate_.

  - Anthropic has an API endpoint you can call. [Here are the docs](https://platform.claude.com/docs/en/build-with-claude/token-counting).

  All of this makes me wonder how Cursor does it. For ClaudeCode, they have the endpoint they can use, but Cursor has to accommodate various models...

## 08 Shell Tool

- In this section, we talked about the _dangers_ of giving the LLM the capabilities to execute stuff in the shell, but also how to make those operations "secure".

  - The solution most people lean towards is running those commands in a _sandbox_.

    - You can use _runtime native_ solutions like the permission systems of Node.js (I believe that one is not released yet) and Deno (available).

    - You can use [Cloudflare Sandbox SDK](https://developers.cloudflare.com/sandbox/) or other vendors.

    - You can use _system native_ solutions like `sandbox-exec` MacOS utility. **This is the solution Cursor uses at the moment**.

Start Day 2 Part 6
