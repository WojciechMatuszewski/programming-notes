# Everything I know about X-Ray

## X-Ray daemon

The _AWS X-Ray SDK_ does not sent the traces directly to _AWS X-Ray service_. It communicates with a background process (the daemon).
This process is the one that communicates with _AWS X-Ray service_ and sends the traces. The [official _AWS X-Ray_ documentation explains the notion of an daemon very well](https://docs.aws.amazon.com/xray/latest/devguide/xray-daemon.html).

**Some services run the daemon for you** – mainly the _AWS Lambda_ and _Elastic Beanstalk_. Note that these services offer native (service-level) integration with _AWS X-Ray_.

If you are familiar with _AWS Lambda extensions_ it might seem that, in the context of _AWS Lambda_, the daemon is run as an extension. That might or might not be the case.
I could not verify or deny that assumption.

## Testing

_AWS X-Ray_ needs the _root segment_ to function. Usually this "root" segment is derived automatically by _AWS X-Ray_ – it detects the context it is run in.
Since the unit / integration tests most likely could be run locally on your machine _AWS X-Ray_ will not be able to automatically generate that "root" segment in that case.

By default, when the _root segment_ is not present an error will be thrown if you attempt to trace an SDK call or add a subsegment via manual tracing.

> Failed to get the current sub/segment from the context.

You will not be able to run your tests without either mocking the _AWS X-Ray_ SDK or configuring the `contextMissingStrategy`.
I really recommend you do the latter. [Please refer to the documentation in regards to the implementation](https://github.com/aws/aws-xray-sdk-node/tree/master/packages/core#context-missing-strategy-configuration).

### X-Ray daemon

With the `contextMissingStrategy` or `AWS_XRAY_CONTEXT_MISSING` environment variable set you do not have to worry about the SDK trying to send data to the _AWS X-Ray daemon_.
There is nothing else to disable in this context.

## Missing traces

Remember that _X-Ray_ samples your traces. Here is the [reference documentation about sampling](https://docs.aws.amazon.com/xray/latest/devguide/xray-concepts.html#xray-concepts-sampling).
Usually this is a good thing. We would not want to send too much data to _X-Ray_. The trace indigestion costs money!

For some services like _AWS Lambda_ the [sampling is configured for you](https://docs.aws.amazon.com/lambda/latest/dg/services-xray.html).

> The default sampling rule is 1 request per second and 5 percent of additional requests. This sampling rate cannot be configured for Lambda functions.

### Still nothing

You configured the sampling? Still nothing?

- **Depending on the service you are using, you might need to configure IAM**.
  The `xray:PutTelemetryRecords` and `xray:PutTraceSegments` are required for the service to be able to communicate with its backend properly.

- When **using _AWS Lambda_ you will need to set the `tracing` parameter to `active`**. [Refer to this documentation during the implementation](https://github.com/aws/aws-xray-sdk-node/blob/master/packages/core/README.md#usage-in-aws-lambda)

### Instrumenting AWS services (SDK)

If you are using **_NodeJS_** it might happen that the "default" way of capturing an SDK client will not work.
According to my research, some of the clients are built differently than others. Because of this difference, you might need to configure instrumentation for that service differently.

A good example of what I'm talking about is **the `DocumentClient`**. [Refer to this GitHub thread](https://github.com/aws/aws-xray-sdk-node/issues/23) for implementation details.

## Confusing statuses

Within the _AWS X-Ray_ console there is the `RESPONSE` column.
You might think that if your API returns 4xx or 5xx the `RESPONSE` column will show those statuses. Sadly this is not the case.

**In the context of _AWS Lambda_** the **`RESPONSE` column shows the status of the _AWS Lambda service_ and NOT the statusCode your API returned**.
I'm not going to lie. This is quite confusing behavior that tripped me many times.

You can read more about this issue (and other issues) in [this great blog post](https://theburningmonk.com/2017/06/aws-x-ray-and-lambda-the-good-the-bad-and-the-ugly/).

## Custom sampling

It is completely possible to configure the sampling rules. **The sampling rules can even be configured PER network request path basis**.
Having so much control is neat. I'm kind of disappointed in myself that I was not aware of the controls _X-Ray_ exposes.

Here is a snippet of how one might change **HTTP sampling rules** for Node.Js base application.

```js
import AWSXray from "aws-xray-sdk";

AWSXray.middleware.setSamplingRules({
  default: {
    fixed_target: 0,
    rate: 0,
  },
  version: 2,
});
```

As for **AWS SDK calls** (the one you instrument using `captureAWSClient` or `captureAWS`) the syntax is a bit different.
We have to specify a **whitelist** that contains specific services and their operations.

The following snippet demonstrates the usage.

```js
import AWSXray from "aws-xray-sdk";

AWSXray.setAWSWhitelist({ services: {} });
```

According to my experimentation **the above setting does not mean "disable tracing for all services"**.
The **tracing will still work if you instrument a service with _X-Ray_**.

```js
import AWSXray from "aws-xray-sdk";

AWSXray.setAWSWhitelist({ services: {} });

import SQS from "aws-sdk/clients/sqs";
const sqs = AWSXray.captureAWSClient(new SQS())

// The following call will be traced
await sqs.sendMessage({...}).promise()
```

What will not be "traced" are the **additional parameters**. In the case of _SQS_ that might be the name of the queue. For _DynamoDB_ that might be _TableName_.

Overall I find this API a bit confusing. I would expect the trace not to be there. Maybe I'm doing something wrong?

## With AWS Lambda

_X-Ray_ is deeply integrated with _AWS Lambda_.

First of all, **regardless of the `Tracing` configuration** _AWS Lambda service_ **runs the _X-Ray daemon_ in the background**.
If you are not sure what _X-Ray daemon_ is, refer to the _X-Ray daemon_ section.

What about the `Tracing` configuration? What is the benefit of having it set to `Active`?

- _AWS Lambda service_ will check the function _execution role_ for you and add missing permissions if that is needed.
  I'm referring to the permissions encapsulated within the `AWSXRayDaemonWriteAccess` managed policy.

- Apply the sampling rules automatically. Otherwise you would be using the default sampling rules which may or might not be sufficient for you.
  **Please note that changing sampling rules directly using the _X-Ray_ SDK will not work**. You cannot change the sampling rules the _AWS Lambda service_ applies.
  The \*\*sampling rules the documentation mentions will kick-in only after

### Using X-Ray tracing outside of the handler

## X-Ray and retries

From the observability perspective is is vital to have a holistic view on what is happening inside your application.
As we all know, network operations might fail. In such case we usually retry the requests.

_AWS X-Ray_ is not perfect in this department. You can visualize some of the retries but not in all occasions.
Let us dig a little bit different.

### Retrying AWS SDK calls

First let us tackle the situation where **the parameters of the API call are invalid / API returned an error**.
Here is one example.

```js
// db is instrumented via .captureAWSClient API
return await db
  .get({
    TableName: "I do not exist",
    Key: {
      pk: "1",
    },
  })
  .promise();
```

The `TableName` will not be accepted as a valid name of a table by _DynamoDB_ API.
If you retry the request with a valid `TableName` **you will be able to see that visually in _AWS X-Ray_ console**.

Let us switch to a situation where **the socket itself times out**.
Here is an example simulating that.

```js
// ddb is instrumented with AWS X-Ray
const db = new ddb.DocumentClient({
  httpOptions: { timeout: 1 },
  maxRetries: 3,
  logger: {
    log: console.log,
  },
});

return await db
  .get({
    TableName: process.env.TABLE_NAME,
    Key: {
      pk: "1",
    },
  })
  .promise();
```

In the **_CloudWatch_ logs you will see that the request was retried three times** – `[AWS dynamodb undefined 0.68s 3 retries] ...`.
In the **_AWS X-Ray_ console you will only see a singular request annotated as "fault"**.

### Retrying HTTP calls

We have a bit more control when it comes to HTTP calls. That control will come in handy soon.

If we were to instrument the `https` _NodeJs_ module as follows:

```js
import AWSXray from "aws-xray-sdk";
AWSXray.captureHTTPsGlobal(https);
```

By default **when your request times and is retried _AWS X-Ray_ will not show another trace for the second request**.
What you will see in the **console is the status from the first request – a "fault"**. This is very unfortunate and misleading.

What if I want to see traces for all the request, even the ones that were a result of a retry?
In this case **you have to delete the appended `X-Amzn-Trace-Id` header before retrying the request**. This header is the reason _AWS X-Ray_ thinks you are making the same request.

Here is an example of doing so using the [`got` npm package](https://duckduckgo.com/?q=npm+got&t=brave&ia=software).

```js
import AWSXray from "aws-xray-sdk";
AWSXray.captureHTTPsGlobal(https);

import got from "got";

const gotInstance = got.extend({
  hooks: {
    beforeRequest: [
      (options) => {
        delete options.headers["X-Amzn-Trace-Id"];
      },
    ],
  },
});

const result = await gotInstance("URL", OPTIONS);
```

With the _trace header_ not present at the time of making the request, _AWS X-Ray_ will be forced to create a new one.
**This ensures that no matter what happens the retried request will be traced**.

TODO: The missing trace log

## Additional reading

- https://theburningmonk.com/2017/06/aws-x-ray-and-lambda-the-good-the-bad-and-the-ugly/
