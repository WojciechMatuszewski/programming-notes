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

Finished Day 1 part 6 53:17
