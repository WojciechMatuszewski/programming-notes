# Epic Web MCP

## MCP Fundamentals

> Based on [this](https://github.com/epicweb-dev/mcp-fundamentals) workshop.

- When you create a stdin-based server, **consider how logging affects the client-server communication**.

  - If you log to stdin, your logs will "interfere" with the messages send between the server and the client. **This can lead to issues**.

  - If you want to log, consider logging to stderr instead. Kind of weird, I know, but that's something you have to do if you do not want to mess with messages sent between the server and the client.

- The TypeScript SDK handles tool errors automatically. Neat!

- Resource vs. Tools. That's a hard one! When you think about it, both can be used to _expose_ data to the LLM. You can return data in a tool call, and you can return data in a resource.

  - Tools can mutate external state and perform computations. Think of tools like AWS Lambda functions.

  - Resources are static, and used for "serving" data to the client application. **Resources are like read-only tools**.

    - Every time you do `@some_file` in Claude Code, you are working with a resource.

  - **You can parametrize the URI for an MCP resource**, so it works kind of like a GET endpoint.

    ```text
    uri: me://tags/{id}
    ```

- Let us say you are building a coding agent MCP server. How do you implement the autocomplete for file names? **You can use the `complete` (and `completions` capability) callback when defining resource**.

  - You can [learn more here](https://modelcontextprotocol.io/specification/2025-03-26/server/utilities/completion).

- TIL that **you can _link_ to a resource in the tool response**. This allows you to keep the respond size of the tool quite minimal, and give the client application the ability to fetch the created resource if needed.

  - [You can learn more here](https://modelcontextprotocol.io/specification/2025-06-18/server/tools#resource-links)

  - Consider a tool that creates a large report. If you return all that data in the tool response, you are going to be polluting the context.

  - In addition, MCP spec supports subscriptions for resources. This could allow the client to always have access to the freshest data.

- When I first learned about prompts, I though they are just _templates_ for content to be pushed to the LLM.

  - TIL that you can return _multiple_ things when configuring a prompt. A good example would be to return a set of related resources, either via `resource_link` or via `resource`.

  - You can even provide completions for prompt parameters!

## Advanced MCP Features

> Based on [this](https://github.com/epicweb-dev/advanced-mcp-features) workshop.

- You can **annotate tools and resources**.

  - **For tools**, this is very **useful for the client application UI where they can present the user with different UI depending on those annotations**. I find the name `openWorldHint` particularly interesting!

    - [Full list of tool annotations here](https://modelcontextprotocol.io/specification/2025-06-18/schema#toolannotations).

- I'm unsure if this capability is there in ai-sdk, but for MCP, you **can specify the `outputSchema` for a tool**.

  - This allows the _server_ (in our case the TypeScript SDK) to validate the `structuredContent` against the `outputSchema`.

  - For backwards compatibility, you ought to also return the result as `text`.

- **Elicitation** is where **the server asks for either additional details or the confirmation from the client mid-tool-execution**.

  - This is **not the same** as **asking the user to confirm a tool invocation**.

    - Imagine **invoking a tool with `destructive: true` hint**. You would probably want a confirmation before running this tool, and you might use _elicitation_ for any follow-ups, like cleaning up related resources.

  - So, it is almost like we have **two HITL patterns, one for tool-follow-ups (the elicitation) and one for tool invocations**.

    - The example we implemented in the workshop lives in, what I would call, a "gray zone".

      - We are deleting a resource and requesting input from the client before we do that. BUT the client should already asked the user to run it BEFORE it invoked it. So we are forcing the user to confirm stuff twice.

  - Interestingly, **you do not have to list _elicitation_ in the _capabilities_ section of your server**.

    - I guess it's because _elicitation_ is more of a _client capability_ rather than _server capability_?

- **The idea of _sampling_ is quite interesting as well**.

  - Basically, as an MCP server, you can _ask_ the client to forward a prompt to the LLM for you.

    - This is great, as it makes the _client_ fully in control when it comes to interactions with the LLM.

  - In the course, after we created a journal entry, we send the _sampling_ request to the client for the LLM to generate tags for that entry.

- When implementing the "tool callback", you can **send notifications back to the client if the tool takes a long time to run**.

  - Great for long-running tasks. In addition, you the "tool callback" function is passed a `AbortSignal` which you can use for cancellation!

- The MCP spec denotes a robust mechanisms for notifying the client about changes.

  - In the workshop, we looked at the `listChanged` on `prompts` and _resource templates_.

  - You can even implement subscriptions!

## MCP Auth

> Based on [this](https://github.com/epicweb-dev/mcp-auth) workshop.

- The `agents` package we use in the workshop does not add CORS headers to "discovery" endpoints like `/.well-known-xx`. We had to add them manually.

- TIL that **the `Response` class has static `.json` method you can use**.

  - Pretty neat! No need to `JSON.stringify` your data anymore.

---

Finished Metadata Discovery 02. We need to wait for more content to be available.
