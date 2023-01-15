# Random Things I've picked up while studying

## Rolling back vs. rolling forward

You might have heard the term _roll back_ before in the context of CI/CD pipelines. It is when you want to bring the system to the old state before a change. **This is the most common way of restoring the system in the case of an emergency**, where a change introduces a bug.

But what happens if the **change latent change – a change introduced a while back that now sits sandwiched between other changes?**
In such cases, the engineer must **pragmatically decide whether it is worth rolling back all the changes, including the latent change, OR is it worth rolling forward**. _Rolling forward_ means creating a change to fix the underlying issue and letting it advance through the CI/CD pipeline. This ensures that we do not roll back valuable releases, like bug fixes.
