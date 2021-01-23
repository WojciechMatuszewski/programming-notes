# Improving Web Performance

Course material from frontendmasters

## Part 1 - Understading

### Psychology of performance

- poor performance will most likely influence revenue or conversion rates

- correlation != causation, there is no one metric that perfectly causes another

- performance is important because Google says so. Google owns a big portion of the internet

- angry and frustrated users wont stick around long

- business model might influence the performance. If you completely rely on ads you website might be slow

- notion of perceived performance

### Measuring performance

#### The old way

- simple page load using the `load` event

- sadly the metric is skewed (using lazy loading)

#### The new ways

- these are harder to break

- four different metrics: _The core web vitals_

- metrics

  - FCP: the time user clicks the link to the time that something meaningful is rendered, eg. title, header or the loader
  - LCP: the time when the largest area on your page loaded (when user thinks your site is _almost_ ready). Usually the LCPs are images
  - CLS: content pushes eg. ad loading asynchronously before the button you want to press. This metric sums all the times it happened (all the times the layout shifted)
  - FID (first input delay): time between the time page looks ready and the time browser can handle user events. Can be a problem when deferring a lot of javascript

#### Tips

- detach dev tools while running lighthouse. Not many users run the website in such a small window
- lighthouse performance is relative to your machine
- keep your browser in the foreground

### Interpreting performance data

- usually, the data will be skewed. Use medians (p75, p95)

## Part 2 - Improving web performance

- `window.performance` browser API

- `performance.getEntries` will return array of really low-level data about performance. Cool stuff

- to gather the web vitals manually use the `PerformanceObserver` API. It is similar to `IntersectionObserver`

  ```js
  var fcpObserver = new PerformanceObserver(function handleFCP(entryList) {
    var entries = entryList.getEntries() || [];
    entries.forEach(function (entry) {
      if (entry.name === "first-contentful-paint") {
        data.fcp = entry.startTime;
        console.log("Recorded FCP Performance: " + data.fcp);
      }
    });
  }).observe({ type: "paint", buffered: true });
  ```

- to send the data to your backend consider listening on `visibilitychange` event. Since using `fetch` would be unreliable in this case, you should perefer the `navigator.sendBeacon` API

  ```js
  window.addEventListener("beforeunload", () => {
    navigator.sendBeacon("/api", payload);
  });
  ```

### Improving FCP
