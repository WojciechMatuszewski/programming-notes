# Step Functions

- `Express Workflows` delivers _at-least-once_ workflow execution.

* `Express Workflows` does not support _activities_ and _callback_ patterns.

- There is no visual information within AWS Console UI that the state was retried.

* be wary when implementing `Task` as `Catch` clauses. The input will be the error object, not the initial input that caused the _exception_.

- `Choice` is very powerful, use it!

* when using `Parallel` watch out for duplicate names. The **state names have to be unique globally, not per branch**.

- with `Parallel` any failure with _cancel_ any other branches that might or might not be affected by the failure.

* remember that the **output from the `Parallel` state will be an array!**.

- you can control the concurrency of `Map` type, use `MaxConcurrency` parameter.

* you can start execution of state machine using `API Gateway`. Please note that the invocation is async, you will not get the response back.

- you can invoke your step functions when **reacting to `CloudTrail` events** with `CloudWatch`. Imagine someone uploading file to s3 and then invoking your workflow!.

* you have to create `Activities` a head of time. They are very similar to `callback` pattern.

## `Callbacks` vs `Activities`

`Callbacks` are event based. With `activities` pooling is used for communication. You should **preffer `callbacks`**.

Keep in mind that with the `callback` you have access to the `context` object which contains useful information about the state machine itself (as well as the `task token`).

## Blue-green

You should not use only `arn` for referencing lambda functions but also include function version. This is because **`step functions` are immutable!**.
Imagine a scenario where you have some executions running and you update your lambda function code. If you use naked `arn` values, that execution will most likely fail due to lambda logic change.

## Patterns

### `try-catch`

You can wrap your whole workflow within a `Parallel` state. With that you do not have to duplicate error handling logic on every step but apply it to the whole `Parallel` state instead.

Remember that you can also introduce retry logic on a single task level and mix and mach your options. Pretty powerfull overall.

### Saga

This pattern is for implementing retries when dealing with transactions. Picture vacations booking where you have to book a flight a car rental and a hotel. All of these systems might fail.

![saga](./assets/saga.png)

As you can see there is also `self-recursion` going on. This is to make sure that the _cancel_ actions are retried.

### `de-dupe`

Remember that the `startExecution` API is idempotent. This mean that you can use step functions to make sure you are processing stuff only once.
You should consider using some kind of `ID` or `MD5` has for the execution ID.
