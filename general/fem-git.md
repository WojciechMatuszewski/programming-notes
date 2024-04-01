# Everything you will need to know about git

Part 3 00:00

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

-   When using `git merge` you will, in some cases, need to provide a commit message, and sometimes you won't

    -   It all depends on how the branches diverged and where is the _most common ancestor_ for both branches.
    -   When merge is of type _fast-forward merge_, you will not need to provide the commit message.

-   The `merge` is different from `rebase`.

    -   The `rebase` allows you to have a "clean" history without _merge commits_.
    -   The `rebase` alters the history of the branch. In some cases, you will need to force push to a remote to sync the
        remote state with what you have locally.

-   The `git reflog` contains **all the changes you have made to the tree**. This also **includes branches you have
    deleted**.
    -   This is quite useful to know, because, **with `git reflog`, you can get data from branches that were already deleted**.
