# Aws Solutions Architect Stuff

Just me trying to learn for an exam 🤷‍♀

## Acloudguru & Linux Academy

### Basics

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

### Lambda

- **LAMBDA IS HA by default! MULTI-AZ!**

* **scales automatically** (can run functions concurrently)

### IAM

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

### S3

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

- s3 enables you to turn on **cross region replication of given bucket**, but **you have to have versioning enabled to be able to enable CRR**

* S3 can be `accelerated`. There is something called **S3 Transfer Acceleration** where users upload to **edge locations** instead directly to
  the Bucket. Then that uploaded object is **transferred automatically to your
  bucket**.

- S3 is a **tiered** storage

  - S3 Standard, stored across multiple devices and multiple facilities. **Standard can tolerate AZ failure**

  - S3-IA/S3 One Zone-IA (**Infrequent Access**): for data that is accessed less
    frequently but requires rapid access when needed

  - S3 Glacier / Glacier Deep Archive: used for data archiving, where you would
    keep files for a loooong time. Retrieval time is configurable (**Deep
    Archive is locked on 12hr retrieval time though**)

* **S3 IS NOT A GLOBAL SERVICE**. It has a universal namespace but the **data stays in the region you created the bucket in (unless you specified CRR**.

#### Versioning

- S3 have the notion of versioning: **stores all versions of an object including all writes even if you delete it!**

* After **enabling** versioning, it **cannot be disabled**, only suspended

- Remember that with versioning your previous versions persists on the bucket. This can lead to exponential growth of size when editing big files.

* When an object is deleted, **bucket may seem empty** but that's not the case. You just placed a _delete marker_ on that object (thus creating a new version). Your **previous versions are still there!**.

- You can restore your deleted objects by **deleting a delete marker**.

#### Life-cycle rules

- can be used to **transition objects to different TIER of storage after X amount of time**, this can be placed on current and previous versions.

* With life-cycle there is notion of expiration. Basically your objects can be **expired after X amount of time**. Object can be in a _expired state_ where it will wait to be permanently deleted.

- Can be used with **conjunction with versioning**.

#### Storage Gateway

- Physical/virtual device which **will replicate your data to AWS**.

* There are 3 flavours of Storage Gateway
  - **File Gateway** : used for storing files as object in S3.
  - **Volume Gateway**: used for storing copies of hard-disk drives in S3.
  - **Tape Gateway**: used to get rid of tapes.

#### Security

- **by default only the account that created the bucket can do stuff with it**

* when you want to assign policies to the resources you do not control, you should be using **resource policies**, in this case know as **bucket policies**. This policies **apply to any identities accessing this bucket**.

- **ACLs are legacy!**. They are attached to bucket or an object.

* you would use **identity policies** to **control access** to s3 resources, this however **only works on identities IN YOUR ACCOUNT!, identities you control**.

- using **resource policy** you can **control access** to s3 resources, works on **identities you DO NOT control. THIS ALSO MEANS ANY IDENTITY**. When resource policy is used **specifically with s3**, it is known as **bucket policy**

So when to use what?

- _identity policy_
  - when **you control the identity**

* _bucket policy_ (resource policy)
  - when **you DO NOT control the identity**

### Snowball

- Big briefcase, **up to 100TB** of storage. Used to move data from one point to
  another, fast

### Snowmobile (not joking)

- A truck with a container that carries snowballs.

### RDS (Relational Database Service)

- AWS service for relational databases

* multiple providers such as: `mysql` or `Oracle`

- **RDS AUTO SCALING in terms of compute does not exist**. You have to provision an EC2 instance and **pay regardless of the db usage**. What is possible is **RDS storage auto scaling**

* RDS **can be in multi AZ configuration**

- there are **read replicas available**.

* data is **replicated synchronously across multiple AZ with read replicas**

- when patching os on EC2, **with multi AZ config, patching is done first to standby in different AZ then failed over onto when main db is down due to patching os**

* you **can connect to** the underlying **EC2 instance that hosts the database**. Use SSH or different means.

- you **DO NOT** have **access to the underlying OS, the DB is running on**.

#### Encryption

- you **can only choose to encrypt your DB during creation phase**

* data **in transit** between **source and read replicas** is **encrypted by default**

- you **cannot encrypt existing DB**. You have to **create an snapshot and encrypt it** and build DB from that snapshot.

* **read replicas** have to be encrypted with the **same key as source**

### CloudFormation

- **infra from json**

* is free, **the service itself does not cost anything**. Only the stuff you deploy with it will cost you money (if you provision stuff which is not free)

### DynamoDB

- **more effective for read heavy workloads**

* **data automatically replicated `synchronously` between 3 AZs!**

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

- **ALB can balance** between **different ports**. This is done by **specifying listeners rules**

* ALB/NLB enable you to create **self-healing architecture**. If you are using **ALB/NLB + ASG combo** your application becomes `self-healing` meaning that if one instance fails it gets replaced and such.

#### Monitoring

- **logs send to CloudWatch in 60sec interval IF there are requests flowing through the load balancer**

* if you **need more information** about the flow that goes through your load balancer you can use **access logs, DISABLED BY DEFAULT!**. Load balancer will **store those logs in s3**. These allow you to get information about **individual requests** like IP address of the client, latencies etc..

### ECS (Elastic Container Service)

- allows you to run **docker containers on EC2 instances** in a **managed way**.

* those **EC2 container instances are launched inside a VPC**

- you can put **docker images** inside **elastic container registry (ECR)**

* ECR is integrated with IAM

- **ALB can balance container instances**

* **cluster** part comes from the fact that you will usually have **multiple ec2 instances running your containers**

- to run stuff you have to create **task definitions**. **Describes the task a.k.a service** you want to run. They are **logical group of containers running on your instance**

* with **tasks definitions** you can specify the **image, way of logging and resource caps**

- **each container instance belongs to only one cluster at given time**

#### Fargate

- **a bit more expensive than ECS itself since AWS is literally managing everything expect tasks for you**

So with **ECS you have to have EC2 instances running**. But with **Fargate you really only care about the containers themselves**. You can wave deploying ASG goodbye. **Fargate is basically container as a service, you only define tasks and that is it**.

### EC2 (Elastic Compute Cloud)

- resizable compute capacity in the cloud, **virtual machines in the cloud**

* different pricing models:

  - **on demand**: pay per time use (per-hour, per-second)
  - **reserved**: capacity reservation, contract with AWS
  - **spot**: like a market, but for instances, when AWS has free
  - **scheduled reservations**: this is for tasks that are well, scheduled, daily monthly, whatever. You sign a contract for 1 year.
    capacity you can bid and buy, **when capacity is needed they will be taken away from you. There are mechanisms which will alert you if that happens!**
  - **dedicated**: psychical machines **only for you**. Mainly used when you have strict licensing on software you are using

* there are also **Spot blocks**. This allows you to have an instance for a **specific X amount of time** using the spot pricing model.

- There are different _health checks_:

  - **System Status Check**: this checks the underlying hyperviser (virtualization tool)

  - **Instance Status Check**: checking the EC2 instance itself

* **I**nput **O**put **P**er **S**econd (IOPS): basically how fast the drive is

- Termination protection is turned off by default

* you can create a **bootstrap script**

- EC2 instance has **metadata**. There are a lot of useful information there.

* using **SPOT** instances there are different **behaviors you can configure** when your instance is about to get interrupted:

  - **stop**
  - **hibernate**

* **To get the metadata info CURL 169.254.169.254/latest/...**
  - `/userdata`: your bootstrap script etc
  - `/dynamic/instance-identity`: stuff about the instance -> IP, instance size, type all that stuff
  - `/metadata/`: has **many options**, IP etc..

- **stopping and starting an instance will MOST LIKELY result in data loss on instance store**. Unless you have dedicated tenancy model on that instance.

* **Tenancy model**. This is something **somewhat different than ec2 pricing models**.
  - **shared**: multiple costumers share the same piece of hardware (same rack, etc)
  - **dedicated**: hardware your EC2 runs on is only yours, but you have to pay more
  - **dedicated host**: you can actually pick the server your EC2 will be deployed into

- **EC2 instance can only have ONE IAM ROLE**

* to specify **launch parameters of EC2** you can use **launch templates**. This allow you to specify **some configuration** for an EC2 like **security groups, AMI ids and such** .

#### Auto Scaling Groups

- launching EC2 based on criteria as a service

* **uses launch templates** to **configure how EC2 is launched**. AMI, Instance Type, KeyPairs, Network stuff, Security Groups etc..

* **DO NOT mistake launch templates with launch configurations**. **Launch templates are used with ASG, launch configuration is something EC2 specific!**.

- controls scaling, where instances are launched, etc

* **can work with instances in multiple subnets**

- instances inside auto scaling group can be monitored as an one entity. By default _Cloud Watch_ is monitoring every instance individually

* you can control the number of instances by manipulating three metrics:
  - **Desired Capacity**: this is the number **ASG will try to maintain**
  - **Min**
  - **Max**

- there is a notion of **scaling policies**. These are **rules, things you want to happen when something regarding EC2 instances happen**, eg.

* **REMEMBER THAT YOU HAVE TO LIST AZs YOU WANT YOU INSTANCES TO BE DEPLOYED INTO!!**

- there is a notion of **health check grace period**. This is the **time** it takes to **spin up new instance**. This time of course is dependant on **bootstrap script** and **how much application code is in the AMI**.

* **individual instances can be protected from scale events**. This is useful when you have a master node that cannot be terminated.

##### Default termination policy

There are number of factors that are taken into consideration while picking which instance to terminate.

- **look at the type of the instance if the allocation strategy is specified**

* if instance uses **old launch template** terminate that instance. **that launch template has to be the oldest. if there are more instances with the same old launch templates, skip this step**

- terminate based on **how close an instance is to next billing hour**. Again **if there are multiple of such instances, skip this step**.

* pick **random instance**

Regardless of these steps, default termination policy will try to terminate instances **in the AZ that has the most amount of instances**.

#### Security Groups

- **changes** to security group **are instant**

* security groups **are stateful**. That means when you **create an inbound rule, outbound rule is created** automatically.

- there are **no deny rules**. Security groups can only **specify allow rules**.

* EC2 can have more than 1 security group attached to it, also you can have multiple EC2s assigned to 1 security group.

- **Security Groups DO NOT APPLY TO s3 BUCKETS, WE ARE IN EC2 VPC SPACE NOW!**

- **All inbound traffic is blocked by default**

* **security group can have other security groups as sources!**. This does not mean that we are _merging_ the rules. **Having other security group as source means that we are allowing traffic from instances ENI which are associated with that group!!!!!!!**

#### EBS (Elastic Block Store)

- basically **virtual harddisk in the cloud**

* persistent storage

- **EBS root volume CAN be encrypted**

* **encryption** works **at rest** and **in transit between the volume and the instance**

- **volumes from encrypted snapshots are also encrypted**

* **automatically replicated** within it's own AZ

- Different versions:

  - **Provisioned IOPS** - the most io operations you can get (databases), most expensive
  - **Cold HDD** - lowest cost, less frequently accessed workloads (file servers)
  - **EBS Magnetic**- previous generation HDD, infrequent access
  - **General Purpose**

* you can take **EBS snapshots**

- **during** the operation of **creating a snapshot** you can **use** your volume **normally**

* **snapshots are incremental**. AWS only diffs whats changed **you can delete every, BUT NOT THE LAST SNAPSHOT** if you want to make sure your data is secure.

> If you make periodic snapshots of a volume, the snapshots are incremental. This means that only the blocks on the device that have changed after your last snapshot are saved in the new snapshot. Even though snapshots are saved incrementally, the snapshot deletion process is designed so that you need to retain only the most recent snapshot in order to restore the volume.

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

* ENI can be **attached** in **different stage of life-cycle of EC2**.
  - **hot attached**: when instance is running
  - **warm attached**: when instance is stopped
  - **cold attached**: when instance is launched

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

- for **non standard metrics like RAM usage** you can install **CloudWatch agent** to push those to custom metric.

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

* **master node can** be **sshed into**

#### Kinesis

- **fully managed**

* ingest big amounts of data in real-time

- you put data into a stream, **that stream contains storage with 24h expiry window, WHICH CAN BE EXTENDED TO 7 DAYS for \\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\$\$**. That means when the data record reaches that 24h window it gets removed. Before that window you can read it, it will not get removed.

* stream can scale almost indefinitely, using **kinesis shards**

- **NOT A QUEUING SYSTEM**

* **consumers can work on the data independently**

#### Kinesis Data Firehose

- allows you to **store data from kinesis stream on persistent storage**

* it can modify the data before storing it

#### Kinesis Data Analytics

- allows you to make **sql queries against data in the stream**

* can ingest from **kinesis stream and kinesis firehose**

- **serverless product**

#### Redshift

- **column-based database**

* **you would store data here after the transactions (changes) has been made**, a good example is kinesis => firehose => redshift

- data should not change, column-based dbs are quite bad at handling changes

* you would use it for lets say find pattern in ages querying on HUGE amount of people data.

- **scales automatically**

* **USED FOR OLAP style data**

- think of Redshift as **end-state repository**

* **always keeps `three` copies of your data**

- provides **incremental/continuos backups** just like EBS.

* **automatically caches repeated queries**

### Networking (VPC)

- **CAN SPAN MULTIPLE AZs**

* **SUBNET CANNOT SPAN MULTIPLE AZs**

#### Elastic IP

- **publicly accessible**, static IP address. Usually used with _NAT-Gateways_

* **FREE OF CHARGE AS LONG AS YOU ARE USING IT**

#### Internet Gateway

- used for **connecting** your **VPC to the internet**.

#### Route tables

- they allow routing to happen within your VPC or make requests to outside your VPC possible

* **usually associated with a given subnets**

- often used to redirect requests to _Internet-Gateway_

* **new subnets are associated with main route table by default**

#### NAT Gateway

- **CANNOT HAVE SECURITY GROUP ATTACHED TO IT!**

* **lives in a public subnet**

- has **elastic static IP address**

* **converts SOURCES` private ip address** to its ip address

- sends its traffic to **internet gateway**. Internet gateway will convert the **private static ip to public ip** and send it to the internet

* basically it **allows** resources in your VPC which **do not have public IP** to **communicate with the internet, ONLY ONE WAY**

- multiple **resources** inside VPC **share the same IP assigned to NAT-Gateway**

* scale automatically

- **ONLY HIGHLY AVAILABLE WITHIN SINGLE AZ**. It is placed in a single subnet in a single AZ. **For true high availability create multiple NAT-Gateways within multiple subnets**

* session aware, that means that responses to the request initialized by your resources inside VPC are allowed. What is disallowed are the requests initialized by outside sources.

- **CANNOT HAVE Security Groups ATTACHED!**

#### NACL

- **CAN ONLY DENY RULES**

* is stateless, that means **it does not _remember_ the relation between incoming and outgoing traffic**.

- **cannot block traffic to a given hostname**

* **WORKS ON A SUBNET LEVEL**

#### Security Group

- **CAN ONLY ALLOW RULES**

* there is an **default explicit deny on everything**

- **remembers the relation between incoming and outgoing traffic**. If you ping and instance with security group attached it will be able to ping you back without having to specify outgoing allow.

* **WORKS ON INSTANCE LEVEL**

#### Peering

- **linking TWO!! VPCs together** (in a scalable way)

* when VPCs are peered, services inside those VPCS can **communicate** with each other using **private IP addresses**

- can **span accounts, regions**

* VPCs are joined using **Network Gateway**

- **CIDR** blocks **cannot overlap**

* **VPC peer has to be accepted by the other side**

- you will probably have to **check SG, NACL** to make sure it works. **Enabling peering DOES NOT MEAN that the connection is made**.

* VPC peering **does NOT allow for _transitive routing_**. That means that if you want to **connect 3 VPCs** you have to **create peering connection between every VPC**. You **cannot communicate with other VPC through peered VPC!**

#### VPC Endpoints

- there is a notion of **VPC endpoint**. This allows the **service that the endpoint points to** to be **accessed by other AWS services without traversing public network**. **NO NATGW or IGW needed!**

* there are **two different types** of VPC Endpoints
  - **Gateway Endpoints**: for **DynamoDB and S3**
  - **Interface Endpoints**: for **everything else**

- **Gateway Endpoints** are **highly available**

* **Gateway Endpoints** use **routing, dns is NOT INVOLVED**

- With **Interface Endpoints** you have to **manually select AZs** to make it **highly available**

* **Interface Endpoints** use **DNS, routing is NOT INVOLVED**

#### IPV6

- **not enabled by default**

* **have to be enabled for the whole VPC**

- **all** IPV6 addresses are **publicly accessible by default**

#### Egress-only Internet Gateway

- **allow only OUTBOUND traffic from IPV6 associated instance**.

* **engress-only** means that it **only allows outbound IPV6 connections**. It's **stateful!**. Which means that it **allows elements in your VPC** to **receive the response back**.

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

### Communication Between Services, Queues

### Random

#### AWS Workspaces

- desktop as a service

* basically you can provision windows or linux desktop to your employers

- you do not have to manage hardware

#### AWS Ops Work

#### WAF (Web Application Firewall)

- allows for monitoring of requests.

* will either accept or deny given request

- looks for CSS or SQL-injection stuff

### Patterns

#### Restricting Access to s3 resources

Use `pre-signed` `s3` urls.
If you are distributing through `CloudFront` create `IAO` and associate that `IAO` with your distribution. Then modify permissions on `s3` to only allow OAI to access those files.

#### EBS and encrypting a volume on a running instance

So you would like to avoid downtime on your EC2 instance. **There is no direct way to change the encryption state of a volume or a snapshot**. You have to create **a new volume**, enable encryption (if it's not enabled by default) and copy the data.
Then you can either swap the volumes or restore volume from newly created snapshot.

#### Adding Encryption to an RDS db

**You cannot add encryption to an existing RDS db**. What you have to do is to **create db snapshot and encrypt it**. From that snapshot you can create a copy of your DB.

#### Custom metrics on EC2

**By default** CloudWatch monitors **CPU, Disk and Network**. If you need RAM metrics for example you can **install CloudWatch agent on a EC2** which will **push data to custom CloudWatch metrics**.

#### Troubleshooting EC2 instances in ASG

First thing you need to do is to **place the instance in a standby state**. When in **standby state**, the instance will be **detached from ELB and target group, it is still part of ASG though**. If you do not want ASG to continue the scaling processes you can **suspend ASG scaling processes**. Keep in mind that you are **still billed for the ec2 that are in standby state**

#### Changing instance type inside ASG

**YOU CANNOT EDIT EXISTING LAUNCH CONFIGURATION**. You have to create a **new launch configuration with new instance type**. To make sure that all of your instances are using this new launch configuration **terminate old instances**. New one will get added using new launch configuration.

**Instead of creating new launch configuration** you **can** also **suspend the scaling process** and **restart existing instances after specificizing new instance type**. This is also a solution but seems pretty meh tbh.

#### ELB and Route53

#### Different IAM roles per container on ECS or Fargate

Since **task definitions allow you to specify IAM roles** you should edit those and setup the IAM roles correctly. The launch type does not matter in this case as **either EC2 or Fargate launch type** support that option.

TODO:

- aws private link
- beanstalk
- chaning schema on dynamodb is easly because dynamo is nosql
- iam query API for programmatic access
- swf
- you can attach iam policies to iam groups
