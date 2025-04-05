# Crafting Interpreters

Working through [this CodeCrafters challenge](https://app.codecrafters.io/courses/interpreter) and reading [this book](https://craftinginterpreters.com/).

## Notes

- **Compiling** is an _implementation technique_ that involves translating a source language to some other, usually lower level, form. When you generate _bytecode_ or _machine code_, you are compiling.

  - _Bytecode_ is code for a _language virtual machine_ (VM).

  - _Machine code_ is code for the chip inside your computer.

- When an implementation _is a compiler_, we mean it translates source code to some other form but does not execute it. You will be taking that output and running it yourself.

- When an implementation _is an interpreter_, we mean it takes in source code and executes it immediately. It runs programs "from source".

  - **A good example here would be a REPL environment**.

    - **Note that a compiled language can also have an interpreter**. That is most likely to provide that REPL environment for debugging and such.

- We used **visitor pattern for parsing/printing expressions**.

  - The visitor pattern creates a layer of indirection between _behavior_ and a concrete type. **Think of adding a new behavior for a class without having to modify that class**.

```ts
interface Element {
  accept(visitor: Visitor): void;
}

interface Visitor {
  visitArcher(element: Archer): void;
  visitMage(element: Mage): void;
}

class Archer implements Element {
  accept(visitor: Visitor) {
    // Notice that each element tells the visitor which method to visit.
    // This is called double-dispatch.
    visitor.visitArcher(this);
  }
}

class Mage implements Element {
  accept(visitor: Visitor) {
    visitor.visitMage(this);
  }
}

class Attack implements Visitor {
  visitorArcher(element: Archer) {
    console.log("Archer attacked!");
  }

  visitMage(element: Mage) {
    console.log("Mage attacked!");
  }
}
```

**If I wanted to add new functionality to both the `Archer` and `Mage`, I would NOT have to touch those classes**. All I would need to do would be to implement a new `Visitor`.

I find this pattern interesting, because I usually default to embedding the functionality alongside with the type.

TODO: Learn about "double dispatch"
