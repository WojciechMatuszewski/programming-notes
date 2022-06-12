# State Machines in JavaScript with XState v2

## Software Modeling

- State machines can help with software modeling by providing a holistic view of the logic flow

  - Your diagrams could be based on the state machines in your code. This way, the diagram is never out of date!

- By modeling, we also mean understanding and predicting the events your application might produce. UIs also produce events – to communicate with other components or update the state of a given component.

### First exercise

- The goal was to familiarize me with a state machine's mental model. Since I've seen them in code before, I could come up with a solution relatively quickly.

- Like David, I prefer the object notation "way" of creating the state machine.

### Second exercise

- Holy moly, the tooling around `xstate` is excellent. The ability to visualize the state machine is a game-changer.

- I like the API, especially how you can tell if a machine can transition from a given state to another.

## Actions

- Actions are side effects. They **do not happen randomly** (looking at you – some of the `useEffects` I've written).

- You can perform actions on `enter`, `exit`, or the transition itself.

  - David recommends performing actions on transitions. I second his recommendation – it aligns with my mental model of state machines.

### Exercise

- I find the `action` name a bit misleading. I would rather have it named `onTransitionToTarget`.

- The order of action execution is NOT surprising – a huge plus!

- The ability to `rise` an event as a globally exported function is a plus.

- During the solution walkthrough, I've noticed that David changed the `target` from `loading` to `.loading`. I understand the underlying purpose of the change, but it seems weird to me. Luckily this syntax is well documented [here](https://xstate.js.org/docs/guides/ids.html#relative-targets).

## Context

- The `context` property is a container for global data that one might need to access in the state machine.

- You can partially update the `context`. The API is quite lovely!

- You can think of the `context` as part of the state machine state. At least, that is the usage the `context` exercise presents.

### Exercise

- I like the `assign` API, but I can also see how, in large machines, the globally available `context` can get unwieldy — something to keep an eye for in terms of complexity.

## Guarded Transitions

- By default, if an event occurs and a given state has a transition to the next state the event points to, the transition will occur.

  - That is fine unless you want to ensure that you can only transition from a given state when a given condition is met – this is where the _guarded transition_ help.

### Exercise

- When David started to explain the difference between the `guards` and the conditionals we could, in theory, use in the `assign` statements, the concept of guards clicked with me.

  - We COULD define conditionals inside the `assign` statements, but then we would be lying about the fact that the state of our machine changes.

  - The guards prevent the state from changing. The state machine will NOT emit a transition if the guard statement returns false.

  - That is not the case. If we used conditionals inside the `assign` statements, the state machine would emit the transition event.

## Compound States

- _Compound States_ is a grouping of states within another state.

  - You can think of two "root" states of _sleeping_ and _awake_ and then multiple groups of states that live "under" those two "root" states.

### Exercise

- With multiple levels of nesting, it can be pretty hard to have a holistic view of the state machine. This is where visualization tools come in handy.

- The more you nest, the more complex the state machine becomes. You might want to consider splitting state machines into multiple smaller ones, then merging them into one "master" state machine.

- As is the theme with _xstate_, there are multiple ways to get the current state of the machine. You can use `tags`, `can`, or `matches`—an interesting choice of an API.

## Parallel States

- Parallel states allow you to "split" your state machine into multiple working features. For example, David presents two "regions" in our state machine – one that controls the player and the other that controls the volume.

- The feature of a parallel state is similar to having two separate state machines. Sometimes it does not make sense to split a state machine into two state machines – that might lead to unnecessary complexity.

### Exercise

- More nesting!

- Great for organizing different states.

## Final States

- Somewhere where the machine ends. There is no more "work" to do.

- There is a difference between the "root" machine final state and the compound state final state.

- Adding an explicit _final_ state is, in my mind, a good practice. It signals that there is no more work to do for a given "event path".

### Exercise

- I have to admit that the state machine is getting quite complex.

- I wonder how much code we would have to write to have the same functionality as we currently have, but without using the state machine. Probably less, but we would risk having logic-related bugs in our code.

## Overall thoughts

- So many ways to do the same thing!

  - I understand that the current API shape is the result of many incremental changes, but I would greatly appreciate it if there were not three ways to define a state or an action.

- The _xstate_ is very powerful and makes you feel like you are writing a bug-free code. Can there be bugs? Sure! But I feel like by using state machines, there is a lesser chance of introducing one.
