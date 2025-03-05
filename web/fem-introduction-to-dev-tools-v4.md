# Frontend Masters Introduction to DevTools V4

> https://masteringdevtools.com/

- The _Computed_ tab in the elements panel is quite useful.

  - You can search for a given CSS property and it will show you _resolved_ value. You do not have to filter through applied classes and such to get to the "real" value.

- **Consider debugging without any active extensions**.

  - In some cases, extensions mess up with the DOM or performance of the website.

- The `$0` in the console references the latest queried element in the console.

  - Pretty neat! Especially when using the "element selector" tool.

- The "scroll into view" option for a given element is quite handy if you work on large lists.

- **When debugging network-related stuff, be mindful of cache**.

  - It is very likely you have cache disabled via the "Disable cache" checkbox in the _Network_ panel.

- **Always disable all extensions when measuring performance AND test on "production" version of the application**.

  - This ensures that the results you get are reproducible.

  - There might be a huge difference in performance between "dev" and "production" mode of application.

  - **Lighthouse is pretty good for catching basic performance and accessibility issues**.

    - Nothing beats measuring against organic traffic, but Lighthouse could be a good start.

- **TIL that hitting "ESCAPE" opens another pane with more tools**.

  - The `CMD+O` is also very handy.

- You might want to configure your build tool to automatically remove `debugger` statements from the production bundle.

  - Nothing bad will happen if you commit a `debugger` statement, but it is worth making sure those do not "escape" our local development environment.

- **You can set _conditional breakpoints_**. This is super handy if there are multiple code paths in a function, but you are only concerned with that one particular one.

- How many times have you ventured into React.js library? Not helpful right? **You can ignore certain paths in dev tools so you never end up in "node_modules" land**.

  - See [this website](https://developer.chrome.com/docs/devtools/settings/ignore-list) for more information.

- "JavaScript bytes" are much more costly than "Image bytes".

  - This is because it takes a lot of time to parse JavaScript code.

- When looking at the performance of a given function in dev tools, **look at the "self time" of the function call**.

  - Some functions might have a very high "total time", but the "self time" might be very short. This indicates that this function _calls_ the problematic function.

## Summary

- DevTools are very useful and have more tools than you most likely are aware of.

- **Conditional debugger statements are a thing**. Use them.

- **"Popping out" devtools to a separate tab can be helpful**. When you have devtools open "inline", the performance of you site might suffer.

- **The "experiments" panel in DevTools might contain a option that is helpful for your situation**.

- Filtering paths in "Sources" panel is a great way to ensure you won't end up in library code when stepping through your code.
