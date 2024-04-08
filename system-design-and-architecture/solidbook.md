# SolidBook

> Notes from [this book](https://solidbook.io/).

Page 153

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

-   Code is **also a vital way for communication**. Yes, the language is quite "foreign", but your colleagues are speaking
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

## A 5000ft View of Software Design

-   To be a good _architect_, you have to understand the _low-level details_ of the system.

    -   You **cannot be a good engineer without understanding the _high-level_ details**
    -   If you do not understand them, you will turn into the _tactical tornado_ programmer. Not a good place to be.

-   **Understanding the _high-level_ details of the applications will enable you to cater to customers better**.

-   Khalil shows a "map" of _software design_. There is a lot to learn!

-   Keep in mind that, **any additions to your code, even when using the "best" _design pattern_, could increase the complexity of the code**.

    -   There are no "free lunches" in software. Everything has a price.

-   The software of today, at the core architectural level, does not differ that much from the software from X years ago.
    -   The tools have changed, but the same patterns still apply.

This chapter was a high-level overview of what's to come, so I did not take that many notes. We will be touching on everything included in this chapter in detail later on.

## Part 2: Humans & Code

-   It is vital to remember that the application has **two types of customers: the users and your fellow developers**.

    -   Applying the principles of XP and software design, helps us cater to both groups.

-   By running various pools and researching what "experts" have to say, Khalil noticed that the **clean code is as much about the humans who read the code as it is about the structure of the code**.

    -   That makes sense, since **most of our time is spent reading code, than writing code**.

-   A good place to start is to **establish a _coding standard_**.

    -   Think code formatting, structure, naming things and other stuff.
    -   **_Coding standard_ avoids bike shedding**. You all agreed on X, so let us stick to X.
        -   Of course, there is additional benefit of consistency. A very vital aspect of software. Humans like patterns.

-   The following are key concepts pertaining to _clean code_ as noted by researching "experts" opinions and various pools.

    -   Runs all the tests.
    -   Contains no duplication.
    -   Maximizes clarity.
    -   Has fewer elements.

        They all make sense. They also could be applied to things outside of software (especially the _has fewer elements_), for example, to things like cars.

-   Writing code is the constant tug-of-war between _structure_ and _developer experience_.

    -   While working on problems, you decided on a given _structure_. This decision, usually, has huge implications on other developers.
        -   The more complex the _structure_ is, the harder it will be to contribute to the project.

-   Depending on the company you work for, _developer experience_ might be an essential aspect to the company surviving.
    -   For others, it might not be that critical (but, I would argue, it should). **How important the DX is, reflects on the culture inside the company**.

## Human-Centered Design For Developers

-   Just like there are the _software design principles_, there is a list of _human-centered design principles_.
    -   Why should you care? Well, in most cases, you work with other people right?
        -   Being mindful about your colleagues will help you achieve a good balance between purist approach (where you only are concerned with structure) and the DX approach (where you are only concerned about DX).

> HDC (Human-Centered Design) is a design philosophy that puts the users' needs, behavior, characteristics, pain points, and motivations first.

-   The **crucial goal of HDC** is to **optimize for _discoverability_ and _understanding_**

    -   Think about the last time you had to onboard to a new team. To be productive as soon as possible, you had to learn how the code works and how it is structured.
        -   If we make that process as easy and as fast as possible, we get ourselves a big benefit that yields dividends over-time.

-   When performing any action (like adding a new feature to the codebase), we go through several stages that get us from "How do I do this" to "Did it work?"

    -   A lot is happening in between those states. The faster we can go through them, the better the structure and design of our code.

-   The **affordance** is about _familiar_ properties of an object, that allows you to perform the desired action WITHOUT learning all about that object.

    -   For example, doors contain handles. Handles are for pushing or pulling, therefore, there is a good chance you will open the door by pushing or pulling.
    -   This **matters in code**. You want **to use concepts that are familiar and well established**. This is very dependent on the language you use.

-   The **signifiers** are **marks or sound or other things that can communicate intended behavior to a person**.

    -   Think of having a "PUSH" or "PULL" sticker on the door.
    -   They go hand-in-hand with affordances.
        -   Design patterns.
        -   Comments.
        -   Tests.
        -   Good variable names.

-   **Constraints limit set of possible actions we can take**.

    -   This is **vital and powerful technique in software development**.
    -   **Constrain what humans need to know about to use the class or module effectively**
        -   Make the _interface_ thin, and the implementation deep.
        -   Also known as creating a _minimal interface_.

-   The **_mapping_ is about the relationship between two sets of things**.

    -   In code, that means grouping items similar to each other by proximity. This is also known as _cohesion_.
    -   The DOM APIs are a good example – the `.getAttribute` is defined on the _element_, it also defines all other similar methods.

-   There will always be tension between the _structure_ AND _developer experience_.
    -   In some cases, putting more emphasis on _structure_ makes sense. In enterprise, since the domain is quite complex, using _structure_ to "tame" that complexity is warranted.
    -   On the other hand, creating a tool for other developers to use might require more emphasis on DX.
    -   It is **your job** to recognize and decide the level of balance between those two forces.

## Organizing things

-   Symptoms of a _poor project structure_ include the following (non-exhaustive list):

    -   Forgetting what the system does.
    -   It is hard to locate features.
    -   You spent a lot of time flipping back and forth between files (low cohesion).

-   Some of the common project structure approaches include the following:

    -   _"Just evolve it over time"_. This one **is a recipe for bike shedding and conflicts if you work in a team**.
        -   Your fellow developers have to spend a lot of time thinking: _where should I put this file?_
            -   Of course, after some time, they give up and put it somewhere random. **You are most likely guilty of this as well**.
    -   _"Package by infrastructure/technology/type"_. Here, **the features are spread across multiple files and the names of the directories does not tell a "story"** of what the project does.
        -   While these problems were also present in _"just evolve it over time"_ approach, here they are even more profound.
        -   Again, with this approach, it is hard to onboard new people.
    -   _"Package by domain"_.
        -   We have the "User" concept in our application, as such everything that has to do with users goes there.
            -   This **MIGHT give you a basic idea of what the application does when you look at the directory structure**.
        -   This approach also MIGHT help you locate features – but you have to understand the top-level domains.
        -   There still might be some file-flipping going on, so **this approach is not that great either**.

-   So, how should we group project files? **Consider structuring your project against _features_**.

    -   Yes, _features_. This is most likely what you are working on right?
        -   From the business perspective, **the _feature_ is the unit of work, NOT some kind of controller that you wrote**.

-   Consider the case of **_screaming architecture_**.
    -   If you open the application, and it _screams_ "React app" or some other technology, the project is poorly structured.
        -   The technology used is **an implementation detail**. The technologies come and go.
    -   However, if you open the application, and it _screams_ "fitness app" or some other domain, you did a good job.
        -   It will be much easier to work in such application.
