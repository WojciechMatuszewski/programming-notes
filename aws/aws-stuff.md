# AWS Stuff

- **Principal** is a person or application that can make authenticated or anonymous request to perform an action on a system. **Often seen in-code in lambda authorizers**

* Security **in the cloud** is **your job**

- Security **of the cloud** is the **AWS job**

- **AZ (Availability Zone)** is a distinct location within an AWS Region. Each
  Region comprises at least two AZs

* **AWS Region** is a geographical area divided into AZs. Each region counts as
  **separate** geographical area.

- **Virtual Private Cloud (VPC)** is a virtual network dedicated to a single AWS
  account. It's logically isolated from other virtual networks in the AWS cloud

* **EFS and S3** are popular storage options

- **Cloudfront** is a CDN

### Throughput vs IOPS

- **throughput** is a measurement of a **number of bits written / read per second**

* **IOPS** is a measurement of a **number of read / write operations per second**

### Lambda

- **LAMBDA IS HA by default! MULTI-AZ!**

* **scales automatically** (can run functions concurrently)

- **YOU CAN ATTACH SECURITY GROUP TO A LAMBDA**

* lambdas **inbound connections are blocked**. When it comes to outbound, **outbound TCP/IP and UDP/IP sockets are supported**.

* by default **lambda is within so called "No VPC" mode**. This means that (actually quite logically) it **will not have access to resources running within private vpc**.

#### Billing

- executions are **billed by the millisecond**

#### Alias

- alias can point to a **single version** or **2 lambda versions**.

* you would use **aliases for shifting traffic between versions**.

#### Execution Role

- this the **role assumed by lambda when invoked**

#### Resource Role

- **lambda** is **treated as a resource** so you can grant **resource-based policy to it**.

* this is **especially useful** when **wanting to grant other account permissions to invoke your function**

- you can allow your function to be invoked by a given lambda service which originates from given account

#### Firecracker

- _Firecracker_ processes your lambdas I/O and network requests and sends them to the _Host kernel_

* lives in a very limited guest envioriment

#### Provisioned Concurrency

- you have to have a **function version** or an **alias that DOES NOT point to $latest** to turn it on.

* it will basically **keep concurrently X clones of execution environments ready** thus making cold starts obsolete (unless you make more requests than you have provisioned for).

- can be **autoscalled by using Application AutoScaling**. Do not mistake this with auto scaling groups.

* Do not confuse this with **reserved capacity**. Here we are **provisioning workers**, in the other case we are **reserving the amount of time a lambda can be invoked**

#### Inside VPC

- **YOU CAN PLACE LAMBDA INSIDE A VPC SUBNET**. The **subnet HAS TO BE PRIVATE though**. If you want to make sure that your lambda within VPC can access internet **place NAT gateway within your public subnet**. This is because **lambdas have ENI assigned that never gets public IP**.

* since Lambdas needs ENI, you might hit **throttling when there are not enough ENIs for given lambda container**. The limit is **350 ENIs per region**. When that throttling happens you will se an `ec2` related throttling message.

#### Monitoring

- **CloudWatch metrics** are **somewhat limited**. They only include **info about errors, invocations, duration etc**.

* for **memory usage information** you can either look into **CloudWatch logs** or you can **create metrics using metric filter (not provided by default)**. The metric filter is very useful if you want to have the graph that includes **allocated memory vs used memory**.

#### X-Ray

- you can **use XRay for distributed tracing**.

* remember that **X-Ray collects data of INCOMING requests. Not only those made within / outside our service**

- **X-Ray works in real-time**. Can be used for real-time monitoring.

- X-Ray **does not sample every request**. This is quite important since that means that the data completeness is not guaranteed and it **should not be used as a audit mechanism**.

##### X-Ray daemon

- listens to traffic on **port 2000 UDP**. This is where the SDKs are sending the data.

* automatically **installed on lambda**.

- on `ECS` you have to **setup an container which will run your daemon**.

#### Lambda @Edge

- **CloudFront will invoke your function when given event happens**. These are **events that has to do with request life-cycle @ CloudFront like origin request or smth like that**.

* you can use it to **perform A/B testing, since you can mutate the request send by the user**

- you can also **perform some redirect logic , maybe check auth**.

* **they are not executed in the edge location**. They are executed in the region closes to the edge location. This is why you have to be mindful of the region whenever you browse the logs.

- these are **NOT running at the edge locations, they are run within the regional CloudFront caches**.

#### CloudFront functions

- these **actually run within the edge location**

* **very constrained in terms of what they can do**. Full comparison table [available here](https://aws.amazon.com/blogs/aws/introducing-cloudfront-functions-run-your-code-at-the-edge-with-low-latency-at-any-scale/?utm_source=newsletter&utm_medium=email&utm_content=offbynone&utm_campaign=Off-by-none%3A%20Issue%20%23140)

#### Event Source Mapping

- this is where **lambda service reads from other service and invokes your function**

* this is **one way a service could _trigger_ your function**

- services that use _event source mapping_: _DynamoDB_, _Kinesis_, _MQ_, _SQS_ and _managed Kafka_.

* _Event Source Mapping_ uses **lambda execution role** for IAM permissions. This is why you have to specify that, for example, your function can read and delete messages from SQS

#### Lambda destinations

- **Lambda destinations are used when lambda is invoked by other services** like: **s3, SNS, SES, Config etc..** and then those **onSuccess or onFailure** events are **send to Lambda, SNS, SQS, EventBridge**.

* work only for **async** and **stream based invocations**. This is completely BS

#### Resource-based policies

- when a service does not use the _event source mapping_ and underlying _execution role_, you have to specify the _resource-based policy_ so that the service can invoke your function directly.

* list of services that use this method of integration is much greater than the _event source mapping_ ones.

- and example policy to enable a given function to be invoked by another function
  ```ts
  beingInvoked.addPermission("invoke-resource-based", {
    // default execution role of a lambda function
    principal: new iam.ArnPrincipal(invoker.role?.roleArn!),
    action: "lambda:InvokeFunction",
  });
  ```

#### Infinite loops and lambda

- with event based architecture, it may happen that you cause an infinite loop of invocations (think s3 events)

* you can troubleshoot the problem by **setting reserved concurrency limit to 0**.
  This will make it so your function is not invoked at all. There is a **built-in option for that within the UI**.

#### Deployments with `CodeDeploy`

- this is a neat solution where you want to perform `blue-green` on different lambda versions.

* you will need to use an **alias**. The traffic **will be shifted by `CodeDeploy`**

- just like a normal `CodeDeploy` deployment, there are **validation hooks** you can use to decide if your new alias is OK.

#### EFS Integration

- your lambda functions can integrate with EFS

* there are some latency considerations. Reading large files will take a couple of seconds (might impact start performance)

- the benefit here is the _elasticity_ of the storage. Keep in mind that EFS can grow almost infinitely.

#### Container deployments

- you can package your lambda code as a container image and push that to ECR

* the size limit here is 10 GB, much bigger than the previous 250 MB of zip file

- **your image has to communicate with lambda runtime**

* this is a good way of deploying code which is subject to heavy regulatory constrains or you need more than 250 MB of space

##### Storage copy issue

- the bigger the image, the longer the _optimization_ phase

* the things that you put into your container are split into 100 MB chunks, this is an side-effect of how Lambda service optimizes your container

- when you are loading large datasets, it takes time, those chunks have to be downloaded. Lambda might timeout before the dataset is loaded

##### Lambda layers and extensions with container deployments

- lambda layers are not directly supported with the container deployments

* to make layers work, you would have to **bake the layer directly into your container** or **use a multi-stage build with a image that contains the layer**

#### Tumbling windows

- available to you when you are consuming **DDB streams or Firehose events**

* enables you to return **state** which **could be consumed by another lambda invocation**

- **state persists within the tumbling window**. The **max window you can set is 15 minutes**

* for doing **agregation work**. Think of **calculating sales numbers for a 15 minute window**

- aggregation is preserved **per shard**

##### Vs Kinesis Analytics

- with **Analytics streams** the **aggregation** is performed **on the whole stream**

* with **Tumbling windows** the **aggregation** is performed **only within a single shard**

#### Custom checkpoints for DDB And Kinesis

- this allows you to **return which message failed in a given batch**

* is **different than bisect on error** because **with bisect on error one message might be processed multiple times**. Here we are **telling the pooler which message failed exactly**

- **DOES NOT WORK WITH SQS**. You still have to use the pattern where you throw the messages that failed the processing

### Step Functions

- **state machines as a service**

* for **orchestrating serverless workflows**

- you can create **manual steps** with step functions. This can be **whi le dealing with hybrid approach** where **there is some level of human interaction needed (like clicking a link)**

* **does not** integrate with **Mechanical Turk**. You will need **SWF for that**.

#### Best Practices

- use timeouts to avoid stuck executions

* avoid passing large payloads (use arns)

- handle exceptions with the `Catch` operator

* there is a **limit in terms of number of entries in the execution history**. It's quite big - 25000. Avoid reaching it (split executions into multiple workflows if needed)

### IAM

- **IAM** is universal, **is global**, does not apply to regions

* **Identity Federation** allows your users to login into AWS Console using 3rd
  party providers like Google etc..

- Supports **PCI DSS** framework. This is some kind of standard for security

- **describe** means **to be able to view, inside the console**.

* **Policies = Permissions**, written in JSON

- **Roles** enable one AWS service do something / interact with another. For
  example virtual machine (EC2) interacting with AWS storage.

* **Root Account**: email address you first sign up to AWS with. This account
  basically has a godmode and can do everything in the console. That's why you
  pretty much never want someone to login on root account, just like in Linux.

- Users can have **programmatic access** to AWS console. This basically allows
  you to pass access key and secret key so that you can interact with developer
  tools.

* Users **can be added to groups**. These groups **can have policies** assigned
  to them.

- Policies have **different types**. Like `Job function` or `AWS managed`.

* **explicit deny ALWAYS overrules ALLOW**. This is very important, especially when working with groups.

#### Real Identities

- both **user and roles** are known as real identities. They both have **ARN** and **can be referenced in other areas of AWS**

* **groups ARE NOT REAL IDENTITIES**. You cannot login onto the group and such. This is such an organizational being.

#### Roles

- **not something you log into**. It does not have username, password or any kind of long-term credentials.

* you should **prefer attaching roles** instead of using aws credentials

- roles can be attached to many services

* roles can be used **in any region**, they are universal

- **roles have underlying policies, which have the notion of the Effect (allow/deny)**

* **policies** are **associated with roles**

- allows you to set **Authentication attributes**: **Usernames, passwords, Access Keys, MFA and Password Policies**

* **authentication** is the process there **you are being verified for being you, being that entity you present yourself as**

##### Service-linked Role

- this is the **role assumed by a service**

* it **allows** given **service to do stuff on your behalf**.

- think of **assigning role to CF, to deploy stuff in sls framework**.

* this role **cannot be restricted by SCP**.

- this role makes your life easier since you do not have to setup permissions for services manually.
  These roles are created when you agree for a given service to create a role within UI, you've done it a lot of times.

#### Assuming Roles

- role which you can assume has two segments
  - **trust policy**. This defines **who can assume a role**
  - **policy document**. This is a standard policy

* **trust relationship** is **ONLY checked** when **assuming a role** (that usually happens once or infrequently)

* assuming a role means **being a completely different identity, defined by assumed role**

- under the hood **assuming a role means using completely new, temporary, credentials created with sts which are associated with the assumed role**

* **all roles that could be assumed are automatically assumed**

- **by default** newly created User has **implicit DENY** on all services. **You should assign roles to lift the implicit deny**

* **assumed credentials** are **valid until the expiration date expires**

- when you are creating **cross account assume role stuff** you can **further protect the policy by adding externalId condition**. This **externalId can only be used when assuming via CLI / SDK**.

##### Revoking Sessions

Since removing a policy from a role which is assumed can be destructive there is another way of removing assumed (short-term) credentials. That is **revoking a session**. This basically **ads a deny all policy for tokens granted before given date**. So you did not remove the permissions directly, only invalidated given tokens. Pretty sweat!

##### Assuming cross account role

##### Assuming EC2 Role

You have to assign a role to a given instance. Then you can ssh into that instance and do stuff that that role allows you to do. This is possible because the `aws-cli` is able to obtain the short term credentials that the role provides. This is done by **querying an instance metadata role**.
`http://169.254.169.254/latest/urity-credentials/NAME_OF_THE_ROLE`

##### Assuming EC2 Role on premises

First thing first **you will need a role that will be shared along all your on prem instances**. Assuming you've done that, you should be using `AWS CodeDeploy` along with `aws-codedeploy-session-helper` to safely fetch credentials to your machines. This is mainly because `CodeDeploy` allows you to tag your instances, track deployments and does a lot of things for you. This approach **involves a cron job**.

##### Switching Roles

- users can switch role **using SwitchRole console menu** or using an **link within an email send by administrator**.

* you **cannot switch roles** when you login **as root user**.

##### `iam:PassRole`

- this is a guardrail on the permissions to grant to other AWS services

* imagine having only some kind of limited role when you can create EC2 Instances. Without `iam:PassRole` you could assign an Administrator Access to the EC2, connect to it and, potentially, do damage.

- you can limit what kind of roles can can be _passed_ or limit things to specific service.

#### Policies

- **applied** to a **IAM Role or a Resource**

* have a **Sid**. This is **basically just a description** of the policy.

- **explicit deny** always **overrides explicit allow**

* **inline policies** should only be applied to a single user. They **should not** be used to **applying policies to multiple identities**. You should **use managed / custom policies for that!**.

- always **prefer explicit deny no matter what**. Use **conditions if necessary**.

* you can **leverage tags** to create access control lists based on tags.

##### Resource arn syntax

There is a standard `arn` indentifier which is assigned to every resource

```
arn:partition:service:region:account:resource
```

You might be wondering why `s3` have `arn` defined as:

```
arn:aws:s3:::example_bucket
```

This is very different from other services where you have to specify most (if not all parts of the syntax).

It turns out that this is somewhat **service specyfic**. For global services you can skip the `region` part, and it just happens that for `s3` you can skip both `region` and `account`.

##### Resource Policy

Now we are in the domain of a given service / resource.

- **IAM is not involved here**, control is **handled by a given service / resource**

###### Cross account

- you can setup **cross account resource policies**.

* this is useful when you **want user to not give up his trusted account roles**.

##### Conditions

- you can create elaborate conditional logic for a given policy

An example for s3-prefix (folder)

```json
{
  "Condition": { "StringLike": { "s3:prefix": ["testuser/*"] } }
}
```

- you can use **variables** within IAM policies

```json
{
  "Condition": { "StringLike": { "s3:prefix": ["${aws:username}/*"] } }
}
```

#### Managed Policies

- **managed policies can be applied to multiple identities at once**. There are 2 versions of managed policy: **customer managed policy** and **aws managed policy**.

* there are 2 notable managed policies: **AdministratorAccess** and **PowerUserAccess (this is for developers)**

##### SERVICE_PowerUser managed role

- usually these roles allow to read & write and delete operations

* think pushing, pulling, creating with `CodeCommit`.

- **this very different than PowerUserAccess managed role**. This role allows for almost everything within every service.

#### Inline vs Managed Policies

- **inline policies** can only be **used by one entity at the time**. You **cannot reuse the same inline policy for multiple identities (can be groups), you would have to create a new one, even if it's the same**.

#### Resource Based Policies

- these are **special subset of policies** which are **attached to AWS Services**.
  Think allowing _APIGW_ to invoke your lambda function. You would **add that policy on the target resource, not the source of the event**

* they have **Principal field**. This is due to the fact that they are evaluated whenever some principal access given resource. IAM role / group policies does not have that because they are applied to principals from the beginning.

- what is interesting is that you can **grant access to every identity within an account using :root suffix**.

  ```json
  {
    "Principal": { "AWS": "arn:aws:iam::1234:root" }
  }
  ```

  or

  ```json
  {
    "Principal": { "AWS": "1234" }
  }
  ```

  This will make it so that every **user and role within 1234 account** have an access.

#### Permissions boundary

- **attaches to the entity**

* additional safety net, it allows you to further restrict what a given role can do

- might be useful while trying to prevent developers to create wide open permissions for resources like lambdas

* it is used for **scoping permissions down, it does NOT provide permissions on its own**

#### Groups

- **CANNOT BE NESTED**. Though the **nesting is not necessary a good idea**. The **explicit deny** can sometimes **override explicit allow**. So in general the nesting should be avoided.

* they are **not a true identities**. A **virtual construct** to **help** you **manage your users**.

- you **can attach policies to them**

* you **CANNOT attach RESOURCE policies to them since they are not a true identity**.

- they are **not allowed to assume IAM Roles**

#### Users

- **UP TO 5000 IAM Users**

* remember that **IAM users are global!**

#### Tags

- you can use `aws:PrincipalTag` to lookup tags **attached to an user or a role**.

#### Certs

- this is a **legacy way of keeping certs**, before ACM was introduced.

* sadly **you cannot migrate IAM certs to ACM directly**. You would have to **download them** and then **import to ACM**.

### AWS Organizations

- **all accounts under organization** can **consolidate their bills** so that you have one **master bill, sum of all small bills**

* there can only be **one master account within organization**. The master account **CANNOT be restricted**

- there can **only be one Root container**.

* the **root container** can only be **controlled by master accounts and Service Control Policies**. If such controls are in place ,they apply to all OUs under the root and all member accounts under given OUs. **Root container / node** is the **account that has OUs underneath**.

- you **can attach SCPS to master account** but **there will be no effect on master account**. As a good practice your master account should not hold any kind of resources.

#### Discounts

- when using **consolidated billing** you can get **volume discounts**. For some **services like S3, EC2**, the more you use them (the volume of data you hold), the less you pay. This is ideal scenario for consolidated billing since all the usage adds up from your other accounts.

#### Sharing

- the **master account** can **turn off RI sharing** for **particular account or whole the whole organization**. There is no private mode or smth.

* for **EC2** the **type and the AZ must be the same**

- for **DB instances all of the DB attributes must be the same (also the AZ)**: engine, instance class, deployment type, license model.

* RI sharing **has to be enabled** for **accounts that purchase the instance, as well the accounts that need the benefit**.

#### Invites

- can be send **through the console or a CLI**.

* you need to have **email or accountID**

- only **one account can join one organization**

* **invitation expire** after **15 days**

- there is a **limit** on how many **invites you can send per day (20)**

* invitations can be **created by any account** as long as **the account has correct IAM permissions**

- you **cannot resend the invitation**. You have to **cancel the previous one, send the new one**

##### Accessing member accounts by using automatically created role

- when you create an account within your organization, that account have `OrganizationAccountAccessRole` created automatically

* that `OrganizationAccountAccessRole` role allows master accounts to assume it and access the member account

##### Accessing member accounts who you invited to organization

- in contrast of creating accounts manually, whenever you invite an account, **`OrganizationAccountAccessRole` is NOT created automatically**.

* what you should do is to **create a role within a member account with a trust relationship to a master account**.

- this can be done **using stacksets** to deploy the IAM config to multiple accounts, or just manually.

#### SCPS

- **by default** the **root account has FullAWSAccess SCP attached**

* they **do not really provide access**. They **restrict what you can do when you HAVE permissions**. Think of them as **boundaries**.

- can be used as **blacklist** or a **whitelist**. The **syntax** is basically the same as IAM policies.

* the **policy evaluation process** is as follows: first you take your iam permissions then you take SCPs. You **take union of permissions from your IAM and SCP, these are the permissions you effectively have**.

- you can have **multiple SCPs which apply**. In such case you take **union of every SCPs with your IAM permissions**.

* remember that **SCP only apply to the accounts that are within given OU**. SCPs **do not apply to outside users (other accounts)**

- the SCPS can be **applied at Root, OU or Account level**. This is very important.

##### Deny List

- by **default** SCPs are configured as **deny list**. This means that **everything is allowed AND IT'S UP TO US TO DENY THINGS WE WANT TO DENY**.

##### Allow List

- reverse of `deny list`. This is where **everything is denied** and **it's up to us to allow things**. This is **usually a bad idea, as things can get messy**.

#### OUs

- there are units called **Organization Units (OU)**. They **can host member accounts or other OUs**. At the top of the **OUs hierarchy there is root container**.

* can **only have one parent**

- an AWS **account can only be a member of one OU**

* you can **move account OUT of OU**. This is useful when wanting to delete an OU.

#### All features

- this is a **free feature**.

* normally, by default only **consolidated billing is `turn on` per se**. There is also **all features mode**.

- the all features mode **can be enabled on a existing organization**.

* this mode **enables you to create Service Control Policies (SCPs)**. These, as described above, allow you to place restrictions on given OUs, member accounts or a root container.

- this is **one way change**. When you enable this **you cannot switch back to only consolidated billing**.

#### Organization Activity

- this is where you can **see which services are used by accounts within Organization**.

* **DO NOT** mistake this for AWS Config. Remember - AWS Config is for looking up different configurations on stuff and checking if they are meeting some requirements

#### Leaving the organization

- **user has to have sufficient permissions to leave the organization**. I think the most **important here are the policies regarding billing**.

#### Moving account between organizations

- if you want to move an account between organizations you will need to
  - have that **account leave the original organization (remember about the ROLES!)**
  - **invite that account** to **new organization**

#### Trusted Access

- you **enable AWS services to perform actions on your behalf within an Organization**.

* this is how you would do **CF stack sets within Organization**.

#### RAM (Resource Access Manager)

- allows you to **share resources within the organization OR WITH OTHER ACCOUNTS**. **Requires** you to **use all-features**

* you can share **a lot of stuff**. Most **notable are subnets and AMIs**. There are some **subnet-related services that CANNOT be placed inside a shared subnet**.

- when you are working with RAM the key keyword is **trusted access**.

* to enable sharing you can either **use AWS Console or AWS CLI (`enable-sharing-with-aws-organization`)**.

#### Cost Explorer

- you can **generate reports**. These reports are a **.csv** file.

* there are **AWS generated and Customer generated tags**. There **tags** can be **used to enhance the rapports (enable filtering)**

- reports **can be generated up to three times a day**

* they are **delivered to s3 bucket**. That **bucket has to be owned by master account**.

- to **have user-defined tags within cost allocation report** you have to make sure to **active given tags before generating a raport**.
  All the tags here are named **cost allocation tags**.

##### Cost & Usage Report

- this is where you can view **the most detailed information when it comes to costs**.

* is **part of cost explorer**

- used for **generating reports**. These **reports have to be generated manually**.

* **reports** are **saved into s3**.

### AWS Budgets

- allows you to **setup alarms** when **costs or usage is exceeded**.

* you can think of this as **billing allarms on steroids**.

- can be very useful when you want to set **alarms for RI utilization and coverage**.

* mainly for setting up alarams and notifications if you go above certain threshold.

- remember that **it takes up to 48hrs for all budget related information to be updated on your account**. That means that **some alerts might fire way after you crossed the threshold of your budget**.

### ACM

- provides **x509 certs (SSL/TSL)**

* only **supported for services that are integrated**.

- you are **not paying for using the certs themselves**

* there are **two types of validation of the domain**. You can either **validate by DNS** or **by email**. With **DNS validation and Route53** you can easily **add DNS record through the console (not automatic)**

- uses **KMS under the hood for keys**

* can **automatically renew certs** BUT only **those which were not imported**.

#### Existing certs

- you can **import existing certs to ACM**. But for gods sake, remember that **imported certs cannot be auto-renewed by ACM**.

#### Private certs

- you can have **private certs for internal APIS**.

#### Multi-region certs

- you **can use** the **same SSL certificate from ACM in more than one region** but it **depends** on whether you are using **ELB or CloudFront**.

* with **ELB** you have to **request a new cert for each region**.

- with **CloudFront** you have to **request cert in n.virginia**. ACM certs within that region can be used for global cloudfront distributions.

#### Zone Apex certs

- you **have to** request a certificate that **has your apex listed** if you want it to **work with `*.example.com` and `example.com`**

- the `*.example.com` certificate **does not cover zone apex domain!**.

### AWS Support Plans

- **developer**: for testing and experimenting with AWS

* **business**: small companies, actually using AWS in production

- **enterprise**: for mission critical and big companies

* **only business and enterprise** come with **full checks for a trusted advisor**.

- **business and enterprise** allow you to have **programmatic(API) access to AWS Support APIS**

* **business** gives you **access to TAM - Technical Account Manager**.

### AWS Config

- so CloudTrail monitors API calls to AWS services (made by your account). AWS Config **monitors your AWS resources configurations**.

- you have too **tick an additional checkbox** to **support monitoring global services (not only in a given region)**

* with AWS Config you can have **history of given resource configurations**

- saves config snapshots to a bucket.

* you can **traverse resource configuration over time**.

* keep in mind that this tool **is NOT used for restricting anything. It merely watches over your resources over-time**

- there is a notion of **relationships**. There are things that are tied **together**. **Think EBS volumes and EC2 instances**.

#### Rules

- you can **create rules**. There are **predefined(AWS) rules** or you can create **custom rules**.
  With this you can for example see if someone enabled inbound port on security group which you deemed non-compliant.

* there is a const **per rule evaluation**.

- with **custom rules** you have to have **custom lambda** written. This is the thing that will be reacting to config changes **and marking them as compliant or not**.

* the rule sends a **CloudWatch event**. You can intercept that event using **CloudWatch rules** and do smth with it.

#### Multi-account

- you can get a **multi-account, multi-region view of configuration and compliance**.

* integrated with AWS Organizations

- does not cost anything for existing Config consumers.

* uses the notion of **aggregator**. This aggregator is responsible for discovery and aggregation of compliance data **from AWS Configs enabled on other accounts**. You **still have to have AWS Config turned on on other accounts**

- there is a **limit** for the amount of **aggregators** you can have. The **default limit is 50**

#### Remediate mechanism

- you can choose an action which is **predefined by AWS**.

* you can also use **CloudWatch events** with eg. a lambda function

- you can create an **Automation rule** within **Systems Manager**.

* **AWS Config** will not remediate anything for you automatically.

#### Notifications

- **notifications** are available only for **SNS** for the **whole AWS Config**. This means that if you want to have **granular notifications, you should use CloudWatch events**

#### Compliance

- you can create **compliance dashboards** by dumping the data into s3 and using **QuickSight**.

### License Manager

- enables you to keep your **licenses in a centralized place**

* you can **create alerts for violations**

- you can define **license limits** eg. **do you want to count vCPUs? or maybe websocket connections?**

* **integrates BUT IS NOT PART OF Systems Manager**.

- you can **track licenses on-prem**

* can be **integrated** with **Service Catalog** for **granular access to given list of products**.

### Systems Manager (SSM)

- helps you manage **large landscapes** like **100s of instances**

* can be used for **both Windows and Linux machines**

- there is an **agent (baked into modern Windows or Linux AMIs)** which you can install. That means that **you can manage on-prem instances aswell!**.

- there are **documents** which are basically **scripts (actions) that SSM can use to do stuff**.

- you **do not have to have SSH ports open**. You also **do not need SSH keys**.

#### Managed instances

- to be able to use SSM **EC2 require IAM roles** and **on-prem instances require `activation`**.
  Remember that **when you have instances managed on prem** they will have **`mi`(managed instance) prefix**. **AWS instances have `i` prefix attached**.

* **activation on prem** is **done using single activation code**. You do not have to create separate code for every instance.

- instances are **integrated with `AWS Config`**. The `AWS Config` gets the inventory data, thus you can see a timeline of installed stuff.

* you will need to create a role for the managed instances SSM agent.

#### Automation

- automate tasks that have to do with EC2.

* you create **automation documents**. These can be used to **eg. create AMIs**.

- can be paired with **EC2 rescue document** to enable auto-repair for your instances.

* for **more complex scenarios like error handling** you should prefer using `Step Functions`.

- can be used to **update your golden AMI**. This is done by launching an instance, running the user script, creating AMI from the instance and shutting the instance down.

#### Session Manager - Running commands on instances

- this **can be** done **without using SSH** for both **Windows** and **Linux** instances

* remember that you **have to have newest SSM agent installed**

- you **instance** **has to have IAM role** that **enables SSM API on that instance**

* you can have **IAM policies (attached to users / roles)** which futher enhance the security of your solution. You can **restrict access to a specific instances** while using `Session manager`, the action is `ssm:StartSession`.

#### Patch Manager

- there is a concept of a **patch manager**. This tool enables you to **discover missing updates, update multiple or single instances**.

* there are multiple **patching strategies available**.

- to apply patches you should **combine maintenance windows with `Run Command`**.

* your **patches can have custom origin!**.

- the `AWS-DefaultPatchBaseline` is actually **for linux instances!**

#### Run Command

- the `Run Command` is very important.

#### Maintenance windows

- can be used for **task automation on schedule**

* has a notion of **tasks** and **targets (EC2) instances**

- you can use **CRON expressions**.

#### State Manager

- tool to control **how and when configurations are applied**

* you could use it to update anti-malware definitions files, turn off / on ssh, RDP all that stuff

- this is basically to make sure your instance is in a desired state. And if something happens to make it go back to that desired state

* can also be used to for **task automation on schedule**. This is done by creating an **association**.

#### Inventory Manager

- collect data about system properties, roles, etc.

* can be visualized using `QuickSight` (**by sending data to s3 first**).

- to **ensure** that you are **tracking all instances**, create a **periodic check with lambda and CloudWatch events**.

#### Parameter Store

- it is quite important to remember that **SSM is public space endpoint (can be a interface endpoint)**. This means that your **EC2 instances have to have internet connectivity / interface endpoint within a vpc**.

- **secure storage** for **configuration data and secrets**

* the data is **structural**. You declare **paths to given values like an url: /dupa/dupa1**

- **keys can be encrypted (secure string)** using **KMS**. They also have **versions, and history of edits**.

* **scalable** and **serverless**.

- very useful for **distributing config inside ASG**

#### Secrets Manager

- while you can store secrets within `Parameter Store`, **`Secrets Manager` rotate the secrets (automatically or programatically)**;

* this is an **ideal solution for passwords for DB for smth**.

- if you need to **rotate secrets programatically**, you can do it by using **lambda**.

* there is a **limit** when it comes to **size of the secret**, that is **65kb**.

##### Advanced Parameters

- you can switch any standard parameter to advanced at any time. **You cannot go back from standart to advanced**

* with advanced parameters your parameter can have **up to 8KB of size**. This is **twice the size of a regular one**.

- with advanced parameters you can get **really granular per parameter policies**.

##### Rotation

- **RDS** secrets **can be rotated automatically** by the service itself

* you can **manually rotate the secret** using a **lambda function**. Be careful though, **always have at least 2 sets of users or credentials you are manually rotating through**. This is because **replacing the credentials for 1 entity might cause disruptions**. There is a **lag between the change you made in your lambda and the SecretsManager propagating this change**

- when you **first** configure the **secrets rotation**, the **secret is rotated**. This means that the value of that secret is different than the initial value that you set that secret to. This might cause confusion where you forget to **update your application to actually use the `Secrets Manager` and you cannot connect to eg. a database.**

#### Billing

- **SSM itself is free** but there are various services that have some cost:
  - **on prem instance management** is NOT free
  - **advanced parameter store** is NOT free.
  - **SSM automation** is NOT free.

#### With `CloudFormation`

- you can use SSM parameters as `Parameters` within `CloudFormation` templates.

* whenever parameters change, you can call `UpdateStack` API and update your stack with new parameters

- might be useful for **EC2 size / AMI configuration**.

#### Configuration Compliance

- scans the fleet of instances for patch compliance and configuration inconsistencies

* it will **pull data** from **Patch Manager** and **State Manager**

### AppConfig

- somewhat an alternative to the `SSM`

* think **SSM but with rolling deployment options** (including rollback)

- there is a neat _extension_ which runs HTTP endpoint which acts as a caching layer between your lambda and the _AppConfig_ service.

### AWS Firewall Manager

- this a **regional service**

* this service **wraps WAF**.

- allows you to create **WAF rules** that are **cross accounts** and **span multiple resources**.

* it **requires you account to be within Organization**. You can create policies there for all accounts (newly joined inherit those)

- it **integrates with AWS Config**. So **whenever a new resource is created** Firewall Manager can **apply rules to that resource**

#### Uses

- applying consistent WAF rules accros all accounts resources inside Organization

* notifying you when AWS Config resource is not-compliant and fixing that for you.

### AWS Service Catalog

- **describes** all **services you offer**. Very **much like online store** but instead of buying eg. food **someone buys products you provide**.

* **region aware** service.

- _Service Catalog_ **can be useful for Tags governance (making sure that resources have tags associated)**.

* can have **triggers** for different events like product deployment etc. **Lambda trigger for product deployments IS NOT AVAILABLE!**

- there is a hierarchy, **_Portfolio_** contains many **_Products_** which are just _CloudFormation_ templates.

#### Constraints

- even if you as a **user do have a readonly access to aws**, you would still be **able to deploy a product from a portfolio** if the **provider provided deploy permissions on portfolio**.

* it is **possible** to have **portfolio without launch constrains**. This means that **user IAM permissions will be used to deploy the product**.

* since the **user is interacting with parameters from CloudFormation**, you as an portfolio admin can **place constrains on those parameters**, like you can only deploy on t2.micro or t3.large or smth like that.

- there are **multiple constrains available**:
  - Launch
  - Notification
  - Template: limits the options of end users when they launch a product
  - StackSet
  - Tag Update

### Access Advisor

- will tell you **what services have the user access to** and also **when he accessed them**. This is a tab within IAM users console

### AWS For Enterprise

#### Workspaces

- gives you **access to a remote workstation**. You can ssh into a machine and do stuff.

* **can be placed within a VPC**. When you restrict the outbound you have created pretty secure environment.

#### AppStream

- gives you **access to a single application (a program)** which usually cost much.

* **can be placed within a VPC**. When you restrict the outbound you have created pretty secure environment.

#### AWS Connect

- basically a **call center with routing**.

* provides **interactive voice responder**

- **integrates** with **AWS Lex**

#### Chime

- basically **the same as Google Hangouts**

### AWS Lex

- to building applications that **listen and respond** with either **text or voice**

* it is used in **Alexa** and **AWS Connect**

### AWS Polly

- used for text-to-speach conversion

* Lex can integrate with Polly to create a chat and speach bot

### AWS Comprehend

- provides text sentiment analysis. Intentions, is it negative or positive, all of that.

### AWS Transcribe

- used for translating audio to text. You will probably use it along with AWS Translate to create subtitles.

### Sage Maker

- **analyze ML data and develop the models**

* you can use **AWS managed Algorithms**

### AWS Elemental

- familiy of tools that have to do with live-streaming or video

* **Media Package** creates **video streams for multiple resolutions**

- **Media Store** is a **media-optimized origin store**. Cuts video latency

* **Media Tailor** is for **placing ads inside a video**. This is **done BEFORE media delivery**. Also provides **metrics for those ads**

- **Media Package** is just a **package of these tools**

### S3

> S3 stands for **Simple Storage Service**

- S3 is an **object storage**, it allows you to upload files.

* Data is spread between multiple devices and facilities

- There is **unlimited** storage. Maximum file-size is 5 TB though.

* **S3 IS NOT A GLOBAL SERVICE**. It has a universal namespace but the **data stays in the region you created the bucket in (unless you specified CRR**.

* By **default** uploaded objects **are NOT public**. You have to set them as public manually

- **if amount of reads/writes is high** consider adding **logical or sequential naming to your S3 objects**. Amazon rewrote the partitioning mechanism for S3 and it does not longer require random prefixes for performance gains

* S3 has an **universal namespace**. That means that Bucket names has to be
  unique globally.

* Object consists of:

  - Key: simply the name of the object
  - Value: data (bytes)
  - Version ID: **S3 allows you to have multiple versions of a file**
  - Metadata
  - Sub-resources

- Buckets can be **replicated to another account or to a different bucket
  (region needs to differ)**.

* S3 can be `accelerated`. There is something called **S3 Transfer Acceleration** where users upload to **edge locations** instead directly to
  the Bucket (your downloads will also be faster). Then that uploaded object is **transferred automatically to your bucket**.
  One thing to keep in mind that you will not always see the speed increase. It is dependant on where you are making the request from.
  Maybe there is no edge pop near you?

- S3 is a **tiered** storage

  - **S3 Standard**, stored across multiple devices and multiple facilities. **Standard can tolerate AZ failure**

    - **RRS (reduced redundancy)**: this one is designed for **non-critical data** that is **stored less redunantly on s3 standard**. You would use this for **data that can be replaced / replayed**.
      While using this might seem like a good option, this **is not the cheapest solution, when it comes to storage, it is the chepest when it comes to the API usage though**.

  - **S3-IA/S3 One Zone-IA** (**Infrequent Access**): for data that is accessed less
    frequently but requires rapid access when needed

  - **S3 Glacier / Glacier Deep Archive**: used for data archiving, where you would
    keep files for a long time.

  - **S3 Intelligent-Tiering**: this is **great for unknown or unpredictable access patterns**. Will **automatically place** your files in **most optimized storage option**. There is a **monthly cost for using this option**

- **tiers** apply to **object**, **not a bucket!**

#### Select

- allows you do **get a piece of data from s3**. Like **some part of the object contents**.

* you can write **expressions in SQL**.

- it can **read multiple formats** and there is **no schema involved**.

* this is **mostly for single files**. If you need **analytics, use Athena**.

#### Glacier

- actually separate from S3 but AWS is integrating it within S3 heavily.

* **Glacier / Glacier Deep Archive** is an **immutable store**. That means that you cannot edit stuff once it's there.

- **Glacier/Deep Archive** have something called **expedited retrieval**. This allows you to **get the data from glacier and not wait Days/hours/Months**. There is a catch though. The **expedited retrieval** is **using resources from shared pool of resources**. There might a case there those resources are unavailable. This is where **provisioned glacier capacity** is used. This **ensures that your expedited retrieval request will not be rejected**. But you have to **pay more (additionally for it)**

* **S3 NON-Glacier** supports **object-lock**. This **along with versioning allows for immutable objects**, but you **have to specify the amount of time the lock thingy is present**.

- **used** by **virtual tape library** (underneath)

* you can **recover files from GLACIER** within **1-5 mins (expedited), 3-5 hrs (standard), 5-12hrs (bulk)**.

- you can **recover files from DEEP ARCHIVE** within **12 hrs (standard) and 48hrs (bulk)**

* you can retrieve specific component of an archive (a single file).

##### Glacier Vaults

- _Glacier Vaults_ are **containers for _Glacier Archives_**

* **glacier vault** can be given **access by using IAM roles**.

- you can create **vault archives UP to 40 tb** (this is usually a .zip file).

###### Vault Lock

- you have **24hrs** after creation **to confirm that the lock you created meets your requirements**. If that's not the case **you can abort it** or **complete it**.

* vault lock is **immutable**. **Once approved cannot be changed or deleted**.

- you can **attach vault lock policies to the vault** eg. nobody can delete anything out of it (or use MFA). This is mainly used for **compliance controls and tightening restrictions on data access**. This is useful for **preventing deletes and such**

#### Consistency

- **strong consistency Read after Write** for **PUTS**. Basically you can read immediately after
  you write to a bucket

- **Eventual Consistency** for **overwrite PUTS and DELETE**. Basically if you
  delete something or override it, it takes a second or two for you changes to
  propagate and take an effect.

* **updates to a single key** are **atomic**. Only one person can update given object at given point of time.

#### Consistency Gotchas

- **404 get, upload, then get** will **result in eventual consistency**. This is due to **how internals of s3 work**. It seems like there is **some cache involved**.

#### Limits

- maximum of **5TB for a single file**

* maximum of **5GB for a single upload** (use **multipart upload**)

#### Versioning

- S3 have the notion of versioning: **stores all versions of an object including all writes even if you delete it!**

* After **enabling** versioning, it **cannot be disabled**, only **suspended**

- Remember that with versioning your previous versions persists on the bucket. This can lead to exponential growth of size when editing big files.

* When an object is deleted, **bucket may seem empty** but that's not the case. You just placed a _delete marker_ on that object (thus creating a new version). Your **previous versions are still there!, you can view them within versions tab!**.

- You can restore your deleted objects by **deleting a delete marker**.

* **you are still billed for old versions**. Event though there might a delete marker the object is not permanently deleted so you have to pay for the storage.

- remember that **elements will be versioned from the moment you enable versioning**. **Any existing object will get the initial version of null**.

* you **cannot set versioning on object level**. Versioning is **always set on the bucket level**.

- versioning has **no impact on signed urls**. They will work as before.

#### Life-cycle rules

- can be used to **transition objects to different TIER of storage after X amount of time**, this can be placed on current and previous versions.

* With life-cycle there is notion of expiration. Basically your objects can be **expired after X amount of time**. Object can be in a _expired state_ where it will wait to be permanently deleted.

- Can be used with **conjunction with versioning**.

* **once placed in a tier, life-cycle rules will not be able to place the item back a tier (where it was)**

- they **apply** to **buckets, prefixes, tags and current or previous versions of the object**.

#### Security

- **by default only the account that created the bucket can do stuff with it**

* if you want to track **policy changes** use **CloudTrail**

##### ALCs

- **considered legacy** but I find it **relevant in day-to-day** job when it comes to s3

* **not granular**, with ACLs you can **grant permissions to AWS accounts and pre-defined groups**

- to **identify the AWS accounts** you will be using something called **Canonical user ID**. The ID is basically an **obfuscated form of AWS account ID**.
  to get your AWS account information, as well as your _Canonical user ID_, click on the dropdown with the account name and go to _Security Credentials_

* to **identify groups** you will be using **URIs**. These are looks something like `http://acs.amazonaws.com/groups/global/AuthenticatedUsers`

- they **can apply** both to **objects** as well as the **bucket itself**

###### Canned ACLS

- **predefined permissions**, like a bundle so that you do not have to fiddle that much with ACLs

##### Resource policy

- when you want to assign policies to the resources you do not control, you should be using **resource policies**, in this case know as **bucket policies**. This policies **apply to any identities accessing this bucket**.

* using **resource policy** you can **control access** to s3 resources, works on **identities you DO NOT control. THIS ALSO MEANS ANY IDENTITY**. When resource policy is used **specifically with s3**, it is known as **bucket policy**

##### Bucket policy

- **much more granular than ACLs**

* allow you to **operate on _indentities_** so users, groups and roles

- you would use **identity policies** to **control access** to s3 resources, this however **only works on identities IN YOUR ACCOUNT!, identities you control**.

##### Endpoint policy

- s3 endpoints allow you to have granular permissions (basically a bucket policy) for a narrow scope of your objects. This policy does not "pollute"
  your main bucket policy because it's attached to the endpoint itself

* you would **use this** if your **main bucket policy is getting out of hand**, as in it's getting too big or unmanageable

- **the effective permissions are an intersection between the main bucket policy and the endpoint policy**

##### Block public access

- can be enabled on the account or the bucket level. Account level should be preferred

* overrules any ACLs and policies that you specify. If you have this enabled, objects will not be publicly accessible, no matter the policy or ALCs
  that you have in place.

###### Overall

- **bucket policies** and **ALC** suffer from the fact that **usually, the uploader is the owner of the object**. Imagine giving someone permissions to upload (from another account eg.) and then you cannot delete that file. **You can guard against that** by **creating resource policy** where you **deny any requests when full permissions were not granted to the owner while uploading given object**.

So when to use what?

- _identity policy_
  - when **you control the identity**. That means **identities within your account**.

* _bucket policy_ (resource policy)
  - when **you DO NOT control the identity**

- by **exposing a role which can be assumed** by third party you are **making sure that YOU as a bucket owner stay as the owner of a given object**. This is why you would want to use _bucket policies_ when you do not control the identity.

* if you are doing anything serious (like you know, your day to day job stuff), encrypt everything. You can start the default SSE-S3.

##### Presigned URLs

- **pre-signed urls** work **per object level**

* **pre-signed urls** always **have the permission of the identity that signed the URL**. They **expire** or can be **made useless sometimes, when the permissions of signing identity changed**

- you can use **referrer IAM condition** to **allow request from specific web pages**.

* you can create **POST** or **PUT** signed urls. The **POST** version allow you to specify **conditions**. The conditions are pretty powerfull,
  you can for example, assert on the size of the object being uploaded.

##### Access logs

- **server access logs are delivered to another s3 bucket**

* the **logs include all bucket / object events**

- with **Server access logs** you can **write access logs to a different bucket**. Remember to **give Log Delivery group permissions to write to dest. bucket!**.

##### Precedence

There is a precedence in which **acls** or **bucket policies** are evaluated.

1. Does the requester has the permission to do stuff on s3?
2. Does the bucket policy has the necessary permissions ?
3. **If the bucket owner is not the owner of the object** does the **object acl** allow for the action?

The step 3 is crucial. Remember that whenever you upload something to a bucket that is not yours, you can control how much access to the underlying object the bucket owner has.

#### Access Points

- having multiple different statements in bucket policy can be painful. Especially when we are talking multiple users

* _access points_ allow you to **associate a special ARN and attach a bucket policy to it**. You can then specify that arn as a _bucketName_ while making requests with SDK

- this way you **can have 1 policy per 1 user per 1 endpoint**

* the **_bucket policy_ still applies**, that is **the resulting permissions are an intersection of the access point policy and the root bucket policy**

- you **have to use sigv4 to access objects through an _Access Point_**. I personally wish it was not the case as they do not support anonymous users.

#### Encryption

- **PER OBJECT** basis

* there are multiple options:
  - **SSE-C**: **keys** are **stored by the customer**. It does **not allow for role separation**. **You manage the AWS master key**.
  - **SSE-KMS**: **keys** are **managed by KMS service**. It **allows for role separation** since **keys are stored in accounts KMS**
  - **SSE-S3**: The **master key is within AWS**. Be careful with this one though. Since S3 has the key, when you make your bucket public, s3 will automatically decrypt the contents.
    This is why you should **prefer SSE-KMS** whenever possible.

- **no additional cost** for enabling encryption

* there is something called **default encryption**. This will ensure that every object you put into a bucket will get encrypted. **You can use AES-256 or KMS**.

- remember that **encryption only applies to objects uploaded AFTER the encryption was turned on**. Existing objects are NOT encrypted.

* you can **override the default encryption PER object basis when uploading**.

##### SSE-S3

- **objects encrypted by S3 using KMS on your behalf**

* **keys** are **stored on the objects themselves**. If you have permission to read the object you can decrypt it

##### SSE-KMS

- **objects encrypted using individual KMS keys**

* **keys are stored within S3 object**

- **decryption of the objects** requires you to have **both iam permissions for the S3 operation and the KMS key permission**. That is why it allows the separation of roles.

#### Replication

- s3 enables you to turn on **cross region replication of given bucket**, but **you have to have versioning enabled ON THE ORIGIN AND THE TARGET BUCKET to be able to enable CRR**

* you **can** also **enable same region replication**

- **does not support SSE-C**. The replication **process is purely server-side**

* when using replication **SOME THINGS DO NOT CARRY OVER to the CRR destination**:
  - **lifecycle rules**
  - any **existing objects before replication was enabled**. You can **change storage class of existing objects** to **force the replication to happen** though.
    Please remember that it has to be **different storage class** in this case.

- **it is possible** for an **object to change storage class and object permissions while in the process of CRR**

* you can **override the owner of an object** when that object (due to CRR) is **going to another bucket**. This might be helpful to implement some kind of security measures

- replication is a **asynchronous process**.

* **SSL** is **enabled by default** when using replication

#### Misc

- there is something called **requester pays**. This is where the **person who downloads the object pays for the transfer**. this is **this feature can only be used with people who have an existing aws account**.

* you can use **BitTorrent** to **distribute s3 content**.

- using API, you can download **specific byte range** for a given object.

#### Events

- these can be send to **SNS, SQS or Lambda**.

* the events are **object based events**. If you want **all s3 events** you should:
  - create **`CloudTrail` for this bucket**
  - use **`CloudWatch rule` with `CloudTrail` integration** and select given destination.

### Storage Gateway

- prefer if you have any kind of integration with `on-prem`.

* does not have SLA

- **SOMETHING YOU DOWNLOAD**

* Physical/virtual device which **will replicate your data to AWS**.

- There are 3 flavours of Storage Gateway
  - **File Gateway** : used for storing files as object in S3 - **NFS, SMB** .
  - **Volume Gateway**: used for storing copies of hard-disk drives in S3 - **iSCSI**.
  - **Tape Gateway**: used to get rid of tapes - **iSCSI**, for use mainly with **backup software**.

* With **Volume Gateway** you can create **point-in-time backups as EBS snapshots**

- if you see **ISCSI** that is **probably Volume Gateway**.

- **useful** when doing any kind of **cloud migrations**

* all data is **encrypted in transit**. Data **at rest** is **by default encrypted using SSE-s3**

#### Volume Gateway

- you **cannot directly access the data using S3 API**. You would need to **use File Gateway to have a native way of accessing your files**.

* you can **create snapshots from the volume**. These can be **turned into EBS snapshots**.

- you can **create mountable ISCSI devices** which you can **mount on prem**.

* this is a **low-latency solution** for **either cached or non-cached gateway**.

- volume gateway creates **asynchronous EBS snapshots of your data to s3**.

##### Cached vs Stored

The main difference is where the most of your files are stored. When we are talking about **cached volumes, most of your files are stored in s3**. There is a small cache buffer that holds the most accesed files.

When using **stored volumes, all of the files are stored on prem and synced with s3 ansynchronously**.

Both offerings store underlying data as **EBS snapshots on s3**.

#### File Gateway

- **exposes itself as NFS/SMB**

* the **latency might be higher than volume** if you **are not using cached gateway**.

- when you have **tapes** `on-prem` **you can still use `File Gateway`**. The usage **depends on the need** (you might not want to archive but work with the data).

* it uses underlying cache for files so that s3 is not queried all the time.

#### Tape Gateway

- used to **either store the tapes within amazon** or **for achriving purposes eg. migrating to glacier**

### FSx

#### FSx for Windows Server

- actually **can also be used with Linux 🤔** (but NFS is still much more preffered for Linux)

* exposes **SMB** or **NFS** file share

- you can **connect** to it **via DNS name**

* **integrates** very well with **AD**

- **DOES NOT** have **caching capabilities**

* **can** be **accessed via DX or VPN from on-prem**

* just **like EFS** it **automatically scales**

#### FSx for Lustre

- this is a **very high performance file storage**

* used for AI stuff

- just like FSx and EFS **it automatically scales**

* **integrated** with **S3**

#### AWS Import / Export

### Snowball

- Big briefcase, **up to 72TB** of **usable storage** storage.

* only used for transferring. There is **no compute included, YOU CANNOT UPLOAD TO s3 using PLAIN SNOWBALL!**

### Snowmobile (not joking)

- A truck with a container that carries snowballs.

### Snowball Edge

- - big briefcase, **up to 80 TB** os **usable storage**.

* there are **three versions**: **compute**, **storage optimized**, **compute optimized with GPU**.

- the versions speak for themselves eg. **you would use compute optimized for performing machine learning analysis at remote location** and then **transferring the data**.

* with the `Snowball Edge` you can **upload stuff to s3 or perform some compute on the data**.

- integrates with **AWS Greengrass (IOT)**

### AWS Import / Export disk

- you **mail** a **psychical drive to AWS**. They make either an AMI, upload to s3, or EBS snapshot out of it.

* you should **always prefer snow**. The I/E Disk you have to encrypt the drive manually, worry about the transporation etc...

### RDS (Relational Database Service)

- AWS service for relational databases

* multiple providers such as: `mysql` or `Oracle`. Please note that **RDS does not support Oracle RAC** and has **semi-support of RMAN**.

- **RDS AUTO SCALING in terms of compute does not exist**. You have to provision an EC2 instance and **pay regardless of the db usage**. What is possible is **RDS storage auto scaling**

* RDS **can be in multi AZ configuration**

* you **can connect to** the underlying **EC2 instance that hosts the database**. Use SSH or different means.

- you **DO NOT** have **access to the underlying OS, the DB is running on**.

#### Read Replicas

- **up to 4 instances** can be associated with the **replication chain**

- with **read replicas** you **CAN SPECIFY to which az to deploy**

* you **can target read replica** but you **CAN NOT target second-master (multi-az)**.
  You can use **R53 weighted policy for that**

- you can **create private hosted zone on R53** with **multivalue answer** to **target multiple read replicas** with one DNS query.

* **read replica can be cross region** - similarly to `Aurora`. This is very important to know!.

- you can have **up to 5 read replicas from a master instance**

* you can have **read replicas of read replicas**

- data **synchronized asynchronously**

* **master has to have backups enabled to be able toprocess.env.AWS_EMF_ENVIRONMENT = "Local";`**

#### Multi-AZ

- when choosing **multi-az deployment** data is **replicated synchronously**

* when choosing **multi-az deployment** you **CANNOT specify which AZ it will be deployed into**

- **with multi-az deployments standby db only comes into play when your primary failed**.

* when patching os on EC2, **with multi AZ config, patching is done first to standby in different AZ then failed over onto when main db is down due to patching os**.

- when **upgrading db engine**, **both master and slave** are **taken offline**. If you want to avoid downtime at all costs, **create read replica** and **update the engine on the replica**. Then promote the replica. Remember that you have to update the **EngineVersion property** in CF

- since the **failover, just like master is within a VPC** you have to **make sure that the failover subnet routing rules are the same as master**. Otherwise you might have a situation where a failover happens and you cannot connect to your instance because routing rules are not configured correctly.

#### Maintenance

- **with multi AZ enabled** any kind of **maintenance, security patching** is **performed** first on the **standby**. When multi AZ is not specified, this process will take DB offline.

* takes place **during maintenance window, specified when creating a RDS db**

- you can **defer** updates. Updates **marked as `required` cannot be deffered indefinetely.**. Updates **marked as `available` CAN be deffered indefinetely**.

* when you **replace the `DBVersion`** both the **primary and secondary will be taken offline**. This is completely different behaviour than patching the underlying OS.
  To make sure your application still works, deploy the secondary stack with RR

#### Encryption

- you **can only choose to encrypt your DB during creation phase**

* data **in transit** between **source and read replicas** is **encrypted by default**

- you **cannot encrypt existing DB**. You have to **create an snapshot and encrypt it** and build DB from that snapshot.

* **read replicas** have to be encrypted with the **same key as source AS LONG AS THE Source and ReadReplica ARE IN THE SAME REGION**

#### Backups and Restore

- you can perform **manual database storage-level snapshots**. These **will NOT be destroyed when you terminate db, YOU have to delete it manually when you want**. You would most likely use **step functions and lambda** to carry out this task.

* **automatic backups are also an option**. These have **retention period UP to 35 days**. You have to **specify backup window, backup window has a duration**.

- **backups are incremental and they BACKUP THE DATA THAT IS ACTUALLY CONSUMED (existing within your db)**

* **automatic backups** will **be deleted after the retention window, deleting db DOES NOT automatically delete backups**

- you can perform **point-in time recovery**. **RDS automatically performs transaction backups every 5 mins, stored on s3**.

* when you **restore a RDS DB, you will actually be creating a brand new DB instance**. You basically have to setup it again, but this time you have data from a snapshot present. **Restored DB has it own CNAME, you have to remember about specifying the correct security group**

- **db manual snapshots can be copied between regions**

* **snapshots from encrypted DBs are encrypted**

#### RDS on VMware

- you can put RDS dbs on VMware.

* this service does what normal RDS does so **patching, availability protection**.

- you can have **multi-az instances** just like on regular RDS.

* **point-in-time restores** are also available.

#### Replication on-prem

- even though it's not natively supported you can replicate your RDS db to on-prem mysql db.

* you will need a VPN connection and `mysql dumps`.

#### Security Groups

- your DB **can have security group associated with it**

* this is especially useful to know whenever you want to **acces your DB within VPC**. Not using correct rules on the secuirt groups **can lead to connection issues**.

#### Storage

- since RDS runs on EC2 you can change the underlying storage type.

* the underlying storage can be of type **GP1, IO1, Magnetic**.

#### Option Groups

- since RDS can run multiple engines, **option groups** allow you to **specify which options of the underlying engines you want to use**.

* a good example for this would be **S3_INTEGRATION** while using **Oracle on RDS**.

#### Cross region failover

- do not mistake this with Multi-AZ failover, which promotes the read replica automatically

* use CloudWatch alarm which will trigger lambda which updates R53, multi-region read replica will be promoted.

#### Streaming

- RDS has some CDC capabilities

* you can **stream** data **to Kinesis** if you are using **PostgreSQL**

#### Notifications

- you can have the RDS service sent you various notifications through SNS

* these notifications are really specific, can be about failures or new resources being created

### Aurora

- **SQL and PostgreSQ** compatible

- **storage layer lives outside of the database itself**

* **automatically replicates the storage between 3AZs (6 replicas of storage)**

- **can** have **multi-az enabled on the database layer (the instance)**

* there a notion of **cluster which has shared storage (up to 64tb)**

- you **pay for the storage you are consuming / consumed**.

- uses **subnet groups**. This is basically telling aurora to **which subnet to deploy to**.

* **supports schema changes** through **fast DDL** (DDL are the operations that have to do with the tables and the structure)

#### Security

- you have to **specify SG, which is applied to the DB**

#### Writers

- **primary node of a cluster**

* **only** instance **used to write to the cluster**

- can be scalled up

#### Readers

- **only** for **making reads to the db cluster**

* **they use the same storage**, so the **replication is synchronous**.

- **you can target them for reads**. They **operate on between the realm of standby and Read-Replica in RDS**. You can create **up to 15 Readers**.

* **can be very easily added**

- **can be used for failover**

* you can have **readers in different AZs**

##### Reader Endpoint

- this endpoint **maps to every reader within an instance (1 url)**. This is quite different than RDS since with RDS you can only target 1 read replica at one time (there is no 1 url that maps to every read replica).

* you can also target an individual instance using special instance ID endpoint.

- this is **not HTTP-call based endpoint (same goes for writer endpoint), you still have to manage connections!**

#### Multi-Master

- you can have multiple **master instances**. That means that you have multiple instances which can **read and write**.

* this means that, **not like using RDS** where **you cannot scale writes**, here **using aurora you can scale writes** using **multi-master configuration**.

- **ALL NODES OF `Aurora multi-master`** have to be **in the same region**. This makes DynamoDB only database that supports multi-region multi master configuration.

#### Parallel Query

- massive **performance benefits** for **long-running queries**

* you have to manually select this option.

#### Failover (tiers)

- this is to help aurora decide on which instance to failover

* there are **15 tiers**

- **failover will be automatically performed to other reader instances**. This will be **done much faster than the multi-az failover on other RDS types**

#### Scaling

- **scaling for writes** means **increasing the size of the primary instance (the master) => scaling up** or **using multi-master**

* **scaling for reads** means **increasing the number of reader nodes => scaling out**

- you can enable **auto-scaling for readers/writers**.

#### Endpoints

- endpoints **correspond to the reader and writer instances**

* **automatically extended if you add more readers**

#### Backups

- when using **standard RDS, the only option is to restore from backup**

* just **like RDS you can specify the backup window and the retention period (35 days)**

#### Backtrack

- this somewhat compares with RDS and restoring from backup. With RDS new DB instance is created, no downtime occurs but you have to switch connection strings.
  With Aurora, there is an DB outage but no new instance is created. Saves you some time in a long run.

* has some cost associated with it, **you have to pay for storing the changes**

- **causes DB outage**

* **rolls back the shared storage**

- **YOU DO NOT HAVE TO CREATE NEW DB INSTANCE**, yes the db will be offline for a short period, but it will save you creating a new instance

* you can backtrack **up to 72hrs in time**

- can be used to **get rid of data corruption**

#### Cloning

- is made from data snapshot

* instances are different but it the cloning itself is much faster.

- you are **only paying for the differences in data** **between** the **clone and the origin**.

#### Monitoring

- usually **using CloudWatch**.

* you can **see the RAM usage without the need to install CloudWatch agent**. This **applies to RDS in general**

#### Replication / High Availability

##### Cross Region Replication

- you can **create cross-region clusters that act as cross-region replicas**

- this is **logical / binlog replication**. It uses **network for the replication part** so there **might be some lag involved**.

* you can create **readers** within that **cross-region cluster**

- you can **failover onto cross-region cluster**

* the **FAILOWER onto cross-region cluster takes FEW MINUTES**

##### Global Database

- it uses **psychical replication: dedicated hardware to do the replication part**.

* the **replication is much faster than CRR read-replica**

- **failover is much faster than CRR read-replica**

##### Crash recovery

- Aurora will **automatically recover from crashing** on parallel threads. This means that **there is also no downtime "almost instantaneously"**.

* this is **not the case with RDS** which, when crashed, **usually takes up to 30 mins to recover**.

##### Dealing with replication lag

- make storage size the same between the master and the replica

* `max_allowed_packet` setting should be the same. Mismatch in this setting can cause replication to fail altogether.

- turn off `Query Cache`

#### Aurora Serverless

- uses **ACUs** which is a **unit of measurement for processing(compute) and memory in Aurora Serverless**

* **charged** based on the **resources used per second**

- used **for unpredictable loads and random surges of traffic**.

* can **rapidly scale**. Behind the db **there are multiple warm instances ready to go**. Because it **uses shared storage component just like normal Aurora** these instances can be used for scaling quickly.

- the **instance itself exist within a single AZ**. In the event of failure, _Aurora Serverless_ will provision new instance in another AZ.
  Failover time will be longer than normal Aurora.

* you can still **utilize normal endpoint just like using regular Aurora**. This might be useful when providing backwards compatibility.

* **operates within a VPCs**

- instead of initializing a connection you connect to a **DATA API (shared proxy)**. This is quite **similar to RDS Proxy**.

##### Aurora Serverless v2 (2020)

- can be a bit more expensive

* scales much faster. There is less over-provisioning going on since the compute increments in 0.5 ACUs instead of doubling the ACUs

- as of writing this, it does not support _DATA API_. This will probably come sooner than later since it's not even GA.

### Neptune

- NoSQL **Graph database**.

* usually used for social-network like websites, think facebook.

- highly available by default

* data is replicated by default, spread across 3 AZs

- you will need an EC2 instance to run int

* one neat trick I saw is to scale this DB up when you have to perform some kind of operations. Then exports the results to Athenta for quering and scale down the database.

### QLDB (Quantum Ledger Database)

- database for tracking pre and post states of data. This is very useful for **banking, payroll, finance, even voting**

* think bitcoin

- **ACID database, which enables you to use SQL queries**.

* fully **serverless**

### DocumentDB (with MongoDB Compatibility)

- cluster lives on instances within a VPC

* you can encrypt at REST using KMS

- supports **point in time recovery** with retention period (just like in RDS, up to 35 days)

#### Scaling

- just like with RDS you have **1 master instance (read/write)** and **up to 15 read replicas (per cluster)**

* you can **scale UP primary for writes**

#### Storage

- uses **shared storage layer, just like Aurora**.

- because of this replication is synchronous.

#### Restoration

- just like with RDS you **are restoring to a separate cluster, there is no backtract**

#### Failover

- **you can failover to read replicas**, just like with Aurora.

### DynamoDB

- **regional** service.

* storing data here is **more expensive than s3, remember!**.

- **more effective for read heavy workloads**

* **data automatically replicated `synchronously` between 3 AZs!**

- store **items which weigh more than 400kb in s3** and use **pointers**

* **compress large attribute values**

- **LSI** gets **the same WCU/RCU and partition key** as the **primary table**

* **to enable backups** you have to have **streams enabled**

#### Reading

- the **least ammount of RCU** you **can consume is 1**. It **does not matter if you specify Projection Expression**.

* one **RCU** is equal to **4KB**. That is **one strongly consistent read OR 2 eventually consistent reads**

- **separate hot and cold data**. This will help you with RCU provisioning.

* you can **save some bandwith** with **projection expression**. Remember **entire row is pulled, but then data extracted by Dynamo**

- you will need to use **projections** on **LSi** and **GSI**.

#### Writing

- one **WCU** is equal to **1KB**. **1 WCU is THE MINIMUM you consume**

#### Filtering

- filtering **does not reduce the RCU consumed**. Similarly to Projection Expressions, **every table with given index is pulled, then it is filtered @Dynamodb**

#### Cost

- you are billed on **overall storage** and **RCU/WCU**

* there is **minimal capacity** which is **1 RCU/WCU per second**. Free tier is **100 MB of storage and 5RCU/WCU - per month**.

#### GSI

- make sure to **project attributes that you use as an index**. Otherwise the **retrival will be costly!**

* you **cannot use strongly consistent reads** with **GSI**. If you find yourself in a situation where you need strongly consistent reads and different query pattern,
  re-create the table (migrate the existingone).

#### LSI

- the LSI **can be strongly consistent**

* **if you are using LSI, your partition can be up to 10 gb**. This **limit does not apply to tables WITHOUT LSI**

#### ACID

- **NOT ACID compliant**. This is due to the fact that you can have different number of attributes for a table row.

- DynamoDB **can be ACID compliant** using **Dynamo Transactions**.

#### Scaling

There are a few approaches when it comes to scaling with dynamoDB

- **adaptive scaling**: where dynamo automatically allocate RCU/WCU to a given partition from the provisioned RCU/WCU pool of all partitions

* **auto scaling**: you specify the **upper and lower bounds of RCU/WCU**

- **on demand scaling**: where you **do not have to set any limits / bounds for RCU/WCU**. DynamoDB will take care of provisioning those dynamically and adjusting to the

#### Dynamo Streams

- dynamo **streams** **hold** data for **24 hrs**.

* the streams are **considered poll based events**. These events are **ordered and guaranteed to hold an order**.

- you **can use TWO lambda functions for dynamoDB stream** but it's **not advisable**. You will face problems with **throttling** and so on. You should use 1 lambda function as a consumer and **implement fanout with kinesis** if you need it.

* while you cannot control the number of shards directly, **one of the things that has an effect on the number of shards is the capacity of the table**. This means that if you encounter a spike in WCU/RCU and you are using _on demand billing_, there might be a concurrency spike in lambdas reading off the stream for that tabl

- items are put onto a shard **based on their partition key**. The **`batchSize` and `batchWindow` might yield lower batches if you are populating the table with items from different "collections"**. Every item will be consumed though, your lambda will be invoked more times.

* the **throughput of the stream** is based on the **number of paritions** given table **has**. This is because, **the more paritions your table has** the more **underlying shards will be allocated for that table**

- DDB **does not use Kinesis under the hood**. **DDB streams use the underlying tech that powers Kinesis** but they do not use Kinesis directly

#### Integration with Kinesis Data Streams

- DDB can stream CDC events to Kinesis Data Streams

* since Kinesis is used, **you have the ability to retain the CDC events up to a year**

- technically you no longer have to write mapping code to put stuff into eg. Elastic Search. You can go Data Streams => Firehose => Elastic Search

* **as of writing this, there is no CloudFormation support, only console or CLI**

- as good as this feature sounds, **you have to calculate your shards yourself**. DDB will not do that for you

##### KCL Adapter

- you can use **DynamoDB Streams Kinesis Adapter** to process the stream data **when you are using KCL**

* the **API exposed** is **very similar** to the **Kinesis one**

##### Dynamodb Streams fan-out

- to combat the 2 lambda subscriber limit you can implement the fanout pattern

* since you probably want to preserve the order, you will most likely **1 subscriber pushing to _Kinesis Data streams_**

##### Triggers

- this is just a lambda function that is invoked whenever a stream event happens

#### TTL

- **attribute** which tells DynamoDB **when given item should be considered as `expired`**.

* DynamoDB will **delete expired items**

- this **value has to be EPOCH time**

#### Partitions

- **HOT partitions** are **thing of a past**. Before, you would need to ensure even distribution of reads/writes across partitions. This is because WCU and RCU was distributed evenly. With that setup your _hot partition_ might end up throttling and dropping requests. Now **with adaptive scaling** that no longer is the case. **Dynamo will automatically given partitions WCU/RCU based on the number of traffic it receives**. It takes away from the pool of WCU/RCU available to the whole table.

- remember that **indexes take up space!**. This is quite important and something you have to consider while creating your 20th GSI ;p.

* **the more data** is within the database **the more partitions there are**.

#### SDK

- when using **SDK**, it **automatically retries (calls are eventually consistent) unless retry queue is too large**.

#### Backups

- you can **backup tables to s3**

* the saved **backup contains table data + RCU/WCU units**

- you can **restore to a different region**

* you can also enable **point-in-time recovery (35 days retention)**

#### Global Tables

- you **need to have streams enabled (new and old)**

* this is a **master-master setup**. This means that you can use the table in other region as your new master. **all `global` tables replicate to each other**

- **removing** the **table from `global-tables` view DOES NOT delete the data**. If you really want to delete the data, you should remove it manually or just remove the table manually.

* you should consider the cost of consistency. The writes to the other global table are more costly than normal writes. These are not done by you, but by the service itself.

#### Encryption

- tables can be encrypted

* either **KMS(AWS Managed)** or **KMS (AWS owned)**

- by default all tables are encrypted by the KMS provided by the aws

#### IAM

- you can really setup **fine-grained control** when it comes to Dynamo

* you can have **policies that allow only access to specific attributes or items**

#### Costs

- you can leverage **reserved capacity for provisioned throughput** to save some \$\$.

* this works very similarly to other offerings. You pay upfront and save in a long term.

### CloudFront

- CloudFront is a **CDN**. Takes content that exists in a central location and distributes that content globally to caches.

* These caches are located to your customers as close as possible.

- Distribution is basically the **collection of Edge Locations**

#### Origins

- Origin is the name given for the thing from which content **originates from**.

* can be **s3 bucket**, **web-server (ec2)**, **Amazon MediaStore**, **APIGW** or **other CloudFront distribution / www site**

- **Origin has to be accessible to the internet**

##### Protocol Policy

- you can enforce **rules on the viewer and the origin**.

* available policies are **`HTTP ONLY`, `Match Viewer`, `HTTP/HTTPS` , `Redirect to HTTPS`**.

- please remember that not every origin supports HTTPs - look at the s3 static website hosting example.

##### S3 website endpoint vs S3 bucket as origin

- you can use both s3 website endpoints (you have to have static website hosting enabled) or plain old s3 endpoint (the bucket itself)

* when you are using **s3 website endpoint** as **origin** you **cannot create OAI**.

- if you want to **restrict access when using s3 website endpoint** you should use **custom header which CF injects**, then **setup bucket policy which looks at that header**.

* **s3 static hosting DOES NOT SUPPORT HTTPS**. That means that you **cannot set `HTTP ONLY` as Origin Protocol Policy**.

#### Distributions

- You **can invalidate cache** content

* There are **2 types of distributions**.

  - Web
  - RTMP (used for video streaming and such)

* when using **RTMP distribution your data has to live on s3**

- When you deploy CloudFront distribution your content is automatically deployed to edge location. You can specify which ones (limited to a country). If you are rich you can deploy to all edge locations

#### Cache

- Edge locations **cache content** (TTL)

* **Cache hit** means that when an user requested a resource (like a webpage), an edge location had that available

- **Regional cache** is like a meta edge location. Basically second level cache, fallback when there is **no cache hit**. If it does not have a copy of a given content it **falls back to the origin** (origin fetch).

* setting up **TTL 0** **does not mean NO caching**. It means that **CF will revalidate each viewer request with the origin**. That basically means that **CF will make a GET to an origin with a special header**. Then **origin signals CF if the origin changed or not**.

#### Protecting Content

- **By default** every CloudFront distribution **comes with a default domain name**. That domain of course works for HTTP and HTTPS. You can register domain and replace it.

* You can restrict the access on two levels (**You can** restrict an access to S3 only but it's no the topic of CloudFront):

  - on a CloudFront level, your bucket is still accessible though
  - on a S3 and CloudFront level, you can only access the website using signed urls.

* Restricting your CloudFront & S3 combo is done by creating **OAI**.

- **OAI** is an _identity_. That _identity_ can be used to restrict access to you S3 bucket. Now whenever user decides to go to your bucket directly they will get 403. To achieve such functionality you add **CloudFront as your OAI identity**

* **OAI can be applied** to **S3, CloudFront, bucket policies**. **OAI origin cannot be a website address of s3 bucket (static website hosting)**.

##### Signed URLS & Signed Cookies

- **configured on BEHAVIOR LEVEL**

* there is a notion of **trusted signers**. These are **AWS Accounts** which hold **CloudFront key pairs** and are used to **sign urls**.

- the account that created the distribution will be, most likely, the trusted signer.

###### Signed URLs

- very similar mechanism to S3, you generate signed URL using your permissions.

* enables the holder of the signed URL to access (read/write) object.

- that is important is that **signed url allows access for only 1 object**.

###### Signed Cookies

- cookies **extend the signed urls functionality** by giving you the ability to **grant permissions to groups of objects**

##### Field Level encryption

- your distribution have public-private key associated with it

* the traffic between **origin location and destination** is **encrypted (on top of HTTPS) using that key pair**.

#### Cache Behaviors

- by default CloudFront provides sane defaults for the cache TTL

* you can specify **per path (\*.img or /dupa) rules for cache TTL**

- you **can get pretty complex since CloudFront gives you a lot of settings**.

* you **attach cache policies to them**. You can create **your own cache policies** or **use the managed ones**

#### Cache Key

- this is how _CloudFront_ knows if the asset you are requesting is within a cache or not

* _cache key_ is **created based on the request parameters** like **url or/and headers/cookies**

- you can **customize this behavior** using **_cache behaviors_** and **attaching _custom cache policy_ to that behavior**

#### Restricting access to CloudFront distribution

- you can further place **restrictions** on **who can access content** available by CloudFront using **signed URLS and signed cookies**.

* The **mechanism is similar to that of s3**, but there is difference **who can manage the underlying key**. With CloudFront **you can protect custom origins, APIGW and static files**.

- Whats more **you can restrict access to a specific IP** using **WAF ACL and CloudFront**

- You can also use **Geo Restriction**. This is **similar** to **R53 geolocation**, but you get the **benefit of CDN (speed)**.

* **SNI** is a way to present **multiple certs to a client**. Client has to **pick which cert. it wants**. Some old browsers do not support this technology.

#### Origin Groups (Failover)

- with Origin Groups you can **create failover behaviors**, remember **CloudFront can have multiple origins**.

* you can **failover on a criteria of Status Codes returned from one origin**.

#### Combining with R53 (multi-region)

This will require 2 layers of DNS, one for the `CloudFront` itself (optional but still) and then an `latency-based` routing policy within `R53` as origin to `CloudFront` distribution.

This way, CF will fetch the data from the **R53 latency-based resolved host**. This is pretty neat!.

#### Injecting headers

- with `CloudFront` you can **inject headers to origin request**. This only works one way
  you **cannot inject headers for the response**

* to **inject response headers** you should **use lambda@edge**.

### API Gateway

- **throttling** can be **configured at multiple levels** including **Global and Service Call**

* you **can** setup **on premise integration**. You need to use **NLB and VPC link** to make it work. You would setup NLB within your VPC and connect to on prem through VPN or Direct Connect. Then that NLB would hit your on prem, APIGW would hit the NLB through VPC link

- you **can** setup **integration with public-facing internet (non-aws) endpoints**

* by default **AWS provides DDOS protection**

- you **can** enable **access logs**. This will enable you to **see the IPs of people calling your API**. This is **not CloudTrail!**. Remember, CloudTrail is about service calls made by identity within your AWS account.

* **access** to an APIGW can be **controlled using multiple means**:
  - **resource policies** which define **access to your API methods from source IPs or VPC Endpoints**
  * **IAM** can be **applied to entrie API or methods**
  - **authorizers**
  * **cognito**

- you **can** actually **use VPC interface endpoint** with APIGW

#### Response Codes

- **4xx means Client errors**. This is where WAF is blocking the request (403), or concurrency throttling happens (429).

* **5xx means Server errors**. This is APIGW or the integration failing.
  - **502 (Bad Gateway)** usually bad request, apigw got malformed response from eg. lambda.
  - **503** is service unavailable
  - **504: Integration failure** is the **timeout on APIGW level** - remember that the default one is 29 seconds

#### Timeouts

- you probably know that lambda has 15min timeout max. But that timeout only applies to async invocations (reading from the queue and such) where user is not waiting for a response. **The default APIGW timeout is 29 seconds**. This can be **configured to be between 5 seconds and 29 seconds**. All you have to do is to **un-tick the "use default timeout option**

* when timeout happens on APIGW level, the **error code will be 504**.

#### Responses

- all done within **gateway responses tab**

* you can **map the response returned from lambda to a different one**. Like **mapping 403 to 404 response**.

- you can add **custom headers to the responses**.

#### Usage Plans

- enable you to create throttling / quota limits per key group. This **enables you to create a tier architecture for your API**, like `Bronze`, `Silver`, `Gold` tiers.

* using **AWS Marketplace** you can **register your API** and effectively **monetize your API using usage plans**.

#### Canary (REST API only)

- you can setup **canary deployment in APIGW**

* to **get new deployment going**, **you have to either delete or promote canary**

### Load Balancers

- different types:
  - **classic load balancer: LEGACY!**
  - **application load balancer** for HTTP/HTTPS stuff
  - **network load balancer** for **connections that are NOT HTTP/HTTPS**

* allocated to **specific VPC** and **AZ inside that VPC**

- **can operate inside multiple AZs**

* you probably want to enable **cross zone balancing when balancing between multiple AZs**. **Without** this setting enabled, traffic is **distributed between the nodes in the AZ the load balancer resides**.

- can have security groups attached

* **can even invoke lambda functions ;o**

- **can perform health checks**

* there is a notion of **Target Groups (up to 1000 targets)**. This basically allows you to specify multiple EC2 instances without having to specify them explicitly.
  **In the context of EC2 Target Group usually points to a Auto Scaling Group**.

- **target groups can be containers, EC2 instances or IP addresses**

* load balancers **can balance** between **multiple target groups**

- load balancer have **listeners (up to 10)**. This **enables** you to **authenticate using social providers through those listeners**.

* **listeners** can have **roles**. This makes **content-based balancing possible**

- ALB/NLB enable you to create **self-healing architecture**. If you are using **ALB/NLB + ASG combo** your application becomes `self-healing` meaning that if one instance fails it gets replaced and such.

* there is a notion of **dynamic host port mapping**. This allows you to for example **run multiple ecs tasks on the same instance**. **When using** this feature **ECS will start containers with a random emphermal port exposed on the instance**. **ALB will take care of mapping between instance port and container port**.

#### ELB

- **ELB CANNOT BALANCE BETWEEN MULTIPLE SUBNETS IN THE SAME AZ**

* general term for **ALB or NLB or Classic Load**

- to **balance between AZs**, ELB creates **Load balancer nodes** within **each AZ**. What is important is that **% of traffic to each node is dependant on number of resources assigned to the ELB node**.

* when placed within a VPC **only nodes consume PRIVATE IP addresses**. The number of nodes depend on the amount of picked AZs **but the ELB itself DOES NOT have reserved PRIVATE ip address**

- you should **always refer to ELBs by FQDN**. That is because **there are multiple nodes of ELB (per AZ)** and that **DNS can correctly resolve to correct IP**.

* you **cannot route outbound traffic through ELB**.

- you are **charged based on running time and traffic**. That means that even if there are no instances which ELB manage, it is still incurring costs.

#### Client Affinity (Sticky sessions)

- you can enable **sticky sessions** on **ELB**. This feature makes it so that ALB tracks the client (for which sever it handed it off to). So it can keep **sending it back to the same server**.

* this is usually done by **setting a special cookie** or **tracing IP details (NLB)**

#### ALB

- work in **layer 7**. That means that they are **HTTP/HTTPS aware**. You can create rules based on url path and also based on the hostname itself (website1.site.com and website2.site.com).

* can have **multiple TSL/SSL Certs** using **SNI**

- **ALB can authenticate users using Cognito with social providers**

* **ALB can balance** between **different ports**. This is done by **specifying listeners rules**

- **exposes a DNS address**

* is **cheaper** than **CLB**

- great for **separating traffic based on their needs**

* **CAN** have **SecurityGroup attached**

- supports **gRPC**, **HTTP/2** and **IPv6**

#### NLB

- work in **layer 4**. They are **software based**. This is the reason behind the extreme performance.

* they **do not modify incoming network packets in any shape of form**, so you **do not have to use proxy-protocol**

- **can balance** on **UDP**. Remember that with **ALB there are 2 connections, one to ALB and one to target**. **NLB forwards the requests straight to instances**.

* **exposes DNS name**, but you **can assign static IP to it (EIP)**.

- **cannot** have **SecurityGroup attached to it**. Again, this is **because it's software based** and you are **assigning SGs to underlying network interfaces**.

* has **cross-zone load balancing disabled by default**.

- when you register instances VIA Instance ID, the underlying (incoming) IP address is preserved, in such case your application does not have to support x-forwarded-for header. But when you register your instances via IP, the underlying incoming IP address will be of the nlb nodes (private ip).

* **supports IPv6**

#### Gateway Load Balancer

- a very **simple passthrough**. Instead of calling a special URL, **you preserve your route table mappings**

* works on **layer 3**.

- from the **architectural perspective** is **similar to NLB + private link combo**

#### Access Logs

- if you **need more information** about the flow that goes through your load balancer you can use **access logs, DISABLED BY DEFAULT!**. Load balancer will **store those logs in s3 (sse-s3 by default)**. These allow you to get information about **individual requests** like IP address of the client, **latencies** etc.

* **access logs are not the same as error logs**.

- logs are **delivered on a time-interval**. You can pick **2 values: 5 mins, or 60 mins**.

#### CloudWatch

- **logs send to CloudWatch in 60sec interval IF there are reqests flowing through the load balancer**

* this is not the same as `access logs`. **Access logs contains details about a single request**, the CloudWatch metrics give you **broader view on the ALB as a whole**.

- the **`access logs` contain information about the latency**, but **if latency is your only concern** you be **much better of using CloudWatch native metrics for that**

#### Health Checks

- load balancer health check **can be carried on HTTP, SSL, HTTPS, TCP** protocols

* the **ELB type healthcheck** is often **preferred than the EC2 one**. This is due to the fact that `ELB` healthcheck actually checks the application and the `EC2` one just indicates that the instance did not crash.

- you can change healthcheck type to EC2 when you want to ensure that your instance will not be terminated whenever you restart given EC2 instance.

#### Scaling Events (mainly AZRebalance)

- this is where **ALB tries to rebalance spread of instances between AZs**.

* this event is quite important since **ALB can exceed, for a brief period of time, the maximum instances count within given group**. This is **due to the fact that new instances are launched first (without terminating the old ones)**.

#### Target Groups

- can be either **group of ec2 instances (usually accompanied with ASG)** or **lists of IPS**. This means that **by using IPS you can load balance with on-prem**.

* with target groups you can **load balance resources within VPC with resources on on-prem**. To make it work you have to use **private-ip only target group** and have **Direct Connect to on-prem**.

### DNS

- domain names are stored in multiple dbs

* there is **DNS resolver** which takes the query (domain-name) and resolve it to ip address.

- there are **Root Servers** which your device trust when trying to resolve given address to an ip.

* **naked domain**: **does not list format (www, ftp ...)**

- **DNS query can be recursive**. One name server will give you back another name server address and so on.

#### Records

- **A/AAAA**: maps a **host, so a www...** into **ipv4(A) or ipv6(AAAA)**

* **CNAME**: allows you to create **aliases (NOT THE SAME AS ALIAS RECORD)** to given **A/AAAA records**. **DOES NOT WORK ON NAKED DOMAINS (google.com vs www.google.com)**

- **Alias**: **ROUTE 53 specific!**. **Behave like CNAMES, instead of pointing to a A/AAAA/IP** it **points to a logical service provided by AWS**. You **can use** on **apex zone (naked) records**.

* **for some integrations** **alias record is FREE!**. You can use this technique to save costs. Remember that you can assign **www.something.com** as alias as well as the **naked domain (without www)**

- when creating an alias it might show as ** A (alias)** or something like this. Do not be afraid. This is what R53 is doing under thee hood.

### Route53

- **global service**. When creating blue/green in another region you do not have to re-create record sets.

* allows you to register domain (priced per domain name)

- there are **public** and **private** **hosted zones**. There is **monthly cost for creating hosted zones**.

* **public zone is created by default when you create or migrate a domain to route53**

- **public zone** is **accessed globally (public internet or VPC)**

* **private zone** has to be **created explicitly**

- **private zones** are **associated with a given VPC**

* there is a notion of **split view**. This basically means **creating the same names in both private and public zones**. **Inside a VPC, private zone always overrides public one**

- you are **charged monthly** for **hosted zones** and also for **resolver queries**.

* you can still use R53 even if your domain is registered within a 3rd party provider. Just update the 3rd party DNS records to use R53 NS records.

#### Route53 Health Checks

- aws has machines that perform health checks to route53 resources you provide.

* **route53 health check can check CloudWatch alarms**

- **route53 health checks can be used to check non-aws resources**

* the health check **CANNOT CHECK EC2 INSTANCES!**

- the health check **CAN CHECK OTHER HEALTH CHECKS**

* **fast: 10secs** or **default: 30 secs** interval

- before setting up health checks, **make sure your firewall settings, especially VPC SG, NACL allow for the requests**.

#### Route53 routing

- just remember that **TTL can bite YOUR ASS!**. **Whenever** you are **considering offloading the balancing to Route53** keep in mind the **TTL**. Since you **cannot attach ASG to Route53** there might be a **case where your instance is no longer there BUT TTL still routes to that IP**.

* you can **nest routing strategies**. This is to create complex routing infrastrucure

- an good **example** would be the need to create **latency based tree where leafs are based on weight**

##### Simple Routing Policy

- **single record** which routes to a **single resource**.

* do not mistake this with multivalue. This is for a single record, you cannot have multiple records with the same name with this routing strategy.

- that single record **can hold up to 8 IPs**.

* this is **NOT LOAD BALANCING**. **DNS request can be cached**. This will cause the **same IP to be hit**

- **no performance control**

- you **cannot attach healthchecks** to a simple routing policy.

##### Multivalue answer

- this routing strategy is for **multiple records (same name)** which routes to **multiple resources**. The **record have to be of the same type**.

* you can have **healthchecks** on **individual record**.

- **when queried** (DNS query) the routing strategy will **return up to 8 healthy records selected at random**.

* this **should not be used for load balancing**.

##### Failover

- allows you to **define additional records with the same name**

* **can be associated with a health check**

- there are **two record types**
  - **primary**: the main record, for example webserver running on EC2. You probably want to create health check for that record
  - **secondary**: the one that will take over if the health check on primary fails. This could be an s3 bucket for example. **Bucket name has to have the same name as the record name**

* if **primary record fails** traffic will be **resolved to secondary record**. The **name stays the same**

- you **can only create SINGLE record FOR PRIMARY and SECONDARY**

##### Weighted

- you **specify weight for a given record name, (you can have multiple records with the same name)**

* **route53** will **calculate the sum of the weights for the records with the same name**

- the **weights** are **used to calculate the percentage of the time these records are returned to the customer**

* **it does not replace load balancing, USING WEIGHTED AS A MEANS OF LOAD BALANCING IS AN ANTI PATTERN!**

- **used** for **testing new application features**

##### Latency

- **can have multiple records with the same name**

* you **specify given region** and **a resource that you want to route to if that region is selected**

- **records are selected per latency basis**

* **BASED PURELY ON NETWORK CONDITIONS, NOT GEOGRAPHY!**

- is based on the **latency between the user and AWS region NOT your service**.

##### Geolocation

- **DNS resolver now only returns the records that match the name AND location**

* this mean that you can **have a case** where **you specify a record but `nslookup` returns no results**. That is because **the location does not match**.

- **location** can be **a country , a region or default**. **Default matches everything**.

* **most specific location** is always **returned first**

- this means you can **present different application with the same DNS name, based on location**

* always **add a default rule**. This is to make sure **whenever DNS cannot pick the location** it will **fallback to the default rule**.

#### Route53 resolver

- **managed DNS resolver**. By default, your VPC comes with Amazon Provided DNS Server

* allows you to **perform DNS lookups across DX or VPN and the other way around (bi-directional)**

- this basically allows you to **create a hybrid cloud env.** where **some of your stuff is within VPC and some on-prem**. You might need to resolve DNS for some of the parts of your application to the stuff on-prem. Before all that you had to roll out your own dns resolver (probably within VPC) to be an intermediary between the AWS managed one and the one on premises.

* it is an **endpoint made of 1 or more ENI within a given VPC**. It **works for every VPC within a region (even if the VPC are not peered)**.

#### Inbound Endpoint

- this endpoint allows you to query **from your network / VPC to another VPC**

#### Outbound Endpoint

- this endpoint allows you to query **from your VPC to another network / VPC**

* allows you to **forward DNS queries to other resolvers**. This can be useful for **forwarding DNS queries to your on-prem DNS resolvers**.

#### Inbound & Outbound

- **bidirectional** queries

#### Transfering domains to R53

- there is a **fee** for transfering the domain.

* you cannot transfer every domain to R53. **There is a list of domains you can transfer**.

- you **can** still use **auto-renew** on **transferred domains**.

#### Logging

- you can enable access logs through `CloudWatch`. `R53` will create a log group and stream logs there.

#### Man in the middle attacks

- R53 **support DNSSEC only for domain registration**. It **DOES NOT SUPPORT DNSSEC FOR DNS service**. If you are worried about man in the middle attacks, use different provider.

* another option is to **run your own DNS server on EC2**.

#### Delegation sets

- when you create a hosted zone, **by default**, that hosted zone will get **different set of name servers assigned**. This can be **problematic when you have to manage multiple hosted zones**.

* you can create **reusable delegation sets (set of name servers)** which you can **use to create hosted zones with**. The **underlying name servers will always be the same (from the delegation set)**.

#### DDOS

- R53 is considered to be **DDOS resilient**.

### ECS (Elastic Container Service)

- allows you to run **docker containers on EC2 instances** in a **managed way**.

* those **EC2 container instances are launched inside a VPC**

- you can put **docker images** inside **elastic container registry (ECR)**

* ECR is integrated with IAM

- **ALB can balance container instances**

* **cluster** part comes from the fact that you will usually have **multiple ec2 instances running your containers**

- you can **force new deployment** if you pushed new image with the same tag, eg. `latest`.

#### X-Ray integration

- to monitor all of your services, you should **deploy X-Ray within a Docker container as ECS service**

#### CloudWatch Events integration

- ECS integrates as a **target** of **CloudWatch Events**

* you can change a lot of stuff when an event is invoked, **even the task count of the cluster**

#### Task Definition

- to run stuff you have to create **task definitions**. **Describes the task a.k.a service** you want to run. They are **logical group of containers running on your instance**

* with **tasks definitions** you can specify the **image, way of logging and resource caps**

- task definition also acts as sort of a bootstrap, **you can specify which command container should run when launched**

* you can **assign IAM role to the task**. This is the role that will be used by the underlying container.

##### Network Modes

- there are multiple network modes available: **none, bridge, awsvpc, host**. The **default is the bridge mode**.

* if you want to assign **private & public IP (ENI)** use **awsvpc mode**. This will also **enable you to assign SG to a specific task if needed**.

- **hosts map directly to the host networking**. That means that you cannot use this if you want to run multiple apache instances on the same host, since each instance needs PORT 80.

* **awsvpc** is the **only networking mode available for Fargate**. Be wary though, **instances have limits of ENI that could be attached to them**.

- **bridge** uses **Docker's built-in virtual network**, it maps the container ports to EC2 ENI directly.

#### Service Definition

- this is the **how will my infra look like which will be running containers** configuration.

* you specify **load balancers, task definitions (arn) and cluster (ec2 instances)** here.

- **each container instance belongs to only one cluster at given time**

#### Auto Scaling (aka Service Auto Scaling)

- it scales **services** which are **instances** which can have **task definitions deployed on them**.

* you can enable **ECS Auto Scaling**. It creates ASG automatically.

- **when creating** you can **specify subnet, VPC and any IAM roles** for a given instance.

* this works on the basis of `CloudWatch`. Note that `ECS` will **not automatically create ELB for you**.

#### Fargate

- should mainly be used for **transient workloads**. You **should not use `fargate` for any kind of databases**.

* **a bit more expensive than ECS itself since AWS is literally managing everything expect tasks for you**

So with **ECS you have to have EC2 instances running**. But with **Fargate you really only care about the containers themselves**. You can wave deploying ASG goodbye. **Fargate is basically container as a service, you only define tasks and that is it**.

- a good usage of fargate would be **load testing**, where you have **docker image of test automation framework** deployed.

#### Forcing new deployment

- this is needed to force the underlying deployment to use the latest pushed image

* you **might need to restart the ECS agent** when you use _Service Auto Scalling_

#### Security

- with **awsvpc** network mode you can **attach security groups to ENIs**.

* there is an **instance (ec2) role** that gives permissions to ECS agent to talk to ECS. This should be the first place where you look when ECS agents are not functioning correctly.

- you can also **attach IAM role to a specific task definition - task role**. This basically is a role attached to a specific container. Whenever you are dealing with **ECS - prefer task role rather than EC2 role**.

* you can also **attach task execution role**. This is the role **used by ECS itself to pull images, publish logs, basically do stuff in your behalf**.

#### Logging

- **make sure** that you have the **`awslogs` driver enabled**

* if you are **deploying on EC2, ensure that the instance can write to CloudWatch**

### EKS

- you can **pull images** from **various sources, not only ECR**

* remember that **EKS service role** needs **permissions to create AWS resources** (just like ECS)

- you will also need **a VPC and a security group for the cluster**. There is a template available within AWS docs.

* there is a **fargate launch type for EKS**.

### EC2 (Elastic Compute Cloud)

- resizable compute capacity in the cloud, **virtual machines in the cloud**

* Termination protection is turned off by default

- you can create a **bootstrap script**

* EC2 instance has **metadata**. There are a lot of useful information there.

- using **SPOT** instances there are different **behaviors you can configure** when your instance is about to get interrupted:

  - **stop**
  - **hibernate**

- **To get the metadata info CURL 169.254.169.254/latest/...**

  - `/userdata`: your bootstrap script etc
  - `/dynamic/instance-identity`: stuff about the instance -> IP, instance size, type all that stuff
  - `/meta-data/`: has **many options**, IP etc..

- remember that the **user data scripts are run only once (when instance is created)**. This can bite you whenever you have some scripts that require internet access and you forgot to configure the access properly. **There is a way to make sure user scripts run on every reboot**. This link can help:
  https://aws.amazon.com/premiumsupport/knowledge-center/execute-user-data-ec2/

* **stopping and starting an instance will MOST LIKELY result in data loss on instance store**. Unless you have dedicated tenancy model on that instance.

- **EC2 instance can only have ONE IAM ROLE**

* to specify **launch parameters of EC2** you can use **launch templates**. This allow you to specify **some configuration** for an EC2 like **security groups, AMI ids and such** .

- you **can add/change** existing **IAM instance role** on a **running instance**

#### Pricing models

- **on demand**: pay per time use (per-hour, per-second)

* **reserved**: capacity reservation, contract with AWS

- **spot**: like a market, but for instances. Instances come for the reserve pool of instances AWS currently is sitting on.

* **dedicated**: psychical machines **only for you**. Mainly used when you have strict licensing on software you are using

- there are also **Spot blocks**. This allows you to have an instance for a **specific X amount of time** using the spot pricing model.

#### Windows instances

- there is a notion of **EC2Config** on Windows instances

* this tool (included within a Windows AMI) is used eg. mapping EBS to correct drive letter.

#### Reservations

- sometimes there is not enough capacity for on-demand instances to be initialized within given region / AZ. This can be annoying whenever you are running ASG. Reservations enable you to make sure you can spin new instances whenever you want (by purchasing reservations).

- **zonal RI** guarantee **capacity reservation** within given az

* **regional RI DOES NOT guarantee capacity within region**. These should be used for **maximizing cost efficiency**.

* **on-demand capacity reservation**. This is useful for a **short term availability guarantee**.

- **scheduled reservations**: this is for tasks that are well, scheduled, daily monthly, whatever. You sign a contract for 1 year.
  capacity you can bid and buy, **when capacity is needed they will be taken away from you. There are mechanisms which will alert you if that happens!**

* you **cannot sell unused CONVERTABLE on the market!**.

#### Health checks

- There are different _health checks_:

  - **System Status Check**: this checks the underlying hypervisor (virtualization tool)

  - **Instance Status Check**: checking the EC2 instance itself

#### Tenancy

- **Tenancy model**. This is something **somewhat different than ec2 pricing models**.

  - **shared**: multiple costumers share the same piece of hardware (same rack, etc)
  - **dedicated**: hardware your EC2 runs on is only yours, but you have to pay more. No other instance from other customer run on this host. **Underlying hardware can change when you stop/start an instance**
  - **dedicated host**: this is basically like **dedicated but extra**. **Even if you stop/start your instance the underlying hardware stays the same**. You are not jumping between racks. You can pick on which underlying hardware you want to run!

#### Hibernation

- to enable hibernation your instance **must based on HVM AMI**.

* you can hibernate an instance. The **hibernation process is moving RAM data to EBS**. There are requirements for this to work
  1. Your instance **cannot be a part of ASG or ECS**. You can always suspend ASG or move given instance into maintance mode.
  2. Your **root volume has to be EBS**
  3. The **EBS volume has to be large enough**

- you **cannot enable hibernation on a existing instance**. You have to tick a box during the creation.

#### Instance Profiles / Roles

- you **cannot directly connect IAM role with an instance**

* to do that **you create instance profile** which **acts like a container for IAM roles for a given instance**

- **when** you **deploy EC2 instance** the **creation of instance profile IS HIDDEN from you**. It **happens behind the scenes**.

#### Fleets

- allows you to **mix and match instance types and sizes**.

* it's **amazon who launches your instances**

> With a single API call, now you can provision capacity across EC2 instance types and across purchase models to achieve desired scale, performance and cost.

- when using SPOT instances you should diversify across instance pools to avoid widespread interruptions. This makes it so that even when there is a demand on particular instance type your SPOT fleet will still live on.

#### AMI

- **WHEN CREATING ON EBS-based EC2** **automatically** creates **EBS snapshots**.

* can be **private or public**

- you can **whitelist aws accounts when the AMI is private**. This allows you to share the AMI between accounts without making the AMI public.

* ami **does not contain instance type information**. Remember that the AMI also contains the AMI permissions.

- you can have **an AMI based off instance store**. Instead of creating EBS snapshots, you have **files on s3 that gets referenced**.

#### Traffic Mirroring

- you can **mirror packets from one instance to another**. That another instance **can monitor the packets**.

* it **captures real packets**. **Not like Flow Logs**.

#### Key pairs

- based on **public/private key** cryptography

* when **migrating AMIs** you **do not have to import any kind of key**. You can **still use your downloaded `key-pair`** since it's the private part of the key.

- if you want to **completely wipe the key from the instance** you should just **terminate that instance**. There is a CLI command to delete the key, but **the CLI command does not delete the key from the instance itself**.

#### Auto recovery

- you can create **CloudWatch alarm** on `SstatusCheckFailed_System` metric and choose **Recover this instance** action.

* the auto recovery will **preserve your instance id, IP Address, EIP, EBS attachments, all that stuff**.

#### Networking (IPs and DNS)

- when **restarting (stop / start)** **private DNS / IPs stay the same**. The **private DNS** has a shape of **ip-X-X-X-X.ec2.internal**

* when **restarting without elastic IP assigned** **public IPs / DNS is reassigned (dynamic properties)**

#### Auto Scaling Groups

- launching EC2 based on criteria as a service

* you can launch underlying instances using 2 types of "templates", **launch templates** and **launch configurations**:
  - both are immutable, but you can create **versions of launch template**
  - **launch template is newer and recommmended by AWS**. You should always **prefer launch template as it has more options**.

- controls scaling, where instances are launched, etc

* **can work with instances in multiple subnets**

- instances inside auto scaling group can be monitored as an one entity. By default _Cloud Watch_ is monitoring every instance individually

- there are multiple versions of **ASG monitoring**
  - **basic**: **5 minute granularity**, by default enabled by **creating ASG from a launch template or from the console**
  - **detailed**: **1 minute granularity, cost additional**. By default enabled by **creating ASG by launch configuration created by CLI or by SDK**

* you can control the number of instances by manipulating three metrics:
  - **Desired Capacity**: this is the number **ASG will try to maintain**
  - **Min**
  - **Max**

- there is a notion of **scaling policies**. These are **rules, things you want to happen when something regarding EC2 instances happen**, eg.

* there are **multiple types of scaling policies: step, simple, custom and target tracking**

- with **target tracking policy (also sometimes refereed as dynamic scaling)** you **specify a scaling metric and a target value**. Think **scaling policy to keep the average CPU utilization of ASG at 40% or something like that**. It **uses CloudWatch alarms and metrics to take actions**.

- the deal with **step scaling policies** is that you can create **multiple predicates for a given policy**. For example: **add 2 capacity unit when CPU hits 80% , add 2 capacity units when CPU hits 60%**.

* **REMEMBER THAT YOU HAVE TO LIST AZs YOU WANT YOU INSTANCES TO BE DEPLOYED INTO!!**

- there is a notion of **health check grace period**. This is the **time** it takes to **spin up new instance**. This time of course is dependant on **bootstrap script** and **how much application code is in the AMI**.

* **individual instances can be protected from scale events**. This is useful when you have a master node that cannot be terminated.

- there is notion of **connection draining**. This is **done by ALB OR NLB**. When **terminating an instance ALB/NLB will wait for that period to pass**.

* **its the ASG who terminates things**. When a scaling even occurs **ELB simply stops sending traffic to that instance**. **ASG does not `listen` for ELB commands to terminate stuff (BY DEFAULT)**. This can be changed. **ASG CAN listen to ELB health checks if you set it up**. But always remember that it is the ASG that terminates stuff by default

- you can have **custom CloudWatch metrics trigger scale events**

* you can **suspend Auto Scalling**. This is **useful** while **debugging**.

- if you want to make requests through API (regarding Auto Scaling), you have to sign them using HMACK-SHA1

##### Updates to the ASG

- you can either use _AutoScalingReplacingUpdate_ or _AutoScalingRollingUpdate_

* use **ReplacingUpdate** with **`willReplace: true`**. This deployment option works very **similar to the `immutable` option from EB**

- there are more policies but these 2 are the most important

##### Lifecycle hooks

- ASG exposes **lifecycle hooks**.

* this allows your instance to **pause** before **initialization / termination**.

- **when is paused** state you can **do your stuff** like **perform a backup or save some data**.
  This can **be really useful** when you want to **do some stuff before your instance gets terminated**.

* this is usually used when your instance is stateful.

- you should probably manually complete given lifecycle. Otherwise it will wait the maximum period.

* lifecycle hooks should be preferred way of adding any kind of dynamic elements eg. ENI

##### Termination policies

- which instance should be terminated and why.

* you can have **multiple termination policies (even define all of them)**. There is an **order to which these policies are applied**. The order is **dictated by how you define your policies**.

- there are multiple termination policies, most of them are self explanatory:
  - `OldestInstance`
  - `NewestInstance`
  - `OldestLaunchConfiguration`
  - `OldestLaunchTemplate`: remember that **you should favour launch templates**.
  - `ClosestToNextInstanceHour`
  - `Default`: remmeber that **default will always apply first to AZ with the most instances**

###### Default termination policy

- the `default` termination policy is will perform other termination policies based on given criteria. There is a flowchart for that

* the steps are
  - `oldestLaunchConfiguration`
  - `ClosestToNextInstanceHour`
  - random instance

##### Instance Protection

- no matter what happens, given instance will not be terminated.

* can be applied to asg or individual instance

- can be terminated manually.

* **instance protection starts when instance is in `InService` state**.

- **instance protection** is **not the same** as **termination protection**. You can still
  - **terminate it manually**
  - use `terminate-instancees` comand
  - you asg can also run `TerminateInstances` action.

* to make sure that your instance **will not be terminated event if marked as unhealthy** use **termination protection**.

##### Suspend processes

- this allow you to suspend asg processes

* useful for where you want to resize your instance - this would require termination, thus asg would most likely spin a new instance.

- you can **even suspend health checks**.

* another use case might be **suspening the ASG during the deployment**. This is due to the fact that when deployment and scale event is underway, the newly added instances (due to scaling) will have the latest **deployed** version of the app running. Often you will end up with 2 application versions in your ASG.

#### Migration

- remember that **security groups are regional**

* you can **copy a security group within a given region**

- you can **export security group config** to a **new region VIA CLI!**

#### Troubleshooting Rolling update

- rolling update settings are set using the `Update Policy` setting. If those are not set correctly, you might encounter unexcepted results

* there are 3 things you might want to do
  - Configure `WaitOnResourceSignals` and `PauseTime` to avoid problems with success signals
  - Configure `MinSuccessfulInstancesPercent` to avoid stack rollback
  - Configure `SuspendProcesses`

#### EBS (Elastic Block Store)

- basically **virtual harddisk in the cloud**

* remember that they are **tied to specific AZ**

- persistent storage

* **automatically replicated** within it's own, **single AZ**. This is on the contrary to EFS and so taking snapshots is highly recommended.

- Different versions:

  - **Provisioned IOPS** - the most io operations you can get (databases), most expensive. Recommended **when you need more than 16k IOPS**
  - **Cold HDD** - lowest cost, less frequently accessed workloads (file servers, databases).
  - **EBS Magnetic** - previous generation HDD, infrequent access
  - **Throughput Optimized HDD** - streaming, big data, log processing. All that stuff also taking costs into consideration.
  - **General Purpose**

So the costs are **usually** **Provisioned > General Purpose > Throughput Optimized HDD > Cold HDD**.

- **I/O grows proportionally to the size of GP ssd**

* **Cold HDD cannot** be used as **boot volume**

- you **can modify the size of an existing mounted volume**.

* you can take **EBS snapshots**

- **during** the operation of **creating a snapshot** you can **use** your volume **normally**

#### Max I/O

This is quite important to know

- **Provisioned IOPS**: up to **64,000 I/O** That is up to **1,000 MiB/s**

- **General Purpose**: up to **16,000 I/O** That is up to **250 MiB/s per volume**

- **Throughput optimized HDD**: up to **500 I/O** That is up to **500 MiB/s per volume**

- **Cold HDD**: up to **250 I/O** That is up to **250 MiB/s per volume**

#### Snapshots

- **snapshots are incremental**. AWS only diffs whats changed **you can delete every, BUT NOT THE LAST SNAPSHOT** if you want to make sure your data is secure.

> If you make periodic snapshots of a volume, the snapshots are incremental. This means that only the blocks on the device that have changed after your last snapshot are saved in the new snapshot. Even though snapshots are saved incrementally, the snapshot deletion process is designed so that you need to retain only the most recent snapshot in order to restore the volume.

- **snapshots** are created **in a given availability zone**. They are **replicated across multiple AZs automatically**. When you want to **restore from snapshot in a different AZ** you have to **copy the existing snapshot to a different AZ**. **When creating a snapshot** you can **pick the `root AZ per say**.

* **EBS snapshots are automatically stored on S3**. These are **inaccessible to you inside S3 console, you can see then through EBS menu in EC2**.

- the **volumes themselves NOT SNAPSHOTS** are **replicated within SINGLE AZ**. If that AZ goes down, the volume is not available.

* you can **create snapshots** either **manually**, **using LifeCycle Manager** or through **CloudWatch jobs**.

- **retention policy DOES NOT carry over when copying.**

* you **should not take use RAID and take snapshots**. This is due to the fact that snapshots are volume based. If multiple volumes are in-sync (by using RAID configuration) the snapshots would most likely be corrupted or have stale data.

#### LifeCycle Manager for EBS

- creating snapshots manually is ok but AWS can take care of this task for you. With `LifeCycle Manager` you can enable creation of automated backups. BUT **YOUR VOLUME HAS TO BE TAGGED**

* this **can be done across different DBS** like SQL or Oracle or other.

- the **frequency can be UP TO 24hrs**. You can pick from the list of available frequencies **but the longest one only 24hrs**.

* your volumes can have **multiple tags**. This means that **all policies specified for given tags will be in effect**. That means that you can have a volume where backups are done every 12hrs and every 24hrs.

- underneath, it **probably uses `CloudWatch Events`** since they allow you to **directly call the EBS API**.

#### Termination

- **by default root EBS volume will be deleted when instance terminates**

* **by default ATTACHED EBS volume will NOT be deleted when instance terminates**

#### Encryption

- you need **specific instance type** to be able to **enable encryption of volumes**. **Not all** instances **support encrypted volumes**.

* **EBS root volume CAN be encrypted**

- **encryption** works **at rest** and **in transit between the volume and the instance**

* **volumes from encrypted snapshots are also encrypted**

- you **cannot remove encryption** from **an encrypted volume**.

* you can only **change the encryption state** while **copying unencrypted snapshot of an unencrypted snapshot**. Note the **COPY** word. This is very important.

- you **can change the KMS logistic (keys)** only when **copying encrypted snapshot of an encrypted snapshot**.

#### EBS Optimized

- this setting is within **advanced settings** and basically makes it so that **you volume has more IO available to it**

* it creates a dedicated connection between your instance and the EBS volume. Please keep in mind that you cannot scale infinitely with Raid 0. You will probably hit a throughput limit on your instance.

#### EBS vs Instance Store

- Instance Store a.k.a **Ephemeral storage**. The data lives on a rack where your virtual machine is booted. **If you reboot it, you WILL NOT loose the data**.

* when you **Stop/Start an instance** **instance store data will be lost**.

- Instance Store is not really persistent, whereas EBS is a persistent, multi AZ storage option.

* you **cannot create snapshots from instance store**

- you **cannot create an AMI directly**. With **instance store** you can only create ami by:
  - **creating bundle from root volume (using EC2 cli)**
  - **uploading** that created bundle **to s3**
  - then that **bundle is registered as AMI (from s3)**

#### Restoring from an EBS

When restoring from an EBS volume, **new volume will not immediately have maximum performance**. This is due to the fact that **not all data is copied instantly to a new volume**. The data is **copied lazily, when you attempt to read from a given resource**. This is why **sometimes, sys admins perform recursive lookup of all files on the volume, this will 'prime' them for real read operation**.

#### Changing The volume

- **root** volume **can be changed without stopping the instance**. When it comes to which volume - it depends on the original volume itself.

#### Raid Configurations

- use **RAID 0 for maximum performance**. Remember that with **RAID 0 preserves all capacity (n \* size of drives) but when you loose 1 drive you loose all the data**.

* use **RAID 1 for maximum fault tolerance**. This is where your **data is mirrored between 2 volumes**

- **DO NOT use RAID 5, 6**

#### Attaching

- you **can only** attach **volumes with instances that are in the same AZ**

* you **can have volumes attached to multiple instance**

#### EFS

- you **only pay for what you use**, but the EFS in itself is **3 times more expensive than EBS** and **MUCH MORE (20x) expensive than s3**

* **E**lastic **F**ile **S**ystem (EFS)

- **Similar to EBS**, but there is one **BIG DIFFERENCE**. EFS instance can be used by multiple EC2 instances, EBS volume can only be used by one EC2 instance.

* think of it as multiple EC2 instances having the same disk

- **automatically scales storage capacity**, when deleting shrinks, when adding resizes.

* just like s3 there are different tiers:

  - infrequent access

* **can be accessed between multiple AZs**

- you **cannot** point **route53 alias to EFS**

* can be accessed by instances **spread across multiple VPCs using VPC peering**

* the data is **distrubited accross multiple AZs**

##### Security

- **mount targets can have security groups associated with them**

* you **cannot** restrict access **through ACL or IAM roles**. **USE SECURITY GROUPS** and **attach them to mount targets**.

##### Compatibility

- EFS is **only supported on Linux instances**

##### Backups

- EFS operates on the notion of **backups not snapshots**

* you can create backups using **AWS Backup service**. This service allows you to create **incremental EFS backups**.

But most important information, remember **there are no so called snapshots when it comes to EFS**.

##### Mount Points / Targets

- there is a notion of **mount targets**. This is the think that **allows multiple EC2 instances to share the same storage**. It **lives inside a subnet**.

* this feature allows you to use the same **efs volume through different AZs**

##### Mounting

- **EFS actually gets a DNS name**

* you can mount it by using **DNS resolution** but there are some **additional requirements**.

- remember that during mounting you also control the encryption (in-transit).

* because EFS gets a DNS name, you sometimes may hear a term: _mounting using FQDN_. This is nothing but just mounting using this DNS name.

##### DataSync

- AWS says that **in theory you can mount EFS on prem** but that **requires really good connection**. Moreover the **traffic** itself **is not encrypted** so **VPN or Direct Connect is recommended**.

* what you can do is to use **AWS DataSync**. This will **sync your on premise env with the EFS volume which you use within the AWS env.**.

![img](./assets/efs-datasync.png)

##### Encryption

- **by default** data is **not encrypted in transit /rest**.

* to enable **encryption in transit** add a **special flag** when **mounting a volume**. This flag is **-o tls**. It uses TLS so no CMK required.

- to eanble **encryption at rest** use **Amazon EFS mount helper**. This **can only be done during mounting**. So if you have an **existing volume**, you would need to **unmount it, specify the setting and mount it back again**

* encryption **at rest has to be enabled before creating the file system**

- **mount helper** is something that you have to **download before you can use it**. It is a **cli tool**.

* you **do not have to have special SG rule for encrypted EFS**

##### Performance

- there are **2 main performance modes**
  - **General Purpose**
  - **MAX I/** - this one is the faster one but there is a **tradeoff of higher latencies for file metadata operations**

##### Throughput

- there are **2 main throughput modes**
  - **busting mode** uses **credits** (recommended by AWS)
  - **provisioned mode**: this is where you have **contant high throughput (not spikes)**

* with **bursting mode** the **throughput of the storage grows with the storage size**. This usually is ok, but you might get throttled on unexpected spikes when the volume size is low.

- with **provisioned mode**, the **throughput is independent of storage size**.

#### Placement Groups

Way of grouping EC2 instances.

##### Clustered

- EC2 are **very close to each other, single AZ**

* you should try to **provision all your instances up front**. There is a limited space, and there might be a problem with adding an instance later on.

- you **should consider** using **enhanced networking** so that you take full advantage of close proximity of instances. This is done to **speed up communication between the instances**, nothing more.

* you should consider this option for **small deployments**

##### Spread

- **can span multiple AZs**

* **max 7 instances** running **per group per AZ**

- instances **spread across underlying hardware**

* you should consider this option for **medium / small deployments**

##### Partition

- **groups of clustered EC2 on separate racks**

* you **cannot use dedicated host** with this option

- you should consider this option for **large deployments**

* ideal for **HDFS, HBase clusters**

- **can be MultiAZ, but has to be the same region**

#### Enhanced Networking

- mean of **optimizing the network interface** rather than the volume itself.

* uses **SR-IOV to squieze maximum I/O NETWORKING performance**.

- EC2 **must** be launched from **HVM AMI**

* EC2 **must** be launched **inside VPC**

- there is **no additional charge**

* available **only on specific instances**

#### ENI (Elastic Network Interface)

- **by default, eth0 is created**

* has **allocated IP address** from the range of subnet. That **allocated IP is A PRIVATE ONE!**

- has **interface ID**

* can have ip addresses changed (multiple private addresses), **the number of ip pools are dependant on EC2 instance size / type**

- EC2 **can have multiple ENI's**. When instance is terminated **ONLY default eth0 is deleted BY DEFAULT**.

* ENI can be **attached** in **different stage of life-cycle of EC2**.
  - **hot attached**: when instance is running
  - **warm attached**: when instance is stopped
  - **cold attached**: when instance is launched
* with **mutliple ENI attached** you can put **multiple SGs on one instance**. That is because you put SG on the ENI and not on the EC2 itself.

### AWS Batch

- allows you to run **processes (async) across one or more instances or ETL jobs**

* **aws will take care of scaling**

- think of **bash scripts or other jobs**

* **jobs** as **units of work (shell scripts, Linux exec, Docker container image)**

* **one job might depend on another**. With AWS Batch you **can make sure that jobs are done in a right order**.

- there is a notion of **job definitions (how jobs are to be run)**

- **you can use SPOT EC2 instances** for maximum cost efficiency.

* there is **no aggregate step at the end**. This is **different than EMR**.

- your job might be **stuck in RUNNABLE state** when **the job does not have adequate permissions** or **you have reached EC2 limit** or **the jobs asked for more resources than the environment can provide**

### SWF (Simple Workflow Service)

- **manage workflow state, similarly to Step Functions**, but this is **not serverless, code is run on ec2**.

* **you** as a user **have to manage the infrastructure that runs the workflow logic and tasks**

- **kinda retired by AWS**.

### Elastic Beanstalk

- **platform as a service**

* aws will literally do everything for you, but **you have full control of underlying resources created**

- will automatically scale your app

* **no additional charge, you only pay for the resources which were provisioned**

- usually used for **web application deployments (with RDS eg.)** or **capacity provisioning and load balancing of a website**

* supports **multiple environments (dev, prod ...)**

- **when you delete an application all the resources associated with it are gone too. That also applies to the databases!**.
  You can make sure the data is still there by **creating DB snapshot** or **creating any parts of your application that you do not want to accidentally delete OUTSIDE Elastic Beanstalk env**.

#### Deployments

- you **cannot deploy on the premise!**

* you can have a **custom AMI** as your blueprint for the application.

* there are several deployment options

- **all at once**: **default** configuration, every instance is affected at once.

* **rolling**: _Elastic Beanstalk_ splits instances into batches and **deploys a version into them, batch by batch**.
  You may choose to add additional instances before the deployment itself and that would be called **rolling with additional batch**.
  Be aware that with this deployment option, you will have a situation where clients are switching between application versions (Load balancer routing)

- **immutable**: _Elastic Beanstalk_ launches **full set of new instances running the new version of the application in a SEPARATE (temporary ASG)**.
  This strategy is quite nice since clients **will not switch between different application versions**.

* **blue/green**: _Elastic Beanstalk_ launches **new version to a separate env. and then switches the DNS**. Before the introduction of per minute/hour billing this method was considered to be not very practical.

#### Other deployment types

- **minimum in-service deployment**: this one is **similar to rolling**, but the orchestration service is greedy and is always **trying to deploy to as many instances as possible** while **keeping the minimum healthy** (**rolling has set batches**).

* **canary**: **similar to blue-green** but **instead of being binary (either blue or green)** you **split traffic between green and blue**.
  Usually you shift this traffic using **weighted routing or lambda aliases**.

#### Updates

- similarly to deployments there are few options here:

* **in-place update** is **performing updates on live instances**

- **disposable** is performing a **rolling update** where **new updates are brough up by terminating old ones (batches)**.

#### Deleting a stack

- while deleting a stack you might face a problem where **SG created by `Beanstalk` is used as a dependency in other SG**

* this will result in a failure when trying to delete the environment

#### Unsupported Platforms

Sometimes it can happen that your runtime is not supported by ElasticBeanstalk by default. But you can still use it!

- you can create **custom AMI** which will be **used to run your application**

* you can create **custom a custom platform**. This is much more involving. It uses **Packer instead of Docker**.

#### Environments

- there are multiple pre-defined environments you can use.

* for `java` you can specify custom java parameters through `Environment properties`

##### Worker Environment

- this is where you have SQS queue, and your instances are listening on `localhost`.

* you can create a special `cron.yml` file which **sends POST requests to `localhost` on a period**.

##### Configuration files

- you should use `Dockerrun.aws.json` for all things related to docker

* you should use `ebextensions` for any other configuration

### Monitoring Services

#### CloudWatch

- **monitoring service** for applications, services, **monitors performance / what is happening WITH resources**

* **by default** CloudWatch **pushes logs every 5 mins**. You can **enable Detailed Monitoring** which will enable **logging in 1 min intervals** but that solution **is a paid feature, per instance**

- the more frequently you publish metrics data the less they stay at that resolution. CloudWatch will after some time aggregate points to a metric with lower resolution eg. 60 secs => 5 mins and so on.

* **namespace groups related metrics**

##### Metrics and dashboards

- You can create **custom CloudWatch metrics**. This will enable you to **eg. monitor RAM on EC2**. With **custom CloudWatch metrics** you can setup **High Resolution metrics** which **will log up to 1 second intervals (the fastest possible)**.

* with **High Res metrics** alarms **will trigger in 10 secs intervals**. Image it triggering every second 😂

- you can create dashboards from metrics. These **dashboards can be shared between accounts and / or within an organization**.

* you **can export dashboards to s3 (the data points)**. This is done **by using cli `get-metric-statistics` command / lambda + cloudwatch rule**.

- **dashboards** are used to **correlate data**. This is just a **fancy way of saying having multiple graphs near each other (as widgets)**.

###### Metrics math

- instead of simply using _count_ and _threshold_ for your alerting, you can create expressions for better alerts

* think creating an alert that _X out of Y request failed within given time window_, not _just X request failed in a given time window_

- while a bit more complex I think these should be favourited, it will make sure that you are not overwhelmed with incidents.

##### Insigths

- there are many flavours of `CloudWatch insights`, there are **Application**, **Container** and just **Insights**

* the most important one being the `CloudWatch Insights`. This is a tool which allows you to filter the logs with expressions and other stuff.

##### Billing dashboard

- remember that **billing metrics are only available within `north virginia (eu-west-1) region`**.

* you can, just like with other metrics, create automatic dashboards.

##### Alarams

- they **react to given metric**

* alarms can **take actions**, like **trigger ASG action / send SNS notification**.

###### Billing Alarms

- you can switch to `n.viriginia` to get to billing dashboards

* you can create **normal, cloudwatch alarams on those billing dashboars**. You can do it per service or overall.

###### Monitoring spend

- you have to enable **Billing ALERTS not Alarams** for you to be able to **track estimated AWS charges**

* there is a **delay between incurring a charge and a AWS Budget notifications**

##### Logs

- **log group is a container for a log streams** which have the same specifications (retention etc...)

- you can **subscribe to a log group using lambda / ES**.

* **log stream is a representation of logs for a single thing**, like ENI (with Flow Logs) or particular EC2.

- you can **filter logs streams** to **create metrics on found items (like log which contains error string)**

* can monitor:

  - **CPU**
  - **Network**
  - **Disk**
  - **Status check**

- for **non standard metrics like RAM usage** you can install **CloudWatch unified log agent** to push those to custom metric.

* you would use **CloudWatch Logs** for creating **application-level alarms**. This could be number of errors and such.

- you can **stream logs to lambda or Elasticsearch (uses lambda underneath)**.

* you can **export logs directly to S3**. You have to remember though about permissions (`getBucketAcl`). Otherwise you will not be able to export the data, even though your bucket might be public.

- the **s3 export is one time operation**. For **real time, use `Kinesis Data streams` or `Lambda`** for **near-real time (buffered) use `Kinesis Firehose`**.

###### Embedded metrics format

- instead of using the sdk to synchronously put metrics data, just log `JSON` data in specified format.

* EMF will handle up to 100 metrics per `JSON` blob.

- use `properties` for data which has high cardinatlity. That means that there are a lot of potential unique `name/value` pairs. Remember that the `CloudWatch` can only store 10 dimmensions per metric.

- use `CloudWatch` insights for searching through the `properties` that you set using EMF.

###### Log retention

- **by default** CloudWatch keeps the logs forever. **But after 15 months you cannot access them through console**.
  This means that if you **want to access your logs after **that period of time, you have to **use an API to retrieve the datapoints**

###### Cross account log data sharing

- you can **send logs to a destination in another account**

* that **destination can only be Kinesis data stream**

- very useful for **creating cross account logs harvesting**

##### Unified log agent

- this is the **new tool for pushing LOGS AND METRICS to `CloudWatch logs`**.

* previously with `CloudWatch agent` you had 2 separate scripts to do what the `unified agent` is doing at the moment`.

- **config file**, by default, **is pushed to SSM**.

* the unified agent is also **more performant**.

##### Events

- there is notion of **events**, which basically provides **near instant stream of system events**

* **events** have to do more with **what is happening on the instance (resource) level**. This could be **EC2 instance changed the state from pending to running** and such.

- you can configure **CloudWatch Events rules**. These allow you to run **scheduled jobs (cron)** and **invoke different targets (services like lambda or step functions (and many others))**.

##### Rules

- remember that you can **create rules NOT EVENTS! to schedule some actions**

* you can even **make some actions based on events comming from different services**.

- you can create a rule which has multiple targets.

* you **cannot use `CloudWatch Alarms` as source**.

- there are multiple predefined events but you can also **integrate with `CloudTrail`** to have access to **every event except `List`, `Get` or `Describe`**.

##### Dashboards

- you can create custom dashboards with given metrics from CloudWatch

* these **dashboards can be cross accounts and cross region**.
  - if you want to make it work **cross account** you have to enable **sharing on each account that will make data available on the monitoring account**
  - if you want to make it work **cross region with a single account** you **do not have to do nothing**. It's **already built-in**.

#### CloudTrail

- **monitors AWS API calls**

* different events can be logged:
  - **data events**: **resource operations** performed **on or within the resource**
  - **management events**:

- **enabled by default for all new accounts**.

* **logs** can be **stored on S3 or can be pushed to CloudWatch**. Log files which are **stored on s3, are using SSE-S3 by default!**.

- **event history** persists **up to 90 days**.

* you can create **one trial** and **apply it to multiple regions**.

- you **can trigger CloudWatch events usng CloudTrail events**. Pretty neat!.

* event are not delivered in real-time. There can be **up to 15mins delay**.

##### Log integrity

- you can check the integrity of the log files using **cli**.

* you need **`digest files`** to do so. These are **delivered every hour to your s3 bucket**.

##### Data Events

- by **default** CloudTrail **does not log S3 put/get events** and **lambda invocations**

* you can **enable data events** to **have those events logged**.

- logging **can be enabled per bucket and per function**

##### Global events

- remember that **by default** CloudTrail **does not _listen_ to global events**, like IAM actions.

* make sure to enable capturing global actions as well.

##### CloudTrail and CloudWatch

- there is no automated way to have your logs delivered to `CloudTrail`.

* you have to **create a lambda that listens to s3 events**. That lambda will **log to cloudwatch while reading CloudTrail files**.

#### Flow Logs

- monitors **metadata of the IP traffic to/from network interfaces within a VPC**

* create **log stream for EVERY interface monitored**

- **logs can** be **stored on S3 or CloudWatch**

* **IT HAS TO DO WITH METADATA NOT THE ACTUAL CONTENTS OF THE IP**

- by default **it exposes the ip of the person who is connecting**

* you filter which traffic you want to log. It can be either rejected, accepted or both.

#### Integration with EB

- you can make `CloudTrail` send events to EventBridge, do achieve that you will need to setup a _trail_ in the region where you want to listen the events to.

* for some services, this setup might not be necessary as they might already send events to both `CloudTrail` and `EventBridge`

#### With `CloudWatch`

You can send `CloudTrail` logs to `CloudWatch`. This might be a powerful combination since you could create alarms for given metrics and have like a _pseudo_ analytics workflow where you are notified when something happens.

### Analytics

#### Athena

- **completely serverless product**

* interactive **query service**

- data lives on s3, it **never changes**

* you can query data without doing any kind of ETL. You are operating on a raw data that lives on s3.

* **schema on read** a.k.a you define what the data YOU would like to look like.

- **schema is not persistent**, schema is only used when you read data (perform queries)

* reduces admin overhead, **no INITIAL data manipulation**

- you only pay for the data you query and the storage on s3 (your source data probably already exist on s3)

* think of **schema as a lens to look through at data**

- mostly used for **ad-hoc queries**, because you only pay for what you use. **It's designed for large-scale queries**

* **query results performed by Athena area automatically written to s3**

- Athena **can read data and write query results given encrypted data by KMS**

* can **query data from CloudWatch, ELB logs, CloudTrail logs, Flow logs**. The native support is for s3, but there are
  **data connectors** that allow you to run queries on other sources. The **data connectors** use AWS Lambda under the hood to run the queries.

- is very **flexible** when it comes to **data, which it needs to process while querying**. It can be **structured / semi-structured / unstructured**

- **if you really care about performance** consider **transforming the s3 objects** to **parquet format**. This is a special format that makes the querying much faster.

##### Integration with Glue

- you can create the _table_ (so the **lens which you will look through**) **manually** or **using AWS Glue**

##### Athena Vs Redshift Spectrum

Both of these tools can be used for DataLake querying, but, and that is very important, **you would use Athena when your data mostly lives on s3 and you DO NOT have to perform joins with other data sources**. This is a **complete opposite of RedshiftSpectrum** where you would **use Redshift Spectrum where you want to join s3 data with existing Redshift tables or create union products**.

##### Athena Vs EMR

**Athena is only for querying** it does **not transform the data**. **EMR has capabilities of transforming the data**.

#### EMR (Elastic Map Reduce)

- splits data into **splits**. This is the **mapping part of the EMR**. Then **nodes process the splitis**.

* allows you to perform **data-processing on large-scaled, semi-structured or unstructured data**

- think **big data** data sets

* there are **nodes**, the **master node splits the work between nodes**

- uses **shared file system through nodes**

* It can use 2 types of storage to perform operations:
  - HDFS: s3 is used to read and write the final data to **(supports both in-transit and and rest encryption)**
  - EMR FS: s3 is used as primary data store to carry out all operations

- **HDFS literally can contain any kind of file format**.

* **use for on-demand, ad-hoc, short-term tasks**

- **master node can** be **sshed into**. You can **ssh into master node** and **use Hive to perform SQL like queryies on the data**.

* usually has to do with **Spark** jobs.

- **can manipulate the data**

* cluster runs on **EC2 that are inside VPC**.

- **master node => core node => task node**

* if **master node dies = everything is RIP**

- if **core node dies** you are facing **potenal data loss**.

* **master node cannot be changed**. If you want **to change underlying instance, you would need to create new cluster**.

- you should **prefer creating Cluster in the same region as the data you will be retrieving / storing**

* **YOU have to provide an application for mapping and reducing**

* **nodes** can be **monitored inside CloudWatch**

#### AWS Data Pipeline

- **serverless product**. It allows for **orchestration within ETL but mainly for moving data**

* allows you to easily **transform (ETL) and move data**

- can be set to **run on given schedule**

* **can integrate with on-premise resources**

- **does not work with streaming data such as Kinesis**

* **integrates with EMR**

- basically it allows you to **simplify ETL jobs**, but if the thing that you want to do is not supported it **may seem limited**.

* **job is relaying on EC2, AWS GLUE does not have this limitation**. It manages the lifecycle of EC2

- in terms of overhead and managment it is like **Elastic Beanstalk of ETL services**.

* one common pattern is to transport data from dynamoDB to other medium, or vice-versa.

- you can also perform **cross region dynamoDB copy**. You would only use this if you do not use global tables.

#### Kinesis

- **fully managed**

* ingest big amounts of data in real-time.

- you put data into a stream, **that stream contains storage with 24h expiry window, WHICH CAN BE EXTENDED TO 7 DAYS or 365 DAYS**. That means when the data record reaches that 24h window it gets removed. Before that window you can read it, it will not get removed. This is why Kinesis is know for the **data immutability**.

* stream can scale almost indefinitely, using **kinesis shards**

- **NOT A QUEUING SYSTEM**

* **consumers can work on the data independently**

- **you have to estimate the number of shards you will be using, you can change the amount later**

* there is a notion of **kinesis partition key**. When sending events you can specify such key. This key is **used to sometimes guarantee the order of events**, but **mainly it acts as a `spreader` of things on the shards**

##### KCL

- Java app which is used to ingest data from Kinesis stream

* you can specify **maxRecords** property. Having it low, can be a cause of a slow reading speed

- if your **application throws errors** you also might experience slow reading speed

##### Kinesis Data Firehose

- allows you to **store data from kinesis stream on persistent storage, IN NEAR REAL-TIME (minimum 1 min. interval)**. The `near real-time` is very important here.
  Let's say you want to implement a search functionality `in real-time` using `ES`. You **would not use Firehose for that**, you would have to write a lambda function kinesis consumer. On the other hand, if the requirement was `near real-time`, it is completely ok to do it using `Firehose`.

* it can **modify the data before storing it**. You should probably **transform the data to parqet format**.

- **you do not have to specify capacity upfront**.

* **data** can be **send to S3, Redshift, AWS ES, Splunk**. Remember that **FIREHOSE DOES NOT INTEGRATE WITH `QuickSight`**.

- very useful for **replayability and disaster recovery**. You can replay your events from s3 directly.

* **there is no need to configure shards!**, this setting is automatic and completely hidden from you.

- you can **transform** the data before sending it to given service. This is done **by lambda function being invoked by the _Firehose_ service**.

##### Kinesis Data Analytics

- this is **data flows real-time**.

* allows you to make **sql queries against data in the stream**

- can ingest from **kinesis stream and kinesis firehose**

* **serverless product**

- can be used for **ETL on streaming data**.

* can have lambda as target (also s3, redshift, rds).

- under the hood it's **managed Apache Flink** backed by EKS.

##### Kinesis Data Streams

- **processing** data **in near real-time**

* you can use **Kinesis Scaling Utility** to **modify the number of shards**. While quite useful **this is not THAT cost effective**

- your targets can be **Kinesis Firehose** and **Kinesis Data Analytics** and even **Data Stream itself**. This enables you to create 0 code infrastructure.

* normally, consumers contend with themselves on per shard basis. With enhanced fanout consumer, that **consumer gets a dedicated 2MB/s egress limit from a shard**.

- the **iteration logic** is handled **behind the scenes** using **DynamoDB tables**.

* this is usually the Kinesis variant everyone is talking about when then are talking about _Kinesis_

###### Under the hood

- at a **high level** a **Dynamo table is used to track your streams state**. This is true both for Lambda integration and the KCL.

* there is a **concept of lease**. A Lease is a **collection of metadata** maintained in a given row of that table.

- each **row corresponds to a shard** in your Kinesis stream.

###### Autoscalling

- while there is **no native autoscalling functionality** you can deploy your own solution

* this involves a cloudwatch alarm and a lambda function which will do the scaling by using sdk

- there are **api limits for the amount of time you can scale your stream**

##### Parallelization factor

- parallel invocations on the **shard level**

* per partition key order is maintained

- this increases processing power per shard

* also works for EFO (enhanced fanout) consumers

###### Monitoring

- you should monitor **IteratorAge** to make sure you do not have any `poison pill` messages within your stream.

* also worth moniroting **MillisBehindLatest**. High `MillisBehindLatest` can indicate that you are dropping records (**you cannot consume them fast enough**).

##### Kinesis Video Streams

- used for **ingestion and storage** of video / audio but also **can be** used as a **live video / audio source**.

* you can view the live video using either **HLS** or **HTTP** but the **HLS option is preffered since it gives your direct URL, much easier**.

- with **HLS** the live video format is **called archived video**. Pretty weird, no idea why.

* can be **integrated with EC2, Rekognition**. From there you could use data streams of firehose to save that data.

###### Enhanced fanout

- this is something that **applies on a consumer level**

* you can have **up to 5 enhanced fanout consumers per shard level**

- Kinesis **pushes (instead of pooling operation done by the regular consumer)** data **to the enhanced fanout consumer** therefore that consumer does not have to pool the data.

* enhanced fanout consumer costs more (can be much more) than the regular consumers.

###### Shards

- unit of scale within Kinesis Data Streams

* **getRecords** can **only be called 5 times per second per shard**. This means that you can have **up to 5 consumers per shard at maximum**

- data is returned at **2 MB / second / shard** rate. That means that the **regular consumer** can **at maximum pool data once per 200ms**.

* there is a maximum of **1 MB of writes / second / shard**.

####### Resharding

- you can either **split** or **merge** shards

* when you want to merge two shards, you have to **ensure that they are adjencent to each other** (based on the hash key)

- you **might end up with an extra shard while resharding**. This is caused by the **difference between `StartingHashKey` and `EndingHashKey` to be low**

###### AWS Kinesis library

- use it. It has automatic batching, proto-buffers and so on build-in. It will save you some money.

* **can be used on prem**.

- **always prefer kinesis consumer/producer libraries to any SDks**. The kinesis library is really fast and optimized.

* you can actually **have the KCL worker consume `DynamoDB` streams**. This is **NOT a Lambda trigger**.

##### Retries & Error handling

- you can configure **`on-failure` destination**. This can be either **SNS or SQS**

* you can **split batch on error**. This is quite useful for finding the **poision pill**.

##### Metrics

- you can configure **shard level metrics**. This bring extra cost and **needs to be enabled manually**.

### Redshift

- **column-based database**

* **you would store data here after the transactions (changes) has been made**, a good example is kinesis => firehose => redshift. Think of Redshift as **end-state repository**

- data should not change, column-based dbs are quite bad at handling changes

- **scales automatically**

* **USED FOR OLAP style data**

- **always keeps `three` copies of your data**

* **automatically caches repeated queries**

- you can enable **Enhanced VPC Routing**. That means that **ALL operations** performed by **Redshift** will **go through your VPC**. Very **useful when you want to make sure your traffic does not go into public internet**. **Usually** created **along with VPC endpoint gateway**

* **automatically caches SOME quires (results)**. It's up to internal Redshift logic switch query to cache but the process it automatic.

#### Distaster Recovery and HA

- **ONLY SINGLE AZ**. There are **no multiple az deployments**.

- you can use **Redshift Snapshots with S3 CRR** or **enable Cross-Region snapshots for the cluster** for **HA**.

* provides **incremental/continuos backups** just like EBS.

- you can restore given table or restore the whole cluster.

* you can have **cross region snapshots**. If your snapshot is encrypted you will need to provide a **KMS grant** for RDS in the destination region.

- Redshift can automatically **copy your snapshots to another region**. There is no action required from your perspective. Of course **you have to enable this feature**.

#### Architecture

- consists of **multiple nodes**. There is a **leader node** an a **compute node**.

* compute nodes do the work, leader node does the loading of the data

- **minimal storage size** is **2 nodes each 160 GB SSD**

#### Costs

- you can purchase `reserved instances` for underlying nodes.

* you can resize the cluster.

#### Workload Managment Groups

- with WML you can create a query queues based on priorities. This means that you can have a system **where short-queries do not get stuck before long-running queries**.

#### Read Replicas

- **REDSHIFT DOES NOT HAVE READ REPLICAS!**

#### Query problems

- you migth be running out of memory.

* connection to the database timed out

- there is a potential deadlock going on.

* when it comes to potential solutions, you should look for:
  - **reducing MTU**. MTU is the size of a packet that can be transferred.
  - **viewing STV_LOCKS and STL_TR_CONFLICT system tables**. This is done to see if there are any update conflicts.
  - **using PG_CANCEL_BACKEND** to cancel any conflicting queries.

#### Ingestion

- data can be loaded from `s3`, `Kinesis Firehose`, `dynamoDB`, `DMS`.

#### Spectrum

- this enables you to **query directly from data files on S3**.

* this is used when you have **DataLake on s3**. Redshift Spectrum then acts as intermediary tool between other analytics tools and the DataLake.

- **you need to have existing Redshift cluster to make it work!**. The **query is submitted to your cluster** but the query itself is run by AWS.

### Virtual Private Cloud (VPC)

- **CAN SPAN MULTIPLE AZs**

* **SUBNET CANNOT SPAN MULTIPLE AZs**

- **maximum subnet mask is /16**

* there are **5 reserved IPS**.
  - network
  - VPC router
  - DNS
  - Future
  - Broadcast

- you can use **custom DNS (can be on-prem)** by changing **DHCP option set**. `DHCP` option sets contain information about **DNS servers**, **NTP servers** and so on.

* VPC **cannot span multiple regions**.

#### Scalling VPCs

- previously you **had to over provision (still good practice though)** on your CIDR range to make sure that there is some room left for expansion

* in late 2017 an option was added to **add secondary IPv4 address ranges (CIDRs) to VPC**. This allows you to expand your VPC

- note that **your CIDRS CANNOT OVERLAP**.

* you **cannot delete the pimary CIDR range, you can only delete the secondaries**

- when expanding an **new entry to default route is added**.

#### Default VPC

- **you probably should keep it**. Some services need default VPC to exist.

* with default VPC **you get: SG, DHCP, Public Subnet, Attached IGWN, NACL**

#### Elastic IP

- **publicly accessible**, static IP address. Usually used with _NAT-Gateways_

* **FREE OF CHARGE AS LONG AS YOU ARE USING IT**

- **elastic IP persists** even through **stopping and restarting an instance**. This is **not the case with 'auto-assign' public ip option**.

* elastic IP allows you to **mask instance failure**. This is possible because **EIP is assigned to a given istance** so **when instance fails, you can re-map it to other instance**.

#### Internet Gateway

- used for **connecting** your **VPC to the internet**.

* **attached to a default VPC at start**

- **NO BANDWIDTH LIMITS**

* if your subnet is associated with a route to IGW, it is by definition a public subnet

- **DOES NOT TRANSLATE PRIVATE IPS**. It only translates the IP addresses of public IPs

* **supports IPv6**. Since **IPV6 is globally public by default** there is something called **Egress Only IGW to control the privacy in IPv6 setting**.

#### Flow Logs

- capture **metadata** about **the traffic flowing in and out of networking interfaces within VPC**

* **FLOW LOGS ARE NOT FOR SNIFFING OR GETTING THE CONTENTS OF THE TRAFFIC**

- can be **attached** to **VPC** or **subnet** or **an network interface**

* **you can filter, which data you want to see**

- **stored inside CloudWatch or S3**

* **has to have permissions to write to the destination (IAM role)**

- **DOES NOT monitor every traffic**

#### Route tables

- they allow routing to happen within your VPC or make requests to outside your VPC possible

* **usually associated with a given subnets**

- often used to redirect requests to _Internet-Gateway_

* **new subnets are associated with main route table by default**

- **local rules** have the **highest priority**

* **local default rules cannot be edited!**. That means that unless you explicitly create a route table for a given subnet, the global route table will associate CIDR blocks with given subnets inside your VPC

- **one route table** can be **associated with multiple subnets**. **One subnet** can **have only 1 route table**

#### NAT Gateway

- **NOT compatible** with **IPv6**. Use **egress-only IGW for that**

* **CANNOT HAVE SECURITY GROUP ATTACHED TO IT!**

- **lives in a public subnet**.

* has **elastic static IP address**. That Ip has to be assigned. This means that NAT gateway is an ideal solution where your IP needs to be whitelisted.

- **converts SOURCES` private ip address** to its ip address

* sends its traffic to **internet gateway**. Internet gateway will convert the **private static ip to public ip** and send it to the internet

- basically it **allows** resources in your VPC which **do not have public IP** to **communicate with the internet, ONLY ONE WAY**

* multiple **resources** inside VPC **share the same IP assigned to NAT-Gateway**

- scale automatically

* **ONLY HIGHLY AVAILABLE WITHIN SINGLE AZ**. It is placed in a single subnet in a single AZ. **For true high availability create multiple NAT-Gateways within multiple subnets**

- **session aware**, that means that responses to the request initialized by your resources inside VPC are allowed. What is disallowed are the requests initialized by outside sources.

* it **can reside in one subnet AND link multiple subnets IN THE SAME AZ**. Normally though **you probably should use one NAT Gateway per subnet**

- **CANNOT SEND TRAFFIC OVER VPC ENDPOINTS** or **VPC peering**

* remember that NAT Gateway is only used to allow traffic to the internet (and back). You cannot connect with a specific instance yourself since the instances themselves do not have an ip address (assuming they live inside private subnet).

##### Billing

- NAT Gateway are quite costly

* they are **billed for just running**

- they are **billed** based on **data transfer rates**

#### NAT Instance

- an **EC2 with a special AMI**.

* has to be **supported by YOU, not AWS**

- **can have SG attached since it's a normal EC2**

* you **can detach ENI** since again, this is a normal EC2 instance.

- can be used **for port forwarding**. This is **not the case with NAT Gateway**

#### NACL

- **CAN ONLY DENY RULES**

* is stateless, that means **it does not _remember_ the relation between incoming and outgoing traffic**.

- **cannot block traffic to a given hostname**

* **can** be **associated with multiple subnets**

* **WORKS ON A SUBNET LEVEL**

- **ephemeral ports** play a huge role here. These are **randomly selected ports to return traffic for a request**. This means that if I **send a HTTP request (port 80)** as an inbound rule I have to **specify ephemeral port rage on inbound rule**. Also remember that it's not only about communication with the internet. Since NACL are subnet level thingy it may be the case that you have to setup ephermal ports in multiple NACLs when talking between subnets.

* **they CANNOT REFERENCE Logical Resources**

- **rule prioritization just like in SG**. **Lowest number wins**

* **default** rule is an **asterix, which is an implicit deny** and **100 rule to allow all traffic**

- remember that **NACL only look at IPS and protocols**. You **cannot use SG to filter based on url** or something similar.

#### Security Group

- **CAN ONLY ALLOW RULES**

* there is an **default explicit deny on everything**

- **remembers the relation between incoming and outgoing traffic**. If you ping and instance with security group attached it will be able to ping you back without having to specify outgoing allow.

* is **attached to an ENI**. This means that **if your instance has multiple ENIs** you can have **multiple security groups on 1 instance**.

- **changes** to security group **are instant**

* security groups **are stateful**. That means when you **create an inbound rule, outbound rule is created** automatically.

- there are **no deny rules**. Security groups can only **specify allow rules**.

- **Security Groups DO NOT APPLY TO s3 BUCKETS, WE ARE IN EC2 VPC SPACE NOW!**

- **All inbound traffic is blocked by default**

* **security group can have other security groups as sources!**. This does not mean that we are _merging_ the rules. **Having other security group as source means that we are allowing traffic from instances ENI which are associated with that group!!!!!!!**

- remember that **secutiy groups only look at IPS and protocols**. You **cannot use SG to filter based on url** or something similar.

#### DNS in a VPC

#### Resolution

- the **DNS name of an instance** will **resolve to the public IP when reaching in public internet** and to the **private IP, when resolving within a VPC**

* this **can be a problem** while **resolving from on-prem**. **Instead of going through a eg. VPN** you **will resolve to public interface**. (**unless using route53 resolver**)

- you can use **private hosted zones** to **override public DNS endpoints**.

* to make sure that the DNS is working within your VPC, **enable `enableDnsHostnames` and `enableDnsSupport`**. They are both set to **true by default**.

#### Peering

- you **cannot use** **anothers VPC NAT Gateway**. So the setup where you have peered VPCs and one is trying to connect to the internet using others NAT Gateway will not work.

* **linking TWO!! VPCs together** (in a scalable way)

- when VPCs are peered, services inside those VPCS can **communicate** with each other using **private IP addresses**

* can **span accounts, regions**

- VPCs are joined using **Network Gateway**

* **CIDR** blocks **cannot overlap**

- **VPC peer has to be accepted by the other side**

* you will probably have to **check SG, NACL** to make sure it works. **Enabling peering DOES NOT MEAN that the connection is made**.

- VPC peering **does NOT allow for _transitive routing_**. That means that if you want to **connect 3 VPCs** you have to **create peering connection between every VPC**. You **cannot communicate with other VPC through peered VPC!**

* works in **LAYER 3** of OSI model

- **VPC can be in a different region**.

* the **traffic** is **not over the internet**. It uses the aws global network.

#### Transit Gateway

- **you can link multiple VPC using Transit Gateway**

* **CIDRS cannot overlap**, but you make much less connections between VPCs in general

- **works** with **direct connect and VPNs**

* you can create **route tables on transit gateway to control the visibility of each VPC**

- there is a **static monthly cost** and also **data transfer cost**.

#### Transit VPC

- **one VPC as pass through**

* can be **used for connecting multiple cloud providers**

- this is a **VPC that contain specific EC2 instances**. You **connect** to transit VPC **using VPN (Virtual Private Gateway and BGP on the transit gateway side)**

* the **connection IS NOT IPsec (vpc peering)**. This is due to routing issues.

![](./assets/transit-vpc.png)

#### VPC Endpoints

- there is a notion of **VPC endpoint**. This allows the **service that the endpoint points to** to be **accessed by other AWS services without traversing public network**. **NO NATGW or IGW needed!**

* **They can only be accessed within a VPC**. That means that you cannot access the endpoint through a VPN, Direct Connect and such!

- **private link is used underneath** for seamless connection to AWS services.

##### Gateway Endpoints

- a gateway that is a **target for a specific route**

* uses **list of predefined IPs** to route traffic

- only **DynamoDB and S3** are supported

* to control the policies you should use **VPC Endpoint Policies (they do not override IAM roles)**

- the **mapping of IPs** occurs on **route table level**. Route table uses **prefix lists**.

* remember that Gateway Endpoint can connect to **multiple s3 buckets** and **multiple dynamodb tables**

- there is **no additional charge** for using gateway endpoints. The only costs are **data transfer costs**
  and, of course, resource costs.

##### Interface Endpoints

- **ENI with a private IP (routing not involved)**

* relays on **DNS** to work

- **for any service that is NOT Dynamo and S3**

* you can use **Security Groups** to **control the access**

- With **Interface Endpoints** you have to **manually select AZs** to make it **highly available**

#### IPV6

- **not enabled by default**

* **have to be enabled for the whole VPC**

- **all** IPV6 addresses are **publicly accessible by default**

* when enabled **is attached to the operating system**

#### Egress-only Internet Gateway

- **allow only OUTBOUND traffic from IPV6 associated instance**.

* **engress-only** means that it **only allows outbound IPV6 connections**. It's **stateful!**. Which means that it **allows elements in your VPC** to **receive the response back**.

#### Private Link

- **THIS IS NOT THE SAME AS VPC PEERING**.

* VPC private link gets **created automatically when you create VPC Interface Endpoint**

- **highly available**

* private link allows you to **share a service that YOU created, not only a AWS resource which is the case with Gateway/Interface endpoints**

- **provider exposes NLB** and the **customer links VPC endpoint network interface to that NLB**

- uses **DNS underneath just like interface endpoints**

* **can be combined with DirectConnect** to **slowly migrate from on-premise**

- **DOES not traverse the public internet**

#### VPN

- quick to setup and relatively cheap. It's **billed hourly**

- **virtual private gateway** is **HA by default**

* **tunnels are encrypted end-to-end**

- uses **internet for transit**

* you do not need IGW for VPN to work.

##### Between VPCS

- you can use **software 3rd party solution** between **2 vpcs**

* you can use **AWs managed (private virtual gateway)** to **software managed**

##### Site-to-site (on-prem => aws)

- allows for **on-prem to VPC** and **VPC to on-prem** connectivity.

* **routing** can either be **static** or **dynamic**.

- if you specify **dynamic routing** the connection will use **BGP**.

* if you specify **dynamic routing** both sides **can exchange routing information**. You **do not have to provide such information yourself, like in a static routing solution**

- the **connection** occurs between **customer gateway** and **virtual private gateway (attached to VPC)**

* you can **create multiple tunnels** to achieve **HA on AWS site**

- you can create **multiple customer gateways** and **multiple tunnels** for **full HA**. This architecture **requires BGP**

* when you setup the connection the **peer identity authentication is established**

- it **uses the public internet** to exchange data.

##### Managed VPN

- **AWS managed** VPN

* this is the **default offering** - the thing that you setup within the console.

- connection **can be of type "on demand"**. This means that the **connection will not be established unless there is some traffic**.

##### Client VPN

- establishing connection using OpenVPN client.

* your entry point is the **Client VPN endpoint**. This is a **regional construct**.

- this allows you to **connect an user to a subnet WITHOUT any router appliance**.

* you can **attach security groups** to **Client VPN endpoint**.

##### VPN CloudHub

- this is just one `Virtual Private Gateway` to multiple `Customer Gateways`

* creates so called `Hub and Spoke` model. **VPC is the Hub** and **each on-prem location is the Spoke**.

##### Customer Gateway

- **represents a physical router on customer side**

* **you have to tell AWS the IP of your router**

- you can also specify **dynamic IP**. This will allow for **better communication in terms of IP discovery between your router and AWS VPC**. This is used mostly **when you have multiple Customer Gateway routers**

##### Virtual Private Gateway

- this is **similar to NAT Gateway**.

* **you have to attach it to specific VPC**

- **used to connect to Customer Gateways**

* **HA by design**

* you **can** have **route table attached** to it. This is qutie useful eg. when you want to route traffic to firewall first AND THEN to your instance.

##### Site-to-site Connection

- this is **logical entity that connects Customer Gateway and Virtual Private Gateway**

* **sometimes refereed to as Tunnel Endpoint**

##### HA

- it's the **customer gateway that is the point of failure**

* you should probably **provide 2 or more customer gateways** so that when one fails you can switch to the other one.

#### Direct Connect

- you have to have **BGP and BGP MD5 enabled** device.

* **psychical connection between AWS and your network**

- you are connecting to **DX (Direct Connect) Locations (psychically!) or Direct Connect Partner**.

* **port speed is very specific**. If you are **ordering from AWS directly, you can pick 1 Gbps or 10Gbps**. If you **need something slower, reach out to Direct Connect Partners**

* **used** when you **need FAST speed and low-latency connections**

- **takes a lot of time to setup**

* **BY DEFAULT THE CONNECTION IS NOT ENCRYPTED (the transit)**. **DIRECT CONNECT BY ITSELF DOES NOT ENCRYPT THE DATA**

##### VIFs (Virtual Interfaces)

- **public VIFs** allow you **to connect to any public AWS services directly**. Like **S3 or Dynamo**

* **private VIFs** will **connect you to a specific VPC**. You **still need Virtual Private Gateway**

- **transit VIF** will **connect to DirectConnect, which connects to Transit Gateway**. You can have **only 1 Transit VIF per Direct Connect**

* you can have **up to 50 PUBLIC + PRIVATE VIFs** on a **single DirectConnect**. This might be a problem. This is where **the usage of DirectConnect Gateway come in**

##### Direct Connect Gateway

- allows you to **fan-out from a single private VIF** to **up to 10 VPCs**. This is super nice since **before that** **private VIF was region locked**.

* the **fanout from DCG** can be used **cross-region**.

- can be **attached to Transit Gateway(UP to 3 cross region)** with the **usage of Transit VIF**.

* within your VPC you should have **virtual private gateway**.

##### DirectConnect + VPN

- this setup basically means **setting up VPN connection over DirectConnect public VIF**

* mainly used **for encryption**. What is because **Direct Connect does not encrypt** the data but **VPN does encrypt it**.

- remember that **DirectConnect routes are always favourited against VPN routes**. You can **change BGP route weight for adjustments**.

##### Combining Direct Connects

- you **can actually combine direct connect connections**

* you have to use something called **LAG - link aggregation group**

- **all connections within a LAG have to have the same bandwidth**

* you can **add existing or new connections to LAG**

- this solution **will increase the overall bandwith**

* you can have **maximum of 4 connections within a LAG**

### AWS Rekognition

- allows you to **recognize faces from images**

* allows you to **create Video Stream processors** which parse the video **consumed from Kinesis and output to other Kinesis stream**.

- **DOES NOT RETURN IMAGE METADATA**

### AWS WorkDocs

- competitor to GoogleDrive and such

* you can store virtually any file there.

- you **can develop rollback feature using API**

* there is something called WorkDocs Content Manager.

### AWS Backup

- operates on a notion of **backup plans**. This is where you specify **how often you want to run the backup**, **whats the retention** and so on.

* you **can create backup plans from JSON files**

- there are also some **options to configure the lifecycle, like moving backups to cold storage after X time after creation**.

* there are also **backup rules, backup plan can have multiple rules**.

- you can **assign multiple resources to the backup plan**. AWS Backup can work with **RDS (except Aurora), DynamoDB tables, EFS,EBS, Storage Gateway volumes**.

* you can create **on demand backups**.

### Caching

#### DAX (in-memory cache for DynamoDB)

- runs **inside VPC**

* uses **cluster architecture**

- works as any cache you expect to work. You get stuff, it's placed in cache, when you get it again you get it from the cache.

* **read-heavy workloads or very low latency**

- two caches, **query cache** and **getItem/batchGetItem cache**

* any reads through DAX are **eventually consistent**

- **can be HA**.

#### ElastiCache

- you **cannot point** route53 **alias record to ElastiCache cluster**

* operates with **products that are NOT dynamoDB**

- supports **Redis** or **Memcached**

* there is **no failover**. You **can run multi AZ** using so called **spread nodes mode**. **Think of this as ASG and EC2**, number of nodes can shirk and get bigger with time.

- you might have a **situation where only SOME of the requests are served by cache**. This is **probably due to the underlying instance being too small**.

#### Redis

- **by default** **data** in a Redis node **resides only in memory and is not persistent**.

* can be used for things as **leaderboards**

- have **pub/sub** capabilities.

* you can **backup existing data** and then **restore that data**. You could also **use AOF (append only file) and resotre from that file**.

- supports **in-transit and at-rest encryption**

* there is an **AUTH for Redis** thingy that can require user to give a token (password) before allowing him to execute any commands

- has a **concept of read replicas just like RDS**. Similarly to Aurora **read replica can be promoted to master**. The **replication between nodes is ASYNCHRONOUS**

##### Clustered mode enabled

- data is **partitioned across shards**, each **shard has 1 primary and up to 5 replicas**

* failover is faster

- better for fault tolerance.

* better for horizontal scaling.

##### Clustered mode disabled

- data **resides in one primary** which can have **up to 5 replicas**.

##### HA

- Redis **supports multi-az with automatic failover**

#### Memcached

- can be used for **database caching** (usually the SQL ones).

* data there is **lost when** instance (or cluster) is **stopped**.

- **DOES NOT support encryption natively (through KMS)**.

* **nodes can be spread a cross multiple AZs**, but there is no failover.

- is **designed** to take **advantage of multiple CPU cors**.

* **nodes** can be **discovered** by using **auto discovery** feature of Memcached.

- can be scaled vertically and horizontally, but when you **scale horizontally** you have to **create new cluster**.

* **data** is **not automatically replicated betweeen the nodes**. You have to make sure that you are **sharding your keys to different nodes**.

### Communication Between Services, Queues

#### SNS

- notification service, **think of fan-out pattern**

* supports **email/json or HTTPS**

- can be used to provide **fully manager messaging service**

* messages can be **up to 256KB in size**

- **basic entity** is the **topic**

* multiple **publishers** can **send messages to a topic**

- **resilient across ALL AZs within a region**

* **messages** can be **encrypted at rest and in-flight**

- can have **resource policies** applied

* **subscribers** can have **filters applied on the SNS topic**. The filtering itself is something named **filter policy**.

##### Non-Lambda integration

- SNS can be integrated with lambda without much hussle, but there is more work to be done when you are not running your code within lambda.

* your application has to be ready to **receive POST calls from SNS**. This means **doing things when SNS requires you to confirm the subscription**.

- you **have to response to a message within 15second window**, otherwise SNS will consider that request a failed attempt.

* you should **read x-amz-sns-message-type header** to get necessary information with what you are dealing with. It could be either subscription confirmation request or a message.

#### SQS

- **nearly unlimited throughput FOR STANDARD QUEUES (not FIFO)**

* **up to 256KB payload**

* when a **consumer reads a message from a queue**, that message will be **_invisible_ for other workers up to X seconds**. This is so called **visibility timeout**. If you **process the message and do not delete it** that message **will be _visible_ again for processing**. This is how **retry mechanism** is implemented.

* you can **change visibility timeout PER ITEM BASIS**. This might come in handy when you know some messages can take longer than usuall. This is **usually done by tagging such message using some kind of JSON header**.

- **by default** your **standard queue** **DOES NOT PRESERVE THE ORDER**. You can also **have duplicates (sometimes)**.

- can have **resource policies**

* you can have ASG react to number of messages inside the queue

- you can also set up **DelaySeconds (Delay Queue)**. This will make sure that **any new message will be invisible for X seconds for consumers before being available for processing**. DO not mistake this with _Visibility timeout_

##### Pooling

- there is a notion of **pooling**.
  - **short pooling**: up to **10 messages** at once. You **constantly have to check the queue**
  - **long pooling**: you **initialize long pool request**. You **wait for that request to finish**. This **request will finish when wait-time exceeds specified time (max 20s) OR queue is not empty**. This will enable you to **avoid empty API calls**

Whats very important to understand is that **LONG POOLING CAN END MUCH EARLIER THAN THE TIMEOUT**. The **connection** is **always open**, it just waits for ANY message to be visible.

- you can **control how long the long pooling takes** by **specifying ReceiveMessageWaitTimeSeconds attribute**

> The length of time, in seconds, for which a ReceiveMessage action waits for a message to arrive

##### HA

- every **message** is stored on **3 hosts across MIN 2 AZs**

##### SQS FIFO

- **FIFO queues** allow for **ordering** but have **limited capacity**.

* **FIFO queues are not supported for lambda integration**. Also, **with Lambda, max batch size = 10!**

- you **cannot** **convert existing queue to FIFO**. You **have to create new one**.

* **DOES NOT SUPPORT DELAY SECONDS**

- you **HAVE TO use Message Group ID**. If you do not have multiple groups just just a single static one.

- **Message Group ID** allows for **parallel consumption of batches**. If you need to preserve order within a given group, but those groups are independend, using message group ID **that is unique PER group of objects** can greatly speed up the processing time.

* you can use **Message deduplication ID** to **ignore (5 mins period) messages that have been send to the queue multiple times**. This is only available for SQS FIFO.

#### DLQ

- this is a **special queue which usually takes the events which were processed unsucessfuly**.

* normally each message has something called **`recieveCount`**. Whenerver you get message delivered that count is incremented

- you can **set that message will be transported to DQL whenever `recieveCount` is greated than ...**

#### Redrive Policy

- this is the **policy used to determine when to move given msg to DLQ**

* very often used with `ReceiveCount` attribute, like **specifying `maxReceiveCount`**.

- you can redrive messages from DQL using [replay-aws-dlq](https://github.com/garryyao/replay-aws-dlq), you probably want to use `npx` for that.

### SES

- for sending **emails, and only emails**.

* you can send **rich text, multimedia etc...**

#### SES vs SNS (email)

- **SNS** is for **simple UTF-8 text based emails**

* **SES** is for **reach, containing multi-media emails**.

### EventBridge

- SNS and SQS combined (more or less)

* can be **integrated with much more services than SQS or SNS (natively, without pooling)**.
  What's more important that you can **integrate natively with 3rd party AWS service providers**.
  Since these are going through APIGW the integration could also be done using SQS or SNS.

- does **not support FIFO ordering**

* **NOT for high amounts of events per second**. You are billed per send an event. For case where you have a lot of events look into Kinesis.

- one important thing to remember here is that **EB supports only 5 targets per rule**. This is something you should keep in mind while designing stuff.

#### Resiliency

- you can use **DQL per rule**. Please note that **this DLQ is only for the EventBridge communication with the underlying service**.
  If your lambda throws an exception, the _EventBridge_ will not care. As far as it's concerned, the event was send successfuly.

* you can specify **custom number of retires** as well as **time spent retrying**.

- when an **event is sent to rule DLQ**, the message is **annotated with additional info**. This should help you with debuging. Also remember to **always have a dlq per rule**.

#### Schema Registry

- source of truth for events that flow through the EventBridge

* you can generate code bindings

- this is a free service

##### Schema Discovery

- is responsible for putting discovered schemas to schema registry

* you literally just have to click 1 button on a given event bus and the service is on

- can be an source of events. Whenever it discovers the schema it will publish to a default EventBus

#### vs SQS

- **EventBridge** has **more integrations**

* **SQS** has much **longer retention period (up to 14 days)** whereas **EventBridge** will **hold messages up to 24hrs**.

#### vs SNS

- **EventBridge** has **bigger number of targets** compared to SNS.

* **EventBridge** has **better filtering capabilities** compared to SNS.

- **SNS** has **much higher throughput** than EventBridge.

#### vs CloudWatch Events

- `EventBridge` pretty much **has the same functionality (and more) as `CloudWatch Events`**.

* there is also `CloudWatach Bus` which is depcracated. This used to function similarly to `EventBridge`.

#### Cron and fixed schedule

- with `EventBridge` can create _cron_ schedules or invoke target at fixed rate. The **minimum granularity** is **1 minute**.

* you **cannot use custom bus for cron / fixed rate schedule**.

#### Event Manipulation

- you can **transform the event**. This of this picking and choosing which fields you want to forward to given service

* you can **return static response** to given service. This of counting service

#### Batching

- there is **not build-in batching** like SQS

* you have to **batch manually, probably the inside the detail type**.

#### Debugging

- debugging is hard but possible

* what I would suggest is to **make cloud watch log group as your target**. You will be able to see all the events!

- you **cannot create the rule for cloudWatch log group through CF**. This is a mess :C

#### DLQ

- you can set up DLQ for a given target of a rule

* the EB will **annotate the message that is pushed to DLQ with basic error info**. This will allow you to debug stuff bettter

#### Decoupling

- one thing that _EventBridge_ is really good at is the decoupling

* one pattern that is to relly on _CloudTrail_ to _EventBridge_ integration in some situations, rather than the native integrations between services. Think multiple buckets hooked into 1 lambda function (putEvent)

### Amazon MQ

- occupies VPC, and only VPC. This is **different than SQS or SNS** where the service CAN be made private, but it's public by default.

* supports **topics** which are **comparable to SNS topics**.

- supports **queues** which are **comparable to SQS queue**.

* supports **virtual topics** which are basically **fanout pattern, but without using SQS and SNS at the same time**.

- usually used when you have some internal requirements. This service is not well supported when it comes it IAM or CloudWatch.

### AWS Workspaces

- desktop as a service

* basically you can provision windows or linux desktop to your employers

- you do not have to manage hardware

### AWS OpsWorks

- there are **three services under the `OpsWorks` umbrella**

* there is a **notion of a recipe**. This is a **unit for work** that **you want the service to perform**

#### AWS OpsWorks Stacks

- **in between Elastic Beanstalk** and **manual deployment**.

* it **trades SOME of the configurability** for automation

- **stack** as **top level construct**. Type of system like dev, prod, test or a specific application.

* **layers** represent **individual pieces of functionality** within a stack, something like ECS Cluster, RDS, or OpsWork Layer

- uses **Chef solo for configuration**

* uses **declarative language for configuration**. You basically tell it what you want to happen, the service will figure the rest out

##### Lifecycle hooks

- there are **five lifecycle hooks** that you can listen. Mainly: **setup, configure, deploy, undeploy, shutdown**

* the **configure event** is usefuly for **regenerating / changing configuration files**. The **event fires when an instance is created or destroyed**

#### AWS OpsWorks for Chef Automate

- _Chef Automate_ is a **managed _Chef_ server** for configuration, compliance. Basically swiss army knife of devopsy stuff

#### AWS OpsWorks for Puppet Enterprise

- a **fully managed configuration mechanism**

* think about it as **more powerful version of _Chef solo_**

#### AutoHealing of instances

- this is **like automatic health checks with an autoscaling (but fixed)**

* you **only have to tick 1 box to make it work**

- if you want notifications about auto-healing you should **look into CloudWatch (events)**. There is no native ops-work integration when it comes to auto-healing and sns.

#### Instances

- there are multiple instance types:
  - 24/7h instances
  - load-based instances
  - time-based instances
    Combination of these should be used for scaling.

* there is a **simple wizard to create scalling scenarios**

#### Updating instances

- **updates** are **applied at the time of launching given instance**. There is **no automatic update process in place, you have to update your instances later on**. You can do that **in 2 ways**:
  - **create and start new instances** to **replace your current ones**.
  - **on Linux-based instances** you can **run `Update Dependencies stack command`**.

#### Blue-green deployments

- just like with `ElasticBeanstalk` there is a possibility to do blue green deployments with `OpsWorks`

* you should **clone the entire stack** and **change DNS to the other stack**. Much more involved than EB but it works.

#### Lifecycle events

- each **layer has lifecycle events**

* the **recipies run when those events happen**

- overall, **there are 5 lifecycle events**. Mainly Setup, Configure, Deploy, UnDeploy and Shutdown

### WAF (Web Application Firewall)

- **layer 7**.

* you can **associate** WAF with **CloudFront, APIGW, ELB**

- **traffic** is **filtered before reaching /\ services**

* use **AWS Shield for DDOS protection, WAF is just to help you filter the request with rules**.

- you can match based on headers, ips any many other stuff

#### WEB ACL

- set of **rules** or **traffic decisions** you apply to a specific project.

* there is a **default rule** for traffic that does not fall under any other condition

- it **filters** the traffic **before that traffic ends up at your servies**

* **not design** for **large scale thread protection**. Conditions have to be explicitly defined, it is not _smart_.

#### IP Sets

- you can create **IP sets**. You can **reference them within your rules**.

* the created IP sets have **arn associated with them**.

- **when creating** you have to use **IP address with CIDR range**. This allows you to block multiple IPS :)

#### Rules

- **regular** rules **match conditions** (you can use IP sets within a condition)

* **rate-based** rules **match for a given frequency**.

- **rate-based** rules can be used to **fight DDOS attacks**.

* you can also **restrict** traffic **based on location**, **prevent SQL Injection**, all sorts of stuff.

### AWS Shield

- mainly for **creating protection against DDOS**.

* while you could be doing this using NACL, the DDOS traffic would be dangerously close to your system (already entered your VPC). You should prefer to elivate the tread as far of your network as possible.

- **offered for free for all AWS accounts automatically**. This is to keep base security level of all accounts within aws.

* it **covers different services than WAF**. One notable fact is that **it does not cover APIGW**

#### Shield Advanced

- gives you more **premium features** but costs **3k/month (costs of WAF included)**.

* there is a concept of **cost protection**. This is where **when you incur a cost on R53, CloudFront and ELB during DDOS attack, you can get your money back**. Pretty neat.

- with Shield Advanced you can **also protect Elastic IPs**.

### CloudSearch

- **AWS own solution**, fully managed by AWS

* still ,**there is an underlying instance hosting the search domain**

- **automatily scales VERTICALLY AND HORIZONTALLY (in that order)**. When scalling to multiple instances whe search index is partitioned into multiple instances.

* you can use **multi-AZ option for HA**.

- **integrated with IAM**.

* feels a bit abandoned

### Elasticsearch Service

- it's **AWS implementation** of **ELK Stack**

* scales **vertically** and **horizontally**.

- for production workloads, ES can use **Multi-AZ**.

* **integrates with IAM** for **resource, identity and IP-based policies**.
  Identity based policies should be used to narrow down the scope to given sercices, while resource based are applied to the domain level.

#### ELK stack

- **E** is for **ingestion and (optionally transform)**: Logstash or Beats

* **L** is for **Elasticsearch**

- **K** is for **visualization**: Kibana. This is mainly used for **anything that is NOT analytics data**. For **analytics data** use `QuickSight`.

* **Kibana** also allows you to **share dashboards**.

### CI/CD

#### CodeCommit

- **hosted Git Repo**

- you can also **use comments** just like you would on github. For the best experience **use comments feature when you are signed as IAM user** (not federated , or temp credentials).

* to begin using the repository, **you have to generate credentials for it**

- as best practice you should use `CodeCommit` as an IAM user

##### Triggers

- can have **up to 10 triggers defined**. Main use case would be to configure **SNS to send updates or lambda** to other developers.

* **limited in scope**

- they **DO NOT** use **CloudWatch Events to evaluate events**

##### Notifications

- **much more granular than _Triggers_**

* from the UI can **only** pick **SNS**

- if you want **more targets - look into CloudWatch events** as the _Notifications_ feature is built ontop of them

##### Protecting branches

- you will need to **use IAM policy**

* the IAM policy will most likely use **Condition with `codecommit:References`**

##### Encryption

- the data that you store in the repo is **automatically encrypted at rest and in transit**

* service is using **it's own CMK for crypto**. That **CMK is regional, it is shared between existing repos**

#### CodeBuild

- Build packages for deployment

* remember to **get the agent from correct region!**.
  This can happen when you are using `user data` to bootstrap the agent
  `wget https://aws-codedeploy-us-east-1.s3.amazonaws.com/latest/install` (look at the region!)

- when working with the buildspec, **do not define any env. variables that start with `CODE_BUILD`**. This prefix is reserved for internal use.
  Ignoring this advice might result in slower build speeds as that prefix also includes memory settings.

* you can configure CodeBuild to have access to VPC. This will require you to specify VPC ID and SGs. Remember that **NAT GW or NAT instance** is **required**.

- there are **metrics** available inside **CloudWatch**. These are about build mainly, success, failure, all that stuff.

##### Artifacts

- you can tell CodeBuild to put artifacts to s3 for you. **By default, CodeBuild will encrypt those artifacts**

* **by default** the **service role of CodeBuild** gets **s3 star IAM permissions**

- **you might have problems with permissions uploading if you have custom bucket policy**

##### Build triggers

- you can have **CRON expression** for **scheduled builds**

##### Envioriment Variables

- there are some **reserved environment variables**. For `CodeBuild`, they have a prefix of `CODEBUILD_`

* you are **able to override / specify the reserved environment variables**, you should not do that though

- you can either use **SSM or regular values** for your environment variables

##### Metrics

- there are **number of default CW metrics**. These include **no. total builds, failed builds, successful builds and the duration of builds**

##### Buildspec file

- you can **pull things from SSM and SecretsManager**

* most of the phases has the `finally` block available to them. You will most likely use that block for cleanup

##### Quotas

- the **maximum timeout** you can specify is **8 hrs**. That means that you can run relatively long running jobs on CodeBuild

##### Output artifacts

- you can **change the name of the output artifacts in the artifacts section of the buildspec.yaml**

* you can **specify the `Namespace`**. This `Namespace` **will add buildId to the path of the partifact** (not filename)

#### CodeDeploy

- Deploy packages to given services (like ElasticBeanstalk)

- allows you to perform **rolling** updates for **ec2 instances**.

- it enables you to perform **blue / green deployment** with **ASG**.

* when it comes to **ECS, it allows you to create blue / green (traffic shifts)** with that aswell.

- can **integrate with AWS Config** to make sure changes are compliant, otherwise they will not deploy.

* you can **schedule jobs using CloudWatch events**.

- it **DOES NOT integrate with CodeCommit directly**. You can **either supply an s3 location or GitHub repo**

* uses the **appspec.yaml** file

##### Deploying to ASG

- it may happen that **scale out an event will occur during the deployment**. In such situations, **you will probably have 2 versions of your application running**.

* you should **suspend asg for the deployment period** or **redeploy your application again** after the initial deployment.

- you **can deploy based on tags** on the **underlying EC2 instances**

##### Deploying to Lambda service

- you **do not have to store your artifacts in s3 if you are using CodeCommit**

* you can either use **traffic shifting** or **canary deployment**

- with **canary deployment** the traffic is **shifted in TWO increments**. So if you see something like `LambdaCanary10Percent10Minutes` that means that the lambda will be taking 10% of the traffic for 10 minutes, then ALL traffic will be shifted to it.

##### Validation hooks

- when deploying with **lambda, using traffic shifts** there are **pre and post deploy hooks** you can use to validate your lambda.

* there are also _validation hooks_ for **ecs**. There is a lot more of them than for lambda functions.

##### Deployment Group

- **configuration set** which will be **used during given deployment**

* think of **alarms for deployment, triggers and such**

- you are able to **edit the deployment group without any issues**. So going from blue-green to in-place is possible.

##### Application specification files

- used to **managing each deployment** along with **lifecycle event hooks**

* current options are: **ECS, Lambda and EC2 (on prem as well)**

- these allow you to specify **lambdas which will be invoked to check the deployment**

* there are **multiple lifecycle events, they differ based on the service**

##### Canary deployment

- you can create a _canary deployment_

* the **_canary deployment_ option is only available for Lambdas and ECS**

##### Health of the instances

- there is **NO SUPPORT FOR ELB Health Checks**

* your mechanism for **checking if the instance is healthy is the `ValidateService` hook**. **This hook is only available for EC2**

##### Automatic rollbacks

- you can setup automatic rollback options for your deployment group

* the automatic rollback **can be triggered** when **deployment fails** or the **alarm thresholds are met**. Of course, you have to associate alarms with deployment group.

##### CodeDeploy triggers

- you can setup **triggers to fire to SNS topic**

* these can **notify you about _Deployment events_ and _instance events_**

#### CodePipeline

- enables orchestration of all the above

* you can implement **manual approval step**. The **pipeline will wait 7 days** in that step, if not approved / rejected , pipeline will timeout.

- you can integrate _CodePipeline_ with your github account by using either _OAuth_ or _personal access_ tokens.

* you can run **stages in parallel** by using **run order** parameter.

- remember that `CodePipeline` **orchestrates your pipeline**. That means that **it also handles `CodeBuild` invocations**.

##### Custom action job workers

- you can have **custom action types** within _CodePipeline_

* those **actions are executed by the worker**. The **worker pools _CodePipeline_ for any jobs**, then **returns results to the _CodePipeline_**

- the worker **has to pool CodePipeline for tasks to perform**

* you would use this for any 3rd party, proprietary tool that needs to be included in the pipeline.

##### Cross region actions

- you **CAN NOT** create **cross region source, third-party and custom actions**

* CodePipeline **copies artifacts from the build region to a given region automatically**. You **have to have the buckets defined beforehand**

- **CodePipeline cannot invoke other CodePipeline directly**. You should look into creating a `source` action (s3) when it comes to cross region deployments.
  If you really need to invoke other CodePipeline, look into custom actions with lambdas

##### EventBridge / CloudWatch events integration

- you can listen to events produced by the _CodePipeline_ using _EventBridge_ or _CloudWatch events_

* one pattern is to use **Systems Manager Automation with CloudWatch / EB based on those events**

##### Artifacts

- you can **deploy to multiple regions**. This means that you can have per-region artifact stores

* you **have to have artifact store in the region the pipeline is defined**

- if you **create the _CodePipeline_ through the console**, **the wizard will create AWS managed CMK to encrypt the artifacts**. This is **not the case if you do it through the CLI**

* you can create your own CMK and manage it yourself

##### Lambda integration

- you can invoke lambdas directly from CodePipeline

* you can pass some data from the CodePipeline to the lambda payload through the _user parameters_

##### Creating ECR Image

- you should use **helper scripts** to **aquire ECR credentials**. You should not be using environment variables

### CodeStar

- _CodeStar_ **ties all the Code services together**. This includes the pipelines and other stuff.

* you configure it through the `template.yaml` file, which is basically the `SAM` file.

- each **project has three roles**. These roles being: **owner, contributor and viewer**

#### User profiles

- CodeStar user profile is **associated with your IAM user**

* you can **upload SSH public key** to be **associated with your profile**. Doing so will allow you to **connect to EC2 instances associated with CodeStar projects**

### Cognito

- authorization and authentication service

#### User Pools

- this is **where your users live**.

* user pool is responsible for **authenticating the users**. It returns the tokens, these tokens can be exchanged for AWS credentials using Identity Pools.

- user pool also **manages the overhead** of **handling SAML/Social sign in provider tokens**

* you **do not have to exchange the tokens using Identity Pool for AWS credentials**. You **might want to control access to your backend resources using those credentials**. (eg. APIGW authorizers)

##### Resource servers

- these allow you to set **custom oAuth scopes**

* **name is misleading**. The **identifier does not have to be an url**.

- you can **set scope authorization on APIGW level**. You can have authorizer check those on per endpoint basis

* these **will not be retrieved when you use amplify js library**. You need to use the _cognito endpoint_ (probably hosted UI). If you really need to get them programatically, you can try the _cognito-js_ library.

#### Identity Pools

- **here is where you grant permissions to users**

* identity pool is **responsible for exchaning user pool tokens for AWS credentials**

- here you can **enable unathenticated access**. This is quite logical that this option is here since remember that Identity Pool is responsible for STS tokens.

#### Multi-region

- Cognito **is a regional service**. You will need to implement the replication yourself.

* One solution I saw is to have **triggers write user data to DynamoDB global tables**. This is still not ideal as the user would have to have different passwords per region, but might be a **solution for active-passive** architecture.

#### SAML

- to get SAML working you will probably need **provider Name and MetaData document**

* instead of doing AssumeRoleWithWebIdentity (like with Google and Facebook) AWS performs **AssumeRoleWithSAML**

- **can be used to access AWS console**. This is **not the case with Web Identity**.

#### Identity Brokers

- sometimes your identity store might **not be compatible with SAML 2.0**. Then you will need to **create your own identity broker**

* this usually works as follows
  - identity broker check if user is already authenticated to your internal system
  - calls _sts:AssumeRole_ or _sts:GetFederationToken_ to get the credentials
  - passes those credentials to users (remember these are temporary credentials)

- you **do not have to create a new user specially for this**. While working with `Identity Brokers` you should use **IAM role** that the federated user will assume.

* remember that **your identity broker is talking to sts, NOT YOUR USERS!**.

#### Identity Federation

- **keys obtained** with the **help of STS or Cognito**

* **SAML 2.0 => Active Directory**

- **Web Identity** means getting the **initial key from FB, Google...** and **exchanging it** for **temporary credentials** within AWS. A **role is assumed after successful token verification**.

##### Active Directory

- you can **either deploy manually on ec2** or use **aws managed AD**

###### AWS Managed AD

- **does not replicate users with on prem**. This is quite important. Users live in 2 different directories, one on prem and one in AWS.

* if you want to synchronize users, you should deploy **self managed AD** and use **replication between it and the on prem**. Then you would establish trust with AWS managed one.

###### AWS AD Connector

- there are 2 versions, **small** and **large**. It **does not have user limits** BUT it has **performance and scalling limits**.

* usually recommended to small to medium companies

- this is just a **proxy**, **there is no MFA**, **you do not have to establish trust relationship**.

###### Simple AD

- use **Simple AD for low-cost, low-scale directory** with **Sambda 4-complatible applications**. It **does not support MFA!**. It cannot serve as a proxy to your on-prem AD. It is a standalone solution.

* there are **user and object limits**. Again, 2 versions: **small and large**. Basically **up to 5k users**.

##### When to use Federation

- when you are **within enterprise**. Big **companies** usually have their **own identity provider (Google, OKTA)...**.

* you have an **app that uses AWS**. With **Cognito you can create identity for users, even guest role!**

#### AWS SSO

- like Okta but AWS managed

* there are multiple identity sources available for SSO - AD, IdP or SSO itself. **You can have only 1 identity source within this service**. You cannot mix and match multiple ones

### CloudFormation

- **Resources section is mandatory**

#### Intrinsic Functions

- these can **only be used within specific parts of template**. These are:
  - resource properties
  - outputs
  - metadata attributes
  - update policy attributes

#### Stack Updates

- based of template diffing

* there are 3 levels of update behaviors: **with no interruption**, **with some interruption** and **replacement**.

- when using CloudFormation documentation, look for `update requirements` for a given property.

* you can specify the **UpdateReplacePolicy** to **ensure that no accidental deletions of the resources happen when CF changes**

##### Change Set

- proposed change for a specific stack. Others can be notified by SNS that such change set was setup.

* it clearly lists the changes that are to be executed (also shows update behavior) when change set is executed.

- you can either `execute` or `delete` it.

* you can have granular permissions - junior members proposing changes , senior members deleting / executing them.

#### CloudFormation Stack Set

- allows you to **easly deploy stacks to multiple places at once**

* you can **deploy stacks across accounts or OUs**

- you can **specify in which accounts / OUs to deploy resources into**

#### CloudFormation Change Set

**Change sets** allows you to **preview changes you made to your template before deploying the template for realz**.

- AWS will basically tell you: this will be modified, this will be deleted, this will be created...

* when you are ready **you can execute given change set to introduce the changes**

#### Nested Stacks

- when you **refference a stack as CF resource**

* you can **use nested stack outputs inside the root stack**

- **recommended** when you want to **isolate information sharing (outputs) to a group, use nested stacks**. **Otherwise** if you want to **share information, export outputs** from a stack.

#### Stack Roles

- **identity** who is creating given stack **needs permissions for the stack creation itself and any underlying resources that that stack creates**.

* you can also create a specific **stack role**. That role will be **used by CF while creating resources**. In such case the underlying identity needs only permissions to create the stack itself, not the resources that that stack might create.

- this allows for **separation of roles**. This works on the similarly to `change sets`. You can enable junior memebers to interact with the stack, modify some of the resources, but full managment of the stack can be off limits to them.

* there is also special **stack policy**. It **may seem like it's overlapping with _stack role_**. While this is **somewhat a case, you might want to allow users do 1 thing and the service itself other thing**.
  The **stack policy only applies on UPDATES**.

#### Capabilities

- depending on the value you provide, **it allows CloudFormation to create IAM resources**

* it **has nothing to do with Stack Persmissions**

- there are **two capabilities**: `CAPABILITY_IAM` or `CAPABILITY_NAMED_IAM`

* this setting exists to prevent you from creating dangerous IAM roles. It forces you to look at what's being created

#### Deletion Policy

- you can use **Snapshot** to **create a snapshot of data for services that support snapshots**

* you can use **retain** or also **delete**. There is no such thing as `force delete`.

- when you are dealing with s3, you should create a lambda function to empty the bucket first. This lambda function can be a **custom resource**.

#### Drift Detection

- allows you to detect changes compared to your reference template

* this is useful when you want to see which department has modified resources created by your template (which should not happen)

#### Custom Resources

- CF rather can creating a specific resource will **send event data to either lambda function or SNS topic**.

* CF expects that the remote entity **responds with a correct response**.

##### Lambda custom resources

- you would use this technique to deal with **import cycles** or creating resources with sdk quickly.

* very useful when you are dealing with a situation where a resource that you want to create
  does not have `CloudFormation` support yet.

#### Wait Conditions

- coordinate stack resource creation with something **that is external to your stack**.

* you could also track configuration process.

- used internally when using **creation policies**. The `creation-policy` is the preffered way of using `wait conditions`.

* it generates **presigned URL** which is used to **communicate that given resource is ready / was created**.

- you can **also use helper scripts (`cfn signal)** for **communication**.

#### Mappings

- this is a `key:value` structure that is designed to be used with `FindInMap` intrinsic function.

* usually used in a context of AMI ids per region.

#### Macros

- similarly to `Custom Resouces` you create a lambda function. This lambda function will be used to **modify the template itself**.

#### Helper scripts

- these are utilities provided by CloudFormation to **orchestrate, notify CF when changes happen to the underlying resources**

* mostly used by non-serverless components like EC2 (there you want your programs to report to CF during setup).

- there are multiple helper scripts
  - `cfn-hup`: **daemon which pools CF for changes**. If changed occured for a given CF block, runs scripts defined by you
  - `cfn-init`: used for **more complex user-init scripts**. Instead of writing scrips, you pass directives to a special program which is OS agnostic.
  - `cfn-signal`: used for `WaitConditions`, coordination between resources

#### Custom resource types

- you (or others) can create **custom CF resource types**

* this is something you **upload to CF registry**

- as a developer of such resource type, **you are responsible for implementing _create_, _delete_ and _update_ hooks**

### AWS Glue

- **serverless, fully managed EXTRACT TRANSFORM AND LOAD (ETL) service**

* automatically provides resources for you

- has an option to deploy a **crawler**. That crawler will **discover (scan) data and populate the Data Catalog**. This **data is usually s3**.

* with **Data Catalog** you can use **Athena / EMR / Redshift to query that catalog**.

- **can generate ETL code** but **only for Scala or Python**.

* has a **central metadata repository (data catalog)**.

- **ETL code** can be written using **Python or Scala**

* as the Load step of the ETL, you can load data to **Redshift, RDS, S3**

#### Job bookmarks

- A **checkpoint**, basically it tells the _AWS Glue_ what data it should process

* You can pick three values:
  - _Disabled_ (default): no checkpoint, every time you run the job, the whole data set is traversed
  - _Enabled_: only traverse the added files
  - _Pause_: you can pick files relatives to job, for example _I want to traverse files that the job 2 did not cover, unit and including job 4_

#### Glue materialized views

- allows you to replicate data from multiple data sources to another data source

* with materialized views **you no longer have to implemented an _ElasticSearch_ connector!**. Materialized views **work with DynamoDB and _ElasticSearch_**.

### AWS Trusted Advisor

- **helps** you to **optimize running costs**

* **checks security-related stuff (IAM)f** and **lists potential problems**

- **shows you** when given service is **using more than a service limit**

* what is very important is that **Trusted Advisor is more about accounts and IAM and unused resources**. This is something completely different than AWS Inspector which has to deal with EC2 instances mostly.

- you can use **AWS Service Quotas** to change **service limits** that you imposed. This is **something that trusted advisor alone cannot do!**.

* it has **reports on underutilized resources which can greatly help with cost savings**.

- you can refresh all or individual cheks. Checks might have different refresh intervals.

* you can check **service limits page**. This page is free for all, you do not have to have business or enterprise level of support

#### Notifications

- the summary can be sent on a **weekly basis as an email**. This is a feature **built-in**

* you can use **CloudWatch events** to listen on **Trusted Advisor** events. You **have to be in north-virginia**, otherwise you will not be able to create the rule.

##### Weekly emails

- you can setup **weekly email notifications**

* this is not something you would use to get a fresh set of data, but it still may be viable to you

##### Alerts

- you can create **alarms based on the _color_ of the resources**. In the console you can see checks that are _red_ or _yellow_

* Trusted Advisor **will automatically refresh checks every 24hrs**. If you need more up to date checks, you have to refresh them yourself

- the **alerts are powered by the CloudWatch metrics which Trusted Advisor exposes**

##### Getting fresh updates

- this involves running **lambda on a cron job, refreshing the _Trusted Advisor_ API**

* you can refresh the results every 15 minutes or so

- by **default** the Trusted Advisor **refreshes the checks every 24 hours**

#### Exposed keys

- with `Trusted Advisor` you can do the same thing with AWS keys that you did with `AWS Health` service.

* `Trusted Advisor` has the _exposed access keys_ tab within security pillar.

#### Tiers

- you do not have access to every pillar that `Trusted Advisor` is checking (Cost Optimization, Performance, Security, Fault Tolerance, Service Limits) if you are on a free plan.

* the **Cost Optimization, Performance is available for those with AWS support plan**

- if **you need all the checks** your account have to have **at least business support plan**

### KMS

- various ways of encryption, but mainly **Server-Side, Client-Side** encryption.

* with **Server-Side encryption** the **decryption also happens on the Server**. This is pretty logical since we would not know how to decrypt something.

- With **KMS** you **can create User Keys** but that process is not necessary. **Depending on the encryption model** you could use **AWS Managed Service Keys in KMS**.

#### Key Rotation

- **AWS Managed** keys are **rotated every 3 years**. You **cannot change** that option

* With **CMK** you can **opt into** rotating your keys **once per year**

- With **CMK and custom _key material_** you can **manually rotate the key whenever you want**

#### CMK

- **CMK** is the **root object used for any encryption within KMS**. CMK has an ARN.

* CMK **never leaves the service**. That means that CMKS **never leave the region**. **CMK CANNOT be exported**

- if you have appropriate persmissions - **resource policy on the CMK**, you can request KMS to encrypt and decrypt data using CMK.

* you can create **regional CMK alias** that points to given key.

- this key **usually not used for the encryption directly**. This is due to **size limit of 4kb that you want to encrypt**. For the encryption itself you should use **data encryption keys**.

* can be created based on imported `Key Material`. Once created, underlying `Key Material` cannot be changed, you would have to create a new key.

- when you **delete the CMK** you **will not be able to decrypt your data back**. The deletion operation is considered dangerous

#### CMK Grants

- allow you to **programatically delegate** the **use of CMK to other principals**. This is much better way of handling dynamic permissions than IAM.

* this action is only for **allowance**. Grants **cannot be used to explicitly deny access**.

- grants can be **easly revoked when neeeded**.

* you would **mainly use this for granting temporary access**

#### Key Material

- aka **Backing Keys**. They are used for creating CMK.

* you can **import your own key material**. You would do that if you have custom compliance requirements.

- they can be **rotated**. The **period** of the rotation **depends if it's AWS Managed (3 years) or Customer Manager (1 year)**

#### Data encryption Keys

- created based on a **request to KMS based on a SPECIFIC CMK**.

* you get **2 versions of this key**. One is **encrypted** and one is in a **plain text** format. You should **encrypt with plain text (discard after)**

- control **access to the CMKs** in KMS. While there are other methods to do so, **you need to use Key Policies if you want to control the access**

* the policies themselves are pretty similar to IAM policies

### Key policies

- the KMS key works a little bit differently in terms of IAM

* **to grant some entity ability to do something with a given key, you have to explicitly add that policy to a given keys policy upon the key creation / update**. So a situation: I create a key and I leave permissions as blank. Even though I might have administrative role, I will not be able to delete that key.

* usually the pattern here is to give `delete` / `PutPolicy` permissions to the `:root` (every entity in a given account), then `encrypt` / `decrypt` to lambdas and other applications

### Encryption context

- metadata you add while encrypting a piece of data

* you can **create IAM conditions based on that metadata**

- can be used **for tenant isolation purposes**. You can you **session policy with a condition for a given context key (eg. tenantId)**

### AWS IOT

- **regional** service

* devices described as _things_

- IOT devices have to be **registered** and then **can communicate with Device Gateway**

* **Device Gateway** creates **device shadows**. This is a **logcal representation of a device (it's state)**. Other **application can communicate with shadows**.

- IOT devices can also **read from their shadow**. This is quite useful when **an application want to send data TO a given device** eg. update it's state.

* **Device Gateway publishes to MQTT topic**. You can subscribe to this topic and create **IOT Rules**. Rules just like CloudWatch rules can have triggers and do stuff like **pushing to kinesis, adding to dynamodb etc..**

- to lower the cost you can use **Basic Indest**. This allows **your IOT device to skip going through the Topic and publish directly to a rule instead**.

#### IOT Analytics

- ingegrates with `Quicksight`.

* it has **different use-cases that Kinesis Analytics**. **Kinesis should be used for real-time or near real time stuff, that also applies to analytics**. With `IOT Analytics` data is stored long term.

### AWS Proton

- enables you to create **envioriment and service templates which are VERSIONED**

* **developers** can **deploy services based of those templates**

- you have the notion of **schema** which is **OpenApi spec which defines the set of inputs**

* this **kinda resambles CDK Constructs**

### QuickSight

- visualisation of data, mainly used for IOT, also for **streaming data**.

* you have to **sign up to it**. It's a bit separate in terms of AWS ecosystem.

- integrates with **Athena, Aurora, Redshift, S3, IOt**.

### CloudHSM

- **H**ardware **S**ecrue **M**odule

* can **generate keys** and **perform crypto operations**

- **KMS uses HSM under the hood**

* **presented as ENI inside VPC**

- if you **really need something that is dedicated ONLY to you**, the **cloud HSM will not cut it**.

#### Permissions

- since AWS only provides hardware here, it's **up to you to manage permissions**.

* there are only **CRUD like IAM persmissions available**. There is **no permission** to **deny key usage**.

#### HA

- you can **create HSM clusters**. These are **NOT HA by default**.

* you can make it HA, **spread it across multiple AZs**.

### Data Lake

- **a lot of data from a different services brought into one - usually s3**

* this is due to having to perform analytics on different sources of data - this is quite cumbersome

- aws integrates **Data-lake formation** which can help you with the creation of data-lakes.

### Migration

#### Migration Strategies

- **rehost (lift and shift)**: usually done using SMS for large legacy applications. Usually quick. After the migration part you might look into re-architecting the solution. This is due to the fact that re-architecting is usually easier within the cloud environment.

- **replatform (lift, tinker and shift)**: this is where you make a few optimizations (easy ones), like migrating your service to managed one like elastic beanstalk

- **repurchase (drop and shop)**: this is a decision to move to a different product

- **refactor / re-architect**: most expensive solution and quite dangerous one. Usually driven by a strong business need.

- **retire**

- **retain**

#### Migration Hub

- this is the **hub for discovery data and tracking the migrations**.

* you can **connect the discovery services to the hub** (like Discovery Service)

- is for checking the statuses and such. Does not perform the migrations / discovery itself.

#### Discovery Agent

- you **install this on your VMs**. It needs root access.

* **collects metrics data (networking, processes) etc** to help you plan the migration

- it is **not supported on all operating systems versions**

#### Discovery Connector

- **only for VMware**. Does not work for unsupported versions of linux and such.

* **when discovery agent is not supported** you can use agentless approach and get the **discovery connector to work**.

- this is basically **an OVA (runs virtual machine)** and collects the data

* when you cannot use Connector or Agent, **another approach would be to use import data via json template to migration hub directly**.

#### Discovery Service

- **contains data from connectors or agents**

* has **integration with Athena**. You **do not have to create special s3 bucket for this**.

- either **agetnless (Discovery Connector) (VMware IS REQUIRED!)** or **agent (does not have to be VMware)**

#### Server Migration Service (SMS)

- **THE recommended service** for migrating VMs. Minimizes downtime.

* you **actually need VMware / Hyper-V or Azure system to use it**.

- used to **actually migrate on prem VMware stuff**. SMS is **agentless**, there is no agent to download thus **SMS connector is needed aswell**.

* **CMK** is **regional**. This is worth knowing especially when encrypting EBS snapshots or EBS volumes.

- **replicates to an AMI** and can also **auto generate CloudFormation templates**

* when **migrating lots of servers** these can be **grouped into applications**. When doing so **SMS will generates AMIs, create CloudFormation templates** and **launch them in a coordiated fashion**.

- **applications can be divided into groups**. Groups can contain multiple servers.

* this service is **free of charge**. Remember that the snapshots, and the underlying created resources will probably cost you some $$.

#### Database Migration Service

- you can **setup custom schema mappings**

* uses **Schema Conversion Tool underneath**

- DMS enables you to **read and write to encrypted sources**. Data is **propagated in a decrypted form** but it **uses SSL for encryption in transit**.

* there are **3 steps for migration**:
  - **allocate a replication instance** which performs all the processes for the migration
  - **specify source and a target DB**
  - **create a task or set of task** to define **which tables and replication processess** you wan to use.

- the **replication instance should be created within DMS console**

* you can define **JSON transformation rules** or **Selection rules**. Basically **transforming and/or filtering** the **data when migrating**.

- you **do not have to create replication instance**. DMS **will do that for you (task)**

* you can **migrate multiple sources to one and vice versa**

- DMS is fast, and there might be a problem with your service throttling the migration by DMS (Good example would be Amazon ES).

##### Encryption

- **by default** DMS **encrypts your data at rest (replication instance)** using **KMS by default**. It can also **use CMK if you specify one**.

#### VM Import / Export

- an **alternative to SMS** which **does much less**.

* with this you can either **import VM as EC2 AMIs** or **import VM envioriment as EC2 instance**.

- you can also **export a VM that was previously imported**

* you can **import disks as EBS snapshots**

- **you should consider using SMS first**.

### VMware on AWS

- this enabled you to have **VMware stack on AWS infrastructure**.

* **can results** in a **networking complexicity**.

- you can use VMware on AWS to run **Oracle RAC** and have **Oracle backups on s3**.

### IDS / IPS Systems

These systems are used to **detect and prevent intrusions** from gettiing to your resources.

- usually **implemented** by **installing IDS agent on every instance within your VPC** or by creating a **reverse proxy layer** which **has IDP agents installed**.

### AWS Artifact

- provides **on demand access to AWS' compliance documentation**

* might be useful when your company is conducting an audit

### AWS Media (Live and Package)

- these tools are used for creatng **livestreaming**.

* combined **with cloudfront you can create distribution of HLS (HTTP Live Streaming)**

### AWS Global Accelerator

- this like **global pseudo-route53 with a sparkle of load balancer which allways returns 2 IPS for your costumers**. This is so that then user device caches the IP and disaster occurs it does not route to the service that is not working. The **returned IP is the same** but **the routing logic is within Global Accelerator**

* gives you **2 anycast IPS**

* **allows** you to create **sticky sessions (client affinity)**

- you **can** configure **health checks**. Health checks **can be PER endpoint (eg. ELB)**

* you can **configure endpoint weights (this is PER ELB basis for example)**

- you can configure **regional traffic dials** which enable you to route **based on weight you set for a global region**

* there is a notion of **endpoint groups**. This are **NLB/ALB/EC2 services OR static IP**

- GA **can use exisiting health checks that exist on ELB**

#### VS Route53

- **GA** has **multi-region failover** which **does not rely on DNS**. This is huge since **you do not have to worry about DNS cache**

* **GA** also **uses AWS backbone network**. This is the same technology as CloudFront. R53 does not work on that layer.

- **GA** is **a bit more expensive than R53** when it comes to **static costs**

#### CloudFront

- **GA** is **does not cache responses**. This is where CloudFront has clear advantage.

* **GA** does **not support Lambda @Edge**. Again, this is where CloudFront has clear advantage.

### AWS Guard Duty

- ingest from FlowLogs, R53, CloudTrail, AI **Thread detection in terms of ACCOUNTS** etc to a centralized place

* you can **invite other accounts to guard duty**. With this setting, Guard Duty will also **ingest from those invited accounts** (**if they accept**).

- it's usually **preferred for instances that have access to the internet**

* has **deep integration with CloudWatch**. You can use **CloudWatch events** to trigger lambda or sns when guard duty notices something.

- will detect if someone is mining bitcoin on your machines (since it's scanning flow/vpc logs)

### Amazon Inspector

- used for **security monitoring of EC2 instances**. Please keep in mind that it **only tackles security stuff, LEAVE STATE CHANGES TO AWS CONFIG!**

* scans for **vulnerabilities and not IAM issues**

- watches out for lack of best practices

* can be used **with or without an agent**. Inspector **will not launch instances for you**.

#### Assessments

- is the thing that **defines things that Inspector will be checking**

* _assessments runs_ can **generate reports** based on the findings.

- you can run **assessments on a schedule**. This is done using _CloudWatch_.

* **if you want to pick which instance** should the _Amazon Inspector_ be working on, you **should tag it**.
  You are then **able to pick the instance based on tags**

#### Uses

- the cool use of this technology I saw was around **validating golden AMI**. This can be done with automated pipeline.

### Health dashboard

- normal dashboard for AWS health is a bit too global. The `Health dashboard` is a dashboard which is presonalized for you.

* you can **subscribe to changes** using **CloudWatch events**.

- the service name is `Health` within `CloudWatch events`.

* this is also a great way of **knowing if your secret key was compromised (pushed to repo)**. This is because the `AWS Health` service tracks github repos.

- you can also track **scheduled change events** like **RDS maintenance** using **CloudWatch events**

### AWS WorkLink

- **AWS Managed VPN app but for mobile**

* WorkLink **makes sure** that the **data is never cached / stored on the real device**. This is great when you are worried that something might leak.

- you can **use your existing browser** but **the pages are rendered on AWS** and **send to you as SVG** 😲

### Device Farm

- enables you to **test your mobile application**, can **also be a website (selenium)**.

* you can literally connect to a real device and use it, like iphone X.

- this is **not a platform to perform load testing**. You should use **device farm for user testing**.

### AWS AppStream

- gives you the ability to **deliver desktop aps by browser** or **installable client**

* can be use to **sassify your product without worrying about piracy**

- you **can have old pc but still run instensive apps**. Everything is **proccessed by a fleet of instancess on AWS**

* you can have admins manage the app stream as it is a centralized service

- **users will always get the newest version of an app**.

### AWS Macie

- service **similar to AWS Inspector but for S3**.

* using **ML and CloudTrail logs it can notify you when something is wrong with your S3 bucket (or data inside that bucket)**

- it can also **tell you how is accessing sensible files most often**.

* very **useful for GDPR**, but it can also detect _ssh keys_ and other stuff

- can generate **dashboards** which report on **high-risk objects, user sessions etc**

* it has some **built-in alerting**. You can also **create your own alerts**.

### Mobile Hub

- aggregator of features tailored for mobile development

* you can click through UI and create Cognito pool, Database, SNS notification channels, REST API, all of that.

### Resource Groups

- by default AWS organizes stuff around services, like dynamodb or lambda

* you can create a resource group which organizes stuff as you want. You can have EC2 along with lambda listed within a console

- mainly for overview purposes, but you can **group by tags** so you there is a possibility for more sophisticated setups.

* you can attach IAM policies to resource groups.

### TCO

- this tool is used to **compare** the **cost of running in premise vs running in the AWS**

### Patterns

<!-- #### Connecting on-prem with VPC -->

#### Migrating from MongoDB

Since mongo is a nosql tech. one might choose **DocumentDB** or **DynamoDB**. It depends on the requirements and the question. Either way **not RDS!**.

#### Copying AMI between regions

You **can copy both EBS-backed or instance-store-backed AMIs**. **AMIs with encrypted snapshot can also be copied**. **IAM RELATED STUFF IS NOT COPIED**.

#### Copying AMI in the same region

This is mainly done to **encrypt root volumes**. You would create an AMI copy within the same region and encrypt the volume while doing the copying.

#### Restricting Access to s3 resources

Use `pre-signed` `s3` urls.
If you are distributing through `CloudFront` create `IAO` and associate that `IAO` with your distribution. Then modify permissions on `s3` to only allow OAI to access those files.

#### EBS and encrypting a volume on a running instance

So you would like to avoid downtime on your EC2 instance. **There is no direct way to change the encryption state of a volume or a snapshot**. You have to **either create an encrypted volume and copy data to it or take a snapshot, encrypt it, and create a new encrypted volume from the snapshot**,
Then you can either swap the volumes or restore volume from newly created snapshot.

#### EBS and Volume termination

There is a **DeleteOnTermination** attribute. **By default root EBS volume will be deleted on instance termination (they have the DeleteOnTermination protection tick on)**. This can be **changed using DeleteOnTermination attribute**. Other **non-root volumes will be NOT deleted by default**. Again **this can be changed by changing the DeleteOnTermination attribute**.

#### Cross VPC EFS

You can easily mount EFS on EC2 instances within a different VPC using **vpc peering or transit gateway**. Remember that you can **mount helper for encryption in transit**.

#### Last resort of Scaling RDS Horizontally (Sharding)

There is a couple of ways to scale the RDS. When you are out of options or it makes sense️™️ you might consider sharding. **Sharding is the process of splitting the database into smaller database and using some kind of router to route queries to specific databases**. As you can image this can be quite admin-heavy process when it comes to AWS.

#### Adding Encryption to an RDS db

**You cannot add encryption to an existing RDS db**. What you have to do is to **create db snapshot and encrypt it**. From that snapshot you can create a copy of your DB.

#### Custom metrics on EC2

**By default** CloudWatch monitors **CPU, Disk and Network**. If you need RAM metrics for example you can **install CloudWatch agent on a EC2** which will **push data to custom CloudWatch metrics**.

#### Troubleshooting EC2 instances in ASG

First thing you need to do is to **place the instance in a standby state**. When in **standby state**, the instance will be **detached from ELB and target group, it is still part of ASG though**. If you do not want ASG to continue the scaling processes you can **suspend ASG scaling processes**. Keep in mind that you are **still billed for the ec2 that are in standby state**

#### ELB pass-through

When you want to achieve pass-through on load balancers you usually want end-to-end encryption when it comes to traffic. Normally, traffic is decrypted at load-balancer level. But sometimes you need end-to-end encryption for compliance reasons. Then **you need to use either classic LB or NLB**. This is because these have some awareness on **layer 4**. This is important since **when you want end-to-end encryption you need to use TCP instead of HTTPS**.
That means that **ALB is not capable of handling that traffic** since it only knows how to handle layer 7 stuff.

#### Terminating SSL certs on the webservers

This one is not that popular but still might be revelant. Normally when you setup a HTTPs listener on ALB the SSL in terminated there. We do not want that.

What you can do is to **setup NLB with TCP listeners** and **put webservers behind that NLB**. This will make it so that the traffic is in a passthrough mode (**as descrbied above**). The certs then can be handled at the webservers.

Another approach you might take is to **use Route53 with multivalue routing**.

#### Getting Source IP with ELB

With **ALB/Classic** you can use **X-Forwarded-For** header. But beware that **this only works for HTTP/HTTPs**. When you are dealing with **eg. TCP** you should **enable Proxy Protocol if you are using Classic ELB (NLB uses proxy protocol by default)**. It **prepends info about clients IP before the actual data**. **Proxy Protocol cannot be added using a console. YOU HAVE TO USE CLI**

Please also **keep in mind that NLB just forwards the traffic**. So **This is not a problem with NLB**. There is a good article on how to make it HA.
https://aws.amazon.com/blogs/security/how-to-add-dns-filtering-to-your-nat-instance-with-squid/

#### Egress HTTP URL rules (DNS Filtering)

Within the AWS there are multiple tools to control egress traffic. You have route tables, security groups, NACLs, Ingress only IGW and so on.
But what happens when you want to control the HTTP URL traffic? Well to do so you should implement a **proxy server**. This server would be reponsible for
for **dns filtering**.

For the implementation, the **aws does not provide such tool out of the box**. You probably should **look into Squid (open source)**. You can **install Squid on NAT Instance** but be aware that **this solution is not HA by default**.

#### Accessing VPC Endpoints

Remember that **VPC Endpoints can only be accessed inside the VPC**. That means that if you have a VPN connection or Direct Connect to your VPC you have to use some kind of proxy. Usually you would use **EC2 to be a proxy for your s3 requests**

#### EC2 reports unhealthy but no ASG event takes place

First of all remember that ASG determines the health of the instance using various factors:

- status checks provided by EC2 itself (default way of checking)
- health checks provided by ELB
- custom health checks

Now, there is a notion of **grace period on ASG**. This is the time it takes for your instances to boot up basically. **There may be point of time where health checks finish before the grace period**. This will mark your instance as unhealthy BUT **ASG will take no action until the grace period is over**.

#### I lost my SSH Key, what now.

There are 2 ways you can tackle this problem. First is using **Systems Manager Automation** with the **AWSSupport-ResetAccess document** to create a **new SSH key**.

The **other one** requires you to **stop the instance**, **detach the root volume** and **modify `authorized_keys file**.

#### Enabling SSH with SG and NACL

#### Changing instance type inside ASG

**YOU CANNOT EDIT EXISTING LAUNCH CONFIGURATION**. You have to create a **new launch configuration with new instance type**. To make sure that all of your instances are using this new launch configuration **terminate old instances**. New one will get added using new launch configuration.

**Instead of creating new launch configuration** you **can** also **suspend the scaling process** and **restart existing instances after specificizing new instance type**. This is also a solution but seems pretty meh tbh.

#### ELB and Route53

You should use simple `ipv4` and alias.

#### Different IAM roles per container on ECS or Fargate

Since **task definitions allow you to specify IAM roles** you should edit those and setup the IAM roles correctly. The launch type does not matter in this case as **either EC2 or Fargate launch type** support that option.

#### EC2 immediately going from pending to terminated on restart

This may be due to these 4 reasons:

- you've **reached your EBS volume limit**
- **EBS snapshot is corrupt**
- root **EBS is encrypted** and you **do not have permissions to access KMS key**
- **AMI is missing some required parts**

#### Migrating VPC to IPv6

There are a couple of steps but always remember about key player here: **Egress-only IGW**.

- route public traffic to IGW

* route private traffic to Egress-only IGW, remember

- add IPv6 CIDR to your subnets / VPC

#### Exposing your dockerized microservices through APIGW

APIGW is great. It allows for caching, throttling, rate limiting etc. You would not want to implement these features all over again within your ECS / EKS cluseters. You can expose your microservices using APIGW and ALB or NLB (depending on the requirements).

All you have to do is put APIGW in front of the load balancer and set the **integration as HTTP proxy**. (this is the default serverless framework method). With PROXY integration you can set endpoint URL as the ALB dns name.

#### Disaster Recovery on AWS

##### RTO (Recovery time objective)

This is basically the period of time your service CAN be unavailable, the downtime.

##### RPO (Recovery point objective)

How much data can we lose? This is the duration in-between data snapshots, backups are made.

##### Backup & Resotre

This is the **low-cost** solution. You mainly use S3 as a backup service and then you start new services using those backups. Nothing is on standy.
This **technique is used mainly where RTO and RPO is measured in hours**

##### Pilot Light

This is **also low cost but more expensive than Backup & Restore**. This is where you have **only the most critical parts (eg. a database) of your system proviosioned somewhere**. It is running and **that part of the system WILL be scalled to support your whole load when disaster occurs**.
This **technique used mainly where RTO and RPO is measured in 10s of minutes**

##### Warm Standby

This is where **your whole system is in place, just SCALLED DOWN**. Recovery is pretty similar to Pilot Light where **you scale up the standy to accomodate for prod traffic**.

##### Multi-Site

This is the **active-active approach (the most expensive one)**. You basically **have 2 prod envs active at all times**. With that you usually have **R53 in front, you can quickly switch the traffic**.

### Other

- **blue / green deployment** is where you switch between one version. Used to move traffic between 2 versions of our stack (blue and green). You can freely switch between them. This is **usually conducted by Route53 weighted policy (1 and 0)**

#### Proxy

You can think of a `man-in-the-middle` when someone is talking about proxies. Sometimes `proxy` will speak on your behalf, sometimes it will just handle the initial connection. Proxy types:

- **Half Proxy**: this type handles the initial setup. It initializes the connection to the backend, proxy passes that to the client. After that it just forwards the traffic, it may do some NAT work but nothing more sophisticated going on here. It's usually useful for discovery purposes, like checking at the connection state where to route the client, then letting the requests pass.

- **Full Proxy**: this type is the `mediator` between the backend and the client. It establishes a connection with a backend and any client requests stop on the proxy, they do not go directly to the backend (or through proxy to the backend). This pattern usually means that the `proxy` is doing some sophisticated work. This is where **ALB and NLB** are. They are full proxies.

#### Terminology

- **resource contention** is where there is a **conflict over access to a shared resource**.

### Whitepapers

#### AWS DDoS Resiliency

First of all you should think about **minimizing the blast area**. This has to do with having most of you infrastructe within a private subnets.

Next, your architecture should be able to **absorb the DDOS attack**. As weird as this might seem, with scallable architecture, you will have more time to think and adjust during the attack.

- **Autoscalling and ELB**
- **Utilize CloudFront**
- use the **Enhanced Networking** when using EC2 instances

Next, think about **safeguarding exposed & hard to scale resources**. There are a few tools which enable you to do that. **R53 with private DNS records**, **CF with OAI and Georestrictions** and finally **WAF for filtering traffic**.

### Notes from the Exam Rediness: DevOps

- how IAM integrates with CodeCommit?

- CodeCommit integration with CloudWatch Events (triggering CodeBuild)

- how would Jenkins would integrate with CodeCommit? (fault tolerance and scalability)

- buildspec and appsec file

- CodeDeploy deployment configuration (4 settings)

- chaining deployments - 1 deployment published to "beta" bucket which kicks off another deployment. If it fails (lambda can test the deployment), we rollback

- cross account pipelines

- deployments

  - in place
  - rolling updates
  - rolling updates: canary
  - blue green
  - red black
  - immutable

- CodeCommit supports pre-commit hooks

- CodeBuild does not handle artifacts protection. You have to eg. add default encryption to s3 bucket where you are going to store the artifacts

- Jenkins EC2 plugin, CodeBuild plugin

- Code Star (template.yaml)

- Update policy for ASG

- WaitConditions and Signals (CF), CreationPolicy

- StackPolicy

- StackSets

- CF helper scripts (cfn-init, cfn-hup). With long cfn-hup interval EC2 deployments(subsequents) might take a while

- EB envioriments (and .ebextensions/config.yaml file)

- OpsWorks Stacks

- Rollback UPDATE_FAILED - resources probably changed outside of CloudFormation

- EB export vs Docker. You cannot lift and shift with EB (unless you are using Docker)

- ELB not allowing you to switch users between ASGs based on weight

- ELB monitoring

  - SurgeQueueLenght - backend systems are not able to process requests (backpressure)
  - SpilloverCount - when the SurgeQueue overflows

- CloudWatch monitors CI/CD services (_CodePipeline_)

- CloudFormation logs (?)

- cross account access (assume role), CI/CD via CodePipeline

- data protection in transit

- host firewalls (?)

- S3 encryption. Glacer encrypts by default

- CloudWatch consuming CloudTrail logs

- AutoScaling lifecycle hooks and policies

- standby states for ec2 instances within asg

- automation vs run script

- CodeDeploy revision and rollbacks (retain the content)
