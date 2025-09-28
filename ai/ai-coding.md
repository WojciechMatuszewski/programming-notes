# Coding with AI

Things I have learned about coding with AI. This might be outdated when you read it—the field moves so fast!

## Claude Code VSCode extension and LSP

One thing that bothered me while using Claude Code was the fact that the tool was not aware of the LSP errors, like type errors in TypeScript codebase.

Apparently using the VSCode extension helps here. **If you run the Claude Code in integrated terminal with Claude Code extension installed Claude Code will be aware of LSP issues**.

See [this video](https://youtu.be/42AzKZRNhsk?t=3255) for more information.

## Research, Plan, Implement

As context windows get larger and larger, you might have heard that "context optimization" does not really matter.

In my experience, this is not the case. The more information the model has in its context (and I am making this claim as something model-agnostic), the less likely the model is to do a good job at the task.

You can learn more about this topic, and how to optimize context, [here](https://github.com/humanlayer/advanced-context-engineering-for-coding-agents/blob/main/ace-fca.md).

To combat this issue, you can implement a three-step process into your workflow:

1. Research the problem. Make sure the LLM outputs an artifact you can read. **It is imperative that you read and validate this document**.
   - After that is done, either clear or compact the conversation.
2. Based on the previous artifact, have the LLM output a plan. **Just like before, make sure to read and validate this document**.
   - After that is done, either clear or compact the conversation.
3. With both artifacts available, proceed to the implementation. Make sure the LLM context is "fresh" and not "polluted" with previous steps. The only things that should be in the LLM context are the research and the plan.

This three-step process can be **extended using subagents tied to different "roles"**. For example, the _research_ (or another phase) can be done in parallel by `security` and `frontend` agents. Then, have a "coordinator" of sorts who merges the findings together.

Parallelism is very powerful, especially since **each agent has its own context**. Using subagents effectively is a very high-leverage skill.
