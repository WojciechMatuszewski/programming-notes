# Scrimba AI Engineer course

WIP

> Notes from [this course](https://scrimba.com/learn/aiengineer).

## Into to AI Engineering

- OpenAI models will not give you the same answer given the same question if you repeat the question.

  - This is quite fascinating. I bet there is some aspect of randomness that influences the answers structure.

- AI models have a **concept of snapshots**.

  - As OpenAI trains their models, they release "updates" to the GPT-X model family.

    - Think of "snapshots" as versions of minor versions of software.

  - The output could be completely different given different snapshots.

- **By default, models have no memory**.

  - If you want the model to have memory, you need to enrich it's data-set with such memory.

    - This is where the _RAG_ technique comes in to play.

- In the first exercise to generate report from stock data, we passed the "raw" API response as the "user" parameter and it worked!

  - I'm pretty amazed that it did. Of course the result is not that great, but the fact that OpenAI can deduce something from unformatted API response is quite amazing!

    - Of course, **passing the "raw" API response will most likely consume more tokens**. This means that the API call to OpenAI will be more expensive.

- The **`max_tokens` settings does not control how concise the answer of the AI is**. All it does it **controls the "cut-off" length of the response**.

  - If you set the `max_tokens` to some low number, the response from the AI might be cut-off mid sentence.

- **Consider using the `stop_sequence`** to control how long the output is.

  - As soon as the model produces a string included in the `stop_sequence` it will return and stop producing further results.

    - If you tune your stop sequence, you might be able to produce a coherent output while saving some cost.

      - A good example is asking for a list of books. Usually the model produces a numbered list. If your stop sequence includes some number, for example `3.`, the model response will only include two books.

      - Another example is the _newline_ symbol. If you want your answer to be a single paragraph, without giving explicit instructions to the model, consider adding `\n` to the `stop_sequence` array.

        - Having said that, this might produce half-finished answer. It might be better to give the model explicit instructions regarding how long the output should be.

- The OpenAI playground is very helpful.

  1. It uses your API key and offers pretty much the same experience as the "chat app".

  - This is nice if you interact with the OpenAI seldomly and the 20 USD / month would not be justified for your purpose.

  2. It allows you to export what you did straight to code.

- The **_temperature_ controls how "daring" / deterministic the answer is**.

  - If you set the temperature high, the answer might be a bit random. **Some times you will get complete random words strung together**.

  - **The higher the temperature, the longer the response time**.

  - **Use lower setting for temperature to lessen the chance of factual errors**.

- There is the **"single shot" and "few shot" approaches when modeling the system message**.

  - The "single shot" means you are not giving any examples for the AI.

  - The "few shot" approach means you are providing some example answers for the AI. **This allows you to influence the answer of the AI more**.

  Consider the following payload. Notice that examples are surrounded by `###` blocks.

  ```js
  const messages = [
    {
      role: "system",
      content: `You are a robotic doorman for an expensive hotel. When a customer greets you, respond to them politely. Use examples provided between ### to set the style and tone of your response.`,
    },
    {
      role: "user",
      // You can add examples in the "user" or the "system" messages.
      content: `Good day!
        ###
        Good evening kind Sir. I do hope you are having the most tremendous day and looking forward to an evening of indulgence in our most delightful of restaurants.
        ###
  
        ###
        Good morning Madam. I do hope you have the most fabulous stay with us here at our hotel. Do let me know how I can be of assistance.
        ###
  
        ###
        Good day ladies and gentleman. And isn't it a glorious day? I do hope you have a splendid day enjoying our hospitality.
        ### `,
    },
  ];
  ```

  - **The main drawback of the "few shot" approach is the cost factor**. Since you have to provide more tokens, the API call is more expensive. **The computation time will also be higher** since the system has to parse more tokens.

- The AI security is a huge topic. For now consider the following.

  - _"Know your customer"_ – you might want to authenticate the user before allowing them to use the AI capabilities.

  - Consider using the OpenAI _moderation_ API to flag input/outputs.

    - If the input or the output is flagged, you can intervene and stop the user.

  - Stress test your application. Try to engineer malicious prompt and so on.

## Open-source AI Models

- The _HuggingFace_ website is THE place to go for open-source models.

  - They also have a nice TypeScript SDK.

- Historically, models and data-sets used to train those models were shared publicly.

  - This changed with ChatGPT.

- Nowadays, there are various closed-source and open-source models.

  - What open-source community lacks in funding, it makes up the difference with sheer amount of contributors.

- **_Inference_ refers to the process of getting predictions or decisions from a pre-trained model**.

  - Think feeding new data to pre-trained model and looking at the results.

- I'm very surprised that one can use the HuggingFace API for free.

  - And they also host various models for you, though not all. That is understandable.

    - There is a small print on the model page if the inference API is turned on/off for a given model.

- TIL **that you can create URLs from `Blob` via `URL.createObjectURL(blob)`**.

  - This is quite useful if you want to "load" some resource and then update DOM element with `src` pointing to that URL.

    - In my case, it was the "speech Blob" returned from the API that I had to turn into URL and provide to the `audio` element.

- **You can download the models locally, to the browser and use the via the `pipeline`**.

  - This is quite amazing.

## Learn Embeddings and Vector Databases

- **The word _embeddings_ refers to a process of placing one object into a different _space_**

  - On our case, it's words that get transform into numerical vectors.

  - The numbers in the vectors represent semantic meaning of the word and how similar it is to other words.

- To generate embeddings, you use a specific model. This model needs to be trained in understanding words and phrases.

  - The embeddings generated by those models are huge! **To store them efficiently, you need _vector databases_**.

- There are a lot of SaaS platforms that offer ready-to-use vector databases.

  - The trend is to build them on-top of already existing tools, like postresql.

- The _vector embeddings_ are used in **_semantic search_ where the program attempts to find similar phrases based on the input**.

  - This is done by comparing the distance between two vectors.

  - Prior to the invention of the semantic search, applications used the _lexical search_ where they relied on exact matches to return the result.

    - Of course, there were a lot of different approaches, but you get the point. The _semantic search_ powered by _vector embeddings_ is quite superior as it also includes the semantic meaning of the question.

  - In the course, to perform the _semantic search_, we are using a stored procedure.

    - Stored procedures are a "pre-compiled database functions".

      - They are nice in theory, but could be problematic in practice.

        1. They are very hard to test as they are coupled to the database.

        2. They cannot be versioned.

        3. They are database dependant.

        Overall, I would not use them in production.

- The initial exercise we did in the course went as follows.

  1. Generate embedding for the user query.

  2. Search for similar data in the database.

  3. Provide the similar data to the OpenAI API call alongside the user question..

  4. Using the `chat.completions` functions, generate the answer.

  This makes it so that AI returns answers only based on the provided context (of course, you have to provide a good `system` prompt for the AI to do that).

- When **producing vector embeddings from text, consider splitting large text into chunks**.

  - The larger the piece of text, the less "accurate" the embeddings will be in terms of information accuracy.

    - **Embeddings for large amount of text will capture the overall context, but will lack the nuance and detail**.

  - The `langchain` framework provides various _text splitter_ functions.

- In the case of OpenAI API, the simplest way to give AI "long term memory" is to append to the `messages` array.

  - **While this approach is the simplest to implement, it does not scale very well**. The input has to fit into a given token limit.

    - The OpenAI documentation recommends summarizing previous parts of conversation, or performing _semantic search_ on a conversation history and only providing the API with relevant parts of the conversation. Clever!

## Learn AI Agents

- **AI Agents are like managers in the factory**.

  - They can use tools to achieve the goals you set for them.

- The concept of _prompt engineering_ is kind of related to AI Agents.

  - The term _prompt engineering_ refers to crafting a good prompts. The prompts that will help the model to give you the best answers.

    1. The prompt has to be specific.

    2. Use technical terms you already know.

    3. Provide context.

    4. Consider **giving examples of answers to related questions**.

  - **The queries you provide to LLMs could be long – that is okay!** We have been trained to reduce our questions to a few words as this is how search engines usually operate – the fewer keywords the better results.

- Agents **could be based on a technique called _ReAct_**. [You can read more about it here](https://cobusgreyling.medium.com/react-synergy-between-reasoning-acting-in-llms-36fc050ae8c7).

  1. First, agent reasons about the input.

  2. Then, agent performs actions based the first step. It might involve calling an API or something else.

  3. Then, the agent will observe the results from the second action. Here, the agent might have enough information to answer the query. If not, the agent will repeat the steps.

- You might be thinking: _how to make the LLM be able to perform tasks?_. **At the very basic level: you cannot**. You can give hints to LLM on what to answer and then, based on the answer, provide more data to keep the conversation going.

  - Here is an **example _system prompt_ that will turn ChatGPT into a mini-agent**.

    <details>
      <summary>Click to expand</summary>

    ```txt
      You cycle through Thought, Action, PAUSE, Observation. At the end of the loop you output a final Answer. Your final answer should be highly specific to the observations you have from running the actions.

      1. Thought: Describe your thoughts about the question you have been asked.
      2. Action: run one of the actions available to you - then return PAUSE.
      3. PAUSE
      4. Observation: will be the result of running those actions.

      Available actions:
      - getCurrentWeather:
          E.g. getCurrentWeather: Salt Lake City
          Returns the current weather of the location specified.
      - getLocation:
          E.g. getLocation: null
          Returns user's location details. No arguments needed.

      Example session:
      Question: Please give me some ideas for activities to do this afternoon.
      Thought: I should look up the user's location so I can give location-specific activity ideas.
      Action: getLocation: null
      PAUSE

      You will be called again with something like this:
      Observation: "New York City, NY"

      Then you loop again:
      Thought: To get even more specific activity ideas, I should get the current weather at the user's location.
      Action: getCurrentWeather: New York City
      PAUSE

      You'll then be called again with something like this:
      Observation: { location: "New York City, NY", forecast: ["sunny"] }

      You then output:
      Answer: <Suggested activities based on sunny weather that are highly specific to New York City and surrounding areas.>
    ```

    </details>

    Then, in the code, you would have to parse the answers, call the functions and add the `Observation: ${result_of_calling_the_action}` in your code.
    **But there is a better way! Enter _OpenAI Functions_**.

- The **OpenAI functions allow you to "embed" capabilities into the data you send to the API**.

  - First, you register the functions. You do that via the `tools` property when you first prompt the model.

  - Then, you parse the response. If the LLM wants to call the function, the response have all the necessary properties for you to understand which function to call – **you no longer have to manually parse the LLM text response!**

  - Then you call the function and provide the data to LLM via the `role: "tools"` prompt.

- **When you are done building the agents using _tools_, consider using the `runTools` SDK method**.

  - What you were doing manually, the SDK will do for you.

  - The only thing you have to do is to provide the functions and some metadata.

TODO: Solo project
