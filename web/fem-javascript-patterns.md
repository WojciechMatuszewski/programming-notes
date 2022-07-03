# JavaScript Patterns FEM

## Workshop

[Here is the link](https://frontendmasters.com/workshops/javascript-patterns/)

## Design patterns

### Module Pattern

While it is a common practice now, it did not use to be the case that we split our JavaScripet files into multiple _modules_. From a time perspective, this practice is relatively new.

#### Using a bundler

A modern bundler, like webpack, will wrap the code in your files with IFFIE. That is the OG way of providing encapsulation.

#### In HTML

Apart from having the bundler wrap your code into IFFIE, you could add the `type="module"` to the `script` tag within the HTML.

#### In Node

You can either **add the `"type": "module"`** to your `package.json` or **use the `.mjs`** extension.

One important thing – **it seems like you have to specify the file extension in the import path**. [According to MDN](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Statements/import), this is to distinguish between `node_module` modules and your locally defined modules.

### Singleton pattern

We are in luck, as by default, the **browser treats modules as non-modifiable singletons**. That being said, there are a few things you might want to consider when creating a singleton.

1. Consider **freezing** the singleton to avoid "out-of-scope mutations". Here, the `Object.freeze` will do – **remember that the `Object.freeze` is shallow!**.

2. To my best knowledge, you can create singletons via objects of classes. Pick your poison.

3. Since modules are singletons by default, you most likely **do not need to use classes or objects** to create singletons (though it might be a good idea to do so, to be more explicit).

```js
// counter.js

let count = 0;

export function add(x) {
  count += x;
  return count;
}
```

```js
// counter.js

let count = 0;
export const counter = Object.freeze({
  add: function (x) {
    count += x;
    return count;
  },
});
```

```js
// counter.js

export const counter = {
  count: 0
  add: function (x) {
    count += x;
    return count;
  },
};
```

If I were to use `Object.freeze` in the last code snippet, the `count` would always be zero, no matter how many times you call the `add` function. Remember that the "frozen" properties cannot change!

### Proxy pattern

The `Proxy` acts as a middleman between you and a given object. With it, one can know whether the code accesses or changes a given object's property. The `Proxy` is widely used in React as a rendering optimization technique, where libraries only re-render the components that use a given piece of the state.

```js
const person = {
  name: "John Doe",
  age: 42,
  email: "john@doe.com",
  country: "Canada",
};

const personProxy = new Proxy(person, {
  get: (target, prop) => {
    console.log(`The value of ${prop} is ${target[prop]}`);
    return Reflect.get(target, prop);
  },
  set: (target, prop, value) => {
    console.log(`Changed ${prop} from ${target[prop]} to ${value}`);
    return Reflect.set(target, prop, value);
  },
});
```

#### About the `Reflect` API

You might be wondering why I have used `Reflect.set` and `Reflect.get` as opposed to `target[prop] = value` and `target[prop]`.

The **main** reason is convenience. In this example, we might as well use the "old" way of handling the traps (`set` and `get`), but if I were to define more of them, the `Reflect` API has methods whose names align with the traps, so it is easier to use.

There any many other differences (the one I like the most is the fact that the `Reflect` API is not throwing errors when `defineProperty` fails). You can read more about the usefulness of the [`Reflect` API here](https://stackoverflow.com/a/25585988).

### Observer pattern

If you have ever added a listener to a button or any other HTML element, you used the _observer pattern_ to do so! The basic premise is to have a list of subscribers. Those subscribers get notified each time the observable gets notified about something.

```js
const observers = [];

export const Observable = Object.freeze({
  subscribe: (subscriber) => {
    observers.push(subscriber);

    return () => observers.filter((sub) => sub !== subscriber);
  },

  notify: (data) => {
    for (let observer of observers) {
      observer(data);
    }
  },
});
```

If you squint, you might also see the _fan out_ pattern here, very reminiscent of the _pub-sub_ pattern implemented by services like AWS SNS. **Keep in mind that the _pub-sub_ pattern is NOT the same as the observer pattern**.

The **main difference** between the pub-sub pattern, and the observer pattern is the fact that **in the observer pattern, the subscriber is strongly coupled with the publisher** – that is **not the case in the pub-sub pattern** where the **publisher pushes messages to a TOPIC that then pushes the message to subscribers**. The topic (or _event channel_ depending on who you ask) is the only thing the subscribers are aware of.

### Factory pattern

A function that returns you a brand new object. The factory pattern is a way to DRY the code you are working with.

```js
const userFactory = () => ({
  // user object properties
});
```

Creating **a considerable amount of objects this way might not be memory efficient IF your objects have functions embedded inside them**. Use the prototype and references as much as possible for efficiency and performance.

### Prototype pattern

The factory patterns usually work well, but, as I mentioned earlier, it is not memory efficient – the functions embedded in the objects are duplicated.

Instead of duplicating the functions and increasing the memory footprint, we can use the `prototype` and inheritance semantics – this is where the `.prototype` or `class` keywords come in handy.

```js
class Dog {
  constructor() {
    // These properties are scoped to a given instance of the class
    this.name = "Name";
    this.breed = "Breed";
  }

  // Everything outside of the `constructor` function is defined on the prototype
  // These functions are SHARED (in the prototype sense) between the instances of the class
  bark() {
    console.log("Woof!");
  }
}
```

#### Inheritance vs composition

There is this old saying that **one should favor composition instead of inheritance**. I stand by this rule, and I believe we should not abuse the `extends` keyword (in the context of classes).

## React patterns

### Container/Presentational components pattern

The idea is to separate the _data_ from the _view_. The **container component fetches the data, while the presentational component is solely responsible for rendering the data**.

I'm not a massive fan of this pattern, especially since the addition of hooks, where we can abstract the container component to a data-fetching hook (most likely with the help of a data-fetching library).

With all that being said, **if you are using class-based components, this pattern is handy**. Since there is no other way to hook into the "lifecycle" of the component than creating a method on the class, fetching the data in a separate component makes a lot of sense.

```js
function Container() {
  const [data, setData] = useState(null);
  useEffect(() => {
    // fetch data
  }, []);

  return <Presentational data={data} />;
}
```

### HOC pattern

The HOC pattern is an attempt at utilizing _composition_ for adding additional behavior to the "root" component. Here is what the HOC pattern looks like from the consumer perspective.

```jsx
function Button() {}

const buttonWithStyles = withStyles(Button);
```

While applying a single HOC usually does not negatively influence the "flow" of the code, in my experience, this is an exception. Usually, developers apply multiple HOC components to a given component, **making it hard to understand where the props are coming from relative to the "root" component**.

```jsx
function Button() {}

const MyNewButton = withData(withStyles(withAuthentication(Button)));
```

It is **easy to accidentally override props when applying multiple HOC components**. Luckily for us, hooks solve this problem by pushing the composition to the body of the component, where you can see what property is coming from which hook.

My advice is: do not use HOC components unless you need to.

> From a TypeScript perspective, HOC components are perfect for training your TypeScript skills. It is not easy to correctly add typings for a HOC.

### Render props pattern

The way HOC passes down properties to the components is implicit and might lead to collisions if multiple HOC components are applied to a given component. To solve this problem, we use the _render props_ pattern.

```js
function Sizer({ children }) {
  return children({ height: 400, width: 300 });
}

// Usage
function Comp() {
  return (
    <Sizer>
      {({ height, width }) => {
        return <div style={{ width, height }} />;
      }}
    </Sizer>
  );
}
```

You can pass the properties using the `children` prop (as in the last snippet) or have specific `renderX` named properties. It is up to you.

Overall, the _render props_ pattern is one of my favorite. Sometimes, you might see people saying developers could replace his pattern with hooks. I cannot disagree more – the _render props_ pattern provides much more than composition. It allows for a concept called "slots", where certain pieces of JSX are rendered in certain places inside a given component.

```js
function Layout({ renderBody, renderAside, renderFooter }) {
  return (
    <main>
      {renderAside()}
      {renderBody()}
      {renderFooter()}
    </main>
  );
}
```

### Hooks pattern

Hooks allow you to express multiple lifecycle hooks in a singular function. So they can completely replace the _container_ component from the container/presentational component pattern.

Make sure you keep the "rules of hooks" in mind.

### Provider pattern

If you ever need to share values (they might be "reactive") with multiple components across your app (the components have to be in a parent-child hierarchy), the `React.Context` API is the thing you will most likely use.

```js
const StateContext = React.createContext(null);
const StateProvider = ({ children }) => {
  const [state, setState] = React.useState(1);

  return (
    <StateContext.Provider value={state}>{children}</StateContext.Provider>
  );
};

const Comp = () => {
  const state = React.useContext(StateContext);
  // logic ...
};
```

Many articles discuss the context API and how to use it effectively. The most important thing you must remember is that the context API allows you to stop drilling props through multiple child components, as would most likely be the case with the "container" component.

### Compound pattern

The compound components pattern showcases the power of the composition and the `context` API. You have one "root" component that manages the state and the children components that read from that state. The combination of the "root" and the children's components compounds into a given UI behavior.

```js
import React from "react";
import { FlyOut } from "./FlyOut";

export default function SearchInput() {
  return (
    <FlyOut>
      <FlyOut.Input placeholder="Enter an address, city, or ZIP code" />
      <FlyOut.List>
        <FlyOut.ListItem value="San Francisco, CA">
          San Francisco, CA
        </FlyOut.ListItem>
        <FlyOut.ListItem value="Seattle, WA">Seattle, WA</FlyOut.ListItem>
        <FlyOut.ListItem value="Austin, TX">Austin, TX</FlyOut.ListItem>
        <FlyOut.ListItem value="Miami, FL">Miami, FL</FlyOut.ListItem>
        <FlyOut.ListItem value="Boulder, CO">Boulder, CO</FlyOut.ListItem>
      </FlyOut.List>
    </FlyOut>
  );
}
```

In this case, the `FlyOut` is the "root". The rest of the components are the child components grouped together. The grouping is a convention, you do not have to group components this way, but such grouping makes it very explicit that these should be used together.

**Notice that you can change the order in which the children of the `Flyout` render in any way you desire**. This is the **most significant benefit over static props**. My rule of thumb is the following: **always favor compound components over static props (especially arrays)**.

## Performance patterns

### Bundling, compiling, minifying, tree-shaking

All of these are pretty standard nowadays. One interesting point is that SWC and esbuild are bundlers, minifies, and compilers in one. If you remember webpack and its configuration, it used to be the case that you had to configure tools for each of these functions separately.

### Static Import

This is your traditional way of importing modules via the `import foo from 'bar'` statements. The bundler might or might not _tree-shake_ the dead code (the effectiveness varies).

### Dynamic Import

This is the pattern where you utilize the `import('...')` syntax. Using the dynamic import syntax, you signal to the bundler that you **do not** want to include the code from that module into the bundle initially. Instead, a separate bundle is created.

Remember that the return value is a `Promise` that you need to resolve to get the module's contents. In React, you would most likely use `Suspense` here.

### Import on Visibility

The _dynamic import_ pattern allows you to control **when** you push the bundle to the user. You might do it when the user interacts with a given element or scrolls to a specific element on the page.

A word of advice: use a library to handle the `IntersectionObserver` API. It is easy to make mistakes and not detach the observer when the element changes, leading to memory leaks.

### Route-base Splitting

Instead of having a single big bundle that encompasses the code for the whole application, you might want to split the code on a per-route basis.

An application view is a good separation point for code. Sometimes we only use specific libraries on a given page, so there is no need to include that code on every page.

Use the _dynamic import_ pattern alongside a router library.

### Browser hints

You can add various keywords and attributes to the `script` tag to control how the browser loads the resource. One of the more popular ones is `prefetch` which tells the browser to `prefetch` the resource if there is a bandwidth for it, and `preload`, which is a more aggressive version of `prefetch` where the browser will always `preload` the resource.

Start on [MDN](https://developer.mozilla.org/en-US/docs/Glossary/Prefetch) and go from there.

## Rendering patterns

### Client-Side Rendering

All the rendering happens in the browser. Until the JavaScript is loaded, you only see the "app shell", which in most cases is not meaningful (from my experience, a "loading" view).

Hydration might take a long time, and if your bundle is big (and let us not kid ourselves – it is), the page will not be interactive for a relatively long period.

The **most significant** tradeoff with rendering on the client is the negative effect on SEO. Your page might rank lower than pages rendered with meaningful content from the start.

### Static Rendering

Here, we are singing the pendulum to the other side – trying to generate as much meaningful HTML as possible at build time. A huge benefit to the SEO but might be cumbersome for deployment – imagine generating many pages during build time (it can take a while).

The **main issue** with static rendering is that the content might get out of date quickly. If that is the case, if you do not have a way to regenerate the content quickly (see the next technique), you will have to **re-deploy the whole website**.

### Incremental Static Regeneration (ISR)

What if you want to pre-render only some of your pages, for example, the homepage, and then render the HTML for others on demand?
This is where the ISR comes in handy. Instead of rendering all the pages at build time, you can do that only for the subset of pages.
The rest of the pages will be rendered dynamically when a user requests the page.

If you decide on using the ISR, you should also consider adding correct cache headers alongside the HTML website payload – you would not want to generate the HTML **every time** a given user visits your page.

### Server-Side Rendering (SSR)

What if we must personalize the HTML you generate for a given page for a particular user? In such a scenario, you cannot create that HTML at build time because you do not have enough context to do it properly.

Using SSR means generating a **personalized HTML** when a given user requests the page. This might take a while, but the response contains meaningful HTML, so your page will have an illusion of loading faster.

### Streaming SSR

Imagine that you could respond quickly with bits of HTML instead of waiting for the full personalized HTML to be generated – this is how the streaming SSR works.

This architecture is the backbone of React Server-Side components and will most likely be the go-to rendering pattern moving forward.

## More resources

- [`patterns.dev`](https://www.patterns.dev/)
