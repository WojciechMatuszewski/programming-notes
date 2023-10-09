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
