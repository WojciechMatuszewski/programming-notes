# JavaScript in the Background

TODO: merge with the other file on other mac.

## Media features

- When playing media, you can create a `MediaMetadata` object to control how the OS displays the information about the audio.

```js
navigator.mediaSession.metadata = new MediaMetadata({
    title: "To much Funk",
    artist: "The Funky Bunch",
    album: "Frontend Masters",
    artwork: [
        {
        src: "/media/thumb.png",
        type: "image/png",
        sizes: "800x800"
        }
    ]
});
```

- Pretty neat stuff!

### Picture in Picture

- We previously talked about how browsers throttle timers when we switch the browser window to the background.

  - That is **not the case for the PiP element and the code it executes**.

    - It would be weird for the video to play at 1 FPS right?

## The service worker

- The **notion of a _scope_ is essential**.

  - Since the path corresponds to an URL, if put your service worker on some nested path, the service worker will only be able to intercept paths under that given path.

  - You **most likely want the service worker file to be at the root of your application**.

- In the service worker you mainly add event listeners and then react to them.

  - Keep in mind that **you do not have access to DOM APIs in the service worker**.

## The Beacon API

- Use for **requests that you do not need the response from**.

  - Mainly for **analytics**.

- The neat thing about this API is that **the browser will send this event even if the page is closed**.

- You **cannot set custom headers with this request**.

  - This means that **sending a stringified object might not work for your server**. What you have to do is to **send is a Blob**.

    ```js
    const data = {}
    const blob = new Blob([JSON.stringify(data)], { type: "application/json" });
    ```

- **Beacons with the same data might be deduped**.

### Push notifications / local notifications

- They require permissions.

  - There are two APIs to create notifications, but only **a single permission to grant the ability to use them**.

    - This creates confusion.

Finished part 3
