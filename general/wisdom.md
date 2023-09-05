# Programming wisdom

## Write code that is easy to delete, not easy to extend

> Based on [this article](https://programmingisterrible.com/post/139222674273/write-code-that-is-easy-to-delete-not-easy-to).

This article goes in various directions, but the underlying theme is the following: **It is more than okay to delete code. Writing code that is easy to delete should be one of your priorities. You will not get the design right when writing the code for the first time**.

Of course, life is more complicated. There are layers to this premise.

1. **Consider not writing code at all**. Think about writing the code as a last resort. Sometimes, you can avoid writing code by ensuring everyone is on the same page. Think of times you wrote some piece of code and then deleted it shortly after because things changed.

2. **Copying and pasting code is healthy to a certain degree**. You would not want to copy and paste every time, but doing so a couple of times will expose patterns you could generalize. The same applies to writing boilerplate.

3. If you have a lot of boilerplate and are ready to collapse code into a module, do so. Feel free to create abstractions over more complicated code. **Thread carefully when creating an abstraction**. A wrong abstraction will do much more harm than good.

> It is not so much that we are hiding detail when we wrap one library in another, but we are separating concerns: requests is about popular http adventures, urllib3 is about giving you the tools to choose your own adventure.

4. **There is nothing wrong with writing a bunch of code, learning from it, and deleting it afterward**. It is much easier to remove one big mistake than many smaller mistakes scattered throughout the codebase. **Keep the experiments within a specific boundary that you can delete without any issues afterward**.
