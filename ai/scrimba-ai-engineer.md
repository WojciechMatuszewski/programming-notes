# Scrimba AI Engineer course

WIP

> Notes from [this course](https://scrimba.com/learn/aiengineer).

- OpenAI models will not give you the same answer given the same question if you repeat the question.

  - This is quite fascinating. I bet there is some aspect of randomness that influences the answers structure.

- AI models have a **concept of snapshots**.

  - As OpenAI trains their models, they release "updates" to the GPT-X model family.

    - Think of "snapshots" as versions of minor versions of software.

  - The output could be completely different given different snapshots.

- **By default, models have no memory**.

  - If you want the model to have memory, you need to enrich it's data-set with such memory.

    - This is where the _RAG_ technique comes in to play.
