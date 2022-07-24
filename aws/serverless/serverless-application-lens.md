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

## Security

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

## Reliability

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

## Performance efficiency

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

##### AWS Step Functions

- When you integrate with Step Functions, you will most likely need to poll for the completion status. **Consider using WebSockets for communicating whether the workflow is done or not**. Using the traditional request/response model might incur additional costs.

- This chapter is a mix of optimization techniques for asynchronous workflows and is NOT specific to Step Functions. I'm not sure why they put all this information in this chapter.

#### Monitoring

- Use CW logs for Step Functions. Keep in mind the additional charges when doing so.

- Use **AWS Compute Optimizer for AWS Lambda functions**. The service can suggest the best memory settings for your AWS Lambda functions.

#### Tradeoffs

- **Configuring AWS Lambda provisioned concurrency will significantly increase the cost of your workload**.

- CW is not free and, while very useful, can be the most significant contributor to your AWS bill.

## Cost optimization

### Cost-effective resources

- **Serverless architectures** are **usually cheaper** since many services operate in the "free if idle" model (sadly, not the case for all resources).

### Expenditure and usage awareness

- **Use tags to gain visibility into different areas of your architecture**. This is especially true if you operate in an organization with multiple teams running multiple workloads.

  - Remember that, for example, in AWS CDK, it is possible to enforce tagging. You can write aspects that inspect all the resources. If a given resource does not have a tag attached, you fail the deployment.

- Make sure the developers understand and have a view into the costs of running the application they are responsible for. Not doing so is a huge mistake. A clear picture of costs will result in more ideas for bringing the cost down.

### Optimizing over time

#### Lambda cost and performance optimization

- Here, the **amount of the function memory is the name of the game**. Apart from rewriting your logic to another language, the memory setting is THE knob you can turn to tweak performance.

  - Depending on your language of choice, **switching to a different VM architecture might also bring cost-related benefits**.

  - When switching to a different VM architecture, consider testing your function before deploying it to production. AWS recommends using aliases for that.

#### Logging ingestion and storage

- **Log what you need, not what you can**.

  - In the beginning, it is easier said than done, but as your application matures, you should have more confidence in what kind of logs you need.

- When using CW, use EMF, **do not use CW API to write logs**.

  - Make sure you **set dimensions on parameters that do not change that often**. The more dimensions your metric has, the more it will cost you to maintain that metric.

  - For data that changes often, **use the `properties` object**. You will be able to filter your logs based on the data in CW logs, but you will not be creating unnecessary dimensions.

#### Leverage VPC endpoints

- The **NAT Gateway pricing is very high**. Keep that in mind whenever you even think about using VPC.

- VPC endpoints/interface endpoint will allow you to skip the NAT Gateway entirely.

  - It will also make the request faster since you make fewer "network hops".

#### DynamoDB on-demand and provisioned capacity

- The guidance here is clear: **use provisioned capacity when you have a predictable load**, **use on-demand capacity when you have a spiky load**.

  - Do **not forget about the autoscaling abilities of DynamoDB**. Even if your workload is semi-spiky, you can still take advantage of the lower cost by running provisioned capacity and leveraging the autoscaling.

#### AWS Step Functions Express Workflows

- **Prefer callbacks or `.sync` tasks over polling resources**. You are billed per state transition. The more time you have to poll a given resource to get the job status, the more you will pay.

- The **standard flavor of SFN has an exactly-once execution model**. Take advantage of this characteristic!

  - That is not the case for the express version of the SFN workflow.

- **You can use express workflow and then implement a given workflow step using another state machine that uses the standard workflows**. Mixing and matching the execution models is a fascinating concept.

#### Direct integrations

- In **some cases**, it might be beneficial to skip the "middleman" AWS Lambda function and use the **direct service integrations** instead.

  - **Keep in mind** that, **skipping the additional function shifts the complexity elsewhere**. In the case of API Gateway, you now have to deal with mapping templates and more complex configurations.

- In **other cases, it does not make any sense to skip the AWS Lambda function**. If you perform some complex logic/calculations in the compute layer, it would be unwise to move all that logic into VTL – it would be much harder to maintain.

#### Code optimization

- The first example of architecture is fascinating. **Instead of downloading the whole S3 object, consider using Athena to perform transformations/introspection of a given S3 object**. You will have to wait for Athena to finish, but the memory requirements of our function are much less than if you were to download the whole object.

- Another feature AWS presents is the **S3 Select**, which we can use **for introspection** of a given S3 object. Note that **S3 Select only supports the `WHERE` SQL clause and limits the scope to a single S3 object**. Pick your tools wisely here.
