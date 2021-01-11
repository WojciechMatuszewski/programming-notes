# Revisiting Async with Kyle Simpson

## Callbacks

```js
function fetchCurrentUser(function onUser(user) {
    var otherOrders;
    fetchArchivedOrders(user.id, function onResp(orders) {
        ....
    })
    // RACE CONDITION \/ /\
    fetchCurrentOrders(user.id, function onResp(orders) {
        ...
    })
    ...
})
```

To understand above code we have to go through all the possible scenarios (which are plenty). This code scales very poorly. The good old callback hell :).

### Thunk pattern

Thunk is a function that has all that it needs to compute it's task. You just need to call it. It's deferral of future computation (think like defer in RxJs). It's a way of describing 'future' stuff, **representation of future value in time independent fashion** Thunk is using **closure to remember it's internal state**.

```js
// th_xx => Thunk

// Thunk making thing (fetchCurrentUser)
var th1 = fetchCurrentUser();
th1(function onUser(user) {
    // concurrency
    var th2 = fetchArchivedOrders(user.id);
    var th3 = fetchCurrentOrders(user.id);
    th2(function onArchivedOrders(archivedOrders) {
        th3(function onCurrentOrders(currentOrders) {
            // both of them has finished
        });
    });
});
```

## Promises

Remember Thunks? Promises are also **time independent representation of future value**.

```js
fetchCurrentUser()
    .then(function onUser(user) {
        return Promise.all([
            fetchArchivedOrders(user.id),
            fetchCurrentOrders(user.id),
        ]);
    })
    .then(function onOrders([archivedOrders, currentOrders]) {
        // ...
    });
```

Promise API is better than callbacks but it also has problems:

- overloading => polymorphic nature of `.then` (passing different things to a `.then` behaves differently depending on the value passed)

- **handlers do not scope** => this one is huge. It encourages putting stuff in outside scope which creates a lot of problems

- swallowed rejections

Firstly, remember that `.then` has actually 2 parameters

```js
function getPromise() {
  return new Promise((resolve, reject) => reject("ERROR"));
}

getPromise()
  .then(
    () => {
      console.log("inside then");
    },
    // this one fires
    err => console.log("im here")
  )
  // this one is ignored
  .catch(err => console.log("err"));
// im here
```

When an error happens in a given step it has to be caught in a next step. This can be misleading.

Kyle thinks we should not use Promises (an API) but the concept.

## Generators

It's a function that can cooperate with another part of the program.
Generators are another example of lazy computation. They do not actually run when you call them. **They create an iterator on being called**.

```js
function* main() {
    yield 1;
    yield 2;
    yield 3;
    // you can also return values
    return 4;
}
var it = main();

it.next(); // all that iterator jazz {...}
```

Another example

```js
function* main(max = 8) {
  for (let i = 0; i <= max; i = i + 2) {
    yield i;
  }
}
[...main()]; //spread consumes iterators
// you can also pass values just like to a normal function
for (let v of main(14)) {
    ...
}
```

`yield` expression is like a placeholder for a value.

```js
function* main() {
    var x = 1 + yield;
    var y = 2 + yield;
    var z = 3 + yield;
    return x + y + z;
}
var it = main();

it.next(); // {value : 4...}
it.next(10); // {value : 5} BUUUT yield 5 got replaced with 10
it.next(20); // {value: 6} BUUUT yield 6 got replaced with 20
```

### Async with generators

```js
function* main() {
    var pr = ajax("...");
    var v = yield pr;
    console.log("im here", v);
}

var it = main();
var pr = it.next().value;

pr.then(function(v) {
    it.next(v);
});
// logs => im here , PROMISE_RESPONSE
```

### Gen runner pattern

A lot of well known libraries implement such pattern (bluebird is a good example). Sadly he did not show how to implement such runner, we probably should use an library.

```js
runner(function* main() {
    var user = yield fetchCurrentUser();

    var [archivedOrders, currentOrders] = yield Promise.all([
        fetchArchivedOrders(user.id),
        fetchCurrentOrders(user.id),
    ]);
    // ...
});
```

Using a library for gen runner is a big trade-off for some people. What can we do about it ?

## Async await

This is exactly **the same thing as gen runner but without an library**. V8 literally transpiles your code into generator and pumps it into gen runner 😂

Async functions implicitly return promise.

```js
async function main() {
  return 42;
}

main().then(...)
```

### Limitations

- async / await are shallow.
  That means that as soon as you wrap a regular function around the code where you are using `await` you can't use `await` keyword.

```js
async function fetchFiles(files) {
    var prs = files.map(getFile);

    prs.forEach(function each(pr) {
        // nope sorry, shallow continuation
        console.log(await pr);
    })
}
```

- `await` only works with Promises (RIP Thunks)
- scheduling (starvation)
  Starvation means that one part of your system accidentally or maliciously can zap all the available resources and prevent other stuff having a chance to run. **You can starve out any other macro task in the system using async/await**.
- external cancellation, you really do not have any way to cancel async / await stuff

## await for ... of

You can actually create async loop with the help of generators.

First, lets create the generator:

```js
async function* fetchUrls(urls) {
  for (let url of urls) {
    let resp = await fetch(url);
    if (resp.status == 200) {
      let text = await resp.text();
      // BOXING !!!
      yield text.toUpperCase();
    } else {
      yield undefined;
    }
  }
}
```

### Consuming async generator with `while`

```js
var it = fetchUrls(favSites);
async function main(varSites) {
    while (true) {
        let res = await it.next();
        if (res.done) break;
        let text = res.value;
        console.log(text);
    }
}
```

### Consuming async generator with `for await`

This is basically the same as the `while` loop. Just some nice syntactic sugar.

```js
async function main(favSites) {
    for await (let text of fetchUrls(favSites)) {
        console.log(text);
    }
}
```
