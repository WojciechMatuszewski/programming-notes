# System Design for Frontend Engineers

You can find the course [material here](https://www.greatfrontend.com/system-design)

## Frontend vs Backend System Design

While the overall, big picture, is similar, there are some differences between the two.

- You are not expected to give "back of the envelope" calculations in the Frontend interview.

- Instead of focusing on the distributed services, one should focus on the _components_ and how they interact with each other in the UI.

- As opposed to thinking about scalability or reliability, one should focus on performance, user experience, accessibility and internalization.

## During the interview

- Consider using the **RADIO** framework. Overall, always start with the requirements first, and then build your way up to something on the screen.

  - Requirements exploration.

  - Architecture and High-level Design.

  - Data Model.

  - Interface Definition (API).

  - Optimizations and Deep Dive.

- Treat the interviewer as a Product Manager. Ask various questions to see what part of the system they are mostly interested in.

  - They might have different backgrounds. One might be more interested in accessibility. The other might be more interested in the API design.

- **Clarify which parts of the product to focus on**. This is essential. The questions will be vague, for example "design facebook". If you do not clarify the scope, you might be in trouble later on.

  - Keep in mind that there are also features which could be deemed as "nice to have".

- **Know your users**. Any business should have a list of "personas" defined. These are your customers. Customers could be very different and have very different needs or expectations. For example, if you are creating an application for creative people, this application might be rich with animations and fluid motions. For banking application that might not be necessary, as people usually care about their banking balance and not that it slides from the right (an example).

## System Design pieces

- There is the **server** which you can mostly treat as a _black box_. Of course, it is vital to clarify what is the format in which the data ends up on the frontend (_JSON_ or maybe _protocol buffers_ or others?), but you should not spent a lot of time talking about it.

- As for the UI part, you might want to compose your architecture using the **View**, **Controller** pieces. You most likely already created applications with this structure, but have not attached any specific names to the files you are interacting with.

  - The **View** is for displaying things to the user.

  - The **Controller** handles the network communication and data flow.

## Things to consider

Based on [this page](https://www.greatfrontend.com/front-end-interview-guidebook/user-interface-questions-cheatsheet).

### Component Organization

- Lean heavily on composition. Modern frameworks use components which can compose really well. If you do, you will arrive at a clean, well defined API and component structure.

  - **Keep the surface area as small as possible**. This principle applies to all programming languages and is not frontend-specific. Like in Go, the interfaces should be small, the same rules applies to React.

### State

- Any application will have state. It is inevitable. After all, we are building an _interactive_ applications.

  - Keep the state small.

    - Identify the _essential_ and the _derived_ state. Deriving the state is usually better as it makes the application more malleable.

  - Separate the state from the "rendering code".

### Semantic HTML

- In todays world, where a lot of people are defaulting to `div` tags, it is very important to understand the principles of semantic HTML.

- You do not have to be an expert. Knowing that `section` should contain a `h` tag or that each form element should have a label will get you really far.

### User Experience

- I would argue that most of the websites should be mobile-friendly. Or at least look decent on mobile.

  - Here, you can talk about _media-queries_ or _container-queries_ and different dynamic units.

  - When working with UIs, it is **vital to ensure interactive elements are big enough for mobile**. Imagine a tiny button user is expected to click on their phone. That is bad user experience.

- The errors should be clearly visible and well worded. The network can fail, the WIFI can be spotty. It is vital to keep the UI integrity together when the app works in such conditions.

- Media elements like images are huge part of the web now. Talk about modern image formats like `webp` or `avif`.

- Network can also be slow. For this, consider **using optimistic updates**. Most of the data-fetching libraries allow you to easily implement this.

### Considerations for the Network

- Network can be unpredictable, the application should handle all states the request could be in (pending, error, success and maybe some other depending on the application logic).

- **Making multiple network requests might lead to race conditions**. It is important to ensure the application always displays the latest state.

- **Consider batching or throttling** if your application is using a lot of network requests. Of course, we are quite limited here by the design of the Server API.

- You might also need to enforce a certain timeout on the requests. Handling the errors related to timeouts is quite important.

### Accessibility

- Accessibility is vital as it enables users with all sorts of disabilities to use your application. **It also makes the site a joy to use for existing users**.

- Remember about various `aria-` roles. The most common are `aria-label`, `aria-role` or, when dealing with forms, `aria-describedby` and `aria-errormessage`.

- It is worth keeping **color contrast** and size of the elements in mind.

### Edge Cases

- Mind the long strings. I'm guilty of forgetting about these a lot. Consider applying the techniques from the [defensive css](https://defensivecss.dev/). Especially [the one talking about long content](https://ishadeed.com/article/css-short-long-content/)

- Since you will most likely be fetching data from the API, create empty states. Not handling those makes for an awkward user experience.

- Performance is also vital. The more nodes are on the page, the worst is the performance. Here you might want to reach for virtualization techniques.

- Nowadays, **caching** is an integral part of any frontend application. The problem is that **caching is hard**. Thread carefully here.

### Security

- Most frameworks sanitize all inputs by default, so `XSS` attacks are not really a thing. Having said that, most frameworks also allow you to put "raw" strings into the dom, for example the `dangerouslySetInnerHTML` function in React.

- For additional security, the page should [employ CSP headers](https://developer.mozilla.org/en-US/docs/Web/HTTP/CSP).

  - CSP headers allow you to specify where the website content is coming from, or block any specific domains from executing code.

### Internationalization

- This might or might not be a requirement. I think it mostly depends on how big the company is.

  - The bigger the company, the wider the market is is trying to capture, the bigger need for internationalization.

- **As long as you use [logical properties](https://developer.mozilla.org/en-US/docs/Web/CSS/CSS_logical_properties_and_values)** and **do not hard-code strings in your application** you should be good to go.
