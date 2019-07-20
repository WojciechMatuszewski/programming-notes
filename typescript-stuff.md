# Typescript Stuff

## Pick and Exclude

### Pick<T,K>

> From T, pick a set of properties whose keys are in the union K

You can, well, _Pick_ specific properties from an interface.

```typescript
interface User {
  name: string;
  email: string;
  password: string;
}

type OnlyName = Pick<User, 'name'>; // {name: string}
```

### Exclude<T, U>

> Exclude from T those types that are assignable to U

It works by diffing two types. That's a common gotcha (at least for me).

```typescript
interface User {
  name: string;
  email: string;
  password: string;
}

// remember, we are diffing 2 types, code below will not do what we want it to do
type WithoutName = Exclude<User, 'name'>; // User, because T is not extending U so Exclude returns User

// now we are talking, we are diffing each key with 'name'
type WithoutName = Exclude<keyof User, 'name'>; // 'email' | 'password'
```

## Combining Exclude And Pick

With combine power of `Exclude` and `Pick` we can do some nice stuff (especially with HOC's). Let's say we want to remove a prop from something in a generic way.

```typescript
// from Root Pick...
type Omit<Root, PropsToOmit> = Pick<
  Root,
  // Exclude these props from Root which can be found in PropsToOmit
  Exclude<keyof Root, keyof PropsToOmit>
>;
```

### Caution warning

What happens if `PropsToOmit` is a single value, let's say `string`. Well then bad things will happen. `keyof string` will actually look at it's prototype chain.

```typescript
type Test = keyof 'something'; // "toString" | "charAt" | "charCodeAt" | "concat" | "indexOf" | "lastIndexOf"
```

## Return Type

You can actually get the return type of a given function. Quite neat!

```typescript
function add(x: number) {
  return x + 1;
}

type SomeType = ReturnType<typeof add>; // number
```

### Caution warning

- you have to cast directly if you want to get _value_ as the type

```typescript
const None = 'None';
function something() {
  return { x: None };
}
type HeyTs = typeof None; // "None"
type SomeType = ReturnType<typeof something>; //{x: string}

function something() {
  return { x: None as typeof None };
}
// .. same operations
// return type is now {x: "None"}
```

- you have to create helpers when dealing with generic functions

```typescript
function identity<T>(prop: T): T {
  return prop;
}
type SomeType = ReturnType<typeof identity>; // unknown
```

Now with helpers (we are creating those because `ReturnType` does not take second argument)

```typescript
type Callable<T> = (...args: any[]) => T;

type MyOwnReturnType<ReturnType, F> = F extends Callable<ReturnType>
  ? ReturnType
  : never;

type SomeType = MyOwnReturnType<string, typeof identity>; // string
```

## Conditional Types

### Very basic example

Just like ternary but with types

```typescript
type SomeType<T> = T extends string ? T : never;
interface Obj {
  name: string;
  age: number;
}

type t1 = SomeType<'ala'>; // 'ala'
type t2 = SomeType<Obj['name']>; // string
type t3 = SomeType<Obj['age']>; // never
```

Using this feature you can create (most of them already ship by default) types

```typescript
// already built-in
type NonNullable<T> = T extends undefined | null ? never : T;
type t1 = NonNullableMy<null>; // never
```

### Inferring the type

This feature is very useful. Basically you can `pluck` a type from generic using conditional types.

**We can place the infer keyword at the position where we want the type to be inferred.**

```typescript
// you could also do args: infer U
type GetFunctionArgumentTypes<F> = F extends (...args: Array<infer U>) => void
  ? U
  : never;

function numberArg(x: number) {}

function arrayMixed(x: [1, 'a', {}]) {}

type t1 = GetFunctionArgumentTypes<typeof numberArg>; // number
type t2 = GetFunctionArgumentTypes<typeof arrayMixed>; // [1, 'a', {}]
```

## Mapped And Lookup Types

You can use `in` and `keyof` to transform interfaces and type-objects

```typescript
type MyPartial<Type> = { [Key in keyof Type]+?: Type[Key] };

interface Something {
  id: number;
  name: string;
  property?: string;
}
type MyPartial<Type> = { [Key in keyof Type]+?: Type[Key] };

type Test = MyPartial<Something>;
/*
  {
    id?: number | undefined
    ...
  }
*/
```

### Keyof and `[keyof]`

Differences are quite big

```typescript
type Test1 = keyof Something; // "id" | "name" | "property"
type Test2 = Something[keyof Something]; // string | number | undefined
```

It's very similar to accessing object values and `Object.keys` in JS.
**It's just that the value is the type itself**

```js
var someObj = {
  prop1: 1,
  prop2: 2,
  prop3: 'someString'
};

Object.keys(someObj); // 'prop1' , 'prop2' ...
someObj['prop1']; // 1
```

### Caution warning

Sometimes typescript is very strange.
It seems that `prop?: number` is not the same as `prop: number |undefined`?.
Let's consider the following

```typescript
interface Something {
  id: number;
  name: string;
  property?: string;
}

type UndefinedAsNever<Type> = {
  [Key in keyof Type]: undefined extends Type[Key] ? never : Type[Key]
};

type Test1 = UndefinedAsNever<Something>;
/*
WTF ???
{
    id: number;
    name: string;
    property?: undefined;
}
*/
// you can merge interfaces btw :)
interface Something {
  id: number;
  name: string;
  // changed this
  property: string | undefined;
}
type Test1 = UndefinedAsNever<Something>;
/*
WEIRD STUFF HUH?
{
    id: number;
    name: string;
    property: never;
}
*/
```

### Plucking nullable (also undefined) keys

Let's say you have an interface

```typescript
interface Something {
  id: number;
  name: string;
  // we want to remove this \/
  property: string | undefined | null;
  // can also be written like this
  property?: string;
}
```

First thing first we probably should _mark_ `property` somehow so that we know that we want to _pluck_ this prop.

Remember our `[keyof Something]` notation?

```typescript
interface Something {
  id: number;
  name: string;
  property?: string;
}

type RemoveUndefinableKeys<Type> = {
  [Key in keyof Type]: undefined extends Type[Key] ? never : Key
}[keyof Type];

type Test1 = RemoveUndefinableKeys<Something>; //"id" | "name" | undefined
```

How does `RemoveUndefinableKeys` work?

```typescript
type RemoveUndefinableKeys<Type> = {
  [Key in keyof Type]: undefined extends Type[Key] ? never : Key
};
// would return
/*
  {
    id: "id",
    name: "name",
    property: undefined
  }
*/
```

Now we _marked_ `property` as the one to be deleted (by undefined type)

Let's add `[keyof Something]` notation (we will basically get only the values from the interface).

```typescript
type RemoveUndefinableKeys<Type> = {
  [Key in keyof Type]: undefined extends Type[Key] ? never : Key
}[keyof Type];
// would return "id" | "name" | undefined
```

See ? no `property` prop.

We can also do this

```typescript
interface Something {
  name: string;
  age: number;
}

type Identity = { [Key in 'name' | 'age']: Something[Key] };
// would return the same Something type
```

So by marking `property` as `undefined` we basically _plucked_ it from the interface.

No we just need to make `Identity` type generic and name it somehow.

```typescript
type RemoveUndefinableKeys<Type> = {
  [Key in keyof Type]: undefined extends Type[Key] ? never : Key
}[keyof Type];

type RemoveUndefinable<Type> = {
  // this is the same as Key in "id" | "name" | undefined
  // undefined will be omitted
  [Key in RemoveUndefinableKeys<Type>]: Type[Key]
};

type Test = RemoveUndefinable<Something>;
/*
would return
{
  id: number;
  name: string;
}
*/
```

### Record

Remember typing dictionaries awkwardly like:

```typescript
type Dict = { [key: string]: number };
```

Ahh... sad times.
As you want to be _the cool kid_ you probably should use this _leet hackrzz_ `Record` stuff :

```typescript
type Dict = Record<string, number>;
```

**Remember that in Javascript all keys are `string`'s !!**

## This keyword

When using strictest possible typescript settings (as you always should) you might night to type `this` keyword. Let's see how this can be done:

```typescript
interface SomeObj {
  someFn: (num: number) => number;
  numberToAdd: number;
}

const someObj: SomeObj = {
  someFn,
  numberToAdd: 4
};

function someFn(num: number) {
  return num + this.numberToAdd; // might cause an error
  // typescript sometimes has problems with inferring the right this
}

// much better implementation would be
function someFn(this: SomeObj, num: number) {
  return num + this.numberToAdd; // much better, we even get autocomplete
}
```

This may look weird, it may seem like `someFn` now takes 2 arguments but that's not the case. First argument (`this` typing) will get compiled away.

## Typeof

Here there is distinct difference between Javascript world and Typescript world.

When using Javascript `typeof` will return underlying type as in the type that you can create in vanilla Javascript. This is familiar territory

```js
typeof []; // "object"
typeof 'something'; // "string"
typeof 3; // "number"
```

But in Typescript `typeof` behaves a little bit differently.
Instead of returning underlying vanilla JS types it will return us the Typescript type.

```typescript
const person = {
  age: 22,
  name: 'Wojtek'
};

type Person = typeof person; // {age: number, name: string}
```

This is very powerful especially with `ReturnType`.

## Type Guards

Using a _Type Guard_ you can tell Typescript which type something is.

### Typeof Type Guard

This is very simple guard. Check this out:

```typescript
function someFn(arg: number | string) {
  if (typeof arg == 'number') {
    // typescript knows we are dealing with a number here
    return arg.toExponential(); // ok
  }
  // typescript knows we are dealing with a string here
  // BUT BEWARE!
  // if we did not return above type would be number | string here
  arg.toLowerCase(); // ok
}
```

### `instanceof` Type Guard

#### In vanilla JS

> it tests if a `.prototype` property of a constructor exists somewhere in another object

Example:

```js
class Foo {
  bar() {}
}

const foo = new Foo();

// we all know this is true
Object.getPrototypeOf(foo) == Foo.prototype)
// above is essentially the same as:
foo instanceof Foo
```

#### In Typescript

This works basically the same as `typeof`. Rarely used (used mainly with classes, brr)

### User Defined Type Guard

Now we are talking. Check it out:

```typescript
interface Response {
  result: any;
  doSmth: any;
}

interface OkResponse extends Response {
  status: 'OK';
}
interface BadResponse extends Response {
  status: 'NOT_OK';
}

// now whenever we use this function typescript is going to set that variable as OkResponse
// if this function returns true, otherwise it will be BadResponse
function isGoodResponse(response: OkResponse | BadResponse) response is OkResponse {
  return response.status == 'OK'
}

```

### `in` Type Guards

You can also use `in` operator as a _boolean checks_ just like you sometimes want to check if some browser feature is available.

```typescript
interface Athlete {
  speed: 99;
  age: 30;
}
interface NormalPerson {
  age: 30;
}
function isAthlete(subject: Athlete | NormalPerson): subject is Athlete {
  return 'speed' in subject;
}
```

## Intersection Types

Instead of `extend`ing interfaces you can use `&` to _merge_ them.

```typescript
interface Order {
  amount: number;
}
interface Stripe {
  cvc: string;
  card: string;
}
interface PayPal {
  email: string;
}

// i think this is much better than interface Stripe extends Order {}
type OrderWithStripe = Order & Stripe;
type OrderWithPayPal = Order & PayPal;

// typescript is great at inferring as well!
const stripeOrder = Object.assign({}, order, stripeData); // OrderWithStripe
```

## Discriminant union

Ever used reducer? You probably used `action.type` or similar property to distinguish between different actions.

To 'gather' all actions you probably did this:

```typescript
type Actions = ADD | DELETE | SOME_ACTION;
```

Thats the _union_ part. Now the _discriminant_ is the **thing that enables typescript (and you) to distinguish between different actions**

```typescript
  reducer(state, action) {
    // type is a common property
    // that lives on all of the actions
    switch(action.type) {
      case DELETE:
      // type inference works because of discriminant unions!
    }
  }
```

## Interface vs Type

- You cannot use `extend` keyword with types but you can use `&` instead

```typescript
interface Item {
  name: string;
}

interface Artist extends Item {
  songs: number;
}

type Artist2 = {
  songs: number;
} & Item;
```

- You can merge declarations with interfaces (you cannot have two types with the same name)

```typescript
interface Artist {
  name: string;
}
interface Artist {
  songs: number;
}
// /\ merged together

// now interface Artist contains name and songs
```

## Function Overloads

You can provide different implementations based on the arguments that we supply. It makes stuff more readable

```typescript
// these are virtual, they will get compiled away
function reverse(dataToReverse: string): string;
function reverse<T>(dataToReverse: T[]): T[];
// real implementation
function reverse<T>(dataToReverse: string | T[]): string | T[] {
  if (typeof dataToReverse == 'string') {
    return dataToReverse
      .split('')
      .reverse()
      .join('');
  }
  return dataToReverse.slice().reverse();
}
```

## Enums

Enums are quite popular with _Ngrx_.
They are not all sunshine and rainbows though.

- they are typescript only concept
- can cause bundle bloat

```typescript
enum Something {}
// you just introduced this to your bundle
'use strict';
var Something;
(function(Something) {})(Something || (Something = {}));
```

Not looking to hot right? Well, there is a solution. A very simple one. Use `const` before `enum`. That way the whole `enum` construct will get compiled away and only picked properties will stay as normal variables. Enum props _get inlined_

```typescript
const enum Something {
  yes = 'Yes',
  no = 'No'
}
let selected = Something.no;

// gets compiled to
'use strict';
let selected = 'No' /* no */;
```

Much better now!

**BUT BEWARE**
Enums cannot be used with _plugin-transform-typescript_ which you are probably using.

## Mocking with Typescript

When testing sometimes you have to mock stuff. It's pretty common procedure, but typescript sometimes makes it difficult.

```typescript
import { Link as MockLink } from 'react-router-dom';

jest.mock('react-router-dom', () => ({
  ...jest.requireActual('react-router-dom'),
  Link: ({ children }: { children: React.ReactNode }) => children
}));

test(/* some test name */, () => {
  // TypeError: !
  MockLink.mockImplementationOnce(() => {/* ... */})
})
```

It is frustrating , we have to help typescript a little bit by casting to a `mock type`

```typescript
import { Link as LinkDep } from 'react-router-dom';

const MockLink = LinkDep as jest.Mock<LinkDep>;

// now you can test in peace
```

Same technique applies to _global mocks_.
