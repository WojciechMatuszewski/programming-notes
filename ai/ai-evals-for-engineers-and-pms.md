# AI Evals For Engineers & PMs

> Taking notes [from this course](https://maven.com/parlance-labs/evals)

> [Supplementary notes](https://vishalbakshi.github.io/blog/#category=AI%20Evals)

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

  - You have to have clear guidelines for the output. Usually, it's a field called "reasoning" and "answer" with "Pass" or "Fail" options.

- After you have the prompt, **you must iterate on the prompt until you are aligned with the judge responses**.

  - You have to **align on both TPR (_true positive rate_) and TNR (_true negative rate_)**.

    - Usually, people skip the TNR alignment (how often the model also agrees that the trace contains failure) and this leads to many issues.

    - In our case:

      - TPR -> how many % of times the label and the LLM agree that the recipe violates dietary restriction.
      - TNR -> how many % of times the label and the LLM disagree that the recipe violates dietary restriction.

### Automated Evaluator iteration loop

- Split the labelled traces into different sets:

  - 10-20% for the "train" set.

  - 20-40% for the "dev" set.

  - Rest for the "test" set. **NEVER** look at the "test" set.

- Evaluate the judge on the "dev" set. **You should consider including a few examples from the "train" set into the prompt**.

  - Calculate the TPR (true positive rate) and TNR (true negative rate) on the "dev" set.

    - Iterate on the "dev" set until those are above 85% or so.

- **You are not done yet!**. You have to account for judge bias. To do that, run the judge on the "test" set.

  - **We purposefully avoided looking at the "test" set to simulate the case where judge has to rate unseen data**.

- It goes without saying that **you should NOT test the evaluator on the "train" set or include examples from the "dev" or "test" sets into the prompt**

### Homework / Wrapping up

- Tuning the scorer was quite fun, but **you can make a lot of mistakes in the process**.

  - Now I understand why they pushed quite hard for having **different datasets**. If you include examples from your "test" dataset into the prompt, you will skew results.

    - **Different datasets must encompass enough PASS/FAIL scenarios**.

- You can use "tune scorer" functionality in Braintrust and it works pretty well!

- When **generating traces** you want to **use the LLM pipeline your production is using**.

  - If you have a separate prompt for answering user queries, and generating traces, there will be a mismatch between the outputs.

- To **label your data, you _might_ want to use another LLM**.

  - While this will save you time, it also have a cost. **If you do not check those labels, you are at mercy of the LLMs accuracy**. Not a fun place to be.

    - Ideally, you would label the data yourself.

> The main purpose of evals is to allow you to improve your product. If TPR (no false negatives) or TNR (no false positives) is perfect, you **should be suspicious**.

- When using LLM-based evaluators **on production, you do not have labels to compare against**.

  - Instead, **you rely on the judge output and correct for the TPR and TNR metrics**.

- **Given TPR and TNR, focus on what failure modes are more "expensive"**.

  - In our case, focusing on TPR misalignment (there is a restriction, the label says "fail", but the LLM says "pass") makes more sense, as we would not want to present someone a recipe with nuts while they mentioned they are allergic to nuts.

## Lesson 5 – More about Automated Evaluators

In this lesson we focused on _evaluating simple RAG_ pipeline.

- The RAG pipeline used the BM25 search algorithm.

- We measured metrics like:

  - "Recall" – fraction where target recipe is in top k results. **In other words, did the correct recipe appear in top K results**.

    - So, `Recall@1` or `Recall@5`.

    - You can **use this metric do decide how many results to fetch while adding context to the LLM**.

      - You would most likely pass those to the re-ranker first.

  - "MRR" – related to "Recall". **It measures WHERE in the top k results the recipe appeared**.

  A good analogy is: _Recall_ answers "Is there a book on the first shelf I check?" and the _MRR_ answers "If it's there, is it the first, second ... book?"

- To **improve MRR and Recall@K** consider **query rewriting**.

  - There are various strategies you can use here. **In the course, since we used the keyword-retrieval search, the "keyword" query rewriter improved the accuracy quite a lot**.

## Lesson 6 – Rag & Complex Architectures

- RAG is really a technique to augment the LLM context with "outside knowledge" to better answers.

- **The retriever, or what you do before "settling" on the additional context, is critical**.

- Usually, you split your dataset into chunks. You retrieve on those chunks rather than on the document as a whole.

- In HW4 (and lesson 5) we already learned how to prepare Evals for RAG retrievals.

  - First, get all the chunks.

  - Then, ask the LLM to generate a salient fact(s) about that chunk.

    - **It goes without saying that you ought to _look_ at the generated output**.

  - Then, ask the LLM to generate a question(s) based on the chunk and the salient fact.

    - **Make sure to filter those questions!** Some of them might not be realistic at all.

  With that data in place, you are able to evaluate your retrieval algorithm!

- Keep in mind the metrics we learned about.

  - Recall@K tells you what % of times the relevant chunk was found in the K retrieved documents.

  - MMR tells you how high up in the ranking is the first relevant chunk. The better the MMR, the better "precision" of the LLM is for key facts.

- **For long documents** consider parallelizing the work across multiple chunks.

  - Retrieve multiple chunks, let the LLM attempt to extract the fact you are searching for from each chunk, and then let another LLM combine the outputs.

### Common Pitfalls

- Focusing on RAG when you have not validated that the RAG is the source of the issues you see.

  - **Look at the data!**

- Blindly striving for perfect recall.

  - The higher the recall, the more context you have to push into the LLM. You can achieve Recall@1 of 100%, but at what cost?

- Skipping error analysis and hoping that changes to underlying tools will solve your issues.

## Lesson 7 – Efficient Continuous Human Review Systems

At the start, the instructors discussed some reflections from the previous lesson.

At one point, they discussed how vague prompts influence LLM output and how high users' expectations are.

**Hamel mentioned that adding UX elements to the process might help here**. So, instead of scheduling an appointment for a user via prompts, consider displaying a "calendar widget" to **avoid ambiguity**.

I really like this idea!

- In the beginning, there was "spreadsheet hell."

  - Tools like Braintrust or similar platforms did not exist yet, so annotation of traces had to be done using other available tools.

- **The advent of vendors like Braintrust does not mean that custom UIs for trace annotation are useless**.

  - It all boils down to your preferences.

- The next part of the lesson was about _designing_ the UI for trace labeling.

- One neat thing they discussed was _how_ to sample traces.

  - You can start with random sampling.

  - You can sample based on the uncertainty of your LLM judge.

  - You can sample based on failures, like user explicit ratings or errors in the tools.

  - You can also sample traces with keywords the traces you graded as "fail" contain.

- It is **essential** to remember that failure modes and criteria for success change.

  - Those vectors MUST be editable to "keep up" with the pace of product.

- **Open-coding** (so free-flowing notes) is great since it does not constraint you to a set of known failure modes. Let the LLM come up with a axial code based on what you wrote!

## Lesson 8 – Const Optimizations and Recap

Parts of this lesson focused on the recap of what we learned so far

- Look at your data.

- _LLM-in-the-loop_ not _Human-in-the-loop_. It is **critical that you are doing most of the error analysis**.

- Building "code evaluations" rather than relying on the LLM-driven ones is a good practice.

  - **The "code evaluations" are much easier to maintain and align** than the LLM-based ones.

- The instructors vibe-coded a custom UI to annotate traces.

  - One of the instructors contributed to a paper on how to go about creating such UIs.

### Accuracy Optimization

- Prompt refinement.

- Perhaps introduce step-by-step reasoning to your prompts (as examples).

- Decompose LLM calls with chain smaller calls. **This will increase your latency**.

  - For example, you might split your LLM call into "extract user intent", "call a tool" and "summarize result".

There are lot's more examples provided in the slides.

### Cost Reduction

- The instructors recommend first _focusing on accuracy_ and having a good practice established around error analysis before jumping in this section.

  - I partly agree, but I also think there are things we can do at the very start of writing our application that _just_ make sense, and that would greatly decrease the costs of running it.

    - One example would be to use cheaper models for easier tasks.

- Use caching! Most, if not all, providers allow you to fine-tune which tokens are cached.

  - **The effective use of caching might require you to restructure your prompt**. Since caching works by _prefix matching_, you would want to put all your static data at first, then the dynamic data.

## Wrapping up

This is a very, and I really mean it, comprehensive course. If you put some effort in, and _do_ the homework, it will teach you a lot.

After taking this course, I'm much more confident in introducing _evaluations_ into products I maintain.

### Recap

- Running AI systems in production is challenging. You have to consider the _three gulfs_

  - The _gulf of comprehension_: Can you understand your data and how the model actually behaves _in production_?

  - The _gulf of specification_: Can we translate, usually quite ambiguous, user prompts into precise specific instructions for the LLM to follow?

  - The _gulf of generalization_: The model is working pretty well on our examples. Can we scale it to work for _general_, organic, data? **This is where evals come in**.

- There are **multiple pitfalls** you might fall into:

  - You can be tempted to _outsource_ the trace annotation and the _error analysis_ process. **You should be looking at the data and it's because YOU know what is the best for your product**.

  - You could be tempted to create LLM judges very early on hoping to have a magical automatic way of improving your product.

    - That approach will not work that well. **Creating a good LLM judge requires lots of examples and fine-tuning**. You can't get the examples without doing _error analysis_ properly.

  - You can start testing your judges on the data you used to craft the prompt the judge on. If you do, your judge will be _very good_ for the situations you are testing it against, but might falter when judging other data.

  - You can start improving your system prematurely. **Be done with _error analysis_ before improving your system. This allows you to look at the big picture**.

- The process of _error analysis_ begins when you look at your data.

  - First, start with **_open coding_ which means writing what is good and bad with the trace**.

  - Then, proceed with **_axial coding_ which means defining taxonomy of failure modes**. This will help you to **narrow your focus for improvements to a single area**.

- **Programmatic evaluators are very useful**. Start with those! Creating a good LLM judge takes a lot of effort. Perhaps you can cover some parts of the functionality with something that requires less work?

- **Use LLM judges for subjective and nuanced criteria**. For example, checking if the response adheres to the product messaging of your brand.

- **Building a _labelling_ and _error analysis_ interfaces MIGHT be worth your time**. Most products, like Braintrust have such capability. **Build when you outgrown those products**.
