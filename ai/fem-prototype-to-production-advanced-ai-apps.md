# Fem Agents in Production

> [Notes from this workshop](https://frontendmasters.com/workshops/advanced-ai-apps/).

> [Workshop notes](https://clumsy-humor-894.notion.site/Agents-in-Production-13754fed51a380da8ca0de6a2361a3a3).

## LLM and Agent Recap

- LLM is a _pattern recognition_ system learned to _predict_ the next most probable word in a series of words.

- **Transformers architecture allows the LLM to "pay attention" to different parts of the sentence**.

  - This allows the model to answer your query with great amount of detail and accuracy.

- An **AI agent uses the LLM as its "brains"**. It has the ability to combine multiple sources of information and consult the LLM with to produce the final response.

## Improving LLMs with Evals

- Evals **are a way to test the output of the LLM**.

  - But testing is tricky. **LLMs are non-deterministic by nature**, so you can't assert that the output will be the same every time.

  - The assertion itself consist of many variables. For example, one might use _similarity_ to the desired output as a metric to gauge whether the assertion passed or not

- You can leverage users to help you test the output of the LLM.

  - Most of the chat applications have the "thumbs up" / "thumbs down" buttons. You can use the data from those to create new evals!

- Scott oped to base his "eval framework" on the [`autoevals` library](https://github.com/braintrustdata/autoevals).

- The first eval we have created was a basic one. We were checking if AI opted to call the correct tool.

  - This is a _boolean_ type of eval. The score is either 0 or 1.

- **When running evals consider saving the system prompt and other variables that might influence the LLM response**.

  - You can then visualize this data and learn from it.

    - In fact, there is a product already doing that – [braintrust.dev](https://www.braintrust.dev/).

  - It gets pretty meta, but you can also get the results of evals and use another LLM to improve the system prompt for you.

- You can run **evals on multiple things that relate to AI**. People are **also running evals on their RAG pipelines**.

  - Scott mentioned [ragas](https://docs.ragas.io/en/stable/) as a good reference for Python.

  - The `autoevals` library also has RAG-evaluation capabilities!

## Rag

- You **use RAG to bridge the gap between AI knowledge-base and the data you want the AI to know**.

  - Data the models were trained on have a certain cutoff date.

    - If, what you are asking the AI for, happened before or after that cutoff date, the AI will not know the answer.

- The **four most basic pieces of RAG pipeline** are:

  - The **_document processing_**.

    - How you process data matters. Usually people chunk the text by X amount of tokens. To preserve the meaning of the words, the chunk X has some overlap with chunk X+1.

      - **This approach is very basic and could be considered _naive_**, but it will get you started.

  - **_Embeddings_** creation.

    - There are separate models to create embeddings – huge vectors of numbers that represent similarity of a given chunk with other chunks.

  - **Data storage**.

    - Most likely some kind of _vector database_ that allows you to effectively perform a _similarity search_.

  - **Querying the embeddings**.

    - You take the user input, turn it into an embedding, and do math to perform a _similarity search_ of some kind.

      - There are different algorithms. **The most common way of retrieval seem to be a _cosine similarity_ search**

- Additional considerations include:

  - **Re-ranking** the search results.

    - The results you get **might be similar in meaning, but lack the necessary relevance to the user query**.

  - Formulating the answer to the user.

  And many, many more...

- Scott **mentioned [kaggle.com] as a reliable source for various datasets**.

  - This will definitely come in handy!

- When querying vector databases, you might see a **parameter called `topK`**.

  - Think of this parameter **as a way to _limit_ the amount of results you get**.

    - It is **called `topK` because it has to do with _similarity_**. The `topK` term is often used in machine learning community.

Finished Part 3 -10:41
