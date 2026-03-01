# FEM Permissions Systems At Scale

1. [Notes based on this workshop](https://frontendmasters.com/workshops/permission-systems/).
2. [GitHub repository](https://github.com/WebDevSimplified/fem-permission-systems-that-scale).
3. [Course website](https://webdevsimplified.github.io/fem-permission-systems-that-scale/).

## Authentication vs. Authorization

Ahh, the never-ending source of confusion!

- **Authentication answers "who are you"**. Imagine _having_ a badge to enter a building.

- **Authorization answers "what can you do"**. Imagine _having a certain permission-level_ to enter a specific floor in a building.

So, when I log-in to a website, I'm going through the process of _authentication_. At the end of the process, I usually get some kind of token back. This token is then send alongside any requests to other services to perform _authorization_ on a given resource.

From my experience, **one of the biggest issues is when engineers try to implement both _authentication_ and _authorization_ together**. This results in messy code and very hard to reason about logic that is hard to talk about.

Another issue I've seen is **engineers using _authorization_ and _authentication_ interchangeably**. This is problematic because it creates confusion.

**Authorization code, usually, changes a lot**. Every time you add a new feature, you might need to amend the authorization checks. In contrast, the **authentication code rarely changes**.

## The Initial Data Model

We started with three entities in the database. On the users, we have the "role" attribute which is an enum. I'm always weary of enums. I wonder how will this scale. What if I want to attach multiple permissions to a single user?

Similar to the `status` and `isLocked` on the document entity. Should those be two separate fields? If we allow multiple values in `status` we could have a document with `["draft", "locked"]` status which I think is much better for extensibility.

## Main Permission Mistakes

1. Scattered checks. This one is **so real**. I've seen those checks get very hard to reason about.

2. Missing permission checks. Builds on the first one. If you have checks duplicated across multiple files, it's _very easy_ to miss adding them.

3. Inconsistent logic. Again, builds on the first one.

Have you ever seen an issue where users can _edit_ but can't save? That's a signal that your permission checks suffer from those issues.

**Kyle argues that missing permission checks are WORSE than anything else**. I agree. IMO it is better for users to not be able to do something, which you can fix pretty fast, rather than everyone having the ability to do something that only a handful of users should be able to do.

---

## Random Next.js learnings

The project we are working on throughout the workshop is using Next.js. Here are some things that stood out to me.

1. The `(<directory_name>)` folders in the `app` directory are will NOT create a new route path. These are so-called _route groups_ that you can use to have multiple `layout.tsx` files per routes without creating additional paths.

Finished part 1 -8:19. Start fixing the issues
