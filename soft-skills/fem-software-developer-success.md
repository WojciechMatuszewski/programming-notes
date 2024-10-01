# FEM Software Developer Success: Soft Skills & Testing

> Going through [this workshop](https://frontendmasters.com/courses/dev-soft-skills/introduction/) and making notes

## Ramping up on the new role

- Ask questions.

  - You are supposed to be asking questions. No matter if you are new or already within the team for some time.

  - **At this stage, your goal should be to get yourself familiar with any rituals, or "tribal knowledge" that the team has**.

- Pair program!

  - This _might_ seem like a waste of time, but in reality, you will get up to speed _much faster_.

## Developer Growth Essentials

- **Time estimation** is critical.

  - You are supposed to have necessary skills to deduce when a given task will be finished.

    - **Tasks ARE NOT fixed entities**. You should be communicating delays and **proposing next steps**.

### Estimation

- Very hard. Everyone gets it wrong from time to time.

  - **Spiking** could be a viable strategy to uncover as many unknowns as possible up-front.

    - In this context, it means doing the work end-to-end as soon as possible, even if you will delete all that you have done at the end.

- **If you have multiple tasks to do, quickly check if there are any dependencies to accomplish a given task**.

  - Perhaps designs are missing. Perhaps you need to align with another team? **The more _dependencies_ the task has, the more time it will take**.

- Communicating more than you though you need is actually beneficial.

  - In the worst case, someone can ignore what you said. No harm done there.

  - In the "best" case, they will notice your update and act on it.

### Feedback Loops

- **It is up to you to look for feedback**. Nobody is obliged to give it to you.

  - It is a good practice to give timely feedback, but you should not be reliant on others.

- **Ask very specific questions**. Vague questions like "how am I doing" are not helpful.

- **ASK FOR HELP**. We all pretty much pretend we know what we are doing. There is NO SHAME in asking for help.

  - **When asking for help, tell what you already tried and what were the results**. If you do not, you might waste other people time.

### Surfacing Your Accomplishments

- Again, **it is up to you to advocate for yourself**.

  - Everyone is busy with their stuff. Do you really expect someone to keep track of notable things you have done for you?

- **Creating a "brag doc"** is a good strategy to accumulate a log of things you are proud of.

> Wojciech: I should be advocating for myself more. There is no shame in doing that.

## Opportunities and Networking

- **Building trust with others and networking** increases your chances of getting new opportunities.

  - The opportunities you might get are a **great source of learning and growth**.

- Remember, collaboration and teamwork are not zero-sum game. You can make it so everybody wins.

  - This is how you create trust.

- **Knowing what to build also plays a role in earning trust**.

  - If you think the project you are working on does not make sense, and you have **a valid reason for it**, speak up!

    - Pushing back should not be viewed as something negative. People will often respect you more when you push back.

### Guide to Building Relationships

- Talk to people face-to-face.

  - Either in real life or through Zoom or similar software.

- **It is the connection and the feeling that you understand the other person** are the main catalysts in any relationship.

  - People want to be heard and understood. That is an **universal thing**.

## Building Authentic Connections

- **Active listening** is all about making the other person feel heard and understood.

  - Asking questions to dig-deeper into what the other person is saying.

    > Something is bothering me
    > Why does that bother you?

  - Labeling emotions.

    > They would not listen me at all!
    > Did that feel really frustrating?

  - **When talking about emotions, ALWAYS give the other person a way out**.

    - This might mean asking for concent.

    - This might mean "following" the other person if they switch the topic.

- In the end, it is not necessarily about having more friends, but seeing the other person in more "human" way.

## Test Driven Development

- Quality code:

  - Meets the user needs.

  - Easy to read.

  - Easy to change.

  - Working code.

In this section, we started to write tests for a legacy codebase. [You can find the codebase here](https://github.com/emilybache/GildedRose-Refactoring-Kata).

- Instead of using `jest`, I decided to stick with the node built-in test runner.

  - I see it pretty much feature complete. It even has mocking!

  - The API is a bit different than the one exposed by `jest`, but I like how there are no global functions like `except`.

- **Before adding new features, consider refactoring existing code to "prepare" it for changes you are going to make**.

  - This preparation step often involves adding tests!

- A _seam_ is a technique of changing the behavior of the code without altering the original code.

  - This could mean using DI, but also wrapping _everything_ with an `else` block and writing the new code above it.

    - If you go the `else` route, consider refactoring the code when you are done.

- In the workshop, Francesca decided to comment-out the existing code and write it from scratch.

  - We had tests to cover us, but I deem this practice quite dangerous. The problem is that we _do not know_ what _we do not know_.

    - Writing tests for existing code will only go so far. We are only humans and we might miss something.

- In the workshop, Francesca decided to _collapse_ the "default" implementation with the implementation for "conjured" items.

  - I do not think that is a good idea. **We are so good at recognizing patterns that it often leads us to _collapse_ things that should not have been _collapsed_**.

    - Resist the temptation to refactor duplication until you are certain, and I really mean certain, that there is no difference between use-cases.

### Working with Existing Code

- Consider using debugger to understand what the test is doing.

- **Do not alter the code unless you understand what the code is doing**.

  - This is often called **Chesterton's Fence** rule.

## Wrapping up

An interesting combination of soft skills and testing. We spent more time talking and thinking about the soft skills, and I believe that is a good thing.

Some key points for more:

- Asking questions is never a bad thing. There are no dumb questions. Nobody cares, and you should not be afraid to ask questions.

- Pairing is a huge productivity booster. It is not for everyone, but usually people do not mind.

- **Estimating is one of the hardest thing you will do as an engineer**.

  - "Shooting an arrow" through the whole workflow end-to-end helps to uncover any unknowns.

  - **If you have multiple tasks assigned to you, check for dependencies first**.

    - You can start working on X, and in the meantime, solve the issue of missing dependency (for example missing designs) for task Y.

- **It is up to you to advocate for yourself**.

- **It is up to you to build a strong network of friends and people who might give you opportunities**.

- **Everyone wants to be heard and understood**. You can facilitate conversations that enable that.
