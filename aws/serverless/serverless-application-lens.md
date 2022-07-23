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
