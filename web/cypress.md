# Cypress stuff

## The power of `cy.task`

When you are first starting out with `Cypress` you might feel a bit constrained. You will quickly notice that you do not have access to the node environment (at least within you test block). **`Cypress` commands run inside the browser**.

There is an escape hatch though. Something that allows you to run arbitrary code in the **`Cypress` backend server**.

```ts
cy.task("nameOfTheTask", serializableData);
```

The main restriction here is that the data has to be _serializable_. This is due to how `Cypress` sends your data to it's backend.

## Retries and assertions

Some of the `Cypress` commands are automatically retried up to a given timeout. This helps with UI tests since sometimes we have to deal with network delays, our machine speed and such.

The most important piece of information to know here is how this mechanism work when combined with assertions.

Let's say you have the following

```js
cy.get(".todo-list li") // command
  .find("label") // command
  .should("contain", "todo A"); // assertion
```

Let's say that the first `command` is successful and the `li` is found. This means that **only the second command will be retried until the assertion passes**. This is very important, let's say we are adding an item before executing the chain

```js
cy.get(".new-todo").type("todo A{enter}"); // add item
// What if there is a delay here and there already is an `li` within `todo-list`?
cy.get(".todo-list li") // command
  .find("label") // command
  .should("contain", "todo A"); // assertion
```

If there already is an `li` within the `.todo-list` but with different label, and there is some delay between clicking _enter_ and `todo A` appearing, the whole chain will fail.

First, the `cy.get(".todo-list li")` will succeed, because there is an item in the list already. The `.find('label').should(...)` will be retried **but in the context of the first item. That item might not have `todo A` as text content!**.

So, how do we make sure our assertion chain passes? There are 2 ways.

### Selector merging

Instead of doing 2 commands, we can perform 1 command on the merged selector.

```js
cy.get(".todo-list li label") // command, will be retried until the assertion passes (with timeout)
  .should("contain", "todo A");
```

This makes sure that we are executing the whole selector each time the `should` fails, not only looking at the label.

### Alternative assertions

I find this solution a bit artificial but either way, it's a solution.

The deal here is to split the chain we had previously to ensure that all preconditions are met, then proceed with our final assertion.

```js
// assume that there is already 1 item in the todo-list
cy.get(".new-todo").type("todo A{enter}"); // add item

cy.get(".todo-list li") // command
  .should("have.length", 2) // assertion
  .find("label") // command
  .should("contain", "todo A"); // assertion
```

Now, the first assertion will pass only if the item is actually added. Then we can proceed with our `label` assertion.
