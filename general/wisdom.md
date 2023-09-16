# Programming wisdom

## Write code that is easy to delete, not easy to extend

> Based on [this article](https://programmingisterrible.com/post/139222674273/write-code-that-is-easy-to-delete-not-easy-to).

This article goes in various directions, but the underlying theme is the following: **It is more than okay to delete code. Writing code that is easy to delete should be one of your priorities. You will not get the design right when writing the code for the first time**.

Of course, life is more complicated. There are layers to this premise.

1. **Consider not writing code at all**. Think about writing the code as a last resort. Sometimes, you can avoid writing code by ensuring everyone is on the same page. Think of times you wrote some piece of code and then deleted it shortly after because things changed.

2. **Copying and pasting code is healthy to a certain degree**. You would not want to copy and paste every time, but doing so a couple of times will expose patterns you could generalize. The same applies to writing boilerplate.

3. If you have a lot of boilerplate and are ready to collapse code into a module, do so. Feel free to create abstractions over more complicated code. **Thread carefully when creating an abstraction**. A wrong abstraction will do much more harm than good.

> It is not so much that we are hiding detail when we wrap one library in another, but we are separating concerns: requests is about popular http adventures, urllib3 is about giving you the tools to choose your own adventure.

4. **There is nothing wrong with writing a bunch of code, learning from it, and deleting it afterward**. It is much easier to remove one big mistake than many smaller mistakes scattered throughout the codebase. **Keep the experiments within a specific boundary that you can delete without any issues afterward**.

## Aging Code and the constant need to rewrite it

> Based on [this article](https://vadimkravcenko.com/shorts/aging-code/).

When working on any codebase, you will most likely interact with the so-called "legacy code". In most cases, this code enabled your employment in the first place – this is where the money is!

Since it is much easier to write the code than to read it, you might be inclined to have a need to rewrite parts of the "legacy". You might justify your decision by the fact that the code is not using the newest technology or it is not "performant enough" (that is rarely the case).

At this exact moment, you should stop and evaluate your thoughts/decisions carefully. **Would it be of benefit for the company to rewrite a given piece of code?** Be very specific here. If you do not have a clear answer to this question, stop now!

See, there is a certain kind of wisdom in the "old code". Since the code age is relatively high, **it is most likely battle-tested and is edge-case free**. These edge cases are the most problematic – you will most likely miss them! By **rewriting the code, you will increase the maintainability**. The code will be different, the bugs will be different, and worst of all, it will not be as battle-tested as the old, well-aged code.

Of course, sometimes the situation is so bad that there is no other way than to rewrite the code. The so-called "big ball of mud" can suck the life out of developers and make projects grind to a complete stop. In such cases, I would also tell caution – start very small, piece by piece, and resist the temptation to rewrite "everything"**.

## Software businessman mindset

> Based on [this article](https://vadimkravcenko.com/qa/how-to-stop-thinking-as-an-engineer-and-start-thinking-like-a-business-man/).

We often think about ourselves as "Software Engineers". We like to solve hard technical problems. We like to get into the "flow" and code away for hours and hours on end.

But the truth is, **in your programming job, you must also be thinking like a businessman**. You will have a unique ability to understand the technical aspects of your business that others do NOT have.

Always remember, that the **code is only a means to the end, not the other way around**. Business is business, and, in the end, it does not matter what is "under the hood". What matters is how much money it brings – this is they money that pays your bills!
