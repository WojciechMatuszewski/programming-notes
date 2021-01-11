# Javascript classes

##### ES6 class example

```javascript
class Workshop {
    // constructor is optional
    constructor(teacher) {
        this.teacher = teacher;
    }
    ask(question) {
        console.log(this.teacher,question_)
    }
}

var deepJS = new Workshop("Kyle");
var reactJS = new Workshop("Suzy");

deepJS.ask("Is class a class?");
reactJS ask('...');
```

Classes can use **_extend_** keyword

```javascript
class Workshop {
  constructor(teacher) {
    this.teacher = teacher;
  }
  ask(question) {
    console.log(this.teacher, question);
  }
}

class AnotherWorkshop extends Workshop {
  speakUp(msg) {
    this.ask(msg);
  }
}
```

It also has **_super_** keyword

```javascript
class Workshop {
    constructor(teacher) {
        this.teacher = teacher;
    }
    ask(question) {
        console.log(this.teacher, question);
    }
}

class AnotherWorkshop extends Workshop {
    ask(msg) {
        super.ask(msg.toUpperCase());
    }
}
```

Classes still have dynamic **_this_**

```javascript
class Workshop {
    constructor(teacher) {
        this.teacher = teacher;
    }
    ask(question) {
        console.log(this.teacher, question);
    }
}
var deepJS = new Workshop("KYle");
setTimeout(deepJS.ask, 100, "Still losing this");
// undefined Still losing this
```

## Prototypes

##### A "constructor call" makes an object linked to its own prototype (there is no copy relationship)

Under the class syntax-sugar

```javascript
function Workshop(teacher) {
    this.teacher = teacher;
}

Workshop.prototype.ask = function(question) {
    console.log(this.teacher, question);
};

var deepJS = new Workshop("Kyle");
deepJS.ask("...");
```

Here Kyle draws a drawing representing how the prototype relations work, takeaways:

- prototype points on an object which contains shared methods
- constructor points back from the prototype object to object which has linkage to prototype (this is used to trick people into thinking that javascript has a notion of true OOP constructor)
- dunder proto points at Object.prototype, it's a function (getter) so when you call it like `deepJS.__proto__` it still has the deepJS **_this_** context
- arrow functions does not have `.prototype` (that's why you cannot use `.new` on it)

When you try to shadow and not using class system, you are stepping on the edge

```javascript
function Workshop(teacher) {
    this.teacher = teacher;
}
Workshop.prototype.ask = function(question) {
    console.log(this.teacher, question);
};
var deepJS = new Workshop("Kyle");
deepJS.ask = function(question) {
    // this.__proto__ => Workshop.prototype
    // binding so that we invoke our method instead of non capitalized one
    this.__proto__.ask.call(this, question.toUpperCase());
};
deepJS.ask("...");
// uppercase console .log
```

##### Object.create [first 2 steps of new algorithm]

- create brand new object (empty)
- link that object to another object

Objects linked using prototype

```javascript
function Workshop(teacher) {
    this.teacher = teacher;
}
Workshop.prototype.ask = function(question) {
    console.log(this.teacher, question);
};

function AnotherWorkshop(teacher) {
    Workshop.call(this, teacher);
}
AnotherWorkshop.prototype = Object.create(Workshop.prototype);
AnotherWorkshop.prototype.speakUp = function(msg) {
    this.ask(msg.toUppercase());
};
var JSRecentParts = new AnotherWorkshop("Kyle");
JSRecentParts.speakUp("Is this actually inheritance");
```

So the call site looks like this

- does `JSRecentParts` has speakUp method? No, go to the prototype
- does `AnotherWorkshop.prototype` has speakUP method? yes, call it (this points to JSRecentParts)
- does `JSRecentParts` has ask method? No go to the prototype
- does `AnotherWorkshop.prototype` has ask method ? no to the another prototype
- does `Workshop.prototype` has ask method? yes, call it (this points to JSRecentParts)

## OLOO: Objects Linked to Other Objects

##### Delegated objects

```javascript
var Workshop = {
    setTeacher(teacher) {
        this.teacher = teacher;
    },
    ask(question) {
        console.log(this.teacher, question);
    },
};

var AnotherWorkshop = Object.assign(Object.create(Workshop), {
    speakUp(msg) {
        this.ask(msg.toUpperCase());
    },
});

var JSRecentParts = Object.create(AnotherWorkshop);
JSRecentParts.setTeacher("Kyle");
```

`JSRecentParts.setTeacher` call goes like this

- does `JSRecentParts` has `setTeacher` method? no, go to the prototype
- does `AnotherWorkshop` has `setTeacher` method ? no, go to the prototype
- does `Workshop` has `setTeacher` method? yes, call it (**_this_** points to `JSRecentParts`)

Another example

```javascript
var AuthController = {
    authenticate() {
        // think about what this keyword points to here
        // it points to LoginFormController, because as Kyle said
        // only think that matters is how the method was called
        // and it was (not implicitly but following the prototype chain) called by LoginFormController
        server.authenticate(
            [this.username, this.password],
            // binding this because server.authenticate is a outside API
            this.handleResponse.bind(this)
        )
    },
    handleResponse(resp) {
        if(!resp.ok) this.displayError(resp.msg)
    }
}

var LoginFormController = Object.assign(
    Object.create(AuthController),
    {
        onSubmit() {
            this.username = ...;
            this.password = ...;
            this.authenticate();
        },
        displayError(msg) {
            alert(msg)
        }
    }
);
```

When you call `LoginFormController.onSubmit` it goes to it's prototype to call `authenticate` with the context of `LoginFormController`. That's why you can access `this.password, this.username` inside AuthController. Binding **_this_** inside `handleResponse` makes sure that we are not going to lose **_this_** binding. **_This_** inside handleResponse still points to `LoginFormController` that's why we can call `displayError`
