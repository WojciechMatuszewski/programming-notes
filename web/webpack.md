# Webpack

## Federation

This one, in my opinion, is one of THE features of `webpack 5`. Like really, imagine sharing code, true micro-frontend style, **without code duplication, 2 versions of react, bundle bloat**. Just... amazing.

The system is still in beta while writing this, but there are plenty [examples](https://github.com/module-federation/module-federation-examples)

### What it's all about

It all comes down to code-sharing. You can `expose` some parts of your code `React Hooks` also work! to other peoples apps. And this is **not a library**, so local development is quite nice.

I've build a sample demo where I've deployed a sample site on `surge` which was exposing a `Button` component.

Then, I've consumed that `Button` on a separate app on a localhost, **this is just insane**.

### Usages

It's still quite early but there are a few usages:

- CMS previews. Image having to enter data inside CMS and wanting to have a preview available.
  You have one team building the actual site, and the other the CMS ingestion site. With federation the second team does not have to replicate the actual site. They can just use federation

- Footers and headers. Nothing to say here really

- A/B testing across multiple sites. One package shares all the logic / components. Acts like a separate site. 🤯

## The `require.context` function

- This is a **webpack-specific API** that does not exist in node.

- This **API allows you to produce paths for nested directory structures**. It is quite powerful.

  ```js
  const pages = require.context("./pages", true, /\.js$/u, "lazy");

  const render = async () => {
    const { pathname } = window.location;

    if (pages.keys().includes(`.${pathname}`)) {
      renderHtml(await pages(`.${pathname}`));
    } else {
      renderHtml("Not found");
    }
  };
  ```

- **The `import` API cannot do the same thing the `require.context` does**.

  - You cannot "produce" paths with `import` ahead of time and use those paths to import components asynchronously.

  - This is super handy when you have a list of well known import paths, but the list is huge and there is no easy way to "pre-compute" them. Specifying each import by hand would be very tedious.
