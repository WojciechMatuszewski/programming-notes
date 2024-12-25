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

TODO: https://www.youtube.com/watch?v=eMlx5fFNoYc&list=PLZHQObOWTQDNU6R1_67000Dx_ZCJB-3pi&index=7
