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

## ESModules

- Use them. They are supported in every major browser.

  - To specify the script as an `ESModule`, use the `type="module"` attribute.

- Before ESModules, everything that you write, every variable and function was global. There was no encapsulation (of course that does not apply to variables defined in functions).

- One thing to note. **When importing files, you have to specify the filename with the extension**. That is not the case if you are using bundlers, but it is a requirement when using native ESModules.

## Navigating via DOM

- You can either have all the pages in the DOM and toggle their visibility, or remove/add elements into the DOM when clicking on the link. Up to you.

  - There is no concept of a "router" in the vanilla JS. But we have **the history API**.

- No matter what approach you choose, you will need to `preventDefault` the `a` tag to ensure the browser does not initiate new navigation with full page reload.

### The history API

- The `history` API allows you to change the URL of the browser without changing anything on the page.

  - It is up to you to react to the URL changes. The event you are interested in is called `popstate`.

    - To be able to navigate when `popstate` fires, you will need to set state when using `pushState`. Otherwise you will not be able to deduce what route the given `popstate` corresponds to.

      - `pushState` with `{route: "/"}` -> `popstate` -> the `event` contains the `route` parameter -> move to the given route.

#### Pseudo components and the page route

- There Max suggest using _Web Components_.

  - In short, **it is your own custom HTML tag**.

    - TIL that `customElements.define` API exists.

    - Custom elements even have a lifecycle methods.

- There is also the **`template` HTML tag**. It is a special tag that exists in the DOM, but the browser ignores it contents.

  - How is this useful? You can **hold different "routes" in the `template` tag. When it is time to display the content, you clone it from the `template`**.

### Shadow DOM

- By default, the CSS applied to the custom element will leak to the "global DOM".

  - By leveraging the _shadow DOM_, you can encapsulate styles. They will only apply to the elements inside the custom components.

- Max loaded the CSS in a very particular way – he used fetch API and injected the result text into the `style` tag.

  - Usually we would be putting the styles in the style tags at the top of the head, but since we want to encapsulate them, we have to put them into the shadow DOM.

## Reactivity using Proxies

- Now that proxies are widely supported, it is a no-brainier to use them!

- Mind the `Reflect` API. Instead of "forwarding" the operation onto the `target`, you can use the `Reflect` API that has the same API as proxy traps (`get`, `set` etc...).

  - Why should you use them? **MDN mentions "but using Reflect saves you from having to remember the syntax that each internal method corresponds to"**.

### Emitting events and changes

- In the workshop we are using custom events to communicate between different components.

  - **We have to explicitly emit those events to the `window` and not the `document`**. The reason is that **each shadow DOM has it's own `document`**.

    - If we were to emit them on the `document` other components would not be able to listen, if the emitting component uses shadow DOM.

## Embedding components in other components

- Since we have the full control over HTML tags (we can creature custom tags), one can render a set of custom tags inside a component, which in itself is a custom tag.

- As for the _props_, Max is using the `dataset` object and adds properties to that object. Then in the child component, we have to parse the property (a string) into a JS object.

## Forms

- Using the frameworks like React and libraries to handle forms caused us to forget that we have quite nice forms API at our disposal.

  - You can get individual fields by the _dot_ notation on the form DOM object.

- **TIL that changing the element value via JS will NOT trigger the `change` event**.

  - This is very handy in our situation as it prevents us from triggering an infinite loop where we change the element value via JS and also listen for the `change` event.

- To **bind the form values with values in JS and vice-versa Max used a Proxy + an event listener on the form fields**.

  - We have also encapsulated the state with a _private class property_ (using the `#` syntax).

## Recap

- It was nice to get my hands dirty with web components. I was not aware that creating custom HTML tags is that simple.

  - To create the elements, one has to register it via the `customElements.define`. Your class/object **must** extend the `HTMLElement` interface.

- Vanilla JS is very powerful. The amount of the surface area can be intimidating, especially for newcomers.

- During the workshop we heavily leveraged proxies. I'm glad they are getting some usage as they are really useful.

- The `history` API is not a black box. I think that if you are using the `react-router` which is partially based on the `history` API (to my best knowledge), it is very beneficial for you to understand the API.

  - The fact that you can "fake" navigation and even provide a state into the `popstate` event is quite amazing!
