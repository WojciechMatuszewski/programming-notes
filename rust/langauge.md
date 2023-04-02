# Rust language

- [Rust for TypeScript Devs](https://frontendmasters.com/workshops/rust-typescript-devs/).

  - Finished the whole thing

## Stack vs the Heap

Since, when writing rust code, one has to think a little bit about the memory management, it is paramount to understand what the _heap_ and the _stack_ is. You might have a vague idea what these things are – after all every basic CS course teaches those things. Having said that, I found myself often guessing where a given variable lands (if it's on a _heap_ or a stack_). Let us explore those concepts in terms of rust code.

- **_stack_ is much faster than the _heap_**.

- **_stack_ is much smaller than the _heap_**.

- **You want your stuff to be on the _stack_ most of the times**. If the _stack_ is faster, then there is no reason to allocate to the _heap_ without a reason.

- **Depending on what you return from the functions, you might be allocating on the _heap_ or on the _stack_**.

  - If you use a structure that lives on the _heap_, you will be allocating on the _heap_.

    - For example returning a `vec` from a function.

    - For example defining variables inside a function that live on the _heap_.

  - If you do not use a structure that lives on the _stack_, you will not be allocating on the _heap.

- **For every thing you store on the heap, there is something (a pointer) stored on the stack**.

  - This helps with cleaning up the memory.

### Borrowing and owning

- When **passing a value to a function, that function will _consume_ the value**.

  - Unless it is a reference, after the function is done, that value will no longer the accessible in the "parent scope".

- You **cannot have more than one mutable OVERLAPPING reference to something**.

  - That also applies to an items of a vector.

  - Keep in mind that the rust compiler is exceptionally smart. Some times you can have multiple mutable references, as long as the previous reference is dropped before you access the next one.

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

## Collections

### What is `collect()`

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

### The lazy nature of the iterators

In JavaScript / TS functions available on the `Array.prototype`, like `map` or `filter` are _push-based_. This means they execute eagerly and they **consume the whole array**. In Rust, it's a bit different. The iterators will be collapsed into a loop. This means **that no matter how many iterators you have, they will only traverse the array once**. In addition, to invoke those iterators, one has to call `collect`.

This is something the Rust community calls "_zero cost abstractions_". Pretty neat concept.

## Boxed values

- Boxed values land on the heap.

  - The heap is much bigger than the stack. It allows you to store bigger pieces of data.

  - Keep in mind that using heap is not free. It is much "slower" than the stack as the compiler has to track who uses what data.

Where would use use the `Box` generic type? Anywhere **where the exact size of the data of a given thing cannot be known at compile time OR the data is large and you want to avoid copying**. This means, we mostly would use this type **while working with traits** or **large pieces of data / recursive structures**.

```rust
trait Driver {
    fn drive(&self);
}

struct CarDriver {}
impl Driver for CarDriver {
    fn drive(&self) {}
}

fn create_driver() -> Box<dyn Driver> { // The size of the return value cannot be known on compile time. There might be multiple structs that implement the Driver trait
    return Box::new(CarDriver {});
}
```

## Traits

I like to think about the `trait` keyword as the `interface` keyword in Go. In addition, one cool thing is that **traits do not have to live in the same file as the data structures that implement them**. This means that **the implementation is separate to the data structure**. Such a powerful concept that allows you to co-locate different _concerns_ together!

### Default implementation for Traits

When I first saw this I was blown away. Remember the whole _separating the data and traits_ stuff? It turns out you can also provide defaults for the trait methods **in the trait itself**. Check this out:

```rust
trait Collidable<T> {
    fn collide(&self, other: &T) -> bool;

    fn collides(&self, others: &[T]) -> bool {
        for other in others {
            if self.collide(other) {
                return true;
            }
        }
        return false;
    }
}
```
