# Joy of React

Going through [this course](https://www.joyofreact.com/) and writing what I've learned.

## React Fundamentals

- **React is platform agnostic**. That is why you have to import from `react` and, if you are writing a web app, `react-dom`.

  - There are different _renderers_ for different environments, like React Native or WebGL.

- There is a slight difference between _transpilation_ and _compilation_. While it does not really matter that much, since I'm curious, here it is:

  - The term _transpilation_ refers to going through a high-level language to another high-level language. You can _transpile_ JavaScript to Python.

  - The term _compilation_ refers to the process of transforming the code into a machine-readable code.

### Understanding JSX

- You might noticed that in some projects, there is only a single `import React from 'react'` in the entry file. All other files do not import React, even though they are using JSX. How come?

  - It used to be the case, that, if you used JSX, you had to explicitly import React. After all, the `JSX` are transformed to the `React.createElement` calls.

  - Since a lot of people did not know how JSX works, this "phantom" React import looked weird and caused issues when removed.

  - The React team decided to release a special _JSX transformer_ that would put the necessary imports for you when you use JSX. You no longer have to import React explicitly (of course, if you want to use hooks, you still have to import them from the `react` package).

- Have you ever wondered how `<div>{somethingHere}</div>` works? Mainly, how does the `{somethingHere}` is represented in JS and why we cannot use `if` statements there? **Josh names those _expression slots_**.

  - The `<div>{somethingHere}</div>` would be translated to the following

    ```js
    React.createElement(
      "div", // element
      {}, // props,
      null, // children
      somethingHere
    )
    ```

    Image putting an `if` statement in place of `somethingHere`. That would not work right? You **cannot pass `if` statements as parameters to a function**! Maybe you could wrap it with an IFFIe, and that works.

    ```js
    React.createElement(
      "div",
      {},
      null,
      (() => {
        if(true) {
          return "foo"
        }
      })()
    )
    ```

    So you **can write complex logic inside the JSX** but you have to wrap it with an IFFIe.

- Understanding how the _expression slots_ work, also explains the common issue with whitespace when writing JSX. Consider the following:

  ```jsx
  const element = (
    <div>
      <strong>Days until Santa returns:</strong>
      {daysUntilSantaReturns}
    </div>
  )
  ```

  This piece of JSX gets translated to the following:

  ```js
  React.createElement(
    "div",
    {},
    React.createElement(
      "strong",
      {},
      "Days until Santa returns:"
    ),
    daysUntilSantaReturns
  )
  ```

  Notice that there is no whitespace after `Santa returns:`. **This is causing the output to be "hugged together" without any spacing**.

- JSX and template-based languages like Handlebars are different things.

  - The template-based languages turn what you write into HTML at compile time.

  - The JSX will turn the syntax into JS. That JS is evaluated when the browser runs your application.

### Components

- Components must start with a capital letter. Like hook names, this is enforced by React.

  - This enables React to distinguish between built-in HTML tag, or a custom component.

- Of course, the JSX "invocations" of the components are also transpiled to `React.createElement` calls.

  ```jsx
  <div>
    <SomeComponent name = {"Wojciech"}/>
  </div>

  // is the same as
  React.createElement(
    'div',
    {}, // props
    React.createElement(
      SomeComponent,
      { name: "Wojciech" },
    ),
  )
  ```

- **While the prop name `children` is "reserved", it is not special in any kind from the `createElement` perspective**.

  - Ever wondered what will happen if you specify both the "component children" and the "html children" attributes? **It will be overwritten**.

    ```jsx
    <div children = "foo">some children</div>

    // outputs
    {
      "type": "div",
      "key": null,
      "ref": null,
      "props": {
        "children": "Hello World"
      },
      "_owner": null,
      "_store": {}
    }
    ```

    Notice that the `foo` is completely missing.

### Application structure

- Every React application usually shares a similar layout structure.

  - There is the `index.js` and the `App` component.

  - Do not put a lot of JSX into the `index.js`. It is meant to be a "setup" file.

### Fragments

Have you ever wondered why you cannot return multiple JSX nodes without wrapping them with some kind of element? **It boils down to how JavaScript (yes JavaScript) works**.

Consider the following function:

```js
function greet() {
  return (
    "Hi there"
    "Wojciech"
  )
}
```

This is not valid JavaScript. We cannot return multiple things from the function (that is possible in Go). Now, let change this scenario to JSX.

```jsx
return (
  <div>hi there</div>
  <div>Wojciech</div>
)
```

This is the same as if writing the following:

```jsx
return (
  React.createElement("div", {}, "hi there")
  React.createElement("div", {}, "Wojciech")
)
```

This does not work either. We cannot return multiple things from a function. **That is why you have to wrap multiple JSX nodes with a parent node**.
And if you want to do this **without introducing unnecessary HTML elements**, you should use the `React.Fragment` or `<>` (short version).

### Iteration

#### Mapping Over Data

- Using the `map`, and other functions which output arrays, to produce lists of JSX elements works because **React will "unpack" an array inside the _expression slot_ for us**.

  ```jsx
  return (
    <div>{["hi", "there"]}</div>
  )
  ```

  Imagine using `map` inside the _expression slot_ instead of hardcoding the data. This is how the syntax works. There is no magic.

  ```jsx
  <ul>
    {items.map(item => (
      <li>{item}</li>
    ))}
  </ul>

  // Transforms to
  React.createElement(
    "ul",
    {},
    items.map(item => (
      React.createElement(
        "li",
        {},
        item
      )
    ))
  )
  ```

#### Keys

- React needs the `key` prop to understand how the data changed between invocations of the component.

- You can imagine that the code to update a text (prop changed inside a component) is different than the code necessary to remove and add a new DOM node (marking todo as done and adding another one).

- In some cases, **when you are not deleting or shifting items, the `key` could be an `index`**. While this works, **why would you bother with an inferior method that can break when refactoring**. Instead **opt-in to something unique for a given item**. This is most likely an `id` but also could be a combination of item properties.

- Please note that the `key` is a "reserved keyword" in React. Even if it seems like you are passing it as a prop, the component will not have this prop defined.

  ```jsx
  <Button key = "123">foo</Button>

  const Button = ({key}) => {
    console.log(key) // undefined
  }
  ```

- You can **pass the `key` prop to a `React.Fragment`, but you have to use the "long" form, i.e the `Fragment` component**.

### Conditional Rendering

- While you cannot put JavaScript statements in the _expression slot_, we can pull those statements outside of JSX and use the result in JSX.

  ```jsx
  /*
    This will not work.
    Think about this in terms of the `React.createElement` calls.
  */
  <div>
    {if(foo) {
      return ""
    }}
  </div>

  let result = "baz";
  if (foo) {
    result = "bar";
  }

  /*
    This will work as you would expect.
    Again – think in terms of the `React.createElement` calls.
  */
  <div>
    {result}
  </div>
  ```

- There is also the `&&` syntax which **is an expression rather than a statement**. As such we can use it inside the _expression slots_ directly.

  - Keep in mind the gotcha with `&&` and numbers.

    ```jsx
    // The following will render `0` in the JSX.
    const items = [];
    const numOfItems = items.length;

    <div>
      {numOfItems && <Button>bar</Button>}
    </div>
    ```

    The above will **render `0`. Why is that**? It is because **the `&&` does not return true of false, it returns either left or right side of the expression**. Check out the [mdn documentation](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Operators/Logical_AND).

  - As a rule of thumb: **always use a ternary operator instead of `&&`**.

    - Keep in mind that the **ternary operator will short-circuit branches depending on the initial condition**.

      ```js
      console.log("foo")
        ? console.log("bar")
        : console.log("other")
      ```

      In the above example, `bar` is never logged as this part of the code is never executed. **The `&&` operator works in a similar fashion**.

### Range Utility

- Very useful for creating arrays. Think when you have a "rating" and you need to create N stars to display the rating visually.

- There is no native function to do this.

- What works for me is the `Array.from({length: X}, mapFN)` syntax.

  - There is also the `Array(X).fill()`.

- Josh recommends creating a special utility function.

- **Keys do not have to be globally or "component" unique**. They **have to be unique within a given scope**. [See the React docs section about keys](https://react.dev/learn/rendering-lists#rules-of-keys).

  ```jsx
  return (
    <div className="grid">
      {rows.map((rowIndex) => {
        return (
          // It's okay, the `rowIndex` is unique within this array.
          <div class="row" key={rowIndex}>
            {cols.map((colIndex) => {
          // It's okay, the `colIndex` is unique within this array.
              return <div className="cell" key={colIndex} />;
            })}
          </div>
        );
      })}
    </div>
  );
  ```

### Styling in React

- Josh recommends using CSS Modules at the beginning.

  - I completely agree. With the rise of RCSs, the CSS-in-JS libraries are a bit problematic, due to how they propagate the theme and all.

    - There are also other solutions that work similar to CSS-in-JS libraries but do not require a runtime component. Those seem to be the most widely adopted these days.

## Working With State

### Event Handlers

- Using the `onX` callbacks instead of attaching the listeners directly via `addEventListener` gives us a few benefits.

  - React is able to optimize the event handlers and can do batching (especially important for state updates).

  - React will automatically clean the event handlers. If we go the `addEventListener` route we have to do it ourselves.

### The `useState` hook

- Each time react "re-renders" the component, it will take that component output and compare with the previous output.

  - If the outputs differ, they React will commit changes to the DOM.

  - This means that **"re-rendering" does not mean updating the DOM**. It means comparing the outputs of the components.

  - The **process of comparing the "snapshots" is called _reconciliation_**.

- The state updates are async. This is to optimize performance. This allows **for batching of state updates to occur**.

### Forms

- The _controlled_ vs. _uncontrolled_ components.

- React dispatches so-called **_synthetic events_**.

  - These events are special objects created by React to ensure compatibility across different browsers.

  - You can still access the "native event" via the `.nativeEvent` property.

- When working with `select` tag, check out the `optgroup` which allows you to group options.

  - This pattern can be useful when also having a default "select a value" option.

    ```html
    <select>
      <option value = "">Select a value</option>
      <optgroup label = "Pets">
        <option value = "dog">Dog</option>
        <option value = "cat">cat</option>
      </optgroup>
    </select>
    ```

- **You might not need state for forms**. If you do not have to display the values from the inputs somewhere else, you could get away with only listening to `onSubmit` callback.

  ```jsx
    <form onSubmit = {event => {
      event.preventDefault();
      console.log(event.currentTarget.elements['name'].value) // the value of the name field.
    }}>
      <label htmlFor = "name">Name</label>
      <input type ="text" id = "name" name = "name"/>
    </form>
  ```

- **Consider NOT disabling the buttons**. There are [various problems with disabling buttons](https://adamsilver.io/blog/the-problem-with-disabled-buttons-and-what-to-do-instead/).

### Props vs. State

- Use props to "funnel" the state around the application.

### Complex state

- Never mutate the React state, even if it appears to be working.

  ```js
  colors[index] = "foo"; // mutation
  setColors([...colors])
  ```

  This might work, but it does not guarantee that this will always work. It is better to flip the operation order, like so.

  ```js
  const newColors = [...colors];
  newColors[index] = "foo";
  setColors(newCoors);
  ```

### Dynamic Key Generation

- In some cases, you will have to generate a random ID for the data you are iterating over.

- The **best solution would be to create some kind of random identifier when you are going to append a new item into the array** (assuming you start with an empty array).

  - There are various libraries that could help here, like `uuid` or `ulidx`.

  - There is also **native way to generate UUIDs, via the `crypto.randomUUID` function**.

- You could also generate the key **based on some unique properties of a given item**, though this might be slippery slope. While the data looks a certain way today, will it look the same way in the future?

### Component Instances

Have you ever wondered, how it is possible to use `useState` at all? Like the value resides?

Because if you think about it, when React invokes our component, it also invokes the `useState` with it's initializer. Would not that mean that each re-render we would "revert back" to the state value that was initially defined? Why is that not the case?

```jsx
function Component() {
  const [state, setState] = useState(0)

  return (
    <button onClick = {() => setState(prev => prev+1)}>click</button>
  )
}
```

Clicking the button will cause React to re-invoke the `Component`. As such, we should start with the `state` 0, but the `state` will be 1.

**The answer to this question lies to a concept called _component instance_**. It is the **_component instance_ that is a long-lived object managed by React that actually holds the state value**. React **will create the _component instance_ when the component is rendered for the first time**. It is the **rendering that creates the instances**.

If you **remove the component from the output, the _component instance_ is destroyed**. You can see this in action whenever you toggle some component that then holds the state. If you toggle it on-off-on the state is lost!

This concept is critical to understanding how the state works. I'm amazed that after spending so many years with React, I did not have this intuition.

## React Hooks

- Some hooks, like `useId` or `useState` are coupled to the _component instance_.

### `useId` Hook

- Generates unique id for every component instance.

- **The id is the same on the server (when SSRing) and on the client (during hydration)**.

  - This is a huge benefit of this hook. **This is not the case for 3rd party libraries like `uuid`**.

  - The **returned value is _stable_. It will not change**.

### Rules of Hooks

- React relies on the _order in which the hooks are called_. The order has to be the same between different invocations of the component.

  - This sounds bad to me. Relying on the order of the code may points to a bad internal design.

- In most cases, you will violate the rules by trying to render hooks conditionally.

  - You can refactor the code to pull the hooks closer to where they are used. Most likely to a separate component.

  - You can move the hooks before the condition. It is okay to use them even though they might be used in a component which is rendered conditionally.

### Refs

- Unlike state, one **should mutate the output of `useRef`**.

- You might be wondering why the `useRef` returns an object with a `current` property. **This is so that one could "escape" closures** or rather always have a fresh value within a given closure.

  - When passing the `ref` as a parameter, we pass the object as a reference. Any mutation of the `current` property will be reflected in the captured reference.

  - [You can read more about it here](https://github.com/WojciechMatuszewski/programming-notes/blob/master/web/react/react.md#why-does-useref-have-the-current-property).

- React will not re-render when the ref changes.

### Side Effects

- **Mind the _strict mode_**! In _strict mode_ the `useEffect` will fire twice.

  - Keep in mind that _strict mode_ is not re-creating components. It will re-run functions multiple times without changing the _component instance_.

    - This is why it is not really possible to simulate it. You would need access to React internals.

- **Before using `useEffect`** try to put the logic you would normally put there **into the event handler, if that is possible**.

  - Of course, **if there are multiple places where the state changes, you should consider `useEffect`**.

- Before using the `useEffect` + `ref` to focus the element, **consider using the `autofocus` attribute.

  - Sadly, **this attribute will not work for client-rendered apps** as the DOM node has to be there when the page first loads.

- When the `useEffect` runs, first it will **run previously queued cleanup functions**.

  - Of course, on the initial render there is no cleanup function.

- **Ask yourself is the thing you want to put into `useEffect` is really an effect**. Questioning your intuition most likely will result in better design.

### Custom Hooks

- Hooks should obey the same rules as normal functions would – **hooks should be deep with a small interface**.

  - This means declaring as much state, variables and other related stuff inside the hook itself.

  - This makes the hook joy to use as the user of the hook does not have to worry about a lot of stuff.

- You can create a `toggle` hook via `useReducer`.

  ```jsx
  function useToggle(initialValue = false){
    return useReducer(
      (state) => !state,
      initialValue
    )
  }
  ```
