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
