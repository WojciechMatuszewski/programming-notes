# FEM Backend System Design

> Learning and taking notes from [this workshop](https://frontendmasters.com/courses/backend-system-design/).

## Requirements

- **Translating business requirements to functional and non-functional requirements is the main focus point**.

  - Instead of jumping into details, **ask yourself: what are we actually trying to solve here**.

  - You might be presented with a problem that seems obvious. **The underlying complexity might not be architecture, but the company landscape and business needs**.

- **Tradeoffs and tradeoffs**. Perhaps using a monolith is fine for this particular case?

My first attempt at specifying requirements for a "todo app" failed. Why? **Because I jumped into thinking about functional and non-functional requirements without stepping back and thinking about the _domain_ or _why_ we are building the app in the first place**.

- Scope the problem.

- Design the high-level architecture.

- Do not make assumptions.

- **Requirements could change over time. Do not "hardcode" constraints unless allowed by the interviewer, ASK**.

- You can "rally" your thoughts around a core concept. For example, for an URL shortener system, that might be the _URL_. For a banking application, that might be the _transaction_.

  - If you focus on one single "entity", you can create functional and non-functional requirements "around" this entity. This also helps with scope.

- At this stage, you might also start thinking about how the design of the application will look like.

  - **By being explicit with questions and requirements, you have a better chance of nailing down the system design**.

---

CAP theorem tells us the tension between:

1. Consistency which is all about reading written data.
2. Availability which is all about making sure every request gets a response.
3. Partition Tolerance which is all about how the system functions when network communication between nodes is is lost or delayed.

The theorem states that you can't get all three. You must, to some extend, "sacrifice" one to get the other two.

C + A is a single server. As soon as you introduce nodes to the system, you introduce _network partitioning_.

C + P could be a distributed server but all reads would need to wait for a given write to propagate to a different nodes. This might be slow.

A + P is a distributed system where reads might return stale data.

Note how this theorem influences system design!

**In distributed systems, the `P` is implied. You can only pick "additional" one to it**.

Consider

1. For a banking app, you might want consistency.

2. For link-shortening service, you might want availability.

3. For medial team ToDO application, you might want consistency or availability.

...

---

Start https://frontendmasters.com/courses/backend-system-design/system-quality/
