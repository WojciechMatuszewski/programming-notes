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

Finished Part 3, 33:49
