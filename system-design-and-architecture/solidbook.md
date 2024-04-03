# SolidBook

> Notes from [this book](https://solidbook.io/).

Page 76

## Introduction

-   Khalil mentions that the goal of the code is to _"create products for customers"_.
    -   Agree with this statement, but I would also add _"to solve problems"_.
    -   Keep in mind that **the code always has two groups of customers – the users and the developers**.
        -   One cannot forget about developers. If you do, you will create a _big ball of mud_ that is tough to with.

## Up to Speed

-   In our trade, we have to consider _complexity_ as the "enemy".
    -   Think how easy it is to make a mess in your room. What about those dishes in the sink? Now, given how inherently
        complex the code is, think how easy it is to go overboard.
-   **Complexity is anything that makes the system hard to understand and modify**.

    -   It creeps up on you. It is **incremental**. If you ignore it, you might be fine for a while, but you WILL reach
        a "breaking point", where maintaining a project is no longer possible.

-   Khalil describes two sides of _complexity_.

    -   There is the **essential complexity**. This one stems from the problem itself. Most of the problems we are solving
        are quite complex if you think about them long enough.
    -   There is the **accidental complexity**. This one is **purely driven by developers**. You increase the accidental
        complexity by suboptimal design decisions or subpar code quality.
        -   You cannot possibly get rid of it completely. But, as a good software engineer, you can reduce this type of
            complexity to a minimum.

-   Here are a few things that contribute to _accidental complexity_.

    -   **Getting the requirements wrong**: as part of your job, you ought to know what you really need to do to solve the
        problem. If what you think is the problem differs from what users think is the problem, you are in trouble.
    -   **Coupling**: If the dependencies between functions resemble spaghetti, you will not be productive in the codebase.
        **You introduce tight coupling by leveraging concrete instead of "virtual", like _interfaces_**.
    -   **Cohesion** is a measure of _how related components are within a particular module_. If the cohesion is low, you
        will have to spend more time understanding what the module does. Inversely, if the cohesion is high, you will have
        an easy time grasping what the code does.

-   Khalil proposes a few strategies to mitigate the complexity.

    -   **Strategic programing**: instead of brute forcing your way through the feature, you have both your fellow
        developers AND customers in mind. This lessens the change of making a mess.
    -   **There is a place and time for _tactical programming_ as well!** If you are building a prototype, or perhaps
        exploring the routes you can take to solve the problem, the _tactical programing_ is a viable technique. **The
        problem arises when you default to _tactical programming_ and never clean up your mess**.

-   Code is **also a vital way for communication**. Yes, the language is quite "foreign", but your collegues are speaking
    it!
    -   Picture yourself touring a country you do not know the main language of. How do you feel? Insecure? Lost at times?
        **This is exactly how developers feel if your code is a mess**. It is, like you are speaking a "foreign", hard to
        understand language.

## Software Craftsmanship

-   The _software craftsmanship_ is **professionalism** in the field of software development.

    -   You do your job well, on time and produce zero defects.
    -   You think pragmatically about problems. You do not let dogmatic ideas rule your work.

-   While you might not work on "mission-critical" systems, think of the people who are.

    -   If the software you wrote for the MRI machine has a bug, people might get killed or injured – this actually
        happened. See the `Therac-25` machine.

-   The "top-down" approach of waterfall resembles factory conditions.

    -   It has the "I'm smarter than you, so do what I say" vibes which leads to hostility and "us vs. them" mindset.
    -   Overall, not a good idea.

-   Lots of people are familiar with the _agile manifesto_, but do you know there is also the _software craftsmanship
    manifesto_?

    -   Consult [this wikipedia page](https://en.wikipedia.org/wiki/Software_craftsmanship) to learn more.

-   Khalil mentions an essential fact: **your employer is NOT responsible for ensuring you grow professionally – you
    are**.

    -   I cannot stress this enough. As a developer, you HAVE TO learn new things, constantly, to improve your craft.

-   The following really resonated with me: **_when you know the domain, the code becomes a metaphor for what happens in
    real life_**.

    -   And guess what – if you have a good DLS in place, it is easier, for EVERYONE, to grasp what is happening in the
        code.

-   The world is not all sunshine and rainbows – most people are not aware of the concept of _software craftsmanship_. So
    the question becomes: _how do I get the buy-in I need to practice its principles_?
    -   Since _money rules the world_, you have to showcase how practicing TDD or XP **makes or saves money in the long term**.
