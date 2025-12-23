# Tactical Agentic Coding

## Hello Agentic Coding

- The "core four" things that _really_ allow you to automate coding are:

  1. The tools.
  2. The context.
  3. The model.
  4. The prompt.

  If you have a software that understands those four (like Claude Code), you can go really far with coding.

- You can run Claude Code as a terminal application, but you can also run it programmatically.

  - This makes it _really_ powerful.

## The 12 Leverage Points

> It's not about what you can do. It's about what you teach your agents to do.

- In the course, we used a very simple prompt to force Claude Code to list all tools it has access to. It had never occurred to me to do that...

  - Such a great information to have!

- **Your applications must have a way to communicate issues to the Agent**.

  - Are you logging the errors of your applications to stdout? If not, the Agent has no way of knowing what the issue might be (unless you copy & paste the issue manually, but that's quite painful to do).

  - In the course, they run the application via script but _inside_ Claude Code. This enables Claude Code to read the application logs.

    - As an alternative, we could write application logs to a file. I believe [`pm2`](https://pm2.keymetrics.io/docs/usage/log-management/) might be a good tool to do that.

- For the longest time, we have been thinking about _other human developers_ when creating code. **Now we also have to think about _agents_**.

  - Luckily for us, what "works" for other people, also works for _agents_, and is common sense: consistency, clear organization and clear names are the keys to working with _agents_ within any codebase.

  - **You want the _agent_ to have the ability to correct itself. This means that they need to have an ability to run and write tests**.

    - Be mindful of HOW the _agent_ writes tests. I've seen agents edit existing tests only to make sure they pass, even if they should not.

- Indy talks about _agentic coding KPIs_

  - The size of the output.

  - The amount of attempts you made to get the output.

  - How many times you can sustain your output (the streak).

  - How much time you have to "babysit" the agent (the presence).

- The **hardest part for us to do is to STOP coding**, but you have to do this to fully leverage your tools capabilities.

## Lesson 3

- We looked at _templates_ which in our case, are `.md` files inside the `commands` directory that are quite generic, but embed `$ARGUMENTS` parameter.

  - **For example, consider a `chore` command that describes which actions to take to _create a plan_ for resolving a "chore"**.

    - This is quite powerful as you **created a prompt based on a prompt**.

- There are so-called **Information Dense Keywords** you can use to _steer_ the model.

  - For example, using "think hard" will trigger more reasoning effort while using Claude models. [This blog post](https://www.anthropic.com/engineering/claude-code-best-practices) touches on this topic.

- We should either use `/clear` or start new terminal sessions quite often. Why?

  - You want to get used to working with multiple agents.

  - **You want to ensure the plan you have is self-contained and does not require any context from the "plan creation" phase**. This way, you ensure the plan the LLM produced is of high quality. If it's not, the LLM will make mistakes, but that's good because it's a signal for you that the _template_ might need adjustments.

- You can take templates pretty far. Consider **spinning up multiple agents programmatically**.

  ```bash
  claude -p "/feature \"Create a new query history side panel in the app. This shows all completed queries. Store in the database. order by newest first. Use a new llm call to generate names for each query we'll display in the side panel. When we click we refocus the main panel on the query. Add a 'show panel' button next to 'upload data'\"" \
  --output-format stream-json \
  --dangerously-skip-permissions \
  --model opus \
  --verbose \
  > feature_plan_raw_agent_logs.jsonl
  ```

## Lesson 4 – afk agents

- The lesson started with Andy showing us how they integrate GitHub issues with remote agents.

  1. You create an issue.

  2. This triggers a webhook.

  3. You have a program listening for events for that webhook.

  4. An agent picks up the event, reads the GitHub issue and attempts to write the code necessary to close the issue. The agent uses _templates_ which we discussed in the previous lessons.

  To me, this is a workflow of the future. **Notice that we do not interact with the agent at all**. Everything is done automatically.

- **Instead of fixing code issues, fix the "system" (templates, prompts and so on) that lead to having the issue in the first place**.

  - Consider the leverage you get by doing this. You might feel good about fixing an issue, but what about fixing the root cause of the issue? Even better!

## Lesson 5 – test the loop

- Andy reminds us that the true value we deliver is the value we deliver to users, not the code that we write. Always good to have this in mind.

- **Agents can help you test your codebase. This enables the agent to close the implementation loop greatly increasing the value of their output**.

  - Ask yourself: how do _you_ test the changes you make. Then **encode the answer into a workflow that an agent could use**.

    - It might take some time for you to figure this out, especially in brown-field codebases, but when you do, it will save you _enormous_ amounts of time.

- Tests become more and more important as your usage of agents progresses.

  - How else the agent could know if the work they done is valid? They must have _some_ way of validating the work.

- **It seems like neither the Playwright MCP server OR the Chrome Dev Tools MCP server support network mocking**.

  - I wanted to test the whether the debounce functionality is working on FE. Usually, to test it reliably, you delay the backend network request via `page.route`. **This API does not work when using Playwright MCP server**.

    - I believe the Chrome Dev Tools do not support this functionality at all.

## Lesson 6 – documentation via agents

- **One agent, one prompt, one focus**. You let the agents _focus_ on a single task. The agent can take advantage of the whole 200k+ context window.

- Let the agent write documentation when it finished the work.

  - **This allows other agents to use the documentation it produced when implementing other features**.

  - In addition to documentation, Andy also has a file that tells the agent to _read_ the relevant documentation files if the agent is working in a given place of the codebase.

  ```md
      ## Conditional Docs Entry Format

  After creating the documentation, add this entry to `.claude/commands/conditional_docs.md`:

  - app_docs/<your_documentation_file>.md
    - Conditions:
      - When working with <feature area>
      - When implementing <related functionality>
      - When troubleshooting <specific issues>
  ```

## Lesson 7

- Andy argues, that at some point, you can completely "outsource yourself". **Z**ero **T**ouch **E**ngineering.

  - In the video, they spun up five workflows to implement various features and fix bugs. The more tasks in parallel, the bigger the bottleneck of human review.

    - My take is: It's definitely a future, but I doubt we can trust agents with coming up with sound architectural decisions just yet.

This lesson was mostly of showcase of the system capabilities. Nothing that practical here.

## Lesson 8

- This lesson was mostly a showcase of what the _adw_ (agentic developer workflow) system is capable of.

  - **It is worth investing in building your own system to greatly increase your leverage**. The future is you _managing_ the agents, and not writing code.

- A good question to ask yourself is: **are you working on the _agentic layer_ or the _application layer_?**

  - You want to be working on the _agentic layer_ as much as possible. **That is where the most ROI comes from nowadays**.

## Lesson 9 – Context Management

- The more tokens in the agent's context, the less "focused" it is.

  - It's getting better and better, but models still have issues with large context windows.

- Consider how many MCP servers do you really need. **Most MCP servers consume huge amounts of tokens, but you rarely use them**.

  - You can use `--mcp-config` and `--strict-mcp-config` flags to really have some control over what you load.

Finished Lesson 9 11:54
