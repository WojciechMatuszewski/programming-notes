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

### Decoupling

- start with the concrete, avoid creating interfaces first

* use **layered approach**:
  - **primitive**: knows how to do 1 thing only. This code has to be testable!
  - **lower-level**: sits on top of **primitive**. This code has to be testable!
  - **high level API**: ease of use for developer, this layer may require integration tests.

- **ONLY focus at the problem at hand**. Get stuff into prod then decouple to make it better.

* **optimize for correctness**. Always, always **focus on the simplest solution!**. Until you have problems with performance you do not have to min-max for it!.

- lower level-api usually consist of things that are not exported

* whe designing an API, **try starting with functions, not methods**. Methods sometimes might not be transparent enough. With **functions, the consumer always knows the required arguments / data**. With methods that is not the case. An example:

```go
type User struct {
 name string
 email string
}
func (u *User) SendEmail() {}
```

This code seem like a good idea **but in reality this piece of code is horribly designed**. There is **zero information to the consumer of what NEEDS to be present on the User `struct` for `SendEmail` to succeed**. Maybe we need only `name`? or maybe we need both `name` and `email`. This is where production bugs are introduced!.

Much better would be writing a plain function. But **even writing a function might lead to bugs**. Lets consider the following function:

```go
func SendEmail(u *User) {}
```

Above function is just a method. A method in disguise which is not really helping. `SendEmail` should look as follows:

```go
func SendEmail(name string, email string) {}
```

Now we are **focusing on concrete data** and **are transparent to the consumer**.

- **concrete implementation** will **provide you will behavior**. That is why you should **always work with concrete first**.

* use **composition of interfaces**. This is very prevalent while working with standard lib eg. `ReadCloser` and such.

- always **refactor from the lowest level**

* keep in mind that **interfaces are valueless**. They do not really exist. What exists is a concrete data that obeys the interface.

Lets say I have `pull` and `store` functions

```go
type Puller interface {
	Pull(d *Data) error
}
func pull(p Puller, data []Data) (int, error) {
}


type Storer interface {
	Store(d *Data) error
}
func store(s Storer, data []Data) (int, error) {
}
```

Pretty simple code. Now lets say we have `Copy` function which takes `PullerStorer` - composition of interfaces.

```go
type PullStorer interface {
	Puller
	Storer
}
func Copy(ps PullStorer, batch int) error {
    data := make([]Data, batch)

    store(ps, data)
    pull(ps, data)
}
```

You may be thinking that we are passing `PullerStorer` to `store` and `pull` but that is NOT the case. What we are passing is **stored** inside`ps`. Again, interfaces do not really exist!

- to fully use composition you can **embed interfaces inside structs**. This is so powerful!

* **readability review** is very important. Consistency is the key to make sure your code reads well!.

Going back to our `Copy` method. As you might have noticed, **it is hiding the cost of initialization**. It does not prevent fraud and misuse. It's the same thing as with `SendEmail`, we can improve it:\

```go
func Copy(p Puller, s Storer, batch int) error {
	data := make([]Data, batch)

    store(s, data)
    pull(p, data)
}
```

Much cleaner. Usually this sorts of things will come up during _readability review_ of your code.

### Conversion and Assertions

- remember that to make interface hold concrete data, that data has to obey interface shape

```go
type Mover interface {
    Move()
}
type Locker interface {
    Lock()
}

type MoveLocker interface {
    Mover
    Locker
}

// has `move` and `lock` behaviors
type bike struct {}

func main() {
    var ml MoveLocker
    var m Mover

    // works!
    ml = bike{}

    // does not work. There is a mismatch between the shapes
    m = ml
}
```

- **type assertions happen at runtime!**. This is quite important since it may lead to panics and bugs while your software is running.

* type assertions can be used to override a default behavior within your API.

This is usually done with **assertions on concrete data**, lets consider how you would override a `fmt.Stringer` implementation

```go
type User struct {
    name string
    email string
}

// implementing fmt.Stringer
func (u *User) String() {
    return fmt.Sprintf("My name is %q and my email is %q", u.name, u.email)
}

func main() {
    u := user{
        name: "Wojtek",
        email: "email@email.pl"
    }

    fmt.Println(u)
    fmt.Println(&u)
}
```

Now, remember that the behavior _sticks_ to the type. We only overridden a pointer-schematics stringer of `User`. Under the hood, `fmt` package actually performs a _type conversion_:

```go
switch v := p.arg.(type) {
 case Stringer:
  handled = true
  defer p.catchPanic(p.arg, verb, "String")
  p.fmtString(v.String(), verb)
  return
 }
```

### Interface Pollution

- normally the **result of starting with interfaces and NOT concrete implementation**

* factory functions should return the concrete values, it should be left up to the caller to decouple if needed.

- you should really think about using `interface` only for mocking purposes. Favor your local env. for eg. a database.

* interfaces should enable users to provide an implementation detail.

### Mocking

This is a very important topic. **YOU DO NOT HAVE TO CREATE INTERFACES FOR EVERYTHING THAT IS REUSABLE**.

Lets consider a `PubSub` system package:

```go
type PubSub struct {

}

func New() *PubSub {
    return &PubSub{}
}

func (ps *PubSub) Publish() {}
func (ps *PubSub) Subscribe() {}
```

First of all, to test it, you should use `Docker` or whatever and test the REAL system.
Mocking here does not make any sense.

What if I want to test other parts of the system that depend on this package. I should be able to assume it works right?. To do that I need a mock of `PubSub`. **If I need a mock I'm going to create it locally!**

```go
type publisher interface {
	Publish(key string, v interface{}) error
	Subscribe(key string) error
}

// mock is a concrete type to help support the mocking of the pubsub package.
type mock struct{}
func (m *mock) Publish(key string, v interface{}) error {

	// ADD YOUR MOCK FOR THE PUBLISH CALL.
	return nil
}

// Subscribe implements the publisher interface for the mock.
func (m *mock) Subscribe(key string) error {

	// ADD YOUR MOCK FOR THE SUBSCRIBE CALL.
	return nil
}
```

The consumer created mock for themselves and we as package creators did not have to include `interfaces` within our package. **We did not create interfaces because we were assured that the system only hold 1 implementation of PubSub**. If there were multiple systems / possible implementations we would probably include the mock.

## Error Handling

- if you really want to shut down your program due to **integrity issues**:
  - choose panic if you need a stack trace
  - use `os.Exit`

* stay away from `else` statement. It introduces more code to read through and may make it less clear to the reader.

- use `naked switch` / `switch` statement for `if/else` logic

```go

switch {
    case 1 == 1:
    case 2 == 2:
    // you know the deal
}
```

- construct `ErrorVariables` when your API can return multiple error types.

```go

var (
    ErrPageNotFound = errors.New("Page not found")
    ErrPageMoved = errors.New("Page moved")
)
```

In the code, the consumer then can `switch` over the error

```go
switch err {
    case ErrPageNotFound:
    //
    case ErrPageMoved:
}
```

The naming itself is very important. **Start with Err** then the description of the error.

- favour the built-in `error` type unless you cannot provide sufficient context to the caller.

With custom error types you probably will have to do a `type switch` to check which error you got. A good example for custom error types is the `json` package (`Marshal/Unmarshal`)

```go
err := json.Unmarshal([]byte(), nil)
if err != nil {
    switch e := err.(type)
    // code
}
```

The `type` within the `err.(type)` expression is a language feature.

- there is also notion of _Behavior as Context_. Instead of asserting on types, we are going to assert of the behavior itself. This can result in less code.

Lets consider the `net` package. That package exposes a lot of custom error types

```go

switch e := err.(type) {
    case *net.OpError:
        if !e.Temporary() {}
    case *net.AddrError:
        if !e.Temporary() {}
    // and so on
}
```

All we really care about is the information that `Temporary` gives us. It tells us if the system has just lost it's integrity.

We can reduce the amount of code by creating an `interface` which has the desired behavior

```go
type temporary interface {
    Temporary()
}

switch e := err.(type){
    case temporary:
        if !e.Temporary()
}
```

Much better right?

- **always prefer pointer schematics when creating custom error types**. Not following this rule might lead to subtle bugs.

Remember, when comparing values, go is going to compare the concrete value.

```go
// errorString implements the `Error` stuff
func New(text string) error {
	return errorString{text}
}

var Test = New("Bad Request")

func main() {
	if err := webCall(); err == Test {
		fmt.Println(err)
		return
	}

	fmt.Println("Life is good")
}

// webCall performs a web operation.
func webCall() error {
	return New("Bad Request")
}
```

Guess what is the output of the program? `Bad Request`. Why? because we are comparing the concrete values. If we change the implementation of `New` to return pointer of `errorString` that equality will never occur, because we are comparing Addresses!

### Wrapping Errors

- logging as insurance policy, this can cause too much noise when it comes to logging.

* balance signal to noise ration when it comes to logs

- make sure to include enough context within a log (tracing and error body)

* if you decide to pass the error froward, **wrap the error with the context**. You should prefer handling errors as low as possible, but in general it's the developers choice.

- use `%v` to get the user context of an error, the one you defined while wrapping. Use `+%v` to get both user context and the stack trace.

## Packaging

- avoid packages of type `utils` or `helpers` or `models` or similar. Package has to have specific purpose.

* make sure your **packages can only import down!**. This is quite important, makes your codebase consistent.
