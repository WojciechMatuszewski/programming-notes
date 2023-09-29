# Vanilla JS

## Inserting DOM nodes

When working with DOM, you might find yourself wanting to insert a lot of DOM nodes into the HTML at once.

```js
let items = [...Array(1000).keys()].map((i) => `Item ${i}`);

let ul = document.getElementById("myList");

for (let item of items) {
  let li = document.createElement("li");
  li.textContent = item;
  ul.appendChild(li);
}
```

While this works, **it will not perform well**. In the code above, we are creating and then instantly putting the node into the DOM. That violates one of the most important rules for working with DOM – **batch actions together**.

A much better way to solve this problem would be to use `documentFragment`. We could insert all the nodes into the `documentFragment` and then commit the fragment to the dom.

```js
let items = [...Array(1000).keys()].map((i) => `Item ${i}`);

let ul = document.getElementById("myList");
let fragment = document.createDocumentFragment()

for (let item of items) {
  let li = document.createElement("li");
  li.textContent = item;
  fragment.appendChild(li);
}

ul.appendChild(fragment)
```

> The key difference is due to the fact that the document fragment isn't part of the active document tree structure. Changes made to the fragment don't affect the document.

In short, the API was made just for this use-case.
