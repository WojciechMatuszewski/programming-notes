# New stuff in ES6+

## Strings interpolation

Quite useful feature, no longer we have to concat strings together using `+`.

```javascript
var num = 3;
// old way
var msg1 = "there were " + num + " cats";
// new way using template strings
var msg2 = `there were ${num} cats`;
```

## Tag literals

You can use your own function for formating or any other stuff you want. You can also use string interpolation inside string interpolation :)

```javascript
// values are the values interpolated inside the string
function formatCurrency(strings, ...values) {
    var str = ""
    ...
        // formatting code here
    ...
    return str
}

var amount = 12.3;
var msg = formatCurrency`The total for your order is ${amount}`;
```

## String Padding

```javascript
var str = "Hello";
// Hello already has 5 characters, nothing to add
str.padStart(5); // "Hello"
// added 3 spaces to Hello since we are missing 3 characters
str.padStart(8); // "   Hello"
// you can use custom characters
str.padStart(8, "*"); // "***Hello"
// you can use multiple characters
str.padStart(8, "123"); // "123Hello"
```

`padEnd` works the same way but from the right to left (characters are still ltr)

## String Trimming

```javascript
var str = "   some stuff \t\t";
str.trim(); // "some stuff"
str.trimStart(); // "some stuff   ";
str.trimEnd(); // "   some stuff";
```

## Destructuring

Destructuring is very useful feature

### Array destructuring

#### Basic array destructing

```javascript
var arr = [1, 2, 3, 4];
// first element
var [first, ...rest] = arr;
// you can pick as many as you want (if there are no more left to pick values will be undefined or an empty array)

// in this case rest is an empty array
var [first, second, third, fourth, ...rest] = arr;

// gather have to be last operation, it cannot be first or in the middle

var [...rest, last] = arr; // error!
```

#### Default values

```javascript
function data() {
  return [1, 2, , 4];
}

// normally third would be undefined but we assigned default value
var [first, second, third = 3, fourth] = data();
```

#### Intermediate variables

```javascript
function data() {
  return [1, 2, 3, 4];
}
var temp;
// temp is assigned to array from data()
// and then that array is getting destructured
var [first, second] = (tmp = data());

// you could also use this syntax
tmp = [first, second] = data();
// tmp is whole array and we still have first and second.
// I prefer the first way though, this seems confusing at first
```

#### Defining variables outside destructuring

```javascript
function data() {
  return [1, 2, 3, 4];
}
// very often we come across syntax that looks something like this
var [first, second, ...rest] = data();
// but it's perfectly valid to move declaration outside destructuring
var first, second, rest;
[first, second, ...rest] = data();

// we could also use an object
var o = {};
[o.first, o.second, o.third, ...o.rest] = data());
```

#### Comma separation

You can skip some values using comma separator

```javascript
function data() {
  return [1, 2, 3, 4];
}
var [first, , third, fourth] = data();
```

#### Swapping values

Let's say we want to swap values without using `tmp` variable

```javascript
var x = 10,
  y = 20;
// here we are swapping values
[y, x] = [x, y];
```

#### Inside parameters

```javascript
// quite useful
function([first,second,third]) {
 ...
}
```

#### Fallbacks

```javascript
function data() {
  return null;
}
// normally this would throw TypeError
var [first, second, third] = data();
// we can prevent that by assigning a fallback value
var [first, second, third] = data() || [];

// fallback with function parameters
// you probably should also supply defaults to individual values
function data([first = 10,second = 20,third = 30] = []) {
    ...
}
```

#### Nested Arrays

```javascript
function data() {
  return [1, [2, 3], 4];
}
// you can also use default values inside destructuring
var [first, [second, third] = [], fourth] = data() || [];
```

### Object destructuring

```javascript
function data() {
  return { a: 1, b: 2, c: 3 };
}
// third is an object (separate)
var { a: first, b: second, ...third } = data();

// with defaults
var {a: first, b:second = 4, ...third = {}}

// you can list same source many times
var {a: first, a: second, a:third} = data();
```

## `Array.find`

Find an item inside array which fulfils the predicate

```javascript
var arr = [{ a: 1 }, { a: 2 }];

arr.find(function match(v) {
  return v && v.a > 1;
});
// {a:2}

arr.find(function match(v) {
  return v && v.a > 10;
});
// undefined

arr.findIndex(function match(v) {
  return v && v.a > 10;
});
// -1 is quite meh here ;/
// -1
```

So what you have to watch out here is that when you do not find any item that fulfils the predicate with `array.find` it returns undefined. But what if you are looking for undefined?

##### An old trick

Remember this?

```javascript
var arr = [10, 20, NaN, 30, 40, 50];
arr.indexOf(30) != -1;
~arr.indexOf(20); // -2
// remember that NaN is the ony type that does not fulfil NaN === NaN
~arr.indexOf(NaN); // -0 (falsy)
```

Yes, i also did that. Well it turns out it was not that good of an idea and there is a better way to do what we were trying to do with this 'tick'

## `Array.includes`

```javascript
var arr = [10, 20, NaN, 30, 40, 50];
arr.includes(NaN); // ACTUALLY TRUE, WOW!
```

`Array.includes` does not do the lying that `===` does sometimes (im looking at you NaN)

## `flat`

This one can help you, well, flatten an array

```javascript
var arr = [1, 2, [3, 4], [5, 6], 7];
arr.flat();
// [1,2,3,4,5,6,7]

// it also takes an optional arg which tells the function how many levels to flatten
// default: 1

// if you are desperate :D
arr.flat(1000);

// you can also pass 0 so it does nothing
arr.flat(0);
```

## `flatMap`

Array => Array of sub arrays => flattened
You could do the mapping then the flattening, but thats not really performant and why bother when we can do both things at the same time

```javascript
[1, 2, 3].flatMap(function all(v) {
  return [v * 2, String(v * 2)];
});
// [2, "2", 4, "4", 6, "6"]
```

You have to remember that the flattening is only one level

#### Removing elements from an array using `flatMap`

We typically thing about the `.map` operator as producing new array (same length). You can actually remove items from an array using flatMap and map the ones that you want

```javascript
[1, 2, 3, 4, 5].flatMap(function doubleEvents(v) {
  return v % 2 == 0 ? [v, v * 2] : [];
});
// [2,4,4,8,6,12]
```

This one is actually neat!

## Iterators

### Declarative iterators

#### Imperative code

So we would have to write something like this if not for declarative approach

```javascript
var str = "Hello";
for(
  let it = str[Symbol.iterator()](), v, result;
  (result = it.next()) && !result.done && (v=result.value || true)
) {
  console.log(v)
}
// "H" "e" ...
```

#### Consuming an iterator

```javascript
var str = "hello";
// we are using default iterator here
var it = str[Symbol.iterator]();
// using iterator explicitly
for (let v of it) {
  console.log(v);
}
// using iterator implicitly
for (v of str) {
  console.log(v);
}
```

#### Spread and iterators

Spread operator consumes iterator

```javascript
var str = "Hello";
var letters = [...str];
letters;
// ["H", ...]
```

###### A neat trick

Ever wanted to create an array from 0 to n without using `Array.from...`?

Create a generator ! (more on them later)

```javascript
Number.prototype[Symbol.iterator] = function* () {
  for (let i = 0; i < this; i++) {
    yield i;
  }
};

var i = 10;

console.log([...i]);
// [1,2,3,4,5,6,7,8,9]
```

## Generators

So instead of creating a function that returns an iterator result `{value: ..., done: true/false}` and handles all that logic, we can create a generator which will do all of that for us.

```javascript
function* main() {
  yield 1;
  yield 2;
  yield 3;
  // we did not yield 4,
  // it equals to actually throwing 4 away, since done will be true
  return 4;
}
var it = main();

it.next(); // {value: 1, done: false}

// remember that spread consumes iterator
[...main()];
// [1,2,3]
```
