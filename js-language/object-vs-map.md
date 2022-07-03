# Objects vs Map in JS

For the longest time, I've used objects for all key-value lookup-related use cases (which also applies to hash maps). As I learned today, such practice might not have been a good one, and I should consider using the `Map` API a bit more.

Let me explain.

## Problems with plain objects

- **Key collisions with the prototype** can happen. Maybe I'm lucky, or maybe I'm aware of a latent bug that haunts the production system I'm responsible for (most likely the latter). If a given key clashes with a property defined on the `prototype`, it will overwrite the prototype property.

  ```js
  const myHashMap = {};
  myHashMap.hasOwnProperty = "foo"; // Oopsie...
  ```

- **Iterating over properties is difficult**. Have you tried using the `for ... in` loop to iterate over the object properties? Did you notice that **the `for ... in` loop iterates over the properties defined in the prototype as well?**

  ```js
  const hashMap = {};
  for (let v in hashMap) {
    console.log(v); // never runs OK!
  }

  Object.prototype.someProperty = "valueFromPrototype";
  for (let v in hashMap) {
    console.log(v); // logs "valueFromPrototype". Not ideal if you ask me.
  }
  ```

- **The API ergonomics are not there**. The `Map` API contains various practical methods that significantly improve its ergonomics. I'm mainly talking about the `size` and `clear` functions that are awesome! You can familiarize yourself with the full breadth of the `Map` API [on MDN](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Map).
