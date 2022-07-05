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

  ```cs
  ul a {
  color: red
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

<!-- Finished part 5 16:30 -->
