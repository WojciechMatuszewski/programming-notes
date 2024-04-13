# SolidBook

> Notes from [this book](https://solidbook.io/).

Page 273

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
    -   If you structure your project this way, you can apply this the following rule: **_if it does not belong to a `feature`, it goes into the `shared` folder_**.
        -   This greatly reduces the _burden of choice_ and makes it more obvious where a given file should land.

-   Consider the case of **_screaming architecture_**.

    -   If you open the application, and it _screams_ "React app" or some other technology, the project is poorly structured.
        -   The technology used is **an implementation detail**. The technologies come and go.
    -   However, if you open the application, and it _screams_ "fitness app" or some other domain, you did a good job.
        -   It will be much easier to work on such application.

-   The most important takeaway is this: **group related files together to increase the _cohesion_ of the code**.

    -   It might be unrealistic to refactor the code structure in large projects.

-   In the **frontend application, features live on pages**. Therefore, consider structuring the project with respect to _pages_.
    -   On the **backend, the features _are_ use cases**. Therefore, consider structuring the project with respect to use cases.

## Documentation & Repositories

-   Use **tests as the MAIN tool for documentation**.

    -   Code comments are okay, but use them sparingly. Only when you need to explain the WHY and not how.
        -   The "how question" should be addressed by tests.
    -   The test should also explain how a given feature works. Ideally, new developers would be learning about the domain _through tests_.

-   **Push as much complexity downwards as possible**.
    -   This will, not only help with onboarding, but also with cognitive load while adding new features.

## Naming things

> This whole section is based on [this book](https://www.namingthings.co/).

-   Think about all the programing courses you have attended. Did they teach you _how to name things_?
    -   Sadly, for most of us, this is not the case, and yet, _naming things_ is **one of the most challenging problems we face as software engineers**.
        -   Naming things is not hard – it is the **act of naming things well that is hard**.

### Consistency & uniqueness

> Each concept should be represented by a single, unique name.

-   **Given how we learn (through repetition), having consistent names is paramount**.

    -   It improves how fast others will be able to grasp as given concept.
    -   It improves how well others can search through the codebase.

-   **Consistency not only applies to naming**. It applies to everything you have in the codebase.

    -   How you structure the application.
    -   How you define functions, properties, and classes.

-   One of the **worst things you can do is to have similar, but different, names for the same concept**.

    -   This confuses everyone and makes them question the whole _convention_.

-   **Do NOT recycle** names.
    -   This is just plain laziness. There is no excuse to do that.
    -   Again, this only creates confusion and makes the code harder to read.

### Understandability

> A name should describe the concept it represents.

-   Remember the concept of the _knowledge in the world_ vs. _knowledge in the head_?

    -   The _knowledge in the head_ is based on memory and current context.
    -   The _knowledge in the world_ is based on culture, something deeply ingrained inside us.
    -   **You want to name things based on _knowledge in the world_ as much as possible**.
        -   Doing so will reduce the amount of time one has to think about the variable to understand what it does.
        -   **You want to name the thing after a real world concept**.

-   The best name for a given construct is the name understandable by everyone, not only programmers.

    -   Yes, there is a required body of knowledge related to the _business domain_, but that is a given.

-   **Avoid overly technical or ambiguous names**.

    -   `Processor`, `Manager`, and `Helper` are all pretty vague.

-   Either **document the _side effects_ the function does, or push the _side effects_ to a separate function**
    -   If the `createUser` also sends a verification email, make sure that fellow developers know that is the case.
        -   You would need to rename the `createUser` to `createUserAndSendVerificationEmail`.
            -   **Using `and` within the name suggest that the function does too many things!**
        -   Or, better, extract the `sendVerificationEmail` to a separate function.

### Specificity

> A name should not be overly vague or overly specific

-   **It is okay to start with overly specific name** granted you refactor the name later on.

    -   That first step of having overly specific name oftentimes signals to us that we ought to split the variable / function in multiple pieces.

-   **Every variable should have _justification_ for its existence**.

    -   The fewer variables, the less context we have to keep in our heads.

-   **A name should describe what is inside the variable – nothing more, nothing less**.
    -   There is no need to get fancy and try to make the variable sound _technical_ or _smart_ (that will only undermine their usefulness).

```js
// Very vague.
function cloneArray(arr1, arr2) {}

// vs.

// "Just right", explicit.
function cloneArray(source, destination) {}
```

```js
async function exists(userId) {
    // Overly specific.
    const tempUserToDetermineIfExistsOrNot = this.userRepo.getUser()
    return !!tempUserToDetermineIfExistsOrNot
}

// vs.
async function exists(userId) {
    const user = this.userRepo.getUser()
    // "Just right". No need for additional variable.
    return !!user
}
```

-   **Name things by their _roles_ in the application**.
    -   Perhaps, instead of using `User` everywhere, we could leverage domain concept names like `Admin`, `Editor` or `Employee`?
        -   Of course, this is highly dependent on the application.

### Brevity

> A name should be neither overly short nor overly long.

-   Code that is too verbose **is hard to read**. The code that is too succinct is **also hard to read**.

    -   There is a constant tension between _compression_ and _context_.

-   To explain something, you need to provide enough context.
    -   But if you provide too much context, we can lose the "main" takeaway.

> **Communicate _what_ with names, explain _why_ with context**

-   The _why_ should be reserved for code comments.

    -   The _what_ is reserved for variable names.

-   **Some practical** rules to follow:
    -   Remove unnecessary words, like `in order to`, `be able to`, `in the event that` and so on. All of these have much shorter equivalents.
    -   Stay away from single-letter variables _unless_ they are a convention, like within loops.
    -   A good example is the `for (let i; i = 0; i++)` code.
        -   The **conventions are _knowledge in the world_**. Keep it consistent!
    -   Grouping things **is paramount as it enables you to make the names shorter**.
        -   Operations on the `User` object will NOT have to repeat the `User` in their name. This makes them easier to read, but accessing them preserves context -> `User.findOne`.

### Searchability

> A name should be easily found across code, documentation, and other resources.

I can't stress enough how important this is. If you can't find something fast, how are you supposed to make changes to it!?

-   Making sure the name is _searchable_ has several benefits

    -   It makes the refactoring easier.
    -   The name can be located in the documentation.
    -   The name can be understood where and how it works throughout the codebase.

-   **This characteristic is greatly influenced by other we have talked so far**.

    -   If the name is too short, or too generic or perhaps outdated, we will struggle with finding what we want.

-   **Practical tips** include:
    -   Avoiding _magic numbers_ and creating constants instead.
    -   Keeping the documentation up to date (this one is quite hard, is it not?).
    -   Ensuring that the names are not too short, nor too long (_context_ vs. _succinctness_)

### Pronounceability

**It is imperative that you can _pronounce_ the name correctly and without effort**.
The name is what you will be referring to when talking with other people.

Some of those people might not be your fellow developers, but other stakeholders.
You want **ALL** groups to understand what you are talking about, not only the people who directly contribute to the code.

-   **Some practical** rules to follow:
    -   Use abbreviations if they are considered _knowledge in the world_.
        -   `ssn`, `vat` and similar are universally understood by anyone.
    -   Use _camelCase_ to signal world breaks in variables.
    -   And my favourite: **booleans should ask a question or make an assertion**.
        -   `potIsEmpty()`, `isPostEmpty()` are good examples.
        -   **DO NOT** try to shorten the names unnecessarily. You gain zero leverage by doing so.
    -   Instead of `tmStmp` use `timeStamp`.

### Austerity

> A name should not be clever or rely on temporary concepts.

This boils down to keeping it professional. In the comments and with naming the variables. Everyone likes to laugh, but let's keep being snarky outside of the code.

-   **Some practical** rules to follow:
    -   Avoid being cute, funny or clever. **The easier the name is to understand**, the better.
    -   Do not include popular culture references.
        -   I would argue that one should not include ANY cultural references.

## Comments

A good comment can be very beneficial.
A bad comment could introduce _uncertainty_ in your understanding of the concept which hinders you from progressing.

**It is hard** to write a good comment. But not all is lost! Given few guidelines, you can increase your chance at writing a good comment.

-   Follow the rule: **code explains what and how, comments explain why**.

    -   For real, focus on the _why_. When writing a comment, always ask yourself WHY we wrote this code. DO NOT repeat what the code is doing.
    -   **Refactor the code to answer the _how_ question** better.

-   Comments, in general, contribute to _"visual noise"_ and clutter the code.

    -   That is why most people skip writing them. **But skipping comments is not a way to go**.
        -   As I mentioned before, always focus on the _why_.

-   One of the cardinal sins of writing comments is **repeating what the code does, but in a different language**.

    ```js
    class UserService {
    // This method gets the user by user id. -> A bad comment. Explains what the code does
    	getUserByUserId(userId: string): Promise<User>;
    }
    ```

## Formatting & Style

-   Code formatting is _very subjective_.

    -   **The most important** for any team is to come up with some kind of convention and stick to it.
    -   If you are working solo, do what works for you.

-   While _very subjective_ in nature, **there are some universal** rules about formatting that applies to everyone:
-   **Use whitespace to create "blocks" of related code** and separate those blocks.

    -   Whitespace is free. It is a very powerful tool in making the code more readable.

-   Here are **a few practical rules** to follow:

    -   _Use obvious spacing rules_. How to apply then is never _explicitly_ taught, but after a while you develop an intuition when one should use them.

        -   I'm referring to spacing between _keywords_, _identifiers_, _operators_ and _literals_.

        -   Keep the **code "density" low**.

    -   Do not have large chunks of code without indentation. **Group related statements together**.

    -   Break things on the horizontal axis if they are too long.
        -   This should be the job of your code formatter.

-   Next, Khalil makes a point about _smaller files_.

    -   I'm torn on this one.
        -   On the one hand, smaller files are easier to read.
        -   On the other, **smaller files often prevent you from seeing the duplication** thus making it hard to create good abstractions.
    -   My advice would be to pick what works best for you and stick to it.
        -   I'm unsure if being dogmatic about file size is a good strategy to have. Life and programing is almost never "black and white".

### Consistent whitespace

-   We touched about _consistency_ before, and this aspect is also quite important in the context of _style and formatting_.
    -   Stick to the same casing style (be it _camelCase_ or _snake_case_).
    -   **It does not matter what casing you choose**. What matters is that you stick to it.

### Storytelling

-   Reading the code is like reading a newspaper. **You want to _front-load_ the most important bits of the code**, and have the less important ones at the bottom.
    -   This **principle is called _Newspaper Code Principle_**.

```js
class RecordingStudio {
    // Quite important method. Let us keep it at the very top.
    public recordSong(...) {
    }

    // ... Other stuff

    // Not that important, but still vital to understand. Can be at the bottom.
    private someInternalDetails() {
    }
}
```

```tsx
// Quite important. It contains what we see on the screen.
function MyAweSomeComponent() {
    const data = getData()
    return (
        // some JSX
    )
}

// Not that important. Can be at the bottom.k
function getData() {
}
```

-   **Do your best to keep the level of abstraction consistent** throughout the function.
    -   Do not start at the very high level, utilizing a _deep_ abstraction only to dive into very minute details at the end.
        -   Hide those details into another abstraction and use that.
        -   This way, the reader knows what to expect.

### Tooling

-   **You should NEVER have to think about formatting (apart from vertical whitespace)** in your codebase.
    -   Use tools like `Prettier` and `ESLint` to do the job for you.
    -   If you do not, use are wasting yours and everyone else's time.

## Types

-   There are two main _type systems_ for any given programming language to choose from:

    -   The **static type system**. Here, the types are a concrete construct used to describe that will be stored in a variable.
    -   The **dynamic type system**. Here, the types are _implied_ and _hidden_ from the programmer.

-   One type system is not necessarily better than the other. It all depends on the people you work with and how you like to work.

    -   Having said that, **usually, the _statically_ and _"strong"_ typed languages will save you more time in the long run**.
        -   They will save you from silly mistakes we all do.
        -   But the most important thing is that **statically typed languages allow you to communicate via contracts and not concrete implementations**.
            -   This is HUGE for code composition.

-   Communicating via _contracts_ means using _interfaces_ or _types_ rather than concrete classes or objects.
    -   This way, you can implement the necessary contract differently in tests and differently in the runtime code.
    -   While you could do the same with _dynamically typed_ language, the language will not check against that contract.
        -   In other worlds, the "machine" will not work for you, you will have to work for the "machine" and ensure the contract is valid.

```ts
interface BookSaver {
    saveBook(): Promise<void>
}

// TypeScript will check the contract here.
// If you pass something that does not contain `saveBook` method, TypeScript will complain.
async function createNewBook(saver: BookSaver) {
    await saver.saveBook()
}
```

-   _Statically typed languages_ are also better at communicating the intent.

    -   They often have a keywords related to a given pattern (or family of patterns).
        -   Keywords like `abstract` or `private`, used in the right context, _scream_ a given pattern right away.

-   Khalil touches on the fact that _dynamically typed_ languages _could_ improve **initial** productivity, but that often comes at a later cost.

    -   I completely agree. I would argue that the temptation of doing "magic stuff" with _dynamically typed_ languages is much higher.

-   **TypeScript** uses a **_structural type system_**.
    -   This means, the type compatibility and equivalence are determined by the computed type's actual structure.

```ts
class Instrument {
    doSomething: () => {}
}

class Guitar extends Instrument {}

class Synth extends Instrument {}

declare function pluckGuitar(guitar: Guitar): void

pluckGuitar(new Guitar())
pluckGuitar(new Synth()) // Should not be allowed but it is! This is Duck Typing!
```

The rest of this chapter talks about the basics of TypeScript. I've skipped making note for those as I'm pretty comfortable with the language already.

## Errors and exceptions

-   **Errors and how we handle them** should not be an afterthought.

    -   They are an integral part of any application, because things fail, all the time.

-   Khalil mentions **two _common_ error-handling approaches that are pretty bad**.

    -   First, you have the "log and return `null`" pattern. It looks like this:

        ```ts
        function CreateUser(email: string, password: string) {
            const isEmailValid = validateEmail(email)
            const isPassowrdValid = validatePassowrd(password)

            if (!isEmailValid) {
                console.log('Invalid email')
                return null
            }

            if (!isPasswordValid) {
                console.log('Invalid password')
                return null
            }
        }
        ```

        This does not work well, because **the caller has no idea what failed**. It is quite bad.

    -   Another approach one might take would be to "log and throw."
        -   Again, not that great since you have to surround everything with `try/catch` blocks.
            -   Also, you **have to dig into the implementation of the function to know if it fails or not**.

-   There is a **difference between an _error_ and an _exception_**.

    -   **An _error_ is something expected**. Something that might happen. Think a duplicate email in the signup form.
        -   These are **domain-specific and ought to use domain-specific language**.
    -   **An _exception_ is something unexpected**. Think network connection timeout.

-   **To handle _errors_, consider the `Either` type**.
    -   If you are coding in Rust, that would be the `Result` type.
    -   In Go, that is returning a _non-nil_ error from the function.
    -   **Whatever you do, return the error rather throwing it**.
        -   I completely agree. After playing with Go and Rust, I can't stand how JS (and TS) handle errors.

```ts
class DomainError extends Error {}

// Return this instead of throwing!
class EmailInvalidError extends DomainError {
    constructor(email: string) {
        super('This email is invalid')
    }
}
```

-   Ok, so we know how to handle _errors_. What about the _exceptions_ ?
-   The key here is to **wrap the code that you do not own with `try/catch` statements and return with a _generic_ error**.
    -   This will greatly increase the _expressiveness_ of your code and will give the caller chance to react to a "generic" error.

```ts
async function createUser(email: string) {
    try {
        await db.createUser(email)
    } catch (error) {
        return new ApplicationError(error)
    }
}
```

-   The main goal of using _domain errors_ or "generic" application errors (derived from exceptions) is to **make the implicit explicit**.
    -   The _implicit_ is not great in programming. You do not want anyone guessing what you had in mind when writing code, so you do not want anyone guessing what the errors could be.

## Features (use-cases) are the key

-   Since _features_ are what users are after, why don't we develop application in a _feature-driven way_?
    -   When you think about it, it makes little sense to do it any differently.
        -   **Using _features_ as our main motivation allows us to be _strategic_ rather than _tactical_ in our approach**.
            -   The _strategy_ does not have to be very forward-looking.

### Code-first approaches to software development

-   There are few notable _code-first_ approaches.
    -   The **tactical programming** is where you mostly brute-force the solution. You code until "it works" and then you stop.
        -   Sadly, you are not done. Your code has bugs and is hard to extend. This violates the goals of software engineering.
    -   The **imperative programming**. The _API-first_, _database-first_ or the _UI-first_ approaches.
    -   Since you want to tackle the **essential complexity** (and reduce the amount of _accidental complexity_), starting with the _API-first_ design is the most optimal solution.
        -   You describe the _contract_ between the UI and the backend applications.
            -   Having said that, all of those have flaws. **You ought to design the domain contracts first** rather than anything else.
            -   By only looking at the API, you **have limited view about the _behaviors_ of the endpoints**.
                -   You can also miss subtle connections in this "idealized" view.
