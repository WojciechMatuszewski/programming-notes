# Streams

Streams are pretty fun. I somehow was unaware of them until... now! So let's get started

## EventEmitter

Streams are so called _EventEmitters_. These are structures which expose events which you can listen to. An example of such _EventEmitter_ would be the result of calling the `http.get`.

```js
const http = require("http");
const request = http.get("foo.com");

request.on("error", console.log);
```

Notice the `on` property and the name of the event. There are multiple of those events.

## Readable/Writable streams

So streams allow you to read from one data source and write to another, you can event transform the data before writing it! This is where the notion of the `pipeline` or much older API of `pipe` comes in.

```js
const fs = require("fs");
const readable = fs.createReadableStream("foo.txt");
const writable = fs.createWritableStream("bar.txt");

const { pipeline } = require("stream");
pipeline(readable, writable, console.log);
```

Here we are just transferring the contents of `foo.txt` to `bar.txt`. The API is pretty simple and the whole process pretty straightforward.

## Creating your own streams

There are multiple ways, one way might be through prototype inheritance with `util`. Or you could just use provided `Readable` and `Writable` instances.

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
Also one might wonder how to preserve `null` as the normal value that can be returned and not the indication that the stream is finished.

## HTTP Requests

The `response` parameter that your `http.request` callback takes is also a _ReadableStream_. This might come in handy when reading from a request.

```js
const request = http.request("foo.com", (response) => {
  pipeline(response, writable, console.log);
});
request.on("error", console.log);
```

You would mostly use that on the server though to _stream_ the response.

## `async iterable` streams

This one is great, especially for any kind of _transformers_. Instead of using the `new Transform` syntax, you just use _generators_

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

Pretty easy right?

## Fetch and streams

Given that the `fetch` API is now built-in to Node.js (via the `unidici` package), you might not need to use third party libraries to fetch some data. Now that we have more-or-less the same API on the browser and on the server for fetching data, let us consider how `fetch` API works with _streams_.

First, know that **when `fetch` returns with `Response` it does not mean that the server "finished" working on the request**. The **fetch will respond with `Response` as soon as the server sends headers**.

This means, that the `Response` could be used as a stream. In the worst case scenario, we will get one single chunk with the whole payload (in such case, the server does not support streaming or is setting incorrect headers), OR we will be able to read the data chunk-by-chunk.

Here is an example of how one might read the `Response` as stream.

```js
const response = await fetch("...");
const readable = response.getReader();

const chunk = await readable.read();
```
