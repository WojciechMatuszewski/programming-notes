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
