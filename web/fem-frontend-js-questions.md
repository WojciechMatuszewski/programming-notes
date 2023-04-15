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

Finished part 2, 39:29
