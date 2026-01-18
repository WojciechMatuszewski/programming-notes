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
let fragment = document.createDocumentFragment();

for (let item of items) {
  let li = document.createElement("li");
  li.textContent = item;
  fragment.appendChild(li);
}

ul.appendChild(fragment);
```

> The key difference is due to the fact that the document fragment isn't part of the active document tree structure. Changes made to the fragment don't affect the document.

In short, the API was made just for this use-case.

## Date format relative to certain date

> Inspired by [this article](https://www.builder.io/blog/relative-time).

The web has made lots of improvements with date handling. The `Intl.X` family of APIs is expanding and enabling developers to do more with less 3rd party code or libraries.

Have you ever had a situation where you had to produce a string akin to `20 minutes ago`? I bet you have.

Some of you might have used a library called `moment.js` to achieve what you needed, but that library is pretty heavy. Nowadays there are numerous replacements for the `moment.js` library that are much smaller, for example the `date-fns` function.

But adding a library to a project comes at a cost. You have to update it from time to time, and ensure there are no security vulnerabilities. What if we could calculate that relative string via vanilla js? Here is where the `Intl.RelativeTimeFormat` comes into the picture.

```js
export function getRelativeTimeString(
  date: Date | number,
  lang = navigator.language
): string {
  // Allow dates or times to be passed
  const timeMs = typeof date === "number" ? date : date.getTime();

  // Get the amount of seconds between the given date and now
  const deltaSeconds = Math.round((timeMs - Date.now()) / 1000);

  // Array representing one minute, hour, day, week, month, etc in seconds
  const cutoffs = [60, 3600, 86400, 86400 * 7, 86400 * 30, 86400 * 365, Infinity];

  // Array equivalent to the above but in the string representation of the units
  const units: Intl.RelativeTimeFormatUnit[] = ["second", "minute", "hour", "day", "week", "month", "year"];

  // Grab the ideal cutoff unit
  const unitIndex = cutoffs.findIndex(cutoff => cutoff > Math.abs(deltaSeconds));

  // Get the divisor to divide from the seconds. E.g. if our unit is "day" our divisor
  // is one day in seconds, so we can divide our seconds by this to get the # of days
  const divisor = unitIndex ? cutoffs[unitIndex - 1] : 1;

  // Intl.RelativeTimeFormat do its magic
  const rtf = new Intl.RelativeTimeFormat(lang, { numeric: "auto" });
  return rtf.format(Math.floor(deltaSeconds / divisor), units[unitIndex]);
}
```

Yes, it requires some work, and ideally we would have a single `Intl.X` method to handle all of this for us. But let us not forget that the alternative is adding a library which is not always the best choice.

## Formatting time durations

Check [this API out](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Intl/DurationFormat). **You can produce "relative time" strings for a given duration!**

I'm so pleased this is now built-in into the browser. Granted, this feature is _very_ new, so it might take some time before you can use it in any application, but nevertheless, it is a win for us!

## Comparing strings

There are multiple ways to compare string in JavaScript.

One might use the `>` or `<` operators but those do not really work as I would expect them to work (at least not the same way lexical comparison works in DynamoDB).

Another way might be to use the `"somestring".localeCompare("someOtherString")`. This works, **but you have no control many variables, for example whether we should favour upper-case strings or not**.

Guess, what. There is an `Intl` class for that. The `Intl.Collator`. [It is pretty magical. Read the docs here](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Intl/Collator).

```js
["a", "z", "g"].sort(new Intl.Collator().compare);
// ["a", "g", "z"]
```

The best part is that **this class enables you to really take control over locales, and all the other various options**. I have to say, I'm impressed how configurable this class is. Kudos for the spec creators.

## Symbols

> Based on [this blog post](https://www.trevorlasn.com/blog/symbols-in-javascript).

The `Symbol` primitive allows you to create unique value that will not clash with other values.

It is **a common pattern to use the `Symbol` primitive for object keys – think metadata**.

```js
const uniqueKey = Symbol("foo");
const user = {
  [uniqueKey]: 123,
  name: "Alex",
};
```

**The benefit of using the `Symbol` as a key** is that they **will not show up in `.keys`, `.entries` and other similar functions**.

```js
Object.keys(user); // ["name"]
Object.entires(user); // [["name", "Alex"]]
for (let key in person) {
  console.log(key); // "name"
}
```

You have to use the **`getOwnPropertySymbols`** to get the symbols from the object.

### The `Symbol.for` API

You can use the **`Symbol.for` API to create shared unique values**.

```js
const uniqueSymbol = Symbol("unique");
const uniqueSymbol2 = Symbol("unique");

uniqueSymbol == uniqueSymbol2; // false

const sharedUniqueSymbol = Symbol.for("unique");
const sharedUniqueSymbol2 = Symbol.for("unique");

sharedUniqueSymbol == sharedUniqueSymbol2; // true
```

## `Object.groupBy`

> [MDN documentation](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Object/groupBy)

Think of this an alternative to the `reduce` function **in certain cases**.

```js
const numbers = [1, 2, 3, 4];

const groupedReduce = numbers.reduce(
  (grouping, number) => {
    if (number % 2 == 0) {
      grouping.even.push(number);
      return grouping;
    }

    grouping.odd.push(number);
    return grouping;
  },
  { even: [], odd: [] },
);

const groupedNew = Object.groupBy(numbers, (number) => {
  if (number % 2 === 0) {
    return "event";
  }

  return "odd";
});
```

Pretty neat!

## Import attributes

> Based on [this blog post](https://2ality.com/2025/01/import-attributes.html).

When using frameworks, like Next.js, you can import CSS files right into your JSX / TSX files like so:

```tsx
import styles from "style.css";
```

This is possible because the frameworks bundler. It is the bundler that transforms those imports.

Would it not be awesome to have the ability to import JSON/CSS and other files directly into JS files _without_ any bundlers or additional tools that interfere with out code?

While not supported in every browser, this is [possible with the _import attributes_ feature](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Statements/import/with).

```js
import styles from "style.css" with { type: "css" };
```

You can also import JSON files!

```js
import styles from "data.json" with { type: "json" };
```

**This will only work if the server serving those files sends the correct `Content-Type` header**.

According to the MDN, under the hood, the browser makes a network request to fetch those files. **I wonder if this will affect webpage performance in any way**.

## New Iterator Methods

> Based on [this blog post](https://2ality.com/2025/06/ecmascript-2025.html#iterator-helper-methods)

> [Another resource](https://allthingssmitty.com/2026/01/12/stop-turning-everything-into-arrays-and-do-less-work-instead/)

There are new set of methods for working with _iterators_, like `drop`, `take` and others.

**The biggest difference between array methods is that iterators DOES NOT create intermediate arrays**. This is huge win for large datasets.

```js
const foo = [1, 2, 3, 4];

foo.values().take(1);

foo.values().map(/*stuff*/);
```

Please note that **iterators can work over ANY ITERABLE data structure, not only arrays**. This means **async generators as well!**

## Handling URLs

### The `URL` constructor

So many times I've seen engineers struggle to "compose" an URL with query parameters. They tend to concatenate multiple strings together using elaborate conditionals making the code quite hard to read.

If you find yourself in similar situation, **consider using the `URL` constructor**. It is _very_ useful.

```js
const u = new URL("/foo/bar/baz", "https://google.com");

if (someCondition) {
  u.searchParams.append("bar", "baz");
}
```

You can even append parts to the `pathname` _after_ you already created the URL. **Sadly, I think you can do that only in Node.js and not in the browser**.

```js
import path from "node:path";

const u = new URL("foo/bar/", "http://google.com");

u.pathname = path.posix.join(u.pathname, "/bazz");
```

### Do not forget about `URLSearchParams` constructor

Only concerned with search params?

```js
const params = new URLSearchParams([["foo", "bar"]]);
params.append("baz", "123");
```

### Matching against the URLs

> [Based on this blog post](https://jschof.dev/posts/2025/11/build-your-own-router/?utm_source=stefanjudis&utm_medium=email&utm_campaign=web-weekly-176-scope-scope-and-enterkeyhint)

Use the `URLPattern` constructor to create a "matcher" you can use to "test" against various URLs.

```js
const pattern = new URLPattern({ search: "foo=bar*" });

console.log(pattern.test("http://google.com?foo=bar&bar=baz")); // true
```

You can match against other `URL` properties like `protocol` or `pathname`.
