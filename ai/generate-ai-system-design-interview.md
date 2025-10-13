# Generative AI System Design Interview

## Introduction and Overview

- Two types of models:

  - **The _discriminative_ models** can both classify and "make predictions". Think whether the imagine depicts a cat or a dog, or what will be users rating of the movie given their previous watch history.

    - **Good for sentiment analysis or movie recommendations**.

  - **The _generative_ models** are all about _generating_ stuff based on the training data.

    - **Good for chatbots and summarization**.

- All the AI buzz we have nowadays revolves mostly around _generative_ models. We sure are addicted to "content" and "new stuff", aren't we?

- Everyone talks about _parameters_ of the model, but we can't forget about FLOPs – the computation complexity.

  - The complexity depends on the architecture. You can have _sparse_ layers with few connections, and _dense_ layers with lots of connections between the neurons.

    - **The more FLOPs the bigger the need for hardware and "power" to train and run the model**.

- **The "power-law"** you might hear about tells us that **the model capabilities grow linearly with the amount of compute we throw at the model durning training**.

  - This is what gives people hope that we can effectively "buy" our way into very advanced LLMs with AGI-like capabilities.

- At this point, the book elaborated on _functional_ and _non-functional_ requirements. It's always good to remember what they are about:

  - Think of _functional_ requirements as what the application does. What kind of value it brings to the user?

  - Think of _non-functional_ requirements as all the "backend" stuff that makes the application work. Basically, how the system performs, not what it does.

- You might encounter a notion of **_model modality_**. The **_modality_ denotes what inputs the model supports**.

  - Can it analyze images? Text? Audio?

- One concept this chapter mentions, that is quite interesting to me, is the fact that **AI models can be, to some extend, trained using data _generated_ by another AI model**.

  - This is mainly done to **improve data diversity and improve scalability**. Of course, the **concern here is about the quality of such data**.

- The book touches on the process of _training the model_. Not really applicable to me, but nevertheless quite fascinating.

- **_Sampling_ means generating output from the trained model**.
