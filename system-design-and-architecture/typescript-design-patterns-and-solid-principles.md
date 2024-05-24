# TypeScript Design Patterns And SOLID Principles

Taking notes while watching [this course](https://www.udemy.com/course/design-patterns-using-typescript/).

## SOLID Design Principles

- The **SOLID principles are not about concrete implementations**.

  - They are more like rules you might want to follow.

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

  Finished 31.
