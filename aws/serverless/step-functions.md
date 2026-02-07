# Step Functions

- `Express Workflows` delivers _at-least-once_ workflow execution.

- `Express Workflows` does not support _activities_ and _callback_ patterns.

- There is no visual information within AWS Console UI that the state was retried.

- be wary when implementing `Task` as `Catch` clauses. The input will be the error object, not the initial input that caused the _exception_.

- `Choice` is very powerful, use it!

- when using `Parallel` watch out for duplicate names. The **state names have to be unique globally, not per branch**.

- with `Parallel` any failure with _cancel_ any other branches that might or might not be affected by the failure.

- remember that the **output from the `Parallel` state will be an array!**.

- you can control the concurrency of `Map` type, use `MaxConcurrency` parameter.

- you can start execution of state machine using `API Gateway`. Please note that the invocation is async, you will not get the response back.

- you can invoke your step functions when **reacting to `CloudTrail` events** with `CloudWatch`. Imagine someone uploading file to s3 and then invoking your workflow!.

- you have to create `Activities` a head of time. They are very similar to `callback` pattern.

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

AWS Step Functions expose two types of service integrations. The _optimized_ and the _sdk_ service integrations. The _optimized_ service integrations predate the _sdk_ ones but a couple of years.

With the addition of the _sdk_ service integrations, you can wire up many AWS services together via AWS Step Functions. This was not the case with _optimized_ service integrations, as only a few of them exist.

### Optimized integrations

The optimized service integrations are "deeper" in functionality than the sdk ones. A good example would be the AWS Lambda optimized integration, where the result is returned to you as JSON and not as an escaped string.

#### The differences between SDK and optimized integrations

- The SDK integrations expose the whole API of a given service. The optimized integrations are not.

- The optimized integrations are "deeper" – they have additional functionality, like parsing the results, baked in.

- For some services, the error handling differs (depending on the service SDK). The optimized integrations will introspect the result and force the step to fail if the call fails, the SDK integration will return the response from the API, and if the status is not 4xx or 5xx, it will mark the step as successful (think AWS Lambda optimized vs sdk integration).

### SDK integrations

The SDK integration is a bare-bones way of calling the service API using the ASL language. There are no additional niceties, as in the case of optimized service integrations.

With the SDK integrations, you have the whole service API surface to pick from. For example, you can perform the DynamoDB `Scan` or `Query` calls, which is impossible with optimized integration.

## Waiting for other tasks

There are two ways you can orchestrate "waiting" for other services to finish their work during the StepFunction executions.

One way is to use `.sync` suffix, another one is to use the `.waitForTaskToken` suffix. Let us explore both of them.

### The `.sync` suffix

The StepFunctions service is able to make an asynchronous process seem synchronous. If you are familiar with JavaScript, it's almost like putting
_await_ in front of the asynchronous function. The flow of the code seems synchronous, but in reality we are performing an asynchronous operation.

How does it work? According to [this video](https://youtu.be/MqVqjn3sZVg?t=2464) the StepFunctions service will listen for a specific EventBridge event, and then handle it from there.

First of all, [not every service support this kind of way of invoking it](https://docs.aws.amazon.com/step-functions/latest/dg/connect-supported-services.html).

Next, you have to consider some caveats with how to handle tasks abortions. Please [refer to the docs](https://docs.aws.amazon.com/step-functions/latest/dg/connect-to-resource.html#connect-wait-token) - the _Run a Job_ section.

### The `.waitForTaskToken` suffix

This, again, is to make an asynchronous process seem synchronous. But the difference between `.waitForTaskToken` and `.sync` is how much work YOU have to do.

The `.waitForTaskToken` is designed to be generic. You have to notify the StepFunctions service that the task is done and the execution should continue (as opposed to the `.sync` suffix where it was the service who decided when the execution should continue).

To signal to the service that the task is done, you will need to call `SendTaskSuccess` or `SendTaskFailure` that the StepFunctions API exposes.

## Creating loops to check for status (_might_ be an anti-pattern)

> I wrote this section before I realize we can replace all the looping with events. Read on to learn how!

Sometimes, you might need to halt the machine execution until a given resource changes status. It might happen that you cannot integrate that resource with the `waitForTaskToken` logic and have to implement the pooling logic manually.

In such scenarios, you need to create a loop in the state machine definition. You can read more about such a machine [here](https://docs.aws.amazon.com/step-functions/latest/dg/tutorial-create-iterate-pattern-section.html#create-iterate-pattern-step-3).

### Replacing looping with events

Whenever you consider introducing a "check loop" in your Step Functions, make sure you have exhausted all other options first.

**You can use a combination of DynamoDB, AWS Lambda, and EventBridge** to handle this as well. [This blog post explains how](https://theburningmonk.com/2026/02/the-anti-polling-pattern-for-step-functions/).

Essentially, you write the `TaskToken` into DynamoDB. Then you can start a job, perhaps triggered via DynamoDB Streams. When the job finishes, you send an event to EventBridge, which invokes an AWS Lambda function. That AWS Lambda function then looks up the token and resumes the Step Function via `SendTaskSuccess`.

**While you might consider this approach "proper," it is quite complex.** The reason to avoid the "check loop" in Step Functions is cost. **While the monetary cost of a "check loop" is X, you also have to consider the cognitive cost of doing things the "proper" way.**

Usually, this is not a problem. The comparison is skewed towards the "right" way since you write it once and then forget about it. BUT, when something breaks, how well are you able to diagnose the issue?
