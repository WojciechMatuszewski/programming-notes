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
