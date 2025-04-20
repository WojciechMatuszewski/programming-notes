# Random Things I've picked up while studying

## Rolling back vs. rolling forward

> Taken from [this AWS Builder's Library article](https://aws.amazon.com/builders-library/cicd-pipeline/).

You might have heard the term _roll back_ before in the context of CI/CD pipelines. It is when you want to bring the system to the old state before a change. **This is the most common way of restoring the system in the case of an emergency**, where a change introduces a bug.

But what happens if the **change latent change – a change introduced a while back that now sits sandwiched between other changes?**
In such cases, the engineer must **pragmatically decide whether it is worth rolling back all the changes, including the latent change, OR is it worth rolling forward**. _Rolling forward_ means creating a change to fix the underlying issue and letting it advance through the CI/CD pipeline. This ensures that we do not roll back valuable releases, like bug fixes.

## Different types of Deployments

### Canary deployment

The idea behind the _canary deployment_ is to **route some percentage of traffic to the new deployment and monitor how the system behaves**.
If everything is ok, you can route more and more traffic to that deployment, until all of the traffic is switched to it.

**The main drawback is the user experience IF things go wrong on the canary**. The users will feel it. If the system suddenly breaks, those users will not have a good time using your app until you route them back to the "old" deployment.

This type of deployment **is great at getting some "real-life" testing done to the new version of the app**.

## Dark Read deployment

Instead of routing the users to new deployment, **you "clone" the request to the new version of the system, but serve the old system to all users**.
This allows you to **monitor how the new deployment behaves without impacting users if things go wrong**.

This deployment type **quite complex and could be costly, but is definitely worth it for critical systems**. If your new deployment explodes, there is no need to rollback – no users are using it!

**Having said that, it is a great way to load test your new deployment**. You can't do that with canary deployments as it would require you to shift _all_ traffic to the new version, rendering the canary deployment an "all at once" deployment.

## TLS Termination

When making a request to a server, you most likely use HTTPS protocol. This ensures the data you sent is encrypted and secure.

To handle the HTTPS request, the server would have to decrypt the data, handle it, encrypt it back and send it to you – **this is a lot of work for a single server to do**.

The **term _TLS Termination_ refers to a point where the data is decrypted and forwarded to the destination**. **This is mostly implemented via _load balancers_, _api gateways_ or _CDNs_**.

When the destination responds, the "TLS Terminator" will re-encrypt the data.

## Different types of Proxies

Keep in mind that the definition change depending on from which _side_ of the request you are looking at.

From the origins perspective, the proxy would be a reverse proxy and so on.

### The _forward proxy_ (or _proxy_)

The **_forward proxy_ sits between you AND the destination**.

A good **metaphor is having an assistant call a restaurant and arrange diner for you**. You have "the assistant in front of you" and the restaurant is the "destination".

From technical perspective, **a good example of a proxy would be any CDN or some tool that filters network traffic**.

### The _reverse proxy_

The **_reverse proxy_ sits between the incoming traffic and the server**.

A good **metaphor is visiting a restaurant and asking for table. The waiter will assign you a seat. From the restaurants perspective, the waiter is a _reverse proxy_**.

From technical perspective, **a good example of a _reverse proxy_ would be a load balancer**.

## The `.env` files

I was watching [this video](https://www.youtube.com/watch?v=j2JRBZaMDSg) where the author makes a compelling case _against_ `.env` files.

- If your `.env` files are not checked into the repository, you might have lots of trouble running the app as it existed at some point in the past.

  - Why would that be useful? Well, if you want to fix an issue, you most likely want to go back to how the app was functioning previously and see what the difference is between the working version and the current one.

- [There is also this follow-up video](https://www.youtube.com/watch?v=5lb3T3R_z2k) where the author talks about the security implications of `.env` files on a server.

  - There have been numerous cases where companies _hosted_ their `.env` files as static files accessible by anyone. In fact, in my work, I see requests to various `.env` files happen all the time! People are actively trying to "probe" for those files in hopes of stealing application secrets.

The author makes a case for using vendor-specific services to host secrets, like _AWS Secrets Manager_. I agree with this approach, but I also think that this solution is not for everyone.

Keep in mind that, depending on the infrastructure, reading secrets might add latency to your application, especially in serverless environments, even when caching those secrets in-memory. In addition, one has to be mindful of the complexity such a solution brings. You usually will need to add some kind of library, or write custom code, to fetch those secrets.
