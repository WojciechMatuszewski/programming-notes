# FrontendMasters Website Accessibility v3

Taking notes based [on this course](https://frontendmasters.com/workshops/accessibility-v3/)

## Screen Readers

- Accessibility is just as important as any other aspect of software development.

- Just like you have _engineering debt_, you could have _accessibility debt_.

- Sadly, not everyone shares the sense of urgency when it comes to accessibility.

  - Accessibility **is the right thing to do** from the civil rights perspective.

  - You might face legal issues if your website is not accessible.

  - By having a well crafted and accessible website, you can be proud of what you accomplished.

  - For business cases, **the more accessible the website, the bigger reach it has**. This could translate to more customers and more revenue.

- Most people with any kind of disability that could prevent them from browsing the web use _screen readers_.

  - This is a software that will read out the contents of the webpage and aid with navigation.

- Screen readers leverage mostly the **_accessibility tree_** rather than the _DOM tree_.

  - Of course, there are exceptions.

  - _Accessibility tree_ is very handy for checking what is the role of a given element. I use it while writing tests all the time!

- While working with images, **include the `alt` attribute with meaningful content**.

  - If you do not, some screen readers might read the filename out loud.

    - Image the screen reader reading the `base64` encoded image. Not fun.

  - **Some images do not warrant the `alt` text. These are mostly purely decorative**.

    - In such case, use the `alt=""` syntax.

  - You might find this [decision tree helpful](https://www.w3.org/WAI/tutorials/images/decision-tree/).

## Debugging

- There are few steps you should consider

  1. Try using the `Tab` key for navigation. Can you navigate / use the website only via keyboard?

  2. Use `Axe` or similar extension to discover accessibility issues.

  3. Zoom your page. Some people use the web with 200% zoom or more.

  4. Run screen reader and try to use the website.

- **[The `Web Developer toolbar`](https://chrispederick.com/work/web-developer/) seem to be very useful extension**.

  - Some of the functions this tool has are amazing!

- Overall, I **would highly recommend [this website for reference](https://web-accessibility-v3.vercel.app/topics/accessibility-debugging/linters-and-devtools)**.

- The best accessibility testing is the manual testing. **You can have a great lighthouse score but still have inaccessible website**.

- If someone insist on removing focus, they might not be aware of the `:focus-visible` selector.

  - Consider educating that person. Prior to `:focus-visible` clicking elements would trigger focus. With `:focus-visible` this is not the case.

  - Removing the _focus outline_ really sucks. It makes the website completely inaccessible for keyboard navigation.

## Accessible HTML

- **No matter what framework you use, the basics always apply**.

  - You ought to leverage semantic HTML tags.

- Use built-in functionality first.

- Reach for _landmark_ elements, like `main`, `nav`, `section` and others.

- Make sure your headings elements are in the correct order.

  - You should not have multiple `h1`s on the page.

- Use **`aria-label` or `aria-labelledby` to give unique labels to `nav` and `section` elements**.

- Lists elements are great.

- Gotcha: **Safari will treat lists with `list-style-type: none` as `divs`**.

  - To fix, consider adding `role="list"` to every list element.

- Setting a language on `html` might be a huge win.

  ```html
  <html lang="eng"></html>
  ```

  **You can override language on the element level as well**.

  ```html
  <html lang="eng">
    <body>
      <p lang="fr"></p>
    </body>
  </html>
  ```

### Debugging HTML exercise

- You cannot have both `role="banner"` and a `header` tag rendered at the same time.

  - There can only be one "banner" on the page.

- The `aside` tag has to be "top-level" meaning it is not wrapped within any other landmark element like `main`.

- **It is perfectly fine to have `hx` tags that are "sr-only"**.

- The `inert` does not work well in React. It seems like passing `inert={true}` on a node will not do anything. I assume it has to do with JSX transpiler rejecting the attribute since it's not a "known" attribute.

  ```jsx
  <div inert={true}>foo</div> // Does not work!
  ```

  What you might want to do instead is to leverage _callback refs_ or using a different value than `true` for the attribute.

  ```jsx
  <div>
    <div
      ref={(node) => {
        node.setAttribute("inert", "true"); // Works
      }}
    >
      foo
    </div>
    <div inert="">Another node</div> // Also works
  </div>
  ```

  This will ensure that the attribute is applied correctly.

- Images should have the right `alt` text. There is no excuse to not having one, especially in the era of AI tools.

## ARIA

- Its best to err on the side of NOT using `aria-` attributes.

  - If you do not know what you are doing, you can make things worse.

- ARIA encompasses _roles_, _states_ or _attributes_.

  - The **_role_ tell us what a given element does**.

    - Is it a button? Then there is `aria-role="button"` and so on.

    - **Semantic HTML tags often have implicit roles**. This is good, you do not want to change them!

      - For example, the `button` tag has an implicit _role_ of "button".

- You can use **ARIA in CSS selectors**.

  ```css
  [role="checkbox"][aria-disabled="true"] {
  }
  ```

  Side note: let us not forget the `data-` attributes! They also could come in very handy.

  ```css
  [data-open="true"] {
  }
  ```

- To provide more context to a given element, you have the `aria-label`, `aria-labelledby` and `aria-describedby` to choose from.

  - The `aria-label` works with a text value.

  - Both the `aria-labelledby` and `aria-describedby` work with element ID.

    - **Use `aria-labelledby` to _name_ the element, use `aria-describedby` to _describe_ the element**.

      - The _name_ would be much shorter than _description_.

    - **`aria-labelledby` will OVERRIDE the existing label**.

### Live Regions

- They are used to **make the screen reader read out things that happen on elements that the user is not currently focused on**.

  - A very good example are toasts messages when there is an auto-save feature.

    - In that case, the users focus is most likely on some kind of input field. It would be neat to make the user aware that their changes were saved then they finished typing. This is where the concept of _live regions_ comes in!

- To create a _region_, use the `region` attribute or the `role` attribute.

  ```html
  <div aria-live="polite">foo</div>

  <div aria-live="assertive">bar</div>

  <div role="status">Loading...</div>
  ```

  There are various additional attributes you could use to supplement the `role` or `aria-live`. The most critical thing to understand is that **when using `assertive`, any changes within the element will STOP what the screen reader is currently saying and announce the change to that element**. The `polite` changes will be announced once screen reader is "idle".

  One additional technique is to **wrap the _regions_ with `sr-only` class**. Again, these are mostly for screen readers.

  ```tsx
    <div>
    <NotificationIcon>
      <div aria-live="polite" className = {"sr-only"}>
      You have 5 notifications
      <div>
    </div>
  ```

- Please do not overuse the _live regions_. Imagine some element constantly re-rendering. It might be the case that the screen reader will announce those re-renders. This will be VERY frustrating for the end-users.

## Focus Management

- Use the keyboard to navigate the application you are working on. Is it even possible?

  - If you use semantic HTML, you should be able to do this. **You do not have to use any `aria-` attributes to make your website accessible to keyboard-only users!**

- Any website should consider adding _skip links_.

  - Here, it is also crucial to understand how `tabIndex` works.

    1. The positive values will influence the _tab order_ of the page.

    2. The zero value will make element "tabbable".

    3. The negative value will **make the element "invisible" to the tab ordering, but still be able to receive focus**.

    In the face of _skip links_, you mostly want the target of the link to have a negative `tabIndex`. It will be `main` or `footer` either way.

- When managing focus, **use the `activeElement` property**. It will tell you what is the currently focused element.

  - To save yourself a lot of hassle, consider using _live expression_ with it. Otherwise you will have to re-query it every time it changes.

## Visual Considerations

- Color and visual contrast is important. Many people might have issues reading the content on your page if the ratio does not meet the guidelines.

- **Test your website with a bigger zoom setting (it should work well up to 200%)**.

  - Ask anyone older to hand you their phone. You will notice that they have enlarged UI.

    - You need to build pages accessible for anyone!

- There are two very common APIs that pertain to user preference.

  - The `prefers-reduced-motion`. This one tells you whether the user is "motion sensitive". If that is the case, you might want to disable / reduce animations.

  - The `prefers-color-scheme`. This one tells you whether the user wants the _light_ or _dark_ theme.

## Wrapping up

- **Accessibility is important and vital to product success**.

  - There is a correlation between accessibility and how "good" the webpage feels. You probably want your web page to feel pretty good to users right?

- Marcy left us with **a great list of awesome tools to add to our workflow**.

  1. [The accessibility insights](https://accessibilityinsights.io/docs/web/overview/).

  2. [The Web Developer Toolbar](https://accessibilityinsights.io/docs/web/overview/). This one is truly amazing.

  3. I've also learned that service like [stark exist](https://www.getstark.co/brave/). Could come in handy.

- Use `aria-labelledby` or `aria-describedby`.

  - I found these to be very useful for modals. When you use those attributes, you can query the modal with `getByRole("modal", "my title")` while testing stuff.

- **Use semantic HTML**.

  - Landmarks are there – the `nav`, `header`, `main`, `aside` and `footer`.

  - Some landmarks are "top-level" landmarks, meaning they should not be nested within other landmarks.

- **Ensure you have correct heading hierarchy**.

  - Use the tools I've listed above to check.

- **Consider adding `role="status` for containers that display any loading indicators**.

  - This enables screen readers to announce that the content is loading.

  - Remember that, by default, the `role="status"` has "polite" assertiveness. This is most likely what you want in this case.

- Be aware that _aria regions_ are a thing. Use them!

- Use the `.sr-only` (or similar) class to provide meaningful description for items that should not contain any visual content.

  - This is very common practice for so-called _Icon Buttons_ – a button that only renders some kind of icon.

    - In this case, also ensure that the button has enough padding around it, so that users visiting on mobile can click it!

- For navigation, use only the keyboard for a day or two. How did it go?

  - If you are having problems using your product, imagine users who have no choice but to use the keyboard.

Overall, great workshop. I wish there was a bit more practice.
