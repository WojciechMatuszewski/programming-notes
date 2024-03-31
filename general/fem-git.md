# Everything you will need to know about git

Part 2 00:00

-   Every git repo comes with a `.git` directory.

    -   This directory contains all the state of the git repo.
    -   **Git does not store diffs, it stores ALL the files for every commit you make**.

-   If you delete the `.git` folder, you no longer have a repo. All your commits are gone (they might still live on some
    remote).
-   Use `git log` to view the history of the repository.

-   The _commit SHA_ is based not only on the contents of the change, but also on the credentials of the user who
    commited.

    -   Things like username and email are also taken into account.
    -   This means that you could have two different SHAs for the same change given different settings.
        ls

-   You can set arbitrary keys in the `.gitconfig`. Some of them are picked up by git (like the `username`).
    -   Use the `git config --add` to add anything to the config.
