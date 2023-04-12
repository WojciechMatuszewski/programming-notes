# React Typescript

## Getting all props for a given HTML node

This is especially useful when building any re-usable components. Picture a situation where you have to build a re-usable `Button` components.
You most likely have some _design-system specific_ props and you would also like to accept all other props the `button` HTML can.

```tsx
type DesignSystemProps =  {
  leadingIcon: Icon,
  variant: string
  // and so on
}

type Props = DesignSystemProps

const MyButton = ({leadingIcon, variant, ...rest}: Props) => {
  // the `rest` is typed as {}
  return <button>...</button>
}
```

There are various ways to do it, but I think the most straightforward one is to use the `ComponentProps` type from React typings.

```tsx
type DesignSystemProps =  {
  leadingIcon: Icon,
  variant: string
  // and so on
}

type Props = DesignSystemProps & ComponentProps<"button">

const MyButton = ({leadingIcon, variant, ...rest}: Props) => {
  // the `rest` contains all the props the `button` can have.
  return <button>...</button>
}
```

Also, when we are at it, **make sure you forward the ref to the button**. Being unable to use the `ref` on the consumer side is a bit soul crushing.

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

## `ref` being immutable

I do not know about you, but whenever I write `React.useRef` I want to be as explicit as possible. This often leads me to write something like this

```tsx
const myRef = React.useRef<HTMLDivElement>(null);
```

Now, this would be fine and all, but whenever you when try to mutate that ref you, the type system will scream at you

> Cannot assign to 'current' because it is a read-only property

Now, this is weird right, should not all _refs_ be mutable? Well, yes, but...

If the **type that you have provided does not contain `null` and the initial value is `null`, you signal that you do not intent to modify the ref value**. This is so called **immutable ref, or DOM ref**.

In the previous example, I've used `HTMLDivElement`, but the notion of _immutable ref_ is not constrained to types related to the DOM nodes.

```tsx
const myRef = React.useRef<number>(null);
myRef.current = 3; // TypeScript error
```

### How to you signal the mutable `ref` then?

There are two cases that you have to consider:

1. The initial value I want the `ref` to have is `null`.
2. The initial value I want the `ref` to have is different than `null`.

For the first case, include the `null` within the type definition of the `useRef`. So `React.useRef<number>(null)` becomes `React.useRef<number | null>(null)`

For the second case, you do not have to change your code at all.

```tsx
const myRef = React.useRef<string>("hey");

myRef.current = "hio";
```

### The union type within the ref type parameter

We've talked about `ref` being immutable, but did you notice that I passed only a singular type as a `useRef` type parameter?

When you specify a union of types

```ts
const myRef = React.useRef<HTMLDivElement | null>(null);
```

the _immutable_ semantics will no longer apply. You will be able to freely mutate the ref (given that the value you are setting the ref to adheres to the types)

Overall, I think we should be as precise as possible while writing code, thus IMO, using _non-union_ types where you will not be mutating the `ref` is a good idea.

## Event handlers and mistakes

Let's say you pass a `fetchMore` function as a prop. Since `fetchMore` has complex parameter signature, and you do not really care about parameters (**but you would not want anyone to pass any parameters to it)**, you annotate it as `VoidFunction`.

The `VoidFunction` annotation should make it impossible to call it parameters right?

```tsx
interface ComponentProps {
  fetchMore: VoidFunction;
}
```

Now within the `Component` someone attaches that `fetchMore` to a `onClick` handler, like so

```tsx
function Component({ fetchMore }: ComponentProps) {
  return <button onClick={fetchMore}>OnClick</button>;
}
```

You run the tests, and ... well, and _SyntheticEvent_ instance is passed to your `fetchMore` function. WTF?

This is because **TypeScript will allow you to assign functions with no parameters to function that take parameters**.

An example

```ts
const x = (a: number = 1): number => a;
const y: () => number = x;
```

This will compile.

As you could guess, this is what is going on here. TypeScript have no problems with the `onClick` passing the _SyntheticEvent_ instance to your function - it should be ignored in the first case right?

### Getting clever with `never`

Okay, let's use `never`. This should work right?

```tsx
interface ComponentProps {
  fetchMore: (arg: never) => void;
}

function Component({ fetchMore }: ComponentProps) {
  return <button onClick={fetchMore}>OnClick</button>;
}
```

And the above compiles. WTF?

Well, this is because the `never` is a bottom type. `never` extends any other type;

```ts
type Foo = [never] extends [React.MouseEvent] ? true : false; // true
```

As a side-note, if you are not sure why I used the `[]` here, please read up on _naked_ vs _clothed_ types.
There is a section within `typescript-stuff.md` on it.

### Improving the `never` solution

To have the cake and eat it too, there is a small adjustment we need to do to our existing definition.

```tsx
interface ComponentProps {
  fetchMore: (arg?: never) => void;
}

function Component({ fetchMore }: ComponentProps) {
  // TypeScript error
  return <button onClick={fetchMore}>OnClick</button>;
}
```

Nice, now we can skip the `arg` parameter if we choose to invoke the prop directly (not using _point-free_ style, like inside the `Component`)
and we achieved our result, the are certain that no parameter will be passed to the function.
If someone were to try to pass anything, he would be greeted with TypeScript error :).

To avoid TypeScript errors in our particular situation, the prop would have to be used as follows

```tsx
function Component({ fetchMore }: ComponentProps) {
  return <button onClick={() => fetchMore()}>OnClick</button>;
}
```

### Form event handler

Did you know you can actually pick the value of every input element of the form from the event that is passed to your `onSubmit` handler?
You can even do it in a semi type-safe way - here is how.

```tsx
import * as React from "react";

interface UserFormElements extends HTMLFormControlsCollection {
  name: HTMLInputElement;
  surname: HTMLInputElement;
}

interface UserForm extends HTMLFormElement {
  readonly elements: UserFormElements;
}

export default function App() {
  function onSubmit(event: React.FormEvent<UserForm>) {
    event.preventDefault();

    const { name, surname } = event.currentTarget.elements;

    console.log(name.value, surname.value);
  }

  return (
    <form onSubmit={onSubmit}>
      <label htmlFor="name">Name</label>
      <input id="name" type="text" placeholder="name" />
      <br />
      <label htmlFor="surname">Surname</label>
      <input id="surname" type="text" placeholder="surname" />
      <br />
      <button type="submit">Submit</button>
    </form>
  );
}
```

Notice the `HTMLFormElement` and `HTMLFormControlsCollection` types. I can easily extend them to type the values what my form is dealing with.
In todays world, you will most likely never have to deal with raw forms like that, but either way I think it's worth knowing how to type them correctly.

Here is an article you can refer to whenever someone, or you, makes a mistake.
<https://epicreact.dev/how-to-type-a-react-form-on-submit-handler/>

## Prop Patterns

### Only one prop or the other, not both

There are two ways to accomplish this. Through overloading or unions.

Let us pretend we are designing a component that can be either _expanded_ or _collapsed_. How should we make it impossible for the user to pass an impossible state state properties?

#### Using unions

This way of handling such problems is definitely the most mainstream one.

```tsx
type Props =
  | { expanded: boolean; collapsed?: never }
  | { expanded?: never; collapsed: boolean };

function Component(props: Props) {...}
```

Now, depending on the editor you are using, you will still get a code completion for the `never` properties.
I would not sweat much about that, unless you are not using tsc before compilation, the consumer of the `Component`
will not be able to pass the impossible state.

#### Using overloading

This technique will make sure you cannot even specify the `never` annotated prop - because there are no such properties.
I would be in favour of this technique if the team I'm working with is comfortable with it.

```tsx
function Component(props: { expanded: boolean }): React.ReactNode;
function Component(props: { collapsed: boolean }): React.ReactNode;
function Component(props: { expanded: boolean } | { collapsed: boolean }) {}
```
