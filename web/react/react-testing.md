# React testing

## Mocking HTTP requests

There are 2 ways to go about it.

1. Mock the `window.fetch`. This might be a solution for you but I'm not really a fan of this.
2. Use _dependency injection_ (very important topic) and pass mocks. This is great when you have such setups, but sometimes you do not.
3. Use `msw`. This library is great! Not only you can mock stuff within your tests but also work on the APIs that may not have been finished yet.

Recently I'm leaning more towards option 3, so there is a quick showcase of the API

```jsx
import { setupServer } from "msw/node";
import { rest } from "msw";

const server = setupServer(rest.post("url", HANDLER));

test("it works", () => {
  // test stuff based on the HANDLER implementation
});

test("edge case", () => {
  server.use("url", SOME_EDGE_CASE_HANDLER);
  // test

  // reset to globally defined handlers. You can also use `beforeEach` for this.
  server.resetHandlers();
});
```

The API is straightforward, it can work with `REST` and `GraphQL`, pretty nice!

## Dreaded `act` function

So you probably are aware of the `act` function right? Right?

You actually might not be, if you are using `@testing-library/react`, all the `fireEvent` and `userEvent` methods are wrapping your interactions with `act`.

But let's say you are faced with the `act` warning - you know which one I'm talking about 😉

So, _React_ is not performing operations synchronously. There is a _scheduler_ package involved. And this is completely fine. Since a lot of worked can be packed in one frame, you do not really notice the UI being updated incrementally.
But guess what, your tests do!

```js
function App() {
  let [ctr, setCtr] = useState(0);
  useEffect(() => {
    setCtr(1);
  }, []);
  return ctr;
}
```

This simple component will have probably two or more _units of work_. When you write your tests like so:

```js
it("should render 1", () => {
  const el = document.createElement("div");
  ReactDOM.render(<App />, el);
  expect(el.innerHTML).toBe("1"); // this fails!
});
```

That assertion will most likely run in between those two _units of work_. This is what `act` is for. To **make sure that your assertions are run AFTER all of the Reacts _units of work_**. And that's really it.

For `async` things you would use `await act(async () => {})`.

## Generating fake data

You most likely need to supply some kind of data to your components when you are testing them. While you could waste time coming up with some mock data (or use _foo-like_ names which are not that good for testing), you could also use tools like `@jackfranklin/test-data-bot` which will build realistic data for you.

Using realistic data for you tests while not having to come up with the names, numbers etc.. will save you time and make you more confident that the stuff you are building actually works.

A sample of the `@jackfranklin/test-data-bot` API

```js
import { build, fake } from "@jackfranklin/test-data-bot";

const fakeCoords = build({
  fields: {
    latitude: fake((f) => f.address.latitude()),
    longitude: fake((f) => f.address.longitude()),
  },
});
```

## `renderWithProviders` function

This is something that you will eventually end up with given large enough codebase (hopefully you are writing tests!). Kent actually recommends creating `test-utils` file (I dread the `utils` name, for me it would be just `testing` or something like that) which _re-exports_ all methods from `@testing-library/react` and overwritten the `render` method.

So, the file would look something like this:

```js
import { render as rtlRender } from "@testing-library/react";

function render(ui, ...options) {
  const Wrapper = ({ children }) => (
    <MyProvider>
      <MySecondProvider>{children}</MySecondProvider>
    </MyProvider>
  );

  return rtlRender(ui, { wrapper: Wrapper, ...options });
}

export * from "@testing-library/react";
// override React Testing Library's render with our own
export { render };
```

Now all you have to do is to _search and replace_ all the files that are importing the `@testing-library/react` and change the import to the newly created file.
This is a neat solution since you will most like end up with a couple of global _providers_ for you app. These should be treated as _implementation details_ when you are testing component in isolation.

Remember that you can customize the value for given providers by passing parameters to the `render` function. This is why it's important to **create your own custom global providers in a certain way**, a way which enabled dependency injection

```jsx
const AppContext = React.createContext();
AppContext.displayName = "AppContext";

function AppProvider({ children, ...props }) {
  const [state, setState] = React.useState();

  const value = React.useMemo(() => ({ state, setState }), [state]);

  return (
    <AppContext.Provider value={value} {...props}>
      {children}
    </AppContext.Provider>
  );
}
```

Notice the **`...props`**. This will allow you to **overwrite the `value` of the context in your tests**. While you **could** do the same with manually creating the `AppContext.Provider` and supplying there, it just feels wrong for me. We are making all this effort to encapsulate context within this one single _provider_ and that throwing all of this away in tests.

## Cypress debugging

Cypress exposes a couple of ways to debug the problem you are currently facing (or maybe you are just writing code and want to see how things are at a given time).

There are couple of ways:

1. Use the `debugger` inside your code / cypress test. You will probably have to use the `.then` API (this is not a _promise API_) to get the current _subject_.

2. Use the `.debug` command. This is the native way of debugging stuff in Cypress. Pretty neat.

3. Use the `.pause` command. Super useful, especially since Cypress is just so fast sometimes that, at least I, cannot keep up when I do something iteratively.
