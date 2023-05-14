# FEM You Do Not Need That Library

Notes based on [this course](https://frontendmasters.com/workshops/pure-javascript/).

## Introduction

- The overuse of libraries is not a new problem. There is so much to learn, that is it impossible to be aware of all the APIs that exist.

  - Having said that, we can at least cover the most common scenarios where we think we need a library/framework, but in reality, we do not.

- The push to use _vanilla JS_ was there even 12 years ago!

- **All the frameworks you are using sit on top of _vanilla JS_**.

- You **do not have to use _vanilla JS_ for everything**. In fact, Max does not recommend it. Frameworks have their place and time.

## DOM

- The in-memory structure that represents the HTML. Contains the data about the nodes and their properties.

  - The DOM **has an API**. The API exposes functions to manipulate the DOM.

- `document` global object refers to the DOM.

  - **EVERY node in the `document` has an API**. Even the comments (TIL).

### Static vs Live collections

- The `getElementsXX` **return a so-called _live collection_**.

- The `querySelectorAll` **returns a so-called _static collection_**.

- The difference between those collections is **how they behave if a mutation occurs on the node that is the part of the collection**.

  - For _static_ collections, the mutation will NOT propagate to the list.

  - For _live_ collections, the mutation WILL propagate to the list.

  - This means that, **for _live_ collections**, the list might change as you are iterating over it!

### "Releasing the thread" AKA yielding to the browser

- Understand that, **as long as you are performing work, the user will not see what you are doing on the screen**.

- This means that **you can create multiple HTML elements and append them one-by-one into the DOM**. The user will not see the "steps" of the operations – they will see the end results, so an HTML with all the elements you have added.

  - This **is how the FLIP animating technique is able to measure and reverse the measurement before running the animation**. Imagine this behavior was not the case. Then, the user would see each step of the FLIP – that would be horrible!

- As soon as you stop executing JS, the browser will update. Not sooner.

### The DOM is NOT the same as HTML

- You can have HTML **without the `body`, but you will see the `body` in the devtools**. Why is that?

  - The reason is that the DOM is a in-memory structure managed by the browser. It does not 100% correlate to the HTML. In this case, most of the browsers will add the `body` node for you.

- When **using devtools you are INSPECTING THE DOM, not the HTML**.

### Scripting and the DOM

- First, understand that, **by default parsing and executing the scripts halts the other processes in the browser**. Those include parsing the HTML and CSS.

  - That is why, in the early days, you were advised to put the scripts at the bottom of the `body` tag, after all the HTML. This ensured that all the HTML was already parsed.

  - Nowadays, we mostly should lean on the `defer` or `async` attributes for the `script` tag. What is the difference?

    - The **`defer` tells the browser to download the script in the background and execute it if the HTML and CSS is parsed**.

    - The **`async` tells the browser to download the script in the backend, and once it is downloaded, execute it, even if there is still HTML and CSS to parse**.

      - If you are not sure which one to use, prefer the `defer` attribute as it is less disruptive.

- The **DOM might not be fully ready when your script executes**. To ensure everything is ready, use the `DOMContentLoaded` event.

  - This event is not the same as `load` event. The `DOMContentLoaded` happens BEFORE rendering the page.

### Event Binding

- There are a lot of events that you can "bind" handlers to.

- There are two ways to do it.

  - The newer `addEventListener`.

  - The older `onxx` setter.

  - One difference is that **you can attach multiple handlers for the same event via the `addEventListener`, but you cannot do it via the `onxx` setter**.

    - Since `onxx` is a setter, it will override the previous handler.

  - Another difference is that **the `addEventListener` supports the _advanced options_ parameter where you can control the phase in which the event fires and so on**.

    - Remember about the `once` parameter as well. It might come in handy!

- You can dispatch your own custom events. Of course, keep them documented and strongly typed. Otherwise it will get out of hand (it always does).

Finished part 3 00:00
