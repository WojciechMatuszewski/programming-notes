# Build UIs that don't suck

Notes [based on these videos](https://tailwindcss.com/build-uis-that-dont-suck).

## #0 Using `aria-labelledby` to reduce the amount of text screen reader reads

> [Watch the lesson here](https://www.youtube.com/watch?v=-h9rH539x1k)

In the video, Adam implements this functionality by making the _clickable_ area bigger, but making the _clickable element_ small.

Someone in the comments pointed out that we could use `aria-labelled` here. I second that approach.

**Problem**: Screen reader reads the the whole content of the `a` tag.

```html
<a href="#" aria-labelledby="title">
  <h2 id="title">Some title</h2>
  <p>
    Lorem ipsum, dolor sit amet consectetur adipisicing elit. Iusto a consequatur deleniti, velit iste inventore
    provident voluptates consectetur sapiente nam quis corporis atque, fuga quidem repudiandae! Odit deleniti
    consequuntur commodi.
  </p>
</a>
```

**Solution**: Use `aria-labelledby` to "scope" the `a` accessible name. Be mindful of the difference between `aria-labelledby` and `aria-describedby`.

```html
<a href="#" aria-labelledby="title">
  <h2 id="title">Some title</h2>
  <p>
    Lorem ipsum, dolor sit amet consectetur adipisicing elit. Iusto a consequatur deleniti, velit iste inventore
    provident voluptates consectetur sapiente nam quis corporis atque, fuga quidem repudiandae! Odit deleniti
    consequuntur commodi.
  </p>
</a>
```

## #1 Rounding corners with science and CSS variables

> [Watch the lesson here](https://www.youtube.com/watch?v=X3-4jwm4Z4Y).

**Issue**: Using `border-radius` on nested elements looks off.

```html
<style>
  .outside {
    width: 200px;
    height: 200px;
    border-radius: 12px;
    background: red;
    display: grid;
    place-items: center;
    padding: 6px;
  }

  .inside {
    width: 100%;
    height: 100%;
    background: blue;
    border-radius: 12px; /* Looks off! */
  }
</style>

<div class="outside">
  <div class="inside"></div>
</div>
```

**Solution**: Subtract the parent padding from the `border-radius` value.

```css
.inside {
  width: 100%;
  height: 100%;
  background: blue;
  border-radius: calc(12px - 6px);
}
```

Now the radii looks nice!

## #2 So you think you can center things?

> [Watch the video here](https://www.youtube.com/watch?v=5QTHG99OYS4&feature=youtu.be)

## #3 What are these, buttons for ants?

> [Watch the video here](https://www.youtube.com/watch?v=soFSSkf4oVY)

## #4 Responsive tables that don't suck

> [Watch the video here](https://www.youtube.com/watch?v=v9nHYcKeBw0)

## #5 Dropdowns, icons, and CSS subgrid

> [Watch the video here](https://www.youtube.com/watch?v=7d0qmca5kzc)
