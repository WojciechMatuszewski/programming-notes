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
