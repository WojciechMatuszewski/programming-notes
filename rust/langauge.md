# Rust language

- [Rust for TypeScript Devs](https://frontendmasters.com/workshops/rust-typescript-devs/).

  - Finished day 1, part 2 37:12

## &str vs String

When I first started learning rust, I was very confused about the two types in the language, mainly the `&str` and the `String` types. To me, both of them represented the same data, but I was looking at them from JS/TS perspective. I did not understand what is going on underneath those tokens.

### The String

- **Heap allocated**. Quite an important detail.

- Can be mutable.

### The &str

- It is a **view into a certain String**, often called _a slice_.

- It is immutable.

---

Imagine you have a huge wall of text you would like to work on. If you were to always work on the `String` type, you would be copying this huge wall of text all over your program. This is very inefficient in terms of memory.

With the `&str` type, you can declare a "view" into particular parts (or the whole thing) of that text block. **Since the `&str` cannot grow**, by passing the `&str` around you are not making any copies. **It is a pointer**.

## Type placeholders

Sometimes you want the compiler to infer to infer the type instead of writing it explicitly. This is super handy when working with collections.

```rust
fn main() {
    let list = vec![1, 2, 3];

    let list: Vec<_> = list
        .iter()
        .map(|x| {
            return x + 1;
        })
        .collect();

    println!("list = {:?}", list);
}
```

Notice the `Vec<_>`. This **tells the compiler to infer the type of the `collect` operation**. Since the map returned numbers, the compiler is smart enough to deduce that the type must be `Vec<i32>`. Of course, you could always specify a different type, and the compiler would run the `into` method implicitly.

> The Rust integer types all implement the From<T> and Into<T> traits to let us convert between them. The From<T> trait has a single from() method and similarly, the Into<T> trait has a single into() method. Implementing these traits is how a type expresses that it can be converted into another type.

From [this book](https://google.github.io/comprehensive-rust/exercises/day-1/implicit-conversions.html).

## What is `collect()`

**`collect()` will consume the iterator and create a new collection**. This is a syntactic sugar over making your own collection and then calling the `.next` on it.

```rust
let list = vec![1, 2, 3];
let mut list = list.iter().map(|x| {
    return x + 1;
});

let mut new_vector: Vec<i32> = vec![];
while let Some(x) = list.next() {
    new_vector.push(x)
}

println!("list = {:?}", new_vector);
```

Mind the borrow checker here! Know that the **iterator refers to the pointer of the vector**. If I were to define the `vec![...]` inline, the compiler would not be happy.

```rust
let mut list = vec![1, 2, 3].iter().map(|x| { // Errors on this line
    return x + 1;
}); // the `vec![...]` is cleared here

let mut new_vector: Vec<i32> = vec![];
while let Some(x) = list.next() { // but we reference it here implicitly by referencing the iterator.
    new_vector.push(x)
}

println!("list = {:?}", new_vector);
```
