# Wisdom told by others

## Make yourself replaceable

Why on earth would you make yourself replaceable? Is not being unique and "the only person fit for the job" good for you?

That is what I was thinking for the longest time. It turns out that this is not exactly the case.

There is a huge benefit in being **replaceable in a given area** – it allows you to move on to bigger and more exciting things.

Imagine being in a team. Only, and only you, know how to operate a given part of the codebase. Feels nice, right? You get that warm and cozy feeling that you will not be fired any time soon because you are needed - WRONG!

Now, instead of pursuing other interesting projects within your company, you are stuck at your team (they might be a great bunch of people, that does not matter). You locked yourself in. There is no going back. You stagnate.

Instead of being a knowledge silo, spread the knowledge to people around you. Make sure they can operate the system just as well, if not better, than you. Would that mean giving them assignments that you deem interesting at first? Yes! The payoff? You are free from the shackles of your domain. You can switch projects – guilt-free!

## Taxonomies fail at the edges

We, engineers, like to keep things organized. We sometimes view the world as if it was black or white – the user is either a blogger or an advertiser, and so on.

In reality, though, most of the things are intertwined. A blog can be written by multiple authors (collaboration). The blogger can also be an advertiser.

It's up to you to find those relations as soon as possible. Otherwise, your system will not be flexible enough to handle them later. The refactoring of taxonomies takes **a lot** of time.

## Let it fail

Have you ever found yourself implementing a fallback strategy on top of a fallback strategy? If so, you most likely introduced latent bugs into your application.

How come?

The biggest problem with fallbacks is that they themselves can fail as well! How do you handle the failure of a fallback? Adding another is going to lead to the same situation.

The answer here is to **let it fail**.

You are much more likely to successfully reboot a server or swap some disks than to implement a robust failure handling mechanism (**retries do not count as a fallback mechanism**).

The more fallbacks you implement, the more complex your code becomes. You **should spend that effort making the thing you are writing the fallback for more resilient instead**.

### Fallback considered harmful

Sometimes the **fallback mechanism can do more harm than good**. Jacob Gabrielson, in his [Amazon Builder's Library post](https://aws.amazon.com/builders-library/avoiding-fallback-in-distributed-systems/), showcases how a fallback implemented for a particular feature caused the whole Amazon website to be down. Yikes!

## Improving anything but the bottleneck

When tasked with "improvement" work, **always address the bottleneck first**. It might be the case that improving the bottleneck is not "fun" or "sexy", and you might procrastinate doing it by improving other parts of the code or the application. This approach is **very wrong**.

In a great book called _Phoenix Project_, we have learned that we should focus all the improvements on the thing that makes our process, application slow (you get the idea). The more I think about it, the more it makes sense. Sadly this was not obvious to me at the start of my career.

### Product work vs. technical minutia

When picking what kind of task you should do next, always ask yourself, how can I increase the number of customers using our product. Why? **Because customers, or the work related to them, is often tied with your bottleneck**. If there were any bottlenecks, your product would be very popular. How do you identify the bottleneck? Sadly I cannot help you here as doing so requires a lot of context.

Keep searching for bottlenecks and improve those, do not waste time on anything else.

## The Product-Minded Software Engineer

You can write the best code in the world, but your product might fail if your code is not helping customers use your product.

This is where the notion of the _product-minded software engineer_ comes into the picture. If you are a product-minded software engineer, the technical decisions do not stem from technical minutia but rather concrete product needs.

You have to **understand the business domain you are in** and have enough courage to ask the _why?_ questions during feature planning. Questions often provoke thinking deeply about a given subject – thinking deeply about a given topic can result in breakthroughs (usually the good ones).

Consider reading [this great article](https://blog.pragmaticengineer.com/the-product-minded-engineer/) on becoming more product-minded in your day-to-day work.

## Big ball of mud

_"Big ball of mud"_ systems are systems that lack architectural cohesiveness and lost the battle to the _code entropy_. Usually, they are not the result of poor developer skills. Rather they become these little "monsters" due to constant churn and time constraints. Unless refactored, they slowly rot to the point where they are unchangeable and unusable.

Thinking and building systems with proper architectural decisions are expensive and time-consuming. Time and money forces are often much more powerful to the decision-makers.

All too often the mantra when creating software goes like this: _"You need to deliver quality software on time and under budget. Therefore, focus first on features and functionality, then on architecture and performance"_. Of course, this approach **leads to the big ball of mud**.

## Premature pessimisation

We often talk about the concept of _premature opitmization_ –there is no point in optimizing the code unless you need to optimize it. This sounds good until your site loads after 30 seconds or a call to a database take a minute to finish.

If that happened to your codebase, you witnessed the [**premature pessimisation**](https://stackoverflow.com/questions/15875252/premature-optimization-and-premature-pessimization-related-to-c-coding-standar) with your own eyes – people quote the idea of _premature optimization_ as means of writing code that is not performing well (it does not have to perform great).

Where you should focus on are the interfaces and the underlying data structures. If you have those right, you will likely never be in a situation where your system is unusable. With the correct and well-defined interfaces and data structures, you can optimize all you want later on.

Of course, that does not mean we should swing the pendulum to the other side and optimize everything we possibly could. If we were to do that, we would not get anything done.

Focus on the interfaces and the data structures. You will not regret it.

## It starts with empathy

> Consider watching [this video](https://www.youtube.com/watch?v=VLfjooAKOqg).

How many times did you judge a person based on yours, maybe inaccurate, judgment? If your teammate is underperforming, you might think he is lazy, stupid, or unmotivated (I'm guilty of this). How often have you said the same to yourself while looking at your performance?

In such situations, consider looking at what is happening through the lens of empathy. **Replace your judgments with curiosity**. If you do that, you will likely arrive at a different explanation of a situation.

So, next time you or someone you work with is underperforming, instead of judging, ask, "Hey, what is going on? How can I help".

## Hyperbolic discounting or how we want the quick rewards

Have you ever worked with a code that does not _feel_ good to work with? You know what I'm talking about – the code is not that great, is complex and there are very few, if any, tests. You are under pressure to deliver, but you are afraid that your changes will introduce issues.

In those cases, you might ask yourself: "who wrote this code?! It is so bad!". And you would be right to say that, the code is bad, but the experienced developer will also ask himself: "what caused this code to end up this way?". And that is a key question you should be asking yourself most of the time.

See, we as humans, want a quick reward. That is why we eat fast food rather than cooking a healthy meal – it is much more convenient to eat something with a lot of fat and sugar from McDonlands rather than to eat clean. This behavior is not limited to mundane every-day things, **it also applies to business and software development**.

In business, unexperienced (to some degree) leaders (or greedy ones) want to get their product out there _FAST_. They will try to pressure you to deliver as fast as possible. Since you will most likely agree (not something I would advocate for), you end up with a situation I described above – code that has low cohesion and is not designed well.

This behavior of _wanting things now_ and disregarding the implications is something called **hyperbolic discounting**. We do not think about the implications of the "reward" in the long term. All we have in front of us are days or weeks after the feature lands. "Oh my business is going to earn some much money" they say, but they fail to realize that they have just set it to fail in the next year.

---

There is much more to be said about the quality of the software, but I find [this talk especially interesting](https://www.youtube.com/watch?v=aRR0EDazxIk). It touches on the subject of hyperbolic discounting and how to fight it with data-driven insights.
