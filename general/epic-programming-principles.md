# Epic Programming Principles

Taking notes while [reading this series of posts](https://www.epicweb.dev/principles).

## ## Core Philosophy

- You are always building _for the people_.

  - Customers come in various forms. They could be a team that you cooperate with, or "regular customers" we usually think of.

    - Regardless, **make sure it is built well and FOR those people**. This **includes people with disabilities**.

- Facilitate in-person connections.

  - Think of the people you like to work with. What are they doing that you feel this way? **Replicate it**, but **always be yourself**.

- Positive gestures come a long way.

- Being kind works!

- Default to assuming a good intent.

- **Value your time**.

  - It is **okay to NOT build something, so you can build something else**.

    - Drop your ego. You do not have all the answers.

- Best problems are those that do not exist or solve themselves.

  - Before starting working on anything, **ask yourself if you REALLY need to solve this problem. If you do, is what you are about to do, the most sustainable way to do it?**

## Craft

- The **so-called "best practices" change with time**.

  - Do not be a hostage of "best practices". **Notice is what is constant and do that, all of the rest will change**.

- **Build what you need, and only that**.

  - You can lay foundations for the future, but do not go further than that – you can't predict the future.

- **Being pragmatic often means cutting corners but in a way that minimizes code debt**.

  - Example: pre-made components that might not look and feel great in your app could unblock a very important use-case for your customers.

- **Do not dismiss things simply because they are unfamiliar, and in turn, look "complex" or "hard"**.

  - Tools that we are unfamiliar with might look scary at first, but, if the tool is well designed, it will increase your velocity, not drag it down.

  - **Simple != familiar**.

- **Take ownership for each dependency**.

  - Dependencies are often times sources of failure. You should evaluate what you depend on carefully.

## Testing & Performance

- Test how the user would interact with your software.

  - For UIs, this means clicking the buttons and asserting on the _view_ rather than internal implementations.

  - For libraries, this means using your library _public_ API instead of importing _private_ functions.

## Debugging & Resilience

- You can prevent lots of problems by ensuring only a select few can _modify_ resources.

- **If applicable** users **should only see their data**.

- Catching errors early, ideally at compile time, is a huge win for productivity.

  - The more errors you catch this way, the lesser the chance for a runtime issue to surface. Runtime issues are expensive to fix!

## Developer Experience

- **Document your work**. You will forget the useful context as time goes one.

  - Having a document you can refer to will come VERY handy. I can guarantee you that you will regret not having written it.

- Depending on what the documentation pertains to, keeping it as close to the code as possible might be beneficial.

  - Documentation can get stale. Stale documentation is confusing and nullifies all the benefits of having it in the first place.

- **Create small, short-lived merge requests**. If necessary, split your work into parts. There are tools to do that!

  - Of course, you do not have to go into the extreme, but the smaller the diff the easier it is to review.

## Career

- **Teaching what you learned can be a great way to solidify your knowledge**.

  - This is one of the reasons I'm writing this document :)

- Communicate the value you bring to the table.

  - When having finished working on a feature, record a demo, share what you did. **Make it product-focused** so everyone can clearly see that it is the company's bottom line that benefits as well!

## Personal Growth

- Prioritize good relationships. If you do, you will surround yourself with people that care for you.

- **Strive for excellence, but do not fixate on perfection**.

  - Perfect does not exist – everything is dependant on the situation or the context. Nevertheless, you can _do your very best_ at what you do!

- **Take responsibility for your actions**. This relates to work and life in general.

  - Accountability is a trait of being an adult.

## Wrapping up

It all makes sense. The principles are _broad_ and not specific to a given technology/stack.

One has to keep in mind, that it is always easier to _talk_ about things outlined in this document, rather than to act on it. Embracing this reality will make your life more enjoyable. You will make mistakes and that is okay!

While I hate to admit it, the part about "personal marketing" is very true, and I feel it becomes ever more important as your career progresses. Sadly, I can't stomach "selling myself" or the work I do. I find it very artificial.
