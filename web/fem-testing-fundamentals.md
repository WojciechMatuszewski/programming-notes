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

Finished part 2, 4:30