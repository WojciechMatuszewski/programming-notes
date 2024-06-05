# Clean Code

Watching the "Clean Code" videos and taking notes.

## Episode 08 – Foundations of the SOLID principles

- The _source code_ is a design.

  - Think of the compiler as the "factory" and the code as the "blueprint" that goes into the factory.

    - If you follow this analogy, you will arrive at the source code being the _design_.

- Software is very interesting.

  - Think about a concrete product like a watch. For watches, you want to pay a great deal of attention to the design as correcting mistakes AFTER you have built the watch is very expensive.

  - Software is _malleable_. So, even after "having a finished product", you can still change it!

    - **Software design is quite expensive looking at how much time it takes vs. how much money programmers make**.

      - So, the structure is inverted. **This inversion of costs has a huge implications**.

        - This means it is **very easy to evolve the software into a design that is hard to work with**. Since you do not have a cost incentive to do a lot of design up-front, you ought to do it _progressively_.

### Designs Smells

- **Rigidity**: when the system is hard to change.

  - It is when the cost of making it change is high. Given the inverted cost structure of software, having such system is like shooting yourself in a foot.

  - Here, you have to **manage the dependency between modules**. If one change requires you to "touch" multiple modules, you are working in a rigid system.

- **Fragility**: when a change to one module causes other modules to misbehave.

  - Imagine a scenario where **you fix a bug, but introduce a bug in other part of the application in the process**. This is viewed as **very unprofessional**, and a sign of structural problems.

  - Again, the solution is to manage the dependencies between modules.

- **Immobility**: when you can't extract a seemingly "generic" module, like a login, into another system.

  - Like in previous cases, **this is caused by tight coupling of dependencies within one system**.

- **Viscosity**: when necessary operations, like testing, or adding a new feature, take a long time.

  - If this happens, **developers tend to ignore it**. They get "used to" the state of things.

    - This is bad, as it lets the codebase rot. If not addressed, similar to cancer, the problem might get _terminal_ and evolve the codebase into a _big ball of mud_.

- **Needles complexity**: when your system has a lot of _anticipatory code_. Think about "hooks" or software written for 100X requests when you are dealing with X and, MAYBE, in a couple of months that will double.

  - This usually **occurs when engineers are _afraid_ of the future**. The **are afraid of the possibility of making changes** since they are already hard.

    - The **solution is a robust set of tests**. Tests allow you to _refactor fearlessly_ and that is THE way to refactor!

### Code Rot

Code rots because of **bad decisions** and the lack of knowledge.

- You have made bad decisions when writing the initial version of the program and you never addressed them.

- You are not aware (or ignorant of) the design principles that ought to guide to when writing code.

**You can't let the code rot**. If you ever see "traces" of it happening, address it immediately – that should be your duty as a professional.

Sadly, not everyone feels up to the task. In such cases, developers leave the company/project and let others handle the mess they have created.

I deem this behavior unprofessional and, in a sense, immoral.

### Dependency Inversion

You can prevent _code rot_, by **visualizing what is the "flow" of dependencies in the module**.

You want the **module that contain the high level policy, to use low level policy modules, never in reverse**.

In other worlds, you want **a dependency of a module to "terminate" at a lower-level abstraction**. This abstraction "hides" the concrete implementation for other modules.

The **SOLID principles are about dependency management**. In OO, the _message_ is the dependency.

## Episode 09 – Single Responsibility Principle

- Bob argues that a **responsibility pertains to the "family of functionality" rather than to be a method**.

  - So, if your class only has methods for interacting with a database, it has a _single_ responsibility.

- Think of a _responsibility_ as **a source of change**.

  - These changes will be _requested_ by the **users the _responsibility_ relates to**.

    - **If your class pertains to more than a single _actor_ it has multiple responsibilities**.

> Responsibility is a family of functions that serves one specific _actor_

- We are talking about actors, not users, because users might have different, and multiple, roles within the organization

### Values of Software

- The **primary value** is **malleability of software**.

  - The software is not a static thing. It can change. Take different forms.

    - Notice how this value interplays with the _secondary value_.

- The **secondary value** is **behavior**.

  - The software meets the current requirements of the current user.

    - The needs of the user change!

- The **profitability of software is tied to the primary value**.

  - While the software might not meet all the user demands now, it can change since the _primary value_ is high.

  - The **inverse, where secondary value is high, and the primary is low** is a sign of a software that might fail in the long run.

### Friction

Imagine a scenario where a class contains two _responsibilities_. One related to the database, one related to formatting strings.

It is likely that there are two separate people working on those two different responsibilities. **If they both make changes to the class, they will face merge conflicts and friction**.

We do not want this to happen.

### SRP

- Now that we understand what _responsibility_ is, we can talk about _Single Responsibility Principle_.

> The module should have **only one** reason to change.

This means you should "gather" together things that change together, and separate things that change separately.

- A good example is _logging_.

  - You might interview logging with the code, but that violates the SRP.

  - Instead, **create a separate functions or classes for logging**. Those construct would then take the "base" data you worked on and add the necessary log statements.

- **SRP does not necessarily mean creating a new files for each _responsibility_**.

  - One has to use their judgement here. How _similar_ are the _responsibilities_?

    - If you have code that does X and code that logs what the X does, I would argue those should live within the same file.

    - On the other hand, code that does X and code that does Y should most likely live in separate files.

### Welcome to engineering

- Here, Bob provided us _potential_ solutions for decoupling a class that had three _responsibilities_.

  - **Sadly, none of the proposed solutions was "final"**.

    - All of them **had tradeoffs**.

- This is the world we live in – the world of perpetual tradeoffs. There are no perfect solutions.

Finished 48:29 – case study

## Episode 10

## Episode 11 Part 1 – Liskov Substitute Principle

- It does not matter what a type is. All it matters are the operations that can be performed on the type.

  - The **_type_ is really just a bag of operations**.

  - The structure of type is _hidden_ behind those operations.

- The idea is to allow the program to use _subtype_ of some class as that class.

  - The code should not be concerned about what is the instance of the concrete data passed to the function.

    - Just like we do not care what the _type is_.

### No need for inheritance

- Notice that the LSP does not rely on inheritance.

  - **As long as a given type has the same set of operations as the parent type, we can use these types interchangeably**.

  - In some languages, this might be achieved through inheritance, but in others, we might employ duck-typing.

### Violation

- When LSP is **violated, it often results in runtime error or "refused request"**.

  - The type does not contain the operation which the "parent" type has.

- It could also mean that the subtype method class does something completely different, like throwing an exception.

- Here, Bob walks us through a classic example of having a `Rectangle` class and trying to create `Square` based-off the _Rectangle_ class.

  - **If you have a function that relies on the `height` no changing after calling `rectangle.setWidth` and you pass a `Square` to it, you created so-called "undefined" behavior in the code**.

    - Your code might crash. Your code might keep working. Only by looking at the code you are able to deduce what is going to happen.

  - To "fix" this issue, you will most likely add an `if` statement to the code which looks at `instance` of the parameter.

    - This **violates the OCP** making your program fragile and tightly coupled.

#### Solution

- To fix the "square as rectangle problem", **consider creating completely different classes for both, and keeping them as separate types**.

  - This is the case of "it appears that this might work, but the reality is different".

    - Remember, in some cases, it is better to "duplicate" code than to try to fix incompatible blocks together (by inheritance or duck typing).

### The "Representative rule"

- You might be tempted to think that, since the rectangle _is a_ square, there ought to be some kind of relationship between those two constructs in code.

  - **But the structures in code are not the "real" thing – they just _represent_ the real things**.

    - The **representatives do not share the same relationships as the things they represent**.

- When writing code **remember that you create a _representation_ of a real word, rather than concrete objects**.
