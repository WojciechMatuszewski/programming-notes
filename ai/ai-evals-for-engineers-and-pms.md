# AI Evals For Engineers & PMs

> Taking notes [from this course](https://maven.com/parlance-labs/evals)

## Lesson 1

- **Evals are measurement of LLM pipeline quality**.

  - It's not a single number. They can comprise different metrics.

  - **A life without evals is not fun**. You have no idea _why_ something is bad, and you play a game of whack-a-mole to improve the system.

    - You change one thing, and another thing breaks.

- **Building LLM products is hard**:

  - (Gulf of comprehension) Developers are blind to the common ways the bot fails because they did not "see" the full variety of user requests.

    - It is _impossible_ to read every single user query, so **missing the patterns is quite common**.

  - (Gulf of Specification) Writing prompts that _seem_ clear to the user, but are ambiguous to the LLM.

    - LLMs can't guess your true intentions. They need explicit instructions.

  - (Gulf of Generalization) The LLM works perfectly for common test queries, but fails unexpectedly for new or unusual queries (which happen more than you think!).

- **Evals are not only "engineering thing"**.

  - You can have multiple people contribute to the effort of labelling the data.

    - If you are not looking at the data, you have no idea what to improve.

- **Including examples in your prompt is good, but testing against them is bad**.

  - It gives you false sense of security. OF COURSE the LLM will perform well, if you gave it an answer beforehand.

- How to write a good system prompt?

  1. Define the LLM's persona and goal. For example, "You are a helpful creative recipe assistant"
  2. Define clear and specific directives and rules.
  3. Provide the LLM with context and background data.
  4. Provide the LLM with a couple of examples to "ground" the LLM.
  5. You might want to give instructions about formatting and output.
  6. You might consider formatting the prompt using XML tags or Markdown or both.

  **Bad system prompt might make your error analysis effort futile**.

- **Define what is "good" and "bad" outcome**.

  - This is quite domain specific, but it will be very helpful for the LLM.
