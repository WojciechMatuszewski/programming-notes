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

Start at Day 1 Part 3 -32:18 (03-single-turn-evals)
