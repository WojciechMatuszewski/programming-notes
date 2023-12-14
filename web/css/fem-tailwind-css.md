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

- Remember about the `accent` group of classes! They are very handy for styling checkboxes.

- There is a difference between `:invalid` and `:user-invalid`.

  - The `:invalid` **applies to all states of the input, even if the user did not interact with it yet**.

    - This makes it hard to work with as the visuals will be applied to all inputs when page loads. To work around it, you might want to use the `:not(:placeholder-shown)`.

  - The `:user-invalid` **is mostly what you want – it only applies when user interacted with the input**.

- There is a difference between `:placeholder-shown` and `:empty`

  - The `:placeholder-shown` is for inputs.

  - The `:empty` matches all elements that actually are empty – they do not have any children.

- The _peer modifier_ could be used to style **next siblings of a given element**. This is quite useful for error messages.

  ```html
  <input
    type="email"
    name="email"
    id="email"
    required
    class="peer"
    placeholder="email"
  />

  <div class="invisible text-red-500 peer-[:user-invalid]:visible">error!</div>
  ```

  At the time of writing this, Tailwind does not natively support the `:user-invalid` so I have to use an escape hatch.

- The _group modifier_ allows you to **style descendants of a given element**.

  ```html
  <div class="group h-24 w-24 bg-red-500">
    <span class="decoration-purple-100 group-hover:underline">foo</span>
  </div>
  ```

  In this case, the `span` will have the underline applied when we hover over the `div`.

- You can use the `:dark` variant to implement different styles based on the "dark mode heuristics".

  - The default way Tailwind detects the dark mode is via media query.

  - Of course, you can change this behavior.

- There is a **handy `container` class** which automatically applies the breakpoints.

  ```css
  .container {
    width: 100%;
  }

  @media (min-width: 640px) {
    .container {
      max-width: 640px;
    }
  }

  <!-- And so on... -->
  ```

- Steve showcases an interesting CSS definition – `columns`. You can use it to create... column-like layouts.

  - It kind of work like masonry layout, but not quite.

  - BTW, there the is the [experimental masonry layout support in CSS Grid](https://developer.mozilla.org/en-US/docs/Web/CSS/CSS_grid_layout/Masonry_layout).

- While playing around with CSS grid, I've noticed that the `grid-cols-NUMBER_OF_COLUMNS` produces CSS snippet that I did not expect.

  ```css
  .grid-cols-2: {
    grid-template-columns: repeat(2, minmax(0, 1fr));
  }
  ```

  Why not `repeat(2, 1fr)` you might ask? **By default, the `repeat(2, 1fr)` will not produce 2 equal-size columns!**.

  The `fr` unit **represents a fraction of available space in the grid container and, by default, minimum width of a grid column is `auto`**.
  This means that, if a column content is bigger than the available space, it might overflow! See [this GitHub issue](https://github.com/rachelandrew/cssgrid-ama/issues/25). This issue is **similar to the flexbox overflow when specifying the `flex-basis`**.

- Writing complex grid layouts might be worth defining in a CSS file rather than inline. It gets hard to read.

- Steve showcases the **`user-select` (`select-X` in tailwind) CSS definition**. Holy smokes it is very useful.

  - You most likely wanted to copy some ID from somewhere right? Did you have a hard time selecting the whole ID string? I know I have.

    - With [the `user-select: all`](https://developer.mozilla.org/en-US/docs/Web/CSS/user-select) the browser would select the whole string if you clicked on it! Pretty neat.
