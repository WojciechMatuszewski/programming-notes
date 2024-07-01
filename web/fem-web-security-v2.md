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

  - The `httpOnly` disables getting cookies from JavaScript.

    - It **still goes over the write as plan text unless you encrypt it**.

  - The `Secure` ensures the cookies are sent over HTTPs only.

  - The `SameSite` restricts how cookies are sent across cross-domains.

### More on the `SameSite` attribute

- The `SameSite=None` means **always send the cookie**.

- The `SameSite=Lax` means **allow cookies to be send with top-level navigation and if the request method is safe**.

  - Nowadays, this is the default.

- The `SameSite=Strict` means **only include the cookie if the request is sent from the same site that set the cookie**.

  - This means that **if you copy & paste the link to the website you are already authenticated to, the cookie will not be send!**

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

- Cross-Site Request Forgery (CSRF): think executing requests as a given user, perhaps by making the user clicking on the image.

  - **You can greatly reduce the chance of being vulnerable to this attack by sticking to HTTP semantics**.

    - Usually, this attack is executed because the service allows state changes on GET requests.

  - Another layer of protection against these attacks is the **usage of CSRF tokens**.

- **If you have pages that spit out stack traces, you are giving a lot of information to the attackers**.

  - From the stack traces, they, most likely, can infer the technology stack you are using.

    - If they can do that, they can focus on exploiting only a single piece of tech.

  - **For example, the `X-Powered-By: Express` header**.

    - If someone discovers "day 0" express vulnerability, the attacker can do a lot of harm since they _know_ you are using the `express` library.

## Data Validation

- Data validation is very important.

- **No matter what you do, you ought to validate the data before it goes into your database**.

  - **Do NOT rely on frontend validation**.

    - For example, when using forms, someone might edit the input names.

- Consider working on a basis of _allowlist_ rather than _denylist_.

  - The problem with _denylist_ is that you might miss something.

  - With _allowlist_ you have to concisely allow certain operations. It is much less likely that someone will slip through the cracks.

## CSRF tokens

- This is a special token the server will check for alongside the data user is submitting via various forms.

  - If the token is invalid, or absent from the payload, the server will reject the request.

- This protects you from the CSRF attack because **you embed that token only on the _forms_ you control**.

  - The attacker might still be able to craft the malicious request, but, unless they can literally look at the HTML you send back to the user when they are on a "valid" site, they will not be able to get that token.

- **Putting CSRF token in a cookie is fine as long as you look for the token in the request payload, and NOT the cookie itself**.

  - This pattern is called **"double submit cookie"**.

Finished Part 4 -28:11
