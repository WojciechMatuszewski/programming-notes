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

Finished https://github.com/johnousterhout/aposd-vs-clean-code?tab=readme-ov-file#johns-rewrite-of-primegenerator
