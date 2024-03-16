# FEM Testing Fundamentals

[You can find the course material here](https://frontendmasters.com/workshops/testing/)

- There are various reasons we write tests for:
    - Improved Code Quality.
    - Increased Productivity.
    - Team Collaboration.
    - Confidence.

- Misko also talks about **writing tests because of laziness**.
    - We are lazy. We want to do something and "be done with it." The "done with it" part is enabled by tests.

- **You want the tests to be easy to write, even easier than running your application,** so that they are the _simplest possible thing to do at a given time_.

- **Writing tests in legacy codebase is quite tricky**.
    - There is a lot of "baggage" in such codebases.

- Writing tests makes you write _different_ code – code that is less coupled.
    - This might feel "bad" to people who are not used to writing tests.

- Tests are literally a superpower.
    - They allow you to refactor.
    - They give you a lot of confidence.
    - They **will** make your code easier to read and reason about.

- You do not want to be dogmatic.
    - Ideally, we would be writing tests BEFORE writing the code.
    - In reality, that might not fly with everyone. If you want to change the culture, try to do it gradually.

- **Testing is all about managing dependencies**.
    - Sometimes you need to mock a dependency.
    - Sometimes mocking the dependency is not the right choice.
    - It all **depends** on the "level" of the test (unit, integration, or e2e) and how critical a given dependency is.

- Testing helps you answer the **WHY** question – arguably, the most important question in software development.
  - WHY are we doing what we are doing?
  - WHY this function works the way it works?
  - WHY do we have these constraints?

- Use the `it.todo` or `test.todo` to **describe the requirements of the function/module you are testing**.
  - Super helpful for keeping the context and also gives you an overview of the work remaining.
    - Of course, you will most likely forget to include some requirements. That is fine. We are all only human.

- TIL that the callback function you provide to `test` or `it` takes `expect` as an argument.

  ```js 
    it("works", ({expect}) => {})
  ```
  
  Misko mentions using this `expect` rather than the "global" one works much better when running multiple tests in parallel.

- When testing a module which uses `fetch`, Misko decided to use _dependency injection_ rather than a library like MSW.
  - This is quite an interesting choice. Personally, I like the MSW approach better.
  - Having said that, I'm very happy to se the _dependency injection_ used here. I think that people still forget how powerful this technique is.


- When dealing with _time_, Misko opted to pass the "delay" function as a parameter.
  - I have to admit, I did not expect this choice. I would expect the "timeoutMs" to be passed as a parameter instead.
  - In my humble opinion, passing the `delay` function as a parameter exposes implementation details.

    ```ts
    fetchMock = vi.fn<Parameters<Fetch>, ReturnType<Fetch>>(mockPromiseFactory);
    delayMock = vi.fn<[number], Promise<void>>(mockPromiseFactory);
    
    api = new GithubApi("TOKEN", fetchMock, delayMock);
    ```

- I'm also not a fan of how Misko abstracted the "setup" functions.
  - While I agree with abstracting the mocks, I'm unsure if abstracting the `api` to the global setup is a good choice.
    - Reason against: each test should declare all the necessary dependencies it needs. Otherwise, they are not completely isolated and might override each other.

    ```ts
    let fetchMock: Mock<Parameters<Fetch>, ReturnType<Fetch>>;
    let delayMock: Mock<[number], Promise<void>>;
    let api: GithubApi;

    beforeEach(() => {
        fetchMock = vi.fn<Parameters<Fetch>, ReturnType<Fetch>>(mockPromiseFactory);
        delayMock = vi.fn<[number], Promise<void>>(mockPromiseFactory);
        // What is the tests run in parallel?
        api = new GithubApi("TOKEN", fetchMock, delayMock);
    });
    ```
    
- I've noticed that Misko does not "wait" for `assert` to be fulfilled while dealing with promises.

    ```ts
    const pendingFetch1 = api.fetchSomething();
    expect(mockFetch).toHaveBeenCalledWith("foo");

    // Instead of
    const pendingFetch2 = api.fetchSomething();

    await vi.waitFor( () => {
        expect(mockFetch).toHaveBeenCalledWith("foo")
    })
    ```
  
    This is quite problematic and source of issues in the workshop. The `expect` could run BEFORE we started to fetch something. As such the tests could end up flaky.

- Misko touches on the importance of **breaking up the _business logic_ and "construction code"**.

- Snapshots are difficult to read. **Consider using the `toMatchInlineSnapshot` if you can**.
  - By putting the snapshot inside the "test code" you will be less tempted to have huge snapshots you have to scroll through.
  - Snapshots **could be useful for legacy systems to ensure that nothing changed**, but if you fear change to a system, you are in a deep trouble.

- Misko opted to use _Storybook_ for component testing via screenshots.
  - Interesting choice. I would expect usage of Playwright component testing here instead.
  - _Storybook_ integrates with _Chromatic_ for screenshot testing.
  - TIL **that _Storybook_ also has a separate test-runner**.
    - From what I'm reading, it is rather simplistic and does not allow you to run "component-like" tests, but still might be useful in some situations.

- When using Playwright, Misko started asserting on classes. **I strongly believe that this approach is not that great**.
  - Classes can change, but the visual output might be the same. If the class changes, the test will break.
    - This is a reverse of "testing as the user sees it." The user is NOT concerned with the names of CSS classes, but rather what is in displayed in the screen.


- Despite using the names of the classes for the selectors, Misko did implement the "POM" model. 
  - I'm a big fan of the "POM" model, but one has to be careful not to overengineer the "POM."
    - **The more indirection layers, the harder it is to reason about the application**.

    ```ts
    test("it works", async ({page}) => {
        const clusterPage = new ClusterPage()
        
        await clusterPage.goto()
        
        // More stuff
    })

    class ClusterPage {
        constructor(page: Page) {
        }
        
        async goto() {
            return await this.page.goto("...")
        }
        
        // Some functions
    }
    ```

- **I/O operations are making the tests slower**.
  - It is quite hard to make the tests CPU-bound.

- Misko touches on the _testing pyramid_ and **how it changes depending on the type of the codebase you are working on**.
  - The **greenfield** projects usually follow the "typical" _testing pyramid_ where you have lots of unit tests with fewer integration tests and even fewer e2e tests.
  - The **legacy** codebases are inverse of that – lots of e2e tests and fewer unit tests.
    - This is because it is, usually, much harder to write unit tests in the legacy codebases due to code rot and other factors. 

- **While in JavaScript you CAN control the dependencies via global state, doing so will prevent you from running the tests in parallel**.
  - A good example is the `fetch` API. You might be tempted to override it in the tests, but that will make the test to be dependent on the global state.

- I love that Misko puts a lot of emphasis on _what makes the code hard to test_ as well. Here is a non-comprehensive list (I agree with all the points).
  - Mixing "new" with logic.
  - Looking for things.
  - Work in constructor. Imagine creating dependencies _inside_ the constructor rather than to have them as parameter. How would you test the class if you do not control the dependencies?
  - Global state.
  - Singletons -> those are pain in the ass.
  - Static methods.
  - Deep inheritance.
  - Too many conditions.
  - Functions that "look for other things." For example, a function takes in `home` as parameter, but only uses `home.garden`.
    - One should make the dependencies as specific as possible. This reduces the amount of mocking you have to do in tests.

## Wrapping up

While I disagree with some of the test-code Misko wrote, I'm very grateful for all the knowledge he passed pertaining to writing more testable code.
The most single important piece of advice seems to be the following: **make sure you control the dependencies** – otherwise, you will not be able to test the application.
