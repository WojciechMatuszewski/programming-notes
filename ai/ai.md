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

---

> Notes from [this blog post](https://hamel.dev/blog/posts/field-guide/)

- **Consider having a binary yes/no grade for evaluations**.

  - The "scale" approach (1 through 5) introduces bias and uncertainty – what would be the difference between 3/5 and 4/5?

- If you want to use _LLM as a judge_ approach, **consider making periodic "alignment" runs that involve humans**.

  - Keep in mind that the scores the LLM gives could drift from what you would consider acceptable. That is why, from time to time, involving humans is imperative to the whole process.

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

### The SSE transport

> [You can read more about SSE and MCPs here](https://pai.dev/why-serverless-mcps-are-especially-hard-6796a09613fe).

When the specification first came out, it called for using SSE (Server-Sent Events) between the server and client.

The main problem with SSE, especially for serverless, is that **SSE is stateful**. This creates an interesting problem for serverless environments – there, we can't reliably maintain long-lived open connections since containers are frozen then booted up again.

If it requires a long-lived connection, it is quite pricey to operate. Just like websockets, there is an overhead associated with keeping multiple clients connected. To address this issue, spec authors are looking to add _streamable HTTP_ as a transport option.

### The Streamable HTTP transport

The _SSE transport_ is deprecated in favour of the _Streamable HTTP_ transport! [You can read more about this transport here](https://modelcontextprotocol.io/specification/draft/basic/transports#streamable-http).

This is a huge win for serverless environments. This transport _could_ be made stateless, but it [can also support stateful connections](https://modelcontextprotocol.io/specification/draft/basic/transports#session-management).

## Prompt caching

In most cases, products have a very robust _system prompt_ that outlines the goals and provides examples. All of this is to ensure the answer to the user query is the best it could possibly be.

Since the _system prompt_ must be sent with every request, if you do not do anything about it, the longer the conversation, the slower the response from the LLM will be – there is more data to process!

Enter _prompt caching_. I find the name misleading, as I initially thought that the solution caches the _whole_ prompt, but that is not the case.

**According to my research, only the "static" part of the prompt is cached**. Think preamble or examples in the system prompt. The "dynamic" part of the prompt is never cached. [OpenAI mentions](https://platform.openai.com/docs/guides/prompt-caching) that they can even cache tool definitions (NOT USE)!

## Notes from [Deep Dive into LLMs like ChatGPT video](https://www.youtube.com/watch?v=7xTGNNLPyMI)

### Pre-training: download and preprocess the internet

> See [this example](https://huggingface.co/spaces/HuggingFaceFW/blogpost-fineweb-v1)

- Gather as much high-quality and diverse documents from the internet as possible.

  - The better the documents, the more "knowledge" we can extract from them.

### Tokenization

- Compress the raw text gathered from the internet into _tokens_.

  - Tokens DO NOT necessarily represent a single word. One word might contain two or more tokens.

- Think of this as assigning IDs to groups of symbols (groups of letters).

  - Those IDs make up a _vocabulary_ of the LLM.

  - The **IDs are arbitrary for each token**.

### Training the Neural Network

- Pick a sequence of tokens (arbitrary and variable length) from the vocabulary and attempt to predict what token comes next.

  - This sequence of tokens that we picked is called _context_. **This is an input to the neural network**.

- Neural network's job is to _predict_ what token comes next. Each token in our vocabulary has assigned a probability.

  - When you start, those probabilities are random.

  - Since you _know_ what token should come next, you have a way to "tune" the neural network by updating those probabilities given the answer.

    - **You "nudge" the probabilities, rather than increasing the probability of the correct token to be 100%**. This is done by applying some math on those probabilities.

- The process of training the neural network is adjusting those probabilities (weights) so that the output is consistent with what we expect.

  - As you can imagine, given the size of the input and how many variations of inputs there might be, this is a very compute-expensive process.

- As you train the neural network, you **should be paying attention to something called "loss"**. It's derived from the output vs. what is the "correct" answer in the text.

  - As you update the parameters, the "loss" should decrease.

#### Neural Network Internals

- Lots of steps that perform various mathematical operations on the inputs. [See the visualization here](https://bbycroft.net/llm).

#### Inference

- Happens **after training**.

  - Process of "flipping a coin" for the next token. **This is what chat-based models like ChatGPT do**.

- This is where _fine-tuning_ comes in – the model you are using has "reasonable defaults", but you want to tweak it to your needs.

#### GPU gold rush

- GPUs are very good at performing parallel operations.

  - When you train neural networks, you perform various mathematical operations that are highly parallelizable.

#### Base models

- When you are done training, you get a _base model_.

  - **A _base model_ is an "internet token autocomplete"**. It is not yet that useful in itself.

    - It CAN predict the next token in the sentence, but it lacks "direction".

      - For example, if you were to "ask" the base model: "What is 2+2" it will go on a random ramble and most likely not give you any answers.

- Since the base models are autocompleting and picking the next word based on probabilities, **you can't trust the output**.

  - It was trained on internet data, so anything goes. While true information most likely appears more frequently, so the probability for such information is higher, the model might pick other tokens!

    - This is what we call **hallucinations**.

### Post-Training: get answers to questions

- This stage is much shorter than the _pre-training_ stage, usually taking less than a day compared to months of _pre-training_.

#### Training on conversations

- First, you generate a vast amount of questions (still very much smaller than the original dataset used to train the _base model_).

- Next, you have people answer them as they would want the LLM to answer them.

  - It is sad that people receive very low compensation for this work. All in the name of pursuit of profit. OpenAI makes billions, people can't make a living...

- Next, you train the model on those questions and answers **using the same method as in the _pre-training_ phase**.

##### Representing conversations as tokens

- So far, the data we input to the model did not follow a strict "convention". It was free-flowing text from the internet. How do we represent the "meaning" of the q&a format using tokens?

  - There is no standardized way to do this. **Each LLM provider seems to have their own format**.

  - It boils down to **"wrapping" different parts of the q&a with special tokens that the LLM has never seen**. This allows us to customize the "meaning" of those tokens!

#### Post-training inference

- The base model will "autocomplete" the next token based on "internet data". The _post-training_ model will "autocomplete" the next token based on the "conversational data" (of course, it has access to the vocabulary from the _pre-training_ phase).

  - Remember: **_inference_ is an act of predicting the next token in a sequence**.

- Think about it: when you ask a ChatGPT a question, it does not magically come up with an answer on its own.

  - It mimics how a human would answer, because that is the data the model was trained on. **This is why the answer feels so human-like**.

  - **If the question (or variation of it) is not in the dataset**, the LLM will derive the answer from its pre-training data. Still, the answer would feel human-like given all the conversations the model "saw".

### LLM physiology

- You can **think of the _parameters_ as "vague recollection"**.

  - Answers based on the neural network parameters (and the data they contain) have a much higher risk of hallucinations.

- You can **think of the _context window_ as "working memory"**.

  - Answers based on the _context_ have a much better chance to be of high quality. The LLM does not have to rely on "recollection".

- **Models can't "see" like you do**.

  - For example, if you were to ask the model to count how many characters are in a given world, the model will most likely give you an incorrect answer.

    - This is because **models "see" tokens, NOT characters**, and each token could consist of multiple characters.

#### Hallucinations

- The _post-training_ data contains answers to questions that are written in a "confident" tone.

- What happens if the question asks for something that does not exist, like a randomly generated person name?

  - The LLM will answer with the most probable next tokens **in the same, very confident way – because of the data it was trained on**.

    - This is why LLMs sound so confident, even when they make up answers.

##### Dealing with hallucinations

- **One way** to deal with hallucinations is to **add "I do not know" answers to questions the LLM is unsure of to the _post-training_ dataset**.

  - How do you know what the LLM is unsure of? You test its "knowledge boundary".

    - Generate a set of questions and answers based on some context, and then probe the LLM multiple times for answers and check whether the answers are correct.

- **Another way** is to "nudge" the LLM to use tools, mainly the web search tool.

  - Similarly to the "solution" above, we add the conversations that leverage "search tool special token" in the question. This allows the LLM to "learn" when to use a given tool and when to lean on the context.

    - Keep in mind that **the result of the _web search tool_ are "dumped" into the context**. You do not see this.

#### Knowledge of "self"

- The knowledge of "self" is "programmed in", just like the "personality" of the LLM.

  - The Q&A dataset contains questions like "Who built you?" and so on. This allows the LLM to "learn" the answers to those questions.

- If you were to ask a _base model_ such a question, it would "hallucinate" the most probable tokens.

  - Since OpenAI is quite famous, and can occur quite frequently in the _pre-training_ dataset, the answer might be that the model was built by OpenAI, even thought it was not.

#### Models need tokens to "think"

Remember: all the model does, is to pick the next token from a list of tokens given their probabilities.

- Let us say you are a human labeller and was given a task to "judge" the answer of LLM to a math question. How do you decide the "best" answer?

  - Focus on how "evenly" the LLM "spreads" the reasoning. How much computation, or reasoning is allocated for each token.

  ```text
  A1: The answer is 3. This is because...

  A2: The total cost of the oranges is 4. 13-4=9, the cost of the 3 apples is 9. 9/3=3, so each...
  ```

- Consider the `A1`. **All the computation is "crammed" into a single token of `3`**.

- Now consider the `A2`. **Notice how "evenly" the LLM spreads the computation**.

  - The answer **creates intermediate computations that the LLM can reference since they are in the LLMs _context window_**.

    - **This matters. The `A1` trains the LLM to guess. The `A2` trains the LLM to "think"**. The **`A2` applies step-by-step reasoning which is very effective way to "probe" the LLM to give you correct answer**.

- Even better, **ask the model to use _tools_ like "code"**.

  - This way, the model does not have to perform arithmetics in their working memory (context).

### SFT (supervised fine-tuning) model - reinforcement learning

**Why do we do this step?** We do it **because we do not know what solution (combination of tokens) works best of the LLM**. LLM has to "discover" that on their own.

- Given a prompt, run the LLM to generate the solution. If the solution is correct, "encourage" the model to favour that way of thinking.

  - The solution comes from the model itself. It DOES NOT come from human labellers.

- This process helps the _model_ to discover what way of thinking works for a given problem.

  - We can't know that. We _do not know_ what happens inside the model. We are only aware of the various mathematical operations.

    - Side note: there is a whole field of science to "peek under the cover" and attempt to understand what is going on inside the model. See ["Mechanistic Interpretability" for more details](https://www.transformer-circuits.pub/2022/mech-interp-essay).

- **The so-called "thinking" is an emergent property of this stage of training**.

  - As the R1 paper showed, LLMs will consume more tokens per answer as the training progresses. This is because the model develops the "wait, let me check the solution step-by-step" intuition.

    - This "step-by-step" thinking greatly increases the correctness of the answer.

- **Using RF allows the model to discover "novel" ways to approach the problem**.

  - A great example of this would be AlphaGO with the famous "move 37".

    - While not _new_, it was a very interesting, and almost creative way to proceed with the game.

### RLHF - Reinforcement Learning from Human Feedback

- So far, we've been looking at the _concrete_. You give an LLM a task, like adding two numbers, and assert that the answer contains what you would expect.

  - But how would you judge things that are "un-verifiable", like telling a good joke? **Here is where _RLHF_ comes in**.

- The "naive" approach would be to have a human look at _every_ output of the LLM and judge that. This is a "naive" approach because it does not scale well.

  - Instead, we will use another LLM that was trained on human-preference data for that particular "thing" that we are trying to get scores on.

    - Given X scores from humans, that LLM will "learn" the preference, and then judge the model we are training.

    - **This way of RLHF could be considered _lossy_, because the judge is an approximation to a human**.

- **You should limit the number of iterations in this process**.

  - It turns out, the "reward model" that we talked about is pretty good at "gaming" the model.

    - Researchers noticed that, after thousands of iteration cycles, the reward model started to give very high scores to things that do not make sense.

### Summary

Such a great video! I'm so thankful that it exists.

Remember:

1. Three stages of training

- _pre-training_ where we process as much text from the internet as possible and build the vocabulary.
- _post-training_ where we "mold" the LLM to be an assistant with q&a data with the help of human labellers.
- _SFT_ and _RLHF_ where we further improve the model.

2. Models are nothing BUT token generators.

3. Models can solve very hard problems and miserably fail at things that seem obvious to us. We do not know exactly why.

4. **Models need tokens to "think"**. You want the answer to "spread" the computation across tokens. Single world answers from the model might have a higher chance of being incorrect.

5. **Never assume the model is correct**. Use the output as first draft or inspiration.
