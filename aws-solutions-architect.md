# Aws Solutions Architect Stuff

Just me trying to learn for an exam 🤷‍♀

## Acloudguru & Linux Academy

### AWS and SA Fundamentals

#### Access Management

- **Principal** is a person or application that can make authenticated or anonymous request to perform an action on a system. **Often seen in-code in lambda authorizers**

* Security **in the cloud** is **your job**

- Security **of the cloud** is the **AWS job**

### Basics

- **AZ (Availability Zone)** is a distinct location within an AWS Region. Each
  Region comprises at least two AZs

* **AWS Region** is a geographical area divided into AZs. Each region counts as
  **separate** geographical area.

- **Virtual Private Cloud (VPC)** is a virtual network dedicated to a single AWS
  account. It's logically isolated from other virtual networks in the AWS cloud

* **EFS and S3** are popular storage options

- **Cloudfront** is a CDN

### Identity Access Management & S3

#### IAM Basics

> IAM allows you to manage users and their level of access to the AWS Console

- **IAM** is universal, does not apply to regions

* **Identity Federation** allows your users to login into AWS Console using 3rd
  party providers like Google etc..

- Supports **PCI DSS** framework. This is some kind of standard for security

* **Policies = Permissions**, written in JSON

- **Roles** enable one AWS service do something / interact with another. For
  example virtual machine (EC2) interacting with AWS storage.

* **Root Account**: email address you first sign up to AWS with. This account
  basically has a godmode and can do everything in the console. That's why you
  pretty much never want someone to login on root account, just like in Linux.

- Users can have **programmatic access** to AWS console. This basically allows
  you to pass access key and secret key so that you can interact with developer
  tools

* Users **can be added to groups**. These groups **can have policies** assigned
  to them.

- Policies have **different types**. Like `Job function` or `AWS managed`.

#### Roles

- you should **prefer attaching roles** instead of using aws credentials

* roles can be attached to many services

- roles can be used **in any region**, they are universal

#### S3 Basics

> S3 stands for **Simple Storage Service**

- S3 is an **object storage**, it allows you to upload files.

* Data is spread between multiple devices and facilities

- There is **unlimited** storage. Maximum file-size is 5 TB though.

* By **default** uploaded objects **are NOT public**. You have to set them as public manually

- S3 has an **universal namespace**. That means that Bucket names has to be
  unique globally.

* Object consists of:

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

* Buckets can be **replicated to another account or to a different bucket
  (region needs to differ)**.

- S3 can be `accelerated`. There is something called **S3 Transfer
  Acceleration** where users upload to **edge locations** instead directly to
  the Bucket. Then that uploaded object is **transferred automatically to your
  bucket**.

* S3 is a **tiered** storage

  - S3 Standard, stored across multiple devices and multiple facilities

  - S3-IA/S3 One Zone-IA (**Infrequent Access**): for data that is accessed less
    frequently but requires rapid access when needed

  - S3 Glacier / Glacier Deep Archive: used for data archiving, where you would
    keep files for a loooong time. Retrieval time is configurable (**Deep
    Archive is locked on 12hr retrieval time though**)

##### Versioning

- S3 have the notion of versioning: **stores all versions of an object including all writes even if you delete it!**

* After **enabling** versioning, it **cannot be disabled**, only suspended

- Remember that with versioning your previous versions persists on the bucket. This can lead to exponential growth of size when editing big files.

* When an object is deleted, **bucket may seem empty** but that's not the case. You just placed a _delete marker_ on that object (thus creating a new version). Your **previous versions are still there!**.

- You can restore your deleted objects by **deleting a delete marker**.

##### Life-cycle rules

- can be used to **transition objects to different TIER of storage after X amount of time**, this can be placed on current and previous versions.

* With life-cycle there is notion of expiration. Basically your objects can be **expired after X amount of time**. Object can be in a _expired state_ where it will wait to be permanently deleted.

- Can be used with **conjunction with versioning**.

##### Storage Gateway

- Physical/virtual device which **will replicate your data to AWS**.

* There are 3 flavours of Storage Gateway
  - **File Gateway** : used for storing files as object in S3.
  - **Volume Gateway**: used for storing copies of hard-disk drives in S3.
  - **Tape Gateway**: used to get rid of tapes.

##### Security

- **by default only the account that created the bucket can do stuff with it**

* when you want to assign policies to the resources you do not control, you should be using **resource policies**, in this case know as **bucket policies**. This policies **apply to any identities accessing this bucket**.

- **ACLs are legacy!**. They are attached to bucket or an object.

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

### Load Balancers

- different types:
  - **classic load balancer: LEGACY!**
  - **application load balancer** for HTTP/HTTPS stuff
  - **network load balancer** for **connections that are NOT HTTP/HTTPS**

* allocated to **specific VPC** and **AZ inside that VPC**

- **can operate inside multiple AZs**

* can have security groups attached

- **can even invoke lambda functions ;o**

* **can perform health checks**

- there is a notion of **Target Groups**. This basically allows you to specify multiple EC2 instances without having to specify them explicitly.
  **In the context of EC2 Target Group usually points to a Auto Scaling Group**.

### EC2 (Elastic Compute Cloud)

- resizable compute capacity in the cloud, **virtual machines in the cloud**

* different pricing models:
  - **on demand**: pay per time use (per-hour, per-second)
  - **reserved**: capacity reservation, contract with AWS
  - **spot**: like a market, but for instances, when AWS has free capacity you can bid and buy, **when capacity is needed they will be taken away from you. There are mechanisms which will alert you if that happens!**
  - **dedicated**: psychical machines **only for you**. Mainly used when you have strict licensing on software you are using

- There are different _health checks_:

  - **System Status Check**: this checks the underlying hyperviser (virtualization tool)

  - **Instance Status Check**: checking the EC2 instance itself

* **I**nput **O**put **P**er **S**econd (IOPS): basically how fast the drive is

- Termination protection is turned off by default

* you can create a **bootstrap script**

- EC2 instance has **metadata**. There are a lot of useful information there.

* **To get the metadata info CURL 169.254.169.254/latest/...**
  - `/userdata`: your bootstrap script etc
  - `/metadata/`: has **many options**, IP etc..

#### Auto Scaling Groups

- launching EC2 based on criteria as a service

* **uses launch templates** to **configure how EC2 is launched**. AMI, Instance Type, KeyPairs, Network stuff, Security Groups etc..

- controls scaling, where instances are launched, etc

* **can work with instances in multiple subnets**

- instances inside auto scaling group can be monitored as an one entity. By default _Cloud Watch_ is monitoring every instance individually

* you can control the number of instances by manipulating three metrics:
  - **Desired Capacity**: this is the number **ASG will try to maintain**
  - **Min**
  - **Max**

- there is a notion of **scaling policies**. These are **rules, things you want to happen when something regarding EC2 instances happen**, eg.

#### Security Groups

- **changes** to security group **are instant**

* security groups **are stateful**. That means when you **create an inbound rule, outbound rule is created** automatically.

- there are **no deny rules**. Security groups can only **specify allow rules**.

* EC2 can have more than 1 security group attached to it, also you can have multiple EC2s assigned to 1 security group.

- **All inbound traffic is blocked by default**

* **security group can have other security groups as sources!**

#### EBS (Elastic Block Store)

- basically **virtual harddisk in the cloud**

* persistent storage

- **EBS root volume CAN be encrypted**

* **automatically replicated** within it's own AZ

- Different versions:

  - **Provisioned IOPS** - the most io operations you can get (databases), most expensive
  - **Cold HDD** - lowest cost, less frequently accessed workloads (file servers)
  - **EBS Magnetic**- previous generation HDD, infrequent access
  - **General Purpose**

#### EBS vs Instance Store

- Instance Store a.k.a **Ephemeral storage**. The data lives on a rack where your virtual machine is booted. **If you reboot it, there is HIGH CHANCE that you will loose that data**. It's not 100% guaranteed though. Your EC2 can always be assigned to the same rack, but that is unlikely.

- Instance Store is not really persistent, whereas EBS is a persistent, multi AZ storage option.

#### EFS

- **E**lastic **F**ile **S**ystem (EFS)

* **Similar to EBS**, but there is one **BIG DIFFERENCE**. EFS instance can be used by multiple EC2 instances, EBS volume can only be used by one EC2 instance.

- think of it as multiple EC2 instances having the same disk

* **automatically scales storage capacity**, when deleting shrinks, when adding resizes

- **CAN ONLY BE USED BY EC2** instances

* just like s3 there are different tiers:

  - infrequent access

- data **can be encrypted** at rest.

#### Placement Groups

- ways of _placing_ your EC2 instances

* **clustered**, **spread**, **partitioned**

  - **clustered**: placing EC2 very close with each other, in a single AZ.
  - **spread**: opposite idea, each instance is placed in **separate racks**. Can be in different AZ but the **same region**. **Limitation of 7 instances per AZ**
  - **partitioned**: similar to spread and clustered, basically you have a **groups of clustered EC2 on separate racks**

#### Enhanced Networking

- EC2 **must** be launched from **HVM AMI**

* EC2 **must** be launched **inside VPC**

#### ENI (Elastic Network Interface)

- **by default, eth0 is created**

* has allocated IP address from the range of subnet

- has **interface ID**

* can have ip addresses changed (multiple private addresses), **the number of ip pools are dependant on EC2 instance size / type**

- EC2 **can have multiple ENI's**. When instance is terminated **ONLY default eth0 is deleted BY DEFAULT**.

### CloudWatch

- **monitoring service** for applications, services, **monitors performance**

- can monitor:
  - **CPU**
  - **Network**
  - **Disk**
  - **Status check**

* **Cloud Trail IS NOT THE SAME AS CloudWatch**
  - **CloudWatch** - performance
  - **Cloud Trail** - CCTV camera, **monitors AWS API calls**

- you can create dashboards from metrics

* there is notion of **events**, which basically provides **near instant stream of system events**

### Analytics

#### Athena

- **completely serverless product**

* interactive **query service**

- data lives on s3, it **never changes**

* **schema on read** a.k.a you define what the data YOU would like to look like.

- **schema is not persistent**, schema is only used when you read data (perform queries)

* reduces admin overhead, **no INITIAL data manipulation**

- you only pay for the data you query and the storage on s3 (your source data probably already exist on s3)

* think of **schema as a lens to look through at data**

- mostly used for **ad-hoc queries**, because you only pay for what you use.

#### EMR (Elastic Map Reduce)

- allows you to perform **analysis on large-scaled, semi-structured or unstructured data**

* think **big data** data sets

- there are **nodes**, the **master node splits the work between nodes**

* uses **shared file system through nodes**

- It can use 2 types of storage to perform operations:
  - HDFS: s3 is used to read and write the final data to
  - EMR FS: s3 is used as primary data store to carry out all operations

* **use for on-demand, ad-hoc, short-term tasks**

- **Athena does not manipulate the data, EMR CAN MANIPULATE THE DATA**

#### Kinesis

- **fully managed**

* ingest big amounts of data in real-time

- you put data into a stream, **that stream contains storage with 24h expiry window, WHICH CAN BE EXTENDED TO 7 DAYS for \\\$\$**. That means when the data record reaches that 24h window it gets removed. Before that window you can read it, it will not get removed.

* stream can scale almost indefinitely, using **kinesis shards**

- **NOT A QUEUING SYSTEM**

* **consumers can work on the data independently**

#### Kinesis Data Firehose

- allows you to **store data from kinesis stream on persistent storage**

* you can make queries and such with it on the data that is accepted by firehose (a.k.a read by it)

- **serverless product**

#### Redshift

- **column-based database**

* **you would store data here after the transactions (changes) has been made**, a good example is kinesis => firehose => redshift

- data should not change, column-based dbs are quite bad at handling changes

* you would use it for lets say find pattern in ages querying on HUGE amount of people data.

- **scales automatically**

* think of Redshift as **end-state repository**

### Networking (VPC)

- **CAN SPAN MULTIPLE AZs**

* **SUBNET CANNOT SPAN MULTIPLE AZs**

#### Elastic IP

- **publicly accessible**, static IP address. Usually used with _NAT-Gateways_

#### Route tables

- they allow routing to happen within your VPC or make requests to outside your VPC possible

* **usually associated with a given subnets**

- often used to redirect requests to _Internet-Gateway_

#### NAT-Gateway

- **lives in a public subnet**

* has **elastic static IP address**

- **converts SOURCES` private ip address** to its ip address

* sends its traffic to **internet gateway**. Internet gateway will convert the **private static ip to public ip** and send it to the internet

- basically it **allows** resources in your VPC which **do not have public IP** to **communicate with the internet, ONLY ONE WAY**

* multiple **resources** inside VPC **share the same IP assigned to NAT-Gateway**

- scale automatically

* **IS NOT HIGHLY AVAILABLE BY DEFAULT**. It is placed in a single subnet in a signel AZ. **For true high availability create multiple NAT-Gateways within multiple subnets**

- session aware, that means that responses to the request initialized by your resources inside VPC are allowed. What is disallowed are the requests initialized by outside sources.

### Caching

#### DAX (in-memory cache for DynamoDB)

- runs inside VPC

* uses **cluster architecture**

- works as any cache you expect to work. You get stuff, it's placed in cache, when you get it again you get it from the cache.

* **read-heavy workloads or very low latency**

- two caches, **query cache** and **getItem/batchGetItem cache**

* any reads through DAX are **eventually consistent**

#### ElastiCache

- operates with **products that are NOT dynamoDB**

* supports **Redis** or **Memcached**

- offloading databases: **caches reads just like DAX**

* storing session data of users

- there is an **AUTH for Redis** thingy that can require user to give a token (password) before allowing him to execute any commands

### Random

#### AWS Workspaces

- desktop as a service

* basically you can provision windows or linux desktop to your employers

- you do not have to manage hardware

#### AWS Ops Work
