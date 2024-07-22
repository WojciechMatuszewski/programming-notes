# About AI

## Basics

- Todays AI models are based on _neural network algorithms_.

  - The simplest algorithm you could use is called _linear regression_. This is where you try to guess a value based on another value.

    - You can think about _linear regression_ as plotting the dataset on x,y axis, then trying to fit a line, so that the line "touches" the most amount of points. Then, you look at the line and guess the next number based on that.

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
