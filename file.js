new Promise(r => {
  r();
})
  .then(() => {
    console.log("1");
  })
  .then(() => {
    console.log("2");
  })
  .then(() => {
    return new Promise(r => setTimeout(r, 100));
  })
  .then(() => {
    console.log("3");
  });

setTimeout(() => {
  console.log("timer");
}, 0);

process.nextTick(() => {
  console.log("next tick");
});

setImmediate(() => {
  console.log("immediate");
});

console.log("regular code");
