# FEM Introduction to Backend Architectures

## Learnings

- Starting with implementation is not a good idea.

  - First, create the blueprint, then start laying the bricks.

- There are a couple of key components of a good architecture.

  - _Interfaces_

  - _Modularity_

  - _Data flow_ which is impacted by the first two.

- Drawing the designs is a skill. **You can go far with good drawings as it makes the design approachable for a lot of people**.

- **Outdated architectural documentation is a problem of culture**.

  - Sadly, this is uncommon. People tend to lean more towards "let us build now!" approach.

### On Modularity

> Each component of the system should have a specific task.

### On Scalability

> Systems should be designed to handle growth of traffic.

### On Robustness

> Systems should be able to handle errors or unexpected situations.

Error handling plays a crucial role here. Sadly, this part of engineering is often neglected in favour of _features_. This strategy might work for some time, but it will backfire when engineers will be clueless when looking at the errors the application produces.

- **Consider building your own "error language"**.

  - It is not for users, but for **other systems that might use your service**.

  - On FE, you can map the errors however you want.

    - The "custom error codes" that you invent, will be very helpful for support team when user reaches out to them. They should be able, given the error code, trace back to the issue within a given service (of course, some details will only be known to the team owning the service).

- **Do not log everything**. Log only what you need, but lean towards logging more and **pruning your logging**.

### On Flexibility

> Systems should be designed to accommodate changes or future features.

### On Key Challenges

- **Think of the _complexity_ of the system when you are building it**.

  - Too much complexity is not a good thing.

  - Do not try to be too smart. **If your team does not understand the design, YOU failed, not them**.

- **Change is constant. Design for _adaptability_**.

  - Do not overload developers, but do not be afraid of making decisions.

- **Security is a must**.

  - Even if the service lives inside "internal architecture", there should be a layer of authorization and authentication.

- **It is up to you to make good technology choices**.

  - Build, learn, experiment.

  - Keep on learning to widen your horizons.

- **Consider the cost of architecture and engineering hours**.

  - Time costs. A lot. Consider the ratio between time spent and the robustness of architecture.

---

- Educate and document the solution you want to build.

  - The more people know how the system works, the better.

    - By educating and "making yourself replaceable", you increase the chances of the system succeeding – brining it money or saving time.

- **Learn how to "sell"**.

  - While it might sound weird, **as an engineer, you are also in business of selling to other people**.

    - You might not be selling "regular products", but **you are selling ideas to others**.

      - To "sell an idea" (some kind of design or a change you want to introduce), you have to convince others that this is a good idea!

## When to Use Backend Architectures

- Drafting architecture is a good idea when you are just starting on a project.

- When you are faced with complex problems. Drafting the architecture would enable to break the problem into multiple parts.

- To reduce the communication overhead between people and teams.

  - **Instead of repeating yourself over and over again, you can link to a doc containing most of the information**.

- **When preparing any technical document (or any other document for that matter), always think about _who you are writing for_, _what_ and _why_**.

  - You **need to have a reason for it**. Do not introduce complexity for the sake of it.

- Consider user feedback – sometimes, the complaints you get are the result of the architecture, not UX.

## How to Implement Backend Architectures

- Three stages

  - Research

  - Implement

  - Iterate / maintenance

- **Aggressively scope the architecture/research**.

- Play around. See how the approach feels.

  - **Make sure to set requirements BEFORE you do that to avoid scope creep**.

- Leave a "breadcrumb" trail of your work.

  - It is important to have a historical record on _why_ you did X and used Y.

- **Team expertise also comes into play here**.

  - It might be much harder to "sell" to the team if you want to use language/approach that the team is not knowledgeable in.

    - **Having said that, remember that the change is sometimes necessary and people will resist it**.

      - If you have a good reason for changing the status quo, you should be able to proceed.

- **Share your work early and often**.

  - The more feedback you get, the better the outcome will be.

  - At this stage, it is time to drop your ego and think about doing what is best for the product and not your career.

### Best practices

- Involve every stakeholders from the beginning.

  - It is never too late to reach out to people to ask for opinion.

- Use modular approach.

  - **Be honest with yourself**. Your first pass might overcomplicate things and that is okay.

- Consider how secure your design is.

  - **Involve people from the security team into the design process**.

## Common Backend Architectures

The silver lining here is that each architecture is viable given it fits the requirements and the team is comfortable running with that architecture.

### Monolith

- Everything is a single unit.

  - That is a huge advantage, as you spent less time switching between different layers and repos.

  - But, it is also a huge disadvantage when you reach certain scale. **The bigger the monolith, the harder it is to scale**.

- Very common since that is usually how you start.

- Monolith is easy to deploy.

- Monolith tend to be consistent. Usually, all the requests are flowing through the same code path.

- Monolith is **limited in scalability**. You can't scale services separately.

- **To deploy, you have to deploy _everything_**. Not a good thing.

- Security concerns might be hard to get right due to surface area and potential blast radius.

#### Use Cases

- Small applications and startups.

  - If the application is very successful, you might need to change the architecture later on.

  - **Nothing stops you from developing a _modular_ monolith** where components are deployed together, but they are not necessarily coupled together.

- Applications where high performance is critical.

  - **Distributed architectures are usually slower due to network overhead**.

### Distributed / service-oriented

- The blast radius of the change is smaller because you update one small component of a larger, loosely connected system.

- It is much easier to scale this architecture.

- It could be easier to add a new piece of functionality to an already existing net of services.

- **Microservices allow for the team to take much smaller surface area**.

  - This means less overhead and more focus on getting the microservice working just right.

- **The fact that you can deploy the systems separately is huge**.

  - Each team can decide when to deploy. There should be no need to coordinate deployments!
