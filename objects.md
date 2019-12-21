# Javascript Objects (oriented)

#### **_This_** keyword

A function's this references the execution context for that call, determined entirely by **how the function was called**

Example

```javascript
// this aware function
function ask(question) {
  console.log(this.teacher, question);
}

function otherClass() {
  var myContext = {
    teacher: "Suzy"
  };
  // calling this aware function with myContext as function context
  ask.call(myContext, "Why");
}

otherClass(); // Suzy, Why
```

##### Implicit **_this_** binding

```javascript
var workshop = {
    teacher:"Kyle",
    ask(question) {
        // because of a call site this keyword will end up
        // pointing to the teacher object key
        console.log(this.teacher, question)
    }
};
// calling ask on workshop so the context is implicitly workshop
workshop.ask("...")
```

You can also share some functionality and assign context accordingly. This allows us to use **_this_** aware functions in different contexts

```javascript
function ask(question) {
  console.log(this.teacher, question);
}

var workshop1 = {
  teacher: "Kyle",
  // ES6
  ask
};

var workshop2 = {
  teacher: "Suzy",
  // ES6
  ask
};

workshop1.ask("..."); // Kyle
workshop2.ask("..."); // Suzy
```

##### Explicit **_this_** binding

Other ways to set up context binding

```javascript

// you can use .call or .apply

// it's easy to remember
// a is for array
// c is for comma

...
ask.call(workshop1, 'some question')
ask.apply(workshop1, ['some question'])
```

##### Hard **_this_** binding

```javascript
var workshop = {
  teacher: "Kyle",
  ask(question) {
    console.log(this.teacher, question);
  }
};

// this is not the call site
// since setTimeout takes a function (callback) as an argument
// call site will be inside it's implementation
// so without binding this we will loose that bound :C
setTimeout(workshop.ask, 10, "Lost this?");

// we hard bound this, everything will work :)
setTimeout(workshop.ask.bind(workshop), 10, "Hard bound this");
```

##### "Constructor calls"

The purpose of the **_new_** keyword is to invoke a function with the **_this_** keyword pointing to a whole new empty object

```javascript
function ask(question) {
  console.log(this.teacher, question);
}
var newEmptyObject = new ask("...");
```

**_new_** keyword does 4 things

- create a brand new empty object out of thin air
- link that object to another object (proto and prototype)
- call function with **_this_** set to the new object
- if function does not return an object, assume return of **_this_**

##### Bind rules precedence

- is the function called by new ?
- is the function called with call or apply ? (bind uses apply under covers)
- is the function called on a context object ?
- default to a global (except strict mode)

##### Arrow functions

```javascript
var workshop = {
    teacher: "Kyle"
    ask(question) {
        setTimeout(() => {
            console.log(this.teacher, question);
        }, 1000);
    }
};
workshop.ask("Is this lexical this")
// Kyle is this lexical this?
```

Arrow functions do not have 'hard bound' **_this_**, it's a meme.
Proper way to thing about it is that arrow function does not define **_this_** keyword at all. There is no such thing as **_this_** inside arrow function. This means that it will lexically resolve to some enclosing scope that is **_this_** aware.

```javascript
var workshop = {
    teacher: "Kyle"
    ask(question) {
        // we land here, this is this aware function
        // and since we invoked ask with a workshop obj context this keyword here points to a teacher key object :)
        setTimeout(() => {
            // from here we are going lexically up one level
            console.log(this.teacher, question);
        }, 1000);
    }
};
workshop.ask("Is this lexical this")
```

###### Gotcha

```javascript
var workshop = {
  teacher: "Kyle",
  // so you would think that ask would look up to the object as it's next scope
  // but nope, just because there are curly braces (object) does not make it a scope
  // it will go to global instead
  ask: question => console.log(this.teacher, question)
};

workshop.ask("What happened to this");
workshop.ask.call(workshop, "still no this?");
```
