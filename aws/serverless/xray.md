# Everything I know about X-Ray

## Missing traces

Remember that `X-Ray` samples your traces. Here is the [reference documentation about sampling](https://docs.aws.amazon.com/xray/latest/devguide/xray-concepts.html#xray-concepts-sampling).
Usually this is a good thing. We would not want to send too much data to `X-Ray`. The trace indigestion costs money!

### Still nothing

You configured the sampling? Still nothing?

- **Depending on the service you are using, you might need to configure IAM**.
  The `xray:PutTelemetryRecords` and `xray:PutTraceSegments` are required for the service to be able to communicate with its backend properly.

- When **using _AWS Lambda_ you will need to set the `tracing` parameter to `active`**. [Refer to this documentation during the implementation](https://github.com/aws/aws-xray-sdk-node/blob/master/packages/core/README.md#usage-in-aws-lambda)

### Instrumenting AWS services (SDK)

If you are using **_NodeJS_** it might happen that the "default" way of capturing an SDK client will not work.
According to my research, some of the clients are built differently than others. Because of this difference, you might need to configure instrumentation for that service differently.

A good example of what I'm talking about is **the `DocumentClient`**. [Refer to this GitHub thread](https://github.com/aws/aws-xray-sdk-node/issues/23) for implementation details.
