# React Typescript

## Polymorphic components

There are the components which change their behavior (in our case _props_) based on some condition (in our case a specific _prop_).

As the `as` prop got popularized by a wave of css-in-js libraries, we often have a need for other props to be _inferred_ from the `as` prop.

Think of a `Button` component example. You might want to pass `as ="a"` so that the props are _inferred_ from the `a` HTML attribute.

```jsx
<Button as = "a" href = "www.google.com"> // works!
<Button href = "www.google.com"> // ts complains about the `href`. By default the `Buttton` is well... a button.
```

Creating those is actually not that hard (even with `forwardRef`). So lets start with a _polymorphic_ `Button` component which does not forward the ref.

### Basic Implementation

For this you need to know that:

- `React.ElementType` describe all possible HTML tags (`a`, `button` etc...) and also our components (class or functional, does not matter)

- To get props from a given `ElementType` you should use `React.ComponentProps<E>` helper type

With these points in mind, the implementation is rather simple

```tsx
type BaseProps<E extends React.ElementType> = {
  as?: E;
  someBaseProp: string;
};

type ComponentProps<E extends React.ElementType> = BaseProps<E> &
  Omit<React.ComponentProps<E>, keyof BaseProps<E>>;

const baseElement = "button";

function Button<E extends React.ComponentType = typeof baseElement>(
  props: ComponentProps<E>
) {
  const Element = props.as ?? baseElement;
  return <Element {...props} />;
}
```

Pretty simple right? All we are really doing is passing around the `as` prop. By default it's a `button`.

### Implementation with `useRef`

Now, things can get tricky here. This is because the `forwardRef` function is generic but you will not be able specify a generic for those properties. An example

```tsx
// this is not a valid syntax
const Button = <E extends any>React.forwardRef(() => {
  // code
})
```

That means that we have to use _casting_. This is possible due to the fact that we used **optional generic parameter for `E` parameter**.

```tsx
// previous code
type ButtonWithRef = <E extends React.ElementType = typeof defaultElement>(
  props: ComponentProps<E>
) => JSX.Element;

const Button = React.forwardRef(
  (props: ButtonBaseProps, ref: React.Ref<unknown>) => {
    // code
    // remember to pas `ref` to the <Element/>
  }
) as ButtonWithRef;
```

There is no possibility for us to know the `React.Ref` parameter so I've opted to use `unknown here`.

### `defaultProps`

You should prefer default value assignment here. Really. Otherwise the definitions gets really awkward. You would have to intersect the `ButtonWithRef` definition with additional `{defaultProps: {}}` type.

This is because by using the _casting_ we are loosing the `defaultProps` typings which `forwardRef` normally provides.
