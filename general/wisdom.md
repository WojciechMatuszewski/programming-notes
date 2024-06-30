# Programming wisdom

## Write code that is easy to delete, not easy to extend

> Based on [this article](https://programmingisterrible.com/post/139222674273/write-code-that-is-easy-to-delete-not-easy-to).

This article goes in various directions, but the underlying theme is the following: **It is more than okay to delete
code. Writing code that is easy to delete should be one of your priorities. You will not get the design right when
writing the code for the first time**.

Of course, life is more complicated. There are layers to this premise.

1. **Consider not writing code at all**. Think about writing the code as a last resort. Sometimes, you can avoid writing
   code by ensuring everyone is on the same page. Think of times you wrote some piece of code and then deleted it
   shortly after because things changed.

2. **Copying and pasting code is healthy to a certain degree**. You would not want to copy and paste every time, but
   doing so a couple of times will expose patterns you could generalize. The same applies to writing boilerplate.

3. If you have a lot of boilerplate and are ready to collapse code into a module, do so. Feel free to create
   abstractions over more complicated code. **Thread carefully when creating an abstraction**. A wrong abstraction will
   do much more harm than good.

> It is not so much that we are hiding detail when we wrap one library in another, but we are separating concerns: requests is about popular http adventures, urllib3 is about giving you the tools to choose your own adventure.

4. **There is nothing wrong with writing a bunch of code, learning from it, and deleting it afterward**. It is much
   easier to remove one big mistake than many smaller mistakes scattered throughout the codebase. **Keep the experiments
   within a specific boundary that you can delete without any issues afterward**.

## Aging Code and the constant need to rewrite it

> Based on [this article](https://vadimkravcenko.com/shorts/aging-code/).

When working on any codebase, you will most likely interact with the so-called "legacy code". In most cases, this code
enabled your employment in the first place – this is where the money is!

Since it is much easier to write the code than to read it, you might be inclined to have a need to rewrite parts of
the "legacy". You might justify your decision by the fact that the code is not using the newest technology or it is
not "performant enough" (that is rarely the case).

At this exact moment, you should stop and evaluate your thoughts/decisions carefully. **Would it be of benefit for the
company to rewrite a given piece of code?** Be very specific here. If you do not have a clear answer to this question,
stop now!

See, there is a certain kind of wisdom in the "old code". Since the code age is relatively high, **it is most likely
battle-tested and is edge-case free**. These edge cases are the most problematic – you will most likely miss them! By \*
\*rewriting the code, you will increase the maintainability\*\*. The code will be different, the bugs will be different,
and worst of all, it will not be as battle-tested as the old, well-aged code.

Of course, sometimes the situation is so bad that there is no other way than to rewrite the code. The so-called "big
ball of mud" can suck the life out of developers and make projects grind to a complete stop. In such cases, I would also
tell caution – start very small, piece by piece, and resist the temptation to rewrite "everything"\*\*.

## Software businessman mindset

> Based on [this article](https://vadimkravcenko.com/qa/how-to-stop-thinking-as-an-engineer-and-start-thinking-like-a-business-man/).

We often think about ourselves as "Software Engineers". We like to solve hard technical problems. We like to get into
the "flow" and code away for hours and hours on end.

But the truth is, **in your programming job, you must also be thinking like a businessman**. You will have a unique
ability to understand the technical aspects of your business that others do NOT have.

Always remember that the **code is only a means to the end, not the other way around**. Business is business, and, in
the end, it does not matter what is "under the hood". What matters is how much money it brings – this is the money that
pays your bills!

## Domain knowledge is more important than coding skills

> Based on [this article](https://vadimkravcenko.com/shorts/things-they-didnt-teach-you/#h-domain-knowledge-is-more-important-than-your-coding-skills).

As previously stated in this document, being a good software engineer is also being a good business man. You **cannot be
a good business man without having a deep understating of the domain you are working in**. Knowledge of the domain will
make your life much easier.

Picture two people, they have the same coding skills, but one person is intimately familiar with the business domain,
and the other is not. If you give them the same task, the results will be widely different.

The person familiar with the domain will most likely ask a lot of good, clarifying questions. **They may even conclude
that writing code is not necessary to do the job they were given**. Since the code is only ONE way to get the job done,
and not writing code is always better than writing it, the outcome will be very beneficial to the company as a whole.

Contrast this with the artifact produced by the person who does not understand the business domain. They might default
to writing code, which in itself is not a bad thing. The worst thing is that the code might work, but it does not really
do what it suppose to. This wastes everyone's time and energy.

So, **whatever you are working on, please make sure you have a deep understating of the domain you are working in**.

## Being a "kind engineer"

> Based on [this youtube video](https://www.youtube.com/watch?v=wTezaqqyzlk) and [this blog post](https://kind.engineering/).

- During the code review, **ask about the why. Do not be dogmatic**.

- People do what they think is right given their current understanding of things.

- **Always assume that others meant good**. People do not wake up one day and choose to do nefarious things.

- Always be honest. **Do not create fake personas at work**. While it might seem beneficial at first, in the long run,
  it will wear you down.

- Being your "true self" is liberating. It allows you to focus on what matters the most – the product, your
  colleagues and the shared goal you are marching towards.

- **Care about the people you are working with**. People are not one-dimensional.

- Being kind also means giving direct feedback. Sometimes that feedback can be negative. That is okay.

- People will respect you if you are honest with them. Even when talking about hard things, it is very important to
  be honest.

- **Encourage feedback. Be vulnerable**. If you want others to share what is on their mind, you have to show that you
  are also able to do that.

- I find that talking openly about your problems, your life (of course, be reasonable here) really does wonders when
  it comes to psychological safety.

- **Start by asking for criticism**. Stop giving it. This shows you are open, you are vulnerable. This build trust.

- **People should be accountable for things they do, but you should not blame them**. There is **literally ZERO benefit
  in blaming someone for something**.

- You will not resolve the issue faster if you blame someone for something.

- _We succeed together and we fail together_.

## Do not solve problems you do not have

> Based on [this article](https://renegadeotter.com/2023/09/10/death-by-a-thousand-microservices.html)

Before embarking on solving any problem, be it related to software or otherwise, always ask yourself this one
question: "What problem do I need to solve?"

Having this in mind will guide you towards the actual solution (in some cases, this solution could be different than the
one you had in mind), but most importantly, it will shield you from overthinking and over-engineering.

Nowadays, we have so many things to think about when building services. One has to take scale, resiliency, backups,
overall architecture, frameworks, CI/CD, load testing, and so the list continues. But in most cases,**what you are
building does not require most of these**. You are not Google or some other huge company (unless you are, but then that
is a different discussion).

Focus on the essentials and product value. Be the person who engineers for the customers, not for their own ego. If you
do, I would argue that you stand a higher chance of achieving success than doing it any other way.

## Make yourself replaceable

Why on earth would you make yourself replaceable? Is not being unique and "the only person fit for the job" good for
you?

That is what I was thinking for the longest time. It turns out that this is not exactly the case.

There is a huge benefit in being **replaceable in a given area** – it allows you to move on to bigger and more exciting
things.

Imagine being in a team. Only, and only you, know how to operate a given part of the codebase. Feels nice, right? You
get that warm and cozy feeling that you will not be fired any time soon because you are needed - WRONG!

Now, instead of pursuing other interesting projects within your company, you are stuck at your team (they might be a
great bunch of people, that does not matter). You locked yourself in. There is no going back. You stagnate.

Instead of being a knowledge silo, spread the knowledge to people around you. Make sure they can operate the system just
as well, if not better, than you. Would that mean giving them assignments that you deem interesting at first? Yes! The
payoff? You are free from the shackles of your domain. You can switch projects – guilt-free!

## Taxonomies fail at the edges

We, engineers, like to keep things organized. We sometimes view the world as if it was black or white – the user is
either a blogger or an advertiser, and so on.

In reality, though, most of the things are intertwined. A blog can be written by multiple authors (collaboration). The
blogger can also be an advertiser.

It's up to you to find those relations as soon as possible. Otherwise, your system will not be flexible enough to handle
them later. The refactoring of taxonomies takes **a lot** of time.

## Let it fail

Have you ever found yourself implementing a fallback strategy on top of a fallback strategy? If so, you most likely
introduced latent bugs into your application.

How come?

The biggest problem with fallbacks is that they themselves can fail as well! How do you handle the failure of a
fallback? Adding another is going to lead to the same situation.

The answer here is to **let it fail**.

You are much more likely to successfully reboot a server or swap some disks than to implement a robust failure handling
mechanism (**retries do not count as a fallback mechanism**).

The more fallbacks you implement, the more complex your code becomes. You **should spend that effort making the thing
you are writing the fallback for more resilient instead**.

### Fallback considered harmful

Sometimes the **fallback mechanism can do more harm than good**. Jacob Gabrielson, in
his [Amazon Builder's Library post](https://aws.amazon.com/builders-library/avoiding-fallback-in-distributed-systems/),
showcases how a fallback implemented for a particular feature caused the whole Amazon website to be down. Yikes!

## Improving anything but the bottleneck

When tasked with "improvement" work, **always address the bottleneck first**. It might be the case that improving the
bottleneck is not "fun" or "sexy", and you might procrastinate doing it by improving other parts of the code or the
application. This approach is **very wrong**.

In a great book called _Phoenix Project_, we have learned that we should focus all the improvements on the thing that
makes our process, application slow (you get the idea). The more I think about it, the more it makes sense. Sadly this
was not obvious to me at the start of my career.

### Product work vs. technical minutia

When picking what kind of task you should do next, always ask yourself – how can I increase the number of customers
using our product. Why? **Because customers, or the work related to them, is often tied with your bottleneck**. If there
were any bottlenecks, your product would be very popular. How do you identify the bottleneck? Sadly I cannot help you
here as doing so requires a lot of context.

Keep searching for bottlenecks and improve those, do not waste time on anything else.

## The Product-Minded Software Engineer

You can write the best code in the world, but your product might fail if your code is not helping customers use your
product.

This is where the notion of the _product-minded software engineer_ comes into the picture. If you are a product-minded
software engineer, the technical decisions do not stem from technical minutia but rather concrete product needs.

You have to **understand the business domain you are in** and have enough courage to ask the _why?_ questions during
feature planning. Questions often provoke thinking deeply about a given subject – thinking deeply about a given topic
can result in breakthroughs (usually the good ones).

Consider reading [this great article](https://blog.pragmaticengineer.com/the-product-minded-engineer/) on becoming more
product-minded in your day-to-day work.

## Big ball of mud

_"Big ball of mud"_ systems are systems that lack architectural cohesiveness and lost the battle to the _code entropy_.
Usually, they are not the result of poor developer skills. Rather they become these little "monsters" due to constant
churn and time constraints. Unless refactored, they slowly rot to the point where they are unchangeable and unusable.

Thinking and building systems with proper architectural decisions are expensive and time-consuming. Time and money
forces are often much more powerful to the decision-makers.

All too often the mantra when creating software goes like this: _"You need to deliver quality software on time and under
budget. Therefore, focus first on features and functionality, then on architecture and performance"_. Of course, this
approach **leads to the big ball of mud**.

## Premature pessimisation

We often talk about the concept of _premature opitmization_ –there is no point in optimizing the code unless you need to
optimize it. This sounds good until your site loads after 30 seconds or a call to a database take a minute to finish.

If that happened to your codebase, you witnessed the [**premature pessimisation
**](https://stackoverflow.com/questions/15875252/premature-optimization-and-premature-pessimization-related-to-c-coding-standar)
with your own eyes – people quote the idea of _premature optimization_ as means of writing code that is not performing
well (it does not have to perform great).

Where you should focus on are the interfaces and the underlying data structures. If you have those right, you will
likely never be in a situation where your system is unusable. With the correct and well-defined interfaces and data
structures, you can optimize all you want later on.

Of course, that does not mean we should swing the pendulum to the other side and optimize everything we possibly could.
If we were to do that, we would not get anything done.

Focus on the interfaces and the data structures. You will not regret it.

## It starts with empathy

> Consider watching [this video](https://www.youtube.com/watch?v=VLfjooAKOqg).

How many times did you judge a person based on yours, maybe inaccurate, judgment? If your teammate is underperforming,
you might think he is lazy, stupid, or unmotivated (I'm guilty of this). How often have you said the same to yourself
while looking at your performance?

In such situations, consider looking at what is happening through the lens of empathy. **Replace your judgments with
curiosity**. If you do that, you will likely arrive at a different explanation of a situation.

So, next time you or someone you work with is underperforming, instead of judging, ask, "Hey, what is going on? How can
I help".

## Hyperbolic discounting or how we want the quick rewards

Have you ever worked with a code that does not _feel_ good to work with? You know what I'm talking about – the code is
not that great, is complex and there are very few, if any, tests. You are under pressure to deliver, but you are afraid
that your changes will introduce issues.

In those cases, you might ask yourself: "who wrote this code?! It is so bad!". And you would be right to say that, the
code is bad, but the experienced developer will also ask himself: "what caused this code to end up this way?". And that
is a key question you should be asking yourself most of the time.

See, we as humans, want a quick reward. That is why we eat fast food rather than cooking a healthy meal – it is much
more convenient to eat something with a lot of fat and sugar from McDonald rather than to eat clean. This behavior is
not limited to mundane everyday things, **it also applies to business and software development**.

In business, unexperienced (to some degree) leaders (or greedy ones) want to get their product out there _FAST_. They
will try to pressure you to deliver as fast as possible. Since you will most likely agree (not something I would
advocate for), you end up with a situation I described above – code that has low cohesion and is not designed well.

This behavior of _wanting things now_ and disregarding the implications is something called **hyperbolic discounting**.
We do not think about the implications of the "reward" in the long term. All we have in front of us are days or weeks
after the feature lands. "Oh my business is going to earn some much money" they say, but they fail to realize that they
have just set it to fail in the next year.

---

There is much more to be said about the quality of the software, but I
find [this talk especially interesting](https://www.youtube.com/watch?v=aRR0EDazxIk). It touches on the subject of
hyperbolic discounting and how to fight it with data-driven insights.

## Always assume good intent

Engineering is a team sport. You would not be able to achieve a lot without the help of others. Since you will be
working with other people, it is essential to have a good relationship with your colleagues.

As we know, people are different. Due to those differences, you will most likely disagree with someone about a specific
topic. You also might be dissatisfied how they are performing at work. These things are natural and one has to learn to
deal with them gracefully.

One strategy I use when faced with disagreement or being let down by a quality of someones work is that **I always
assume the person means good. It might be that I do not have a full context, and from my point of view, given what I
know, this is not the case**.

When defaulting to "this person means good", we can employ empathy and compassion which "humanizes" the person we
disagree with. It changes the person from a "someone who I have a problem with" to "a friend who I have differences of
opinion with". **Such change has profound effects**.

You will be talking with that person differently. Not as a frustrated colleague, but rather concerned friend. Concerned
about the future of the team, the code quality and similar. **This also makes the other person feel safe and gives them
the chance to open up**. In most cases, the issue you have with a given person only stems from misunderstanding or
cultural differences.

So, always assume good intent. No matter what other people at work do. It is very unlikely that the person woke up one
day and decided to be a bad person. This almost never happens.

## Sense of Urgency vs Sense of Purpose

> Based on [this wonderful piece](https://medium.com/@kimber_lockhart/don-t-create-a-sense-of-urgency-foster-a-sense-of-purpose-724e309ecdb0)

Have you ever been on a team where the management was pushing for a certain deadline? Most of us have. Now ask yourself
this: were you aware of WHY the deadline is set for a certain day? Or maybe the deadline felt arbitrary to you?

If the deadline felt arbitrary, you experienced management trying to create a _sense of urgency_. Management does this
to supposedly speed up the development process. They think that, by creating a _sense of urgency_ the project will
finish faster. While this might be the case in some scenarios, **creating a _sense of urgency_ will most likely
back-fire**.

Why? The _sense of urgency_ creates detachment from the _project purpose_. **The _sense of purpose_ is the thing that
drives innovation and excellence within the team**. The **_sense of urgency_ proliferates burnout and mistakes due to
rushing**.

Have you ever coded for work during the weekend without anyone asking you to do so? If so, you most likely cared for the
project so much that you decided it would be neat to fix this one little but, or maybe create this little feature in
your free time, because that would make the project better. **This is the _sense of purpose_**. You knew WHY of the
project, you were able to emphasize with users and decided to take the time out of your day and work a bit more. Nobody
forced you to do this.

Contrast this with the dynamics of the _sense of urgency_. Your boss asked you to code during the weekend. Because you
do not know WHY the deadline is so tight weekend coding is required, you are not looking forward to sacrificing your
weekend for job-related stuff (why should I do X when I could spend my time doing Y?). **Now you rush to implement what
you have been asked to do and deliver the feature. Since you rushed, you made some shortcuts**. In addition, your boss
is probably micromanaging you ensuring that you actually work during the weekend. Not a fun experience.

**Always create a _sense of purpose_ (ensure engineers and you know the WHY)**. Skip deadlines that seem arbitrary,
always ask the WHY question until.

## On the separation between the design and web teams

Have you ever worked in environment where the "design team" handed off designs to the "web team"? That how it works in
most organizations. I'm not saying this way of working is good or bad (in fact, I quite enjoy this dynamic), but **what
is important is how the "design team" and "web team" cooperate together**.

Imagine a scenario where the "design team" literally hands-off ready set of mockups to the "web team". There were no
consultations made. The process of creating the designs was purely on the "design team" side. How does that feel to you
as a member of the "web team"? Do you feel like you have a say in the products direction? Not really, right? This
dynamic between two teams is dysfunctional as the **"design team" cannot be solely responsible for the UI mockups. It is
the user requirement that should drive the UI, not the other way around**.

Let us reverse this scenario. Imagine a world where you, as a member of the "web team", implement the 80% of the UI –
the common stuff. You do this without bothering the "design team". You can do that since you understand what you are
working on. You have enough domain knowledge to decide what would "feel right" for the user (of course you might be
wrong, but that is not the point). Of course you do not work within a silo. You "work in public", you notify people who
might be interested in your progress, you listen and gather feedback.

Now, when that tricky 20% of UI design work comes, you lean heavily on the "design team" **while trying to learn their
process as well!**. In the end, you should learn how the "design team" operates, what tools they use and learn to use
those tools as well!

**The UI work is just as collaborative as any. Remember that it is the requirement that should drive the UI look and
feel, not the other way around**.

## On estimates

Why do we estimate? Do we need to estimate how long a given task will take?

Having an estimate gives our bosses a way to calculate risk. Their job is to make a decision given information about
those estimates. But that leaves us in an awkward position.

If you make a bad estimate, your boss will most likely make a bad decision. Some might argue that the fault lies solely
on you. Missing an estimate also has psychological effects – it leads to stress and could cause developer burnout.

So, do we need to estimate? I would argue that **we do not need to estimate everything. We only need to estimate the
most critical tasks, which SHOULD appear very seldomly**.

If you manage to build trust and collaborative spirit in your organization, you will notice that, without estimation,
the tasks are done much quicker. Why? Because nobody is stressed about the deadline and most people want to do their
best work. That also means delivering things as fast as possible while keeping architectural and other aspects in mind.

**Focus on building trust and inclusion instead of having deadlines**. You might consider the deadline a motivational
factor, but it only creates a **fake sense of urgency**.

Now, some deadlines are so-called _"hard deadlines"_. Think of running ads on TV or something similar. In that case,
**the most reasonable way to estimate would be to provide stakeholders with percentages for a given timeframe**.

"The probability we will get this done by Friday is about 30%. The probability we will get this done by next Tuesday is
about 50%", and so on.

This gives your boss enough information to decide, but if something unexpected happens, you tell the truth – that you
cannot predict the future.

## When joining new team

Here are some ideas about what to do when joining a new team.

- **Use the product**, a lot. Take some time and try to understand who the users are. What are the "personas". How do we
  cater to each of them?

- **Understand where the system fails**. Doing so will give you A LOT of ideas on what to improve. Discuss those ideas
  in a larger group and **implement them
  ** — [related blog post](https://blog.staysaasy.com/p/new-hires-learn-how-the-system-breaks?utm_source=profile&utm_medium=reader2).

- **Start adding tests**. Other engineers will love you. If the codebase does not have that many tests (from my
  experience, that is the case for almost all codebases), start writing them.

- **Update onboarding documents**. They are outdated. Unless the project's README is up-to-date (it never is), your
  first contribution could be "Update the README" PR.

- **Do the tedious, so-called "dirty work"**. Yes, it might not be super fun, but guess what? You will learn a tone
  while doing so.

## The spectrum of speed and robustness

Engineering is an art of balancing between tradeoffs. One axis is the tradeoff between the speed and robustness.

Do you want something out _very_ fast? I could skip the testing, the documentation and thinking about the big picture. I
could _just_ make the feature work.

Do you want something out a bit slower, but more robust in implementation? Something that will not collapse under the
weight of extension? Well, then I need to write tests, think about the big picture and spend some time on design.

So, which approach is better? **Of course, there are no silver bullets**. One might argue that the latter approach is
always better. One might argue otherwise.

The thing is, **the best approach is the blend of those two, but you have to keep your engineering discipline in check
**. **Some things are non-negotiable**, like thinking about the big picture or writing tests. Some thing could be
skipped. The **most important thing is to not let yourself become the
so-called [_tactical tornado_](https://medium.com/@parallelit/tactical-tornado-f5e0414087af)** – if you do, your
colleagues might start to resent you.

The best engineers live on the spectrum. They have a set of things that are "non-negotiable" at the core of their
practice, but they are flexible to move from one end to another (and stay in the middle most of the times).

It is up to you, and nobody else to decide what approach is the right one. Do not let product managers, your bosses and
anyone else to dictate where on the spectrum you need to be to deliver the feature. Remember – **you are the one doing
the job, not your boss or product-manager**.

Of course, speaking in "nevers" and "always" is not a good approach either. Your work, like life, if _fluid_. The better
you are at "bending and twisting", the more situations you can handle.

## Conway Law and software

> [Based on this article](https://medium.com/@fwynyk/conways-law-in-team-topolgies-did-you-really-get-it-69c1a4d702af)

The law states

> Organizations which design systems are constrained to produce designs which are
> copies of the communication structures of these organizations.

If you think about it, you have most likely already seen this law in action.

Perhaps, you worked on a team responsible for the _design system_ and the _component library_ used on the website of
your company?
Then, you most likely were part of the _design team_ or similar named team that operated as "SaaS" within the company.
Other teams depended on the
library you maintained.

Or perhaps you worked on a team where you had to talk to the "infrastructure team" for any kind of database-related
inquiries? Then, you most likely
had a single database and tried to do a lot with only that single instance (given that spinning another DB would incur a
high cost of communicating with the "infrastructure team").

The kicker is that **you can leverage the Conway Law to ensure the team builds with the architecture you had in mind**.

— Need a modular, loosely coupled architecture? Organize the teams to work independently, make the culture open and
informal.

— Need a monolithic, centralized architecture? Have a "architecture team" or similar so that all other teams depend on
that team.

**Thinking about the team structure first, then about the architecture is so-called "reverse Conway maneuver"**.

## Divergence & convergence in software

> [Based on this article](https://khalilstemmler.com/letters/divergence-convergence-spaghetti-code/).

Have you ever met the so-called _tactical tornado_ programmers? I bet you did.

Those people focus on code and only on the code. What matters is what works, not how it is built.
Of course, there is a time and place for working in such a manner.
When building an MVP, you most likely lean a bit more towards the "tactical tornado" style of writing code.

**The problem surfaces when ALL you do is to write code in that way**.
The act of programming tactically akin to _divergence_.

You make a mess, you prototype, you probe the system to see what works.
**After _divergence_ there ought to be a period of _convergence_** where you "clean up" the code – create abstractions,
refactor based on what you learned in the previous "phase".

If you stay in the _divergence_ phase for too long, you will make a mess. A mess that everyone will be scared to
touched. A mess that will hinder product development.
**A good engineer is able to cycle through _divergence_ and _convergence_ phases** and ensures that they ALWAYS have a
_convergence_ phase after the _divergence_ phase.

## Abstractions and Illusions

> [Based on this talk](https://www.youtube.com/watch?v=aWZFRk-w3ng)

As software engineers, we are hired to solve problems. In most cases, people who pay for our work do not care _how_ we
solve those problems.
But if we are not careful, our peers might suffer from our recklessness and tactical thinking. **This happens when you
solve the problems without building good abstractions**.

The art of building good abstractions is hard to master. It is because **abstractions can become _illusions_** which are
quite dangerous.

- Those _illusions_ will make you assume incorrect facts about the system.
- Those _illusions_ will make you design for X when, in reality, you should design for Y.

Let us consider RPC. RPC often is though as "local calls" with zero network latency. **That is not the case**!
Now, consider a fellow software engineer who builds a system with this assumption. The RPC gives him the _illusion_ of "
safety", but the reality is different.

So, how do we design good abstractions that are NOT "illusions"?.

- **Do not abstract too much**. Abstract what is necessary, but allow for flexibility.
- This is VERY hard to get right.
- **Do not name things by the pieces it is made out**.
- This makes the abstraction useless. It "leaks implementation" details.
- A good example is the "gas pedal" in cars. For petrol-powered vehicles, it kind of makes sense. But what about
  the electric vehicles?
- **Create a new vocabulary for the abstraction**.
- This shifts the "level" at which you communicate through the interface. The vocabulary should be rather generic.

## Vanity metrics

> [Based on this great article](https://www.ampedcommunity.com/p/vanity-metrics-lol)

Take a look at the metrics your team is tracking (you do have metrics you track, right?)

Are some of them related to commits, or the number of related features?
If so, I'm afraid the metrics you are tracking are so-called _vanity metrics_.

_Vanity metrics_ **will, most likely, not influence your product in a very positive way**.
Your customers, people who pay for your product, DO NOT care how many commits you have pushed.
While they might care about the features, they most likely care about the _quality_ much more than then _quantity_.

You might ask: "Okay, but what is the problem with tracking those things?"
Well, the **biggest problem with _vanity metrics_ is that they create wrong intensives**.

Go ahead and watch the video linked in the article. The ad is hilarious, and it showcases just how far we can go
if some kind of _vanity_ metrics is introduced. Everyone started swearing, but the leadership probably expected the
reverse to happen.

**Make sure the metrics you are tracking are as "close" to your product as possible**.

- Number of paying users.
- Growth of paying users over time.
- Churn.
- Bug count.

A lot of _vanity metrics_ should be treated as something, we software engineers do by default.
Writing tests, code reviews and so on. These things are part of our job – **that
is what you are paid to do**.

## Playing a "dumb" engineer or a manager

> These are my random thoughts. Probably all nonsense.

**First of all, there are NO "dumb" managers or engineers, just like there are no dumb questions**.

So, let us say you want to extract some knowledge from another team or peers. It might be how the project is doing in terms of a deadline,
or it might be something else. Either way, the **objective is to extract information without making the other party feel rushed or under pressure** more than they already are.

To do that, you can leverage the fact that humans **like to feel like experts by explaining something to others**.
Think about the time your colleague asked you to explain something to them. It did feel good right? It felt like you are valuable and needed!

So, even if you know the answer to the question, try playing "dumb". Pretend that you do not have the necessary knowledge required in your work.

_"Hey. I'm unsure if I understand the complexity of project X. Would you mind elaborating on it?"_

Then, when the person answers, you can add: _"oh this seems complex. Given all of that, are you sure you will make it by Y date?"_

You "tricked" the other person giving you an honest answer. Ideally, you would not need to _"trick"_ anyone, but oftentimes, it is necessary.
The answer you get is the honest, and the most accurate one (most likely).

Why? Because you "gained" trust of the other person by asking for help (and giving your peer the chance to show his expertise).

Playing _"dumb"_ will not hurt anyone. **Please do not overdo it**. It is a good tactic, but is might backfire in the long term.

## Definition of fixed and definition of done

Each task should have a _"definition of done"_.

Having such definition prevents scope creep and ensures we de-scoped the feature into the smallest deliverable.
What about the bugs? Have you ever considered having a _"definition of fixed"_ in place? What would that look like?

I've noticed this trend where engineers, upon discovering the issue, change the code and announce that the bug was fixed.
**I deem the bug only partially fixed in this case, because we did not add any guardrails to prevent us from introducing this bug into the code again**.

How certain are we that a new person joining the team will not repeat the mistake that we did? – **without a test, we can't**.
That is why one of the best things you can do is to **start with a failing test, then fix the issue**.

The test will guide you towards the solution. In addition, the **test will act as documentation for other programmers**.
**Unless the production is on fire, always start with a failing test**. This requires discipline. This requires putting a bit of thought
into what you are doing BEFORE jumping into code.

## Cost vs. Price in software

> Based on [this great article](https://www.germanvelasco.com/blog/refactoring-is-a-habit).

When buying an item, you have to account for its _price_, but also for its _cost_ down the line.

> Good habits have a price. Bad habits have a cost. Either way, you pay.

It turns out **this analogy is also applicable in software**.

You can think of the codebase "health" as human health. If you eat healthy and exercise, you incur a price. A price to buy the gym membership. A price to pay for the healthy food.

Living a healthy life, reduces the risk of _paying the cost_ later – when you are older.

**You can think of refactoring or not refactoring in the codebase on the same basis**.

You can refactor the codebase bit by bit, after every feature. You pay the _price_ of time now.

Alternatively, you could skip refactoring now, and _pay the cost_ of having to maintain a _"big ball of mud"_.

**Either way, you are going to pay**. But which "payment" is better for the product you are working on?

Well, in most, if not all cases, you want your product to provide value for a long time. You want to be able to add new features to cater to different users, right?

If that is the case, I would argue **that paying the _price_ now is much cheaper** than paying the _cost_ later on.

## Seek First to Understand

> Based on [this great article](https://dannorth.net/seek-first-to-understand/)

When joining a new project or a company, there are bound to be things you wish were different.

Perhaps the team processes are manual and their build takes what feels like eternity?

Perhaps there are organizational issues within the company structure?

In fact, in most cases, there will be so many things you wish were different, that you will have a hard time deciding where to start.

**If you try to "fix" all the things, you will be spread too thin and achieve mediocre results** – the law of _raspberry jam_: _"The wider you spread it, the thinner it gets"_.

**Instead of trying to fix all the things, focus on understanding SOME issues deeply and zoom in on handful of those**.

The author of the article recommends focusing on three things and sucking up the rest. **Having such discipline is very hard** – engineers are people who like to solve problems!

Do no spread yourself too thin. Focus your energy and effort on a few issues and fix them as best as you can!

## Communication vs. Reporting Structures

> Based on [this great blog post](https://kevinyien.com/blog/communication-structures.html).

Over time, every company develops a hierarchical reporting structure. It's normal.

**While having a hierarchical reporting structure is expected, the _communication structure_ should be as flat as possible**.

If you, an engineer, believe that an input from an executive team is needed, you should be able to ping ANY executive and get a answer. The culture within the company should make you comfortable to reach out to anyone. **In the end, we are here to solve problems and get the job done**.

Of course, some people might say they do not have capacity to answer or help. That is okay. **What can't happen is the "you are not worth my time" type of answer**. If you notice this happening, the organization has deep problems. Perhaps that is an opportunity for you to change things around?

## No Wrong Doors

> Based on [this great blog post](https://lethain.com/no-wrong-doors).

Have you ever found yourself in a situation where you were redirected to another team before finding the right person to help you with something?

This behavior, while typical, is problematic. **It forces people to navigate the organizational design and to find the right people/team for questions they might have**. The problem with here is that **this process takes a lot of time, especially if you are new to the company**.

It would be much more optimal to **have the person you asked initially help you start the discussion with the right people**. If they do that, they **create a three-way conversation, where everyone learns**.

This _"let me help you even though you should aks X"_ approach is called "No Wrong Doors" policy. No matter where you initially end up, the team is willing to provide help even if you are in a "wrong" place.

Ultimately, we work together. We should help each other as much as possible. You never know when you will be in a similar situation!

## Tradeoffs as picking the "less bothersome drawback"

> [Based on this great blog post](https://buttondown.email/hillelwayne/archive/ive-been-thinking-about-tradeoffs-all-wrong/).

When making decisions, we often consider the positive aspects of each side—"X is faster, Y uses less space."

In software, it is not the positives that drag you down. The negative consequences of your decisions have the most significant impact on what you are working on. The negatives influence how you design things and how you move forward.

So, if you **compare alternatives only by their positive aspects, you are missing the main contributor to your future problems — the negatives**!

Instead of framing the question as "X gives me this benefit and Y gives me this benefit," think of "X is bad at this, and Y is bad at this." This way, you can consider the whole picture. You can employ effort to minimize the "bad thing" that your decision brought.

## Context > Deadlines

> [Based on this great blog post](https://newsletter.canopy.is/p/the-4-keys-to-creating-team-accountability)

Have you ever been in a situation where you were told that "X has to be done in Y time"? – I would argue most of us has.

I'm not bashing against deadlines, as they are an important to leadership and product.

- The leadership team is subject to the board expectations. The leadership people and people in the board do not like uncertainty. One way to reduce uncertainty is to have deadlines.

- The product team needs to know when a feature will be ready to plan for it's release. Perhaps they have an email campaign planned? Or maybe even a TV ad?

**The deadline becomes problematic when people responsible for the actual work lack the context WHY we have the deadline in the first place**.

- Without context, the deadline creates a sense of _false urgency_.

- Without context, the deadline is an abstract concept.

**Providing context alongside the deadline makes you question the actual deadline**.

- Before providing context, you are, most likely, going to question the deadline – "do we really need to implement X by Y? What are the implications of missing the deadline?". This could result in dropping the deadline altogether. A net win!

- Providing the context means making the deadline a tangible thing.

  - Perhaps we have TV ads lined up for X date?

  - Perhaps this very big customer of us has something planned for X date and needs Y to accomplish that? If we do not we will loose a lot of revenue.

**The more meaning the deadline has, the more motivated your team will be**. Of course, this not applies to everyone, and that is fine.
