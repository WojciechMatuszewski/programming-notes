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
