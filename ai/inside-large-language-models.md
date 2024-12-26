# AI for Beginners: Inside Large Language Models

> Learning from [this course](https://zerotomastery.io/courses/ai-for-beginners-large-language-models/).

- Real words are often broken down into multiple tokens.

- **Setting the `temperature` parameter to 0 DOES NOT not remove all randomness from the output**.

- I **really** liked the example where we asked the ChatGPT to roll a dice. It turns out, that the chance ChatGPT will respond with `4` is much greater than any other number. **It boils down to the training data**.

  - Again, LLM will attempt to pick _the most probable_ next token whilst factoring the `temperature` into the equation.

- When reading about LLMs, you will hear people refer to things like **parameters** and **layers**.

  - Think about **layers as set of connected _components_ that takes in input data and forwards the result to the next layer**.

    - One **component can forward its output to multiple other components in another layer**.

    - Usually, the _components_ DO NOT forward their data to other components in the same layer. **This seems to be a design choice (according to ChatGPT)**.

  - Think about **parameters as the total number of connection between different components (weights) and their biases**.

    - The **_bias_ is used to _alter_ the activation function**. Think of this as **some information that gets added to the _components_ input which is used to calculate whether the _component_ should be "activated"**.

      - **If the component is "activated" it will output its data to another components in the next layer**.

- Usage of **_transformers_ unlocked the ability for the LLMs to understand the context of the message**.

  - When speaking/writing you do this instinctively. You refer to the facts that you have just said, without having to repeat them.

  - Transformers are designed to allow to _parallel processing_. This is why GPUs are in such a high demand these days.

  - **The transformers architecture unlocked the ability to "see" which token is the next probable one**.

- "The first step in a transformer is to associate each token with a high-dimensional vector – what we call its embedding"

  - If you map the vector into 3d space, the **direction of the vector correlates to its meaning**.

    - So, the bigger the embedding, the more "semantics" the vector can have.

  - Take a _generic_ embedding for a word "tower".

    - **The aim of the _transformer_ is to "enrich" this embedding to represent its meaning in a sentence**.

      - So, if I were to say "miniature tower", the "tower" embedding would be different than if I were to say "large tower".

- When we talk about _training an LLM_, we usually think of two steps:

  - The _pre-training_ phase. **In this phase, we create those weights and connections between layer _components_**.

    - At the end of this phase, you **have a _base model_**.

  - The _fine-tuning_ phase.

    - At this stage, you **feed the model more curated data**, so it learns how to "answer questions correctly".

- In one of the videos, the instructor showcased the example of the so-called **_reversal curse_**.

  - LLMs are not able to make connection between `B` and `A` if they know that `A` = `B`.

    - For example, LLM did not know the answer to the question "Who Mary Lee Pfeiffer is", but they know the answer to the question "Who is the Tom Cruise mom?".

- **It is fascinating to me that we have a whole field of work called _mechanistic interpretability_ that aims to _understand_ what is actually going on inside the LLMs layers**

  - We can _see_ the calculations happening, but we do not understand them. We do not fully understand what they mean.

  - The _mechanistic interpretability_ aims to map out the _features_ of the LLM and how they activate for a given token.

    - **A _feature_ is a collection of _components_ (neurons) that activate in a given layer when presented with a given token**. Some neurons might fire "more" and more might fire "less".

- The course talked about the _scaling laws_ of the LLMs.

  - **The more parameters the model has, the lesser the "test loss" is**.

    - Think of the "test loss" as "how well the LLM can interpret and answer questions about new data".

  - **The more training data you have, the better the LLM is**.

    - The AI-companies are syphoning more and more data, even creating _synthetic_ training data.

  - **The more calculations model does during training, the better it gets**. This is often called _training compute_.

    - This is quite problematic. Using more compute requires more energy, thus requiring more money.

  - The question now is whether the _scaling laws_ will hold up. That is whether increasing compute, parameters and dataset size will produce better and better models without diminishing returns.

## Summary

A great course on the basics of LLMs. As an additional watching material, I would recommend everyone to go through the [3brown1blue playlist related to neural networks](https://www.youtube.com/watch?v=aircAruvnKk&list=PLZHQObOWTQDNU6R1_67000Dx_ZCJB-3pi).
