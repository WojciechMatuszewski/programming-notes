# Events and their phases

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
