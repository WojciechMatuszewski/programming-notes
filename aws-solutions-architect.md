# Aws Solutions Architect Stuff

Just me trying to learn for an exam 🤷‍♀

## Acloudguru & Linux Academy

### AWS and SA Fundamentals

#### Access Management

- **Principal** is a person or application that can make authenticated or anonymous request to perform an action on a system. **Often seen in-code in lambda authorizers**

- Security **in the cloud** is **your job**
- Security **of the cloud** is the **AWS job**

### Basics

- **AZ (Availability Zone)** is a distinct location within an AWS Region. Each
  Region comprises at least two AZs

- **AWS Region** is a geographical area divided into AZs. Each region counts as
  **separate** geographical area.

- **Virtual Private Cloud (VPC)** is a virtual network dedicated to a single AWS
  account. It's logically isolated from other virtual networks in the AWS cloud

- **EFS and S3** are popular storage options

- **Cloudfront** is a CDN

### Identity Access Management & S3

#### IAM Basics

> IAM allows you to manage users and their level of access to the AWS Console

- **IAM** is universal, does not apply to regions

- **Identity Federation** allows your users to login into AWS Console using 3rd
  party providers like Google etc..

- Supports **PCI DSS** framework. This is some kind of standard for security

- **Policies = Permissions**, written in JSON

- **Roles** enable one AWS service do something / interact with another. For
  example virtual machine (EC2) interacting with AWS storage.

- **Root Account**: email address you first sign up to AWS with. This account
  basically has a godmode and can do everything in the console. That's why you
  pretty much never want someone to login on root account, just like in Linux.

- Users can have **programmatic access** to AWS console. This basically allows
  you to pass access key and secret key so that you can interact with developer
  tools

- Users **can be added to groups**. These groups **can have policies** assigned
  to them.

- Policies have **different types**. Like `Job function` or `AWS managed`.

#### S3 Basics

> S3 stands for **Simple Storage Service**

- S3 is an **object storage**, it allows you to upload files.

- Data is spread between multiple devices and facilities

- There is **unlimited** storage. Maximum file-size is 5 TB though.

- S3 has an **universal namespace**. That means that Bucket names has to be
  unique globally.

- Object consists of:

  - Key: simply the name of the object
  - Value: data (bytes)
  - Version ID: **S3 allows you to have multiple versions of a file**
  - Metadata
  - Sub-resources

- When it comes to consistency:

  - **Read after Write** for **PUTS**. Basically you can read immediately after
    you write to a bucket

  - **Eventual Consistency** for **overwrite PUTS and DELETE**. Basically if you
    delete something or override it, it takes a second or two for you changes to
    propagate and take an effect.

- Buckets can be **replicated to another account or to a different bucket
  (region needs to differ)**.
- S3 can be `accelerated`. There is something called **S3 Transfer
  Acceleration** where users upload to **edge locations** instead directly to
  the Bucket. Then that uploaded object is **transferred automatically to your
  bucket**.

- S3 is a **tiered** storage

  - S3 Standard, stored across multiple devices and multiple facilities

  - S3-IA/S3 One Zone-IA (**Infrequent Access**): for data that is accessed less
    frequently but requires rapid access when needed

  - S3 Glacier / Glacier Deep Archive: used for data archiving, where you would
    keep files for a loooong time. Retrieval time is configurable (**Deep
    Archive is locked on 12hr retrieval time though**)

### Snowball

- Big briefcase, **up to 100TB** of storage. Used to move data from one point to
  another fast

### Snowmobile (not joking)

- A truck with a container that carries snowballs.

### CloudFront

- CloudFront is a **CDN**. Takes content that exists in a central location and distributes that content globally to caches.

* These caches are located to your customers as close as possible.

- Origin is the name given for the thing from which content **originates from**,
  can be an S3 bucket, web-server or other AWS services.

* Origin has to be accessible to the internet

- Edge locations **cache content** (TTL)

* Distribution is basically the **collection of Edge Locations**

- You **can** invalidate cache content

* There are **2 types of distributions**.
  - Web
  - RTMP (used for video streaming and such)

- When you deploy CloudFront distribution your content is automatically deployed to edge location. You can specify which ones (limited to a country). If you are rich you can deploy to all edge locations

* **Cache hit** means that when an user requested a resource (like a webpage), an edge location had that available

- **Regional cache** is like a meta edge location. Basically second level cache, fallback when there is **no cache hit**. If it does not have a copy of a given content it **falls back to the origin** (origin fetch).

* **By default** every CloudFront distribution **comes with a default domain name**. That domain of course works for HTTP and HTTPS. You can register domain and replace it.

- You can restrict the access on two levels (**You can** restrict an access to S3 only but it's no the topic of CloudFront):

  - on a CloudFront level, your bucket is still accessible though
  - on a S3 and CloudFront level, you can only access the website using signed urls.

- Restricting your CloudFront & S3 combo is done by creating **OAI**.

* **OAI** is an _identity_. That _identity_ can be used to restrict access to you S3 bucket. Now whenever user decides to go to your bucket directly they will get 403. To achieve such functionality you add **CloudFront as your OAI identity**
