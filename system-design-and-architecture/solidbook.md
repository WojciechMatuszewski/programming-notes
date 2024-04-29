# SolidBook

> Notes from [this book](https://solidbook.io/).

Page 483

## Introduction

- Khalil mentions that the goal of the code is to _"create products for customers"_.
  - Agree with this statement, but I would also add _"to solve problems"_.
  - Keep in mind that **the code always has two groups of customers – the users and the developers**.
    - One cannot forget about developers. If you do, you will create a _big ball of mud_ that is tough to with.

## Up to Speed

- In our trade, we have to consider _complexity_ as the "enemy".
  - Think how easy it is to make a mess in your room. What about those dishes in the sink? Now, given how inherently
    complex the code is, think how easy it is to go overboard.
- **Complexity is anything that makes the system hard to understand and modify**.

  - It creeps up on you. It is **incremental**. If you ignore it, you might be fine for a while, but you WILL reach
    a "breaking point", where maintaining a project is no longer possible.

- Khalil describes two sides of _complexity_.

  - There is the **essential complexity**. This one stems from the problem itself. Most of the problems we are solving
    are quite complex if you think about them long enough.
  - There is the **accidental complexity**. This one is **purely driven by developers**. You increase the accidental
    complexity by suboptimal design decisions or subpar code quality.
    - You cannot possibly get rid of it completely. But, as a good software engineer, you can reduce this type of
      complexity to a minimum.

- Here are a few things that contribute to _accidental complexity_.

  - **Getting the requirements wrong**: as part of your job, you ought to know what you really need to do to solve the
    problem. If what you think is the problem differs from what users think is the problem, you are in trouble.
  - **Coupling**: If the dependencies between functions resemble spaghetti, you will not be productive in the codebase.
    **You introduce tight coupling by leveraging concrete instead of "virtual", like _interfaces_**.
  - **Cohesion** is a measure of _how related components are within a particular module_. If the cohesion is low, you
    will have to spend more time understanding what the module does. Inversely, if the cohesion is high, you will have
    an easy time grasping what the code does.

- Khalil proposes a few strategies to mitigate the complexity.

  - **Strategic programming**: instead of brute forcing your way through the feature, you have both your fellow
    developers AND customers in mind. This lessens the change of making a mess.
  - **There is a place and time for _tactical programming_ as well!** If you are building a prototype, or perhaps
    exploring the routes you can take to solve the problem, the _tactical programming_ is a viable technique. **The
    problem arises when you default to _tactical programming_ and never clean up your mess**.

- Code is **also a vital way for communication**. Yes, the language is quite "foreign", but your colleagues are speaking
  it!
  - Picture yourself touring a country you do not know the main language of. How do you feel? Insecure? Lost at times?
    **This is exactly how developers feel if your code is a mess**. It is, like you are speaking a "foreign", hard to
    understand language.

## Software Craftsmanship

- The _software craftsmanship_ is **professionalism** in the field of software development.

  - You do your job well, on time and produce zero defects.
  - You think pragmatically about problems. You do not let dogmatic ideas rule your work.

- While you might not work on "mission-critical" systems, think of the people who are.

  - If the software you wrote for the MRI machine has a bug, people might get killed or injured – this actually
    happened. See the `Therac-25` machine.

- The "top-down" approach of waterfall resembles factory conditions.

  - It has the "I'm smarter than you, so do what I say" vibes which leads to hostility and "us vs. them" mindset.
  - Overall, not a good idea.

- Lots of people are familiar with the _agile manifesto_, but do you know there is also the _software craftsmanship
  manifesto_?

  - Consult [this wikipedia page](https://en.wikipedia.org/wiki/Software_craftsmanship) to learn more.

- Khalil mentions an essential fact: **your employer is NOT responsible for ensuring you grow professionally – you
  are**.

  - I cannot stress this enough. As a developer, you HAVE TO learn new things, constantly, to improve your craft.

- The following really resonated with me: **_when you know the domain, the code becomes a metaphor for what happens in
  real life_**.

  - And guess what – if you have a good DLS in place, it is easier, for EVERYONE, to grasp what is happening in the
    code.

- The world is not all sunshine and rainbows – most people are not aware of the concept of _software craftsmanship_. So
  the question becomes: _how do I get the buy-in I need to practice its principles_?
  - Since _money rules the world_, you have to showcase how practicing TDD or XP **makes or saves money in the long term**.

## A 5000ft View of Software Design

- To be a good _architect_, you have to understand the _low-level details_ of the system.

  - You **cannot be a good engineer without understanding the _high-level_ details**
  - If you do not understand them, you will turn into the _tactical tornado_ programmer. Not a good place to be.

- **Understanding the _high-level_ details of the applications will enable you to cater to customers better**.

- Khalil shows a "map" of _software design_. There is a lot to learn!

- Keep in mind that, **any additions to your code, even when using the "best" _design pattern_, could increase the complexity of the code**.

  - There are no "free lunches" in software. Everything has a price.

- The software of today, at the core architectural level, does not differ that much from the software from X years ago.
  - The tools have changed, but the same patterns still apply.

This chapter was a high-level overview of what's to come, so I did not take that many notes. We will be touching on everything included in this chapter in detail later on.

## Part 2: Humans & Code

- It is vital to remember that the application has **two types of customers: the users and your fellow developers**.

  - Applying the principles of XP and software design, helps us cater to both groups.

- By running various pools and researching what "experts" have to say, Khalil noticed that the **clean code is as much about the humans who read the code as it is about the structure of the code**.

  - That makes sense, since **most of our time is spent reading code, than writing code**.

- A good place to start is to **establish a _coding standard_**.

  - Think code formatting, structure, naming things and other stuff.
  - **_Coding standard_ avoids bike shedding**. You all agreed on X, so let us stick to X.
    - Of course, there is additional benefit of consistency. A very vital aspect of software. Humans like patterns.

- The following are key concepts pertaining to _clean code_ as noted by researching "experts" opinions and various pools.

  - Runs all the tests.
  - Contains no duplication.
  - Maximizes clarity.
  - Has fewer elements.

    They all make sense. They also could be applied to things outside of software (especially the _has fewer elements_), for example, to things like cars.

- Writing code is the constant tug-of-war between _structure_ and _developer experience_.

  - While working on problems, you decided on a given _structure_. This decision, usually, has huge implications on other developers.
    - The more complex the _structure_ is, the harder it will be to contribute to the project.

- Depending on the company you work for, _developer experience_ might be an essential aspect to the company surviving.
  - For others, it might not be that critical (but, I would argue, it should). **How important the DX is, reflects on the culture inside the company**.

## Human-Centered Design For Developers

- Just like there are the _software design principles_, there is a list of _human-centered design principles_.
  - Why should you care? Well, in most cases, you work with other people right?
    - Being mindful about your colleagues will help you achieve a good balance between purist approach (where you only are concerned with structure) and the DX approach (where you are only concerned about DX).

> HDC (Human-Centered Design) is a design philosophy that puts the users' needs, behavior, characteristics, pain points, and motivations first.

- The **crucial goal of HDC** is to **optimize for _discoverability_ and _understanding_**

  - Think about the last time you had to onboard to a new team. To be productive as soon as possible, you had to learn how the code works and how it is structured.
    - If we make that process as easy and as fast as possible, we get ourselves a big benefit that yields dividends over-time.

- When performing any action (like adding a new feature to the codebase), we go through several stages that get us from "How do I do this" to "Did it work?"

  - A lot is happening in between those states. The faster we can go through them, the better the structure and design of our code.

- The **affordance** is about _familiar_ properties of an object, that allows you to perform the desired action WITHOUT learning all about that object.

  - For example, doors contain handles. Handles are for pushing or pulling, therefore, there is a good chance you will open the door by pushing or pulling.
  - This **matters in code**. You want **to use concepts that are familiar and well established**. This is very dependent on the language you use.

- The **signifiers** are **marks or sound or other things that can communicate intended behavior to a person**.

  - Think of having a "PUSH" or "PULL" sticker on the door.
  - They go hand-in-hand with affordances.
    - Design patterns.
    - Comments.
    - Tests.
    - Good variable names.

- **Constraints limit set of possible actions we can take**.

  - This is **vital and powerful technique in software development**.
  - **Constrain what humans need to know about to use the class or module effectively**
    - Make the _interface_ thin, and the implementation deep.
    - Also known as creating a _minimal interface_.

- The **_mapping_ is about the relationship between two sets of things**.

  - In code, that means grouping items similar to each other by proximity. This is also known as _cohesion_.
  - The DOM APIs are a good example – the `.getAttribute` is defined on the _element_, it also defines all other similar methods.

- There will always be tension between the _structure_ AND _developer experience_.
  - In some cases, putting more emphasis on _structure_ makes sense. In enterprise, since the domain is quite complex, using _structure_ to "tame" that complexity is warranted.
  - On the other hand, creating a tool for other developers to use might require more emphasis on DX.
  - It is **your job** to recognize and decide the level of balance between those two forces.

## Organizing things

- Symptoms of a _poor project structure_ include the following (non-exhaustive list):

  - Forgetting what the system does.
  - It is hard to locate features.
  - You spent a lot of time flipping back and forth between files (low cohesion).

- Some of the common project structure approaches include the following:

  - _"Just evolve it over time"_. This one **is a recipe for bike shedding and conflicts if you work in a team**.
    - Your fellow developers have to spend a lot of time thinking: _where should I put this file?_
      - Of course, after some time, they give up and put it somewhere random. **You are most likely guilty of this as well**.
  - _"Package by infrastructure/technology/type"_. Here, **the features are spread across multiple files and the names of the directories does not tell a "story"** of what the project does.
    - While these problems were also present in _"just evolve it over time"_ approach, here they are even more profound.
    - Again, with this approach, it is hard to onboard new people.
  - _"Package by domain"_.
    - We have the "User" concept in our application, as such everything that has to do with users goes there.
      - This **MIGHT give you a basic idea of what the application does when you look at the directory structure**.
    - This approach also MIGHT help you locate features – but you have to understand the top-level domains.
    - There still might be some file-flipping going on, so **this approach is not that great either**.

- So, how should we group project files? **Consider structuring your project against _features_**.

  - Yes, _features_. This is most likely what you are working on right?
    - From the business perspective, **the _feature_ is the unit of work, NOT some kind of controller that you wrote**.
  - If you structure your project this way, you can apply this the following rule: **_if it does not belong to a `feature`, it goes into the `shared` folder_**.
    - This greatly reduces the _burden of choice_ and makes it more obvious where a given file should land.

- Consider the case of **_screaming architecture_**.

  - If you open the application, and it _screams_ "React app" or some other technology, the project is poorly structured.
    - The technology used is **an implementation detail**. The technologies come and go.
  - However, if you open the application, and it _screams_ "fitness app" or some other domain, you did a good job.
    - It will be much easier to work on such application.

- The most important takeaway is this: **group related files together to increase the _cohesion_ of the code**.

  - It might be unrealistic to refactor the code structure in large projects.

- In the **frontend application, features live on pages**. Therefore, consider structuring the project with respect to _pages_.
  - On the **backend, the features _are_ use cases**. Therefore, consider structuring the project with respect to use cases.

## Documentation & Repositories

- Use **tests as the MAIN tool for documentation**.

  - Code comments are okay, but use them sparingly. Only when you need to explain the WHY and not how.
    - The "how question" should be addressed by tests.
  - The test should also explain how a given feature works. Ideally, new developers would be learning about the domain _through tests_.

- **Push as much complexity downwards as possible**.
  - This will, not only help with onboarding, but also with cognitive load while adding new features.

## Naming things

> This whole section is based on [this book](https://www.namingthings.co/).

- Think about all the programming courses you have attended. Did they teach you _how to name things_?
  - Sadly, for most of us, this is not the case, and yet, _naming things_ is **one of the most challenging problems we face as software engineers**.
    - Naming things is not hard – it is the **act of naming things well that is hard**.

### Consistency & uniqueness

> Each concept should be represented by a single, unique name.

- **Given how we learn (through repetition), having consistent names is paramount**.

  - It improves how fast others will be able to grasp as given concept.
  - It improves how well others can search through the codebase.

- **Consistency not only applies to naming**. It applies to everything you have in the codebase.

  - How you structure the application.
  - How you define functions, properties, and classes.

- One of the **worst things you can do is to have similar, but different, names for the same concept**.

  - This confuses everyone and makes them question the whole _convention_.

- **Do NOT recycle** names.
  - This is just plain laziness. There is no excuse to do that.
  - Again, this only creates confusion and makes the code harder to read.

### Understandability

> A name should describe the concept it represents.

- Remember the concept of the _knowledge in the world_ vs. _knowledge in the head_?

  - The _knowledge in the head_ is based on memory and current context.
  - The _knowledge in the world_ is based on culture, something deeply ingrained inside us.
  - **You want to name things based on _knowledge in the world_ as much as possible**.
    - Doing so will reduce the amount of time one has to think about the variable to understand what it does.
    - **You want to name the thing after a real world concept**.

- The best name for a given construct is the name understandable by everyone, not only programmers.

  - Yes, there is a required body of knowledge related to the _business domain_, but that is a given.

- **Avoid overly technical or ambiguous names**.

  - `Processor`, `Manager`, and `Helper` are all pretty vague.

- Either **document the _side effects_ the function does, or push the _side effects_ to a separate function**
  - If the `createUser` also sends a verification email, make sure that fellow developers know that is the case.
    - You would need to rename the `createUser` to `createUserAndSendVerificationEmail`.
      - **Using `and` within the name suggest that the function does too many things!**
    - Or, better, extract the `sendVerificationEmail` to a separate function.

### Specificity

> A name should not be overly vague or overly specific

- **It is okay to start with overly specific name** granted you refactor the name later on.

  - That first step of having overly specific name oftentimes signals to us that we ought to split the variable / function in multiple pieces.

- **Every variable should have _justification_ for its existence**.

  - The fewer variables, the less context we have to keep in our heads.

- **A name should describe what is inside the variable – nothing more, nothing less**.
  - There is no need to get fancy and try to make the variable sound _technical_ or _smart_ (that will only undermine their usefulness).

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
  const tempUserToDetermineIfExistsOrNot = this.userRepo.getUser();
  return !!tempUserToDetermineIfExistsOrNot;
}

// vs.
async function exists(userId) {
  const user = this.userRepo.getUser();
  // "Just right". No need for additional variable.
  return !!user;
}
```

- **Name things by their _roles_ in the application**.
  - Perhaps, instead of using `User` everywhere, we could leverage domain concept names like `Admin`, `Editor` or `Employee`?
    - Of course, this is highly dependent on the application.

### Brevity

> A name should be neither overly short nor overly long.

- Code that is too verbose **is hard to read**. The code that is too succinct is **also hard to read**.

  - There is a constant tension between _compression_ and _context_.

- To explain something, you need to provide enough context.
  - But if you provide too much context, we can lose the "main" takeaway.

> **Communicate _what_ with names, explain _why_ with context**

- The _why_ should be reserved for code comments.

  - The _what_ is reserved for variable names.

- **Some practical** rules to follow:
  - Remove unnecessary words, like `in order to`, `be able to`, `in the event that` and so on. All of these have much shorter equivalents.
  - Stay away from single-letter variables _unless_ they are a convention, like within loops.
  - A good example is the `for (let i; i = 0; i++)` code.
    - The **conventions are _knowledge in the world_**. Keep it consistent!
  - Grouping things **is paramount as it enables you to make the names shorter**.
    - Operations on the `User` object will NOT have to repeat the `User` in their name. This makes them easier to read, but accessing them preserves context -> `User.findOne`.

### Searchability

> A name should be easily found across code, documentation, and other resources.

I can't stress enough how important this is. If you can't find something fast, how are you supposed to make changes to it!?

- Making sure the name is _searchable_ has several benefits

  - It makes the refactoring easier.
  - The name can be located in the documentation.
  - The name can be understood where and how it works throughout the codebase.

- **This characteristic is greatly influenced by other we have talked so far**.

  - If the name is too short, or too generic or perhaps outdated, we will struggle with finding what we want.

- **Practical tips** include:
  - Avoiding _magic numbers_ and creating constants instead.
  - Keeping the documentation up to date (this one is quite hard, is it not?).
  - Ensuring that the names are not too short, nor too long (_context_ vs. _succinctness_)

### Pronounceability

**It is imperative that you can _pronounce_ the name correctly and without effort**.
The name is what you will be referring to when talking with other people.

Some of those people might not be your fellow developers, but other stakeholders.
You want **ALL** groups to understand what you are talking about, not only the people who directly contribute to the code.

- **Some practical** rules to follow:
  - Use abbreviations if they are considered _knowledge in the world_.
    - `ssn`, `vat` and similar are universally understood by anyone.
  - Use _camelCase_ to signal world breaks in variables.
  - And my favorite: **booleans should ask a question or make an assertion**.
    - `potIsEmpty()`, `isPostEmpty()` are good examples.
    - **DO NOT** try to shorten the names unnecessarily. You gain zero leverage by doing so.
  - Instead of `tmStmp` use `timeStamp`.

### Austerity

> A name should not be clever or rely on temporary concepts.

This boils down to keeping it professional. In the comments and with naming the variables. Everyone likes to laugh, but let's keep being snarky outside of the code.

- **Some practical** rules to follow:
  - Avoid being cute, funny or clever. **The easier the name is to understand**, the better.
  - Do not include popular culture references.
    - I would argue that one should not include ANY cultural references.

## Comments

A good comment can be very beneficial.
A bad comment could introduce _uncertainty_ in your understanding of the concept which hinders you from progressing.

**It is hard** to write a good comment. But not all is lost! Given few guidelines, you can increase your chance at writing a good comment.

- Follow the rule: **code explains what and how, comments explain why**.

  - For real, focus on the _why_. When writing a comment, always ask yourself WHY we wrote this code. DO NOT repeat what the code is doing.
  - **Refactor the code to answer the _how_ question** better.

- Comments, in general, contribute to _"visual noise"_ and clutter the code.

  - That is why most people skip writing them. **But skipping comments is not a way to go**.
    - As I mentioned before, always focus on the _why_.

- One of the cardinal sins of writing comments is **repeating what the code does, but in a different language**.

  ```js
  class UserService {
  // This method gets the user by user id. -> A bad comment. Explains what the code does
  	getUserByUserId(userId: string): Promise<User>;
  }
  ```

## Formatting & Style

- Code formatting is _very subjective_.

  - **The most important** for any team is to come up with some kind of convention and stick to it.
  - If you are working solo, do what works for you.

- While _very subjective_ in nature, **there are some universal** rules about formatting that applies to everyone:
- **Use whitespace to create "blocks" of related code** and separate those blocks.

  - Whitespace is free. It is a very powerful tool in making the code more readable.

- Here are **a few practical rules** to follow:

  - _Use obvious spacing rules_. How to apply then is never _explicitly_ taught, but after a while you develop an intuition when one should use them.

    - I'm referring to spacing between _keywords_, _identifiers_, _operators_ and _literals_.

    - Keep the **code "density" low**.

  - Do not have large chunks of code without indentation. **Group related statements together**.

  - Break things on the horizontal axis if they are too long.
    - This should be the job of your code formatter.

- Next, Khalil makes a point about _smaller files_.

  - I'm torn on this one.
    - On the one hand, smaller files are easier to read.
    - On the other, **smaller files often prevent you from seeing the duplication** thus making it hard to create good abstractions.
  - My advice would be to pick what works best for you and stick to it.
    - I'm unsure if being dogmatic about file size is a good strategy to have. Life and programming is almost never "black and white".

### Consistent whitespace

- We touched about _consistency_ before, and this aspect is also quite important in the context of _style and formatting_.
  - Stick to the same casing style (be it _camelCase_ or _snake_case_).
  - **It does not matter what casing you choose**. What matters is that you stick to it.

### Storytelling

- Reading the code is like reading a newspaper. **You want to _front-load_ the most important bits of the code**, and have the less important ones at the bottom.
  - This **principle is called _Newspaper Code Principle_**.

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

- **Do your best to keep the level of abstraction consistent** throughout the function.
  - Do not start at the very high level, utilizing a _deep_ abstraction only to dive into very minute details at the end.
    - Hide those details into another abstraction and use that.
    - This way, the reader knows what to expect.

### Tooling

- **You should NEVER have to think about formatting (apart from vertical whitespace)** in your codebase.
  - Use tools like `Prettier` and `ESLint` to do the job for you.
  - If you do not, use are wasting yours and everyone else's time.

## Types

- There are two main _type systems_ for any given programming language to choose from:

  - The **static type system**. Here, the types are a concrete construct used to describe that will be stored in a variable.
  - The **dynamic type system**. Here, the types are _implied_ and _hidden_ from the programmer.

- One type system is not necessarily better than the other. It all depends on the people you work with and how you like to work.

  - Having said that, **usually, the _statically_ and _"strong"_ typed languages will save you more time in the long run**.
    - They will save you from silly mistakes we all do.
    - But the most important thing is that **statically typed languages allow you to communicate via contracts and not concrete implementations**.
      - This is HUGE for code composition.

- Communicating via _contracts_ means using _interfaces_ or _types_ rather than concrete classes or objects.
  - This way, you can implement the necessary contract differently in tests and differently in the runtime code.
  - While you could do the same with _dynamically typed_ language, the language will not check against that contract.
    - In other worlds, the "machine" will not work for you, you will have to work for the "machine" and ensure the contract is valid.

```ts
interface BookSaver {
  saveBook(): Promise<void>;
}

// TypeScript will check the contract here.
// If you pass something that does not contain `saveBook` method, TypeScript will complain.
async function createNewBook(saver: BookSaver) {
  await saver.saveBook();
}
```

- _Statically typed languages_ are also better at communicating the intent.

  - They often have a keywords related to a given pattern (or family of patterns).
    - Keywords like `abstract` or `private`, used in the right context, _scream_ a given pattern right away.

- Khalil touches on the fact that _dynamically typed_ languages _could_ improve **initial** productivity, but that often comes at a later cost.

  - I completely agree. I would argue that the temptation of doing "magic stuff" with _dynamically typed_ languages is much higher.

- **TypeScript** uses a **_structural type system_**.
  - This means, the type compatibility and equivalence are determined by the computed type's actual structure.

```ts
class Instrument {
  doSomething: () => {};
}

class Guitar extends Instrument {}

class Synth extends Instrument {}

declare function pluckGuitar(guitar: Guitar): void;

pluckGuitar(new Guitar());
pluckGuitar(new Synth()); // Should not be allowed but it is! This is Duck Typing!
```

The rest of this chapter talks about the basics of TypeScript. I've skipped making note for those as I'm pretty comfortable with the language already.

## Errors and exceptions

- **Errors and how we handle them** should not be an afterthought.

  - They are an integral part of any application, because things fail, all the time.

- Khalil mentions **two _common_ error-handling approaches that are pretty bad**.

  - First, you have the "log and return `null`" pattern. It looks like this:

    ```ts
    function CreateUser(email: string, password: string) {
      const isEmailValid = validateEmail(email);
      const isPasswordValid = validatePassword(password);

      if (!isEmailValid) {
        console.log("Invalid email");
        return null;
      }

      if (!isPasswordValid) {
        console.log("Invalid password");
        return null;
      }
    }
    ```

    This does not work well, because **the caller has no idea what failed**. It is quite bad.

  - Another approach one might take would be to "log and throw."
    - Again, not that great since you have to surround everything with `try/catch` blocks.
      - Also, you **have to dig into the implementation of the function to know if it fails or not**.

- There is a **difference between an _error_ and an _exception_**.

  - **An _error_ is something expected**. Something that might happen. Think a duplicate email in the signup form.
    - These are **domain-specific and ought to use domain-specific language**.
  - **An _exception_ is something unexpected**. Think network connection timeout.

- **To handle _errors_, consider the `Either` type**.
  - If you are coding in Rust, that would be the `Result` type.
  - In Go, that is returning a _non-nil_ error from the function.
  - **Whatever you do, return the error rather throwing it**.
    - I completely agree. After playing with Go and Rust, I can't stand how JS (and TS) handle errors.

```ts
class DomainError extends Error {}

// Return this instead of throwing!
class EmailInvalidError extends DomainError {
  constructor(email: string) {
    super("This email is invalid");
  }
}
```

- Ok, so we know how to handle _errors_. What about the _exceptions_ ?
- The key here is to **wrap the code that you do not own with `try/catch` statements and return with a _generic_ error**.
  - This will greatly increase the _expressiveness_ of your code and will give the caller chance to react to a "generic" error.

```ts
async function createUser(email: string) {
  try {
    await db.createUser(email);
  } catch (error) {
    return new ApplicationError(error);
  }
}
```

- The main goal of using _domain errors_ or "generic" application errors (derived from exceptions) is to **make the implicit explicit**.
  - The _implicit_ is not great in programming. You do not want anyone guessing what you had in mind when writing code, so you do not want anyone guessing what the errors could be.

## Features (use-cases) are the key

- Since _features_ are what users are after, why don't we develop application in a _feature-driven way_?
  - When you think about it, it makes little sense to do it any differently.
    - **Using _features_ as our main motivation allows us to be _strategic_ rather than _tactical_ in our approach**.
      - The _strategy_ does not have to be very forward-looking.

### Code-first approaches to software development

- There are few notable _code-first_ approaches.

  - The **tactical programming** is where you mostly brute-force the solution. You code until "it works" and then you stop.
    - Sadly, you are not done. Your code has bugs and is hard to extend. This violates the goals of software engineering.
  - The **imperative programming**. The _API-first_, _database-first_ or the _UI-first_ approaches.
  - Since you want to tackle the **essential complexity** (and reduce the amount of _accidental complexity_), starting with the _API-first_ design is the most optimal solution.
    - You describe the _contract_ between the UI and the backend applications.
      - Having said that, all of those have flaws. **You ought to design the domain contracts first** rather than anything else.
      - By only looking at the API, you **have limited view about the _behaviors_ of the endpoints**.
        - You can also miss subtle connections in this "idealized" view. You are at risk of **missing to uncover the _essential complexity_**.

- **_Code-first_ approaches tend to create "Anemic Domain Models"**.
  - These are models where the _domain objects_ contain little or no business logic.
    - You think you are encapsulating the behavior, but in reality, you put all the burden of data-validation and calculations on the caller.
      - **If you perform a lot of checking / calculations is the _services_ layer, your domain model is _anemic_**.

### A feature/use-case driven philosophy for software design

- First, you have the entropy. The _big idea_, the _knowledge in your head_ how things ought to work.

  - Then, **your job is to translate that into a well-defined _domain model_ DRIVEN by _features_**.
    - You know, the things that the customer expects.

- As an expert, while talking with customers, you need to sniff for _features_ (feature = use case).
- **If you are hyper focused on the _feature_, you will notice when you write pointless code**.
  - Code that does not serve the feature (be it functional or non-functional) is only adding to _accidental complexity_.

### The anatomy of the feature

#### Data

- The following are the **most common types of data**:
  - The **input data** is necessary to _execute_ the feature.
    - You might need to perform validation or bend the data in your logic. That is okay.
  - The **dependencies** are sometimes needed to accommodate features.
    - Think fetching from a network or persisting to a database.
  - **Output data** is critical to get right.
    - Remember that **features can fail**. You ought to **encode the _errors_ and _exceptions_ is an expressive way** (hint: the `Result` type).
  - **Output events** could be used to notify other systems that something happened.
    - Using events **allows you to _chain_ features together but keep the decoupled**.

#### Behavior

- **Side effects** are everywhere. The **key is to push them towards _edge_ of the operation**.
  - This allows you to perform "all or nothing" updates. Very handy if you need to notify the other system via events when operation succeeded.
  - Here, a good pattern to have in mind is the **_Transactional Outbox_** pattern.
    - Think sending events based on changes to a given table, rather than performing a _transaction_ (side effect + event);

#### Name-spacing

- Just like you break your code into smaller pieces, **the _domain_ could also be broken down to smaller _subdomains_**.
  - A _subdomain_ is a smaller piece of the entire problem space.
    - Dealing with smaller pieces of the problem, helps with mental overhead. In addition, such split increases _cohesion_ and _discoverability_.

### The lifecycle of a feature

- Khalil mentions the following _lifecycle of a feature_:
  - Discovery
  - Planning
  - Estimation
  - Architecture
  - Acceptance tests
  - Design & implementation

I do agree with the above, but **keep in mind that all of this happens in a short cycle**.The last thing you want is to introduce a waterfall approach in your team.
I'm not saying that the waterfall approach is bad, but rather that, for the project you work on, it _most likely_ is not the right choice.

- **I like the concept of the _outer_ and the _inner_ feedback loop**.
  - The outer loop could be an end-to-end test validating the whole feature.
  - The inner loop could be a specific unit test validating a small piece of functionality.
    - **When both loops are passing, you introduced just enough _essential complexity_**. Stop right there and refactor!
      - Do not add any more code.

## Planning

- **There is no room for heroism in software development**.

  - This is a marathon, not a race.
  - There is no place for _all-nighters_ to meet the deadlines.
    - Software development is a marathon, not a race.

- In XP, you are in a constant loop of _prioritizing_, _coordinating_ and _recalibrating_.

  - Prioritize – do the most important stories first.

  - Coordinate – keep everyone synced up and in the know.

    - Especially critical in remote-work scenario. It is better to repeat yourself several times, than not communicating enough.

  - Recalibrate – get back on track when we get off track.
    - Things change. Estimates change. **It is very likely that you will get off track at some point**. That is okay!
      - Again, communicate with others.

### How planning works

- First, you _scope_ the project.
  - Here, you are probing for _functional_ and _non-functional_ requirements **at the high level**.
    - You should be roughly able to estimate how much time the work will take, but that estimate might not be accurate.

After scoping the project, the estimate might be quite long. The customer might object the estimate. **That is to be expected**.
This is where you **work with the customer on the _essential_ rather than _nice-to-have_ features**.

What feature would bring the most value? – that is the question to ask at this stage.

- Next, Khalil talks about _measuring velocity_ and how it **allows us to be honest about the amount of work we are able to do**.

  - I like the principle, but I never been to a place where this concept was executed properly.
    - Using _story points_ is a good idea, but in experience, it falls short (we tend to vote the same amount of points as our colleagues did).

- The concept of using _velocity_ as the measure of how much work the team can take on is called _yesterday's weather_.
  - You can get as much done as you did last week, or last day.

## Customers

- **The customer is an integral part of the team**.
  - They might not participate in the _"day-to-day"_ of coding and meetings, but their voice is **essential** in getting the priorities right.

> It is not stakeholder knowledge but developers' ignorance that gets deployed into production

- **Writing software is like driving the car with two people in it**.

  - You have the **driver – the engineers** concerned with the technical and the "manual".
    - They have to operate the wheel, change gears and stop on red.
  - You have the **navigator – the customer**. They **chart the direction of the trip**.
    - The road can change, the directions can change.

- You might think of _the customer_ as a single person.

  - That is not the case – **your customer is the _whole team_** consisting of domain experts, product managers, business analysts and so on.

- In most teams, that customer is the product manager.
  - **Having a _"proxy"_ for the customer is not ideal, but it is better than nothing**.
    - It is like **playing the _telephone game_** where each "gateway" for communication distorts the original message.
      - The less of those "gateways" the better.

To finish up our _"driving a car"_ analogy, consider the following quote:

> If you get lost driving, it is not the car's fault, it is the driver's

**You are the driver**. It is your responsibility to follow the direction, but also take your surroundings into consideration.

## Learning the Domain

- **The code should use _"the common language"_** to describe what it does.

  - The names, and the vocabulary seen throughout the codebase ought to be **dictated by the domain** we are working in.

- **Customers ARE the domain experts**. They inherently understand the **essential complexity**.

  - It is through _them_ we should be discovering features.

- The _model_ in the _domain model_ is an **imprecise, but good enough** description of the problem.

  - It allows **everyone to be on the same page**.
  - A good example of a _model_ is _"the cloud"_.
    - We know what it means, the PM knows what it means (more or less). This shared understanding gives us a common language we can use to communicate with each other.

- **Developing a shared mental model is critical** for any application success.

  - You want to get rid of as many "chains" of communication as possible, so developers ought to directly work with _domain experts_.
    - But, if there is no "common ground" between such communication, developers might have problems inferring the requirements as they usually _think_ differently than the _domain experts_.

- Khalil lists the following benefits of the _shared mental model_:

  1. **More clarity and less time spent on low-value work** – you spent less time going back and fourth trying to understand the requirements.
  2. **Solve problems faster** – you communicate between with your peers.
  3. **Perennial and maintainable** – code translates to real-world. Things in real-world do not change that fast. This makes the code less volatile.

- **One of the processes to create a _shared domain model_ is _event storming_**.

  - You gather everyone interested in the application and ask them how the process _could_ look like.
    - We write the sentences in past-tense as events or commands.
    - At the end, you have a timeline of events/commands with key entities!
  - You might notice, that, in the process, **different team boundaries start to emerge**.
    - This might be a good opportunity to create those teams if they do not exist already!

- **Commands are your features**.

  - Commands will trigger events that then other parts of the system (or other systems) subscribe to.

- **We perform commands _against_ aggregates**.

  - They encapsulate rules and decide whether a particular command should succeed or fail.

- The **subdomains could be coupled together to some degree**.

  - For example, ordering an item would most likely involve billing.
    - This _might_ be solved using a event-driven system.
    - **Creating a "map" of relationships between different subdomains _might_ also be helpful here**.

- You will find **queries especially useful for web applications**.

  - Most of web applications fetch data rather than change it.
  - **To correctly scope any UI work, you ought to have designs + the queries at the ready**.
    - Otherwise, your estimates will be wildly off due to the lack of information.

- **There is an alternative technique to _event storming_ called _event modelling_**.
  - It works on the similar basis to _event storming_, in addition to:
    1. It incorporates rough sketches of the UI into the _timeline_.
    2. It **allows you to create different _swimlanes_ for subdomains**.
    - This is very vital, as some events might flip-and-fourth between different subdomains making it challenging to place them on one timeline.

## Stories

- **It is your job** to help customers identify, estimate and implement stories.

  - The _story_ describes an unit of functionality **without low-level details**.

- **It is the customers' job** to **write the stories**.

  - You should take a backseat and let the _domain experts_ do the work.
    - These should be **high-level descriptions of how things ought to work**, completely devoid of technical details.

- The **story should follow the _INVEST_ rule**.

  - **I**ndependent – this allows you to parallelize development.
  - **N**egotiable – they can change over time.
  - **V**alue – they must provide value to the customer.
    - The non-functional stories are valuable as well!
  - **E**stimable – if the story is too large to estimate, split the story into sever small stories.
  - **S**hort – related to the above.
  - **T**estable – the _magic question_ here is _"How will I know when I've done that?"_.
    - Use _preconditions_ (GIVEN), _actions_ (WHEN), and _post-conditions_ (THEN) to drive tests.

- When writing stories **avoid ambiguous statements** like _quickly_ or _as soon as possible_.
  - Do your best to **be concrete**, this way, **you might be able to bake in non-functional requirements into the story as well!**

> As a day trader, I want to be notified when the price of a stock on my watch list goes up over 5% within **1 to 10 seconds of it occurring** so that I can decide whether I should make a trade or not.

## Estimates & Story Points

- **It is developer's job to estimate how much the story will take**.

  - The customer can arrange the stories by _value_, but they should not take part in estimation.
    - Estimating is a technical task. Do not let others dictate how you will work.
    - **Estimates are NOT commitments, but most people will think of them as commitments**.

- Use _some kind of number_ to estimate how long the story will take.

  - **Those numbers DO NOT represent hours or days**, they are arbitrary **estimation of the effort**.
    - The higher the effort, the longer the story will take, but again, **we do not precisely know how long that will be**.

- Stories are _malleable_. You can merge, or split them as you wish.

  - As long as it makes estimating easier, anything goes.

- **Use a _"spike"_ if you are dealing with a lot of unknowns**.

  - If you have not used a particular library before.
  - If you need to introduce a new technology to get stuff done.

- **The _"spike"_ is a POC. Consider it a _throwaway code_** that you learn from.

  - You can always use parts of the code you wrote in the actual implementation!

- While assigning estimates takes some time, **it is a great way to share knowledge amongst the team**.

  - Consider a scenario where two people disagree on the amount of story points for a given task. **Perhaps one developer knows more about the dependencies of the task than the other?**
    - If that is the case, the team benefits from the knowledge transfer and is able to more accurately estimate the story.

- To **improve your estimates, look at the historical data**.
  - This will require you to measure how long a given story took.
    - There are multiple tools that allow you to do this.

## Release Planning

- **You want to release the most valuable functionality or have the most technical risk associated with them first**.

  - This will not only make your life less stressful, but also make the customers happy.
  - **Always focus on the _essential_**. Be in complexity or features.

- **Not all bugs need fixing right away**, but **the most critical ones ought to be**.

  - For others, **create a story for a bug** and let the customer decide how critical it is.
  - In addition to making the application "work as expected", **fixing a bug involves writing a test for that particular issue**.
    - **Do NOT start fixing a bug without having a failing test**.
      - The time spent writing the test will yield dividends in the future if you need to refactor that bit of code.

- **Build the infrastructure as you go**.
  - This will make it so you build **only the necessary infrastructure**.
    - The bigger the infra layer, the bigger the overhead.
  - Deploying only what you need ties in nicely with the concept of a _walking skeleton_.

## Iteration Planning

- While the _release planning_ is mostly driven by customers, the _iteration planning_ is mostly driven by developers.

  - **Here, we dive deeper into the technical world** of infrastructure and concrete code-related tasks.

- In the places I've worked at so far, this meeting was called _planning_ or _weekly sync_.

  - You most likely already do this in your job as well!

- Consider letting people assign themselves to tasks.

  - This practice is not without issues, as it _might_ create silos of knowledge, but it can also motivate your peers to dive into areas they are not familiar with.
    - It all depends on the people you work with.

- Breaking down tasks into small pieces allows you to shuffle them around between engineers.
  - The smaller the unit of work (without going into extremes), the better.

## Understanding a Story

- **Avoid jumping into technical conclusions**.

  - When asking about a story, **ask more questions than you speak**.
    - Goes without saying, but avoid speaking technical jargon.
      - This rule applies to _every_ discussion you have. Even with other fellow developers.
        - By using technical jargon, you _assume_ that the other person also understands that jargon. That is not always the case!
  - The technical minutia does not matter at this point. **Your job is to uncover the essential technical complexity**.
    - HOW we solve that complexity is up to us.

- Consider using _pseudo-code_ for documenting what you learned while talking with domain experts.

  - This _pseudo-code_ should be understood by everyone that has SOME technical knowledge, but is not necessarily an engineer.
    - **This _pseudo-code_ might be verbose** and quite lengthy – that is okay!
      - It is the _essential complexity_ showing its head. Usually, without writing everything down, you will forget/miss some details.

- **Ask about _what could go wrong_**.

  - This is a great way to discover many edge cases that every application has.
  - It also allows you to uncover _dependencies_ of a story.
    - System A usually will integrate with System B. **You story will most likely span both systems, even if only in a small way**.

- [Here is a template](https://cis.bentley.edu/lwaguespack/CS360_Site/Downloads_files/Use%20Case%20Template%20%28Cockburn%29.pdf) for an _use-case_ task.
  - Notice it contains a lot of sections. **Again, all of this is to uncover _essential complexity_**.

## Acceptance tests

- **Acceptance tests, when passing on production, are THE definition of done**.

  - Do not settle for nothing less. The task is NOT done when the code is "working".
    - How do you know it is "working"? Did you consult with a _domain expert_?
      - Even if you did that, how do you know it will keep on working? By manual checks? **Manual checks are a waste of time!**

- In the early days of XP, we had tools that tried to encapsulate the acceptance tests in a DSL. That DSL would translate to a set of "runnable" tests.

  - While it looks nice on paper, it did not work well.

    - Why would a customer had to learn a specific DLS only to describe how the feature should work like?

  - **To solve this "problem", use BDD – TDD but focused on the _behavior_ rather than technical details**.
    - BDD utilizes the _given_, _when_, and _then_ "sections" **written in a plain english**.
      - This allows customers / domain experts to encapsulate behavior and it translates nicely into tests!

- In an ideal world, the customers would finish writing acceptance tests in the first days of the iteration and then handoff those to developers.
  - In reality, be prepared to write them yourselves.
    - It is not an ideal situation as your knowledge of the problem space might be limited, but it is much better than having no tests at all!

## Programming Paradigms

Programming is complex. It is like talking an arcane language to a computer and getting responses back.

Add to that the fact that programming is _purely virtual_, meaning you can create anything you wish to create, it is only natural that complexity starts to creep in on us.

> Just because the tool lets you do something does not mean you _should_ use it to do that thing.

- **Your job is to LIMIT the number of ways one thing can be done, ideally to a SINGLE way**. This "golden path" allows you to enforce consistency and
  helps everyone to build a _shared mental model_ of the program.

- It **does not matter what _programming paradigm_ you use**. What matters is that you use it as it was intended to be used.
  - Do not be dogmatic either. It is expected of you to be able to use different programming paradigms.
    - **You are not a "some_language programmer". You are a software engineer who solves problems**.

Consider the following quote:

> Testing shows the presence, not the absence, of bugs

We can't be certain that our programs work unless we prove that _mathematically_. Good luck doing that on your React app!

There are multiple _programming paradigms_:

- The **structural programming**.

  - You use _sequences_, _loops_ and _iteration_ to get the job done.
    - Problem 1: **it will require you to either define a global state or pass state between subroutines**.
    - Problem 2: **complexity rises as you use more loops, iterations and sequences**. This is called _cyclomatic complexity_.

- The **object-oriented programming**.

  - Some times though of as the most controversial of them all due to many ways one can go about using it.
  - **The original OO implementations thought of _methods_ on an object as means of handling a message that it understands**.

    - This has huge implications.
      - First, it showcases that **changing methods is not a good idea**. You most likely can encapsulate both calls into one.
      - Second, and what shocked me the most, is that it **goes against the popular teachings about OO**, where objects are "bags" we can put stuff into.
      - Third, **notice that it goes against using inheritance in most cases**.
        - Using inheritance does not create a "web of objects" but rather hierarchical structures. It **creates a hard code dependency** which we do not want!

  - There are many good things OO brings to the table, but arguably, the most important of them is the **_dependency inversion_** technique.

    - This allows you to **_code against the interface_ rather than implementation**.
      - You operate in the world of _abstract_ rather than concrete.
        - This makes your program more flexible, less coupled and more maintainable!

  - Drill this into your mind: **the OO is about creating a _loosely coupled "web" of objects_ that communicate via _methods_**.

- The **functional programming**.

  - It might not seem like it on the surface, but **functional programming has a lot in common with OO**.

    - Instead of _dependency inversion_ you use _parameterization_.
    - Instead of decomposition to smaller objects, you decompose the code into smaller functions.

  - **One massive benefit of functional programming is the REDUCTION of state**.
    - You will, most likely, still have SOME state in your application, but it will be pushed to the very edges.
      - This makes testing and debugging much easier!

- Khalil touches on the _imperative_ vs. _declarative_ programming.

  - The _imperative_ style focuses on "how to do things". You process the code line-by-line.
    - **Code is mostly comprised of statements** like variable declarations and loops.
  - The _declarative_ style focuses on "what to do".
    - **Code is mostly comprised of expressions** like function calls.

- Be pragmatic when picking which _programming paradigm_ you use. Always consider your peers. Perhaps they are unfamiliar with _functional programming_?

  - **Design principles are paradigm-agnostic**. You can implement every design pattern in any code.

## An Object-Oriented Architecture

- The **"double loop of tests (unit and acceptance) is a cheat-code for development**.

  - Imagine how much time you save when you set up the acceptance tests up front.
    - When you are done, you are _really done_. The bugs might still be there, but the chance of introducing regressions is rather small.

- Before you do anything, you have to have a _walking skeleton_ in place – the initial object-oriented architecture.

- A _functional requirements_ answer the question of _"what the application should do"_.

  - It should allow users to upload photos.
  - It should allow users to chat with each other.

- A _non-functional requirements_ answer the question of _"how the application should behave"_.

  - It should have low latency, even when used by thousands of users.
  - It should be SOC2 compliant.

- Here, Khalil touches on **the _principle of last responsible moment_**.

  - This means **delaying any significant decisions until the last possible moment when we have narrowed down our options**.

    - There is inherit **cost of delay you have to factor**.

  - A good example here is a decision to **deploy a _modular monolith_ rather than starting with microservices**.
    - You MIGHT need microservices architecture at some point, but you MOST LIKELY do not need it when you start.

- We want to build _testable_, _flexible_ and _maintainable_ software.

  - **In most cases, using a _layered architecture_ will fit the bill**.

- Arguably, the best known _layered architecture_ pattern is the MVC. **MVC is not without its fair share of problems**.

  - We tend to _overload_ the model with all sorts of logic that should not be there – validation, business logic and so on.
  - We tend to mix "core code" and "infrastructure code".
    - **This makes it quite hard to test the "core code" (the most valuable in our application)** because it is coupled to the "infrastructure code" – code that is not that valuable.

- If you "upgrade" the MVC to add more layers, you will be able to separate the "core code" and "infrastructure code".

  - Treat the application like an onion with multiple layers.
    - The _core domain logic_.
    - The _application features_.
    - The _adapter logic_.
    - The _infrastructure details_.

- **Use _interfaces_ to communicate between layers**.
  - Do your best to stay in "abstract". This allows you to **leverage the _dependency inversion_ to the fullest!**

## Testing Strategies

- There are various _test "levels"_.

  - The unit tests.
  - The integration tests.
  - The end-to-end tests.
  - And probably more...

  I find that people tend to care too much into which category a given test belongs to. Okay, the test might be an _unit_ test. Now how is that information helping us? **If you follow the domain-driven layered architecture, your tests will naturally fall into each category, and it will not matter which one is it**.

- There are patterns to follow for each "category" of the tests.
  - For example, the _Page Object Pattern_ suites end-to-end tests well.

### What should we test and how

- Unit test all the domain layer code.

  - These are your helpers, middleware, _Value Objects_ and _Entities_.
    - As a reminder, a _Value Object_ represents a small, cohesive unit of data like money or street address.
      - Usually, _Value Objects_ are their own separate type (think type alias).
    - The _Entity_ is a mutable object with unique identity, like a _person_.

- Unit test all the use cases.

  - You will most likely need to use mocks / stubs here.
    - Ensure that the mock or stub has been called with the value you expect.
  - As a reminder, the **_use case_ depends on _repositories_ and other _adapters_**.

- Integration tests for adapters (incoming and outgoing)

  - If there is a network request to be made, use tools like _msw_ or _supertest_ to "catch" the requests and return values as if the _use case_ did.

- Write end-to-end tests.
  - **Do not mock anything** if possible.
  - It is **critical** to ensure the end-to-end tests run at the most "production-line" environment as possible.

## The Walking Skeleton

- The _walking skeleton_ is **just enough infrastructure and "code structure" to get you started**.

  - You build it during _iteration zero_, **before you develop any features**.

- It **allows you to validate that all the architectural pieces you have in place work end-to-end**.
  - This is very valuable, as you will be building upon those pieces all the time!
  - In addition, **validating functionality end-to-end early on** helps you uncover any potential issues you overlooked in your initial design.
    - It is not a question of IF this happens but WHEN this happens.

In this section, Khalil walks us through the process of **creating the _walking skeleton_ by using end-to-end BDD test that DRIVES the implementation**. It is quite fascinating seeing an end-to-end test used in very TDD-like scenario.

- **While you can certainly overdo it, adding a layer of indirection is usually very helpful**.
  - Consider the _Page Object_ pattern for end-to-end tests. By "hiding" a lot of code behind a method call, you make your tests less brittle!

## Pair Programming

I only skimmed this chapter since I'm well acquainted with the topic.

The most important thing to understand about _pair programming_ is this: **it will save you a lot of time in the long run** as it drastically shortens the feedback loop – you are coding and having your code reviewed at the same time!

## Test-Driven Development Workflow

- **TDD is a tool that enables you to complete your work faster**.

  - By _complete_ I really do mean _complete_. No more "follow up" tickets.

- **The technique Khalil recommends is _Double Loop TDD_**.

  - It comprises of the "inner loop" – the _unit test_, and the "outer loop" – the _acceptance test_.
  - If both pass, you are done with the feature!

- **TDD requires discipline** to get right.

  - You will be tempted to write more code than necessary. You ought to resist that temptation!

- **TDD is also quite difficult because our code is often not written in a "test-first" mindset**.

  - Ever had to manipulate the `process.env` in your tests? That is a sign of mixed _infrastructure code_ with what you are testing.
    - Use _dependency inversion_ to combat this issue.
  - Ever had to setup a huge test harness to run the _system under test_? That is a sign of _single responsibility principle violation_.
    - Make your code more modular to combat this issue.
      - Perhaps some things should live in different domain?

- **There are _two_ TDD "schools" of though**.

  - The "Chicago" TDD where you test _inside out_.
    - Starting from the very core you build your way _outwards_.
  - The "London" TDD where you test _outside in_.
    - Starting from the very edge you build your way _inwards_.
  - Some people tend to be dogmatic about the approach, but **the true masters of TDD mix the approaches as they see fit**.
    - Once again, just like in life, nothing is purely white or black.
      - Being dogmatic about something is not helpful as it prevents you from learning more about "the other side".

- The **refactoring part of TDD is where the design happens**.
  - When you have a passing test, now it is time to refactor, collapse duplication and so on.
    - If you are disciplined about the process, you would also make any structural and architectural changes in this step.

## Test-Driven Development Basics
