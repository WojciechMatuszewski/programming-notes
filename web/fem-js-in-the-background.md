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

### Background Sync

- The browser can **wake your service worker up** when it deems it appropriate.

- Keep in mind that **the service worker CAN NOT use local-storage**.

  - To sync the data, you **most likely will need to use the indexedDB**.

- **This API works only in SOME browsers** and **it might not be available due to users browser permissions**.

  - Libraries like **workbox have their internal fallbacks if the sync API is not available**.

  - For me, using Brave, I cannot use this feature.

#### Periodic background sync

- Based on the best-effort scenario.

  - The browser will not execute your code if the battery is low or similar.

  - **Developers and users have NO control if the sync is executed or not**. It is purely a browser thing.

- To know if you can register a background-sync or a periodic-background-sync use the `navigator.permissions.query` API.

### Background Fetch

- **Ask the browser to download some files**.

  - You **do NOT download these files to the file system**. Instead you download them to the app itself.

    - The UI on the browser is a bit misleading. On Brave it looks like downloading files to a file system.

### Push Notifications

- Do not ask to allow for push notifications instantly.

  - Follow best practices here.

- Goes thought the **push server which belongs to the browser**.

  - You send requests to the browser push server. This creates the notification the user sees in the browser.

- Different browsers **have different heuristics they apply when deciding whether to allow you to request push notifications permissions**.

  - For example, if Chrome thinks you abuse the feature, it will NOT display the popup for a new user.

- The flow of the notification is a bit complex.

  1. You have to save the browser endpoint + the private/public keys.
  2. You send the payload to the browser endpoint.
  3. The browser wakes up the service worker and sends the event to the service worker.
  4. The service worker creates a push message that the end-user will see.

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

## Push notifications / local notifications

- They require permissions.

  - There are two APIs to create notifications, but only **a single permission to grant the ability to use them**.

    - This creates confusion.
