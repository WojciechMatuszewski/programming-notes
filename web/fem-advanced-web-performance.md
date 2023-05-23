# FEM Advanced Web Performance

Taking notes while watching [this course](https://frontendmasters.com/workshops/advanced-web-perf/)

## What is the problem

- The **median** of time-to-interactive on mobile is 12.5 seconds. 12.5 seconds!

  - That is so long. Try timing it with a clock.

- The slower your page, the most likely the users are to bounce.

- The slower the cellular network, the bigger the latency.

  - Most of the parts of the world is 4G/3G.

- Before you embark on improving performance, **measure first**.

- There are two categories of metrics that pertain to performance.

  - The _user-centric_ metrics.

    - These measure **how fast things appear on the screen and how fast the user can interact with the page**.

    - How user perceives the performance.

  - The _browser-centric_ metrics.

    - These measure **lower-level latencies, like TTFB**.

- Keep in mind that **you can define your own business-centric metric**. Like _time-to-first-tweet_ or similar.

  - It is good to have those kind of metrics so you can base OKRs on those.

## Tools and Charts

- Most of the times you will be looking at the waterfalls charts.

- The most accessible tool is the Chrome developer console.

  - There is the network waterfall chart and also the "performance" chart with screenshots.

- There is also webpagetest.org. This one is **the most advanced tool to test performance with real devices**.

- If you analyze the charts, **you will come to the conclusion that most of the time waiting is spent on the frontend parsing assets**.

  - It is your responsibility, as a FE developer, to optimize that part.

- **Improving performance WILL help your bottom line**.

## What happens when we browse the web

- Before we even start downloading assets, the browser has to open a lot of connections and perform the DNS resolution to get the IP for your website.

  - This takes time. Of course some of those operations are cached on subsequent visits.

    > <www.frontendmasters.com> (the times are only an example)
      1 – DNS query       –   100ms
      2 – TCP connection  –   120ms
      3 – SSL Negotiation –   150ms
      4 – HTTP Request
          header – body
      6 – Server process(backend time)    – 200ms to 500ms
      7 – Browser HTML parsing
      8 – Resource Discovery & Priority
      9 – Render (layout, paint)

- Max gives a really interesting trivia background on why we push on HTTPs so much nowadays.

  - One aspect is security. That is given.

  - The other is about evolving the spec. Some old routers, when they do not know the protocol, they will modify the packet and discard the protocol information. Since the packet is encrypted with HTTPs, they will not be able to "see inside" that packet and they will forward it. Having the packets encrypted allows those routers to be compatible with changing the spec.

### Different specs of HTTP

- There is the **HTTP/1.1**.

  - You can only make **one request per TCP connection**.

  - There is a **limit of parallel request to the server**.

  - You can encode stuff with gzip.

- There is the **HTTP/2**. Created with performance in mind.

  - It compressed headers.

  - **You can re-use TCP connections**. This is a great optimization.

- There is the **HTTP/3**.

  - Even faster than HTTP/2.

  - It **uses UDP instead of TCP**. There is a **thin level of abstraction on top of the UDP**.

### Browser Cache

- Browsers populate the cache with downloaded resources. **It is the headers that come with each resource that control the browser cache**.

- When the resource is expired, the browser will ask the server if the resource changed. If it did, the browser will download the resource again. If it did not, it will serve the resource from the cache.

  - Note that **this dance of asking the server if the resource is valid is a separate request that does not contain the resource**.

    - This is why you **most likely want to set very long cache headers and then when the resource changes, "bust" the cache with cache headers**.

- There is also the **back/forward cache (bfcache)**.

  - This cache is responsible for caching the whole views when you are going back/forward in the browser.

  - **Some event listeners might make your website incompatible with bfcache**. One of such events is the `unload` event.

    - Keep in mind that **the `beforeunload` event is compatible with bfcache on Chrome and Safari, but not on Firefox**.

### Server Worker

- If you utilize caching, the service worker can act as a "local" server for the page assets.

  - There are multiple ways to go about it. I would personally recommend the [workbox library](https://developer.chrome.com/docs/workbox/the-ways-of-workbox/).

## Parsing & executing resources

- **By default, executing the `script` tag will block the browser parsing further HTML until the script is executed**.

  - This is pretty bad for performance.

    - Nowadays we have the `defer` and `async` attributes available to us.

- **By default, parsing the CSS will block rendering**.

## Basic Performance Optimizations

- Enable GZIP on text-based files.

- Make static content expire late in the future.

  - So that the browser does not have to ask the server if the files are valid or not. Bust the cache via hashes appended to files.

- Use CDN for static content.

  - So that the assets are "closer" to the user.

- Consider avoiding putting a lot of data in the cookies.

  - Browser sends the cookies with **every request**.

- Avoid redirects. They force the browser to "start from zero" in the process on rendering the page and resource discovery.

- JavaScript **is very expensive to run when you think about it**.

  - Defer or remove it as much as possible.

  - Compress and obfuscate it. It improves the filesize. Remember about bundling everything together.

- Embrace responsive images.

  - Do not serve very large images on phones. Serve the size the image the user needs.

    - You have the `picture` and `srcset` at your disposal!

## Hacking Performance

### Hacking first load

- Avoid more than one roundtrip.

  > Here Max goes on a tangent related to units of measurement. There is a confusion that 1 KB is 1024 Bytes, but in fact it is 1000 Bytes.
    The 1024 Bytes is 1 KiB. Max traces this to Windows that does not adhere to the official units of measurement and uses KB (and other similar units) in the place of KiB (and other similar units).

  - **On most Linux-based systems**, the **maximum packet size is 14.6 KiB**. If you can **fit the "above the fold" HTML and assets there**, you will benefit greatly in terms of performance metrics.

    - The reason is that, in a TCP connection, packets might arrive out of order. Consolidating those takes time. If you send a single packet, there is nothing to order.

- Use the **HSTS header**.

  - This tells the browser that, no matter the URL on a given domain, it is always safe to redirect to HTTPs.

### Hacking LCP

- **Help** the browser **discover assets earlier** by using the **rel=`preload` syntax** on links and images.

  - This applies to images, fonts, stylesheets.

  - **Do not overuse this**! If you add "preload" to everything, it will not work well.

- Hand-pick the **`fetchpriority` for very important assets**.

  - Again, setting everything to the highest possible priority does not work. It makes everything "baseline" again.

  - The **browser already sets the `fetchpriority` implicitly for assets**.

  - You can also **help the browser and change the defaults by letting the `fetchpriority` to `low` on some assets**.

- Utilize the **HTTP Early Hints**.

  - Instead of the browser waiting for the server to respond with HTML, the **server can, while working on the HTML, send the browser instructions about what assets the browser will need**. The browser **can start downloading those assets meanwhile the server works on the response**.

### Hacking Data Transfer

- Utilize the HTTP/3 standard.

- Use Zopfli.

  - It **takes more time to compress than gzip, has the same efficiency when decompressing and produces archives that are about 8% smaller than gzip**.

- Use Brotli

  - About **25% better than gzip**.

### Hacking Resource Loading

- Use modern image formats like WebP, AVIF and so on.

- Use the service worker with _stale-while-revalidate_ strategy.

  - Serve from the cache, but in the background, ensure the asset is up to date. If it's not, replace it on the next load.

  - Nowadays, you can use **the _stale-while-revalidate_ header value**.

- Use the `preconnect` of `dns-prefetch` to save up time when the browser requests given asset.

  - Keep in mind that **`preload` will also fetch the file**. Sometimes **you do not want to `preload` the files, especially since they will be shown later in the page lifecycle**.

- Use the **native `loading="lazy"` for images**.

  - It used to be that you needed a JS library for this. Nowadays it is built-in inside the browser.

- **Change how the fonts load via the `display` property**.

  - There are many values available to you. Play with them to avoid the _Flash of Unstyled Text_.

## Hacking Interaction Experience

- **Avoid client-side rendering**.

  - As nice as it is, it is hurting your performance in a big way.

  - Keep in mind that Next.js will do SSR on the first load, then do client-side navigation.

- Move heavy tasks to WebAssembly and workers.

- **Stop serving legacy code if you can**.

  - This one is huge. See what kind of code the bundler you use produce. Maybe you are still on ES5?

## Using Performance APIs

- Using RUM tools like DataDog.

- Using the `performance` object.

  - The `performance.timing` **contains a lot of timestamps related to the page lifecycle**.

- There is the `PerformanceObserver` object. It works similar to the `IntersectionObserver`, but for performance-related timings.

  - It can notify you when the browser decided what was the "first-input delay" and so on.

## Animation Frames

- It is not guaranteed that the timers run by the time you specified in the function.

  - They are put in a specific queue. If we block the event loop, the timer might never fire.

- If you want to execute code **at the next frame, use the `requestAnimationFrame` API**.

- There is also `requestIdleCallback` API as well.

- You can **split intensive work in-between idle moments** by **`requestIdleCallback`**.

  - It **even has the `timeout` property**. A nice API to control the flow of synchronous job.

- There is also the **`setImmediate`** where **you tell the browser to execute your code at the end of the event loop, but before the timers**.

## Summary

- The web performance is also about your bottom line.

- There are a LOT of tools to optimize web performance.

- **First measure things**.

- It is vital to understand how browsers cache works.

  - Setting a very long expiry time on assets makes sense if you include their hash in the name.

  - Mind the bfcache and the implications of some event listeners on it.

- CSS will block HTML rendering.

- JS will bloc HTML parsing (worse than CSS).

- If you tried to solve performance problems long time ago, it has gotten much easier.

  - You no longer have to lean on external libraries that much. A great example is the `loading="lazy"` attribute on images.
