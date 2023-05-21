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

Finished part 3 52:28
