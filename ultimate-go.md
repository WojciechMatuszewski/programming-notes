# Ultimate Go


## Pointers 
### Pointers Part 1 (Pass by Value)

When you create a go-routine a form of sandbox forms. 
This sandbox can only perform operations within the given slice of memory.

By operating on ``pointers`` you can basically step out from such sandbox and pass address around to different parts of a program.

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
Since we passed ``val`` by value (passing copy), we only mutated ``val`` in the context ``increment`` function.

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
This is **pass by value**. It just happens that the **data we are passing is the underlying memory address of ``counter``**.
Pointer variables are **not special**. They just are used to store addresses.

Next thing is the ``*int`` syntax. Can be confusing so lets explain different variations:
- ``val`` means **what is inside the box** (value of)
- ``&val`` means **where is the box** (address)
- ``*val`` means **find me a box given this address** (indirect memory read/write)

Pointers pose a huge thread though. **Pointers introduce side-effects to your codebase**, use them carefully, especially in multi-threaded environment.

### Pointers Par 3 (Escape Analysis)
Go does not have constructors, they hide *the cost*.

*Escape analysis* is the term used to describe the mechanics of golang compiler, when deciding whenever to put stuff on stack or on heap.
It happens when you return stuff from eg. a function, so called **sharing**.

Try to abstain from *pointer schematics* during construction. This may lead to your code being difficult to read.
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
The *sharing* part should be clearly visible near ``return`` statement. When assigning pointer schematics during construction that is not always the case.
### Pointers Part 4 (Stack Growth)

### Pointers Part 5 (Garbage Collection)
Garbage collection causes *stop the world event* to happen. This event disables any writes and makes it possible for GC to do it's job.
We as programmers want to be sure we are writing our programs in GC-friendly way so that *stop the world* event takes the least amount of time possible.
### Constants
- only exist @compile-time

There exists a mechanism similar to coercion in JavaScript, called *type promotion*.
An example
```go
package main
func main() {
    const val int8 = 2
    test := 2 * val
}
```
Now, after multiplication `test` will be of **type int8**. It got promoted during the operation.
There is a notion of *kind* when it comes to types. `test`s kind was just an `int`, but it got promoted.

## Data-Oriented Design

### Arrays 1 (Mechanical Sympathy)
- main memory is very slow to access
- small is fast, the lower memory footprint is the better performance will occur (L1, L2, L3 cache).
- predictable access patterns are very important. That why you should prefer to use arrays. They have such predictable access patterns
and can be optimized at hardware level (prefetcher).

### Arrays 2 (Semantics)

There is a gotcha with ``range`` loops.

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
Here, ``for i := range friends`` is using **pointer semantics**. What does that mean?
That means that **``friends is NOT copied``** during the operation.
``friends[1] = "Jack"`` will mutate the original array and inside the ``for loop`` we will see ``Jack`` instead of ``"Betty"``.

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

The only difference here is the ``v`` variable. **Now the ``range loop`` will use value semantics**. That means that the array is **copied!**.
So, we are still mutating the original array BUT the output will be ``"Betty"``

### Array 3 (Slices)
``slice`` is very similar to ``C++`` vectors.

- built in function ``make`` to create ``slices``

- ``slices`` **should not be shared**. They are designed to be used within value semantics context.

### Array 3 (Appending Slices)

- slice initialized with ``var`` => *zero value ``slice``* (nil slice)

- slice initialized with ``data := []string{}`` => *empty slice*

Since ``var []string`` **DOES NOT** mean *empty slice* it should be returned when an error happened. This way
the consumer will not be confused about  the return result.

- start with *nil slice* when you do not know the length of array.

There is notion of ``capacity`` when it comes to ``slices``. You can specify  ``capacity`` as 3rd argument of ``make`` when creating slices.

``append`` call has a cost. When you can prefer specifying ``length and capacity`` and using index for assigment.
