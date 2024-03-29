# Defensive CSS

Going through all the tips on [this website](https://defensivecss.dev/).

## Flexbox Wrapping

By default, if you set the `display: flex` on the element, the children of that element **will not wrap** if they exceed the parent width. It is important to note that, **by default, the children of the `flex` container have the `flex-basis` set to `auto`**. This means that **the `width` or `height` will not be taken into the account – the element contents will**.

You will most likely see this happening while having children with text or icons. **To fix the issue, use the `flex-wrap` property**.

```css
.container {
  display: flex;
  width: 100px;
  gap: 1px;
  border: 1px solid black;
  overflow: scroll;
}

.item {
  /* The width here will be "ignored" by default. Use the `flex-basis` property! */
  width: 100px;
  height: 20px;
  background: red;
}
```

## Image Distortion

You have two options for specifying how the image behaves in relation to its container.

1. The `object-fit` property.
2. The `background-size` property.

Note that the `object-fit` is used on a _replaced element_ like `img` or `video`. This means that the element is part of the DOM. The `background-size` is used for... backgrounds. This means that the `img` you use is NOT a part of the DOM tree because it renders as a background for a given element.

There are multiple values for those two properties. Here are the takeaways.

1. The `object-fit: fill` might distort the image. You most likely do not want that. **This is the default value**.

2. The `object-fit: contain` might create a "letterbox" effect.

3. The `object-fit: cover` might zoom-in the image to cover the whole container.

In the end, you will get the best results by tailoring the image size to a given content. Here, anything that enables you to resize a given image on the fly would be helpful.

**Also keep in mind that, by default, the image will have the height and the width of the image size**. To make the above rules work, you must set the `width` and the `height` to be `100%` of the parent container.

## Long Content

Most of the web apps allow users to submit their own content. They might also display the users name and surname somewhere. In such cases, it is imperative to ensure that, no matter the length of the content, the website looks good.

1. Use the `text-overflow: ellipsis` in combination with `white-space: nowrap` and `overflow: hidden`.

   **If you want to produce ellipsis on a given line, use the `-webkit-line-clamp` with `display: -webkit-box`**.

2. Use the `overflow-wrap: break-word` or any other value.

**While you can manually produce ellipsis and hide the text, ask yourself if it is really necessary**. The problem with hiding text is that **it could lead to accessability audit failure**. The hidden text will not be visible for keyboard users/regular users and **might also hide visibility of the text for users who have a high zoom value set in their browsers**.

[This article](https://www.tpgi.com/the-ballad-of-text-overflow/) talks about the problem with `text-overflow` accessibility. The **solution is to embrace the fact that some parts of the webpage might not look symmetrical**, and that is okay!

### Flexbox Considerations

A bit of tangential topic, but still important. Let us consider the following markup and styles.

```html
<style>
  .container {
    border: 1px solid red;
    width: 100px;
    display: flex;
    align-items: center;
  }

  .username {
    text-overflow: ellipsis;
    overflow: hidden;
    white-space: nowrap;
  }
</style>

<div class="container">
  <div class="username_meta">
    <h2 class="username">Very long name name</h2>
  </div>
  <button>Click me</button>
</div>
```

In this case, **the name of the user will overflow out of the flex container**. Remember that **flex items will NOT shrink below its minimum content size**. To address this issue, we use the `min-width: 0` property on the `username_meta` class.

The **`min-width` will override the `flex-basis` and `width`**. Trying to fix this issue with `width` or `flex-basis` will not fly.

## Content Spacing

This blog post touches on the importance of ensuring that elements have enough space between them to "breath". **Side note: you should not embed margins on components but rather have the layout take care of spacing components out**.

This could be achieved in many ways.

1. Using the `margin` property. Bonus for using _logical properties_ like `margin-inline` or `margin-block`.

2. Using the `gap` property on either `grid` or `flex` children.

## `auto-fit` vs `auto-fill` in `minmax`

There is this "famous" CSS Grid syntax snippet you probably seen.

```css
.container {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(200px, 1fr));
}
```

The **`auto-fit` or `auto-fill` is to a CSS grid like `flex-wrap: wrap` is to CSS flexbox**. Of course, there are some differences, but in the simplest terms, using those two keywords will allow you to create "responsive" grids.

The **main difference between `auto-fill` and `auto-fit` is how the items stretch when there is more available space than the `min` value in `minmax`**.

1. The `auto-fill` will NOT expand the grid items if there is an available space. You can think of this as _filling available grid item placeholders_.

2. The `auto-fit` WILL expand the grid items if there is an available space. You can think of this as _fitting the available grid item placeholders_ to the amount of children the container has.

What are the implications? **While using the `auto-fit` property, the last item in the grid, if it is pushed into separate row, might look stretched**. You most likely want to use `auto-fill` and reserve the `auto-fit` for rare occasions.

### With the `min` function

So you have your responsive grid defined like so.

```css
.container {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(300px, 1fr));
}
```

What happens if the viewport gets smaller than 300px? Then, the page will start to scroll – we most likely do not want that. How we can go about making sure our grid is "truly responsive"? **We can leverage the `min` function for the minimum value of the `minmax`**

```css
.container {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(min(100%, 300px), 1fr));
}
```

This way, the minimum value is either `300px` or `100%` of the viewport depending on the viewport size. Pretty neat!

## Background Repeat

If you use the **`background-image`, by default, the image will REPEAT if the element is large enough**. You most likely do not want that. Hopefully you cough this behavior while testing, but **keep in mind that most elements resize based on content/viewport size**. As such you might have missed it in your application.

To solve this issue, consider using the `background-repeat: no-repeat` on the element that has the `background-image` property.

## CSS Variable Fallback

When writing TypeScript/JavaScript code, you most likely used the `??` or the _default assignment_ operator. We use those to apply defaults to variable if the variable does not exist.

```ts
const { data: posts = [] } = useGetPosts();

// or

const { data } = useGetPosts();
const posts = data ?? [];
```

**We can also apply default to CSS Variables**. This ability comes in very handy when we create those variables dynamically.

```css
.container {
  max-width: calc(100% - var(--actions-width, 70px));
}
```

If, for some reason, the script does not load, or there is some kind of error, the `max-width` will have the value of `70px` (this assumes we set the `actions-width` dynamically via JavaScript).

I've used this technique a couple of times with great success. It allowed me to simplify the code in many places.

## Fixed Sizes

I would argue that, the more experience you have creating UIs, the less you rely on fixed sized for elements. Of course, there are times where having a fixed size makes sense, but those are usually few and far between.

Instead of using fixed sizes for element dimensions (`width` and `height` are the biggest offenders, but this could also apply to spacing), consider the following tactics.

1. Use the `min-width` or `min-height` instead of `width` or `height`. Let the elements grow!

2. Use the `min`, `minmax` or `clamp` functions. These allow us to create _fluid layouts_ that look well on all screen sizes.

Side note: I feel like the `clamp` is very much underused in the day-to-day development (I'm also guilty forgetting it exists).

## Minimum Content Size in CSS Grid

Remember the issue we had with `flexbox` children when the content inside the children was too long? Where the child would "escape" its parent since **flex items, by default, will not shrink below its original content size**.

The same thing happens in CSS Grid.

## Minimum size in Flexbox and Grid

We have already touched on this, but I believe it is worth repeating – the **default size of a `flex` child is `auto` meaning it will NOT shrink below its initial size. The same applies for grid (but the width of the element is `min-content`)**.

So, given the following CSS

```css
.grid {
  display: grid;
  grid-template-columns: 1fr 250px;
}
```

It is possible to create an overflow effect on that 1fr column. **Add the `min-width` to the column or use the `min(0, 1fr)` for the column size**.

## Sticky Position Gotcha

Concerns over using the `position:sticky` as it relates to accessability on mobile devices aside, there is one gotcha you might have encountered while working with `position: sticky`.

Sometimes, you add the necessary declarations, but they do not seem to have an effect.

```css
.foo {
  position: sticky;
  top: 0;
}
```

**Always look how much space is the `position:sticky` element taking. If the `position: sticky` is "not working", it might be due to stretching of the element**. This occurs in flexbox and grid layouts!

The solution is to ensure the element has **`align-self: start`**. This will ensure it does not "stretch" to match the "highest" element in the container.

## Scroll chaining

Picture this: you open a modal on a page with a scrollbar. The modal also scrolls. You scroll to the bottom of the modal and then the page starts scrolling! Super frustrating right?

In most cases, you do not want this behavior. **Use the `overscroll-behavior` to change this. Now the "scroll" events will be encapsulated within the element that has this definition until you move your mouse to another scrollable element**.

This is was one of those "I really wish I know this earlier" moments for me when I first encountered this property. It would save me so much time in the past!

## Vertical media queries

There is the `@media(min-width)` query, and **there is also the `@media(min-height)` query**. To be honest, it is the first time I'm writing it. It never occurred to me that one can also query the height of the viewport.

## Accidental hover on mobile

This one is a good one. How many times have you written CSS similar to the following.

```css
.card:hover {
  /* some styles */
}
```

**The problem is that this class will also trigger whenever someone clicks on the `.card` on mobile devices**. This behavior is **dependant on the browser, but in most cases, it is NOT what you want**.

A more robust implementation of the above would look like this

```css
@media (hover: hover) {
  .card:hover {
    /* some styles */
  }
}
```

There are other "device inputs" you can query based on. CSS is vast and there is a lot to learn :)

---

While the TailwindCSS docs do not mention it, the generated CSS **seem to be correct with a special flag on: `hoverOnlyWhenSupported`**. [See this PR for more details](https://github.com/tailwindlabs/tailwindcss/pull/8394).

## Minimum button size

While the action buttons on the website you are working on might look good, it is not always the case for other languages.

Let us consider the word "Done". Given a button with the copy of "Done" and some padding, the button would look ok – not too big, not too small.
Now consider other languages where word "Done" might be a single symbol. If that is the case, the button will look too small!

That is why you **should strongly consider adding a `min-width` to all of your buttons**. This will ensure it they look repesentable no matter the inner content.

```css
.button {
  /* Of course, this is only an example */
  min-width: 100px;
}
```

## Bottom line

The [defensive.css](defensivecss.dev) is a **great resource**. Some of the tips were new, some I already kew about. Either way,
