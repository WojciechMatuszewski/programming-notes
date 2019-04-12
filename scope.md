# Javascript Scope a.k.a where to look for things

#### Javascript is not interpreted language, its compiled language (parsed)

Let's say you have a syntax error somewhere in your code. If you run that piece of code you will see only that syntax error,no previous code was run. That means javascript first 'processes' then runs your code.

## Scope example

#### Javascript organizes scopes with functions and blocks (let, const)

```javascript
// create globally scoped variable
var teacher = "Kyle";

// function keyword is a declaration
// it also makes a 'marble'
function otherClass() {
  // scope manager creates a new 'bucket' (functions local scope)

  // this shadows a variable created above, of course these are completely different variables (they are scoped differently)
  var teacher = "Suzy";
  // by shadowing global variable we have no way to access the global variable in this functions scope

  console.log("Welcome!");
}
// we step back to global scope

// next formal declaration, things happen just as before :)
function ask() {
  // scope manager creates a new 'bucket' (functions local scope)

  // create scoped variable in this scope
  var question = "Why";
  console.log(question);
}

// identifier in source position
// hey global scope, i've got a source reference called otherClass, ever heard of it?
// we pull the value out, that value is a function that other class points to
// execution moves inside that function
otherClass(); // Welcome!
ask(); // Why
```

All the scopes are defined at compile time. Scopes are used at run time but never redefined.

#### Script mode vs non-strict mode

```javascript
// auto-created variable teacher would be available on global scope level

function something() {
  // normally you would get Reference Error (with strict mode)
  // but without it scope manager (global) auto-creates a variable
  // it's worth pointing out that, for auto-globals to work, variable needs to be in target position
  teacher = "Kyle";

  // this is a source reference, auto-globals will not work here
  // even without script mode we will get Reference Error
  someFunc();
}

something();
```

#### Scope with function expressions

```javascript
// function declaration
function teacher() {}

// function expression
var myTeacher = function anotherTeacher() {
  // we can reference anotherTeacher because this scope has a
  // reference to anotherTeacher (see differences below)
  console.log(anotherTeacher);
};

console.log(teacher); // ok, function declared in this scope
console.log(anotherTeacher); // Reference Error, function not declared in this scope, it can only be referenced within it's own scope
```

One of key differences is that function expressions attach their name to their own scope. (function declaration attaches their name to the enclosing scope, in this example global)

## Named function expressions

```javascript
// function expression (anonymous), not a function declaration because keyword 'function' is not the first thing in the statement
var clickHandler = function() {};

// named function expression
// you can self-reference the keyHandler within it's scope,
// that reference is *read-only*
var keyHandler = function keyHandler() {};
```

#### Why should you prefer named function expressions?

- you can self-reference within the function (e.g recursion)
- named function expressions show up in the stack traces 🙌
- code becomes more self-documenting.

#### What about arrow functions?

Arrow functions are still anonymous, should prefer named functions, but do not stress about it too much.

## Undefined vs undeclared

Undefined means a variable exists, but at the moment has no value.

Undeclared is never formally declared in any scope we know (we do not have a marble for it).

## Lexical scope

That's the thing that we talked about so far. That's the process of compiler and scope manager figuring thing out. It's related to compiler, author-time decisions. Decided at compile time.

#### Lexical scope example

```javascript
var teacher = "Kyle";

function otherClass() {
  var teacher = "Suzy";

  function ask(question) {
    // teacher reference got 'locked in', it's fixed and predictable
    console.log(teacher, question);
  }

  ask("Why?");
}
```

Lexical scope is popular because it's easily optimizeable, compiler gets to know where all the variables come from at compile time.

#### Dynamic scope in javascript

Dynamic scope does not exist in javascript, **but** there is a mechanism that gives us the same type of flexibility as dynamic scoping (more on that later).

#### Function scoping

First lets examine this piece of code

```javascript
var teacher = "Kyle";

// variable is re-assigned not re-declared !
var teacher = "Suzy";
console.log(teacher); //Suzy

console.log(teacher); // Suzy -- oops
```

The problem is not that the variable could be re-assigned but the problem is the naming collision in the same scope.

How to fix this problem? Well just make sure they use different scopes

```javascript
var teacher = "Kyle";

function anotherTeacher() {
  // now we are using a different scope, using shadowing
  var teacher = "Suzy";
  console.log(teacher); // Suzy
}

anotherTeacher();
console.log(teacher); // Kyle
```

We resolved our problem but still have a naming collision (shadowing)
and we added additional function (kinda meh ;/).

We can improve our solution even more by using an IFFIE

```javascript
var teacher = "Kyle";

// IFFIE
// this is not a function declaration because
// the word function is not the first word in the expression
// this IFFIE does not pollute global scope
(function anotherTeacher() {
  var teacher = "Suzy";
  console.log(teacher); // Suzy
})();

anotherTeacher();
console.log(teacher); // Kyle
```

#### Block scoping

Instead of writing and IFFIE let's write a block

```javascript
// var is function scoped
var teacher = "Kyle";

{
  // let is block scoped
  let teacher = "Suzy";
  console.log(teacher); // Suzy
}

console.log(teacher); // Kyle
```

Block scoping is useful with encapsulating a logic in e.g if statements

```javascript
function diff(x, y) {
  if (x > y) {
    // sadly this will be function scoped inside diff function
    // but the premise is to signal to the reader that this variable should only be accessible within this if statement
    var tmp = x;
    // to get the desired behavior use let (or const in this example)
    /*
      const tmp = x
    */
    x = y;
    y = tmp;
  }
  return y - x;
}
```

Var keyword can be useful inside try/catch blocks

```javascript
function lookupRecord(searchStr) {
  // variable id is available at function level
  try {
    var id = getRecord(searchStr);
  } catch (e) {
    // this is not re-declaration, this only re-assigns
    var id = -1;
  }
  return id;
}
```

As a good practice you should use block scoping for variables that only _exists_ for a few lines

```javascript
function formatStr(str) {

  // now we encapsulated prefix and rest inside their own block scope
  // we are not polluting function scope
  {
    let prefix, rest;
    prefix = str.slice(0,3);
    rest = str.slice(3);
    str = prefix.toUpperCase() + rest;
  }

  if (...) {

  }

  return ...
}
```

#### const vs let

```javascript
var teacher = "suzy";
teacher = "Kyle"; // ok

const myTeacher = teacher;
myTeacher = "Suzy"; // error

const teachers = ["Kyle", "Suzy"];
teachers[1] = "Brian"; // allowed!
```

const keyword allows array and object mutation but not re-assignment. It's kind of confusing because a lot of people think that "const" means variable that cannot change.

###### Takeaways

Use it only when you have primitive value type and never going to re-assign that value, otherwise you might as well use let or var

#### Hoisting

###### Hoisting per-say does not exists. It's a phrase we (JS developers) introduced to explain some behavior of JavaScript

```javascript
// we could explain what really happens here (2 passes,
//scope manager talking with scope of current function),
// but we are kind of lazy and we are going to introduce a magic concept of hoisting
student;
teacher;
// as you know during the first pass scope manager will see these and declare them
// (put them in a bucket and create a marble)
// all statements without keywords like var or function will get ignored (during the first pass)
var student = "you";
var teacher = "Kyle";
```

Javascript does not magically move you code around. How would that even work? Some kind of magical look-ahead mechanism would have to be in place.

If you want to find variable declarations further down the code chain the only way to do it is with **parsing**

Hoisting is not really a thing it's a convention to omit what really happens.

How does function behave ?

```javascript
function teacher() {
  return "Kyle";
}

var otherTeacher;

teacher(); // Kyle
otherTeacher(); // TypeError

otherTeacher = function() {
  return "Suzy";
};
```

With function expressions you have to have them declared (assigned) before you call them.

Another example

```javascript
var teacher = "Kyle";

otherTeacher(); // undefined

function otherTeacher() {
  // remember about 2 pass
  // so teacher here is undefined as a value
  // but it has been defined in a parsing phase (shadowing the outer teacher)
  console.log(teacher);
  // scope manager will define this statement as a variable but assign it's value only when the second pass will get here
  var teacher = "Suzy";
}
```

Let's tackle the **_meme_** that let/const does not "hoist"
Of course it does!

```javascript
{
  // here scope manager (block scope) will declare this variable
  // but instead of assigning undefined to it as to a var
  // variable will be in so called TDZ state
  teacher = "Kyle"; // TDZ error
  let teacher;
}
```

Another more through proof

```javascript
{
  // so if variable teacher did not exist here we would go to the outer scope
  // and look there for that variable (it's declared there with var [function scope])
  // but instead we get a TDZ error because variable exists, but it's in TDZ state
  teacher = "Kyle"; // TDZ error
  let teacher;
}

var teacher = "Kyle";

{
  // same as above
  console.log(teacher); // TDZ error!
  let teacher = "Suzy";
}
```

###### Takeaways

- hoisting is a made up term
- var "hoists" inside a function scope
- let/const "hoists" inside a block scope
- when var "hoists" it's initialized with undefined
- when const/let "hoists" it's not initialized (TDZ) [different that undefined]

#### Closures

Closure is like a backpack that function carries. You can put a lot of stuff in that backpack. To be more technical closure is when a function "remembers" it's lexical scope even when the function is executed outside that lexical scope

Basic example of a closure

```javascript
// this is a basic example of a closure
function ask(question) {
  // we pack question variable inside waitASec's backpack
  setTimeout(function waitASec() {
    // this function is executed outside it's lexical scope
    // setTimeout is a browser API
    console.log(question);
  }, 1000);
}
ask("What is closure");
```

Closure is per scope basis.

Closure is not using "snapshots" of the thing that it's closing over. It's the real value.
Example

```javascript
var teacher = "Kyle";

var myTeacher = function() {
  console.log(teacher);
};

teacher = "Suzy";

myTeacher(); // Suzy
```

Simple gotcha

```javascript
for (var i = 1; i <= 3; i++) {
  // this function closes over i, sure
  // but how is that i variable scoped?
  // i is function scoped so it goes (in this case)
  // to a global scope, so i accessed here will always print 4 (in this case)
  setTimeout(function() {
    console.log(i);
    // you can even change timeout value to 0, does not matter
    // same behavior will persists
  }, 1000);
}
```

How you can fix the gotcha?
You can declare i (iterator) as a block scoped variable

```javascript
for(let i = 1; i <= 3; i++) {
  // now i is blocked scoped, setTimeout will close over i
  // which has correct(1,2,3) value at the time of invoking setTimeout
  ...
}
```

More robust solution
If you cannot use let or you are feeling creative you can use IFFIE's :)

```javascript
for (var i = 1; i <= 3; i++) {
  // you probably should rename argument taken by the IFFIE
  // (left the i here for reference)
  (function(i) {
    setTimeout(function() {
      console.log(i);
    }, 0);
  })(i);
}
```

#### Modules

Modules encapsulates some parts of their own API and exposes others.

##### Revealing module pattern (classic)

```javascript
var workshop = function(function Module(teacher) {

  var publicAPI = {ask, };
  return publicAPI;

  // ask is closed over teacher
  function ask(question) {
    console.log(teacher, question);
  }

  // we could have added more functions here and decide if we want to
  // put them in the publicAPI object or not

})("Kyle")

workshop.ask("...");
workshop.teacher // this variable is hidden
```

The purpose of a module is that you have some state that you closed over and you are controlling access of the state by exposing publicAPI

##### Module factory

```javascript
function WorkshopModule(teacher) {
  var publicAPI = { ask };
  return publicAPI;

  function ask(question) {
    console.log(teacher, question);
  }
}

var workshop = WorkshopModule("kyle");
workshop.ask("...");
```

Factory is creating modules that are independent of each other.

##### ES6 modules

```javascript
var teacher = "Kyle";

export default function ask(question) {
  console.log(teacher, question);
}
```
