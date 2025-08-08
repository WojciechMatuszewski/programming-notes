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

Start part 6
