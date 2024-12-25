# React 19

## `useActionState`

The `useActionState` (previously `useFormState`) allows you to track the _state_ of the action.

A lot of examples you are going to see will most likely use this hook in a context of _form submission_, but know that you can call the action returned from this hook however you wish – it does not have to be a form!

### Resetting the form

> Based on [this blog post](https://www.nico.fyi/blog/reset-state-from-react-useactionstate)

One of the things we get from `useActionState` is the `state` of the action. The `state` is either the _initial_ state or the value we returned from the action. In your component, you will most likely use this `state` to display something to the user.

```tsx
const initialState = "";

const [state, dispatch, isPending] = useActionState(async (previousState: string, state: string) => {
  await submitName();
  return state;
}, initialState);

// Somewhere down in the component

<p>Your name is {state}</p>;
```

**How do we reset this `state` variable**?

When using a form library, like `react-hook-form`, you can call `reset()` and the values of the form will be reset to their default values. The `useActionState` does not give us a `reset` function, so what can we do?

#### First attempt – the `reset` button

My initial inclination was to reach for the `type="reset"` HTML button, like so.

```tsx
export function Page() {
  const initialName = "Wojciech";
  const [state, dispatch, isPending] = useActionState(async (previousState: string, state: string) => {
    await submitName();
    return state;
  }, initialName);

  return (
    <Fragment>
      <form
        className={"p-4 border"}
        action={(formData) => {
          dispatch(formData.get("name") as string);
        }}
      >
        <legend>Name form</legend>
        <fieldset>
          <label htmlFor="name"></label>
          <input type="text" name="name" id="name" defaultValue={state} />
          <button className={"block mt-4"} type="submit">
            {isPending ? "Submitting..." : "Submit me"}
          </button>
          <button type="reset">Reset</button>
        </fieldset>
      </form>
      <p>Your name is {state}</p>
    </Fragment>
  );
}
```

The `type="reset"` button works on the same basis as [the `type="reset"` input element](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/input/reset) – it will reset all the controls in the form to their initial values.

**But that only works for _uncontrolled_ inputs**. When we click `reset`, the `action` callback is not fired, so the value of the input is preserved.

#### Second attempt – calling the action again

We can call the action again, this time with the `initialName`.

```tsx
export default function Page2() {
  const initialName = "Wojciech";

  const [state, dispatch, isPending] = useActionState(async (previousState: string, state: string) => {
    if (state === initialName) {
      return initialName;
    }

    await submitName();
    return state;
  }, initialName);

  return (
    <Fragment>
      <form
        className={"p-4 border"}
        action={(formData) => {
          dispatch(formData.get("name") as string);
        }}
      >
        <legend>Name form</legend>
        <fieldset>
          <label htmlFor="name"></label>
          <input type="text" name="name" id="name" defaultValue={state} />
          <button className={"block mt-4"} type="submit">
            {isPending ? "Submitting..." : "Submit me"}
          </button>
          <button
            type="reset"
            onClick={() => {
              dispatch(initialName);
            }}
          >
            Reset
          </button>
          {/* <input type="reset" value={initialState} /> */}
        </fieldset>
      </form>
      <p>Your name is {state}</p>
    </Fragment>
  );
}
```

Notice the condition within the `useActionState` callback. Every `dispatch` call will trigger the `isPending` if your perform any async-work. This might or might not be what you want.

#### Built-in automatic reset

**If you do not provide the `value` or the `defaultValue`, React will reset the form for you after action is done**.

```tsx
export default function Page2() {
  const initialName = "Wojciech";

  const [state, dispatch, isPending] = useActionState(async (previousState: string, state: string) => {
    await submitName();
    return state;
  }, initialName);

  return (
    <Fragment>
      <form
        className={"p-4 border"}
        action={(formData) => {
          dispatch(formData.get("name") as string);
        }}
      >
        <legend>Name form</legend>
        <fieldset>
          <label htmlFor="name"></label>
          <input type="text" name="name" id="name" />
          <button className={"block mt-4"} type="submit">
            {isPending ? "Submitting..." : "Submit me"}
          </button>
        </fieldset>
      </form>
      <p>Your name is {state}</p>
    </Fragment>
  );
}
```

This is quite nice. Prior to that change, you had to manually reset the form via `.reset` called on the ref to the form node. Not that great DX.

## `cache`

> [Based on this great video](https://www.youtube.com/watch?v=MxjCLqdk4G4).

I first **thought that I could use the `cache` function to memoize promises for `Suspense` on the client-side**. It turns out that the **`cache` function is NOT a good candidate for this use case**.

The reason is that the **`cache` function memoizes values only for the duration of the request**. So, usage of this function should be scoped to _server components_. You **can use the `cache` in _client components_ ([reference](https://youtu.be/uttgAAZUdYk?t=3023)), but the function is a noop in that environment**.

I like to think about the `cache` as **_data loader_ WITHOUT the batching**.

[According to this video](https://youtu.be/uttgAAZUdYk?t=1674), the `cache` is internally implemented via `AsyncLocalStorage`.
