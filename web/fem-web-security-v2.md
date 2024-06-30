# FEM Web Security v2

## Introduction

- The best part of the web is also it's worst part.

  - You can do A LOT with JS and web.

- There is a tension between getting stuff out and caring about security.

  - It is hard to developers to care about the security when they are under pressure.

## Cookies

> Small pieces of data stored on the client-side to track and identify users.

- Keeping track of authenticated user.

- Tracking users for advertisements purposes.

- The original implementation was fast and loose.

- **There was no spec for cookies till 2011**.

- There are various options for setting cookies.

  - The `httpOnly` disables getting a cookie from JavaScript.

    - It **still goes over the write as plan text unless you encrypt it**.

### Signing cookies

- Means "blending" together the cookie value and the _secret value_ only your sever should know.

  - **Signing a cookie means ensuring it was not tempered with**.

- You should not put this _secret value_ in code in plain text.

  - At the very least, consider using environment variables.

### Sessions

- Consider **separating the session from the identity**.

  - Holding sessions in a database might be a good idea. It gives you the ability to expire (or remove) sessions later on.

- Instead of passing potential sensitive data to a cookie, use the `sessionId`.

  - Then, on the backend, get the sessionId and retrieve any data the session holds.

- **You should sanitize the inputs related to any DB query and especially the ones related to authorization and authentication**.

### CORS

- The **browser checks three things**.

  1. The **protocol**, so `https` or `http`.

  2. The **domain**, for example `frontendmasters.com`.

  3. The **port**, for example 443 or 80.

  If those things are the same, resources have the same origin.

### Common vulnerabilities related to cookies

- Session hijacking: think impersonating a user: think man-in-the-middle attack.

- Cross-Site Scripting (XSS): think malicious script injected via input field or an URL.

- Cross-Site Request Forgery (CSRF): think executing requests as a given user. That user is unaware that this is happening.

- **If you have pages that spit out stack traces, you are giving a lot of information to the attackers**.

  - From the stack traces, they, most likely, can infer the technology stack you are using.

    - If they can do that, they can focus on exploiting only a single piece of tech.

  - **For example, the `X-Powered-By: Express` header**.

    - If someone discovers "day 0" express vulnerability, the attacker can do a lot of harm since they _know_ you are using the `express` library.

Finished part 2 21:13

https://frontendmasters.com/workshops/web-security-v2/
