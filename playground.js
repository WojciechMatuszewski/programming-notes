const { setTimeout } = require("timers/promises");

const makeRequest = async ({ signal }) => {
  // Makes the request...
};

const cancelTask = new AbortController();
const cancelTimeout = new AbortController();

const timeout = async () => {
  try {
    await setTimeout(1000, undefined, { signal: cancelTimeout.signal });
    cancelTask.abort();
  } catch {
    return;
  }
};

const task = async () => {
  try {
    await makeRequest({ signal: cancelTask.signal });
  } finally {
    cancelTimeout.abort();
  }
};

const main = async () => {
  Promise.race([timeout, task]);
};

main();
