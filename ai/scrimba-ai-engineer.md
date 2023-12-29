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
