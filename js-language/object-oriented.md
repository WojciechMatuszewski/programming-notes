# Object-oriented-js Hard Parts

## `Object.create`

It creates brand new empty object, using an existing object as the prototype of the newly created object. It's used heavily in OLOO pattern

```javascript
var myBasePrototype = {
  sayName() {
    console.log("Wojtek");
  },
  sayLastName() {
    console.log("Matuszewski");
  },
};
var obj = Object.create(myBasePrototype);
obj.sayName();

console.log(obj.__proto__); // => myBasePrototype
```

Object.create sets `__proto__` to a given object.

## Object factory

Going with the 'prefer composition over inheritance' we can create a factory function which will generate objects for us

```javascript
function userCreator(name, score) {
  return {
    name,
    score,
    increment() {
      this.score += 1;
    },
  };
}
const user1 = userCreator("Phil", 4);
user1.increment();
console.log(user1.score);
```

##### Big gotcha

Remember about call-site context. With above example we created `this`-aware function. We called that function with `user1` context. But you might get question like this:`

```javascript
function userCreator(name, score) {
  return {
    name,
    score,
    increment() {
      this.score += 1;
    },
    sayScore() {
      console.log(this.score);
    },
  };
}
const user1 = userCreator("Phil", 4);
user1.increment();

var reporter = user1.sayScore;
reporter(); // undefined
// sayScore was called without a context (ALWAYS LOOK AT THE CALL SITE!)

// to fix this we could use .call
reporter.call(user1); // 5
```

## Prototype

### Factory functions using `Object.create`

You can combine factory functions with prototype to get some nice results.

```javascript
// let us create a shared-base of functionalities
var userFunctionStore = {
  // this functions  has to be this-aware, they are meant to be reusable
  increment() {
    this.score++;
  },
  login() {
    console.log("You are logged in");
  },
};

function userCreator(name, score) {
  var newUser = Object.create(userFunctionStore);
  newUser.name = name;
  newUser.score = score;
  return newUser;
}
```

But this gets repetitive, we can shorten the process of creating that object.

### `new` keyword

This keyword automates a lot of things

- make an object for us (there is no longer a need for `Object.create`)
- we would not need to return an object by ourselves, it will return it automatically
- sets context of the function to the newly created object (the one, which we previously created with Object.create)

But how do we specify the bond we have previously been creating with `Object.create`?

JavaScript gives us a special place to place our shared functions or properties called `.prototype`

#### Functions are Objects and Functions at the same time

This should seem obvious by now, especially when you write some React code.

```javascript
interface Props {
  num: number;
}
const Component: React.FC<Props> = () => {
  return <div>Works</div>;
};
// so without Functions being an Objects we would not be able to set props directly on that Function :)
Component.defaultProps = {
  num: 3
};
```

When you create a function it also gets an object attached to it, that object is created with **`.prototype`** prop on it

###### Gotcha

```javascript
function triangle(a, b, c) {
  this.a = a;
  this.b = b;
  this.c = c;
}

// so when you want to set base prototype of a function

// do not be a fool and do this

Object.setPrototypeOf(triangle, {});

// above wont work
// we are setting prototype so we have to access it directly

// why triangle.prototype ?
// well, Object.setPrototypeOf sets __proto__ (ohh i love js)
// and since functions in their object representation do not have __proto__
// we have to set it on their .prototype (property which they have and that property is an object so it has .__proto__ on it)
Object.setPrototypeOf(triangle.prototype, {});
// now everything should work
```

This `.prototype` will be used to put all of our shared functions or properties.

### `new` keyword under the hood

```javascript
// userCreator is a function object combo, that object has .prototype property
function userCreator(name, score) {
  this.name = name;
  this.score = score;
}

userCreator.prototype.increment = function () {
  this.score++;
};

userCreator.prototype.login = function () {
  console.log("login");
};

// calling userCreator function
// mutating context of that function
const user1 = new userCreator("eva", 9);
user1.increment();
```

So what happened while creating a `new userCreator`?

- new keyword creates brand new empty object and sets `userCreator` context to that object
- sets `__proto__` prop of that empty object to `userCreator.prototype` (link, reference)
- auto returns the context object (this)

### Scoping issues with `new`

```javascript
...
userCreator.prototype.increment = function() {
  function add1() {
    this.score++;
  }
  // where does this keyword points to ?
  // well who knows? probably window
  // it's certainly not to the right 'this'
  add1()

  // how would we solve it?

  // well the solution is pretty clear
  // use arrow function instead of normal function or apply/bind


  // arrow function will look up the enclosing scope for this keyword
  // and since outer scope is a function with the 'right' this, everything will work :)
  const add1() => this.score++;
  add1()
}
```

## ES6 Class

`class` keyword is just a cover up, syntactic sugar on top of true prototype behavior

We are doing even less work with `class` keyword

```javascript
class UserCreator {
  //  assignment using this keyword inside the function
  constructor(name, score) {
    this.name = name;
    this.score = score;
  }

  // adding methods to a prototype
  increment() {
    this.score++;
  }
  login() {
    console.log("login");
  }
}

const user1 = new UserCreator("Eva", 9);
```

## Default prototype

### Objects

```javascript
var obj = {
  num: 3,
};
obj.num; // 3
obj.hasOwnProperty("num"); // this comes from __proto__
```

As soon as JS starts, it adds function object combo (built-in).
It's called `Object` but its also a function. That object has a prototype. That Prototype has few methods including `hasOwnProperty`. By default `obj`'s `__proto__` points to that Object prototype

### Functions

As with the objects, JS also adds Function function object combo

```javascript
function multiplyBy2(num) {
  return num * 2;
}

// lookup goes like this
// => multiplyBy2 (object from) => multiplyBy2.__proto__ => Function.prototype
multiplyBy2.toString(); // where is this method

Function.prototype; // {toString: ... and many others}

// what about
multiplyBy2.hasOwnProperty("score");
// multiplyBy2 (object form) => multiplyBy2.__proto__ =>
// Function.prototype => Function.prototype.__proto__ => Object.prototype
```

## Subclassing

### Most manual subclassing

```javascript
function userCreator(name, score) {
  const newUser = Object.create(userFunctions);
  newUser.name = name;
  newUser.score = score;
  return newUser;
}

var userFunctions = {
  sayName() {
    console.log(this.name);
  },
  increment() {
    this.score++;
  }
};

function paidUserCreator(paidName, paidScore, accountBalance) {
  const newPaidUser = userCreator(paidName, paidScore);
  // establish a link to paidUserFunctions
  // setPrototypeOf sets .__proto__
  // this is important to understand
  Object.setPrototypeOf(newPaidUser, paidUserFunctions);
  newPaidUser.accountBalance = accountBalance;
  return newPaidUser;
}

var paidUserFunctions() {
  ...
}
// establish a link to userFunctions
// sets .__proto__ again
Object.setPrototypeOf(paidUserFunctions, userFunctions)


// so with paidUsers lookup goes like this
// object => paidUsers.__proto__ => paidUsers.__proto__.prototype
```

### Subclassing with `new` keyword

```javascript
... userCreator code from above ...

function paidUserCreator(paidName, paidScore, accountBalance) {
  // monkey-patching this with properties from userCreator
  userCreator.call(this, paidName, paidScore);
  this.accountBalance = accountBalance;
  // remember that new keyword implicitly returns this object (it's context)
}

// remember that functions (object representation) do not have .__proto__
// setPrototypeOf could also be used here
paidUserCreator.prototype = Object.create(userCreator.prototype)
// Object.setPrototypeOf(paidUserCreator.prototype, userCreator.prototype)
// adding new stuff to existing prototype
paidUserCreator.prototype.increaseBalance = function() {
  this.accountBalance++;
}
```

### Subclassing with `class` keyword

```javascript
class userCreator {
  constructor(name, score) {
    // assigning props to this context created by new keyword
    this.name = name;
    this.score = score;
  }
  // adding props to .prototype
  sayName() {
    console.log(this.name);
  }
  increment() {
    this.score++;
  }
}

// using extends does is equivalent to doing one of this
// - paidUserCreator.prototype = Object.create(userCreator.prototype)
// - Object.setPrototypeOf(paidUserCreator.prototype, userCreator.prototype)

// extends also sets .__proto__ on paidUserCreator to userCreator
// so that when we call super it knows 'where to go'
class paidUserCreator extends userCreator {
  constructor(paidName, paidScore, accountBalance) {
    // use userCreator to help with creating .this context (it knows that it's userCreator because .__proto__ is set to it)
    super(paidName, paidScore);
    // context is created inside userCreator and returned here
    // now we can add more properties to it
    this.accountBalance = accountBalance;
  }
  increaseBalance() {
    this.accountBalance++;
  }
}

var paidUser1 = new paidUserCreator("Alyssa", 8, 25);
```

- when you invoke `new paidUserCreator()` context (aka this) is being 'born' inside **userCreator**, that's why you are not able to do anything before calling super(), otherwise `.this` would be undefined

```javascript
// this is not a valid syntax but that's what is happening behing the scenes

// call a function like it was called with new but without actually using new keyword
this = Reflect.construct(userCreator, [paidName, paidScore], paidUserCreator);

// also you might think about what's happening like this
this = new userCreator(paidName, paidScore);
// with this object returned by new class has .__proto__ which points to userCreator.prototype
// but super hijacks that mechanic and does
this.setPrototypeOf(this.prototype, paidUserCreator.prototype);
```

Super does not do any side-effects, before we were modifying an object inside a function body with .call. **Construction of the context is happening inside userCreator not paidUserCreator constructor**
