# JS-TS monorepos FE

- the `@` is a convention for naming packages.

- use `composite` option for building monorepos with _TypeScript_. With this option enabled, _TypeScript_ can determine better which files needs to be re-build when change happens in one or more packages.

- the `references` option allows you to build multiple projects at the same time. Any projects that you specify here, has to have `composite` within their `tsconfig.json`. Dependencies between packages are also managed. If you import _package A_ from _package B_, the import will refer to the build output of _package A_.

- just like you can extend the `tsconfig` you can extend the `.babelrc` config.

- specify only **functionality agnostic tools** on the root workspace level. Something like `rimraf` or `cross-env` or `eslint`. You will have different versions of libraries
  in different packages. Imagine trying to update _TypeScript_ in 30+ packages when _TypeScript_ is specified as root dependency, good luck!

- _lerna_ also works with npm. Nice to know but may be less relevant with _npm 7.0_ which supports _workspaces_.

- _lerna_ can run your commands in **parallel**. This is huge. _lerna_ will figure all the inner dependencies in between packages and decide how to build stuff most efficiently.
