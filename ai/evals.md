# Notes about AI Evals

## Your AI Product Needs Evals

> Notes from [this blog post](https://hamel.dev/blog/posts/evals/)

Using LLMs in your application is not that different from operating any other software.

You must have a way to:

1. Evaluate quality – for example tests.
2. Handle debugging issues – for example logging and inspecting data.
3. Change the behavior of the system – for example prompt engineering and fine-tuning.

**Most people focus on no.3 which prevents them from improving their application**.

Without evals:

1. Addressing one change, might lead to failure modes in another. Just like in regular software!
2. Will have limited visibility into the system.
3. Will have very long and unwieldy prompts that attempt to "fix" the LLM for "just that one case" that cause your trouble.

### Level 1: Unit Tests

Like "regular" unit tests, those must run very quickly and be scoped to a given feature or "functionality" you want to make sure works.

For example, checking if the LLM uses correct nomenclature when talking to an user. Or that the LLM output does NOT include UUIDs (why would the user need to know an UUID of something?).

The author mentions that teams can outgrow these tests and move to another layer, but mentions that **this step is very vital in getting started**.

**The biggest mental shift here is that those tests DO NOT have to pass 100% of the time**. Remember – even with "temperature" settings set to 0, the LLMs are non deterministic. **The pass-rate is a PRODUCT decision**. We have to weight in the tradeoff between time spent and the "accuracy gain".

### Level 2: Human & Model Eval

At this point, you need to be orchestrating your calls with traces, so you can view the input and the output of the LLM call in some graphical interface. This way:

1. You can allow other people to look at the data.
2. **You can make decisions based on what you see** and potentially improve things by amending the system prompt or other "knobs" you have at your disposal.

If you start assessing the output of the LLM using another LLM:

1. **Use the best model you can afford for the critique**.
2. Make sure to **align with a human** from time to time.

Keep in mind that all of this is very domain and product-specific. How come a model can know all the constraints and goals of the product and your current situation? It can't. **That is why it is imperative to align with a human!**

### Level 3: A/B Testing

You have a hypothesis brought up by evaluating your system. Now it is time to "deploy" this hypothesis to users. This is usually done by A/B testing.

### Synthetic data

You can ask the LLM to generate synthetic data you can run evals on. You **can also ask the LLM to generate data which then could be used to generate data to run the evals on.** This is a bit meta, but it works! especially when you have evals already in place.

## Creating a LLM-as-a-Judge That Drives Business Results

> Based on [this blog post](https://hamel.dev/blog/posts/llm-judge/index.html)
