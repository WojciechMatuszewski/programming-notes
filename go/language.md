# Language

## `errors.Unwrap` vs `errors.Cause`

As you probably know, you should *wrap* your errors to provide additional context to them.
Now, you do not have to do it all the time, every time, but in most situations it's a good practice™️.

Sometimes you need to unwind the errors though, and there are two methods to do so.

### The `Unwrap` method

This method is coming from the Go std lib. To my surprise it **does not work recursively**.
This is a drawback for me, I usually have a couple of levels of nesting within my applications
when it comes to errors.

### The `Cause` method

This one is coming from the `pkg/errors`. **It carries out the unwrapping recursively** which is a huge plus.
Definitely will be using this one from now on.

## `pprof` handlers

If you are using the default multiplexer for your http server, the `ServeMux`
type, you can introduce metics to your server by using the `pprof` package.

All you really have to do is to
* Import the package as a side effect
```go
_ "net/http/pprof" // register the /debug/pprof handlers
```
* Register the default multiplexer

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

Now you can create your *dimmensions* if you will

```go
reqNum := expvar.NewInt("requests")
reqNum.Add(1)
```

This is a very powerful technique which enables you to improve observability in your app. 
