# Ultimate Go

## Pointers

### Pointers Part 1 (Pass by Value)

When you create a go-routine a form of sandbox forms.
This sandbox can only perform operations within the given slice of memory.

By operating on `pointers` you can basically step out from such sandbox and pass address around to different parts of a program.

Crossing the sandbox, by copying the data (passing by value).

```go
package main
func main() {
    count := 3
    increment(count)
    // count here is still 3
}
func increment(val int) {
    // mutate memory but only here, in this context. No side-effects occur!
    val++
}
```

Since we passed `val` by value (passing copy), we only mutated `val` in the context `increment` function.

**The cost is that you have multiple copies of given data in your program**.

### Pointers Part 2 (Sharing Data)

Pointer semantics serve a purpose of sharing data between different boundaries.
**Sharing** is the key word here, if you do not need it, do not use it.

Passing and address so, **pass by value or reference**?

```go
package main
func main() {
    count := 3
    increment(&count)
    // count here is 4
}
func increment(val *int) {
    // mutate the memory OUTSIDE of this boundary.
    *val++
}
```

This is **pass by value**. It just happens that the **data we are passing is the underlying memory address of `counter`**.
Pointer variables are **not special**. They just are used to store addresses.

Next thing is the `*int` syntax. Can be confusing so lets explain different variations:

- `val` means **what is inside the box** (value of)
- `&val` means **where is the box** (address)
- `*val` means **find me a box given this address** (indirect memory read/write)

Pointers pose a huge thread though. **Pointers introduce side-effects to your codebase**, use them carefully, especially in multi-threaded environment.

### Pointers Par 3 (Escape Analysis)

Go does not have constructors, they hide _the cost_.

_Escape analysis_ is the term used to describe the mechanics of golang compiler, when deciding whenever to put stuff on stack or on heap.
It happens when you return stuff from eg. a function, so called **sharing**.

Try to abstain from _pointer schematics_ during construction. This may lead to your code being difficult to read.

```go
package main
type user struct {
    name string
    email string
}
func main() {
    u := user{
        // ..
    }
    // consider NOT doing this \/
    u2 := &user{
    // ..
    }
}
```

The _sharing_ part should be clearly visible near `return` statement. When assigning pointer schematics during construction that is not always the case.

### Pointers Part 4 (Stack Growth)

### Pointers Part 5 (Garbage Collection)

Garbage collection causes _stop the world event_ to happen. This event disables any writes and makes it possible for GC to do it's job.
We as programmers want to be sure we are writing our programs in GC-friendly way so that _stop the world_ event takes the least amount of time possible.

### Constants

- only exist @compile-time

There exists a mechanism similar to coercion in JavaScript, called _type promotion_.
An example

```go
package main
func main() {
    const val int8 = 2
    test := 2 * val
}
```

Now, after multiplication `test` will be of **type int8**. It got promoted during the operation.
There is a notion of _kind_ when it comes to types. `test`s kind was just an `int`, but it got promoted.

## Data-Oriented Design

### Arrays 1 (Mechanical Sympathy)

- main memory is very slow to access
- small is fast, the lower memory footprint is the better performance will occur (L1, L2, L3 cache).
- predictable access patterns are very important. That why you should prefer to use arrays. They have such predictable access patterns
  and can be optimized at hardware level (prefetcher).

### Arrays 2 (Semantics)

There is a gotcha with `range` loops.

```go
package main
func main() {
    friends := [5]string{"Annie", "Betty", "Charley", "Doug", "Edward"}
    	for i := range friends {
    		friends[1] = "Jack"

    		if i == 1 {
    			fmt.Printf("Aft[%s]\n", friends[1])
    		}
    	}
}
```

Here, `for i := range friends` is using **pointer semantics**. What does that mean?
That means that **`friends is NOT copied`** during the operation.
`friends[1] = "Jack"` will mutate the original array and inside the `for loop` we will see `Jack` instead of `"Betty"`.

This was somewhat expected, but check this out:

```go
package main
func main() {
    friends := [5]string{"Annie", "Betty", "Charley", "Doug", "Edward"}
    for i, v := range friends {
            friends[1] = "Jack"

            if i == 1 {
                fmt.Printf("v[%s]\n", v)
            }
        }
}
```

The only difference here is the `v` variable. **Now the `range loop` will use value semantics**. That means that the array is **copied!**.
So, we are still mutating the original array BUT the output will be `"Betty"`

### Array 3 (Slices)

`slice` is very similar to `C++` vectors.

- built in function `make` to create `slices`

- `slices` **should not be shared**. They are designed to be used within value semantics context.

### Array 3 (Appending Slices)

- slice initialized with `var` => _zero value `slice`_ (nil slice)

- slice initialized with `data := []string{}` => _empty slice_

Since `var []string` **DOES NOT** mean _empty slice_ it should be returned when an error happened. This way
the consumer will not be confused about the return result.

- start with _nil slice_ when you do not know the length of array.

There is notion of `capacity` when it comes to `slices`. You can specify `capacity` as 3rd argument of `make` when creating slices.

`append` call has a cost. When you can prefer specifying `length and capacity` and using index for assignment.

**When slice has reached it capacity limit, `append` will create new backing array with capacity doubled!**.
This can cause confusion. Lets consider this example

```go
users := make([]string, 3) // length 3, capacity 3
// you could also be more specific and write make([]string, 3, 3)
users = append(users, "Wojtek")
fmt.Println(cap(users))
```

`capacity` is now 6!. And how is our `users` slice looks like now?
`["","","","Wojtek"]`. **Notice the empty strings**.

### Slices of Slices

You can take a slice of slice. The syntax is similar to the one in `python`.

```go
    users := string[]{"Wojtek", "Mateusz", "Ela"}
    users2 := users[1:3]
```

There are some gotchas associated with this technique though.

- taking `slice` out of existing `slice` this way **shares existing backing array**

This means that **operations that mutate `users2` will probably mutate `users`**

```go
	users := []string{"Wojtek", "Mateusz", "Ela"}
	users2 := users[1:2]

	users2[0] = "CHANGED"
	fmt.Println(users2[0], users[1])
```

This logs `CHANGED` and `CHANGED`. This is very easy to miss so **watch out for such behavior**

Another thing that is worth knowing that the **range syntax** (x:y) is **not inclusive on the end range**.

```go
  users := []string{"Wojtek", "Mateusz", "Ela"}
  users2 := users[1:2]
  // users2 => [Mateusz]
```

So the `x:y` syntax takes **from x to y but NOT INCLUDING y**.

Another thing that comes in to play when taking `slice` out of `slice` is the `capacity`. The `capacity` changes on the _derived_ slice.

```go
    users := []string{"Wojtek", "Mateusz", "Ela"}
    users2 := users[1:2]

    fmt.Println(cap(users2))
```

The new capacity is `2`. There is a direct correlation between the starting point of new `slice` and length of original `slice`.
The new `capacity` is `lenOfOldSlice - startingPointOfSlicing`.
So in our case that `capacity` is 2.

#### Append and mutations

So by now you should be familiar how `append` works. Mainly how the share semantics come into play depending on the `capacity`.

What will happen if I `append` to `users2`?

```go
	users := []string{"Wojtek", "Mateusz", "Ela"}
	users2 := users[1:2]

	users2 = append(users2, "CHANGED")
	fmt.Println(users2[1], users[2])
```

Guess what? **I've must made a mutation using `append`**. Since the capacity is `3` append will not create a new `backing array`. It will _`push`_ `"CHANGED"` into `users2` effectively mutating `users`.

There is a way to deal with such situations though, that is making use of **third argument when taking slices of slices**. So the syntax is **not** `x:y` **but** `x:y:z`. The `z` part tells golang **where to stop when expanding in terms of capacity**. Normally, as discussed earlier, the `capacity` would be calculated from `lenOfOldSlice - startingPointOfSlicing`, but with this syntax its actually `2`.

So when `length = capacity` `append` will create new `backing array`, and we will be free of side-effects.

```go
	users := []string{"Wojtek", "Mateusz", "Ela"}
	users2 := users[1:2:2]

	users2 = append(users2, "CHANGED")
	fmt.Println(users2[1], users[2])
```

Logs: `CHANGED Ela`

### Maps

- **you can use `make` to create `zero-value` maps**.

* **you cannot iterate over `zero-value map`**. If you are in such situation you have to initialize those `key-values`.

What is actually interesting, `for range` loops are **random in terms of ordering**. That means that initialization order does not matter and your `key, value` when `ranging` can be different every time you loop (in terms of ordering).

You can also use `literal initialization` for initializing maps.

- there is an `delete` operator for maps to delete given key.

**Not everything can be a key to a map**. This is expected, you would not want to use `slices` as keys of your map, that would not make sense.

## Composition

### Grouping Types

- There is **no sub-classing** in Go!

* Avoid something called **type pollution**. This is where you define a type only to be extended somewhere else (probably embedded). A little of copy-pasting will not kill you. With **embedding you CAN create a coupling problem**. Of course that is not the case with every situation where you need / want to use embedding.

- Stop thinking about what things are and **start thinking about that things DO**. Always **favour behaviors (interfaces) not structs (what things are)**
