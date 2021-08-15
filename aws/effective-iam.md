# Effective IAM

Taking notes while reading the [Effective IAM book](https://www.effectiveiam.com/)

## Control Access to Any Resource

- Use the latest `Version` of the policy document.
  This is a good recommendation. I remember debugging a problem where some feature of the IAM policy language was not working.
  The reason was my not specifying the version, thus missing the feature I needed completely.
  On the side-note, **if you do not specify the policy version, the `2008-10-17` version is applied**. This version might not include features you rely on like **variables**. For more info please [refer to this documentation page](https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies_elements_version.html).

- Some of the policy elements **are required but are inferred in some contexts, thus are really required ONLY in specific scenarios**. Let us take a policy attached to the IAM user or role.
  In such cases, the `Principal` field is inferred, thus you do not have to specify this field. The [_JSON policy document structure_ of this documentation page elaborates](https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies.html#policies_id-based)

  > If you are creating an IAM permissions policy to attach to a user or role, you cannot include this element. The principal is implied as that user or role.

- The _AWS Managed Policies_ might not be a good choice. Since they have to be usable in every customer account, **the managed policies usually specify wildcards at the resource level**. Not good for the security perspective.

- **Prefer** starting with **denying every action for every identity** and then **loosening the strictness for a hand picked identities**.
  This way you lessen the chance of making an error and making the permissions to permissive.

- Use conditions. They exist, they are very useful.

## Why AWS IAM is so difficult to use

- AWS IAM has many features, so might say that there is too many features.

- The way the permissions are resolved is complex. For example, one might grant access to the bucket via _Resource Policy_ or _Identity Policy_.

- Some services, like S3 has service-specific access control systems - in the S3 case _S3 Object ACLs_.

- Achieving least privilege requires deep knowledge of AWS itself. Since AWS is vast, it's a quite a big ask.
