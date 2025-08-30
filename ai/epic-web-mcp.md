# Epic Web MCP

## MCP Fundamentals

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
