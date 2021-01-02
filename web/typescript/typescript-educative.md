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

## `strictNullChecks`

This kinda blows my mind since I always opted to the strictest possible setting so I did not know.

When you have this setting as `false`, you can **pass `undefined` and `null` as legit parameters** even though the signature might require a `number` lets say:

```ts
function someFunc(num: number) {
    return 1;
}
// compiles, WAT!?
someFunc(undefined);
```

When you have some paramters listed as optional, Typescript wont scream at you when accessing them without a null check.

```ts
function someFunc(num?: number) {
    // just... wow!
    return num.toFixed();
}

someFunc(undefined);
```

So to sum up, **USE `strictNullCheck` OPTION**.
