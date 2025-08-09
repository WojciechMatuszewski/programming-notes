# FEM Tailwind CSS V2

## Learnings

- In essence, the "Tailwind compiler" will look at your files and attempt to gather all classes that you use.

  - Then, given those classes, it will produce a CSS file that contains definitions for those classes.

  - **This works well only if you are not using _string interpolation_**. If the class it not "whole", the compiler won't "notice it" and it might NOT be included in that final CSS file.

- Tailwind CSS will inject _CSS Layers_ into the CSS file.

  - If you want to extend the styles, consider adding the styles to a given layer.

- **Some anti-patterns** to consider:

  - Creating one-off classes.

    ```css
    @layer components {
      .my-button {
        background: blue;
      }
    }
    ```

    The problem with this one is that many utility "hooks", like `:hover` won't work if you write `.my-button:hover`.

    You probably want to do it like so:

    ```css
    @utility my-button {
      background: blue;
    }
    ```

  - Using arbitrary values in classes, aka the `[]` syntax.

    ```jsx
    <div className="bg-[#121131] p-[123px]"></div>
    ```

    This might be problematic while refactoring the codebase. Imagine switching colors defined in variables and having to update those one-off cases.

    I believe one should stick with values defined by Tailwind, and if you need other variables that you use, especially for spacing, make sure to add them to `@utility` or other relevant layer.

    ```css
    @theme {
      --color-brand: xxx;
    }
    ```

- When you add a border to an element upon hovering over it, the element might _shift_ a little bit (border adds size to the element).

  - **The alternative is to use `outline` which will NOT shift the element when applied dynamically**.

    - It **used to be the case** that `outline` did not respect rounded borders. That is no longer the case!

    - Since older browsers might still have an issue with rounded `outline`, **some CSS libraries implement "rings" which is a `box-shadow` that looks like a border**.

      - Again, this is all to make sure you can have a rounded border on an element without "shifting" it when applying it dynamically.

- The neat thing about "the `ring` technique" is the fact that you can stack them!

  - That is not possible with `outline` without wrappers. The element can only have one outline.

    ```html
    <button class="rounded-md ring-blue-500 ring-offset-2 ring-offset-green-300 hover:ring-2" />
    ```

- It's always good to remind ourselves about the existence of `:user-invalid` and `:user-valid` pseudo classes.

  - The `:invalid` is problematic because **by, default, even if users did not interact with the form, the field would be treated as `:invalid`**.

    - The `:user-invalid` works a bit differently. **The `:user-invalid` styles only apply AFTER user interacted with the input and that input value failed the native validation**.

      - By "native validation", I mean things like `type="email"` or `required`.

- Another "gem" people might not be aware of are the `:focus-within` and `:focus-visible` pseudo classes.

  - `focus-visible` is like `focus` but only when you focus the element via keyboard (or other non-pointer).

  - `focus-within` is **looks for `focus` on any of it's children**.

    - It's NOT akin to having `focus-visible` on the child, but `focus`.

      - If you need something like `focus-within-visible` you can use `has` selector, like `has-focus-visible:xxx`

- The `has` is not an _alternative_ to `group`, but rather a supplement for it.

  - With `has` you can style the parent based on the children.

  - With `group` you can style the children based on the parent.

  The best part? **You can combine the `has` and `group` together!**.

  ```html
  <!-- Switch from `display: none` to `display: block` when a particular input in the group is checked -->
  <div className="group-[:has(#yes:checked)]:block hidden" />
  ```

- One of the neat things in Tailwind v4 (available in v3 via plugins) is a native support for _container queries_.

  - _Container queries_ are one of the most awesome features released in recent years!

## Wrapping up

A high-level overview of the library. If you are new to Tailwind, I would highly recommend this course to get up-to-speed on it.
