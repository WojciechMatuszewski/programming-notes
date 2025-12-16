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

- In the course, we used a very simple prompt to force Claude Code to list all tools it has access to. It never occurred to me to do that...

  - Such a great information to have!

- **Your applications must have a way to communicate issues to the Agent**.

  - Are you logging the errors of your applications to stdout? If not, the Agent has no way of knowing what the issue might be (unless you copy & paste the issue manually, but that's quite painful to do).

  - In the course, the run the application via script but _inside_ Claude Code. This enabled Claude Code to read the application logs.

    - As an alternative, we could write application logs to a file. I believe [`pm2`](https://pm2.keymetrics.io/docs/usage/log-management/) might be a good tool to do that.

Finished lesson 2 19:16
