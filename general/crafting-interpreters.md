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
