# AI Prompting

> 1. [Prompt engineering guide](https://github.com/brexhq/prompt-engineering?tab=readme-ov-file#what-is-a-prompt)
> 2. [Prompt engineering guide](https://www.promptingguide.ai/)

## Hidden prompts

The **_hidden prompt_ is the input you give to the LLM before providing the user input**. This could be used to set the "scene" for the conversation, or provide some examples (see the section below), if you know what the user is going to ask about.

In addition, you depending on the use-case, you might want to ask the bot to think step-by-step. That technique proved to produce more reliable answers.

## Providing examples in prompts (related to X-shot prompting)

- It is often helpful to provide some examples in your prompt. This helps LLM to generate a good response.

  ```txt
  Create three slogs for a business with unique features.

  Business: Bookstore with cats
  Slogans: "Pull-fect Pages", "Books and Whiskers", "Novels and Nuzzles"

  // More examples

  Business: Coffee shop with live music
  Slogans:
  ```

  Examples help to "ground" the LLM into reality and steer the algorithm into the right direction.

## X-shot Prompting

The `_X_` here refers to **the amount of examples you provided to the LLM before asking the question**. While not strictly necessary, consider keeping the format of the "question" and the "response" in the example the same.

Depending on the complexity of the task, you might need to provide one example or a few examples. **Providing even a large number of examples does not guarantee the correct LLM response**.

I've read somewhere that **there is a point of diminishing returns when using this technique**. You want to provide SOME examples, but at some point, around 10 mark, the LLM output might be degraded.

## ReAct model

The **`ReAct model` is a framework you can ask the LLM to follow to simulate the process human might go through when researching for information**.

There are a couple of **"stages" the LLM goes through in a loop to come up to an answer**:

- The `question`.

- The `thought`.

- The `action`.

- The `observation`.

The `action` step leverages _tools_ to retrieve information. **Keep in mind that the efficiency of the LLM here is heavily dependant on the quality of the available _tools_**.

## Prompt caching

In most cases, products have a very robust _system prompt_ that outlines the goals and provides examples. All of this is to ensure the answer to the user query is the best it could possibly be.

Since the _system prompt_ must be sent with every request, if you do not do anything about it, the longer the conversation, the slower the response from the LLM will be – there is more data to process!

Enter _prompt caching_. I find the name misleading, as I initially thought that the solution caches the _whole_ prompt, but that is not the case.

**According to my research, only the "static" part of the prompt is cached**. Think preamble or examples in the system prompt. The "dynamic" part of the prompt is never cached. [OpenAI mentions](https://platform.openai.com/docs/guides/prompt-caching) that they can even cache tool definitions (NOT USE)!

## Context Placement

Some time ago, [researchers released a paper](https://arxiv.org/abs/2307.03172) talking about how _where_ you put the information in the prompt affects the LLM output.

1. **If you provide too much context, the LLM will perform worse than without any context at all**.
2. LLMs seem to have the same "recency bias" that humans have. We are more likely to remember things we were told at the _start_ of the sentence than in the middle of it.
3. Embedding information at the start > at the end.

Consider example where you want to provide the LLM with additional information about the user. You can put that information in a couple of places:

1. In the system prompt. **It will work nicely, but will prevent you from caching previous user messages if the data is dynamic**.
2. Before the last user message. **It will work ok, and allows you to cache previous user messages alongside system prompt** (assuming the system prompt does not contain dynamic data).

Usually, the option number two is better. We are trading some accuracy, but gaining on the shorter latency and lower costs.

## Chain of Thought (COT)

Nowadays, most of the LLMs are so-called "reasoning models". People who made them, baked the COT technique into them, to ensure the output is as good as possible. COT is nothing but asking the LLM to "think step-by-step".

Nowadays, if you are using the "reasoning" or "thinking" model, adding "think step-by-step" into your prompts might not be necessary. However, if you are using "regular" model, using COT will most likely make the output of the LLM more robust and accurate.

Finish part 5
