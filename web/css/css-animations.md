# CSS Animations

## Animating the height `auto` property (or any other size-related property)

Since the beginning of CSS (?), we have been unable to animate the height of the element from "auto" to 0 (or
vice-versa). That lead us to use various "hacks" to create toggles. With time, these "hacks" became the go-to solutions
for creating a "hide/show" animation.

In an ideal world, we could animate the height using CSS, but that is not possible. If you want to animate the height,
you shall use JS, and the best way to do this is to use the so-called FLIP technique.

### Using FLIP

With the FLIP technique, you first measure the element, apply the animations, and reverse them – all in a single frame.
This allows you to create an effect of smooth animation. Do not get fooled by the "simplicity" of the steps listed
here – to **correctly animate the height using FLIP, you will most likely need to apply reverse transforms**. You are
better off using a library that will perform the math for you.

To learn more about the FLIP technique, checkout [_Method4_ in this blog post](https://carlanderson.xyz/how-to-animate-on-height-auto)

### Using the `max-height` property

Instead of animating the `height`, we could animate the `max-height` property. This avoids the need to apply the scale
transformations – the content will shrink (remember about the `overflow` property here!) as expected.

One note about the `max-height`: **you most likely want to measure the element you animate before you animate it**. The
reason is that if you apply a very high `max-height` and then animate it to 0, the animation will look weird – the
timing function applies to the whole `max-height` range!

Apart from the issues with timing functions, the `max-height` can sometimes mess with your layout. Be mindful of what
you put the `max-height` on – we would not want the content to create unnecessary scroll bars!

### Using the CSS Grid layout

This one is entirely new for me. Instead of animating height-related properties, we could animate the grid tracks.
Combine it with the `overflow: hidden` property, and you have a beautiful "collapse" animation.

```js
document.getElementById("app").innerHTML = `
<button type = "button" id = "toggle">Toggle</button>
<div class = "box">
  <div class = "box-inner">
    <p>content</p>
    <p>content</p>
    <p>content</p>
    <p>content</p>
  </div>
</div>
`;

const button = document.getElementById("toggle");
const box = document.querySelector(".box");

button.addEventListener("click", function onButtonClick() {
  const isHidden = box.classList.contains("hidden");
  if (isHidden) {
    box.classList.remove("hidden");
    return;
  }

  box.classList.add("hidden");
});
```

```css
.box {
  display: grid;
  grid-template-rows: 1fr;
  background: red;
  transition: grid-template-rows 0.5s ease-in-out;
}

.box-inner {
  overflow: hidden;
}

.hidden {
  grid-template-rows: 0fr;
}
```

This technique is used in [this free course](https://www.epicweb.dev/tutorials/fluid-hover-cards-with-tailwind-css).

Looking at the performance snapshot from Chrome, animating height this way will cause **layout and style recalculation** but **it seems to be much cheaper than reflow**.

### Using the `calc-size`

> Based on [this article](https://developer.chrome.com/docs/css-ui/animate-to-height-auto/)

**The API of the `calc-size` is mind-blowing**.

Relatively new addition to the CSS, it allows you to animate from `0` to `max-height` or any other other "auto" property. **Please note that this is not the same as animating from `display: none` to `display: block`** – for that we have **`@starting-style`**.

```css
.animate {
  width: 30px;
  overflow: clip;
  display: block;
  white-space: nowrap;
  transition: width 0.3s ease;

  &:hover {
    width: calc-size(max-content, size);
  }
}
```

When you hover over the element, it will animate from `30px` to `max-content`. No `grid` or FLIP needed.

**When the `calc-size` really shines, is the ability to perform operations on the `size` keyword**. Check this out.

```css
.animate {
  width: 30px;
  overflow: clip;
  display: block;
  white-space: nowrap;
  transition: width 0.3s ease;

  &:hover {
    width: calc-size(max-content, size / 2);
  }
}
```

When you hover over the element, it will animate from `30px` to **half of the `max-content`**. How amazing is that?

**Combine this with [`round` CSS function](https://developer.mozilla.org/en-US/docs/Web/CSS/round) and the sky is the limit**.

```css
p {
  width: calc-size(max-content, round(up, size, 50px));
}
```

The `p` all have the `width` of `max-content` but rounded to the nearest multiply of `50`. **Again, how amazing is that?**

### Using `interpolate-size`

> [Read more about this property here](https://developer.mozilla.org/en-US/docs/Web/CSS/interpolate-size)

This property allows you **to enable native animations between length/percentage values and _intrinsic sizes_ such as `auto` or `fit-content`**. At the time of writing, this property is not YET widely supported, but it is only a matter of time when it will be.

You _might_ consider adding `interpolate-size: allow-keywords;` to your CSS reset, but be mindful of users who prefer the "reduced motion" experience.

```html
<!doctype html>
<html lang="en">
  <head>
    <meta charset="UTF-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
    <title>Document</title>
    <style>
      #box {
        /* From "auto" size */
        interpolate-size: allow-keywords;
        background: red;

        transition:
          block-size 1s,
          inline-size 1s;

        &:hover {
          /* To a well-defined size */
          block-size: 80px;
          inline-size: 80px;
        }
      }
    </style>
  </head>
  <body>
    <div id="box">box</div>
  </body>
</html>
```

## Animating display via `transition-behavior`

> Read more on [MDN here](https://developer.mozilla.org/en-US/docs/Web/CSS/transition-behavior).

Since the `display` is a "discrete" _animation type_, by default, it will not work with any transitions. This is how it worked for pretty much forever, and required you to use JavaScript to create "show/appear" animations when toggling display.

**That is no longer the case if you use the `transition-behavior: allow-discrete`**.

```css
dialog {
  display: none;
  opacity: 0;

  transition-duration: 1s;
  transition-property: opacity, display;

  /* The key to making this whole thing work \/ */
  transition-behavior: allow-discrete;
}

dialog[open] {
  display: block;
  opacity: 1;

  @starting-style {
    opacity: 0;
  }
}
```

The `@starting-style`, in this example, as I understand it, is the value you want to animate from when the `display` changes.

**One gotcha to keep in mind is that you can't nest within the _pseudo-classes_**. So you if wish to animate the `backdrop`, you need to move the `@starting-style` outside of the selector.

```css
dialog {
  display: none;
  opacity: 0;

  transition-duration: 1s;
  transition-property: opacity, display;

  /* The key to making this whole thing work \/ */
  transition-behavior: allow-discrete;
}

dialog[open] {
  display: block;
  opacity: 1;

  @starting-style {
    opacity: 0;
  }
}

dialog::backdrop {
  display: none;
  background: black;
  opacity: 0

  transition-duration: 1s;
  transition-property: opacity, display;

  /* The key to making this whole thing work \/ */
  transition-behavior: allow-discrete;
}

dialog[open]::backdrop {
  display: block;
  opacity: 0.5;
}

@starting-style {
  dialog[open]::backdrop {
    opacity: 0;
  }
}
```

**Notice that the definition inside the `@starting-style` can be different for different "directions"**.

### Note about `display` property

It is worth noting _how_ the browser "animates" the `display` property.

- **When animating from `display: none` to `display: block` (or any other value)**, browser will flip the element to `block` at the very start of the animation, so it is visible. **This is why you usually want to animate `opacity` alongside the `display` property**. Otherwise the element will "pop in" out of nowhere!

- **When animating from `display: block` to `none`**, browser will **flip the element to `none` at the END of the animation**. Again, this is to ensure the element is visible thought the animation.

## The Popover API

The [popover API](https://developer.mozilla.org/en-US/docs/Web/API/Popover_API) introduced a wave of very useful features.

First is the fact that it allows you to natively created popovers on a webpage **without using JS**.

The following snippet will take you very far.

```html
<button popovertarget="mypopover">click me</button>
<div id="mypopover" popover>popover content</div>
```

At the time of writing, the default is to place the popover at the middle of the viewport.

Then, you can control the look and feel via different states of the `popover` element.

```css
/* Starting styles */
[popover] {
  opacity: 0;
  transition: all 0.7s allow-discrete;
}

/* Animate to */
[popover]:popover-open {
  opacity: 1;
}

/* When toggling the button */
@starting-style {
  [popover]:popover-open {
    opacity: 0;
  }
}
```

Since, **by default the popover is NOT anchored to the target**, you will most likely want to combine this with the [CSS Anchor Positioning](https://www.oidaisdes.org/popover-api-accessibility.en/).

```html
<button popovertarget="mypopover" id="target">click me</button>
<div id="mypopover" anchor="target" popover>popover content</div>
```

Notice that I did not have to apply any additional styles to make the functionality work. Of course, If I need to customize the positioning, I can do that!

- To learn more, consult the [MDN article](https://developer.mozilla.org/en-US/docs/Web/API/Popover_API) about the Popover API.

- To learn more about positioning the elements anchored to another element, read [this blog post](https://developer.chrome.com/docs/css-ui/anchor-positioning-api#use_anchor_with_top-layer_elements_like_popover_and_dialog).
