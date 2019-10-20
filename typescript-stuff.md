# Typescript Stuff

## Testing

What can we test in the realm of TypeScript? Code ? sure, but what about the
Types.

Is it possible to write units for TypeScript Types? Well indeed it is possible.

There are number of libraries that check types, either with special _generic
type_ or a comment inside the code.

The strategy is as follows:

- write a test for type declaration -> this just means using that function on a
  very simple data (but that data has to check type-implementation)
- run `tsc` on test files

### dtslint

This tool was built by Microsoft. A sample test:

```ts
var stooges = [
  { name: 'moe', age: 40 },
  { name: 'larry', age: 50 },
  { name: 'curly', age: 60 }
];
_.pluck(stooges, 'name'); // $ExpectType string[]
```

In this example we are using a special comment that ensures that this is the
return type.

You can even test with different versions of TS:

```ts
// TypeScript Version: 2.1
export function pluck<K extends keyof T, T>(array: T[], key: K): Array<T[K]>;
```

[More on this topic](https://medium.com/hackernoon/testing-types-an-introduction-to-dtslint-b178f9b18ac8)

## Assert Signatures

So with `type guards` you are returning `true` or `false`. This then determines
the outcome of the type.

But `assert signatures` **are quite different**

### Different schematics

Type Guard:

```ts
function isDefined<T>(x: T): x is NonNullable<T> {
  return x != undefined;
}
```

Assert Signature:

```ts
function isDefined<T>(x: T): asserts x is NonNullable<T> {
    if(x == undefined) {
        throw AssertionError('Not defined!')
    }
}
```

The signature differs greatly and there is actually more to the
`assert signatures` signature than presented here.

### Two types of Assert Signatures

There are actually 2 variants

- for checking a condition
- for telling TypeScript that specific variable or property has a different type

So it all basically boils down to that `Assert Signatures` does not return
anything, they throw this `AssertionError` whenever something is wrong.

`Type Guards` on the other hand return `true` or `false` based on they inputs.

The signature `asserts something` or `asserts x is something` tells the reader
of the code that **that function will only return if the assertion holds**

```ts
function checkIfString(input:any) asserts input is string {
    if (typeof input != 'string') throw Error('must be a string')
}

function doSomething(val: number | string) {
    checkIfString(val)
    val // string here!
}

```

## Null Coalescing

This is more of a JavaScript thingy but hey, we are all probably writing only
TypeScript now :)

So, do you remember the deal with `&&` and `||` ?

- with `&&` you guard the right value with left value (checking _truthiness_)

- with `||` you either will get left or right value depending on their
  _truthiness_

And _truthiness_ is the key-word here.

So the deal with `Null Coalescing` is that it only checks for `null` and
`undefined`

```js
console.log(0 || 'something'); // something
console.log(0 ?? 'something'); // 0
```

This can help in cases where you have valid non-truthy values as your _guardian
values_ but you still want to check for `null` and `undefined`

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

With combine power of `Exclude` and `Pick` we can do some nice stuff (especially
with HOC's). Let's say we want to remove a prop from something in a generic way.

```typescript
// from Root Pick...
type Omit<Root, PropsToOmit> = Pick<
  Root,
  // Exclude these props from Root which can be found in PropsToOmit
  Exclude<keyof Root, keyof PropsToOmit>
>;
```

### Caution warning

What happens if `PropsToOmit` is a single value, let's say `string`. Well then
bad things will happen. `keyof string` will actually look at it's prototype
chain.

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

Now with helpers (we are creating those because `ReturnType` does not take
second argument)

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

This feature is very useful. Basically you can `pluck` a type from generic using
conditional types.

**We can place the infer keyword at the position where we want the type to be
inferred.**

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

It's very similar to accessing object values and `Object.keys` in JS. **It's
just that the value is the type itself**

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

Sometimes typescript is very strange. It seems that `prop?: number` is not the
same as `prop: number | undefined`?. Let's consider the following

```typescript
interface Something {
  id: number;
  name: string;
  property?: string;
}

type UndefinedAsNever<Type> = {
  [Key in keyof Type]: undefined extends Type[Key] ? never : Type[Key];
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

First thing first we probably should _mark_ `property` somehow so that we know
that we want to _pluck_ this prop.

Remember our `[keyof Something]` notation?

```typescript
interface Something {
  id: number;
  name: string;
  property?: string;
}

type RemoveUndefinableKeys<Type> = {
  [Key in keyof Type]: undefined extends Type[Key] ? never : Key;
}[keyof Type];

type Test1 = RemoveUndefinableKeys<Something>; //"id" | "name" | undefined
```

How does `RemoveUndefinableKeys` work?

```typescript
type RemoveUndefinableKeys<Type> = {
  [Key in keyof Type]: undefined extends Type[Key] ? never : Key;
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

Let's add `[keyof Something]` notation (we will basically get only the values
from the interface).

```typescript
type RemoveUndefinableKeys<Type> = {
  [Key in keyof Type]: undefined extends Type[Key] ? never : Key;
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

So by marking `property` as `undefined` we basically _plucked_ it from the
interface.

No we just need to make `Identity` type generic and name it somehow.

```typescript
type RemoveUndefinableKeys<Type> = {
  [Key in keyof Type]: undefined extends Type[Key] ? never : Key;
}[keyof Type];

type RemoveUndefinable<Type> = {
  // this is the same as Key in "id" | "name" | undefined
  // undefined will be omitted
  [Key in RemoveUndefinableKeys<Type>]: Type[Key];
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

Ahh... sad times. As you want to be _the cool kid_ you probably should use this
_leet hackrzz_ `Record` stuff :

```typescript
type Dict = Record<string, number>;
```

**Remember that in Javascript all keys are `string`'s !!**

## This keyword

When using strictest possible typescript settings (as you always should) you
might night to type `this` keyword. Let's see how this can be done:

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

This may look weird, it may seem like `someFn` now takes 2 arguments but that's
not the case. First argument (`this` typing) will get compiled away.

## Typeof

Here there is distinct difference between Javascript world and Typescript world.

When using Javascript `typeof` will return underlying type as in the type that
you can create in vanilla Javascript. This is familiar territory

```js
typeof []; // "object"
typeof 'something'; // "string"
typeof 3; // "number"
```

But in Typescript `typeof` behaves a little bit differently. Instead of
returning underlying vanilla JS types it will return us the Typescript type.

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

> it tests if a `.prototype` property of a constructor exists somewhere in
> another object

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

This works basically the same as `typeof`. Rarely used (used mainly with
classes, brr)

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

You can also use `in` operator as a _boolean checks_ just like you sometimes
want to check if some browser feature is available.

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

Ever used reducer? You probably used `action.type` or similar property to
distinguish between different actions.

To 'gather' all actions you probably did this:

```typescript
type Actions = ADD | DELETE | SOME_ACTION;
```

Thats the _union_ part. Now the _discriminant_ is the **thing that enables
typescript (and you) to distinguish between different actions**

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

- You can merge declarations with interfaces (you cannot have two types with the
  same name)

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

You can provide different implementations based on the arguments that we supply.
It makes stuff more readable

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

## Declare keyword

This keyword is used for telling typescript that a **javascript construct**
(like a function, variable etc) has already been defined elsewhere. (as a part
of runtime environment)

```ts
declare function add(x: number, y: number): number;

// somewhere in js file for example

function add(x, y) {
  return x + y;
}
```

This allows you to have JS codebase covered with types that are separate. Users
who use typescript can benefit from type completion while users using vanilla
still have access to your library.

## Enums

Enums are quite popular with _Ngrx_. They are not all sunshine and rainbows
though.

- they are typescript only concept
- can cause bundle bloat

```typescript
enum Something {}
// you just introduced this to your bundle
'use strict';
var Something;
(function(Something) {})(Something || (Something = {}));
```

Not looking to hot right? Well, there is a solution. A very simple one. Use
`const` before `enum`. That way the whole `enum` construct will get compiled
away and only picked properties will stay as normal variables. Enum props _get
inlined_

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

**BUT BEWARE** Enums cannot be used with _plugin-transform-typescript_ which you
are probably using.

## Mocking with Typescript

When testing sometimes you have to mock stuff. It's pretty common procedure, but
typescript sometimes makes it difficult.

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

It is frustrating , we have to help typescript a little bit by casting to a
`mock type`

```typescript
import { Link as LinkDep } from 'react-router-dom';

const MockLink = LinkDep as jest.Mock<LinkDep>;

// now you can test in peace
```

Same technique applies to _global mocks_.

## Typing `get` function

Typing such functions is a nightmare. But we can make our lives easier with a
couple of tricks.

[### All credit goes to this article](https://medium.com/@jamesscottmcnamara/type-yoga-typing-flexible-functions-with-typescripts-advanced-features-b5a282878b74)

And btw, we are going all in when it comes to functional programming so our
`get` function will be fp ready :)

### HasKey

Turns out you can create object types out of thin air. Check this out:

```typescript
type HasKey<K extends string, V> = { [_ in K]: V };
type Testing = HasKey<'wojtek', number>;
/*
  {
    wojtek: number
  }
*/
```

This is just mind bending stuff. Very clever usage of the `in` keyword. How does
it work?

- We are extending `string` because generic will be type literal (like `wojtek`
  or `ala`)

- `{[_ in K]: V}` means an object with keys in K with value V

Lets say you use `|` then typing `K` what will happen? Well you will get 2 props
on an object with value: `any` (or any value you passed to generic).

Again, very clever stuff

### Basic Implementation

With `HasKey` we can start our basic implementation.

```typescript
declare function get<K extends string>(
  key: K
): <Obj extends HasKey<K>>(obj: Obj) => Obj[K];
```

Our function could be used as such

```ts
get('name')({ name: 'wojtek' }); // all ok
get('name')({ someprop: 'someprop' }); // Typescript is not happy, error!
```

### KeyAt

You might think that we achieved what we wanted:

> You just have to declare more overloads right?

Not really, sadly this function is far from complete. The inferring system might
have problems with more complex types.

To fix this we introduce another type: `KeyAt`

```typescript
type KeyAt<Obj, K extends string> = Obj extends HasKey<K> ? Obj[K] : never;
interface SomeInterface {
  wojtek: 'ala 123'
}
KeyAt<SomeInterface, 'wojtek'> // 'ala 123', literal type!
```

This allows us to _pluck a given type_ out of object. This makes sure that
return our function has return value correctly typed.

```ts
declare function get<K extends string>(
  key: K
): <Obj extends HasKey<K>>(obj: Obj) => KeyAt<Obj, K>;
```

Personally i would name this type `TypeAt` but I'm going to roll with this name
paying an homage to original author :).

### Traversals

Our function can also filter stuff. We basically want to work _lenses-like_.

Example:

```ts
get(matching(friend => friend.friends > 5), 'name')(obj.friends);
```

With our current implementation this operation is impossible. How would we
enable such functionality?

Let's type `matching` first:

```ts
interface Traversal<Item> {}

declare function matching<A>(
  filteringFunction: (a: A) => boolean
): Traversal<A>;
```

We have to change our implementation a bit to introduce `matching`.

```ts
declare function get<Item, K extends string>(
  traversal: Traversal<Item>,
  key: K
): <Obj extends Array<HasKey<K>>>(obj: Obj) => Array<KeyAt<Obj, K>>;

const popularFriends = get(
  matching((user: User) => user.friends.length > 5),
  'name'
)(user.friends);
```

But there is a problem. Our `popularFriends` are typed as `never[]`.

Going back to our declaration of `KeyAt` we typed it so that `Obj` has to be
`HasKey<>` not `Array` of that type.

That is easily fixable. Just change `obj: Obj` to `obj: Obj[]`.

### Unpacking

Very useful stuff for our function (which we want to be able to accept multiple
containers) and overall (I really wonder why they would not put it inside TS
utility types already).

```ts
// power of conditionals and infer baby!
export type Unpack<F> = F extends Array<infer Item>
  ? Item
  : F extends Set<infer Item>
  ? Item
  : F extends Map<any, Item>
  ? Item
  : F extends { [n: string]: infer Item }
  ? Item
  : F extends { [n: number]: infer Item }
  ? Item
  : never;
```

## Tuples and Currying

This is going to be wild ride so strap on.

### Head

This type will let us pluck off the head of the `tuple Type`. Will come in handy
later

```ts
type Head<A extends any[]> = A extends [infer First, ...any[]] ? First : never;
type Test = Head<[1, 2, 3, 4]>; // 1
```

This type is using `infer` to get the correct type.

### Tail

We implemented `Head` it's time for `Tail` now. As of writing this we cannot
just get the last type out of the tuple.

Lets try the naive approach

```ts
// NOT WORKING
type Tail<A extends any[]> = A extends [any, ...infer tail] ...
```

Sadly we cannot use spread and `infer` together. To implement this type we can
use _function types_. We are going to work on function parameters where we can
directly _infer_ from the rest of arguments.

```ts
type Tail<A extends any[]> = ((...t: A) => any) extends ((
  _: any,
  ...tail: infer TailType
) => any)
  ? TailType
  : never;

type Test = Tail<[1, 2, 3, 4]>; // [2,3,4]
```

We are sort of creating a _virtual type-only function_ that will allow us to
work with parameters which we can type and infer from freely (we cannot do that
using arrays).

### HasTail

Since classical curried functions are taking one argument at a time we have to
know when we should stop and return the return type. `HasTail` type will allows
us to do so

```ts
type HasTail<A extends any[]> = A extends ([] | [any]) ? false : true;
type Test = HasTail<[]>; // false
type Test2 = HasTail<[1, 2]>; // true
```

Pretty straight forward right? Unless our tuple is empty or only has 1 element
left we can keep going with currying.

### CurryV0

With those simple types we can define our **strict curry** type

```ts
type CurryV0<Parameters extends any[], ReturnType> = (
  arg: Head<Parameters>
) => HasTail<Parameters> extends true
  ? CurryV0<Tail<Parameters>, ReturnType>
  : ReturnType;

declare function curry<P extends any[], R>(f: (...args: P) => R): CurryV0<P, R>;

function addTwo(x: number, y: number) {
  return x + y;
}

const curried = curry(addTwo);

curried(1)(2); // works like a charm!
```

I specifically am very verbose with names to make this type clear. This type is
using recursion to gradually (with each call) pluck off one parameter at a time.

### Last

Our curry implementation is great! But, we can always improve on things right?.
So what if we want to support _loose curry_ ? (like partial application). This
would prove to be very difficult using our current tools.

One type that might help us reach that goal of partial application is the `Last`
type.

Instead of plucking off tail from tuple we will only pluck the last type.

```ts
type Last<P extends any[]> = {
  0: Last<Tail<P>>;
  1: Head<P>;
}[HasTail<P> extends true ? 1 : 0];
type Test = Last<[1, 2, 3, 4]>; // 4
```

So this might be hard to digest but stay with me. This type is using recursion

- if there is a tail, pass that tail recursively to `Last`
- if there is no tail use that value, stop recursion

The picking if we need to stop the recursion happens in `[]`. This is what's
called `indexed type accessor`.

You might think we can do the type using normal turnery like :

```ts
type Last2<P extends any[]> = HasTail<P> ? Last2<Tail<P>> : Head<P>
```

This restriction stems from TS itself, you can though reference a type from
within an object type just like we are doing with our first `Last`
implementation.

### Length

This type will allow us to have basic information about arguments that are
passed in and such. This, in terms, will allows us to implement partial
application.

```ts
type Length<T extends any[]> = T['length'];

type Test = Length<[1, 2, 3, 4]>; // 4
```

`Length` type will work as a pseudo-counter.

### Prepend

This will allow us to prepend a type to a tuple type, which will allow us to
know which parameters has already been supplied. To implement this type we will,
again, make us of `function types` trick.

```ts
type Prepend<TypeToPrepend, Tuple extends any[]> = ((
  head: TypeToPrepend,
  ...tail: Tuple
) => any) extends ((...args: infer U) => any)
  ? U
  : never;

type Test = Prepend<number, [1, 2, 3]>; // [number, 1,2,3]
```

Just to make sure you know how this is working. As mentioned before, we cannot
use array destructuring (or spread for that matter) to assign types or use
`infer` keyword.

To circumvent this restrictions we have to operate on functions parameters.

In this type we are basically checking if `Function == Function` (which will
always be the case) but in the act of checking we can merge `TypeToPrepend` and
`TupleType` using `infer U` on the second function.

### Drop

Just like we can `Prepend` type to a tuple type we are also in need of the
ability to remove a number of arguments from the tuple type.

To achieve such functionality we will use **recursive indexed types** (already
seen before)

```ts
type Drop<
  ElementsToDrop extends Number,
  TupleToDropFrom extends any[],
  Iterator extends any[] = []
> = {
  0: Drop<ElementsToDrop, Tail<TupleToDropFrom>, Prepend<Iterator, any>>;
  1: TupleToDropFrom;
}[Length<Iterator> extends ElementsToDrop ? 1 : 0];

type Test = Drop<2, [1, 2, 3, 4]>; // [3,4]
```

One might be curious about that `any` type passed to `Prepend`. Well this
`Iterator` type only accts as to-be-thrown-away accumulator, one might say:
recursion stop predicate. Since we do not care about the type passed to iterator
we default to `any`.

## Matching exact shape

> TypeScript is a structural type system. This means as long as your data
> structure satisfies a contract, TypeScript will allow it. Even if you have too
> many keys declared.

This definitely can be a problem. Example

```typescript
type Person = {
  first: string;
  last: string;
};

const tooFew = { first: 'Stefan' };
const tooMany = { first: 'Stefan', last: 'Joe', other: 'something' };

declare function savePerson(person: Person): void;

savePerson(tooFew); // Error!
savePerson(tooMany); // OK -> WTF ;C
```

We can use clever generic type along with `Exclude` to make sure our params are
the exact shape of given `type / interface`

```ts
type ValidateShape<T, Shape> = T extends Shape
  ? Exclude<keyof T, keyof Shape> extends never
    ? T
    : never
  : never;
```

- check if T is _`Shape`-compatible_ with `extends`
- see if there is 1:1 match between `T` and `Shape` when it comes to properties.
- otherwise return never

Pretty neat stuff!. With this we can re-write our example to:

```ts
declare function savePerson<T>(person: ValidateShape<T, Person>): void;

savePerson(tooFew); // Error -> returns never
savePerson(tooMany); // Error -> returns never
```

## Optional Chaining && Null Coalescing

As of writing this, these count as _future_.

### Optional Chaining

If you ever worked with Angular you are probably familiar with optional
chaining. It is very interesting that Angular had that from the version 2
onwards and Typescript is getting them now 🤔.

So an example:

```js
const someObj = {
  prop1: {
    prop2: undefined
  }
};

const value = someObj.prop1 && someObj.prop1.prop2 & //...
const value = someObj?.prop1?.prop2 // ..
```

Syntax with `?` is much cleaner, especially with nested objects and properties.
You no longer have to worry about checks with `&&`. `?` operator takes care that
for you.
