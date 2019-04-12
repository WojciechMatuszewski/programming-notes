# Javascript types

> In JavaScript everything is an object --- every inexperienced, **wrongly thinking**, JS dev

## Primitive types

#### List of primitive types

- undefined
- string
- number
- boolean
- object
- symbol

#### Other things that may behave like types (but are not)

- undeclared
- null
- function
- array
- bigint

#### These are objects

- object
- function
- array

#### In JavaScript, variables do not have types, values do

```javascript
var v;
typeof v; // "undefined"
v = "1";
typeof v; // "string"
v = 2;
typeof v; // "number"
v = true;
typeof v; // "boolean"
v = {};
typeof v; // "object"
v = Symbol();
typeof v; // "symbol"
```

We are not asking what's the type of variable `v`, we are asking what's the type of value that `v` holds

#### About undefined

`undefined` does not mean variables does not yet have value, it means _does not currently have a value_

#### About `typeof` operator

`typeof` will always return strings, it cannot return anything else

```javascript
typeof doesNotExist; // "undefined"

// this might cause a bug
var v = null;
typeof v; // "object"

v = function() {};
typeof v; // "function"

// array type is subset of object
v = [1, 2, 3];
typeof v; // "object"
```

#### `undefined` vs `undeclared` vs `uninitialized`

`undeclared` - never been created, in any scope we have access to.
`undefined` - we defined a variable but at the moment it does not have a value.
`uninitialized`- (TDZ state, see scope section and block variables). Block scope variables does not get initialized to undefined when they never initially get set to `undefined`

## Special values

#### NaN

It does **not** mean _not a number_, it's more like _invalid number_

```javascript
var myAge = Number("0o46"); // 38
var myNextAge = Number("39"); // 39
var myCatsAge = Number("n/a"); // NaN
myAge - "my son's age"; // NaN

// NaN is a special value, it does not equal to itself
myCatsAge === myCatsAge; // false

isNaN(myAge); // false
isNaN(myCatsAge); // true
isNaN("my son's age"); // true HMMMMMMMM, weird

Number.isNaN(myCatsAge); // true
Number.isNaN("my son's age"); // false  Number.isNaN to the rescue
```

So why does `isNaN` shows strings as `NaN` values? Because it coerces to `number` before it does the check and string coerced to number may result in `NaN` value (like in this example)

#### Negative Zero

```javascript
var trentRate = -0;
trendRate === -0;

trendRate.toString(); // "0" OOPS!
trendRate === 0; // false OOPS!
trendRate < 0; // false
trendRate > 0; // false

// Object.is for the rescue
Object.is(trendRate, -0); // true
Object.is(trendRate, 0); // false
```

## Fundamental Objects (Built-In Objects or Native Functions)

You should use new

- Object
- Array
- Function
- **Date**
- RegExp
- Error

Do not use new (use as functions not as constructors)

- String
- Number
- Boolean

## Abstract Operations (coercion)

#### ToPrimitive(hint [optional])

Any time when we have non-primitive and it needs to become a primitive, conceptually we have to do some steps to turn it into primitive and that is called `ToPrimitive`

It also takes (optional) type hint. It's telling the `ToPrimitive` something like _it would be nice to get back [hint]_. It does not guarantee to return the hint.

This might be a recurring process.

##### How does ToPrimitive works

Any non-primitive value in JavaScript have 2 methods on them which are quite useful for `ToPrimitive`

- valueOf()
- toString()

The order in which these get invoked while doing `ToPrimitive` algorithm steps depends on the _hint_ provided to `ToPrimitive`

#### ToString

Takes any value and gives us representation of that value in string form.

##### Gotchas

Most of primitive types behave as you expect but there is a gotcha with -0 (what a surprise )

```javascript
-0.toString() === "0" // true OOPS!
```

Weird things happen when using on Array (so you probably should not use in real code [maybe for debugging])

```javascript
[].toString() // => ""
[1,2,3].toString() // => "1,2,3"
[null,undefined].toString() // => ","
[[[],[],.[]], []].toString() // => ",,,"
[,,,,].toString() // ",,,"
```

With Objects, it's better than with Arrays but still weird

```javascript
{}.toString() // => "[object Object]"
{a:2}.toString() // => "[object Object]"
// you can also override toString() (you should use Symbol)
{toString(){return: "X";}}.toString() // => "X"
```

#### ToNumber

##### Gotchas

```javascript
"".toNumber() === 0; // true OOPS!
null.toNumber() === 0; // true meh
Number.isNaN(undefined.toNumber()) === true; // true OOPS!
```

#### ToBoolean

It does **not** invoke ToPrimitive or ToNumber algorithm. It just does a lookup to the falsy list.

Falsy values

- 0, -0
- null
- NaN
- false
- undefined

If it's not on the list it's always truthy

## Coercion

You can use `Number` as a function or `+` to convert stuff to number

```javascript
Number("42"); // 42 (most explicit form)
+"42"; // 42 (kinda meh, still works but looks weird)
```

You can use `Boolean` or `!!` to convert stuff to booleans

```javascript
!!"something"; // => true
Boolean("something"); // => true
```

Watch out for empty strings

```javascript
Boolean(""); // => false
Boolean(" \t\n"); // => true OOPS!
```

### Boxing

How come we can do `.toString()` or `.length` or similar function calls on a primitive ?

Thats called `boxing`. It's a form of implicit coercion
Actually whats happening is JavaScript implicitly coerces the primitive values to their object counter parts so we have methods on them.

### == vs ===

There is a common misconception about `==` vs `===`. It is though that
`==` checks value and `===` checks value and type **WHICH IS COMPLETELY FALSE!!!!!**

##### They both "check the type"

First step of _Abstract Equality Comparison_ algorithm says that if
Type(x) is the same as Type(y) then return x === y (strict equality)

##### `===` checks the type

First step of _Strict Equality Comparison_ algorithm says that
if Type(x) is different from Type(y) return false

##### The real difference

The reals difference between these operators is that one allows coercion to happen ( == ) other do not ( === )

##### Comparing identity

Comparison operator works by comparing identity not structure

```javascript
var workshop1 = {
  name: "..."
};
var workshop2 = {
  name: "..."
};

console.log(workshop1 === workshop2); // false
console.log(workshop1 == workshop2); // false
```

Comparison is false because they are different objects in memory

##### `null` and `undefined`

`null` and `undefined` are coarsely equal to each other.
You can use `==` to match against both of them

```javascript
var test = null;
var test2 = undefined;
console.log(test == test2); // true
```

##### Avoid

- `==` with 0 or `""` or `" "`
- avoid `==` with non-primitives
- avoid `== true` or `== false`, allow ToBoolean or use `===`

##### The case for preferring `==`

`==` is **not** about comparisons with unknown types, it's about comparisons with known type(s), optionally where conversions are helpful

If you know the type(s) in a comparisons:

- if both types are the same, `==` is identical to `===`, using `===` would be **unnecessary** in this case
- if the types are different, using one `===` would be **broken**

If you do not know the type(s) in a comparison:

- the uncertainty of now knowing the types should be oblivious to the reader
- the most obvious signal is ===
