# Build AI Agents that Never Sleeps

Based on [this course](https://master.dev/workshops/background-agents/)

## Welcome to background agents

When you interact with agents synchronously, you are bound by the "prompt speed". Yes, the agent might work for hours on end, but it's you who has to start it. The work happens only when you explicitly say so.

Background agents flip this relationship. The agent is now autonomous and can react not only to your events (the input prompt), but also to things that happen in the "outside world", which is usually some kind of system.

## Give the agent a goal

In this section, Scott talked about how _giving the agent a view into the current state of the "world"_ is a requirement for building a good background agent. The agent has to know what is going on so it can make proper decisions based on that state.

Since we control the harness, we can inject the "current state of the world" into the prompt. Remember: the prompt might not come from us typing it. It might be composed based on an event that triggered a certain code path.

I've noticed that Scott, in pretty much every agent-related workshop they lead, uses some kind of _durable execution_ layer. In this case, it's [Inngest](ingest). It's quite interesting to me, because I have not had a chance (or felt the need) to use this technology at my work. Perhaps I'm missing something?

In the app, we have one spot where we ask the LLM to derive the next action we ought to take based on the current state. Scott implemented this flow by calling a regular text model. The nice thing about text models, as opposed to models like Jev, is that they can _elaborate_ on their choice. You can force the model to output the `reason` field, **which you can't really get from a classifier model like Jev**.

## Make progress durable

That is what using a "durable execution layer" gives us – we can _resume_ the agentic loop after bad things happen, like a server crash.

But this level of robustness does not come without a price. **Every "step" in the workflow needs to be idempotent** and **the workflow itself needs to be deterministic**.

- **Deterministic** means that the same input always produces the same output.

  - `2+3` is deterministic. `Date.now()` is not.

- **Idempotent** means that running a given function once produces the same results as running it multiple times.

  - Those function calls _can_ have side effects, but the side effect must not change the state of the world after it runs for the first time.

In addition to what we have talked about so far, you will also need to reason a bit more about the code. The question "should this be its own step?" will be asked often.

A good heuristic seems to be to put every `async` operation within a step. You _might_ want to put multiple `async` operations within a step, but only if you are okay with _both_ of those operations running together when the step retries.

LLM calls are also a good example of something that you most likely want _inside_ the `step` function. LLM calls are not deterministic, even with a strict schema.

## Wait for the world

The abstraction we use for making the workflow durable and resumable exposes the `waitForEvent` function, which we can use to pause the execution until a given event happens or a certain amount of time elapses (whichever comes first).

We have to be mindful of how _long_ we want to wait, and of the race condition where the event we want to wait for fires _before_ we call `waitForEvent`. Nevertheless, the Inngest SDK makes it quite easy.

## Put a human in control

Scott makes a **great** observation that making your agent aware of the fact that certain tools require human approval _might_ influence which tools it calls. For example, it might decide against calling such tools.

**That is why you want to make sure the agent is NOT aware of the HITL implementation**. No tool should be exposed as "requires human approval". All that logic should live in the harness and should be transparent to the agent.

In the app, you have an "inbox" where all the proposals go to. You are supposed to either approve or reject the proposal. That got me thinking: "How do I ensure I always have enough context to make the change?"

See, I often find myself flooded with questions from the LLM, especially when it works on a plan. That's okay – I would rather answer those questions now than be surprised after implementation.

BUT, given that I'm working on multiple things at the same time, I often feel like I do not have enough context to _answer_ those questions. The trend I see here is to produce more _visual_ artifacts than written text – HTML pages, videos, etc.

Of course, visuals cost much more tokens to produce than text, so there is a trade-off here.

## Retries & Handoff

Retries are an interesting topic. Whenever we retry, we have to consider _idempotency_. You can say that the service / function / whatever is idempotent if, given the same input, it returns the same results. For backend services, this also means we do not want to change the "state of the world" in the second or nth request – we "just" want to return the result as it was returned the first time.

The problem arises when we are dealing with a service that is NOT _idempotent_. When a request returns 503 (Service Unavailable), what do you do?

1. You might choose to retry the request, but then you might re-create some entity, or re-send the notification.
2. You might choose to _skip_ making another request, but then you risk not performing the side-effect that the request is supposed to cause.

I would err on the side of retrying the request and ensuring the downstream can handle potential duplicates. IMO, it's better to send two notifications than not send any.

## Wrapping up

This is yet another course led by Scott where we use some kind of "durable execution" engine. I find it interesting that the technology choice seems to be constant, at least from Scott's perspective.

If durable workflows are so awesome, I wonder _why_ I never felt the need to use them? I guess it's probably because I was not aware that I needed them when I _could_ have used them, and went with an alternative that might not have been such a good choice.
