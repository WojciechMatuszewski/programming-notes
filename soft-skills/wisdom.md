# Wisdom told by others

## Make yourself replaceable

Why on earth would you make yourself replaceable? Is not being unique and "the only person fit for the job" good for you?

That is what I was thinking for the longest time. It turns out that this is not exactly the case.

There is a huge benefit in being **replaceable in a given area** – it allows you to move on to bigger and more exciting things.

Imagine being in a team. Only, and only you, know how to operate a given part of the codebase. Feels nice, right? You get that warm and cozy feeling that you will not be fired any time soon because you are needed - WRONG!

Now, instead of pursuing other interesting projects within your company, you are stuck at your team (they might be a great bunch of people, that does not matter). You locked yourself in. There is no going back. You stagnate.

Instead of being a knowledge silo, spread the knowledge to people around you. Make sure they can operate the system just as well, if not better, than you.
Would that mean giving them assignments that you deem interesting at first? Yes! The payoff? You are free from the shackles of your domain. You can switch projects – guilt-free!

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

- TODO: Big ball of mud
