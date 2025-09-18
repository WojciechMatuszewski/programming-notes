# AI Evals For Engineers & PMs

> Taking notes [from this course](https://maven.com/parlance-labs/evals)

## Lesson 1 – Fundamentals & Lifecycle LLM Application Evaluation

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

### Wrapping up

- A good system prompt consists of multiple parts.

  - You have to specify the _persona_ of the LLM.
  - You have to provide clear directions and rules for the LLM to follow.
  - You must ground the LLM with examples.
  - Consider adding instructions on how the output ought to be formatted.

- **Evals are very important** and without them, you will be playing a game of whack-a-mole.

  - You will gain **deep understanding about how the LLM is performing** by **looking at the data**.

- LLMs do not "think" like humans. It's hard for us to "shift" our mindset given how similar to humans they can answer the queries.

## Lesson 2 – Systematic Error Analysis

- When attempting to improve LLMs output, you can follow the cycle of: _Analyze_ -> _Measure_ -> _Improve_ -> _Analyze_...

- The term **theoretical saturation** means that you reached a certain threshold and you see that looking at new traces do not reveal any new error types.

  - That's good! This means you have a diverse dataset to work with, and run evals on.

- **Generating synthetic data** is quite tricky. **Do NOT** blindly ask the LLM to generate you sample user queries.

  - It will be favouring queries that might "adhere" to the system prompt.

- **To generate good synthetic data**:

  - First, **focus on dimensions** of a query. **This is product-specific**.

    - For example, the _persona_ of the user, or perhaps the _scenario_ or particular _feature_ of the bot.

    - Gather all the dimensions into tuples to create combinations.

      - For example: ["Confident cook", "Short on time", "Breakfast for two"]. We have _persona_, _scenario_ and _feature_ for a bot that deals with cooking.

  - **Ask the LLM to filter tuples that are unrealistic** and review the results. **DO NOT SKIP THIS**, this is critical.

  - Use the tuples to generate the traces! Notice how realistic they are.

- **Open Coding** means reading each trace and writing brief notes about what went good and want went bad.

  - **If you do this enough times**, you will **notice some patterns start to emerge**.

  - Given a long trace, **focus on first failure you noticed**.

    - A failure in one place could lead to a cascade of failures later on in the conversation. Those "later" failures are just noise!

- **One thing that struck me during the Open Coding demo** was **how fast** they went through each trace.

  - They really did not try to understand _everything_. They zeroed-in on the conversation between the bot and the user and tried to find any "errors" in the behavior.

    - You can't spend that much time on this exercise, given how many traces you have to go through.

- After you have annotated your traces, **perform axial coding**.

  - This means **collapsing the annotations into different "failure mode" category**.

    - You may consider using LLMs for this to give you _something_ to work with.

### Wrapping up

- You can generate a high-quality synthetic data by following the three step process:

  1. Generate "dimension tuples" for your product. For example, for the recipe bot, we focused on _time_, _recipe type_, _allergens_ and _time constraint_.
  2. **Filter combinations that do not make any sense**.
  3. Ask the LLM to generate user queries based on those dimensions.

- Open coding and axial coding are very valuable.

  - _Open coding_ is just writing comments about stuff.
  - _Axial coding_ is adding labels and creating initial _taxonomy_.

- After you've done those to the **point of saturation**, you can start creating dashboards and metrics based on your findings.

  - At this point, you should have actionable insights into your product and how to improve it.o

- You can do all of this manually, or via code or via tools, like Braintrust.

  - I've done this manually and using Braintrust. Braintrust is quite nice!

## Lesson 3 – More Error Analysis & Collaborative Evaluation

- A **trace** is a full record of everything that happens in a response to a single user query.

  - Including messages, tool calls, and data retrievals.

- Why annotate at the _trace_ level? Because the trace allows you to "zoom out" and look at the bigger picture.

  - **Focus on first failure you see**. It is often the case that the first failure causes a cascade of other failures. There is no point in looking at other failures, as they might be different depending on the first failure.

- Using **axial codes** allows you to quickly see _how many_ problems you have in a given "bucket".

  - This allows you to prioritize. **Without axial codes, it would be hard to know which issues to tackle first**.

- For multi-turn queries, the **key** here is to _isolate_ failures.

  - You **will need to read through the conversation until you notice the first failure**. Then, you can "pluck" that failure mode from the conversation and generate more synthetic data that mimics the query that caused the failure.

- **It is vital to have some criteria on which you evaluate a trace**.

  - For example, you want your bot to be _helpful_. What does that mean? **You must avoid ambagious** when doing error analysis.

    - You also should **elaborate on "FAIL" criteria**. When should we consider the response "bad"?

### Wrapping up

This lesson was focused on how to conduct a _collaborative_ error analysis.

- Make sure you have a team of experts.

- **Do not allow ambiguity to creep into the process**. This is true no matter if you do error analysis yourself of with a team.

- With more people involved in the process, there will be more time spent arguing about axial codes and failure modes.

  - **There must be one person that has the right to make a decision to push the process forward**. There is no point in discussing one trace for 20 minutes.

    - If you notice people arguing, that means there is misalignment in failure/pass rubrics.

- **Look at the whole trace instead of a single span**.

- Focus on the first failure you see.

  - If you need to generate more data, focus on that single span that lead to failure. Can you "isolate" it, just like you would when trying to reproduce a "regular" bug in software?

## Lesson 4 – Automated Evaluators

- Before we attempt to create an automated evaluator, we **have to know what is the difference between _specification_ and _generalization_ failure mode in our application**.

  - The **_specification_ failure is where the prompt or instructions were unclear or incomplete**. For example, the bot did not not provide amenity photo because the prompt did not ask it to.

    - **These should be fixed by you, manually**. Fix the prompt first!

  - The **_generalization_ failure is where LLM fails to apply clear instructions correctly across various diverse inputs**. For example, the LLM hallucinates a tool call, OR fails to include a constraint from the user query when calling a tool.

    - **These are prime candidates for automated evaluators**.

- There are two main "types" of automated evaluators.

  - The **code-based evals** which are cheap, deterministic and interpretable.

    - You run the LLM, get it's output and assert on the output. For example, checking if a certain word exists in the output.

  - The **LLM-as-Judge evals** which are much more expensive, and subjective.

    - Those are best used for things that are _subjective_, like "is the tone appropriate?" or "is the summary faithful?".

      - **Lot's of challenges associated with them**. They can be bias, inconsistent and cost you a lot of $$.

- **Writing prompt for LLM-as-a-Judge evaluator is quite hard**.

  - You have to have clear task & evaluation criteria. **Focus on one specific failure mode here**

  - You have to have **precise** pass/fail definitions.

  - You have to provide a few examples. **Make sure that you do not run the eval on those examples afterwards**.

    - Of course the evaluator will perform great if you run it against the examples.

  - You have to have clear guidelines for the output. Usually, it's a field called "reasoning" and "answer" with "Pass" or "Fail" options.

- After you have the prompt, **you must iterate on the prompt until you are aligned with the judge responses**.

  - You have to **align on both TPR (_true positive rate_) and TNR (_true negative rate_)**.

    - Usually, people skip the TNR alignment (how often the model also agrees that the trace contains failure) and this leads to many issues.

  - Do this on the _dev set_ (approximate 20% of the traces).

TODO: HW3
