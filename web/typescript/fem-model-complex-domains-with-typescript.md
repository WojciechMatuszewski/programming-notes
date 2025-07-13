# FEM Model Complex Domains with TypeScript

Course covers the concept of _domain-driven design_ and how to implement it using TypeScript.

See [this repository](https://github.com/mike-north/peashoot#) for examples and code samples.

## Value Objects

A _value object_ represents an **immutable** "bag of data" without any _identity_.

- They do not have an `id` attached to them.

- Two of them are equal if their attributes are the same.

  - You can't compare _entities_ like that!

- They have their own validation logic.

  - The R in RGB can't be greater than 255.

### Examples

Some examples of _value objects_ include:

- Money, where we store the amount and currency.

- Color, where we store the RGB values.

### Usage

You would create _value objects_ to help you express business logic within the application. If your application operates on colors, using primitive types like strings might not be suitable for the task.

## Entities

An _entity_ represent an _identity_, like person or a dog. _Entities_ can have relationships between them, and play a key role in your application.

- Have an unique `id` attached to them.

- **Are mutable in time**, just like a person could change their views on a particular issue.

- The **are equal only if they have the same `id`**.

### Examples

Anything that you can attach an _identity_ to: a person, a dog, your grandmother and so on.

## Exercise: the "Temperature Date Calculator"

First, we started talking about what _entities_ and _value objects_ we have to create to make the calculator function.

- We definitely want some kind of way to represent a temperature. We opted into creating a _value object_ to represent that data.

  - Thing to note: we do not want to store the temperature value in multiple units. We probably want to store it in Kelvin and translate it when needed.

- We also want to know a range of temperatures for a given location. We opted into creating a _value object_ to represent that data.

- We also want to represent how temperate ranges change over time.

  - Since this data _can change_, we opted for an _entity_.

## Exercise: Talking with domain experts

To create _value objects_ and _entities_, you must understand what the _domain_ is about. Without having a deep understanding of what problem you are trying to solve, you won't be able to "connect the dots" and write code that actually solves the problem.

**To understand the _domain_, you should consider talking with _domain experts_**. These are people who know A LOT about the thing you are working on.

It all depends on what you are working on:

- For financial applications, you might want to talk with accountants and bankers.

- For medical applications, you might want to talk with people working at the hospital.

When talking with a domain expert, **your job is to understand and translate their knowledge to a set of _entities_ and the relationships those entities have with each other**.

- You MUST NOT use technical jargon.

- You MUST NOT "force" a solution on the domain expert. After all, they are the expert, and you are trying to learn from them.

- You MUST stay curious and open.

- Visualizing things _really_ helps.

  - It allows everyone to have a broader view of the problem and promotes collaboration.

## Bounded contexts

You can think of a _bounded context_ as a set of related concepts that all share a common "theme" or rules and terminology.

Consider talking about "orders". When you are in "sales context" that might mean customer's purchase. If you are in "logistics concepts", it might refer to a shipping instructions.

_Bounded contexts_ allow everyone to "speak the same language" and avoid ambiguity. There is no reason to clarify some of the well-known terms, because everyone already assumes what they mean and that they mean that exact thing.

## Aggregates

I understand _aggregates_ as clusters of associated objects that one should treat as a single unit for the purpose of data changes.

Let us say you want to "communicate" between bounded contexts. You would not modify the _price value object_.

Consider any gym application. You have everything related to a workout (a single _bounded context_) and everything related to the social aspect of that application (like and comments that pertain to a given workout).

The _aggregate root_ would be the workout itself. You would modify an exercise you did _through_ the workout, and not via accessing the exercise directly.

**The main benefit of going "through" a "single point of entry" is the ability to enforce constraints and ensure consistency**.

You most likely would have some logic related to validation in that _workout_ entity. If you did not go "through" the workout entity to update the exercises, you might, by accident, introduce a data corruption.

---

- TIL about `git lfs`. Using `git lfs` allows you to store _pointers_ to large files in a separate server. Then, when you are ready, you can pull those assets locally.

  - **The main benefit here is the reduced size of the repository**. All your git operations will be much faster!
