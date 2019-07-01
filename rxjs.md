# RxJS

## Properly handling errors using `catchError`

So you've been using `catchError` just treating it like `.catch` in a promise-base API and it all seemed all good and sweet. But sometimes you encountered a bug where a stream would not be called again after an error. _But you caught the error with `catchError` and returned a new stream_, what could go wrong?🤔. Well just know that:
**`catchError` replaces whole stream, WHOLE STREAM**. Now lets see an example:

```typescript
source$.pipe(
  // switchMap can fail
  switchMap(something => from(resourceGetterFn(something))).pipe(
    // resolveResourceResponse can fail
    mergeMap(response => resolveResourceResponse(response))
  )
);
```

Now, what would happen when we did this:

```typescript
source$.pipe(
  // previous code with switchMap etc
  catchError(_ => {
    return of();
  })
);
```

So, error is propagated and is caught by `catchError`, that's all and good. But again **CATCH ERROR REPLACES WHOLE STREAM!**(and we are returning an empty Observable). That means, **after an error, that operator is just an empty Observable**.

### Solution

Solution would be... well reading the docs and such (and actually understanding what code you are writing). To solve this problem we just need to move `catchError` **inside switchMap**.

```typescript
source$.pipe(
  // switchMap can fail
  switchMap(something => from(resourceGetterFn(something))).pipe(
    mergeMap(response => resolveResourceResponse(response)),
    catchError(_ => {
      return of();
    })
  )
);
```

There, no magic, no weird copy-paste from stack. That's all.

_Reference: [this great article](https://medium.com/city-pantry/handling-errors-in-ngrx-effects-a95d918490d9)_
