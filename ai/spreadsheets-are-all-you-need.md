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

The key takeaways for me are:

- The tokenization is a separate step from the "main" pipeline. You have to have a model specially trained to tokenize the input. This increases the complexity of the overall solution.

- The tokenization algorithm used in GPT-2 is very english-centric. It might have some issues with other languages, like Japanese language.
