# About HTTP protocol & networking

## CORS

CORS is the mechanism which enables one webapp (like your localhost) to share some resource with another webapp (like your endpoint API).

If those 2 apps have **the same origin** they can easily share those resources with no hassle at all.

Problems begin when they are on different origins.

So what does _different origin_ mean?

- different domain like `google.com` and `twitter.com`

- different subdomains like `localhost:3000` and `localhost:3000.api/v1`

- different ports like `:3000` and `:4000`

- different protocols like `http` and `https`

To make it work you have to follow the CORS standard.

### Step by Step

Suppose we have 2 apps: A and B. They want to share resources. App A makes a POST request to app B:

- `preflight` request is made (before the actual request) **also known as OPTIONS call**.

- app B now have the responsibility of verifying either this request is valid or not.

- app B sets some additional headers to that request and sends it back.

- now browser knows if the request is valid or not. The actual `POST` request is made.

### Simple Request vs Preflight Request

> You can [read more about _simple requests_ here](https://developer.mozilla.org/en-US/docs/Web/HTTP/CORS?utm_source=stefanjudis&utm_medium=email&utm_campaign=web-weekly-142-so-many-new-browser-features#simple_requests).

So we've seen how the preflight mechanism works. But the next question on your mind probably is: is this happening every time I send a request?

That is not the case. When request meets certain criteria, browser will skip the `preflight` request.

### Caching

In the headers of the "main" request, you can specify for how long the `preflight` request should be cached for.

```
Access-Control-Max-Age: 86400
```

**According to the MDN, the default is 5 seconds**, and there is an upper limit of time for each browser.

Caching those requests is a good idea, as they add to the network overhead. Even better if your API can be on the same origin as your website. This way, you do not have to deal with CORS at all!

## The `Query` method

> More info [here](https://www.ietf.org/archive/id/draft-ietf-httpbis-safe-method-w-body-05.html)

From what I understand, the **purpose of the `QUERY` is to make a GET-like request with body**.

This would be huge for GraphQL, or in cases where you have a ton of query-string parameters.

In the past, for cases like those, you would use `POST` or encode everything in the URL. This has a couple of drawbacks.

- If you encode everything in the URL, the URL becomes unreadable.

- The `POST` request is for changing the resource, not getting it. It does not have the same semantic meaning as `GET`.
