# React TypeScript v2

- Refactoring from `PropTypes` to TypeScript is not that hard.

  - The overlap is there. Keep in mind that TypeScript has much more extensive type system.

  - Something I've noticed is that the instructor **types the event handlers props explicitly using the `XXEventHandler` type from React**.

    - IMO it is a good practice for strictness sense, but I'm not sure if that is needed. In most cases, the `VoidFunction` will do.

    - I've also noticed the the instructor is also typing the state handlers explicitly. I'm not sure that is a good idea, although it might be if you require the prop to use the callback form as well as the "regular" form where you only pass a number.

- Keep in mind that there are subtle differences between the `interface` and the `type` keyword, though the line gets more blurry with each year passing.

  - It comes down to consistency. Does your team uses `interface`? Then use that. If not, then align on a single way you type your components.

  - The key difference is the extensibility and how different keywords interact with each other.

    - The `interface` can be _augmented_ – **a very important property when writing a library**.

    - The `type` is more "self-contained", it cannot be augmented.

    - For [all the differences, look at this chart](https://twitter.com/karoljmajewski/status/1082413696075382785).

- There are **multiple types for nodes in React**. There include, but are not limited to

  - `JSX.Element`
  - `React.ReactNode`
  - `React.ReactChildren`
  - `React.ReactElement`

  You should **favour the `React.ReactNode` as it is the most correct and the widest of them all**.

  - You can also use the `React.PropsWithChildren<T>` generic which is much better than the `React.FC` or similar types I've seen in the wild.
    The `React.FC` is inferior as it hides the fact that is also allows for children. Your components should not allow for props which they do not use.

- A great tip made by instructor (not really related to TypeScript)

  > If you have a complex logic within the state setters, consider pulling that state setter outside of the component and testing it as a separate function. This way you can cover the whole logic without introducing complexity on the unit test level of your component.

- TIL that **the `input` element has the `valueAsNumber` property**.

    ```jsx
    <input type = "number" onChange = {e => e.currentTarget.valueAsNumber}>
    ```

Stopped at part 3
