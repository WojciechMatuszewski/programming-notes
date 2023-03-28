# Full Stack for Front-End Engineers v3

## Modern computing

- The UI, the Server and the Database

- Often people call themselves "full stack engineers", but they do not really understand what is going on under the hood.

  - There is nothing bad with that. But there is something to be said about understanding underlying components – it can help you!

- It is **almost impossible to be good at all the layers of the stack**.

  - There is too much too learn. You **have to pick your battles!**

## Diving in terminal and command line

- The fastest way of communicating with the operating system.

- If you are proficient with commands, you can make a lot of things fast.

### Shells

- Shell is the application that communicates with the operating system.

  - Terminal runs the shell.

## Servers

- Servers usually run on a dedicates hardware to optimize for performance.

- In almost all cases, the servers are redundant (there are a couple of them).

- Nowadays, most of the people run applications on servers in the cloud.

  - The secret of all of this working is: **virtualization**.

    - **Virtualization is an act of dividing a server resources into virtual computers**.

    - You can buy **virtual servers called VPS (_virtual private server_)**.

### Operating systems

- There is the **Kernel (this one talks to the hardware)**, **Shell (this one talks to the Kernel)** and the **User** layer

### Security

- You can use passwords, biometric authorization or ssh keys.

- Passwords can be easily broken if they are bad.

> Here Jem talks about hash functions and the usage of salt to bolster the randomness.

- Interestingly, the Digital Ocean VPS allows `root` access while on EC2 one has to log-in as a _ec2-user_.

  - It's been a long time since I last did ssh onto EC2 machine.

## Understanding the internet

### How does the internet work?

- Mainly through collaboration and rules.

- The request travels through different layers.

  - computer -> network card -> router -> ISP -> Tier 1 ISP -> Data center -> server cluster -> load balancer -> server

> Here Jem talks about TCP, UDP, ICMP and network packets

### How does the domains work

- The _DNS_ and the _Nameservers_.

- The server talks to the browser via the IP.

- There are different type of _DNS records_

  - The **A record maps the IP to the name**.

  - The **CNAME maps the name to another name**.

## Web Servers

- The main players are _Apache_ and _NGINX_.

  - The _NGINX_ can do a lot of things. It can be a web server, _reverse proxy_ or a _forward proxy_.

  - The _NGINX_ will route the request to a given "piece of the stack" on your server.

## Version control

- You should be using it.

- Git uses hash functions to encode content.

## Security

- SSH keys, Firewalls, keeping the software up-to-date, use two-factor-authentication and VPNs.

- **Ports allows us to map specific endpoint to a process or a network service**.

  - Without ports, we would have run out of network addresses much sooner than we did.

### Firewall

- Firewalls allows you to close ports.

  - This is super helpful since **there are tools to scan your machine for open ports, like _nmap_**.

  - An open port, that should not be open, is an attack vector.

  - There are **different tools for configuring Firewall rules**. One of them is the `ufw`.

### Permissions

- Do your best to scope the permissions down.

  - I think everyone is guilty of doing `chmod 777 FILE`. While it might work, it is a bit dangerous.

  - Do understand what the numbers means. Behind every number there is the `rwx` permission set.

## CI/CD

- It is about _validating your code_.

  - **Testing is at core of CI/CD**.

> Here Jem creates a CRON job to invoke a given script

## Diving into the terminal

### Streams (stdin, stdout, stderr)

- Since all of the unix commands implements those interfaces, we can chain them together.

- The `2>&1` means: **redirect both stderr and stdout**.

  - The only time I used this is to get an output from jest: which tests failed without having to scroll endlessly through logs.

### Finding things

- There are many commands to help you find things. The most common among them are `find` and `grep`.

  - `find` is for searching through **file names**, the `grep` is for searching through **file contents**.

    - There is also `zgrep` which allows you to find inside zipped files. Pretty amazing.

## Files / Databases

- Saving everything to a file is inefficient and you cannot shard data across multiple files.

- There are relational and non-relational databases.

> Here Jem demonstrates the basics of SQL language.

## HTTP

- Has two distinct parts, the request and the response.

- **Cookies are just another header in the request part of HTTP**.

### HTTPS

- Secure HTTPS.

  - Prevents the man-in-the-middle attack.

- In most cases, you want to either disable HTTP or redirect the HTTP traffic to HTTPS.

### HTTP/2

- The **HTTP made a request for every file we need**.

  - This is a bit inefficient as the webpages usually require multiple files. Initializing multiple network connections for each file creates big overhead.

- **HTTP/2 allows for _multiplexing_**. This means that multiple files can be fetched on a single connection.

  - Keep in mind that HTTP/2 **is not free!**. It takes a bit more CPU to handle this protocol.

## Containers

- Packaging the applications into small chunks.

  - The chunk contains everything that the application needs to start running. All the dependencies and such.

  - This makes it easy to share / duplicate services. There is no problem of "it runs fine on my machine".

### Orchestration

> Here Jem talks about k8s and load balancers.
