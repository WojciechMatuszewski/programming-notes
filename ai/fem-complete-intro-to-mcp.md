# Frontend Masters Complete Intro to MCP

> [Course website](https://mcp.holt.courses/)

## Learnings

- The STDIO and SSE protocols are now considered legacy.

  - Your go-to should be the HTTP Streaming Transport protocol.

- Brian **mentions that tool descriptions should be short and to the point**.

  - "Add two numbers" vs. "Add two _real_ numbers." The word "real" could steer the LLM in a very weird direction.

- You can think of an MCP as a server that _deterministically_ solves issues or provides more information to the LLM.

  - The LLM is inherently non-deterministic.

- Regardless of which transport mechanism you choose, the messages use the JSON-RPC 2.0 format.

  - While REST focuses on _resources_, like _users_ or _posts_, RPC focuses on _methods_, like _add_ or _subtract_.

    - Instead of _using a user resource_, you _call an add function_ remotely.

## Resources

- After getting started with the basics of MCP servers, we learned about **resources**.

  - Resources as a **push model vs. pull model of RAG**.

    - Instead of giving the LLM the ability to fetch some data about X, you give the LLM all the information upfront.

    - For example, when you attach an image to the chat, that is a resource (but not an MCP-server one). Think of MCP-server resources like those attachments.

  - The **_resources_ are static**, but there are also **_resource templates_ which can be dynamic**.

- TIL that **you can "select" all the statements used to create the database via `SELECT sql FROM ...`**.

  - The output is so-called DDL, with statements like `CREATE TABLE` and so on.

## Prompts

- Think of prompts as **pre-made commands you can give to the LLM that also support variables**.

  - It's like a function that takes a set of parameters and returns a string with those parameters interpolated into the string.

## Roots

- Roots **describe directories to which the LLM has access**.

  - This is **mostly for "local" MCP clients**, for example, the Cursor client, which can potentially modify files in a given directory.

  - Not that widely supported.

## Sampling

**This does not mean what you think it means!** A better word for it would be "generations" or "completions."

The word _sampling_ refers to the ability for the MCP server to **call the LLM that drives the MCP client**.

To be honest, I do not see where this could be helpful. Perhaps the MCP server can use _sampling_ as suggestions for the user as next steps?

1. Generate report MCP tool.
2. Sampling -> Prompt the LLM with potential follow-ups.

But then the MCP server would need to understand the contents of the report to generate those follow-ups. Otherwise, they would be overly generic and not that helpful.

---

- **Brian recommends being quite specific with the tools you expose to the LLM**.

  - I second this opinion. LLMs will only get better at choosing the right tools, and if your tool is "too wide," the LLM can hallucinate and call it in very weird ways.

  - Tools can become "too wide" if you naively map your API to an MCP server.

    - The API was, most likely, originally made for humans, and humans reason a bit better than LLMs.

- Keep in mind that most LLMs are trained to do whatever you want them to do without considering tradeoffs.

  - They are mostly a "yes" machines that have a very hard time disagreeing with you.

    - **This is quite problematic**. Real engineering means arguing about the solution and thinking about tradeoffs.
