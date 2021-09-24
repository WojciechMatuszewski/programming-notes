# Language

## Typed and untyped `nil`

Have you ever tried to use a "plain" `nil` within your program? Like so

```go
package main
func main() {
  var x = nil
  _ = x
}
```

The above would fail with

> use of untyped nil

Interesting, right? But why is that the case. Well it turns out the `nil` has to have an underlying type.

### `nil` as predefined value

It turns out **`nil` is not a keyword!. It's is treated as predefined value**.
That means that you can cause a lot of mayhem by doing

```go
package main() {
  func main() {
    var nil = errors.New("boom")
  }
}
```

This is similar to how you **were able to up to a certain point** break others _JavaScript_ programs by doing

```js
var undefined = "lol";
```

### Zero values of types

The `nil` value is very useful for different types that we have in go

- For _pointers_, methods can be called on nil receivers
- For _slices_, `nil` is a perfectly valid zero value
- For _maps_, `nil` is a perfect for read-only use case
- For _channels_, `nil` is essential to some concurrency patterns (`nil` channel)
- For _functions_, `nil` is just there for completeness sake
- For _interfaces_, `nil` is used to signal that everything is fine (in the case of errors), or to use a default behavior

### The typed `nil` gotcha

The case where `nil != nil` is very particular and has to do with how _interfaces_ are represented.

For the interface to equal to `nil`, the **interface type and the value has to be nil**.
This means that a **`nil` pointer to an interface is not equal bare `nil`**

```go
var p *int
var i interface{}

i = p

if i != nil {
    fmt.Println("not a nil")
}

// Outputs `not a nil`
```

The output might be confusing, but if you think about, it makes sense.
Here the interface is "backed" by pair of `(*int, nil)`, not a `(nil, nil)`. In such cases, it will never be equal to `nil`

This can be especially problematic around returning errors, where you **should never return concrete error types, always return the `error` interface**.

## `errors.Unwrap` vs `errors.Cause`

As you probably know, you should _wrap_ your errors to provide additional context to them.
Now, you do not have to do it all the time, every time, but in most situations it's a good practice™️.

Sometimes you need to unwind the errors though, and there are two methods to do so.

### The `Unwrap` method

This method is coming from the Go std lib. To my surprise it **does not work recursively**.
This is a drawback for me, I usually have a couple of levels of nesting within my applications
when it comes to errors.

### The `Cause` method

This one is coming from the `pkg/errors`. **It carries out the unwrapping recursively** which is a huge plus.
Definitely will be using this one from now on.

## Making sure given struct implements all methods of an interface

I find it very odd (?) that in a language that is typed in a static manner (maybe that is the reason?) there is no `extends` clause.
What the `extend` clause allows for, is to ensure that a given construct implements all methods of a given _interface_.
Since there is no `extends` in Go, how one might ensure that given `struct` implements all methods of an `interface`?

According to the [official Go FAQ](https://golang.org/doc/faq#guarantee_satisfies_interface) one will need to create a special variable that
uses _type casting_ that `struct` on the `interface`.

Here is an example:

```go

type Human interface {
  speak()
  drinkWater()
}

type Myself struct {}

// Cast the `Myself` struct onto the `Human` interface.

var _ Human = Myself{}

```

Since the `Myself` struct does not implement the `speak` and `drinkWater` methods the compiler will thrown an error.

What is more, the same problem occurs when using _interfaces_ in Rust.

## `pprof` handlers

If you are using the default multiplexer for your http server, the `ServeMux`
type, you can introduce metics to your server by using the `pprof` package.

All you really have to do is to

- Import the package as a side effect

```go
_ "net/http/pprof" // register the /debug/pprof handlers
```

- Register the default multiplexer

```go
package main

log.Printf("main: Debug service listening on %s", cfg.Web.Debug)
err := http.ListenAndServe(cfg.Web.Debug, http.DefaultServeMux)
log.Printf("main: Debug service ended %v", err)
```

Keep in mind that the `cfg.Web.Debug` address should be protected, and only accessible by your team.

You can now open your browser and navigate to `debug/pprof/` 🤯

## `exprvar` metrics on http handlers

There is so much tooling built-in to Go ❣️

I already mentioned `pprof`, you did you know you can add your custom metrics there?

You have to import `expvar` package for side effects

```go
import _ "expvar"
```

Now you can create your _dimmensions_ if you will

```go
reqNum := expvar.NewInt("requests")
reqNum.Add(1)
```

This is a very powerful technique which enables you to improve observability in your app.

## Handling errors from deferred functions

We all have written similar code at one point or another

```go
func cleanup() error {
  fmt.Println("cleanup")

  return errors.New("boom")
}

func getMessage() (string, error) {
  defer cleanup()

  return "message", nil
}

func main() {
  message, err := getMessage()
  if err != nil {
    // code
  }
}
```

The problem is that we are not handling the errors that the `cleanup` function produces.
These are usually safe to ignore, like while working with files or similar scenarios, where an error returned from `file.Close()`
would mean we are in deep trouble and that error is the least thing we should be worried about.

But what if you know you _should_ be handling such errors, that they are not the same class of errors like those returned from `file.Close`?
Well, I've seen people do this

```go
func cleanup() error {
  fmt.Println("cleanup")

  return errors.New("boom")
}

func getMessage() (message string, err error) {
  defer func() {
    err = cleanup()
  }()

  return "message", nil
}

func main() {
  message, err := getMessage()
  if err != nil {
    // code
  }
}
```

And this is completely fine, **but this structure lurks a huge opportunity to make a mistake**.
Let us change the code a bit to add call to another function within the `getMessage`.

```go
func cleanup() error {
  fmt.Println("cleanup")

  // Now this function does not return any errors.
  return nil
}

func doAnotherThing() error {
  fmt.Println("doing another thing")
  return errors.New("boom")
}

func getMessage() (message string, err error) {
  defer func() {
    err = cleanup()
  }()

  err = doAnotherThing()

  return "message", nil
}

func main() {
  message, err := getMessage()
  if err != nil {
    // code
  }

}
```

If you run the code ... **the `if err != nil` code in the `main` function will never run!**.
This is because the `cleanup` function overwritten the error since it's deferred. Oops!

So to ensure that you do not make such mistakes, **always check for errors within deferred functions before assigning them to the _named returned values_**.

```go
func getMessage() (message string, err error) {
  defer func() {
    if tempErr = cleanup(); tempErr != nil {
       err = tempErr
    }
  }()

  err = doAnotherThing()

  return "message", nil
}
```

This section was inspired by [this article](https://trstringer.com/golang-deferred-function-error-handling/).
