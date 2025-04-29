# APOSD vs Clean Code

> [Reading this discussion](https://github.com/johnousterhout/aposd-vs-clean-code) and making notes in the progress

- I prefer John's "deep" vs. "shallow" functions framework over the "one thing" UB attempts to articulate.

  - I agree with John – the "one thing" idea is great in principle, but in practice, it is hard to establish boundaries and know when it is taken too far. It is ambiguous in its definition – what does "one thing" _really mean_?

- When discussing the implementation of `isNotMultipleOfAnyPreviousPrimeFactor`, John makes a compelling argument against the inner `isMultipleOfNthPrimeFactor` – that, while in theory, it makes the code look nice, its name does not "scream" the side-effect that it contains.

  - I'm completely aligned. Code that makes it apparent that it modifies something _must_ make it apparent for the reader!

- John is _purely_ focused on reducing complexity. In my experience, that is the way to go.

  - While you can make _some_ assumptions about the knowledge of the other person, you _can't_ rely on them.

    - Here comes the concept of "knowledge of the world" vs. "knowledge in the head".

      - "Knowledge of the world" is something based on culture, something that _everyone_ usually knows.

      - "Knowledge in the head" is based on various factors, like skill in a particular area.

- John leans _heavily_ on comments. UB is more reserved.

  - I tend to stand behind UB in this matter. While I do not necessarily agree with "all comments are evil", I would rather think more about how to name a variable than to write a comment.

    - Comments that elaborate on the code, rather than on the _WHY_ behind the code, are a bane of my existence. They might not cause a lot of "bad", but they sure as hell are distracting.

  - Having said that, I also see John's point. If you skip comments altogether, you are at risk of _moving comments into variable names_, which produces overly long and, as John calls them, "megasyllabic names".

- UB suggested dropping a couple of sentences from a comment John added for the `isMultipleOfNthPrimeFactor` function. I agree.

  - Perhaps it's because of my short attention span, but I believe we should be deliberate with what we put inside the comments. If I sense, even slightly, that the comment repeats the code, I usually delete it.

    For example, take the following comment:

    ```ts
    // Returns if candidate is a multiple of primes[n], false otherwise
    ```

    To me, this comment is completely unnecessary. It's the role of the function name to communicate what it does.

    Of course, sometimes such comments are unavoidable. We do not live in a perfect world. But in most cases, I believe our time would be better spent thinking about the name of the function rather than how to word what the function does.

  - Having said that, I agree with some points John makes AS LONG AS his points pertain to comments that explain _why_ things work.

- When John presented their rewrite with comments, UB found a couple of "bugs" within the comments.

  - **My problem with comments is that I can't write tests against them**. I **can't catch "bugs" in the comments the same way I could catch bugs in code**.

- John, instead of TDD, advertises an approach UB calls "bundling" - instead of working in very small units, programmers should focus on tactical thinking and write larger chunks of code before writing tests for it.

  - I second this approach, but I believe it is more nuanced that "always TDD" vs. "always bundling".

    - In some cases, TDD is very useful, for example when implementing an algorithm.

    - In some cases, "bundling" might be more useful, for example when implementing a feature.

  - John argues that the "bundling" approach promotes tactical thinking more. It allows us to think about the design more, since you can actually write _most_ of the code that you want to write before refactoring it.

- Something John said made me think: "By refusing to write comments, you are hiding important information that you have and that others need".

  - I mostly agree. Perhaps my stance on comments sits in between those two. I will add code comments on things that I deem "complex" or "out of ordinary".

    - Remember: **the complexity is in the eye of the reader**. It does not matter what I think. If someone finds the code confusing, I must change it to make it more obvious.

## Summary

It was fascinating read! I love the format. I like how they choose to _write_ their thoughts rather than argue on some kind of podcast.
This allows everyone to digest what other said and respond with thoughtful manner.

Reflecting back, I see that my stance on engineering and the topics discussed often sits in-between what John and UB think.

Take comments for example. I'm more reserved that John, but I do not think comments are inherently waste of time.
When it comes to splitting code into smaller functions, I will do that, but there is a balance to be had. I read Johns book multiple times, and I love the idea of "deep interfaces".

I'm very grateful to had the opportunity to read this document!
