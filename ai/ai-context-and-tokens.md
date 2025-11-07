# Everything about AI context and tokens

## Saving on tokens via virtual filesystem

> Based on [this blog post](https://www.anthropic.com/engineering/code-execution-with-mcp)

Let’s say you are creating a chatbot whose role is to help users build interfaces on your page. Think _Lovelace_ or _Replit_.

You provide the LLM with lots of tools, and those tools most likely have quite complex schema definitions and large descriptions. Your product is complex, and you need to _somehow_ encode your business rules in a way the interface LLM can understand.

Everything works okay when you test it. The chatbot is _relatively_ fast, and since you are testing locally, you do not really look at the costs.

**But then you release the feature to users, and you realize that the chatbot is slow and is costing you a ton of money**.

You look at the traces, and the UI struggles to load. **All those tool definitions take up the majority of your context, and they are huge!**.

What can you do about it?

### Using virtual filesystem to expose tools

The blog post I linked above mentions that LLMs are quite good at exploring the file system. Why not use this to our advantage?

The idea is simple: **instead of exposing all the tools, expose a virtual filesystem that "defines" the tools and allows the LLM to "invoke" a given tool**.

So, instead of exposing `createComponent`, `updateComponent` tools, and so on, you expose `fs.list`, `fs.describe`, and `fs.invoke` tools. This way, for example, when a user wants to update a certain component, the LLM will use `fs.describe` and `fs.invoke` on that particular tool. They do not need to know about the `createComponent` tool at all!

I have to admit. The idea is quite brilliant. **It's very smart way of doing _progressive disclosure_** which is a great way to save tokens.
