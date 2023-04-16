# Front-End & JavaScript Questions

> [From this course](https://frontendmasters.com/workshops/frontend-js-questions/).

## Question 1

- **By default, the scripts are blocking the rendering and parsing of the HTML, even when they are downloading**.

  - That is why in the past some snippets told you to put the `<script>` tag at the very end of the HTML.

- The **async attribute** causes the script to **not block the parsing while it is downloaded**.

  - The parsing is blocked when the script executes.

- The **defer attribute** causes the script to **be executed when the whole HTML is parsed, it also does not block when the script is downloading**.

- When you provide both `async` and `defer`, the browser will have to pick one. Most likely it will be `async`.

  - I was unsure how the script with both `async` and `defer` works. Now I know that it will behave as if it was `async`.

## Question 2

- TIL that there is something called **CCSOM Tree** which is **basically a DOM Tree but for CSS**.

- The **combination of DOM & CSSOM Trees (with only visible elements) creates a _Render Tree_**.

## Question 3

This question is all about the DNS, and I know almost nothing about these things.

- **First**, the browser **sends the request to the DNS resolver for the IP address**.

- Then the **DNS resolver does a recursive search** to get the **_Authoritative Name Server IP_**.

  - This _Authoritative Name Server_ has the IP entry for the website.

- Keep in mind that this **dance of DNS resolving is cached, all depending on the TTL settings**.

## Question 4

A classic question of "what gets logged". This one requires the knowledge of the event loop, which behaves differently in different browser engines.

- Keep in mind that **`Promise.resolve()` is synchronous**. It will instantly invoke the `then` function.

- The **function you pass to the `Promise` constructor is also synchronous**.

- The `new Promise(() => console.log("foo"))` will log the `foo` instantly, no delay.

- **As long as you remember that the _micro-tasks_ are executed in-between the turns of the event loop, you should be good**.

  - This means that the _timers_ are executed AFTER the micro-tasks (mainly Promises).

## Question 5

- The **`dns-prefetch` performs domain name resolution in the background**.

- The **`preconnect` takes the `dns-prefetch` a bit further as it also does the TCP/TLS handshake**.

  - Make sure you are **not abusing these attributes**. It takes bandwidth to do these things and, if congestion occurs, **some critical resources might need to wait before they get the chance to perform these actions**.

- The **`prefetch` is to tell the browser to download the resource for later**.

  - The **browser might ignore your request**. If the network bandwidth is not there the _hint_ might be ignored.

- The **`preload` is a more aggressive version of the `prefetch`**. This is where you tell the browser it _must_ download a given resource.

## Question 6

This one showcases the issue with the _spread operator_ – the fact that it is a shallow copy and will preserve the references of nested values. Brings the old memories of writing reducers in Angular where a library like [_immer_](https://github.com/immerjs/immer) did not exist.

## Question 7

This one is about the order of the [_navigation performance-related events_](https://developer.mozilla.org/en-US/docs/Web/API/Performance_API/Navigation_timing). This is usually something you get from the API docs.

It is quite fascinating how many ways there are to measure performance of the website. I was unaware this API exists.

- It seems to me like the `loadEventEnd` is a good event to perform some work when you **want to ensure that all the resources are loaded**.

## Question 8

- The `no-cache` header is a bit misleading to me. I would imagine it would not cache anything (or touch the cache at all), but it **will still validate the cache response with the origin server before using it**.

  > It will always make the request to the origin server BEFORE using the resource. If the resource server responds with 304, the browser will use the resource it has, otherwise it will use the one from the resource server.

  - Use the **`no-store` value when you do not want to cache at all!**.

Overall, I think the knowledge of cache headers is quite important. It is very easy to get them wrong or forget that they even exist.

## Question 9

Interesting history lesson here. It **used to be the case that if two objects referenced each other, you could introduce a memory leak**. Older browsers were using **reference counting** as the heuristic to determine if a given piece of memory can be garbage collected.

This is **not the case now**. The instructor states that **if there are no references to items in the global execution context, the browser will GC those items** (even if they have a reference to each other).

## Question 10

This question touches on the topic of **"optimized" and "non-optimized" animations**. It is worth knowing that **animating properties like `width` or `left` might cause freezes and stutters**. There are websites that will tell you which property is "optimized" for animation (AKA will not cause the layout work). [This MDN page is a good resource](https://developer.mozilla.org/en-US/docs/Learn/Performance/CSS#animating_on_the_gpu)

## Question 11

- There are two phases of even propagation: the **capturing and the bubbling phase**.

- **By default** the `addEventListener` **will invoke the callback in the context of the bubbling phase**.

  - You can **change this behavior by using the 3rd parameter of `addEventListener`**. Either by providing an object with a `capture` property or a boolean value.

- The **event flow is bidirectional – it first goes to the target (capturing), then back from the target to the root document (bubbling)**.

## Question 12

This one is about CSS specificity. I'm not sure how useful it is to memorize those rules. There are bunch of sites that will tell you the specificity of a given selector. I think that the most important thing is to know that _CSS specificity_ is a thing and that some rules can override each other.

## Question 13

- The **difference between the `Map` and `WeakMap`** is how it "handles" references and what can be a key.

  - You can think of the **`Map` as copying over the values into it's own memory**. Even if you **remove the underlying value, the `Map` will still hold onto that value in the memory**.

  - The **`WeakMap`** works more like it was **holding references to the underlying values. If you "delete" the underlying value, the `WeakMap` will also remove it**.

```js
var k1 = {a: 1};
var k2 = {b: 2};

var map = new Map();
var wm = new WeakMap();

map.set(k1, 'k1');
wm.set(k2, 'k2');

k1 = null;
map.forEach(function (val, key) {
    console.log(key, val); // You can still access `k1` value, even though we "removed it" from the execution context.
});

k2 = null;
wm.get(k2) // undefined. The value is gone.
```

## Question 14

A question about [_Web Vitals_](https://web.dev/vitals/). There are a lot of acronyms and I'm unsure if it's worth remembering them. I mean you can always look it up in the docs.

## Question 15

This one is about **CSP header** value. I have to admit I was unaware that such thing even existed before I jointed Stedi. It is very interesting to me that one can **limit the origins the browser can load the resources from**. This is a great boost to the security posture of your application, **but it also can be a pain in the ass sometimes**.

The strategy seem to disallow all origins by default (by providing the `default-src: "none"`) and then extend the various resource types with hand-picked origins.

You can [read more about CSP here](https://developer.mozilla.org/en-US/docs/Web/HTTP/CSP). Notice the "weird" syntax where each "rule" is separated by `;`.

## Question 16

- If you do not specify some attributes on the `a` tags, and the page your link is pointing at is malicious, you might help that malicious site hack your users.

  - By default, though some browsers like Chrome are already removing the access to those properties, you can check the origin the user came from to your website. This is done via the `document.referrer` (there is also `document.opener`).

  - You **should be setting the `rel="noreferrer"` on your links to ensure that the target website cannot gain any information where the user came from**. This is a good practice from the security and privacy perspective.

- If you specify the `noreferrer` it also implies the `noopener`.

## Question 17

Ahh, JS generators. Such an underused and underappreciated feature of the language.

Since I'm not that familiar with them, I was unsure about the answer, but I turned out to be the correct one. One thing that caught me off guard was assigning to the `yield`, like so.

```js
function* gen() {
  const result = yield;

  console.log("You passed", result);

  return "done":
}

const it = gen();
it.next() // undefined
it.next("42") // "You passed 42" (iterator ends)
```

## Question 18

- `all` will resolve all promises but until a point of first rejection.
- `race` – will resolve the first promise that resolves or rejects.
- `any` – will wait for the first promise that resolves successfully. If none of the promise resolve, then it will reject.
- `allSettled` – will resolve all promises and then give you either a value or a rejection reason for each of them.

## Question 19

- The **`bfcache` is a cache of pages the user visited in a given session**. The cache is there to **allow snappier "go back" navigation**.

- There are some **events that will make a given page unfit for `bfcache`**. One of them is **the `unload` event**.

  - Keep in mind that **this is not the same as `beforeunload` event that you might use to warn the user that they have not saved their work**. You cannot cancel the `unload` event.

    - The `beforeunload` and the `unload` is not that reliable. The end-solution to the problem where the user might forget to save their work is to have auto-save functionality with versioning.

## Question 20

This question is about various attacks a malicious party might try. The most famous one would be XSS or CSRF. Luckily for us, the XSS is pretty much impossible if you use a framework.

- Use "same-site" cookies as they are very secure and should be the default

- Understand the mechanism of the _CSRF token_.

  - This is a token which is used to validate that the requests are coming from a legit client and are not _forged_ by an attacker.

  - First, the server generates such token and sends it alongside other data in the cookie. The client has to make the requests with that token for the server to accept their request.

## Question 21

This one is about loading fonts and the `font-display` attribute. There are a lot of them and what value you specify for `font-display` could influence how "fast" your page feels.

- I did not know that the `block` value renders an **invisible font**. Interesting.

  - I have not seen this in the wild for ever? This is mostly because the **`auto` is the default**.

- You most likely want to use `swap` with custom CSS to match the default font as close as possible to your custom font.

## Question 22

- The **`HttpOnly` on the cookie DOES NOT mean that the cookie will only be set for HTTP domains**.

  - It means that the **client cannot access that cookie**. The `HttpOnly` makes the cookie _server-side_ only.

## Question 23

- The `first-of-type` will target all elements that match this rule of different nesting levels.

    ```html
    <div>
      <ul>
        <!-- ul:first-of-type > li:first-of-type -->
        <li>foo</li>
        <ul>
          <!-- ul:first-of-type > li:first-of-type -->
          <li>bar</li>
        </ul>
      </ul>
    </div>
    ```

- If you want to target the "real" first occurrence, use the `first-child`

## Question 24

- This question is about the `Strict-Transport-Security` header and something called **HSTS**. I was completely unaware that such thing exists.

- With **HSTS you can guarantee that the browser will automatically redirect the user from HTTP to HTTPs**. This could save you some work on your server as it is the browser doing that redirect.

  - When you **perform the initial redirect on the server, attach the `Strict-Transport-Security` header to the response and the browser will handle the rest**.

## Question 25

This one was about rendering layers and how they are created. The browser creates those layers to promote the animations to GPU or to enhance performance – imagine your `position:fixed` element causing the layout re-calculation on scroll, that would be horrible for performance!

## Question 26

This one was about various image formats. The web progressed a LOT from the `.jpeg` days – there are a lot of choices.
It appears that the `AVIF` if the image format of the future?

## Question 27

This one was about CORS. I have to admit, I'm still not so well educated on this subject.

- To allow cookies with requests, you have to explicitly denote that in the CORS headers.

- Keep in mind that the "CORS request" is the "OPTIONS" request. **Not all requests require the preceding OPTIONS request**.

  - If the [request is _simple_](https://developer.mozilla.org/en-US/docs/Web/HTTP/CORS#simple_requests) the browser will not bother with the "OPTIONS" request.

## Question 28

Another question about the event loop. The takeaway here for me is that **you can block the queue if a microtask (promise-queue) schedules another micro-task**. Keep in mind that some things run synchronously, like the `defer` in Go.

- The `Promise.resolve()` is synchronous

- The `Promise(callback)` will call the `callback` synchronously.

## Question 29

- HTTP 1.1 can only do one request over a given TCP connection at a given time. You have to have multiple TCP connections for multiple resources.

- HTTP 2.0 can do **multiple requests over the same TCP connection**. This means that there are less TCP-related resources to manage.

- The HTTP 3.0 can do invents a new protocol called **QUIC**. This greatly increases performance as you can do multiple requests across multiple connections.

## Question 30

This one was about the `this` keyword and its value under different contexts. I find the golden rule Kyle Simpson thought me here very valuable.

> The only thing that matters is how the function is called.

Do you see any object to the left of the `.`? Then that will be the `this` value. If not, fallback to global.
