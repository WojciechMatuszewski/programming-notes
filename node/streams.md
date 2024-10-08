# Streams

Streams are pretty fun. I somehow was unaware of them until... now! So let's get started

## Append-only, but you can trick the system

Streams are _append-only_.
You **cannot change the data chunk once you send it to the stream**.
What you can do, is to create an _illusion_ of _out-of-order streaming_ via JavaScript.

```js
controller.equeue("<span id = 'gretting'>Hi</span>");

controller.equeue("<span>Wojciech</span>");

controller.equeue(`
    <script>
        document.getElementById('greeting').textContent = 'goodbye';
    </script>
`);
```

Given the stream above, it will look as if the first chunk changed. That is not the case.
We only changed the contents of the chunk via JavaScript, but since the chunks are sent immediately after each other,
to the user, it will seem as if we "streaming out of order."

---

A side note: this is how `React.Suspense` works.
It will render the fallback whenever possible, and continue to render the rest of the HTML.
Then, when the component is no longer suspended, it will update the "placeholder" with the right view via JavaScript.

## EventEmitter

Streams are so called _EventEmitters_. These are structures which expose events which you can listen to. An example of
such _EventEmitter_ would be the result of calling the `http.get`.

```js
const http = require("http");
const request = http.get("foo.com");

request.on("error", console.log);
```

Notice the `on` property and the name of the event. There are multiple of those events.

## Readable/Writable streams

So streams allow you to read from one data source and write to another, you can event transform the data before writing
it! This is where the notion of the `pipeline` or much older API of `pipe` comes in.

```js
const fs = require("fs");
const readable = fs.createReadableStream("foo.txt");
const writable = fs.createWritableStream("bar.txt");

const { pipeline } = require("stream");
pipeline(readable, writable, console.log);
```

Here we are just transferring the contents of `foo.txt` to `bar.txt`. The API is pretty simple and the whole process
pretty straightforward.

## Creating your own streams

There are multiple ways, one way might be through prototype inheritance with `util`. Or you could just use
provided `Readable` and `Writable` instances.

This is how I would create a simple `Readable` stream

```js
const { Readable } = require("stream");

let reads = 0;
const myStream = new Readable({
  read() {
    reads += 1;

    // Do not return, the `null` indicate that you are done
    if (reads === 0) {
      this.push(null);
    } else {
      this.push("SOME_CONTENT");
    }
  },
});
```

Notice the `push` method. Pretty weird, my first instinct was to just `return` stuff.
Also, one might wonder how to preserve `null` as the normal value that can be returned and not the indication that the
stream is finished.

## HTTP Requests

The `response` parameter that your `http.request` callback takes is also a _ReadableStream_. This might come in handy
when reading from a request.

```js
const request = http.request("foo.com", (response) => {
  pipeline(response, writable, console.log);
});
request.on("error", console.log);
```

You would mostly use that on the server though to _stream_ the response.

## `async iterable` streams

This one is great, especially for any kind of _transformers_. Instead of using the `new Transform` syntax, you just use
_generators_

```js
const { pipeline } = require("stream");

async function* upperCase(readable) {
  for await (chunk of readable) {
    yield chunk.toString().toUpperCase();
  }
}

pipeline(process.stdin, upperCase, process.stdout, (err) => {
  if (err) {
    console.log(err);
  }
});
```

Pretty easy, right?

## Fetch and streams

Given that the `fetch` API is now built-in to Node.js (via the `unidici` package), you might not need to use third party
libraries to fetch some data. Now that we have more-or-less the same API on the browser and on the server for fetching
data, let us consider how `fetch` API works with _streams_.

First, know that **when `fetch` returns with `Response` it does not mean that the server "finished" working on the
request**. The **fetch will respond with `Response` as soon as the server sends headers**.

This means, that the `Response` could be used as a stream. In the worst case scenario, we will get one single chunk with
the whole payload (in such case, the server does not support streaming or is setting incorrect headers), OR we will be
able to read the data chunk-by-chunk.

Here is an example of how one might read the `Response` as stream.

```js
const response = await fetch("...");
const readable = response.getReader();

const chunk = await readable.read();
```

## For-of-loop

**In Node** (it does not seem to be working in the browser for me), you can use the `for await (let chunk of stream)`
syntax to consume the stream.

```js
const stream = new ReadableStream({
  start(controller) {
    controller.enqueue("foo");
    controller.enqueue("bar");
  },
});

for await (const chunk of stream) {
  console.log(chunk);
}
```

Pretty neat!

## Consuming streams via `.on` callbacks

**For the `end` event to fire, you have to bind the `data` event callback**. [This is mentioned in Node.js documentation](https://nodejs.org/api/stream.html#event-end).

```js
const readable = new Readable({});

readable.on("data", () => {});

readable.on("end", () => {}); // will not fire unless you registered the `data` event callback.
```

So, if you see your stream "hanging" and seemingly not working, check if you have the `data` callback registered.

## Web Streams vs. Node streams

> More info [here](https://www.platformatichq.com/node-principles#understanding-native-node.js-apis-vs.-web-standard-apis)

For the longest time, in JavaScript, streams only existed in Node.js. **With the introduction of the `fetch` API, browsers started to implement so-called _web streams_**. As you can imagine, this lead to having two, similar in functionality, but different in terms of API, ways of handling streams.

**At the time of writing, one can use `Readable.toWeb` or `Readable.fromWeb` in Node.js**. This is quite nice, as it makes for unified experience across the web and the terminal.

While I only dipped my toe in the world of streams, I can already see a couple of differences.

1. Pausing _Web Streams_ does not seem to be possible.

Here is how you would read a file line-by-line via _web streams_ in Node.js

```js
async function* webStreamsFileGenerator() {
  const filePath = fileURLToPath(import.meta.resolve("./file.txt"));
  const fileStream = Readable.toWeb(Readable.from(createInterface(createReadStream(filePath))));

  for await (const chunk of fileStream) {
    yield chunk;
  }
}
```

**Notice that we do not have to add any separate code to read chunks of the `fileStream`**. the `fileStream` does not have the `pause` method.

Now, consider the code that does the same thing, but uses Node.js streams.

```js
function getChunkFromStream(stream) {
  return new Promise((resolve, reject) => {
    stream.once("data", (chunk) => {
      stream.pause();
      resolve(chunk);
    });

    stream.once("end", () => {
      resolve(null);
    });

    stream.once("error", (error) => {
      reject(error);
    });
  }).finally(() => {
    stream.resume();
  });
}

async function* nodeStreamsFileGenerator() {
  const filePath = fileURLToPath(import.meta.resolve("./file.txt"));
  const fileStream = Readable.from(createInterface(createReadStream(filePath)));

  let chunk;
  while ((chunk = await getChunkFromStream(fileStream))) {
    yield chunk;
  }
}

const generator = await nodeStreamsFileGenerator();
for await (const chunk of generator) {
  console.log(chunk);
}
```

**Notice that I had to `pause` the stream in order to create a generator**. Without the `pause` method on the stream, creating that generator would be quite hard. **Node.js streams seem to expose more functions on the stream itself**.
