# Fundamental methods and functions

## `arguments` and an `array-like`

Let's start with something that was not really known to me.
Did you know that `arguments` is an `array-like` ?
So what is an `array-like` ?

```javascript
function someFunc(arg1, arg2) {
    console.log(arguments.length);
    // the result of this console log is actually an object
    // well you can say : but Wojtek! Array's are objects as well!
    // to that i say : arguments is an array-like
    // it does have length and indexed values but it does not have
    // array methods like .push or .slice or .splice ...etc

    arguments.push(3); // ERR!
}

someFunc(1, 2);
```

So what to do about the `array-like` ?

Enter `slice` and `call`

## `slice`

Quick recap what does `slice` do
It slices an array from given index to given index (optional, if omitted rest of array items will be included)

```javascript
var animals = ["ant", "bison", "camel", "duck", "elephant"];
console.log(animals.slice(2));
// expected output: Array ["camel", "duck", "elephant"]

console.log(animals.slice(2, 4));
// expected output: Array ["camel", "duck"]

console.log(animals.slice(1, 5));
// expected output: Array ["bison", "camel", "duck", "elephant"]

console.log(animals.slice());
// expected output: whole array
```

But the most important thing, **`slice` does not mutate the original array (at least when you have only one dimension)**

## `call`

`call` is used to invoke method with a different context.

```javascript
var obj = {
    name: "Wojtek",
    say() {
        console.log(this.name);
    },
};
obj.say(); // 'Wojtek'
obj.say.call({ name: "Robert" }); // 'Robert'
```

## Using `call` and `slice` together

So remember how you could just call `.slice()` to basically copy the array ? we are going to use that feature with `.call()`

```javascript
function someFunction(arg1, arg2) {
    var argumentsArray = Array.prototype.slice.call(arguments);
    // and so now we have array of arguments which actually have Array methods like push etc...
}
```

It works for all `array-likes`. One of the useful usages is

```javascript
var links = [].slice.call(document.querySelectorAll("a"), 0);
// normally links would be a NodeList now it's an array
```

### New kind in town (`Array.from`)

If you can use ES6 (which should not be a problem, we have babel)
you probably should use newer `Array.from`. The syntax is very explicit and does not look like a black magic.

Argument can be made for `.slice` method for it's performance but we are talking about thousands operations which will probably never happen in a real-word app

## `bind`

Very useful when you want to be sure that function will be called with correct context (new function will be created with given context bound to that function). You cannot 'unbind' given function.

```javascript
var me = {
    name: "Wojtek",
    sayName() {
        console.log(this.name);
    },
};

me.sayName(); // Wojtek

var someoneElse = {
    name: "not me",
};

me.sayName.bind(someoneElse)(); // not me
```

With second argument you can pre-pend arguments. Let's look at an example.

```javascript
function curry(fn) {
    return function curried(...args) {
        // if you cannot use gather remember about Array.prototype.slice.call(arguments)
        // or Array.from(arguments)

        return args.length >= fn.length
            ? fn.apply(null, args)
            : // here we are pre-pending args and when invoked function will take new args as well it will have old args from this call
                curried.bind(null, ...args);
    };
}
```

## `splice`

With splice you can add/remove elements from array. **It mutates the array!**. It returns removed / added elements

```javascript
var arr = [1, 2, 3, 4];
var removed = arr.splice(0, 2);
// => [1,2]

// you can also add stuff
var replaced = arr.splice(0, 0, 1, 2);
// => [] (we did not replace anything)
// => array is now [1,2,3,4]
```

## `reverse`

Well, the name speaks for itself but **it mutates the array!**
Remember the weird behavior when you wanted to do pipe from compose and tried reversing the array? Well that's your answer

```javascript
function composeTwo(f, g) {
    return function composed(arg) {
        return f(g(arg));
    };
}
function compose(...fns) {
    return fns.reduce(composeTwo);
}

function pipe(...fns) {
    // it would be very bad for us to mutate arguments passed
    // so we are using the old trick using .slice.call
    return compose.apply(compose, Array.prototype.slice.call(fns).reverse());
}

// now we have pipe and compose funcions :)
```

## Note to remember

If you find yourself thinking if a function (apply vs call) takes an array or comma separated list of arguments remember this:

- a is for array
- c is for comma

- b and c love each other (bind takes comma separated list of params)
