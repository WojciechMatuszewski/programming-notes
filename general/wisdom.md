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

Always remember, that the **code is only a means to the end, not the other way around**. Business is business, and, in the end, it does not matter what is "under the hood". What matters is how much money it brings – this is the money that pays your bills!

## Domain knowledge is more important than coding skills

> Based on [this article](https://vadimkravcenko.com/shorts/things-they-didnt-teach-you/#h-domain-knowledge-is-more-important-than-your-coding-skills).

As previously stated in this document, being a good software engineer is also being a good business man. You **cannot be a good business man without having a deep understating of the domain you are working in**. Knowledge of the domain will make your life much easier.

Picture two people, they have the same coding skills, but one person is intimately familiar with the business domain, and the other is not. If you give them the same task, the results will be widely different.

The person familiar with the domain, will most likely ask a lot of good, clarifying questions. **They may even conclude that writing code is not necessary to do the job they were given**. Since the code is only ONE way to get the job done, and not writing code is always better than writing it, the outcome will be very beneficial to the company as a whole.

Contrast this with the artifact produced by the person who does not understand the business domain. They might default to writing code, which in itself is not a bad thing. The worst thing is that the code might work, but it does not really do what it suppose to. This wastes everyone's time and energy.

So, **whatever you are working on, please make sure you have a deep understating of the domain you are working in**.

## Being a "kind engineer"

> Based on [this youtube video](https://www.youtube.com/watch?v=wTezaqqyzlk) and [this blog post](https://kind.engineering/).

- During the code review, **ask about the why. Do not be dogmatic**.

  - People do what they think is right given their current understanding of things.

    - **Always assume that others meant good**. People do not wake up one day and choose to do nefarious things.

- Always be honest. **Do not create fake personas at work**. While it might seem beneficial at first, in the long run, it will wear you down.

  - Being your "true self" is liberating. It allows you to focus on what matters the most – the product, your colleagues and the shared goal you are marching towards.

  - **Care about the people you are working with**. People are not one-dimensional.

- Being kind also means giving direct feedback. Sometimes that feedback can be negative. That is okay.

  - People will respect you if you are honest with them. Even when talking about hard things, it is very important to be honest.

- **Encourage feedback. Be vulnerable**. If you want others to share what is on their mind, you have to show that you are also able to do that.

  - I find that talking openly about your problems, your life (of course, be reasonable here) really does wonders when it comes to psychological safety.

  - **Start by asking for criticism**. Stop giving it. This shows you are open, you are vulnerable. This build trust.

- **People should be accountable for things they do, but you should not blame them**. There is **literally ZERO benefit in blaming someone for something**.

  - You will not resolve the issue faster if you blame someone for something.

  - _We succeed together and we fail together_.
