# About AI

## Basics

> See [this video](https://www.youtube.com/watch?v=zjkBMFhNj_g).

- You can think of LLMs as two files: huge file with all the _parameters_ of the model, and a bit of code to interpret the parameters.

  - The "magic" is embedded within those parameters. **Getting that huge parameters file is very expensive**.

    - You have to "compress" a very large chunk of "internet" into numbers. The bigger the chunk of the internet, the "smarter" the model feels, because it saw more patterns.

### Back propagation

> [Learn more here](https://youtu.be/Ilg3gGewQ5U?list=PLZHQObOWTQDNU6R1_67000Dx_ZCJB-3pi&t=369).

During training, at first, the output layer might produce very inaccurate results. This is expected!

To "help" the layers to produce desired output, we **propagate the adjustments to the weights of the layers going from right to left**. We can propagate those recursively focusing on neurons that influence the results the most.

But **you can't focus on a single output in the dataset**. If you were to do that, the output would _always_ point to that single result. **You have to calculate the adjustments for ALL outputs and then average them together and THEN apply the adjustments**.

### Hidden layers

The term "hidden layers" refers to the neural layers in-between the input and the output layer.

### Fine-tuning

When you are done training the model, the model can only generate next word in a given sentence. It does not know how to "answer" user queries in a way that you are used to when using chat-based AI applications.

**This is where the process of fine-tuning comes in**. You can **train the model on the additional dataset. This dataset includes Q&A style inputs**. Usually, companies hire a bunch of people to ask the model something and expect an answer in return.

#### RLHF

The RLHF means _reinforcement learning from human feedback_.

Image the LLM is tasked with writing a haiku. Your job is to "label" the haiku that you deem most "preferable". Then, the LLM learns from that data to produce even better haikus.

**As you can imagine, if there are humans involved, there is also some bias attached to the answers**. The more diverse "labellers" you hire, the less likely the bias impact will be, but still, it will _always_ be there.

## Vectors

When working with AI and AI-related tools, you will hear the word "vector" quite often.

**The term _vector_ most likely refers to the data structure the model uses to _understand_ the content**.

Some time ago, bunch of smart people came with an idea to **represent content (words, images) with numbers**. Think of vectors as **collection of floating point numbers**. The **more elements a given collection has, the bigger its _dimensionality_**.

Nowadays, **most of AI model providers have their _embedding_ endpoint**. This endpoint is used to create vectors for a given piece of content. Then, **you can compare two or more vectors together** to denote if they are _semantically close_ to each other – remember, at it's very basis, AI is guessing the next word based on what came prior.

## RAG

- **R**etrieval **A**ugmented **G**eneration means **adding additional set of data into the LLM "knowledge base"**. [Here is a great video about this topic](https://www.youtube.com/watch?v=T-D1OfcDW1M).

  - A good example would be asking the LLM about the planet with the highest amount of moons. **The data LLM has could be outdated** as such it might give you wrong answer. Now, if you **augment** the data LLM has with sources from, let us say NASA, the LLM would be able to give a correct answer.

    - **The LLM would first ask the "content store" for the answer**. If the answer is there, the LLM would use that as a data source. Otherwise it uses the knowledge it already has.

- In the context of AI, the **word embeddings** are representation of words as array of numbers called **vectors**.

  - You might think of embeddings as "classifications". The modal will classify some word to a given number.

  - The numbers in the vector represent how similar each word is to another word. For example, the vector for _"I took my cat for a walk_" would be similar in terms of numbers to the _"I took my dog for a walk"_.

  - The **embeddings are then feed into some kind of _similarity search_ engine** which LLM use to retrieve the final answer.

## Prompt Engineering

> 1. [Prompt engineering guide](https://github.com/brexhq/prompt-engineering?tab=readme-ov-file#what-is-a-prompt)
> 2. [Prompt engineering guide](https://www.promptingguide.ai/)

- The "prompts" are the starting points for the LLM. They are inputs that trigger the model to generate text.

  - The "better" the prompt, the better the LLM response will be.

### Hidden prompts

The **_hidden prompt_ is the input you give to the LLM before providing the user input**. This could be used to set the "scene" for the conversation, or provide some examples (see the section below), if you know what the user is going to ask about.

In addition, you depending on the use-case, you might want to ask the bot to think step-by-step. That technique proved to produce more reliable answers.

### Providing examples in prompts

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

### X-shot Prompting

The `_X_` here refers to **the amount of examples you provided to the LLM before asking the question**. While not strictly necessary, consider keeping the format of the "question" and the "response" in the example the same.

Depending on the complexity of the task, you might need to provide one example or a few examples. **Providing even a large number of examples does not guarantee the correct LLM response**.

### ReAct model

The **`ReAct model` is a framework you can ask the LLM to follow to simulate the process human might go through when researching for information**.

There are a couple of **"stages" the LLM goes through in a loop to come up to an answer**:

- The `question`.

- The `thought`.

- The `action`.

- The `observation`.

The `action` step leverages _tools_ to retrieve information. **Keep in mind that the efficiency of the LLM here is heavily dependant on the quality of the available _tools_**.

## AI Agents

- Agents _create a chain of thought_ and **interact with tools, and the LLM, on our behalf**.

  - This makes the Agent a bit of a black box making observability a bit of a pain.

  - This also means that **the context window might get pretty large** – you do not control what kind of questions, and in what format, the agent sends to the LLM.

## Evals

- **Evals are a way to "grade" the output of the LLM**.

---

> Notes from [_"Deep dive": Generative AI Evaluation Frameworks_](https://www.youtube.com/watch?v=bLHQEG4V8-E)

- "Evals" as a series of input/expected output pairs. We do not check the _exact_ match, but rather if the output _contains_ a given string.

- Involving non-engineers into the process is quite important.

  - You can have the PM to write those pairs, engineer to provide results, and then PM to "grade" them.

- **You can use an LLM to "grade" the output of the another LLM**.

---

> Notes from [_"Evaluation for Large Language Models and Generative AI - A Deep Dive"_](https://www.youtube.com/watch?v=iQl03pQlYWY).

> [Another resource from the same source](https://github.com/guidance-ai/guidance/blob/main/notebooks/testing_lms.ipynb).

- **Exact matching** is cheap, but has a lot of problems.

  - The main reason is non-determinism. **Even the slight change in the prompt could cause the LLM to have different answer**.

  - The operations GPU make are also non-deterministic in nature. This means, that even if you set the _`temperature`_ to `0`, the choice between "top token" might be different.

- **Similarity approach** _could_ look at **how much generated text** is in the **reference text**.

  - One such method is called "BLEU" which stands for "Bilingual Evaluation Understudy".

  - This method **is not that great when you want to consider the meaning or sentence structure**.

    - There _might_ be a lot of overlap, but does the sentence make sense?

- **Functional correctness** is where you **check for properties of the output**. The checking is done either manually or via LLM.

  - Given "make the output concise", is it concise?

  - Given "make it sound polite", is it polite?

- **Model based approach** is where you craft a prompt to another LLM to grade the output.

  - LLMs are really good at detecting sentiment or judging whether the answer is X.

  - **There are special models trained to be the "judge"** for such testing.

    - You should consider **using different model for evaluation and testing since models tend to favour their own answer**. I do not fully understand how is that even a thing, but apparently it is.

- In a word where AI is often used for RAG, **evaluating the accuracy of RAG is critical**.

  - Split RAG into two parts - the _retrieval_ and the _augmentation_.

    - For _retrieval_, see if, for a given query, the "retriever" returned the most relevant documents.

      - This is deterministic, as the vector values does not change, unless you re-calculate them with a different model.

      - In this case, **exact matching seem to be a good approach**.

    - For **_augmentation_, consider using model-based approach**.

## AI Gateways

- Similar to an _API Gateway_ but specialized for AI.

- In some cases, allows you to pick which model you want to make request to.

**Basically an _API Gateway_ but for AI**. The metrics are tailored for AI no matter what provider you are using.

## Model Context Protocol

> You can [read more about it here](https://www.anthropic.com/news/model-context-protocol).

From what I understand, Anthropic wants to standardize how LLMs communicate with external "things" via tools.

The idea is to create a _server_ and a _client_. The _server_ exposes the functionality. The _client_ connects to the server and uses that functionality.

The _client_ is a desktop LLM application. The _client_ has to know where the _server_ lives in the system to be able to use it. [You can read more about the architecture here](https://modelcontextprotocol.io/introduction).

## Modality & Multi-modal LLMs

The _multi-modal_ refers to the various ways YOU can interpret and interact with the world. Think _hearing_, _seeing_, _feeling_.

As for the LLMs, this would refer to the ability to **reason based on various input types, like images, text or voice**. At the time of writing this, most LLMs are multi-modal. They can interpret images, text and voice and produce coherent outputs.

## MCP server

I like to think about the MCP server as a way for LLM to _discover_ and _call_ tools in a standardized way.

- **MCP is a protocol** that defines the endpoints the server should have.
- **MCP is also set of SDKs** released by Anthropic.

Imagine every company hosting their own MCP server. All you need to do is to "plug in" to their implementation. That would be wild, would it? Well, it seems like we are on a path to be able to do that.
