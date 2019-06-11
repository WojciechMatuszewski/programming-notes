# React State Management

## Quirks with `this.setState`

Lets say you have `count` inside your state. What will the following output?

```jsx
this.setState({ count: this.state.count + 1 });
this.setState({ count: this.state.count + 1 });
this.setState({ count: this.state.count + 1 });
console.log(this.state.count);
```

It's going to be **0, this.setState is asynchronous** (it's placed to a different queue than a simple task such as console.log)

Consider following code, what will happen when `increment` gets invoked?

```jsx
class Counter extends Component {
    constructor() {
        ...
    }

    increment() {
        this.setState({count: this.state.count + 1})
        this.setState({count: this.state.count + 1})
        this.setState({count: this.state.count + 1})
    }

    render() {
        return <span>{this.state.count}</span>
    }
}
```

**Ui will show 1**. React is batching calls. It's almost like doing `Object.assign`

```js
Object.assign(
  {},
  yourFirstCallToSetState,
  yourSecondCallToSetState,
  yourThirdCallToSetState // wins
);
```

What will happen if we pass a function to `setState?`

```jsx
...
  increment() {
        this.setState((state) => ({count: state.count + 1}))
        this.setState((state) => ({count: state.count + 1}))
        this.setState((state) => ({count: state.count + 1}))
    }
...
```

Result **will actually be 3**.

> When you pass functions to `this.setState()`, it plays through each of them

## Patterns and anti-patterns with state

- **DO NOT** use `this.state` for derivations of props
- **DO NOT** use `state` for things you are not going to render
- **DO** use defaults for data in `state`

## Getters and setters

Remember doing `get` and `set` in Angular? Well you can also do it here.

```jsx
get someMethod() {
    ...
}
render() {
    return (
        <div>this.someMethod</div>
    )
}
```

## Prop-drilling

Not much to say about it here. Just use context or redux or mobx

## State Architecture Patterns

### Lifting State with the _Container Pattern_

- **Container components** manage state and pass it to presentational components
- **Presentational components** receive actions and pass them back to the container
- **Presentational components** only have a `render()` method or they are stateless functional components

### Higher Order Components

Container factory.

- **Injector**
- **Enhancer**
- **Injector and Enhancer**
