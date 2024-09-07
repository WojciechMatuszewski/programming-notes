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

Finished "When to use backend architecture designs"
