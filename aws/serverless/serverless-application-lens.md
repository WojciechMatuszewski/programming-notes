# Serverless Application Lens

Taking notes while reading [this document](https://docs.aws.amazon.com/wellarchitected/latest/serverless-applications-lens/the-pillars-of-the-well-architected-framework.html).

## Operational excellence

### Operate

#### Metrics and alerts

- Understand the critical metrics for the services you use.

- Use either the powertools or EMF to emit custom metrics.

- Create dashboards to gain a holistic view of different areas of your application/business.

- Different services have "built-in" metrics that you can monitor. They are pretty helpful in understanding how healthy your system is.

#### Centralized and structured logging

- Use a single format for logs and stick with it.

  - You most likely want to have some logging abstraction that does the "normalization" of the output for you.

- **Use correlation IDs** and propagate them throughout the system to have the whole picture.

- Consider **using extensions and tapping into the AWS Lambda Logs API** to move AWS Lambda function logs to other destinations. Remember that **if done through CW, you will pay CW storage, ingestion, and data transfer costs!**.

#### Distributed tracing

- AWS recommends using X-Ray, but I would argue that there might be better services like Lumigo.

  - If you decide to use X-Ray, consider using the _"Service map"_ view.

- You might find the **subsegments and annotations** useful. Utilize them to gain additional insight into how your workload is performing.

  - **Annotations are handy since X-Ray will index them for search purposes**. You can add up to 50 annotations per trace.

#### Prototyping

- Use IaC to create many different environments.

  - Use different accounts on the per-team or per-developer level.

#### Configuration

- For **non-secret values that might change at deployment time**, use **environment variables**. A good example would be the URL of the API Gateway or the name of a DynamoDB table.

- For **values that might change dynamically**, use **SSM Parameter Store**. Remember about caching those calls.

  - Consider using the SSM Parameter Store extension.

- I'm glad that AWS is not pushing the Secrets Manager down our throats here, as the service can get quite expensive, and the SSM Parameter Store ("basic" throughput) is free.

#### Testing

- Use the actual services instead of mocking.

#### Deploying

- Use a framework. It does not matter which one. What matters is that you do not write plain CloudFormation.

- Create separate stages to reduce the blast radius of your change. Run tests when a given stage deploys.

- AWS recommends AWS CodePipeline and AWS CodeBuild. I'm not surprised.

  - In my opinion, AWS CodePipeline requires a LOT of configuration. Sure, you might use AWS CDK and have that configuration mostly hidden from you, but if shit hits the fan, are you comfortable not knowing how your deployment pipeline is working?

  - On the other hand, AWS CodePipeline makes some of the things pretty "easy", like deploying to multiple environments and integrating with IAM (the recent addition of GitHub OIDC makes integrating IAM with GH Actions pretty easy, though).

- Before making a decision, I suggest you [read this article](https://serverlessfirst.com/switch-codepipeline-to-github-actions/).

- Favor **safe deployments over all-at-once systems**. That means introducing your changes gradually.

  - In the context of serverless applications, you might want to use different aliases for your functions or canary deployments with API Gateway.

- Favor **safe deployments over all-at-once systems**. That means introducing your changes gradually.

  - In the context of serverless applications, you might want to use different aliases for your functions or canary deployments with API Gateway.

### Security

#### Identity and access management

- **If possible, use temporary AWS credentials**.

  - There are many ways to get those. One might use AWS Cognito or assume an AWS IAM role.

- In recent years, **AWS has been rolling out the support for tag-based conditions in IAM policies for various services**. These are great for applying taxonomies to different resources.

- Consider using the **resource policies over embedding the inline policies on a given identity**. Doing so keeps the permissions "close" to the resource they relate to.

  - Keep in mind that **for cross-account access, both the resource and the identity policy have to allow for an action**.

  - For access in the context of the same account, either the resource or the identity policy has to allow for action.

- Consult [this documentation page](https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies_evaluation-logic.html) when in doubt about how the evaluation of permissions works.

- **AWS Lambda Authorizers** are very powerful and allow you to integrate with any identity provider.

#### Detective controls

- In most cases, it's either the vulnerable dependencies or pushing secrets to a given repository that bring you trouble. Make sure you have some way of handling these problems.

  - Consider using services like Snyk or others that scan your code for secrets and other security-related issues.

#### Infrastructure protection

- Consider network boundaries, especially when interacting with services deployed to a VPC.

  - Do not be that guy who leaves his security group wide open.

#### Data protection

- Do **not log sensitive data**. Keep in mind that **API Gateway Access Logs and Execution Logs might contain sensitive data**.

  - Ensure that you have the settings tuned for those.

- AWS is bullish on encryption in this section. You should be good if the costs of KMS are baked into the service you use. Remember that data encryption might cost a LOT depending on your traffic.

  - Encryption at rest is the name of the game.

### Reliability

#### Foundations

- Throttling plays a vital role in ensuring your service is reliable.

  - You can implement throttling via the **reserved concurrency** settings of **API keys and usage plans** (remember that the API keys should not be used as an authorization mechanism).

  - Throttling can also be helpful in an event where your architecture creates an infinite loop of requests.

- In the context of AWS Lambda, **keep in mind that the maximum AWS Lambda function concurrency is shared between all the functions in your account in a given region**.

- If you find yourself in a situation where throttling is an issue, **consider moving your workload to an asynchronous model**. In such a model, you have much more control over the concurrency (for example, via the Amazon Kinesis streams by increasing the shard count and changing the parallelization factor).

#### Failure management

- The problem with synchronous workflows is that you do not have many options on how to handle the failure. You have to retry and, if the issue persists, yield to the user with an error message.

  - With asynchronous or stream-based workflows, you can deploy DLQs or use _destinations_. There, the failed piece of work (the payload) could be inspected and corrected.

  - **Do not forget to set maximum retry attempts or similar settings** when integrating with other services to implement error handling. The last thing you want is for your event to be gone.

- One of the most robust patterns for error-handling is **the saga pattern**. The pattern is based on having the orchestrator (in this case, a state machine) handle the partial failures and rollback the committed work.

#### Limits

- Limits are there for a reason. Understand them, and have them in the back of your mind when you architect.

  - If you have problems with limits, consider moving to asynchronous flows where you have control over how much throughput the system can process.

### Performance efficiency

#### Selection

##### APIGW

- Know the **difference between the regional and edge optimized** APIs.

  - Use **regional if you want to use your own Cloudfront distribution**.

  - The document says that the regional APIGW enables HTTP2 by default. Intriguingly, this is not the default setting for Cloudfront distributions.

##### AWS Lambda

- TIL that **AWS Lambda integrates with Application Auto Scaling**. You can use Application Auto Scaling to manage provisioned concurrency.

- This section contains a convenient decision tree that should help you do decide whether your AWS Lambda function should be inside a VPC or not (hint: it most likely should not).

##### AWS Step Functions

- Know the difference between the express and the standard step machine workflows.

  - The most surprising fact from this section is that **the express workflows have UNLIMITED state transitions**. I rarely see a quota of "unlimited" – a very bold statement from AWS.

- Keep in mind that, for both the standard and the express workflows, you can enable publishing logs to AWS CloudWatch. Of course, **this will increase the cost of operating the service**.

#### Optimize

##### APIGW

- Use **data compression along with correct content-encoding headers**. I think this feature of APIGW is often overlooked.

##### AWS Lambda

- Since you will most likely integrate with AWS Lambda, it is **vital that you tune the timeout settings of the function correctly**.

  - It cannot be too long as that might lead to unnecessary compute time during errors (if the timeout were lower, the function execution would have been killed sooner).

  - It cannot be too short as that would create a risk where some of the work is dropped.

  - If relevant, AWS documents those gotchas in a given "integration guide" (for example, the SQS one).

- Remember about the "basic" stuff like per-container caching (initialization) and the fact that **AWS Lambda extensions might slow your function down as the resources the function is allocated with are shared between the function and the extension**.

- Consider using the **RDS Proxy** for connection polling.
