# AI Prompting

> 1. [Prompt engineering guide](https://github.com/brexhq/prompt-engineering?tab=readme-ov-file#what-is-a-prompt)
> 2. [Prompt engineering guide](https://www.promptingguide.ai/)

## Hidden prompts

The **_hidden prompt_ is the input you give to the LLM before providing the user input**. This could be used to set the "scene" for the conversation, or provide some examples (see the section below), if you know what the user is going to ask about.

In addition, depending on the use-case, you might want to ask the bot to think step-by-step. That technique proved to produce more reliable answers.

## Providing examples in prompts (related to X-shot prompting)

- It is often helpful to provide some examples in your prompt. This helps LLM to generate a good response.

  ```txt
  Create three slogans for a business with unique features.

  Business: Bookstore with cats
  Slogans: "Pull-fect Pages", "Books and Whiskers", "Novels and Nuzzles"

  // More examples

  Business: Coffee shop with live music
  Slogans:
  ```

  Examples help to "ground" the LLM into reality and steer the algorithm in the right direction.

## X-shot Prompting

The `X` here refers to **the number of examples you provided to the LLM before asking the question**. While not strictly necessary, consider keeping the format of the "question" and the "response" in the example the same.

Depending on the complexity of the task, you might need to provide one example or a few examples. **Providing even a large number of examples does not guarantee the correct LLM response**.

I've read somewhere that **there is a point of diminishing returns when using this technique**. You want to provide SOME examples, but at some point, around the 10-example mark, the LLM output might be degraded.

## ReAct model

The **`ReAct model` is a framework you can ask the LLM to follow to simulate the process a human might go through when researching information**.

There are a couple of **"stages" the LLM goes through in a loop to come up with an answer**:

- The `question`.

- The `thought`.

- The `action`.

- The `observation`.

The `action` step leverages _tools_ to retrieve information. **Keep in mind that the efficiency of the LLM here is heavily dependent on the quality of the available _tools_**.

## Prompt caching

> This section is [**based on the Anthropic documentation**](https://docs.claude.com/en/docs/build-with-claude/prompt-caching). It's the provider I'm most familiar with.

Key things to understand:

1. The order in which you define the `messages`, `tools`, and `system` properties in the request `body` does not matter.

2. The `cache_control` property is called a "cache breakpoint."

3. Caching is not free. **Pricing varies based on the model you are using.**

It took me a while to understand how the `cache_control` checkpoint system works.

The "I got it" moment came when I learned that **Anthropic "glues" the `tools`, `system`, and `messages` together and then looks for the furthest valid `cache_control` checkpoint** to determine caching heuristics.

For example, suppose you add a `cache_control` checkpoint at the level of the system prompt. On the first relevant request, this would populate the cache with the contents of `tools` and `system`. If the `system` changes, the cache is invalidated.

**Since you can define multiple `cache_control` checkpoints**, suppose you now also add one to the last `tool`. Now, if the `system` changes, Anthropic is able to _reuse_ the cache for `tools`.

## Context Placement

Some time ago, [researchers released a paper](https://arxiv.org/abs/2307.03172) talking about how _where_ you put the information in the prompt affects the LLM output.

1. **If you provide too much context, the LLM will perform worse than without any context at all**.
2. LLMs seem to have the same "recency bias" that humans have. We are more likely to remember things we were told at the _start_ of the sentence than in the middle of it.
3. Embedding information at the start > at the end.

Consider an example where you want to provide the LLM with additional information about the user. You can put that information in a couple of places:

1. In the system prompt. **It will work nicely, but will prevent you from caching previous user messages if the data is dynamic**.
2. Before the last user message. **It will work ok, and allows you to cache previous user messages alongside system prompt** (assuming the system prompt does not contain dynamic data).

Usually, the option number two is better. We are trading some accuracy, but gaining shorter latency and lower costs.

## Chain of Thought (COT)

Nowadays, most of the LLMs are so-called "reasoning models". People who made them, baked the COT technique into them, to ensure the output is as good as possible. COT is nothing but asking the LLM to "think step-by-step".

Nowadays, if you are using the "reasoning" or "thinking" model, adding "think step-by-step" into your prompts might not be necessary. However, if you are using "regular" model, using COT will most likely make the output of the LLM more robust and accurate.

## Delimiters / XML Tags

**How you format the prompt matters**.

From what I've seen and read about:

1. Using markdown to format the prompt is quite effective.

2. [Anthropic recommends using XML tags to "gather" prompt content](https://docs.claude.com/en/docs/build-with-claude/prompt-engineering/use-xml-tags). Think examples or "additional data".
   1. **The tags themselves do not matter, they can be made up**. What matters is that you are consistent and close opened tags and so on.

## Personas

Another way to increase your chances of LLM outputting quality data is to ground the LLM into a persona.

```text
You are a senior frontend engineer ...
```

It is not a magic bullet, but it might help with:

1. Aligning the LLM with the task.
2. Making sure the output contains certain vocabulary used by a given profession.

**Be aware: the quality will, most likely, go down if you mis-align persona with the task**.

If you say that the LLM is a "expert cardiologist" and then ask it to review some code, it might not do that well.
