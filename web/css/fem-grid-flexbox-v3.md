# FEM CSS Grid & Flexbox v3

> Taking notes while watching [this workshop](https://frontendmasters.com/workshops/css-grid-flexbox-v3/).

Finished Part 10 46:36
https://grid-flexbox.css.education/ch11.html

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

-   **Vertical (block-level) margins do NOT collapse** when using grid.

    -   You might need to disable margins when using grip and favor the `gap` property.

-   The `minmax` allows you to create effect similar to `flex-basis` where one column shrinks _first_ and then all other columns follow.
    -   But if you want that kind of behavior, `flexbox` might be a better choice.

## Subgrid

-   The `subgrid` allows you to **make the children of the grid items align in accordance to the "outer" grid**.

    -   This is very handy with _cards_ where you want the headings, the content and the footer of the card to align.

```css
.parent {
    display: grid;
    grid-template-columns: repeat(3, 1fr);
    grid-template-rows: repeat(3, 1fr);
}

.child {
    display: grid;
    grid-template-columns: 1fr;
    grid-template-rows: subgrid;
}
```

-   You might have heard about `display: contents` and how, prior to `subgrid` availability, could allow us to achieve the same thing.
    -   **While `display: contents` COULD, in some cases, help us to achieve the same design, in most cases, it is not a substitute for `subgrid`**.
        -   The `display: contents` removes the "box" from the element and **removes that element from the _accessibility tree_**.
            -   This is now how `subgrid` works!

## Putting it all together

-   When you use `subgrid`, **make sure to set `grid-rows` on the element that has `subgrid` definition`**
    -   Otherwise, all the items will stack on top of each other, no matter what you set for `grid-row` or `grid-column`.

```css
.parent {
    grid-template-columns: repeat(2, 1fr);
    grid-template-rows: repeat(3, auto);
}

.child {
    grid-template-columns: 1fr 1fr 1fr;
    grid-template-rows: subgrid;

    /*Without this, you will not go far*/
    grid-row: span 3;
}
```

## Flexbox

-   For a series of boxes that are different sizes.
    -   Not possible to do with grid.
-   For a series of boxes that are NOT in an evenly sized grid.
-   _Flexbox_ supposed to be _flexible_.

-   The `flex-flow` is a shorthand for `flex-wrap` and `flex-direction`.

-   With flexbox, you can use the `order` property to order given boxes ahead of other boxes.

-   Remember about `flex-basis`!

    -   It works similar to `width` (or `height` depending on the `flex-direction`), but it is _flexible_.
        -   The `width` or `height` are "hardcoded" values. They will not "flex."

-   While working on the navbar, Jen used the `text-align` to center an image.
    -   That is quite interesting as images are not text right?
    -   It works **because the `text-align` aligns all _inline-level_ content inside a given block**.
        -   The `img` tag, by default, is an _inline_ element.

## Responsive images

-   The `picture` tag (combined with `srcset` property on `img`) **allows you to load only a single image best suited for a given resolution/screen**.

    -   You DO NOT want to hide images with CSS.
        -   Even if you hide them from users, **the browser will still load those images**. This will delay the loading of your page.
    -   You DO NOT want to download big images only to scale them to smaller screens.
        -   While this might "work" visually, it will hinder the performance of your application!

-   There are multiple formats for images. **Pick the ones that best suit your needs**.

    -   It is okay to combine multiple formats!

-   Pixel density is a thing.
