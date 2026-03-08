# FEM Permissions Systems At Scale

1. [Notes based on this workshop](https://frontendmasters.com/workshops/permission-systems/).
2. [GitHub repository](https://github.com/WebDevSimplified/fem-permission-systems-that-scale).
3. [Course website](https://webdevsimplified.github.io/fem-permission-systems-that-scale/).

## Authentication vs. Authorization

Ahh, the never-ending source of confusion!

- **Authentication answers "who are you"**. Imagine _having_ a badge to enter a building.

- **Authorization answers "what can you do"**. Imagine _having a certain permission-level_ to enter a specific floor in a building.

So, when I log-in to a website, I'm going through the process of _authentication_. At the end of the process, I usually get some kind of token back. This token is then sent alongside any requests to other services to perform _authorization_ on a given resource.

From my experience, **one of the biggest issues is when engineers try to implement both _authentication_ and _authorization_ together**. This results in messy code and very hard to reason about logic that is hard to talk about.

Another issue I've seen is **engineers using _authorization_ and _authentication_ interchangeably**. This is problematic because it creates confusion.

**Authorization code, usually, changes a lot**. Every time you add a new feature, you might need to amend the authorization checks. In contrast, the **authentication code rarely changes**.

## The Initial Data Model

We started with three entities in the database. On the users, we have the "role" attribute which is an enum. I'm always wary of enums. I wonder how this will scale. What if I want to attach multiple permissions to a single user?

Similar to the `status` and `isLocked` on the document entity. Should those be two separate fields? If we allow multiple values in `status` we could have a document with `["draft", "locked"]` status which I think is much better for extensibility.

## Main Permission Mistakes

1. Scattered checks. This one is **so real**. I've seen those checks get very hard to reason about.

2. Missing permission checks. Builds on the first one. If you have checks duplicated across multiple files, it's _very easy_ to miss adding them.

3. Inconsistent logic. Again, builds on the first one.

Have you ever seen an issue where users can _edit_ but can't save? That's a signal that your permission checks suffer from those issues.

**Kyle argues that missing permission checks are WORSE than anything else**. I agree. IMO it is better for users to not be able to do something, which you can fix pretty fast, rather than everyone having the ability to do something that only a handful of users should be able to do.

## The first attempted fix

The first thing we did was to actually write those permissions checks inline. Think of manually checking `user.role` and then conditionally rendering based on that.

We _might_ have "fixed" some of the permission issues this way, but there is still a LOT of work we had to do. More importantly, the core issue remains: the checks are scattered throughout the codebase.

## Centralized permissions system

To help with duplication, we've created functions like `createDocumentService` or `updateDocumentService`. All of this to **reduce duplication**.

Now all of our checks of permissions happen only in those functions. DAL and other layers are not concerned with the permissions at all.

I'm still not a fan of this approach since we _have to remember_ to call the right function. Since the function names (can change, but then to what?) are quite similar, it might be the case that we will call the wrong function (without the `Service` suffix), and create a bug.

## Introducing Role-Based Access Control (RBAC)

Introducing RBAC requires us to shift our mental-model. Instead of manually checking for permissions, now it's the _roles_ that contain them!

```
User -> Role -> Permissions
```

The main benefits of using RBAC are:

- Centralization. You can have all the logic in a single place.
- Future-proofing. User can have _multiple_ roles and roles can have _multiple_ permissions. You can add / remove them as you see fit.
- Code is easier to read. You can do a "simple" `.includes` on the list of permissions and see if a given entity has access.

The drawbacks include:

- Additional complexity. You most likely need to add additional schema in your database to handle the Permissions (and Roles).

- Depending on how you structure your code, complex permission relationships might be hard to encode.

  Consider the `project:read` permission. We quickly expanded to `project:read:all` and `project:own-department`.

  If the permissions logic is even more complex, are you okay with extending the types like that? What about conditionals?

In the workshop, we **quickly hit limits of the "native" implementation of RBAC**. As soon as our permissions started to depend on other things, not only on the "role" of the user, the system quickly fell apart.

**If you notice your permission system quickly expanding the "permissions enum", it might be a sign that RBAC is not the way to go**. This is mostly due to permissions also depending on "surrounding context" rather than strictly on the "role" property on a user.

## Introducing Attribute-Based Access Control (ABAC)

This section started to look more and more similar to AWS handling permissions.

```
RBAC: User -> Role -> Permission
ABAC: (Subject + Resource + Action + Environment) -> Policy -> Decision
```

Notice that this **solves the issue we had with RBAC where we were not really able to check for "surrounding context" in a single permission call**. We can now create "templates" that we can "overlay" on top of users data to derive the `true/false` values.

While the implementation is a bit more complex that RBAC, it allows us to extend it without having to make additional "helper" functions to check for permissions.

```ts
type Resources = {
  project: {
    action: "create" | "read" | "update" | "delete";
    condition: Pick<Project, "department">;
  };
  document: {
    action: "create" | "read" | "update" | "delete";
    condition: Pick<Document, "projectId" | "creatorId" | "status" | "isLocked">;
  };
};

type Permission<Res extends keyof Resources> = {
  action: Resources[Res]["action"];
  condition?: Partial<Resources[Res]["condition"]>;
};

type PermissionStore = {
  [Res in keyof Resources]: Permission<Res>[];
};

class PermissionBuilder {
  #permissions: PermissionStore = {
    document: [],
    project: [],
  };

  allow<Res extends keyof Resources>(
    resource: Res,
    action: Permission<Res>["action"],
    condition?: Permission<Res>["condition"],
  ) {
    this.#permissions[resource].push({ action, condition });
    return this;
  }

  build() {
    const permissions = this.#permissions;

    return {
      can<Res extends keyof Resources>(
        resource: Res,
        action: Resources[Res]["action"],
        data?: Resources[Res]["condition"],
      ) {
        return permissions[resource].some((perm) => {
          if (perm.action !== action) return false;

          const validData =
            perm.condition == null ||
            data == null || // unsure about this check here. IMO that's a bug
            Object.entries(perm.condition).every(([key, value]) => {
              return data[key as keyof typeof perm.condition] === value;
            });

          return validData;
        });
      },
    };
  }
}
```

This allows us to check if we have a matching pair of `action` and `condition` for a given _resource_, like `document` or a `project`.

If you look at the `Resources` type, it is quite similar to how AWS does IAM policy statements + conditions block.

## Looking at individual resource fields

Currently, our permissions grant "general" `read`, `update` etc.. actions that pertain to _all_ fields of the _resource_. How do we scope it down to only allow, for example, reading the document `title` field? – we can extend our `allow` function to take additional parameter!

In the workshop, we did not check for the valid `action` given the fields. Nothing stops us from creating a `create` permission for document `title` field. On the one hand, that makes sense, on the other, I'm unsure.

## Permission libraries

Instead of writing the code yourself, you can rely on libraries that manage ABAC for you.

1. **In-code libraries**. This is what we've been working on. You write the policies in your language of choice. A popular choice for JavaScript/TypeScript is [CASL](https://casl.js.org/v6/en/). The main drawback here is that you _can't_ really share those policies across teams working in different languages.

2. **DSL's for writing policies**. You have a completely separate language for writing policies that **you can share and centralize**. A good example is the [Cedar](https://www.cedarpolicy.com/en) language that AWS uses under the hood.

## The spectrum

1. Our first attempt was a big ball of mess. We did not have centralized place for permissions and it was _very easy_ to introduce an issue.

2. Centralizing permissions helped a bit. We still had to write a lot of code and handle each case separately, but it was definitely an improvement.

3. Introducing RBAC looked promising, but we quickly hit roadblocks when dealing with "related" data and how it influences permissions (for example, certain users can only edit documents from their own department).

4. ABAC was the "solution" to all of our problems, BUT it introduced code complexity (albeit neatly tucked away in a single module).

5. Then we reached out for a library that "compressed" that complexity even further.

**Note**: While we were mostly working in-code, we also had to **amend our database queries**. So in a sense, you have to make sure you have permissions present in two layers: the application and database layer.

In an ideal world, you would have only a single layer that would sit between your application and the database that performed all the checks. But then how do you make sure you are not leaking permissions information metadata _into_ your database?

## Random Next.js learnings

The project we are working on throughout the workshop is using Next.js. Here are some things that stood out to me.

1. The `(<directory_name>)` folders in the `app` directory will NOT create a new route path. These are so-called _route groups_ that you can use to have multiple `layout.tsx` files per routes without creating additional paths.

2. The `cache` function is _per-request_. It's interesting to see that you do not have to pass any `requestId` to it or anything to work.
