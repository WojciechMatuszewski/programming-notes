# Events

Usually we are working with frameworks like _Angular_, _React_ or _Vue_ that abstract interactions with the DOM event model.
While this is certainly a productivity boost, I think it's at least worth knowing about 2 most important concepts - _event capturing_ and _event bubbling_.

## Bi-directional nature of events

Most of you are most likely thinking about the events as a uni-directional flow.
You might think they start at the element the event was triggered on, and _bubble up_ to the root of the document.
That is partially true and the behavior I described is called _event bubbling_.

To illustrate, let's say you have 2 elements inside each other.

```html
<div class="box-1">
  <div class="box-2">Box 2</div>
</div>
```

And some _JavaScript_ that adds the listeners

```js
box1.addEventListener("click", (e) => {
  console.log("box-1");
});

box2.addEventListener("click", () => console.log("box-2"));
```

If I click the element with the text "Box 2" I would expect to see this in my console

```shell
> box-2
> box-1
```

This would be uni-directional _event bubbling_ we talked about.

But... did you know that there is a third parameter you can pass to the `addEventListener` API?
That parameter indicates if the listener is registered for the _event bubbling_ phase (false) or the _event capturing_ phase (true).

What will happen if I add `true` to both of the listeners?

```js
box1.addEventListener(
  "click",
  (e) => {
    console.log("box-1");
  },
  true
);

box2.addEventListener("click", () => console.log("box-2"), true);
```

By clicking the element with the text "Box 2", in my console I will see

```shell
> box-1
> box-2
```

Completely in reverse right?

Well **the _event capturing_ phase is the reason the DOM event system is bi-directional**.

One interesting outcome of all of this, is that you can block the listener registered on the element with the text "Box 2" from firing all together.
All I have to do to make it work is to specify the other listener to be registered within the _event capturing_ phase, then stop the propagation, like so

```js
box1.addEventListener(
  "click",
  (e) => {
    e.stopPropagation();
    console.log("box-1");
  },
  true
);

box2.addEventListener("click", () => console.log("box-2"));
```

Now, when I click the element with text "Box 2", in my console I will see

```shell
> box -1
```

Despite me clicking the element with text "Box 2", only 1 listener was fired and it was not the one registered on that element.
This is behavior might seem like a bug, but as always, is pretty logical when you are aware how the DOM event system works.

## Configuration of the `addEventListener` API

It turns out the third parameter of the `addEventListener` function is overloaded. It can be either a boolean or an object.
With the object passed as a third parameter, you can configure

- capture
- once
- passive

`capture` is for specifying the phase under which the listener will be registered. Talked about this in the second above.

`once` means that the **event listener will automatically remove itself after the first time it fires**. This is soo useful, you do not have to keep the reference anymore.

`passive` is for performance gains. Granted these are probably insignificant, but, nevertheless you can specify those. With `passive` you tell the browser that you will not use `event.preventDefault`, in return you get those aforementioned performance optimizations.

### Class as the listener

It turns out, instead of a function, one can specify a class as a second parameter to the `addEventListener` function.
This might come in handy when you want to ensure referential equality needed for **cleaning up** the listener.

Learn more about the _class as the listener_ [here](https://www.stefanjudis.com/today-i-learned/addeventlistener-accepts-functions-and-objects/?utm_source=stefanjudis).

## Canceling the event listeners

When you first encounter the `AbortController` API, you will most likely read about its ability to cancel network requests. And this is great! It allows us to eliminate the "flickering state" issues when fetching in `useEffect` with a dependency array.

But **the `AbortController` can also be used with event listeners!** (and even in your custom code, but that is beyond this entry). By passing the `signal` to the `addEventListeners` you can later "detach" or "remove" the underlying listener whenever you call `.abort`.

```js
const controller = new AbortController();
const { signal } = controller;

window.addEventListener("resize", () => fire(), { signal });

controller.abort(); // remove the listener
```

Notice that I've created a new instance of the listener in the `addEventListener` callback. **With the `AbortController` API, you do not have to worry about the referential equality of the callbacks!**. Pretty neat if you ask me.

- You can read more about many other `AbortController` use cases [here](https://whistlr.info/2022/abortcontroller-is-your-friend/).

## Custom events and the `EventEmitter`

There is an API for creating custom events that you can incorporate to your application. **The events can be scoped to a given "emitter" or global**.
This makes some of the npm packages obsolete. Of course the packages might provide some additional sugar on top, but in most cases, you will not need them!

### Global events

Use the `dispatchEvent` API available on the `window`.

```js
window.dispatchEvent(new Event("..."))
```

Then you can listen in any other part of your application.

```js
window.addEventListener("foo", () => {
  // ...
})
```

### "Scoped" events

You can **either scope the event to a given element, or create an instance of an "emitter"**. Usually you want your events to be scoped to a given context.

```js
const emitter = new EventTarget();

// You can also query for an element here instead of using `window`.
emitter.dispatchEvent(new Event("..."));
```

```js
emitter.addEventListener("..")
```

## The `setPointerCapture` function

Imagine a scenario where you want to implement a "drag" feature for some UI element. From the experiences you already had, you most likely assume you should be able to "drag" the element, even if your thumb/mouse moves OUTSIDE of this element.

To achieve this effect, you will **need to listen to future pointer events that might happen "outside" of a given element**. To achieve this, you should consider using the `setPointerCapture` function on the event.

```js
const box = document.querySelector(".box");

const handlePointerMove = (event) => {
  event.target.style.transform = `translateX(${event.clientX - 50}px)`;
};

box.addEventListener("pointerdown", (event) => {
  event.target.setPointerCapture(event.pointerId);
  event.target.addEventListener("pointermove", handlePointerMove);
});

box.addEventListener("pointerup", (event) => {
  event.target.removeEventListener("pointermove", handlePointerMove);
  event.target.releasePointerCapture(event.poinsterId);
});
```

[MDN has a very good documentation regarding this event](https://developer.mozilla.org/en-US/docs/Web/API/Element/setPointerCapture). Check out the demo – start dragging, then move your pointer down. Notice that you can still drag the element, despite having the pointer outside of its bounds!
