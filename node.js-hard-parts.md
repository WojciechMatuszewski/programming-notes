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

## How does `c++` work with `JavaScript`?

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
