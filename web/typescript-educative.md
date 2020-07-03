# Typescript course on `educative.io`

## Aliasing `any`

This is something that I have heard about but never got to actually do (because I usually just type the stuff I'm working on).

The idea is that instead of using `any` you alias it to a descriptive name.

```ts
type IDoNotKnowTheTypeYet = any;

function someFunc(param: IDoNotKnowTheTypeYet) {
  // stuff
}
```

This way, the `any` is much more _visible_ when it comes to code review, at least I think it is.
