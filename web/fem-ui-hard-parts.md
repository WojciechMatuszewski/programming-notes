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

### Going from JavaScript to C++

- When the JS engine boots up, it populates the memory with globally accessible objects that **allow you to reference the DOM – the data in C++**.

  - One of such objects is the **`document` object**. The `document` object allows you to query / manipulate elements.

### User Actions

- By assigning to certain properties on the object returned from running the `querySelector` you can tell the DOM to _call back_ a function when user interacts with an element like an input.

- The **_Event API_ will trigger an event to a given _DOM node_**.

  - This node then puts **the _callback_ into the _callback queue_**. The **_callback queue_ is a macrotask queue**.

    - This model also showcases the power of the closure. If you think about it, when you assign a function to a given property of the node, you are packaging everything that a given function has to have, in terms of memory, to run successfully. Quite fascinating.

  - The **event loop will pick up the function from the _callback queue_** and run it.

- It takes a lot of work to perform this seemingly simple task. All the data and information that has to travel from the JS to C++ through the Web IDL.

## One-way Data Binding

- There is a lot of things that the user can do that could affect the view.

  - It is not feasible to have a separate case for each of those things.

  - That is why **most frameworks derive the view from the state. It is the state that is the single source of truth**.

  - That is why React will re-render **everything** by default. It is so much easier to reason about the UI this way.

  - Of course this only a single approach, but it is very popular.

- The goal here is to declare the relationship between the data and the view **once**.

  - After that is done, the only thing that we have to do is to update the data. The UI should update automatically!

## Virtual DOM

- The virtual DOM is a in-memory representation of all the nodes from the HTML.

  - This allows us to make changes in-memory, and then, when all the changes are made, commit to the "regular" DOM.

    - This is an optimization technique. Applying changes one-by-one to the "regular" DOM is quite slow.

- By comparing the previous and the next virtual DOM, we can deduce which nodes needs to change.

  - This process **requires us to traverse the whole virtual DOM**. This can get quite slow on large trees.

### Block DOM

> Based on [this article](https://millionjs.org/blog/virtual-dom).

- The Block DOM is a **representation of state of the HTML DOM nodes**.

  - It **holds the state of the nodes, not the nodes themselves**.

    - This means that, **to know what changed, we only have to compare the state. This DOES NOT require us to do any kind of traversals**.

      - All of this results in a much more performant solution. But this **is not a silver bullet**. Websites with a lot of dynamic content would not benefit that much from this technique. You can see huge difference in websites with a lot of static content and some reactivity here and there.

### The function as a component

- In the _One-way Data Binding_ section we explored the notion of having a single source of truth of values for a given elements.

  - The model of adding a callback to an input where the callback is the single source of truth breaks down as soon as you want to manipulate (remove/add) elements on the page.

  - To encapsulate even more logic, we will **not** be binding to the events, but rather re-running our function (something that resembles the components) every so often.

    - In the function itself, we will be replacing the "state of the world" (the UI) with the most up-to-date elements, derived from the data in JS.

### Making the code more "visual"

- One of the most appealing facts about the JSX (even the "raw" representation, in it's functional form) is that it _visually_ describes the HTML.

  - Of course, it is not the HTML itself. It is much more powerful. But my being _visual_ it really helps with keeping the right mental model.

- At this point in the course, Will starts to slowly convert to the JSX syntax by representing the "to be created HTML" via arrays, like so

    ```js
      const divInfo = ["div", `Hi, ${name}`]
    ```

  The first item is the node "kind" and the second are the "props", in this particular case, the `children`.

  - The evolution of this approach is the `createVDOM` function

      ```js
      function createVDOM() {
        return [
          ["input", name, handle],
          ["div", `Hello, ${name}`],
        ]
      }
      ```

### Composition & Functional Components

- Will first starts with using the `map` function on the array returned by the `createVDOM` function.

  - This allows for flexibility – we no longer have to manually assign variables to the results of the `convert` function.

  - At this stage we are still running the "update" function at a interval. Not ideal.

- Here we are still re-creating all the elements. This is wasteful. Will eludes to some kind of diffing that would allow us to improve the performance.

- The next evolution is to, instead of hardcoding the elements within the VDOM array, use functions that would return pieces of that array. **This is how component works in React!**.

  - We need to use recursion here. Keep in mind that the function can return an array of arrays. For each array we have to call our "convert" function that will create DOM nodes.

### Performance optimizations

- Will first starts with encapsulating the DOM creation and update in a single function.

  - This allows us to trigger this function **only when something changed**. We are still missing granular updates.

  - We no longer run the "update" function at a given interval. This makes the whole UI faster.

- The next step is to introduce the `findDiff` function. The version Will presents is using the `JSON.stringify` comparison. That makes sense as the whole structure of the VDOM is serializable.

  - I wonder how that relates to RSCs and the JSON-serializable format they are in.

  - Note that the function diffs **array items and not the whole VDOM**. This allows for granularity on node-by-node basis.

## My closing thoughts

I'm a bit conflicted whether I got any value from this workshop. While the content and the explanation was top notch, I believe it was a bit too shallow. On the other hand, it showcased how one might build a simplistic FE framework.

---

## Different models of Reactivity

> Based on [this great article](https://www.builder.io/blog/reactivity-across-frameworks)

- At the very basic level, the reactivity model **can be _fine-grained_ or _coarse-grained_**.

  - With _coarse-grained_ reactivity model, the framework have to re-run the whole parts of the tree to understand what changed.

  - The _fine-grained_ reactivity model is more performant, in a sense that much less diffing operations are required. There is no need to re-run the components for the whole parts of the tree. We exactly know what has changed.

- **Some frameworks like Svelte use a technique called _dirty checking_**. This technique involves **running a "scanner" function that checks if something changed**. This is **different than using listeners (at least from what I read on the internet)**.

  - Other notable framework that uses this technique is Angular.

- **Qwik does the reactivity at the DOM element level, not component level**. This is great since, when a change happens, the framework does not have to re-run the whole component. It can update a given value in the DOM.

  - This **only works with mutations that do not add/remove DOM nodes**. But I bet they are working on improving that. When you add/remove DOM nodes, currently, they have to re-run the component.

- **Solid works on similar basis to Qwik, but the framework also handles removals/additions in course-grained fashion**.

  - Solid executes the component exactly once. This is wild to think about in the world where React will execute your component multiple times.

- After reading the above, you might think that _coarse-grained_ reactivity is bad and should not be used.

  - As with everything in software world, that is not the case. **The main benefit of _coarse-grained_ reactivity model is that it "just" works**. It is very hard to mess things up, when you re-create the state of the world every time something updates.

    - Of course, not everything is sunshine and rainbows. By using _course-grained_ reactivity model, you pay the price with performance.
