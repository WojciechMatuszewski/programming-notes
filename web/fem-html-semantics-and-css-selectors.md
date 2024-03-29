# FrontendMasters – HTML Semantics & CSS Selectors

Notes from [this workshop](https://frontendmasters.com/workshops/semantics-selectors/).

> Slides & website [here](https://semantics-selectors.css.education/).

## Why do HTML CSS matter?

- HTML is not responsible for the display of the page!

  - HTML is NOT used for page styling.

  - HTML comes with many hidden features for SEO and such.

- We use very few number elements that are available to us. This impacts SEO and accessibility.

## Lists and Selector Review

- The `box-sizing: border-box` makes the most sense but is not the default, most likely, due to legacy considerations. Imagine all those websites breaking if the `border-box` is the new default (as opposed to the `content-box`).

- **Some CSS properties**, like `font-family` or `color`, when applied to parent, **will propagate to children**.

- **To capture the hierarchical nature of CSS**, one might want to **read CSS selectors from right to left**.

  ```css
  ul a {
    color: red;
  }
  ```

  Means _"select all elements that are descendants of the ul element"_.

### Child selectors

- Some of the selectors the instructor showed us:

  - The **direct child selector** -> `>`

  - The **general sibling selector** -> `~` (the siblings can be separated by multiple elements).

  - The **direct sibling selector** -> `+` (the siblings must NOT be separated by any other elements).

### Attribute-based selectors

We also went through some of the attribute-based selectors. It is pretty helpful to remind yourself of their existence.

- The attribute selector -> `[attribute_name]` (the `attribute_name` MUST be present on the element). Think `href` on the `a` tag.

- The attribute value selector -> `[attribute_name="attribute_value"]`. Think `href` on the `a` tag with a specific value.

- You might also select based on **parts of the value** by using the **substring selector** -> `[attribute_name*="attribute_value"]`.

### Pseudo-children selectors

One cannot forget about the _pseudo children selectors_ (the `nth-child` and so on). Here are a few notable ones.

- `nth-children(even)` or `nth-child(odd)`. Pretty sweet! I used to do the magic `2n+1` thingy for those before.

- `nth-children(2n+1)` -> **start with the first element, then, select every second after that**. I like substituting the `n` with real numbers like 0, 1, 2...

- `nth-last-child` -> **instead of starting from the top, start from the bottom**. You can apply the same modifiers as in previous examples. I have never used this one before.

### Description lists

From what I could gather, the _description lists_ are for grouping `key:value` pairs of content. Think categories and items that belong to a given category.

To style the _description lists_, the `grid` type of layout should most likely be the "go-to" as it nicely works with the `key:value` pair model of the elements within the _description list_.

The instructor discussed the difference between the `first-child` and the `first-of-type` selectors in this section. The **`first-child` selects the first child of a given element**. The **`first-of-type` selects the first type of a given element within a given element (it does not have to be a first child)**.

```html
<div>
  <span>foo</span> span::first-child OK, span::first-of-type OK
  <h1>bar</h1>
  h1::first-of-type OK, h1::first-child NOPE
</div>
```

### Menu lists

- The **`menu` element is designed to wrap interactive elements, like buttons, but NOT LINKS!**. Very useful to know. I was in a situation where I wrapped multiple buttons with `div` instead of the `menu` tag.

  - The **`menu` tag is a new list type, like the `ol` or `ul`**. To properly wrap the buttons with the `menu` tag, you would do something like the following.

  ```html
  <menu>
    <li><button>First</button></li>
    <li><button>Second</button></li>
    <li><button>Third</button></li>
  </menu>
  ```

### Takeaways from the exercise

- The `dl` is quite powerful, and if you know that the tag exists, you will most likely use it more often than not.

- Do not forget about the `article` or `section` tags – they are there for a reason. Remember – the `article` is for content that makes sense even if you were to pull it out of the website. The `section` is for content that would not make sense in such a scenario (it is more contextual).

## Parts of the Web Page

- You **can use `header` and `footer` tags inside the `article` tag**.

- `div` and `span` elements do not mean anything.

- Proper usage of `section` and `article` tags is very powerful. It makes the HTML more readable.

### Pseudo-classes – `is()`, `where()`, `not()`, `has()`

- The `:is` and `:where` selectors are used to eliminating repetition that sometimes you have to deal with when selecting multiple children of a given element.

  ```css
  :is(article) h2,
  h3,
  h4,
  h5,
  h6 {
    color: var(--fire-red);
  }

  // the same as \/
  article h2,
  article h3,
  article h4,
  article h5,
  article h6 {
    color: var(--fire-red);
  }
  ```

  - The only difference between the `:is` and `:where` selectors is the specificity. The **`:where` selector has a specificity of 0, whereas the `:is` adheres to the regular specificity rules of selectors**.

- The instructor showed a fascinating use of the newly `:has` selector – specifying the number of columns based on the list items. Consider the following.

  ```css
  ul:has(li:nth-child(6)) {
    columns: 2;
  }

  ul:has(li:nth-child(11)) {
    columns: 3;
  }
  ```

  Pretty amazing stuff if you ask me.

## Inheritance

- Some properties pass their properties to their children – for example, the `color` property.

  - You can view the complete list of properties that inherit [here](https://www.w3.org/TR/CSS21/propidx.html)

- The `:root` selector is the `html` selector. We usually define variables within that selector to ensure we can access them globally.

  - Keep in mind that **the `:root` has a higher specificity than the "bare" `html` selector**. Because of that, people usually use the `:root` instead of the `html` selector.

- For properties that do not inherit, you can **force inheritance programmatically via the `YOUR_PROPERTY: inherit`**.

  - For the inheritance to work, **you have to have a direct parent-child relationship**. The `:inherit` declaration must be defined on the immediate child.

- Inheritance is excellent, but sometimes you might want to **turn the inheritance off**. Luckily, some properties can help you with that.

  - Use **`all: initial` to strip ALL inherited properties**. This also means the styles inherited by the browser agent!

  - Use **`all: revert` to "revert" any custom inherited properties**. Use this one to _revert_ to the original browser agent styles.

  - You can also use these values for specific properties, like `color` or `font-family`.

## Cascade

- The cascade mechanism (the order in which the styles are evaluated) affects and can overwrite rules with very high specificity.

### Layers

- Layers are a container of CSS rules. You can programmatically order the layers the way you want, giving you control over the cascade mechanism (in the author land).

  ```cs
  @layer addme, standard

  @layer standard {
    // definition
  }

  @layer addme {
    // definition
  }
  ```

  The `addme` layer will be evaluated first, despite being second in the CSS file.
