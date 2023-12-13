# FEM Tailwind CSS

- The `@apply` allows you to compose tailwind-specific classes in your own class definitions.

  ```css
  .my-custom-class {
    @apply text-2xl text-white;
  }
  ```

  This is a nice way of encapsulating some repeated declarations into one. I would advice caution here – the DRY is good, but only if it's applied to some things that are actually the same, and not only pretend to be the same at this very moment.

  One example use-case for the `@apply` are component libraries. Check [the daisyUI out](https://github.com/saadeghi/daisyui/blob/master/src/components/styled/alert.css).

- Keep in mind that there are escape hatches available to you if some of the defaults do not work.

  1. You can always change / replace the defaults within the config file.

  2. You can use the brackets syntax.

  ```html
  <div class="h-72 w-[375px] bg-blue-200 [margin-inline:40px]">
    <div class="mx-20 h-40 w-20 bg-red-200"></div>
  </div>
  ```

  Notice that I can either start with the tailwind-specific prefix (`m-[375px]`) or write entirely custom definition in the brackets (`[margin-inline:40px]`).

- For colors, consider using [this page](https://uicolors.app/create). It is pretty great!

- The initial CSS you get via tailwind uses, in modern browsers, `@layer` directives (I assume the _postcss_ preprocessor will do it's thing when building for production). This means that you can extend them!

  ```css
  @tailwind base;
  @tailwind components;
  @tailwind utilities;

  @layer components {
    .btn {
      @apply rounded border-2 px-2 py-1;
    }
  }
  ```

  And now you can go ahead and use `btn` class in your HTML. **I did not have to put the `.btn` definition inside the `@layer`**. This is only to ensure the proper hierarchy of CSS specificity is preserved.

Finished day 2, 28:54 https://frontendmasters.com/workshops/tailwindcss/ https://tailwind-workshop.vercel.app/variants
