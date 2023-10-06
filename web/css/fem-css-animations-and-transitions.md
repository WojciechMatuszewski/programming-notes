# CSS Animations and Transitions

You can access the course [material here](https://frontendmasters.com/courses/css-animations/).

## Animations fundamentals

- Animations can guide the user. They make the app a delight to use.

- We can also use animations for style & branding.

- **Animation delay happens only once**. If your animation has many cycles, the delay will only be applied at the very start of the animations, not for each "cycle".

- CSS Variables play a vital role in creating animations.

  - They can encapsulate easing, durations and other properties related to animations.

  - Using CSS Variables enables you to create consistent and smooth animations.

  - You can set the value of CSS Variables via JavaScript.

- You can animate a lot of things in CSS. **That does not mean you should animate everything**.

  - Some properties are more performant to animate than others. For example, **animating `transform` or `opacity` is much less expensive than animating `height` or `width`**.

    - Animating `width` or `height` will cause a reflow of the page. The browser is forced to re-calculate positions of the DOM elements on the page. This is expensive!

    - Use the [csstriggers](https://csstriggers.com/) website to check if animating a given property is _safe_.

## Transitions

- Each browser has a default transition easing function.

- The `transition` is a short-hand. I find defining each "sub" property a bit more clear.

- **If you animate multiple properties, you can pick different durations for each property**. The same applies to the easing function.

    ```css
    transition-property: background, opacity;
    <!-- background, opacity -->
    transition-duration: 1s, 3s;
    <!-- background, opacity -->
    transition-timing-function: linear, eas-in-out
    ```

- You **either have to define the `transition` on each element or define it on the parent and use `transition:inherit` on the children**.

  - You are not transitioning variables, you are transitioning properties. This means that the transition happens when you _use_ the variable, not where you change its value.

## Keyframes

- The `@keyframes` give you a lot more control than only using the `transition`s.

- You can define the easing function for each "step" in the keyframe.

- Remember about the **`animation-fill-mode`**. This property denotes what should happen when the animation is about to start or finish. Should the element stay at the "end" position or snap back to its original state.

- I've **noticed that, at least in Brave, when you have the dev tools open, the animation might "jump" at the very end**.

  - This is pretty weird – it does not matter what you animate.

- You can **pause animations via the `animation-play-state`**. This can be useful when you do not want the animations to play when user interacts with a given element.

- **If you are using `animation-delay` you might find the `animation-fill-mode: both` handy**. Instead of jumping to the "start state" after the delay period, the element will preserve the "start state" during the delay period and then stay at the "end frame".

### Keyframes vs Transition

Both can produce animations, but the `transition` is only for animating from the start to the end state. There is no way to define a `x%` step.

As such, consider using `transition` for "simple" animations, and the `keyframes` for more complex needs. Or better yet, always stick to `keyframes` since you can always re-create what you did with `transition` in `keyframes` but not vice-versa.

## Choreography

- You can control the animations and how they interplay with each other via CSS Variables.

  - For example, assign an "index" variable to each element, and then use it for the `animation-delay` property.

- You can define CSS Variables in the HTML, inline.

  ```html
  <div style = "--variable-name: 3">foo</div>
  ```

- You **can not have two animations running at the same time that modify the same property**.

  - Since you mostly should be modifying the `translate`, this means that only one `translate` animation can play at a given time.

    - One **solution to this issue is to wrap the elements you are animating with another element**. This allows you to **put the other animation on the parent**.

## State

- To hold information about different states, one might use the `data-` attributes.

  - We are using the `data-` attribute as those are easily accessed via the `element.dataset.X` where `X` is the attribute name.

- The **main benefit of the `data-` attribute is that a given attribute can only hold a single value**.

  - You might be tempted to use classes for these kinds of interactions, but remember that element can have multiple classes attached to it!

## The FLIP technique (Layout animations)

- Relatively old technique **for animating layout-related properties like `height` and `width`**.

  - Yes, ideally we should not be doing this, but sometimes, for those really nice animations, we do not have any other choice!

- If you are using `translate` and `scale`, **consider animating the pseudo-elements to prevent the children of a given container to stretch**.

  - This is pretty good advice. I always wondered how they did it.

  - Of course, **you do not have to scale some elements**. It all depends on how you want the animation look like. For some, you might opt into transitioning the `transform` property.

## Reactive Animations

- Reactive animations depend on the input of the user. They can change mid-animations depending on what the user does.

- You can send the values from the JavaScript land to CSS land.

  - Remember that the values you get in CSS might be _unitless_. **Use `calc` to transform an _unitless_ value into a value with an unit**.

    ```css
    transform: translateX(
      calc(var(--x) * 1px)
    )
    ```

- David demonstrates the **`lerp`** technique which allows us to create animations based on the movement of some element, like a pointer.

  - Pretty neat.

- Not everyone wants animations. You can detect user settings via the `@media` query.

  ```css
  @media (prefers-reduced-motion: reduce) {
    <!-- styles -->
  }
  ```

- What blew my mind is the way David animated the text reveal from the right.

  - He used the `clip-path` with a polygon and then started revealing the text under that polygon.

  - This created an effect as if the animation reveals a letter at a time. You mind find [this page useful](https://bennettfeely.com/clippy/).

    ```css
    @keyframes reveal-text {
      from {
        clip-path: polygon(0% 0%, 0% 0%, 0% 100%, 0% 100%);
      }

      to {
        clip-path: polygon(0% 0%, 100% 0%, 100% 100%, 0% 100%);
      }
    }
    ```
