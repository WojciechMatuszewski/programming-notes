# FEM UI Hard Parts

[Notes from this course](https://frontendmasters.com/workshops/hard-parts-ui/).

## Introduction

- All the work related to the UI appears deceivingly simple. How hard can it be to update something after an interaction?

  - Turns out there is much more to that than meets the eye.

- People underestimate HTML. It is quite intuitive and ver forgiving language.

  - The **translates to a collection of nodes represented by an object. That object is then feed to the rendering engine**.

    - This is **the origin of the _Document Object Model_**. You have the object that represents the HTML Document.

### HTML, C++ and DOM

- Browsers operate in C++ or C.

- The HTML translates to DOM which is then internally a C++/C (browser) data structure.

  - It is the **layout and rendering engine that translates DOM onto the screen**.

- The **layout engine** works out page layout for specific browser and screen.

- The **rendering engine** produces the composite "image" for graphics card.

## CSS (briefly)

- It turns out that, like in the case of HTML, **CSS also has a "model" called CSSOM**.

  - The _CSSOM_ also contains information on what is displayed on the page.

    - The _CSSOM_ contains styling rules that the rendering and layout engine take into the account.

    - The **rules in _CSSOM_ are mapped to a node in the DOM**.

## Enabling user interaction

- For any change, there must be some underlying piece of data that changes.

  - Pixels DO NOT equal data.

- Since we cannot "run" the DOM that holds the data, we have to user other means.

  - This is where JavaScript comes in.

### Data and code in JavaScript

- Browsers only load HTML. It is **the HTML parser that spins up the JS engine**.

- When JS starts up, the callstack is empty. As soon as you invoke a function, it will be added to that callstack.

- **Some of the APIs you use in JavaScript, like `console.log` are NOT JS native features**.

  - In this particular case, the `console.log` is a browser feature. When invoked, it does not do much in JS, all the work is done in the browser.

- As eluded earlier, it is the JS that can interact with the DOM. **You can interact with the DOM via the `document` object**.

- The **WebIDL is a format that denotes how JS interacts with the browser API features**.

### Going from C++ land to JavaScript

- When you invoke functions like `querySelector` JS will reach out to the browser APIs. The browser holds all the elements in memory as C++ objects.

  - How come we can get the element back, if it is declared in C++?

  - It is the **Web IDL** that helps to do the translation from C++ land to JS land.

    - Keep in mind that the returned object still has a "hidden link" to the C++ land.

- If you think about it, **logging the output of the `querySelector` prints the _command_, the _HTML syntax_ rather than the data itself**.

  - This is quite fascinating. I cannot think of a similar situation.

### User Actions

- By assigning to certain properties on the object returned from running the `querySelector` you can tell the DOM to _call back_ a function when user interacts with an element like an input.

- The **_Event API_ will trigger an event to a given _DOM node_**.

  - This node then puts **the _callback_ into the _callback queue_**. The **_callback queue_ is a macrotask queue**.

    - This model also showcases the power of the closure. If you think about it, when you assign a function to a given property of the node, you are packaging everything that a given function has to have, in terms of memory, to run successfully. Quite fascinating.

  - The **event loop will pick up the function from the _callback queue_** and run it.

- It takes a lot of work to perform this seemingly simple task. All the data and information that has to travel from the JS to C++ through the Web IDL.

## One-way Data Binding

Finished Part 4 day 1 00:00
