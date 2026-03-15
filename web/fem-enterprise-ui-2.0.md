# FEM Enterprise UI 2.0

## Architecture Patterns

- **Start small and simple**.

### The Monolith

Most people associate "the monolith" architecture with something negative. I'm unsure its a fair thing to do.

You can go really far with this architecture. It is often "the default" for many teams.

- Monoliths are the _simplest_ things to deploy. There are no dependencies!

- All components of the architecture (routing, state, styles, tests, dependencies) are tightly coupled. While it is usually not a good thing, in this case, it actually is! It is easier to reason about things that way.

- Most tooling we use is set-up in a _single repository_ and single "application" in mind. This makes it easier to start adding new tooling to the monolith.

For example, Next.js applications are, usually, monoliths.

---

Of course, **as the monolith grows, you might start to see the following problems:**

- Team collisions, where multiple teams edit the same files.

  - Team dynamics also might play a role here. You have multiple teams with different bosses and timelines attempting to meet their objectives.

- As codebase grows, the build-time and deploy-time also increases. It might get to the point where deploying takes forever.

- Blast radius of a given change starts to get pretty big. If you do not have clear code boundaries, you might accidentally change a function that is used by completely different team and cause an incident.

- Upgrading dependencies is almost impossible, given the amount of changes you would need to do. Good luck upgrading framework versions!

### The Three Axes

We can look at the following

- Runtime architecture.

  - Is is a single application (in this case a frontend application) monolith or a multiple microfrontends stitched together?

  - Remember: the more moving parts, the more "freedom", but the more complexity.

- Repository topology.

  - _mono_ or _poly_ repository?

  - _Poly_ repositories means each team own their own repository. **Here you get a LOT of freedom, but that might also be the problem**. Think people using completely different frameworks.

  - _Mono_ repositories means each team contributes to the same repository, but the repository contains multiple applications or packages. **Less freedom, but usually that's a good thing**. You can enforce consistency more easily.

- Deployment topology.

  - **The worst thing that you can have here is multiple applications with lock-step deployment scheme**.

  - If you want to have separate applications (_poly_ repository), perhaps it would make sense making sure you can actually deploy them independently?

Finished part 1 50:04
