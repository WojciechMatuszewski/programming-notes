# TypeScript Design Patterns And SOLID Principles

Taking notes while watching [this course](https://www.udemy.com/course/design-patterns-using-typescript/).

## SOLID Design Principles

- The **SOLID principles are not about concrete implementations**.

  - They are more like **rules you might want to follow**.

The most important thing – **be pragmatic**. Use these as a _guiding star_ rather than hard, unmovable rules that you can't break.

### Single Responsibility Principle (SRP)

> A class should have only one reason to change

- The **goal here is to isolate code as much as possible**.

  - The more responsibilities a given piece of code has (_class_ or a _function_), the more complex it is.

    - The more complex something is, the harder it is for someone to work with it.

- **While splitting the code, be mindful of _cohesion_**.

  - The more granular the functions, the bigger the risk of lower _cohesion_.

    - **It is possible to have granular functions that preserve high _cohesion_**. As long as you keep related code together (same file or module), you should be fine.

Here is an example of a _class_ mixing two responsibilities – handling blog post CRUD and formatting the blog post into HTML.

```ts
class BlogPost {
  constructor(
    private title: string,
    private content: string,
  ) {}

  createPost() {}

  updatePost() {}

  displayHTML() {
    return `<h1>${this.title}<h1>`;
  }
}
```

The `displayHTML` method does not fit into that class. We ought to create a separate class (or a function) that deals with HTML.

- SRP has **multiple benefits**.

  - The code becomes **easier to maintain** since a change is less likely to impact other modules of your code.

  - The code becomes **easier to understand** since you are concerned only with one thing when looking at a given class.

  - The code becomes **easier to test** since the module/class is simpler.

  - The code becomes **less coupled** since the SRP forces you to leverage dependencies to achieve functionality rather than embedding code inside.

  - The code becomes **easier to reuse** since the modules/classes are more granular.

- SRP is not free of **drawbacks**.

  - You will end up with **a lot more classes/modules**. This could make **searching for particular piece of code harder and make the _cohesion_ lower**.

  - You will **run a risk of over-engineering**. It is very **tempting to take code-reuse and granularity too far**.

  - You **might end up with _shallow interfaces_**. _Shallow interfaces_ are problematic as they increase the overall complexity of the code **from the caller perspective**.

    - There is a fine balance to achieve between having granular, single-purpose functions, and _wide interfaces_.

### Open Closed Principle (OCP)

> Software entities should be open for extension, but closed for modification.

- Two things to unwrap here.

  - **Open for extension** means you can add new behavior (feature).

  - **Closed for modification** means that once the entity **functionality** has been written, **to modify the entity you ought to extend it, not modify it**.

    - This **does not mean you will never modify a class**. All it means that once you have written the implementation, you ought to re-use it to create new features.

Here is the code that violates this principle.

```ts
class Discount {
  giveDiscount(customerType: "premium" | "regular"): number {
    if (customerType === "premium") {
      return 20;
    }

    return 10;
  }
}
```

**Notice that to add a new `customerType` I would need to modify the `Discount` class**. This violates the OCP. I should be _extending_ this class rather than modifying it!

Now, here is an example of _extending_ the `Discount` class to add a new `customerType`.

```ts
interface Customer {
  getDiscount(): number;
}

class RegularCustomer implements Customer {
  constructor() {}

  getDiscount() {
    return 10;
  }
}

function giveDiscount(customer: Customer) {
  return customer.getDiscount();
}
```

**Notice the usage of an `interface`**. It is the `interface` that allows us to dive into _abstract_ and enable _extension_ rather than _modification_. Now I can add as many different customers as I want – the `giveDiscount` will not change!

- OCP has **number of benefits**.

  - You **reduce the risk of bugs** as **extending is much easier than modifying**.

    - When you extend something, you are starting "fresh". You do not have to worry about existing implementation.

  - The code is **easier to maintain**. Again, since you do not have to modify existing code, but extend it, making changes is much easier.

  - The **code written to adhere to the OCP is easier to test**. It turns out, modeling your code in that way also makes it easier to test!

- As with everything, the OCP **is not free from drawbacks**.

  - **OCP introduces a layer of indirection**. This **could increase the overall code complexity**.

  - If you are **rigid in following this principle, you might over-engineer your code**.

    - When using OCP, you might fall into the trap of thinking that you _really_ can't modify the existing code. If that is the case, your classes should be designed only once. **That is not true**.

  - **Using OCP might requite a shift in thinking for some developers**. If your peers are not accustomed to writing code in this quite specific way, you should consider educating them.

### Liskov Substitution Principle (LSP)

The _Liskov Substitution Principle_ is all about **having the ability to use the _child class_ in a place where _parent class_ is required**.

```ts
abstract class PaymentProcessor {
  processPayment(amount: number): void;
}

class CreditCardProcessor {
  constructor() {
    super();
  }

  processPayment(amount: number) {
    // code
  }
}

function executePayment(processor: PaymentProcessor) {
  // code
}
```

Notice that I can call the `executePayment` with `CreditCardProcessor`. I could also create different kinds of "processors".

**In the `executePayment` I do not have to check which instance the `processor` is – I should not care about that**. If you have to check the instances, you are not following LSP.

This rule might _appear_ to be similar to the _OCP_, and there are some similarities, but they pertain to different things.

Remember, the _OCP_ is about _extension_, the _LSP_ is about _interchangeability_.

You can follow one rule, but violate another.

#### Following _OCP_ but violating _LSP_

```ts
interface Person {}

class Boss implements Person {
  constructor() {}

  doBossStuff() {}
}

class Peon implements Person {
  constructor() {}

  doPeonStuff() {}
}

function doStuff(person: Person) {
  // I'm checking for the instance.
  // There is no _interchangeability_ between `Person`, `Boss` and `Peon`
  if (person instanceof Boss) {
    person.doBossStuff();
  }

  if (person instanceof Peon) {
    person.doPeonStuff();
  }
}
```

In the example above, it would be better to have `doStuff` on the `Person` interface. This way, each _subclass_ of `Person` could implement the same function.

- LSP has **number of benefits**.

  - It greatly increases the flexibility of your code.

    - Since your functions take _interfaces_ or _abstract classes_ as parameters, your functions are more generic and _deeper_.

- LSP is **not free of drawbacks**.

  - You are introducing an indirection to the code. This means the code will be more complex.

  - Applying LSP requires careful design. **You might end up with more code that necessary**.

    - **Be very careful**. You are at risk of over-engineering a solution!

#### Following _LSP_ but violating _OCP_

```
class Android implements Platform {
    @Override String name() { return "Android"; }
    @Override String version() { return "7.1"; }
}
class HumanReadablePlatformSerializer implements PlatformSerializer {
    String toJson(Platform platform) {
        return "{ "
                + "\"name\": \"" + platform.name() + "\","
                + "\"version\": \"" + platform.version() + "\","
                + "\"most-popular\": " + isMostPopular(platform) + ","
                + "}"
    }

    boolean isMostPopular(Platform platform) {
        return (platform instanceof Android)
    }
}
```

Notice that, if the platform popularity changes, we would need to update the `isMostPopular` method. This violates the _OCP_ rule.
