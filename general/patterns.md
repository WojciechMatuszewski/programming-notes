# Programming patterns

## PubSub

This pattern usually utilizes a queue that sits between publishers and subscribers. The **publishers are not aware of subscribers and vice-versa** which implies **asynchronous communication**. In AWS world, one could implement the _PubSub pattern_ via either SNS or SQS. Here is a good [article on implementing this pattern using SNS or EventBridge](https://docs.aws.amazon.com/prescriptive-guidance/latest/modernization-integrating-microservices/pub-sub.html). Please note that this article blurs the line between the _PubSub_ and _Observer pattern_, but I think it is still a good introduction to the topic.

In JS, consider extending the the native `EventTarget` class. You could also dispatch the events on the global `window`, but that might not be good for the encapsulation and could result in leaky abstractions. [Check out this article for the implementation](https://frontendmasters.com/blog/vanilla-javascript-reactivity/#custom-events-native-browser-api-for-pubsub).

## Observer Pattern

This pattern is somewhat similar to _PubSub pattern_, but **there is no middle man involved**. The **_observer_ and the _subject_ (the one emitting the event) know about each other**, which implies **synchronous communication**. In AWS world, one could implement the Observer pattern via SNS. SNS synchronously invokes the subscribers (but the subscribers are still not aware of publishers).

In addition, there is whole field of _reactive programming_. Libraries like RxJS utilize the _observer pattern_ to handle streams of data and handle _back pressure_.

## Sidecar

At it's core, the _sidecar_ pattern is there to ensure that one unit does one thing, and only that one thing.

Think sending logs to some kind of observability tool. Should your "core" application do it? Perhaps we can offload that concern to a _sidecar_ that does _only_ that?

Your can run the _sidecar_ as a completely isolated service, OR a program on the same host as your main application. There is no silver bullet, and one solution is not necessarily better than the other.

Some developers might refer to this pattern as "side kick".
