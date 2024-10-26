# Programming

General thoughts about programming. This is more technical-focused equivalent of `thoughts.md`.

## The slippery slope of positional arguments

> Based on [this great blog post](https://fhur.me/posts/2024/why-you-should-not-default-to-positional-parameters)

Let us consider the following function:

```ts
activateEntity(1, true, true, false);
```

Ask yourself, what does this function do?

You might answer with: "it activates the entity, right?"

That would be, assuming the function does what the name implies, correct.
But are you able to infer anything else than that? What does the `true, true, false` mean? What about that single number passed as first argument?

If you have trouble answering this question, other might as well. **If the interface of the function poses questions, it means that the interface is leaky or unclear**.
In this particular case, you would need to go to the file the function is declared in, and look at the name of the parameters (do not get me started at "my IDE prepends the names of the parameters as ghost text near arguments" stuff).

How could we refactor this code to be easier to understand? Not in "how this function works" sense, but rather "what can I do with this function" sense?
Well, we can use an object as an argument. **Using an object as an argument makes the function self-documenting**.

```ts
activateEntity({
  id: 1,
  confirmEmail: true,
  isAdmin: true,
});
```

One might argue that passing boolean flags to the function is not a good practice, and I would agree, but let us leave that for another day.

**Notice how it is much easier to infer WHAT the function does, and HOW it might be useful in different situations**. Adding new parameters is also much easier.
We improved readability and extendability of this function in one refactoring, and all it took was to change how many arguments the function takes.

**Positional arguments** are "dangerous" because they **have a momentum of its own**.
You have to change a function, so you pile more and more positional arguments, making the code less readable.

**It always starts simple, but then reality steps in, and you end up with unreadable mess of a function with ten positional arguments**.
