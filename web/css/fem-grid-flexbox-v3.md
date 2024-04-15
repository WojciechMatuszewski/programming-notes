# FEM CSS Grid & Flexbox v3

> Taking notes while watching [this workshop](https://frontendmasters.com/workshops/css-grid-flexbox-v3/).

Finished Part 5 23:43
https://grid-flexbox.css.education/ch6.html

## Intro to Grid

-   Grid is best for:

    -   A series of boxes that are the same width.
    -   Siblings that need to overlap or cover each other.
    -   Layouts that span multiple rows and/or columns.

-   `1fr` stands for _one fraction_.

-   The `gap` is a shorthand for both directions (columns and rows).

    -   You can specify each direction separately.
        -   `column-gap: 1rem;`
        -   `row-gap: 2rem;`

-   When you think about it, _grid tries to be quite helpful_. In some cases, you do not have to define _rows_ at all!

    -   Let us say you have four items, and you define two columns. The grid will "assign" the other two items to the second row.
    -   -   This is called _intrinsic grid behavior_.

-   You can go really far with `grid-template-columns` and skipping all other properties.

-   **You can position a grid child with `grid-column` or `grid-row` but have that child display as `block` or `inline`**.

    -   This is quite a powerful technique.

-   To avoid repetition while writing `grid-template-columns: 1fr 1fr 1fr` and so on, use the `repeat` keyword.
    -   `grid-template-columns: repeat(3, 1fr)`;

## Grid span syntax

-   Instead of writing `grid-row: 1/4` you can write `grid-row: 1/span 3`. The same applies for `grid-column`.

    -   I personally find it a bit more confusing than the "number" syntax.

-   The `grid-rows: span 3` would translate to: _"wherever you are in space, span three rows"_.

## Grid area syntax

-   First, you define the _names_ of the "areas" in the grid.
    -   This allows you to construct a very expressive definition for columns and rows.
    -

```css
.container {
    /*We have defined a grid with two columns and three rows*/
    grid-template-areas:
        'feature a'
        'feature b'
        'feature c';
}
```

-   Next, you have to assign the names to elements in the HTML. You do that via the `grid-area` property.

-   To define widths of the columns, **you can combine the `grid-template-columns` with `grid-areas`**.
    -   Otherwise, you will need to repeat yourself in `grid-areas` which is not ideal.

```css
.grid-container {
    grid-template-areas:
        'giant-first giant-first giant-first'
        'semi-second semi-second third'
        'semi-second semi-second fourth'
        'giant-last giant-last giant-last';
}

/*Is the same as*/

.grid-container {
    grid-template-columns: 2fr 1fr;
    /*Notice that I did not have to repeat the second column twice*/
    grid-template-areas:
        'giant-first giant-first'
        'semi-second  third'
        'semi-second  fourth'
        'giant-last  giant-last';
}
```

## Grid and media queries

-   There are two overarching approaches you can take.

    -   The **mobile-first** approach.
    -   The **desktop-first** approach.

-   I tend to lean more towards the **mobile-first approach as when you design for mobile, you have to ensure your site has all the content that it will even need**.
    -   Then you can sprinkle additional bits here and there when widening the viewport.
    -   If we went the other way around (desktop-first), you are at risk of removing critical content from your site.

Sidenote (and as a reminder): the `box-sizing: border-box` makes it so that the `width` includes the `padding`, `margin` and the `border` properties.
If you do not use that rule, the `width` will refer only to padding and can make your life harder when assigning `width` to elements.

## Grid tricks

-   Grid allows you to **overlap items without using `position: absolute`**.
    -   All you have to do is to ensure the `grid-column` and `grid-row` overlap between multiple items.
