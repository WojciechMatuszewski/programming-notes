# Serverless Gotchas

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
