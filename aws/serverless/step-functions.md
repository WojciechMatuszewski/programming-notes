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

`Callbacks` are event based. With `activities` pooling is used for communication. You should **prefer `callbacks`**.

Keep in mind that with the `callback` you have access to the `context` object which contains useful information about the state machine itself (as well as the `task token`).

So basically, think of the `callback` vs `activity` as `push` vs `pool` debate.

## Blue-green

You should not use only `arn` for referencing lambda functions but also include function version. This is because **`step functions` are immutable!**.
Imagine a scenario where you have some executions running and you update your lambda function code. If you use naked `arn` values, that execution will most likely fail due to lambda logic change.

## DR Strategy

Imagine the whole region going down while your execution is in progress. What should you do? Let us look at [this video](https://youtu.be/MqVqjn3sZVg?t=786) for answers.

1. Have an event registry, most likely EventBridge, that would allow you to retry the failed event in another region
2. You could move the execution data from CloudWatch logs to a different region and hydrate the state from there

It is vital to have an active-active setup where the other region is exercised regularly. This way there will be no surprises while switching regions.
In general there is no built in way of doing DR. It's a hard problem to solve for this service.

## JSON Path expressions

You would not believe how powerful the transformations you can do with ASL are. The StepFunctions support full JSON Path syntax.
You can flatten array, filter them, map them. Here is a [video that showcases the power of JSON Path with StepFunctions](https://youtu.be/MqVqjn3sZVg?t=2015).

## Patterns

### `try-catch`

You can wrap your whole workflow within a `Parallel` state. With that you do not have to duplicate error handling logic on every step but apply it to the whole `Parallel` state instead.

Remember that you can also introduce retry logic on a single task level and mix and mach your options. Pretty powerful overall.

### Saga

This pattern is for implementing retries when dealing with transactions. Picture vacations booking where you have to book a flight a car rental and a hotel. All of these systems might fail.

![saga](./assets/saga.png)

As you can see there is also `self-recursion` going on. This is to make sure that the _cancel_ actions are retried.

### `de-dupe`

Remember that the `startExecution` API is idempotent. This mean that you can use step functions to make sure you are processing stuff only once.
You should consider using some kind of `ID` or `MD5` has for the execution ID.

## Service Integrations

- when a fields **ends with `.$`** this means that the **value can use the input or context**.

  ```json
  {
    "Parameters": {
      "ExecutionID.$": "$$.Execution.Id",
      "SomethingFromInput.$": "$.valueFromInput"
    }
  }
  ```

- if you use `.$` on a field name, **you cannot combine paths with strings, it's all or nothing**

  ```json
  {
    // NOT VALID
    "pk.$": "KEY#$$.Execution.Id"
  }
  ```

- you **cannot use `StartSyncExecution` with _regular_ step function**, you have to use the _express_ one for that

### DynamoDB integration

Very useful but a bit "underpowered". Lacks the `Query` API that would open a door for a whole set of new use-cases.
Think fan-out patterns after reading the DDB and much more.

### Waiting for other tasks

There are two ways you can orchestrate "waiting" for other services to finish their work during the StepFunction executions.

One way is to use `.sync` suffix, another one is to use the `.waitForTaskToken` suffix. Let us explore both of them.

#### The `.sync` suffix

The StepFunctions service is able to make an asynchronous process seem synchronous. If you are familiar with JavaScript, it's almost like putting
_await_ in front of the asynchronous function. The flow of the code seems synchronous, but in reality we are performing an asynchronous operation.

How does it work? According to [this video](https://youtu.be/MqVqjn3sZVg?t=2464) the StepFunctions service will listen for a specific EventBridge event, and then handle it from there.

First of all, [not every service support this kind of way of invoking it](https://docs.aws.amazon.com/step-functions/latest/dg/connect-supported-services.html).

Next, you have to consider some caveats with how to handle tasks abortions. Please [refer to the docs](https://docs.aws.amazon.com/step-functions/latest/dg/connect-to-resource.html#connect-wait-token) - the _Run a Job_ section.

#### The `.waitForTaskToken` suffix

This, again, is to make an asynchronous process seem synchronous. But the difference between `.waitForTaskToken` and `.sync` is how much work YOU have to do.

The `.waitForTaskToken` is designed to be generic. You have to notify the StepFunctions service that the task is done and the execution should continue (as opposed to the `.sync` suffix where it was the service who decided when the execution should continue).

To signal to the service that the task is done, you will need to call `SendTaskSuccess` or `SendTaskFailure` that the StepFunctions API exposes.
