# Useful Web APIs

## `structuredClone`

How many times you had to _deep clone_ an object in JavaScript? Probably quite a lot, right?

```js
const deepClone = (data) => {
  return JSON.parse(JSON.stringify(data));
};
```

We have been taught to use the `JSON` family of APIs for a long time. This is justified because **there were no better tools to do this until recently**.

The `JSON.parse(JSON.stringify)` methods works, **but it has a few key problems**.

- It does not parse `Date` or similar objects.

- It does **not handle circular references**.

- It does not work with [_transferable objects_](https://developer.mozilla.org/en-US/docs/Web/API/Web_Workers_API/Transferable_objects) – those are used to share data between the main thread and workers.

Read more about the `structuredClone` capabilities [here](https://www.builder.io/blog/structured-clone). It has been supported for a while, and, unless you are shipping a very old JS, consider using it whenever you need to deeply clone some data!

## The `Intl` family of APIs

Holy smokes, I cannot emphasize enough how useful some of the APIs within the `Intl` family are. **The problem is that most of the people I know are not aware of them!** It will be up to you to educate your fellow colleagues that they do not necessary need a library to format dates or make some word plural.

[Here is the link to the MDN](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Intl) containing all the information about the `Intl` namespace. I found the following very useful:

- The `Intl.DateTimeFormat` APIs.

- The `Intl.PluralRules` APIs.

- The `Intl.Segmenter` APIs.

- The `Intl.Collator` APIs.

  - You would use this to sort large arrays of strings. The `"foo".localCompare` is quite slow.

    - [See this blog post for more details](https://claritydev.net/blog/faster-string-sorting-intl-collator)

Of course, this list is highly subjective. It all depends on the requirements that you have. The most important part is to know these APIs exists – you do not have to memorize them!
