# Node.js hard parts

## The very basics

Things needed to run webpages (excluding the obvious stuff)

- html
- javascript
- css
- data

These things are not on your computer by default. These things are located on the `server`.

But how does the `server` knowns what to send back? It's programmed using some kind of programming language.

To be able to send some message back (response) we need to have access to some of the underlying features of the computer that runs the server. Pure JavaScript is not the best language for that.

Which language has these abilities? `c++` does!. So we are going to combine powers of `c++` with JavaScript to create the ultimate platform: `Node.js`.

## How does `c++` work with `JavaScript`

### Basic explanation

Someone's computer sends an `'internet message'`. It lands on twitter server (computer).

That inbound `message` arrives to `'computer internals'` [twitter computer], like networking etc, (think of this place as something JavaScript cannot access).

We are going to pull out this `message` using `Node c++ features` and `JavaScript labels` which are in reality `c++ commands`.

### Rough workflow

So based on _basic explanation_ workflow will look like this:
JS -> Node -> Computer feature (e.g network, file system)

### JavaScript features in context of the workflow

- saves data and functionality (code)
- uses that data by running functionality (code) on it
- has a ton of build-in labels that trigger Node features that are built in `c++` to use our computer's internals

## Javascript labels & http

We can set up `JavaScript` label, a `Node.js` feature to wait for requests for html / css / js etc..
How? The most powerful build in `Node.js` feature of all: http

- ht -> hypertext (structure of web-pages)
- t -> transfer (self-explanatory)
- p -> protocol (format of data, how do you communicate)

## Using http feature of Node to set up an open socket

**Socket is a word for open channel for data to go in and out**

```javascript
const server = http.createServer();
server.listen(80);
```

`http.createServer` is a built-in `Node.js` label.
Inbound web request -> run code to send back message.
**Default http port is: 80**

### How do we handle incoming _message_ ?

We do not know when the inbound request would come - we have to rely on Node to trigger JS code to run.
How do we 'store' some code to run it later? using a `function`.

```javascript
function doOnIncoming(incomingData, functionsToSetOutgoingData) {
  functionsToSetOutgoingData.end("Welcome");
}
const server = http.createServer(doOnIncoming);
server.listen(80);
```

### Incoming message and Node.js help

As soon as this message comes in, node will invoke our callback.
But we really, really do not want to handle raw _http message_ in `JavaScript`.

Instead `Node` will auto-create 2 very important objects which will help us with handling that _incoming message_.

**These objects do not have names**

#### First auto-created object

It's going to package up in nice object format important information's about the _incoming message_:

- url and more information's about the _incoming message_

#### Second auto-created object

This object has a bunch of utility-functions (_remember that these are JavaScript labels_)
Good example of such function is `end function`. Whatever we pass to that function will be used to data to the _response message_ (**this response message is auto-created by Node**)

#### How are these two objects used ?

These 2 objects will be passed to your callback. In our previous snipped we called them `incomingData` and `functionsToSetOutgoingData`.
It's important to understand that these can be called anyway we want. **The so-familiar `req` and `res` is just a convention**

## HTTP format

Three parts of _http message_ (then sending .get)

1. **Request line** (in this case GET with url)
2. **Headers**, metadata about senders browser, pc etc.. (used to know what to send back)
3. **Body** (in this case it's empty, we are sending GET request)

### Example and walkthrough

```javascript
const tweets = [...]
function doOnIncoming(incomingData, functionsToSetOutgoingData) {
  // reading tweet number from url
  const tweetNeeded = incomingData.url.slice(8)-1;
  functionsToSetOutgoingData.end(tweets[tweetNeeded])
}
const server = http.createServer(doOnIncoming);
server.listen(80)
```

`http.createServer()` does 3 things, 2 in node, one in `JavaScript`

- setups up `http` feature in node which basically opens a network channel (in computers internals) known as **socket**
- store callback to auto-run when _incoming message_ comes
- returns (in `JavaScript` land) an object which holds bunch of functions. These functions enable you to edit the background `c++/Node.js http` feature

`server.listen` has an effect in `Node.js` land. It sets the port which in turns opens an entry point to our computer

`libUV` library is used to talk to computer network features by node.

_incoming message_ is brought into `Node.js` word using `libUV`, `Node.js` auto-creates another _message_, this time this message will be used as response to the _incoming message_

In this stage `Node.js` also auto-creates 2 objects (we talked about them before), **they are passed to callback function as an arguments**

In our case these two would look something like this

```javascript
// first one known as request
var obj1 = {
  url: '..',
  method: 'GET'
}

// second object allows us to control the auto-generated response
var obj2 = {
  // basically a lot of 'labels'
  end: () => ...,
  write: () => ...,
}
```

`Node.js` is going to invoke our callback with these objects as parameters.

By using `.end` we are telling `Node.js` that we are done editing the auto-created message and then `Node.js` sends it back as a response.

_Response a.k.a return back message_ is also in http format.

## Handling errors

Errors are inevitable, how do we handle these ?
To handle them correctly we need a deeper understanding how does the background `http Node.js feature` work.

### How does `http.createServer` really work

It turns out when the new request comes in `Node.js` does not really run our function passed to `createServer` automatically. It gets run when `request` event is emitted within `Node.js`

Since this is an event, we can attach listeners to `object returned by http.createServer()`.

```javascript
function doOnIncoming(...) {}
// infoOnError is a parameter !!!!!
function doOnError(infoOnError, rawAccessToSocket) {
  console.error(infoOnError)
}
const server = http.createServer();
server.listen(80);

// attaching listeners
server.on("request", doOnIncoming);
server.on("clientError", doOnError);
```

When `clientError` event gets dispatched, `Node.js` creates another object, it's called `Error`. It contains information's about there does the error originate from etc..

That object is going to be the input to our `doOnError` function

Another input to our `doOnError` function is a raw socket (**that raw socket is NOT in http format**). Not that very useful since it's not in a http format

## `Node.js` file system

### Using fs module

```javascript
function cleanTweets(tweetsToClean) {}
function useImportedTweets(errorData, data) {
    const cleanedTweetsJson = cleanTweets(data);
    const tweetsObj = JSON.parse(cleanedTweetsJson);
    console.log(tweetsObj.tweet2);
}

// dot represents current location in file system
// node is going to firstly look up the folder in which you switched node on in
fs.readFile("./tweets.json", useImportedTweets);
```

`libUV` also helps with file-system access. It's much more involved actually. When talking about http `libUV` does not create a thread to focus on listening to incoming messages. In case of file system access `libUV` sets up a thread manually (it's done because there is too much variety in how different OS handles files)

```javascript
// this is called error-first pattern
function useImportedTweets(errorData, data) {}
```

### Streams

Reading & cleaning data in batches.
Node is using `event pattern` for batching stuff.
Now we can do two things at the same time (pulling data and cleaning).

```javascript
let cleanedTweets = "";
function cleanTweets(tweetsToClean) {
  //..
}
function doOnNewBatch(data) {
  cleanedTweets += cleanTweets(data);
}
const accessTweetsArchive = fs.createReadStream("./tweetsArchive.json");
accessTweetsArchive.on("data", doOnNewBatch);
```

Invoking `createReadStream` is like telling `libUV` to setup a dedicated thread to go and get the desired data and start pulling it.

File is divided in chunks by `Node.js`. Every 64 (you can change that, but 64 is by default) bytes `Node.js` is going to dispatch an `data` event.

But what if the doOnNewBatch takes a lot of time and we have new `data` event dispatched? How do we ensure the correct order of execution?

#### Callback queue

`Node` introduces callback queue. If the call-stack is empty `Node` is going to check the callback queue periodically. That checking is done by something also present in browsers, mainly `event loop`. Here it's `libUV` that implements that.

So the logic now is simple. We put `doOnNewBatch` on the callback queue. When call-stack is empty we are putting that callback on the call-stack and running it

#### Many different queues

```javascript
function useImportedTweets(errorData, data) {
    // parsing
    console.log(tweets.tweet1);
}
function immediately() {
    console.log("run me last");
}
function printHello() {
    console.log("hello");
}
function blockFor500ms() {
    // BLOCK JS thread DIRECTLY for 500ms
}
// remember this is not Web-API, Node has his own implementation
setTimeout(printHello, 0);
fs.readFile(".file", useImportedTweets);
blockFor500ms();
console.log("meFirst");
// setImmediate is another Node feature
setImmediate(immediately);
```

- `setTimeout(printHello,0)` goes to _timer queue_
- evaluating `fs.readFile`, nothing added nowhere yet
- `blockFor500ms` goes to call-stack
- 200ms in, `useImportedTweets` gets called (data comes back),
  `useImportedTweets` goes to _I/O callback queue_ (the one you saw earlier)
- blocking code finishes, but there is still more global code to run
- `console.log` gets added to call-stack and runs
- `setImmediate` callback fn gets added to _check queue_

You want to hear a joke?
Lets consider `setImmediate`. The name implies that it will be run, well, immediately. But behold `Node.js` logic. _Check queue_ has the least priority of all the queues so it will most certain run as the last one :D

#### Queue priority

**Excluding microtask queue and close queue**
_timer queue_ > _I/O callback queue_ > _check queue_

## Event loop

Let's summarize our knowledge of the _event loop_

- **_event loop_ has _phases_ – different queues that contain _commands_**

* With each **transition** to **another phase**, the **`process.nextTick` and the microtask (promises) are drained**. This happens to a certain extend, the _libuv_ makes sure we are not starving the _event loop_ with blocking calls
