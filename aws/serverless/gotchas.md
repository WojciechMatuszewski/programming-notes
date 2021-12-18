# Serverless Gotchas

## Lambda container freeze and connection reuse

Let's say you are making an HTTP requests within your lambda function.
You've heard that you should reuse connections to ensure performance.

Your code might look as follows

```js
import https from "https";
const agent = new https.Agent({ keepAlive: true });

export const handler = () => {
  https.get({ agent });
  // .. rest of the code
};
```

The code works, and is faster one some requests.
You decide to turn on _X-Ray_ and orchestrate the agent so that you can see the traces.
With that you notice that some requests are marked as "failed" with an error that says _Socket hang up_.

Turns out this error is a natural consequence of how the Lambda container reuse works.

### Lambda container reuse

For performance reasons, Lambda will re-use the containers that run Lambdas code.
The container can live up to 2 hours.

Important part is that **Lambda Service will freeze the container after your code is done executing, that also includes all open connections**.
Looking at our example, that would mean that the agent's underlying connection will also freeze.

So the connection is frozen. Next request comes, the container is unfrozen. The Agent could not know if the connection is closed or not,
because everything was frozen. Tries to make the request, the connection might be closed (in the error scenario is). If so, you will see the error being reported.

### How others handle this

To the best of my knowledge, other clients just retry when this happens.
This seems to be a standard practice to do so.

### X-Ray only shows the error

This is a weirdness of X-Ray. From my experience, if you orchestrate the http client with X-Ray, the service will only
show a trace for the first outgoing request. If that original request fails, and is successfully retried, the AWS Console, will still be displaying the failed request.

### Potential solutions

1. Retry the request and carry on

2. Initialize the connection / agent within the handler

3. Consider using RDS Proxy or Redis for connection management

## Logging EB events into the log group

This might come in handy while debugging. Some target is not receiving events and you want to know how the payload looks like, maybe your filter pattern is not right?

To enable this kind of setup without using the console (keeping everything as IAC), do have to:

- have a log group which name starts with `/aws/events/`

- **add a resource policy to that log group to allow EB to deliver logs to it**

- have a policy on the rule itself allowing to write logs to that log group

No. 2 and 3 are limitations of the `CloudFormation` currently.

You will have to use a **custom resource to add the necessary resource policy**.
You can view this thread for more info https://github.com/aws-cloudformation/aws-cloudformation-coverage-roadmap/issues/351

The feature itself is already implemented in CDK

## Log group and the ARN

Please watch out whenever you reference the ARN of the logGroup anywhere.
The **ARN provided by the CloudFormation has `*` at the end**. This **makes this ARN useless for some services**.

I did not find any other way around this rather than computing the ARN manually.

## My CloudWatch rules are not working

Usually there are 2 causes for this

1. Insufficient IAM permissions to invoke the target
2. CloudTrial is not configured correctly
3. Rule is not created in a correct region

### Insufficient IAM permissions to invoke the target

To properly configure the target for your _CloudWatch event rule_ you will need to specify a role that will be used when a given _rule_ is triggered.
If you do that through the console, that role is automatically created for your. Some services, like _CloudWatch logs_ require you to specify _resource based policy_ along with the role.

To debug the issues, I would suggest digging into _CloudTrial_ where you can filter API calls made by _events.amazonaws.com_.
These will contain info why your event is not triggered - **but only if your filter pattern is correct**.

### _CloudTrial_ is not configured correctly

If your account lives within an _AWS organizations_ context, you might see a _CloudTrial_ already created for your account.
This is most likely an _Organization CloudTrial_. These are pretty similar to the _Regular CloudTrial_ but in the context of events there is a significant difference.

Mainly, **if you create _CloudWatch Rule_ that is based on events from _CloudTrial_ and the only _CloudTrial_ you see in your account is the _Organization CloudTrial_, your rule will not work**.
You will see events in the _events history_ but your rule will not be triggered. You have to create _CloudTrial_ for your account to make your rules (that are based on _CloudTrial_) work.

While working with events produced by **services that are global, your _CloudTrial_ has to include global services**.

### Rule is not created in a correct region

If your event is produced by a **service that operates globally**, eg. IAM, your rule **must live in us-east-1 region**.
This is easy to overlook as we do not usually think in terms of regions.

## I cannot get a hold of `Task.Token` when using sfn and Lambda integration

I was recently surprised by this so here it goes.
**To be able to specify the `Task.Token` within your `Parameters` you have to use `.waitForTaskToken` type of integration**.

I have no idea why it is the case. When I tried specifying the whole `Context` with `.sync` (default) execution model,
only input was there.

Of course, it is written in the docs, so I cannot blame AWS, but either way it is kinda weird for me.

### Adding intermittent state to extract context details

What if you do not want to use `.waitForTaskToken` semantics within your Lambda task and still have the access to the context object?
In this case you will need to create intermittent `Pass` state which extracts the context object for your.

The definitions for the workflow would look something like this

```json
{
  "StartAt": "PreEvent",
  "States": {
    "PreEvent": {
      "Type": "Pass",
      "Parameters": {
        "context.$": "$$"
      },
      "Next": "Event"
    },
    "Event": {
      "Type": "Task",
      "Resource": "arn:aws:states:::lambda:invoke",
      "Parameters": {
        "FunctionName": "${EventLambdaName}",
        "Payload": {
          "Input.$": "$.context.Execution.Id"
        }
      },
      "End": true
    }
  }
}
```

Notice that I'm not using `$$` within the `Payload` block of the Lambda task.
All the context is available to me because I've passed it to the previous `Pass` state as input (using the `Parameters` block)

**Please note that I still do not have access to the `Task.Token` variable. It is only available when a given task is of type `.waitForTaskToken`**

## CFN is not resolving my secrets manager reference. Help

I've recently encountered this at work and found the situation interesting, hope you find it interesting as well.

Whenever you are working with services like `SecretsManager` or `ParameterStore` you most likely do not want to resolve their values statically whenever your `CloudFormation` template is created. Otherwise, these would be visible in plain text! Ideally, these would be resolved at deployment time by `CloudFormation` itself.

Do carry out that process, `CloudFormation` uses references that looks similar to this one:

```shell
{{resolve:secretsmanager:arn:aws:secretsmanager:us-east-1:567235585022:secret:secret-RhKGCW}}
```

The idea is to these be visible in your `CloudFormation` template whenever it's created, then the service would take care of resolving them whenever a deployment happens. All good.

At work, I was aware of this fact and thus I though that my template would be deployed nicely, without any issues. Sadly to my surprise, the resource I deployed, instead of resolved `SecretsManager` value, was passed the reference to that value. As if the `CloudFormation` decided not to resolve the reference in this particular case.

**It turns out, the `CloudFormation` will resolve those references only in particular contexts**. That is **not all resources and their properties support the resolution of such references**. In my case, deploying custom resource and passing the reference as a property on the resource parameters, the resolution was not supported.

I found [this documentation page](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/dynamic-references.html) very helpful. It clearly states that "Dynamic references for secure values, such as secretsmanager, aren't currently supported in custom resources.".

### Solution to the problem

TBD, I still have no idea :D

## Many layers of AWS Lambda retries

Your Lambda function could invoked asynchronously or synchronously depending on the service the event originates from.

Lambda service might or might not retry the event if the handler fails.
These kind of retries should not be confused with the retry policy that one can configure on the services that push events to the Lambda service.

### APIGW synchronous invocation

I would argue that this is the most common way your lambda function be invoked.
Since the **invocation is synchronous, neither APIGW nor Lambda service will not retry your request for your if the handler fails**.

It is completely up to caller on how to deal with the retries.

### Asynchronous Lambda invocations

Lambda can be invoked asynchronously, either directly via the API by specifying the `InvocationType` of `Event` (SQS queue is used underneath)
or via various services like _SNS_, _S3_, _SES_, _CloudFormation_, _CloudWatch / EventBridge_, _AWS Config_, _AWS CodeCommit_.

**AWS Lambda asynchronous invocations are, by default, retried twice. That is to say that if your handler fails, the Lambda service will retry that event by invoking your handler again**.

One could specify the retry settings on the integration level as well. **These settings only apply to the event delivery and has nothing to do with how many times your handler is retried**.

### Lambda + EventSourceMapping

This is where the _EventSourceMapping_ pools for events from specified service. This type of integration is usually used whenever we are dealing with services that operate on the notion of a stream.

**Since it's the EventSourceMapping that invokes your function, it's the EventSourceMapping that controls the retries**.
This means that the default behavior of asynchronous invocation retry does not apply here.

## Role is invalid or cannot be assumed

You will see this error whenever you first delete the CDK bootstrapped stack and **then** try to delete the stack deployed using the CDK (associated with the bootstrap stack you have just deleted).

Remember that the CDK bootstrapping process creates the roles for the CloudFormation to assume when deploying your stack. You can change this behavior, and I've written an article on how to do that: https://dev.to/aws-builders/deploying-aws-cdk-apps-using-short-lived-credentials-and-github-actions-59g6

How to get around this problem? The solution is **to re-create the IAM role with the same name**. Remember about the trust relationship – it has to point to the CloudFormation service.
