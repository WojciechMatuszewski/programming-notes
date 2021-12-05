# Effective IAM

Taking notes while reading the [Effective IAM book](https://www.effectiveiam.com/)

## Control Access to Any Resource

- In IAM, it's crucial to understand the **notion of the _principal_**. The _principal_ is an entity **authenticated by AWS and assigned privileges within AWS**.

- Use the latest `Version` of the policy document.
  This is a good recommendation. I remember debugging a problem where some feature of the IAM policy language was not working.
  The reason was my not specifying the version, thus missing the feature I needed completely.

  On the side-note, **if you do not specify the policy version, the `2008-10-17` version is applied**. This version might not include features you rely on like **variables**. For more info please [refer to this documentation page](https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies_elements_version.html).

- Some of the policy elements **are required but are inferred in some contexts, thus are really required ONLY in specific scenarios**. Let us take a policy attached to the IAM user or role as an example.

  In such cases, the `Principal` field is inferred, thus you do not have to specify this field. The [_JSON policy document structure_ of this documentation page elaborates](https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies.html#policies_id-based)

  > If you are creating an IAM permissions policy to attach to a user or role, you cannot include this element. The principal is implied as that user or role.

- The _AWS Managed Policies_ might not be a good choice. Since they have to be usable in every customer account, **the managed policies usually specify wildcards at the resource level**. Not good for the security perspective.

- **Prefer** starting with **denying every action for every identity** and then **loosening the strictness for a hand picked identities**.
  This way you lessen the chance of making an error and making the permissions to permissive.

- Use conditions. They exist, they are very useful.

- If you can, **leverage resource policies**. The "closer" the policies are to the actual resource you want to protect, the better.

- In **cross-account** scenarios, you need to allow **actions on the resource and identity levels**. I remember the first time stumbling upon this. I was a bit confused. Of course, everything is in the docs :)

## Why AWS IAM is so difficult to use

- AWS IAM has many features, some might say that there is too many features.

- The way the permissions are resolved is complex. For example, one might grant access to the bucket via _Resource Policy_ or _Identity Policy_.

  - What is more, the **_Resource Policy_ can grant access even though the _Permissions Boundary_ explicitly denies it**.

  - The famous policy evaluation diagram does not account for every possible scenario. For example, the _S3 Object ACLs_ are not considered.

- Some services, like S3 has service-specific access control systems - in the S3 case _S3 Object ACLs_.

  - Thankfully AWS offered an alternative. As of late 2021 people can use **Object Ownership**.

- Achieving least privilege requires deep knowledge of AWS itself. Since AWS is vast, it's a quite a big ask.

## Scale with security domains and a control loop

- AWS Accounts could be used to create secure boundaries around resources.

- As a rule of thumb: **create an AWS account for each major use case**. That might be a new team or a particular microservice.

- Use AWS OUs for grouping accounts together.

- After a while, you might want to start splitting your accounts even further, looking at the _delivery methods_ (like dev, preprod, prod).
  Having such granularity will help you will cost management.

- Use SCPs. They are great for governing global capabilities.

- Use tools for generating policies. For example, to ensure that your organization can only use services that are PCI compliant, instead of manually checking, use the `aws-allowlister` or similar tools.

## Simplify AWS security by using the best parts

This chapter contains concrete recommendations and features you should focus on while implementing least privilege.

- **Separate use cases using AWS accounts**. As eluded earlier, AWS accounts are great for creating secure boundaries.

- Use **policy condition keys**. The `Condition` section is very powerful. Use it!

- **Control access to data with resource policies**. Resources that might be shared between accounts usually allow you to define a _resource policy_ on them. Use those! Instead of dealing with X roles and users, you know only need to be concerned with concrete resources.

- **Take advantage of KMS, it's integrated with a lot of services**. The KMS service exposes you two ways of creating and managing keys.
  I'm talking about **AWS managed CMK** and **Customer managed CMK**.

  - In theory, using KMS **, one might re-create the behavior of resource policies on a resource that does not support them!**. To achieve this, **encrypt the data with custom CMK and apply a policy on the key itself**.

- By using KMS, you can effectively retrofit IAM resource permissions for the service that does not support that feature.
  All you have to do is to encrypt the data service is holding / reading, then apply the key policy on the CMK and you are done.

- **Use IaC for controlling the IAM related things (and effectively all your cloud resources)**. This point does not need explanation.

## Understand what your policies actually do

- There are three AWS-native tools that can help you with that. Look into `AWS Policy Validator`, `AWS Access Analyzer` and `AWS IAM simulator`.
  The `AWS IAM simulator` exposes an API, neat!
