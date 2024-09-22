# FEM ChatGPT API for Web Developers

> Learning from [this course](https://frontendmasters.com/courses/chatgpt-api/introduction/)

## AI and Web Development

- GPT stands for **Generative Pre-Trained Transformer**.

- ChatGPT is an LLM – a _large language model_ trained to work and understand text.

  - In essence, ChatGPT is _guessing_ which word should come after a given word.

- The ChatGPT is a black box. You have no way of "peeking in" to see what is inside.

- When the _prompt_ and the _result_ are measured in _tokens_.

  - _Token_ is something you pay for.

  - _Token_ is a sequence of characters or sub-words that the model uses as the basic unit of processing and understanding the natural language.

  - **The number of tokens is a variable of the model**. [Check out the _tokenizer_ website from OpenAI to learn more](https://platform.openai.com/tokenizer).

- The **data they used to train those models with have a cutoff at some point in time**. This means that the results you get will not be the most up to date.

  - This might be problematic if you have specific question about a new technology, but it is usually fine for general-level questions.

- We used to have the ability to add plugins in the ChatGPT web console to augment the results we get from the LLM.

  - Nowadays, you pick specific GPTs that were trained to tailor for more specific tasks.

- When it comes to security, the biggest risk seem to be _prompt injection_ or _jailbreaks_. [You can read more about prompt injection here](https://github.com/greshake/llm-security).

  - The creators of LLMs work hard on prevent the LLM to produce malicious output.

  - People have found a ways to "outsmart" the LLM resulting in weird outputs, like instructions to create a bomb from every-day use items.

- **There is also a concept of _Agent_** – a tool that is able to _execute tasks_ on the behalf of the user.

  - Agent works in tandem with the LLM, repeatably asking the LLM questions and acting on the answers.

- The **_temperature_ defines how often the LLM picks the "most probable" next word**.

  - You might think that _always_ picking the most probable word would be the best strategy, but that results in repetition and makes the response feel "artificial".

  - I'm unsure if there is a scientific basis for this, but usually, the temperature around 0.7 seems to be good enough.

## ChatGPT Clone

In this section, we touched on OpenAI API and built a basic app – a great overview for someone who has not interacted with OpenAI API yet.

- By default, **when using _Chat Completions API_**, every request to the API does not include the "context" of the previous answer.

  - This means **that you have to send the entire (or summarized) context of the conversation so-far alongside the next API call**.

    - This could get costly!

  - **As an alternative**, one might consider the **_Assistants_ API** which handles this problem for you.

## Prompt Engineering

- Currently, to get the most of LLMs, you might consider writing instructions in specific style.

  - In some cases, the LLM can _hallucinate_ (make up the answer). While there is no way to completely avoid this, your best bet is to be very explicit and tell the LLM that, it is better to say _"I do not know"_ than to hallucinate an answer.

- A good strategy is to give the LLM some examples in the prompt.

  - This is called _"few-shot"_ approach.

## Embeddings & Fine Tuning

- _Fine Tunning_ means taking an LLM and training it further on specific set of data.

  - In other words, **is is the act of making the LLM more specialized** to tailor to particular need.

- _Embeddings_ are numerical vectors that represent text.

  - The numbers represent the "similarity" or _semantic relationship_ of the text to other words.

  - **To create embeddings, you would use a different LLM than the one for holding conversation**.

    - This is an example of _fine tuning_ where you have an _embeddings model_ which is specialized for creating embeddings from text.

  - **You can also create embeddings for images**. The LLM would then work on raw pixel data.

- You would store the embeddings in a _vector database_.

- There are frameworks to work with LLMs. One of the most known is _LangChain_.

## Wrapping up
