# Web Performance Fundamentals V2

> Notes from [this course](https://frontendmasters.com/courses/web-perf-v2/introduction/)

## Importance of Web Performance

- From the users perspective, poor web performance leads to poor UX.

  - Users have certain expectations about your website based on their previous experience.

- **According to research Todd cites, 40% of users will abandon a site within three seconds**.

- The **better your site performance, the better the SEO**.

  - **Google takes web vitals into the account** when ranking pages.

## Measuring Web Performance

- To _read_ how performant your web page is, you will most likely interact with **waterfall charts**.

- There are A LOT of metrics related to web performance.

  - Some are considered _legacy_, like `DOMContentLoaded` or `load` event.

    - The `DOMContentLoaded` tells you that the HTML has loaded. Images and other interactive elements might still be missing.

    - The `load` event fires when all _known_ resources have been downloaded and rendered.

      - This event does not include _lazy-loaded_ files and dynamically rendered content.

  - The **_legacy_ metrics are problematic because they do not account for various technologies the site could be built in**.

    - When built as SPA, the websites `DOMContentLoaded` fires _very fast_, but the experience might be poor due to client-side rendering.

### Core Web Vitals

- To address issues with _legacy_ metrics, Google came up with **_Core Web Vitals_**.

  - The aim here is to create **metrics that correspond directly to UX and are independent of how the website is built**.

- Many metrics fall under the _Core Web Vitals_ umbrella

  - The `LCP`

  - The `CLS`

  - The `INP`

  - The `FID`

  - The `TTFB`

  - The `FCP`

  - Those **metrics might be different when you let the page load and do not interact with it vs. when you start interacting with the page as soon as you can**.

- You can download the [Web Vitals extension](https://chromewebstore.google.com/detail/web-vitals/ahfhijdlegdabablpippeagghigmibma) to capture those metrics easily.

## Capturing Performance Metrics

- **For custom performance metrics**, to answer questions like: "How long did this function run for?", consider using **Performance API**.

  - The **problem with the Performance API is that the act of measuring can slow things down**.

    - Adding a bunch of logs and calling a bunch of functions only to measure performance will impact the thing you are measuring.

  - To **get performance information when the browser is "idle" consider using [the `PerformanceObserver`](https://developer.mozilla.org/en-US/docs/Web/API/PerformanceObserver/PerformanceObserver)**.

- **Google has a `web-vitals` package** that **handles all the edge cases of using the `PerformanceObserver` and the Performance API for you**.

  - If you want to measure _web performance_, this is usually the best choice.

- **Interestingly, some web engines do not support all the metrics**.

  - The `Webkit` engine used by Safari, does not support `LCP`, `CLS`, and `INP`.

  - The `Gecko` engine used by Firefox, does not support `CLS` and `INP`.

  - **This means that you should measure on a browser using the `Blink` engine as it supports all the metrics**.

    - If you measure only on `Webkit` or `Gecko`, the score you get might be good, but, in reality, it is much lower because it does not include all the metrics.

## Testing Performance

- The **best performance metrics you will get are from the real, organic traffic**.

  - Tools like Datadog RUM and similar providers might help here.

- **If you want to test performance locally, ensure you throttle your CPU and network to appropriate levels**. You can do that in dev console.

  - You most likely have much more powerful machine than most of your users.

- **Google published real-world performance data from _logged-in users_ for top X websites that you can query**.

  - This data is called "CRUX" – _Chrome User Experience_ data.

## Setting Performance Goals

- _Fast_ is subjective. **It depends on the audience**.

- **There has been a lot of studies around "psychology of waiting**. To me, the most surprising facts are:

  - People will remember the "slow experience" slower than it actually was.

  - People **are willing and wait to wait for "value"**.

    - Todd gave an example of tax-filling application that specifically slows things down to make the user trust that the application is "checking all the things".

- **You, as an engineer, SHOULD NOT decide what is "fast enough**.

  - It is **your users, competitors and the SEO score** that decides whether you should invest in performance optimizations or not.

## Improving Time to First Byte

- **Compress the resources** you sent over the wire. GZIP or Brotli are good choices.

  - Brotli is usually the _better_ choice.

  - **Compression does not work that well for smaller files. Always measure first**.

    - Keep in mind that there is an overhead in compressing files.

- **Use newer HTTP versions**.

  - Nowadays, the HTTP v3 is the best for performance, but it is pretty now, so the overhead of using it might be hight.

- **Host your website closer to your users**.

  - In the era of Cloudflare workers and Edge Lambda Functions, there is no excuse in having a single location for your website.

    - **If you go this route, be mindful where you data lives**. If **your database is hosted in US, but your website sits in EU, the website <-> DB latency will be high**.

## Improving First Contentful Paint

- **CSS and Fonts are, USUALLY, RENDER BLOCKING**.

  - The `@import` statements in CSS files.

    - Theses create waterfalls, because browser first has to fetch file X, parse it, and then fetch file Y, the file X references.

    - **You can use `link rel="preload"` for critical CSS files**.

  - The `url` in `@font-face`.

    - **You can use `link rel="preload"` for font files**.

      - **Consider hosting the fonts yourself, ideally on the same domain**. This will be even faster than preloading from Google CDN.

  - Using a bundler helps here a lot. It helps because it will concatenate files together reducing the request waterfalls.

- You have two options for lazy-loading scripts.

  - Using the `script async`. This one **will download the JS file alongside other resources**. **As soon as the file is ready, it will BLOCK all parsing until the file is executed**.

    - This **can create race conditions with other resources**.

      - Think of a situation where CSS file is 90% done downloading, and the script has just completed downloading. The browser will defer downloading the CSS file until JS is done executing.

  - Using the `script defer`. This one **will download the JS file alongside the other resources and execute it only when `DOMContentLoaded` is about to happen**.

    - This means that there is **no race condition** with other files. The `DOMContentLoaded` fires when all HTML was parsed, so CSS files are already downloaded.

  - The `type="module"` scripts are _always_ deferred.

## Improving Largest Contentful Paint

- **The more non-critical things you defer loading, the better**.

  - Ideally, you would be able to **only load the assets related to LCP at first load**, and then lazy-load the rest.

- **Consider using `loading="lazy"` on `iframe` or `img` tags**.

  - In addition, **consider using `link rel="preload"` on critical images**.

    - This will make them download as soon as it is possible, and not when they are discovered by HTML parser.

- **Changing image formats and optimizing the size of the image for a given viewport helps to shrink the byte size of the image A LOT**.

  - Even if you can't change the image format, you can run your image through some kind of optimizer.

  - **Use the `picture` with `source` tags with `srcset` attributes**.

Finished https://frontendmasters.com/courses/web-perf-v2/caching/
