# Random Stuff

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

That said, you can actually do very interesting stuff with this API, described on Philip Walton's blog. [Link to article](https://philipwalton.com/articles/idle-until-urgent/)

## Magic Webpack Comments

There soon may be deprecated due to `webpack 5` releasing but did you know that you can use multiple of them? like

```js
const Tilt = React.lazy(() => import(/* webpackChunkName: "tilt", webpackPrefetch: true */ "../tilt"));
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

## Concurrent vs. Parallel

There **is a difference between _concurrent_ and _parallel_ execution**.

- Think about **_concurrency_ as interleaving tasks with each other. Task MIGHT execute at the same time**.

  - A good example here would be **multitasking on a single-core machine**. Here, two threads are _making progress_ but the CPU switches between them.

  - The **tasks need to be _interruptible_ to allow concurrency**.

- Think about **_parallelism_ as multiple tasks executing AT THE SAME TIME**.

  - A good example here would be **multitasking on a multiple-core machine**. Here, two threads are _independent_ of each other. They make progress AT THE SAME TIME.

## Writing new code first when refactoring

> [Based on this talk](https://www.epicweb.dev/talks/6-safe-refactorings-for-untested-legacy-code)

You stare at the code. It's clunky and complex. Lots of logic, confusing variable names and duplication all over the place. What do you do? You _refactor the code_.

I bet you are not a stranger to these circumstances. We should be refactoring a lot and frequently. But one can refactor in a "safe" way, or in a not-so-safe way.

One of the things that struck me while watching the "6 Safe Refactorings for Untested Legacy Code" talk is the stress Nicolas puts on _writing the new code first_ and then replacing the old code with it.

Of course, this makes total sense when you think about it. This allows you to work "in chunks," preferably making commits along the way. The blast radius of your refactor is reduced because you can stop at any given time.

Sadly, I've noticed that I sometimes lack the discipline to proceed with refactorings this way. The allure of improving the code _right now_ is quite strong, and while I start with the new code first, I often fail to apply the refactorings chunk-by-chunk.

This talk has been a good reminder about how important discipline is when working with convoluted code.
