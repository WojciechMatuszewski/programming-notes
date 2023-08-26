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

## Generic components

In TypeScript you can provide the generic slots for arrow functions and the "regular" functions. The syntax, while a bit different, is more or less similar.

```ts
const myGenericArrowFn = <T>(a: T) => {};
function myGenericFn<T>(a: T) {}
```

**While in JSX files, the syntax to define a generic slot for an arrow function looks a bit different**. This is most likely due to the potential confusion of you editor where it is unable to distinguish a JSX tag vs a generic slot syntax.

```ts
const MyComponent = <T,>(props:T) => {}
```

**Notice the `,` here**. Pretty weird, but it works.

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

- **`React.ElementType` describes all possible HTML tags (`a`, `button` etc...) and also our components (class or functional, does not matter)**

  - This is **different than `JSX.IntrinsicElements` as this type describes only the HTML elements and not custom components**.

- To get props from a given `ElementType` you should use `React.ComponentProps<E>` helper type

  - You can also use variations of this type, like `React.ComponentPropsWithoutRef`.

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

### The `Slot` component as alternative to `as`

The `as` prop and the whole notion of _polymorphic components_ is pretty complex, especially the types. Would it not be nice to have an interface that is just _as good as_ the `as` prop, but have it be much less complex? **Enter the `Slot` component**.

> Do not mistake the `Slot` component with the notion of _component slots_ as props. These are two different things!

The premise is the following: **if you pass a certain prop, you are responsible for rendering the correct element type. The `Slot` component will only merge the props**.

```tsx
// Imagine that the `Button` implements some complex logic and passes the props to `children`
<Button asChild = {true}>
  <a href = "#">My button rendered as link with merged props</a>
</Button>
```

And here is how the `Slot` component looks and is used.

```tsx
function Slot({
  children,
  ...props
}: { children: React.ReactNode } & React.HTMLAttributes<HTMLElement>) {
  if (React.Children.count(children) !== 1) {
    throw new Error("boom");
  }

  if (React.isValidElement(children)) {
    return React.cloneElement(children, {
      ...props,
      ...children.props
    });
  }

  return null;
}
```

The role of the `Slot` component is to **merge the props passed to it with the props of the child if wraps**. The user of the `Button` component would never be exposed to the `Slot` component.

```tsx
function Button({children, ...props}) {
  const myButtonComplexStateAndProps = {}
  if (props.asChild) {
    return <Slot {...props} {...myButtonComplexStateAndProps}>{children}</Slot>
  }

  return <button {...props} {...myButtonComplexStateAndProps}>{children}</button>
}
```

This makes the interface a bit more explicit. You also use one of the most powerful features React has to offer – composability! While typing the `asChild={true}` might be a bit weird at first, this approach is starting to get traction. The most notable uses (and probably inventors) of this pattern is the [`radix-ui` library](https://github.com/radix-ui/primitives).

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

Now, depending on the editor you are using, you will still get a code completion for the `never` properties. I would not sweat much about that, unless you are not using tsc before compilation, the consumer of the `Component` will not be able to pass the impossible state.

**Keep in mind that TypeScript might complain when you destructure the props without checking of the discriminant first**.

```tsx
type ModalProps =
  | {
    variant: "no-title";
  }
  | {
    variant: "title";
    title: string;
  };

// Works
export const Modal = (props: ModalProps) => {
  if (props.variant === "no-title") {
    return <div>No title</div>;
  } else {
    return <div>Title: {props.title}</div>;
  }
};

// Does not work
export const Modal = ({variant, ...props}: ModalProps) => {
  if (props.variant === "no-title") {
    return <div>No title</div>;
  } else {
    return <div>Title: {props.title}</div>;
  }
};
```

TypeScript is not smart enough to handle this workflow yet.

#### Using overloading

This technique will make sure you cannot even specify the `never` annotated prop - because there are no such properties.
I would be in favour of this technique if the team I'm working with is comfortable with it.

```tsx
function Component(props: { expanded: boolean }): React.ReactNode;
function Component(props: { collapsed: boolean }): React.ReactNode;
function Component(props: { expanded: boolean } | { collapsed: boolean }) {}
```

## Type-helpers

React exposes various type-helpers to help you type your components/hooks faster.

### Automatic `children` with `PropsWithChildren`

Writing `children: React.ReactNode` might get cumbersome after a while. To save you a little bit of typing, one might use `PropsWithChildren` type-helper.

```tsx
const Component = ({someProp, children}: React.PropsWithChildren<{someProp: number}>) => {}
```

Of course, I **strongly think you should NOT use this type-helper for every component**. As it was the case with `React.FC`, marking every component as taking `children` prop is misleading as some of them will not do anything with the `children` prop.

### Mirror HTML element without `ref` prop with `ComponentPropsWithoutRef`

It is quite common, especially when creating a design system or a component library, to have to "mirror" an HTML element. Your design system might include custom buttons, sliders and inputs. It is vital to have the props be correct so that the customers of the library are happy.

To help you achieve that goal, you should **consider using `ComponentPropsWithoutRef`** when creating custom HTML-like components.

```tsx
interface Props extends React.ComponentPropsWithoutRef<"button"> {
  scale: "small" | "large" | "medium"
}

const MyCustomButton = ({scale, ...buttonHTMLProps}: Props) => {
  return <button {...buttonHTMLProps}>SomeButton</button>
}
```

Of course, **this only applies if you DO NOT want to `forwardRef` which I argue you SHOULD do in this case**. It is quite frustrating as a library consumer not to have the ability to get the `ref` of the underlying element. **For that, use the `React.ComponentPropsWithRef`**.

### Strongly typed `useRef` with `ElementRef`

When using the `useRef` hook, you most likely want to pass a type to the generic slot of the hook.

```ts
const ref = useRef<SomeTypeHere>()
```

The problem is, that sometimes, it is hard to know what the `SomeTypeHere` supposed to be. Usually it is a HTMLElement type, but sometimes it could be a custom type, especially when you use `useImperativeHandle`.

This is **where the `ElementRef` type-helper comes in**. The `ElementRef` will help you to derive the right type for the ref. Check this out.

```tsx
import { ElementRef, forwardRef, useImperativeHandle, useRef } from "react";

function App() {
  // Works well
  const someElementRef = useRef<ElementRef<"audio">>()

  // Also works well!
  const componentWithForwardedRef = useRef<ElementRef<typeof ComponentWithForwardedRef>>()

  // This one as well!
  const componentWithImperativeHandle = useRef<ElementRef<typeof ComponentWithImperativeHandle>>();

  return (
    <div>it works</div>
  );
}

const ComponentWithForwardedRef = forwardRef<HTMLDivElement>((props, ref) => {
  return <div>works</div>
})


const ComponentWithImperativeHandle = forwardRef<{ someFunc: VoidFunction }>((props, ref) => {
  useImperativeHandle(ref, () => {
    return {
      someFunc: () => { }
    }
  })
  return <div>works</div>
})

export default App;
```
