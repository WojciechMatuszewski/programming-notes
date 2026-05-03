# Claude Code

## Programmatic Tool Calling (PTC)

> See the [docs here](https://platform.claude.com/docs/en/agents-and-tools/tool-use/programmatic-tool-calling).

How you can reduce latency and optimize the context usage while working with tools.

### The baseline

In a traditional approach, each tool call results in back-and-forth between you and the model.

1. You send the message to Claude.
2. Claude decides to use a tool.
3. You send the tool results.
4. Claude reasons about the results and _might_ return with an answer OR another request for a tool call.

As you can see, the more tool calls, the bigger the context. The model has to know the output of the previous tool call to proceed.

Some models have the ability to call multiple tools in parallel, **but note that those tool calls must be independent of each other**.

If you want to perform two independent actions in one roundtrip, it makes sense to call tools in parallel. But, as soon as there is some relationship between the output of first tool with the input of the second, you are back to having to call tools sequentially in each roundtrip.

Unless...

### With PTC

PTC is quite fascinating. **PTC does NOT append the intermediary tool call results in the model context, saving you tokens and reducing latency**.

1. You send a message to Claude.
2. Claude writes a Python script. This script contains calls to functions that look as your tools. The script is asynchronous, so Claude can ask for multiple results at the same time.
3. You run the tools and provide the results back. **Those results does NOT land in Claude context yet**.
4. Claude continues executing the Python script.
5. Python script finished. **Now Claude "sees" only the output of the Python script, and not tool results you provided**.
6. Claude responds back.

You can think of this as having a Python script, that contains I/O stubs that have signature of your tools. As soon as we `await` on them, Claude will ask you for result, and replace the call-site with that result.

The Python script executes in an environment managed by Anthropic. This means that, in theory, they have access to some of the data you provided.

**PTC does not support zero-data retention policy**. The results must be persisted in that Python script so Claude can actually run it and derive the end result.

#### Runtime alternatives

That tool that creates that Python script is managed by Anthropic. If you want, you can omit it, and implement your own "execute code" tool.

You can read more about alternative runtimes [here](https://platform.claude.com/docs/en/agents-and-tools/tool-use/programmatic-tool-calling#alternative-implementations).

## Using `/rewind`

While reading tips [from this repository](https://github.com/shanraisshan/claude-code-best-practice/tree/main), I stumbled upon [this guide](https://github.com/shanraisshan/claude-code-best-practice/blob/main/tips/claude-thariq-tips-16-apr-26.md) that mentioned the usage of `/rewind`.

Multiple times I've been in a situation where Claude failed to perform a task, and I had to re-prompt it to do something differently. This means that the context now contains "data" from that failed attempt. **Depending on the situation, this may or may not be what you want**.

The "failed attempt" contains nuance that Claude could use when retrying. On the other hand, the "failed attempt" might pollute context.

**If the failed attempt does not contain useful information, consider using `/rewind` to go back to the point before that attempt and re-prompt from there**. The later messages get dropped and you can start fresh with all the knowledge you gained from the "failed attempt"
