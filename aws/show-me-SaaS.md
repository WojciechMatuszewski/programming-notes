# Notes from "Show Me SaaS" show

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
