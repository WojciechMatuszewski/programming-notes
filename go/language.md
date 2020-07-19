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


