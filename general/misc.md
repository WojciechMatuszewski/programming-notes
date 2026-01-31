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

# The complexity of "soft deletes" and what to do about it

> [Based on this blog post](https://atlas9.dev/blog/soft-delete.html)

## Why

Why do we even have to implement this feature in the first place? Make sure you understand the _why_ before you embark on the implementation.

Perhaps you _assume_ you need to implement this feature, and it might not be _truly_ necessary?

Remember – it is always better to ask a couple of times, and make sure people are aligned on the requirements, than to start implementing complex features you might not need.

## The "first thing that came to my mind" way

The first thing that could come to your mind might be adding `expires_at` or `is_archived` properties to the items in your database.

Consider the following:

1. Unless you use a schema-less database, where each item can have different properties, you might need to append this property to each existing item in your database (with `null` or other equivalent value for non-deleted items). Doing so is usually quite cumbersome.

2. Unless your database has TTL process built-in, where it deletes records given a date in a certain attribute is greater than current date (a good example of this would be [DynamoDB](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/TTL.html)), you will need to introduce a CRON that deletes those items for you. That's yet another piece of application code you need to write.

3. Any migration that you write, will either have to migrate those "soft-deleted" items, OR skip them. This is an additional complexity in your migration code.

4. Restoring the item might not be as simple as setting that `expires_at` to `null`. Consider the calls your service makes when you insert a new item. Now consider how this could change over time. It is _very_ hard to know which additional calls you need to make for certain item to mark it as `expires_at=null`.

## Separate data store with expired records (events)

Consider the following setup: one data store for "regular" items, one for "expired" items. When you delete an item from the "regular" datastore, you send an event and another system writes that "deleted" item to another data store.

1. More infrastructure to manage.

2. The migrations on the "regular" items are much simpler, as you have only the items you care for there.

3. The "soft deleted" items can be serialized to JSON (since you are using separate datastore). This means you are not constrained to fixed schema of your "regular" database.

4. Recovering the item is "as simple as" removing it from the "expired" items datastore and creating a new item in your application. This allows you to treat the item recovery as adding a new item in your database. This means that all the logic that is currently relevant will be executed for this item.

A good tools to implement this might be SQS, EventBridge or DynamoDB streams for events and S3 for archiving items.

## Replica that does not process deletes

This is an interesting idea author outlined that I have not considered yet – having a read-only replica of your data that does _not_ process deletes.

1. The cost of replica will grow with the size of data.

2. I'm unclear how migration would work. Would we apply migrations on all items? What if I delete an item from the primary, and then run the migration? Since the replica does not process deletes, what would be the outcome?

   - We would most likely "sync" migrations. Perhaps we could run two migrations. One one the primary, and in the background, one on the replica.

## Summary

Author recommends going the events routes with two different stores. I second this recommendation, but I'm very curious about how this "read-only replica" approach could work.
