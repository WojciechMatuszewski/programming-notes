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
