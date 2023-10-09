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
