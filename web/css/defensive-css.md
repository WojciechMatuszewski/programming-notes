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
