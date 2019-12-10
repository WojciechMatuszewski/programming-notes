# Aws Solutions Architect Stuff

Just me trying to learn for an exam 🤷‍♀

## Acloudguru & Linux Academy

### AWS and SA Fundamentals

#### Access Management

-   **Principal** is a person or application that can make authenticated or anonymous request to perform an action on a system. **Often seen in-code in lambda authorizers**

*   Security **in the cloud** is **your job**

-   Security **of the cloud** is the **AWS job**

### Basics

-   **AZ (Availability Zone)** is a distinct location within an AWS Region. Each
    Region comprises at least two AZs

*   **AWS Region** is a geographical area divided into AZs. Each region counts as
    **separate** geographical area.

-   **Virtual Private Cloud (VPC)** is a virtual network dedicated to a single AWS
    account. It's logically isolated from other virtual networks in the AWS cloud

*   **EFS and S3** are popular storage options

-   **Cloudfront** is a CDN

### Identity Access Management & S3

#### IAM Basics

> IAM allows you to manage users and their level of access to the AWS Console

-   **IAM** is universal, does not apply to regions

*   **Identity Federation** allows your users to login into AWS Console using 3rd
    party providers like Google etc..

-   Supports **PCI DSS** framework. This is some kind of standard for security

*   **Policies = Permissions**, written in JSON

-   **Roles** enable one AWS service do something / interact with another. For
    example virtual machine (EC2) interacting with AWS storage.

*   **Root Account**: email address you first sign up to AWS with. This account
    basically has a godmode and can do everything in the console. That's why you
    pretty much never want someone to login on root account, just like in Linux.

-   Users can have **programmatic access** to AWS console. This basically allows
    you to pass access key and secret key so that you can interact with developer
    tools

*   Users **can be added to groups**. These groups **can have policies** assigned
    to them.

-   Policies have **different types**. Like `Job function` or `AWS managed`.

#### S3 Basics

> S3 stands for **Simple Storage Service**

-   S3 is an **object storage**, it allows you to upload files.

*   Data is spread between multiple devices and facilities

-   There is **unlimited** storage. Maximum file-size is 5 TB though.

*   By **default** uploaded objects **are NOT public**. You have to set them as public manually

-   S3 has an **universal namespace**. That means that Bucket names has to be
    unique globally.

*   Object consists of:

    -   Key: simply the name of the object
    -   Value: data (bytes)
    -   Version ID: **S3 allows you to have multiple versions of a file**
    -   Metadata
    -   Sub-resources

-   When it comes to consistency:

    -   **Read after Write** for **PUTS**. Basically you can read immediately after
        you write to a bucket

    -   **Eventual Consistency** for **overwrite PUTS and DELETE**. Basically if you
        delete something or override it, it takes a second or two for you changes to
        propagate and take an effect.

*   Buckets can be **replicated to another account or to a different bucket
    (region needs to differ)**.

-   S3 can be `accelerated`. There is something called **S3 Transfer
    Acceleration** where users upload to **edge locations** instead directly to
    the Bucket. Then that uploaded object is **transferred automatically to your
    bucket**.

*   S3 is a **tiered** storage

    -   S3 Standard, stored across multiple devices and multiple facilities

    -   S3-IA/S3 One Zone-IA (**Infrequent Access**): for data that is accessed less
        frequently but requires rapid access when needed

    -   S3 Glacier / Glacier Deep Archive: used for data archiving, where you would
        keep files for a loooong time. Retrieval time is configurable (**Deep
        Archive is locked on 12hr retrieval time though**)

##### Versioning

-   S3 have the notion of versioning: **stores all versions of an object including all writes even if you delete it!**

*   After **enabling** versioning, it **cannot be disabled**, only suspended

-   Remember that with versioning your previous versions persists on the bucket. This can lead to exponential growth of size when editing big files.

*   When an object is deleted, **bucket may seem empty** but that's not the case. You just placed a _delete marker_ on that object (thus creating a new version). Your **previous versions are still there!**.

-   You can restore your deleted objects by **deleting a delete marker**.

##### Life-cycle rules

-   can be used to **transition objects to different TIER of storage after X amount of time**, this can be placed on current and previous versions.

*   With life-cycle there is notion of expiration. Basically your objects can be **expired after X amount of time**. Object can be in a _expired state_ where it will wait to be permanently deleted.

-   Can be used with **conjunction with versioning**.

##### Storage Gateway

-   Physical/virtual device which **will replicate your data to AWS**.

*   There are 3 flavours of Storage Gateway
    -   **File Gateway** : used for storing files as object in S3.
    -   **Volume Gateway**: used for storing copies of hard-disk drives in S3.
    -   **Tape Gateway**: used to get rid of tapes.

### Snowball

-   Big briefcase, **up to 100TB** of storage. Used to move data from one point to
    another fast

### Snowmobile (not joking)

-   A truck with a container that carries snowballs.

### CloudFront

-   CloudFront is a **CDN**. Takes content that exists in a central location and distributes that content globally to caches.

*   These caches are located to your customers as close as possible.

-   Origin is the name given for the thing from which content **originates from**,
    can be an S3 bucket, web-server or other AWS services.

*   Origin has to be accessible to the internet

-   Edge locations **cache content** (TTL)

*   Distribution is basically the **collection of Edge Locations**

-   You **can** invalidate cache content

*   There are **2 types of distributions**.
    -   Web
    -   RTMP (used for video streaming and such)

-   When you deploy CloudFront distribution your content is automatically deployed to edge location. You can specify which ones (limited to a country). If you are rich you can deploy to all edge locations

*   **Cache hit** means that when an user requested a resource (like a webpage), an edge location had that available

-   **Regional cache** is like a meta edge location. Basically second level cache, fallback when there is **no cache hit**. If it does not have a copy of a given content it **falls back to the origin** (origin fetch).

*   **By default** every CloudFront distribution **comes with a default domain name**. That domain of course works for HTTP and HTTPS. You can register domain and replace it.

-   You can restrict the access on two levels (**You can** restrict an access to S3 only but it's no the topic of CloudFront):

    -   on a CloudFront level, your bucket is still accessible though
    -   on a S3 and CloudFront level, you can only access the website using signed urls.

-   Restricting your CloudFront & S3 combo is done by creating **OAI**.

*   **OAI** is an _identity_. That _identity_ can be used to restrict access to you S3 bucket. Now whenever user decides to go to your bucket directly they will get 403. To achieve such functionality you add **CloudFront as your OAI identity**

### EC2 (Elastic Compute Cloud)

-   resizable compute capacity in the cloud, **virtual machines in the cloud**

*   different pricing models:
    -   **on demand**: pay per time use (per-hour, per-second)
    -   **reserved**: capacity reservation, contract with AWS
    -   **spot**: like a market, but for instances, when AWS has free capacity you can bid and buy, **when capacity is needed they will be taken away from you. There are mechanisms which will alert you if that happens!**
    -   **dedicated**: psychical machines **only for you**. Mainly used when you have strict licensing on software you are using

-   There are different _health checks_:

    -   **System Status Check**: this checks the underlying hyperviser (virtualization tool)

    -   **Instance Status Check**: checking the EC2 instance itself

*   **I**nput **O**put **P**er **S**econd (IOPS): basically how fast the drive is

-   Termination protection is turned off by default

#### Security Groups

-   **changes** to security group **are instant**

*   security groups **are stateful**. That means when you **create an inbound rule, outbound rule is created** automatically.

-   there are **no deny rules**. Security groups can only **specify allow rules**.

*   EC2 can have more than 1 security group attached to it, also you can have multiple EC2s assigned to 1 security group.

-   **All inbound traffic is blocked by default**

#### EBS (Elastic Block Store)

-   basically **virtual harddisk in the cloud**

*   persistent storage

-   **EBS root volume CAN be encrypted**

*   **automatically replicated** within it's own AZ

-   Different versions:

    -   **Provisioned IOPS** - the most io operations you can get (databases), most expensive
    -   **Cold HDD** - lowest cost, less frequently accessed workloads (file servers)
    -   **EBS Magnetic**- previous generation HDD, infrequent access
    -   **General Purpose**

#### EBS vs Instance Store

-   Instance Store a.k.a **Ephemeral storage**. The data lives on a rack where your virtual machine is booted. **If you reboot it, there is HIGH CHANCE that you will loose that data**. It's not 100% guaranteed though. Your EC2 can always be assigned to the same rack, but that is unlikely.

-   Instance Store is not really persistent, whereas EBS is a persistent, multi AZ storage option.

### CloudWatch
