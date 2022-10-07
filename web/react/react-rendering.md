# React rendering

This document contains tidbits of my knowledge about rendering behavior in React. It does not focus on HOW things work but on the "I did not know that" side of things.

## It's all just a function

The `<Component/>` syntax is invoking the given function underneath. There is no magic here. Even the _class components_ could be considered a function (especially if you use the ES6 class keyword).

All in all, the rendering is a process of calling `React.createElement` API multiple times. Let us take a look at an example.

```jsx
import React from "react";

function Component() {
  const Tag = "a";

  return <Tag>link</Tag>;
}
```

Will the example code snippet crash? Nope. It's a valid JSX syntax. The `<Tag>` will end up being transpiled to `React.createElement('a', ...)` call. **In fact, this is how, more or less, styled components render your elements!**.

If you are curious about how _styled-components_ work under the hood, [here is a great resource on that](https://www.joshwcomeau.com/react/demystifying-styled-components/).

### Behavior with hooks

If the `<Component/>` syntax means calling a function, does that mean I can also do `Component()` and call it a day? What happens in such situations?

Let us look at an example.

```jsx
import React from "react";

function Component() {
  return <div>Hello</div>;
}

export default function App() {
  const output = Component();

  return output;
}
```

The output will be "Hello". This means that **in theory, we can 'render' our components by directly calling the _component functions_**. As with most things in tech, this type of "rendering" has a big drawback, which might lead to subtle (or not very subtle) bugs.

I'm referring to how hooks _bind_ to a given component instance and how this process breaks down whenever we call the _component function_ directly. Let us look at an example where calling a React _function component_ leads to an application crash.

```jsx
import { Fragment, useReducer, useState } from "react";

function Component() {
  useState(0);
  return <div>Hello</div>;
}

export default function App() {
  const [shouldRender, toggle] = useReducer((a) => !a, true);

  return (
    <Fragment>
      {shouldRender ? Component() : null}
      <button onClick={toggle}>Click me</button>
    </Fragment>
  );
}
```

The first render works as expected. The button and the `Hello` string are displayed. But as soon as we click the "Click me" button, things break with a `Rendered fewer hooks than expected. This may be caused by an accidental early return statement` error. As I eluded earlier, the behavior and the error we observe have to do with how hooks _bind_ (or are managed) by a given component. **If you call the React function component directly, all the hooks within that component are managed by the parent component**.

In our case, it's the `App` component that keeps track of the hooks calls. As we know, the number of times hooks are called in each render has to be a constant - it cannot change (this is why you cannot call hooks conditionally).

It might not be obvious, but in our example, we are doing just that - calling a hook conditionally. From the `App` component point of view, the `useState` hook within the `Component` is called conditionally since it's the `App` that manages that hook.

```jsx
import { Fragment, useReducer, useState } from "react";

function Component() {
  useState(0);
  return <div>Hello</div>;
}

export default function App() {
  const [shouldRender, toggle] = useReducer((a) => !a, true);

  return (
    <Fragment>
      {/* Calling `useState` conditionally because it's the `App` that manages the hooks inside the `Component` function.
      This is because we are calling the `Component` react function directly, instead of using a JSX syntax. */}
      {shouldRender ? Component() : null}
      <button onClick={toggle}>click</button>
    </Fragment>
  );
}
```

To fix the issue, use the `<Component/>` syntax to render the component. This will allow the React to bind the `useState` within the `Component` to it's instance.

For more reading, please refer to [this blog post by Kent C. Dodds](https://kentcdodds.com/blog/dont-call-a-react-function-component).

## My component renders multiple times. Help

So you heard about this thing called `React.Strict` and that you should use it to see if your application is _concurrent features_ ready. But now, you have noticed that some of your components re-render twice seemingly at random, **during development** (if this happens when running a production built there is either a bug in React or your application has issues).

This behavior that you are noticing is intentional and, in fact, is the way your application is tested to see if it's ready for _concurrent features_.

This stems from the fact that, React might _suspend_ rendering of your component and re-initialize it in the future. If you have things that cause _side effects_ within your component, and they are not wrapped with `useEffect`, you will have problems when the _concurrent features_ are live.

So to flush those cases as early as possible, the `React.Strict` simulates the behavior of _time slicing_ without employing any magic that might be behind this feature.

So if you evert wondered why your components are called at random during the development after you have added the `React.Strict` flag, now you know.

## Render bail-out

Knowing when React will bail-out of committing your component to the DOM is very handy in any optimization / performance-related work.

For React to deem your component "unchanged" (React does not need to commit the result of invoking your component to the DOM), the component must meet the following criteria.

- Your component is already mounted.

- If your component consumes the React context, there were no context changes.

- Your component itself did not schedule an update.

- Your component props did not change.

As you can see, the list is not short, but as long as you try your best to pass primitive props, avoid oversharing the React context, you should be good on that front.

If you wish to learn more about the bail-out criteria and how React renders your components in depth, I suggest you read [this great article](https://www.zhenghao.io/posts/react-rerender). It explains the rendering mechanics much better than I ever could.

### What "render" means

When reading articles about React you will often see the word "render" used in many different contexts. Some people think about the "render" as the act of changing the DOM (incorrect). Some think of "render" as invoking the _render_ method (or the function component), which is the correct way of thinking about that term.

So, keep in mind that the **word "render" does not actually mean changing the DOM, that part is called the "commit"**. The word "render" refers to a React phase where React extracts the output of your components. Since invoking a JS function is usually much less expensive than committing stuff to a DOM, this will most likely be a fast operation.

The takeaway is **not to fear the amount of "renders", instead focus on SLOW "renders" instead**. If the computation in your component takes a couple of seconds, no wonder your application feels slow.

> Learn more about different [phases here](https://blog.isquaredsoftware.com/2020/05/blogged-answers-a-mostly-complete-guide-to-react-rendering-behavior/#render-and-commit-phases).

### Rendering with React 18

The React 18 introduces _interruptible renders_. This allows React to **pause the "render" work and yield back to the browser, making the "render" asynchronous**. Having said that, the **"commit" part (so changing the DOM) is still synchronous**.

Also, take note that **one could also schedule "renders" in the "commit" phase**. These are **synchronous renders** as you most likely manipulated the DOM directly. If React would keep the asynchronous nature of the "render" here, you might end up with a flash of weird-looking DOM as your changes were not fully applied yet.

To schedule a synchronous "render" during the "commit" phase, use the `useLayoutEffect` hook.

> TODO: <https://blog.isquaredsoftware.com/2020/05/blogged-answers-a-mostly-complete-guide-to-react-rendering-behavior/#improving-rendering-performance>
