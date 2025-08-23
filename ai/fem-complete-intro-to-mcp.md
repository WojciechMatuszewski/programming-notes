# Frontend Masters Complete Intro to MCP

> [Course website](https://mcp.holt.courses/)

## Learnings

- The STDIO and SSE protocols can be considered legacy nowadays.

  - Your go-to should be the HTTP Streaming Transport protocol.

- Brian **mentions that the tool descriptions should be short and to the point**.

  - "Add two numbers" vs. "Add two _real_ numbers". The word "real" could steer the LLM into very weird direction.

- You can think of an MCP as a server which _deterministically_ solves issues or gives back more information to the LLM.

  - The LLM is inherently non-deterministic.

- Regardless of which transport mechanism you choose, the messages use JSON-RPC 2.0 format.

  - While REST is focused on _resources_, like _users_ or _posts_, the RPC is focused on _methods_, like _add_ or _subtract_.

    - Instead of _using a user resource_ you _call an add function_ remotely.

## Resources

- After getting started with basics of MCP servers, we learned about **resources**.

  - Resources as **push model vs. pull model of RAG**.

    - Instead of giving the LLM the ability to fetch some data about X, you give the LLM all the information upfront.

    - For example, when you attach an image to the chat, that is a resource (but not the MCP-server one). Think of MCP-server resources like those attachments.

  - The **_resources_ are static**, but there are also **_resource templates_ which can be dynamic**.

- TIL that **you can "select" all the statement used to create the database via `SELECT sql FROM ...`**.

  - The output is so-called DDL with statements like `CREATE TABLE` and so on.

## Prompts

- Think of prompts as **pre-made commands you can give to LLM that also support variables**.

  - It's like a function that takes a set of parameters and returns a string with those parameters interpolated into the string

## Roots

- Roots **describe directories which the LLM have access to**.

  - This is **mostly for "local" MCP clients**, for example Cursor client which can potentially modify a files in X directory.

  - Not that widely supported.

## Sampling

**This does not mean what you think it means!**. A better world for it would be "generations" or "completions".

The word _sampling_ refers to the ability for the MCP server to **call the LLM that drives the MCP client**.

To be honest, I do not see where this could be helpful. Perhaps the MCP server can use _sampling_ as suggestions for the user as next steps?

1. Generate report MCP tool.
2. Sampling -> Prompt to the LLM with potential follow-ups.

But then the MCP server would need to understand the contents of the report to generate those follow-ups. Otherwise they would be overly generic and not that helpful.

Start part 6
