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
  yourThirdCallToSetState, // wins
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

No correctly type HOCS you probably should use `utility-types` package, it's great!.
Also, there are great overall guides and cheatsheets

[Guide One](https://github.com/piotrwitek/react-redux-typescript-guide#higher-order-components)
[Second Guide](https://github.com/typescript-cheatsheets/react-typescript-cheatsheet)

### Render Props

Beloved render props.

- you can be using `render` or `children` or any other name for that matter.

### Flux pattern

Nowadays you would probably use context for the problems this pattern is trying to solve but i think it's good to know it either way.

Very simple implementation of flux-like store

```typescript
import { EventEmitter } from "events";
import { users } from "../default-state.json";

export default class UserStore extends EventEmitter {
  users = users;

  createUser = ({ name, email }) => {
    const user = {
      id: Date.now().toString(),
      name,
      email,
    };
    this.users = [...this.users, user];
    this.emit("change", this.users);
  };

  updateUser = (updatedUser) => {
    this.users = users.map((user: any) => {
      return user.id === updatedUser.id ? updatedUser : user;
    });
    this.emit("change", this.users);
  };
}
```

Then you could be using this inside a component like this:

```jsx
React.useEffect(() => {
  const listener = (users) => setState({ users });
  UserStore.on("change", listener);
  return () => void UserStore.removeListener("change", listener);
}, []);
```

### Context API

> Context provides a way to pass data through the component tree without having to pass props down manually at every level.

> All consumers that are descendants of a Provider will re-render whenever the Provider's value prop changes. **The propagation from Provider ot it's descendant consumers is not subject to the shouldComponentUpdate method, so the consumer is updated even when an ancestor component bails out of the update.**

That's why you probably should wrap your `Provider.value` with `useMemo`

```js
const ProviderValue = React.useMemo(
  () => ({
    users: state.users,
    onCreateUser: createUser,
    onUpdateUser: updateUser,
  }),
  [state.users],
);
```

## Redux

### Bind Action Creators

I've actually never used this one. It's used to shorten the `dispatch(action_here)` notation

```js
const createAddAction = (amount) => {
  return { type: "add", payload: { amount } };
};
const dispatchAdd = bindActionCreators(createAddAction, store.dispatch);
// now you can call
dispatchAdd(4);
// instead of
store.dispatch(createAddAction(4));
```

### Normalizing Data

Not much to say here. Use `normalizr`

### Structure

Steve went for `components/containes` structure. It's quite ok actually. He also separates `reducers` and `actions` to their own folders

### Automatically bind action creators

Remember that redux has the ability to automatically bind action creators

```js
const CREATE_CARD = 'CREATE_CARD';

function createCard(cardData, cardId) {
  return {
    type: CREATE_CARD,
    payload{ cardData, cardId}
  }
}

// somewhere else

connect(
  null,
  {createCard} // => great huh ?
)(component)
```

### Performance

Use `reselect`. To learn this properly you probably should write yourn own app. It's similar to Angular (ngrx) selectors.

## "Preserving" the old state

Suppose you are working on a functionality where the user can add and remove some pieces of data. The data could be modified by the user before its removed. Users can add a new item into the data, but as soon as they do, the new item has an "initial" value rather than the value it had when they removed it.

How would you ensure that whenever user adds a new item, that item displays the data it had when the user removed it?

You might be tempted to reach out for another piece of state to track the "deleted" array. Or maybe you encode the state into the URL somehow and use the forward/backwards mechanism.

All of these are valid solutions, but the problem with them is that **they are quite complex**. There exists a quite beautiful solution that does not require that much effort. **Instead of tracking the deleted items, track how many items are displayed at a given time. Keep the items data in separate piece of state you change when the user makes changes**.

```jsx
function Component() {
  const [data, setData] = useState(["initial", "initial", "initial"]);
  const [displayedDataCount, setDisplayedDataCount] = useState(0);

  function onChangeData(index, payload) {
    const newData = [...data];
    newData[index] = payload;
    setData(newData);
  }

  function onAddData() {
    // And for remove we would subtract one here
    setDisplayedDataCount(displayedDataCount + 1);
  }

  // render based on the `displayedDataCount`. NOT the `data` variable
}
```

This solution is pretty great. Yes, we have to have additional piece of state, but that state is only a single primitive value. And there is a risk of that state getting out of sync with the "data" state, but that would also be the risk with any other solution.

## Local and Server state

> Based on [this blog post](https://tkdodo.eu/blog/deriving-client-state-from-server-state).

Sometimes, I find myself wanting to _sync_ the local state with the server state.

```tsx
const { data: users } = useQuery({});

const [selectedUserId, setSelectedUserId] = useState(null);
```

Let's go through a couple of challenges we might have here and how to solve them

### Initialize state from server

How do I initialize the `selectedUserId` to be the first user?

```tsx
const { data: users, isLoading } = useQuery({});

const [selectedUserId, setSelectedUserId] = useState(null);

useEffect(() => {
  if (!isLoading) {
    setSelectedUserId(users[0]);
  }
}, [isLoading]);
```

**Whenever you see a `useEffect` that has `setState` call inside it, your alarm bells should go off**. In some cases, it is necessary, but in most cases, it is not.

```tsx
const { data: users, isLoading } = useQuery({});

const [selectedUserId, setSelectedUserId] = useState(null);
const selectedFromState = users.find((user) => user.id === selectedUserId);
const defaultUserId = users[0]?.id;

const selected = selectedFromState ?? defaultUserId;

return [selected, setSelectedUserId];
```

Notice that I _derived_ the initial state from the server state! No need for the `useEffect`.

The most important, to me, piece of advice to remember: **the local state does not necessarily have to be in-sync with the server state – the _derived state_ must be!**
