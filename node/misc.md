# Misc changes / stuff

## The native `.env` file support

- In v20.6.0 Node.js added [native `.env` file support](https://nodejs.org/en/blog/release/v20.6.0).

- Usage via the `--env-file` flag, like so: `node foo.js --env-file=config.env`

  - **You can define any `NODE_OPTION` in the `.env` file as well!** Pretty neat stuff.

- This **might make the `dotenv` package obsolete**.
