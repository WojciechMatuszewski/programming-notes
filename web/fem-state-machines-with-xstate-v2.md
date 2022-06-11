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

<!-- Finished on actions solution -->
