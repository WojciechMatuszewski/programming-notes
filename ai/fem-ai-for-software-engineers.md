# AI for Software Engineers

> Notes from [this workshop](https://frontendmasters.com/workshops/engineering-and-ai/)

## Concepts

- In the beginning of the workshop, we have created a basic _converted_ that, given some parameters, like amount of money and the age of the account, returned information whether the user should be given a refund or not.

  - **This _converter_ is called a _model_**. There are different types of models. In our case, we created a **_decision tree model_**.

- The goal is to create a _converter_ that is the most accurate against our sample data.

  - This is quite tricky, because we have to ensure the sample data is representative of a real-world population (whatever that might be).

- Deploying any model to "production" is a multi-step process that often involves a lot of data-science related work and experimentation.

  - One example would be _pre-processing the data_ to see how the _converter_ behaves.

    - Perhaps we do not care about cents and look only at the amount of dollars as one of the parameters to decide whether to grant a refund or not.

- You can _tune the model_ after the training phase.

  - **_Tuning_ the model is quite important. It allows you to prevent "drift" where model starts to product inaccurate results**.

    - Remember that the circumstances under which the model underwent the training might be different that the current reality.

- **If we can create _converters_ based on numbers, we can also convert other things to numbers and create _converters_ for those things**.

  - Images are, set of pixels, so they can be converted to values.

  - Words are a bit harder to convert to words, but it could be done. Think _embeddings_.

## Neural networks

- As an example, we have developed a grid with filled with "multipliers".

  - If you multiply the "image grid" with the "multipliers grid" and all numbers together, you will get either -1 or 1.

    - 1 indicates a smile, -1 indicates a sad face.

- **Running the _converter_ on some input, and getting the result back** is called **inference**.

  - So, in our case, multiplying the "multipliers" grid with the grid representing an image, and getting the result back, would be called _inference_.

- **Out "multipliers" grid contains "weights"**. The grid itself is **called a "layer"**.

### Figuring out the weights

- First we started with our grid filled with zeros.

  - When, we tweaked the "significant" positions one by one.

  - Then, we applied those changes all at once.

    - Surprisingly, the net-result was positive. We got more accurate results! Note that **we did not get 100% accurate**.

      - This was surprising to me, because I would not think that, a _sum_ of changes in the "right direction" would point us to the right direction.

    - **This technique is called _gradient descent_**.

- After applying the _gradient descent_, we have noticed that our output stopped being binary (1 or -1).

  - This is where **we applied the sigmoid function** to **"squash" the numbers to percentages**.

    - Now, we will never get 0% for "definitely not a smile", and we will never get 100% for "most certainly a smile", but the values will be close to 0% and 100% respectively.

### Combining neural networks

- If you have one neural network, for example for detecting smiles, you can use that network to generate another which would generate images that contain a smile.

  - You build something basic that generates the image, and then feed that input to the existing neural network. You keep adjusting the weights in the new neural network until you are satisfied with the results.

## LLMs and Embeddings

### Representing the data (focused on text)

- Text is thought of as a series of _tokens_.

  - The **_token_ is the smallest unit of semantic meaning**. It **could be the whole word, but it also could be a sub-word**.

- **Embeddings express the model's understating of the input**.

  - If looked at in isolation, those are opaque.

    - We can visualize how LLM understands the input if we take lots of embeddings into the account.

### Self-attention

- When we talk, **we understand that words could have different meaning depending on what other words "surround" that given word**.

  - **This "understanding" is called _attention_**. That is to say, the LLM understands the correlation of a given word to another word in a sentence.

  - The **term _self-attention_ refers to the ability for the LLM to _embed_ the relevance to a given word of every other word in a sentence**.

    For example, when saying "I have no interest in", the relevant words for "interest" are "no" and "in".

    For example, when saying "The rising interest rate", the relevant words for "interest" are "rising" and "rate".

    For example, when saying "The dog chewed the bone because it was hungry", the relevant words for "it" are "dog", "chewed", and "hungry".

### Pre-training

- This is a process of building an _understanding of the dataset_ – the ability to predict its pre-training data.

  - In the context of words, that would be guessing which word should come next, and checking with the pre-training data.

    - **The model adjust itself depending on the result**.

- The outcome of this training is a **_foundational_ model**.

  - The _foundational_ model can **generate documents** that **resemble its pre-training data**.

- This is not yet the chat-like LLMs that you are used to.

- Think of pre-training like getting basic education. You know the basics of maths, biology and other subjects, but you lack the _specialization_.

### LLM Knowledge

- We do not have a good way of discovering what the model "knows".

  - We are programming our models to generate next word in the sentence, so it is understandable that we do not know what it will generate.

    - **This would explain why models "hallucinate"**.

### Guiding the model (fine-tuning)

- This process often involves showing the model lots of examples of how to respond to different queries.

  - You can make the model talk like a pirate or act like a chatbot assistant.

- Think of fine-tuning as graduating from colleague with a degree in certain subject.

  - You are now much more knowledgeable in certain subject.

### Prompting

- **Prompting is like writing code but in natural language**.

  - You _guide_ the model to provide you with the best answer.

- There are many ways to increase the likelihood of getting a good and valid answer.

  - _Few-shot_ prompting where you provide valid data in your query to "ground" the model.

Start Part 8
