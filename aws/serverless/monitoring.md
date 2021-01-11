# Monitoring

## Separate successful invocations from failed ones

There is an _Duration_ metric for _AWS Lambda_. This metric will tell you how long the overall setup and invocation of your lambda was.
There is a problem though. This metric also takes into consideration invocations which resulted in an error. **Having errors included in your _duration_ metric, will skew the metric**. This is because the failed invocations (either by caught or uncaught exception) are much shorter (in terms of duration, latency).

To ensure that your metric is not laying to you, you should create at least 2 _duration-based_ metrics. One metric for successful invocations and one for the failed ones.

You could try using _EMF_ and instrumenting with `performance.now` within the handler of the function. While this might be what you want, you have to take into consideration, that this method will not capture the _init phase_ of your lambda function.
Another solution would be to write a **custom _extension_**. With this you have access to the `init` phase (can be notified of it).

## Metrics math

This feature allows you to create metric out of multiple other metrics. You have _functions_ at your disposal.
You can reduce the costs by combining _AWS native_ metrics into 1 custom metric.

## Composite alarms

This alarms are based on other alarms.
For example you might have 2 alarms on your EC2 instance. One for CPU Utilization and one for Memory Utilization. With composite alarms, you can fire an SNS notification if CPU Utilization is at X and Memory Utilization is at Y.
