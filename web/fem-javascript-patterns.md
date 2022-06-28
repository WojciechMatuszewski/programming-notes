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

<!-- Finished part 4 29:00 -->
