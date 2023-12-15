# Random Stuff

## CORS

CORS is the mechanism which enables one webapp (like your localhost) to share
some resource with another webapp (like your endpoint API).

If those 2 apps have **the same origin** they can easily share those resources
with no hassle at all.

Problems begin when they are on different origins.

So what does _different origin_ mean?

- different domain like `google.com` and `twitter.com`
- different subdomains like `localhost:3000` and `localhost:3000.api/v1`
- different ports like `:3000` and `:4000`
- different protocols like `http` and `https`

To make it work you have to follow the CORS standard.

So how does it work?

Suppose we have 2 apps: A and B. They want to share resources. App A makes a
POST request to app B:

- `preflight` request is made (before the actual request) **also known as
  OPTIONS call**
- app B now have the responsibility of verifying either this request is valid or
  not.
- app B sets some additional headers to that request and sends it back.
- now browser knows if the request is valid or not. The actual `POST` request is
  made

### Simple Request vs Preflight Request

So we've seen how the preflight mechanism works. But the next question on your
mind probably is: is this happening every time I send a request?

Well, no.

Some request are marked as `simple` by the browser and the preflight (`OPTIONS`)
request is not send.

### Caching

Browsers can actually cache preflight responses. You usually specify that in a
header.

## Observers API

You probably know of `intersectionObserver`. An API to check if given object is
visible at the screen currently.

But there are other _observers_ too:

- `mutationObserver`: used to watch for `DOM` _mutations_.

- `resizeObserver`: a new kid on the block. Tries to achieve the holy grail of
  being notified when **given element** resizes to given width/height.

There is actually very interesting article on `resizeObserver` by Philip Walton
[Link to article](https://philipwalton.com/articles/responsive-components-a-solution-to-the-container-queries-problem/)

## Check for Idle period

This would be nice would not it? Having a way to know when browser is finished
doing stuff so that we can fire off some kind of computation.

There is an API for that: `requestIdleCallback`, but sadly is not all green when
it comes to browsers.

That said, you can actually do very interesting stuff with this API, described
on Philip Walton's blog
[Link to article](https://philipwalton.com/articles/idle-until-urgent/)

## Magic Webpack Comments

There soon may be deprecated due to `webpack 5` releasing but did you know that you can use multiple of them? like

```js
const Tilt = React.lazy(() =>
  import(/* webpackChunkName: "tilt", webpackPrefetch: true */ "../tilt")
);
```

This works !

## Web workers with webpack

So web workers are great. They enable you to offload the work on different thread. Nice!.

But did you know that there is an webpack plugin which enables you to turn any js file into _web worker_?. That plugin is called `workerize-loader`.
Let's say you have a module called `expensive.js` and you want to import that module as _web worker_.

```js
import expensiveWorkerized from "workerize-loader!./expensive";
```

**BOOM!**. Thats all. Granted now methods exposed by `expensiveWorkerized` are _async_ but that should not be a problem.
