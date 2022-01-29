# Mastering AWS Kinesis Data Streams

Session from O'REILLY live training.

- Instructor: Anahit Pogosova.
- Link: https://learning.oreilly.com/live-events/mastering-aws-kinesis-data-streams/0636920059729/0636920059728/

## Introduction

- Fully managed service to stream data.

- Data is stored **from 24 hours to up to 365 days**.

- Data available in **near real-time**.

- Consists of shards. Each **shard is an ordered queue**.

  - _Resharding_ might be a good strategy to react to traffic.

- Kinesis uses **_Partition key_ to decide which shard the data goes to**.

  - It might be good idea to distribute the records across multiple shards.

- There are **two capacity modes**. You can switch between them.

  - The **on-demand mode** is **not autoscaling!**

  - The **on-demand mode** is good for exploration and playing around.

- Many services integrate directly with Kinesis.

  - Remember that **KCL uses DynamoDB table under the hood to keep track of where you are in the stream**.

- A lof of **AWS services use Kinesis under the hood**. To name a few:

  - _CloudWatch Logs_
  - _CloudWatch Events_
  - _IoT Events_
  - _Kinesis Firehose_

- One neat use-case is the _DynamoDB CDC_ to _AWS Kinesis_.

## Writing data with AWS SDK

- There are two API operations one can use to write to the stream.

  - Using the `putRecord` API call
  - Using the `putRecords` API call

- Using **`putRecords` API call is much more efficient**.

- **No matter the API you choose**, the **records are counted separately towards the shard throughput limit!**.

  - The `putRecords` optimization is on the HTTP layer, and will do nothing if you experience throttling.

- Consider using some kind of HTTP mocking library to see what kind of HTTP calls SDK is doing.
  - Anahit is using `nock` in the course.

### Failures

- It is always a good idea to implement retries on the client side.

  - The **default SDK settings have very long timeouts**. Consider changing the options to more reasonable ones.

  - It is important to **understand the difference between _connectTimeout_ and the _timeout_ settings**.

    - The **_connectTimeout_**: timeout for establishing a **new connection** on a socket.

    - The **_timeout_**: read timeout for **existing** socket.

- To **test retry behavior use HTTP mocking library, not mocks**.

  - Mocks are hard to deal with, especially with _jest_.

- The `putRecords` API can fail in two ways

  - One of the way is for the API to throw an error. This would happen if the SDK could not establish a connection with the Kinesis API
    or the API call itself would fail.

  - The second is **when a throttling error occurred on Kinesis side**. This time the SDK **will not** throw an error.
    Instead, the response will contain the **`FailedRecordCount`** property. The same response scheme is used in other batch API calls for other AWS services.

  - Since the second case is not throwing an error, **make sure you handle it**. Otherwise you might loose your data!

  - The **array of records returned by the `putRecords` API call** has **the same order as the array of your events**. Use this fact to **deduce which events failed and which succeeded**. (TBH it is a bit weird. Could not they provide us with an ID of some sorts?).

### Duplicates

- Duplicates do happen. You **should build your customers in a manner that can handle duplicate records**.
