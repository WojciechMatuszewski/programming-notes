# Notes from "Show Me SaaS" show

## Glossary

- Shuffle sharding: relates to the _noisy neighbor_ problem and the _shared resource_ model in SaaS applications.
  This process exists to prevent resource contention in architectures that might experience a _noisy neighbor_ problem.

  > You can learn more about shuffle sharding [in this video](https://www.twitch.tv/videos/1578416865?filter=archives&sort=time).

  - Simple sharding: randomly distribute tenants across a couple of resources. Assign a single tenant to a single resource.

  - Shuffle sharding: randomly assign a single tenant to multiple resources. This is where the term "shuffle" comes from – think randomly shuffling a deck of cards and picking four out of 52 resources. The resources which you assign to a tenant A might overlap in some percentage with the resources assigned to the tenant B.

    Shuffle sharding enables you to dynamically re-balance the assignment to resources. Imagine that the resource RA is overloaded. We can ensure that other tenants use other resources RB, RC... assigned to them.

## Episode 2 – Inside SaaS Identity

- Knowing who you are as a user, connecting that to the tenancy model, and flowing that information through the application.

- JWT token is essential in the identity story. It allows you to inject a lot of context into the request.

  - You should **not** be putting a lot of info into the JWT token. Use `claims` for attributes, do **not** handle authorization via JWT.

  - Steer clear from adding _application_ data to the JWT.

- You might want to cache context extracted from JWT.

- Multi-region identity is complicated.

  - Active-active disaster recovery identity sounds lovely, but it is tough to get right.

  - Sometimes, your identity has to have regional awareness. Think GDPR.

- Each tenant uses a single Amazon Cognito User Pool per tenant or single, shared User Pool ("pool" model).

  - The information about Amazon Cognito User Pools are stored in the DynamoDB table for fast lookups.

- It would be best if you were not putting off thinking about multi-tenancy in terms of identity in your application.

## Episode 3 - Onboarding Automation: Fact or Friction

- Onboarding can be complex. It is not only probing for contact information. It is also about wiring the tenant to the existing architecture.

  - Depending on your architecture (pooled vs. silo), the onboarding can be complex technically.

- The presented example – VPC per tenant – is quite fascinating.

  - The application lives in each VPC. That means multiple applications are running, completely separated for different tenants.

- You have to keep resource limits in mind. Some resources have hard limits, and you cannot create more of them.

- It is nice to see the CloudFormation -> SNS pattern to wait for the deployment to finish.

  - It is a pity that CloudFormation does not support EventBridge thought.

## Episode 4 - Designing Multi-Tenant Microservices

- You might find your microservices split works differently for a single-tenant vs. multi-tenant environment.

- Noisy neighbor problem is significant to handle in a multi-tenant environment.

  - Distribution of load and profiling comes into play here.

- It is important **not to leak** multi-tenancy into many layers of your code. Your code has to understand the relation between the request and the tenant. How contained can you make it?

- Create small, reusable functions that deal with the tenancy. Use them inside your "main" code.

- **Do NOT** create a microservice especially for handling the concept of a tenancy. It is hazardous to do so as that microservice becomes a singular point of failure. It is better to duplicate code than have such a microservice.

## Episode 5 - Life in the SaaS Lane: Routing Multi-Tenant Traffic

- Two ways to route tenants. The **Domain-driven routing** and the **Data-driven routing**.

  - **Domain-driven** is the model Vercel promotes.

  - **Data-driven** is the model where the routing information is present within the JWT token.

- Your application might need to deploy a _routing infrastructure_ to handle the routing concerns.

- These solutions relate to the **Data-driven routing**.

- One also must consider how the support staff can get into a tenant environment for debugging / support cases.

- Throttling is often part of the routing story.

## Episode 6 - Tenant Isolation Strategies: Can't Touch This

- What do we mean by saying _isolation_?

  - Isolation means enforcing the rule that you cannot access other party resources.

  - Authorization is smaller than isolation.

- Tenant isolation models are often hard to change, whereas the authorization characteristics of the application can change fluently.

- A couple of levels of isolation.

  - **Full-stack isolation**. Very complex from an architectural standpoint.

  - **Resource level isolation** (think only a database isolated to a given tenant). It is a bit easier to implement than the full-stack isolation, but it still might be complex.

  - **Item level isolation**. Think DynamoDB leading keys policies.

- You have to build different layers of security around the isolation story. You cannot rely on a single database statement with a `WHERE` clause.

- Isolation can influence your technology choices.

- Not every AWS service does not have an IAM condition suitable for tenant isolation.

- In AWS land, one might use **dynamic policy generation** scoped to a given tenant. Use either APIGW authorizer (maybe even with an API key) or in-lambda assume role calls.

  - You might need to deploy the compute with a broader scope that allows scoping down to a specific tenant.

## Episode 7 – SaaS metrics: Shaping your SaaS architecture with data

- We should inject the aspect of tenancy in all log statements.

- Each application should be tracking product metrics and the developer metrics. Combining those two will give you a holistic view of your application.

- Metrics on a micro-service level are pretty important. It is vital to understand how tenants use a given micro-service.

  - Customers communicate by feature usage. If a given feature is not used, it is a strong signal that you should focus on something else.

  - > How many features are actually in use?

    That is impossible to know if you are not instrumenting (at a particular scale).

- When it comes to observability and instrumentation, tooling is critical. You have to have the tooling to drill into details for a specific tenant.

- Gathering metrics in one place is important from an operational standpoint. This is where the notion of **single pane of glass** comes from – having all your easily accessible in one place.

  - One of the **key metrics** is the **cost per tenant**. You need to make sure that you are making some money.

- Metrics are the lifeblood of big SaaS companies. When you are starting, you have to push to make metrics a high priority of your business. If you ignore this aspect, you might have problems growing your application later.
