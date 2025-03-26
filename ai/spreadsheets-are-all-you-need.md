# Spreadsheets are all you need

Going through lessons uploaded to [this youtube channel](https://www.youtube.com/@Spreadsheetsareallyouneed).

## Lesson 1

> [Video link](https://www.youtube.com/watch?v=FyeN5tXMnJ8).

This lesson gives a birds eye view of the core architecture of GPT-2.

- The tokenization process, where the input is split into "tokens".

  - A token is not necessarily a word. It can contain punctuation and whitespace.

- The generation of embeddings.

  - The video mentions two types of embeddings: _text_ and _position_ embeddings.

  - Remember that an embedding is a high-dimensional vector of numbers that represents a given token **and the relationship to other tokens in the input**.

- The Multi-head Attention and Multilayer Perceptron loop.

  - The attention mechanism attempts to figure out what is the most important words in the sentence and how they relate to each other. **This applies to ALL words in the sentence**.

    - This is a critical step. Without it, the answers you get from AI would not be that good.

  - The Multilayer Perceptron attempts to uncover the _meaning_ of the word in a given context. It leverages data from the Multi-head Attention layer.

- The _Language Head_ job is to pick the next token to complete the sentence.

  - This is where the _temperature_ setting comes in. Be mindful that the temperature of 0 does not necessarily mean that the model is deterministic.

## Lesson 2

> [Video link](https://www.youtube.com/watch?v=PvZN3-WqAOI).

This lesson explains how _byte-pair encoding_ or _tokenization_ works in GPT-2.

Key takeaways:

- The tokenization is a separate step from the "main" pipeline. The algorithm used to produce the tokens require a pre-computed vocabulary that is created during the training process.

- The tokenization algorithm used in GPT-2 is very english-centric. It might have some issues with other languages, like Japanese language.

- **You can tokenize images or anything else that you can assign a number to**.

  - Nowadays, we have models that work with audio, images and text at the same time.

    - These are called _multimodal LLMs_.

## Lesson 3

> [Video link](https://www.youtube.com/watch?v=v6yD5SOxOXI)

This video build on the knowledge from the second video. It talks about _embeddings_.

- We were able to _tokenize_ the input, but the tokens themselves does not mean anything.

  - That is a problem. **Context is very important for an LLM**. How we can "embed" context of a given word into numbers?

Enter _embeddings_: high-dimensional vectors of numbers that **represent features of a given token**.

- The "features" of a word allow us to derive its meaning in a given context.

- **You can perform operations on those vectors and get another vector**.

  - This is quite useful. The prime example is: `king - man + woman = queen`. This does not make sense on its own, but when viewed through the lens of embeddings, it all checks out!

- You can deduce how "similar" two given vectors are using math. The most common way to do this is using _cosine similarity_.

  - Since we are operating in n-dimensional space, measuring distance between two points could be problematic. Using _cosine similarity_ allows us to measure an _angle_ "between" different points.

- How are embeddings values created? **Embeddings are created by another LLM specifically trained to perform this task**.

- Similarly to tokens, **you can create embeddings for images and audio**.
