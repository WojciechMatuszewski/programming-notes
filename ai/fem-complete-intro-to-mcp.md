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

Start part 4
