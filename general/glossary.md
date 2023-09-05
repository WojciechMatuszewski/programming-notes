# Glossary

Terms, words I was not familiar with at first.

- **Load shedding**: act of dropping work to preserve the system availability.
  It's much better, in the scenario of overload, to prioritize some of the requests and process those and drop others than to try
  to fullfil every request. [More on this topic here](https://aws.amazon.com/builders-library/using-load-shedding-to-avoid-overload/)

- **Heuristics**: in software, it refers to the "way of doing things". Usually, derived from experience and learning based techniques.

- **Lock-step deployment**: where you have two or more services which need to be deployed together. This indicates tight coupling and might lead to problems and deadlocks later on. An excellent example of a _lock-step deployment_ problem is **an API backed by AWS Lambda FURL and the frontend** – the Lambda FURL URL tightly coupled to the function name. If that changes, the URL changes (the solution might be a DNS record).

TODO: <https://meiert.com/en/blog/the-web-development-glossary-3k/>
