# Everything you will need to know about git

> Going through [this workshop](https://frontendmasters.com/workshops/git/).

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

-   **Use the `git stash -m "YOUR MESSAGE`** instead of the "bare" `git stash`.

    -   Adding the message will come in very handy when using `git stash list`. Otherwise, it is hard to figure out what the stash contains.
        -   Of course, one could use the `git stash show` to view the diff.

-   `git rebase` could be dangerous as **you could erase some commits from the history**.

    -   In addition, if you had one conflict during the `rebase`, any more additions to that file, you will need to resolve the conflict again.
    -   This is where the `rerere` setting comes in.
        -   It will re-use the conflict resolution you have made in the previous rebase when rebasing again.

-   Use the `git rebase -i` to manipulate the history.

    -   I find it handy for _squashing_ commits, but there are other options as well!

-   Let us consider a situation where you have a bug on your branch. You know a commit where the bug is not there, but you do not know when it was introduced.
    -   Here, consider using the `git bisect`. It will help you narrow down the change that introduced the bug.
    -   **In addition, you can provide a script for `git bisect` to run (the `run` flag), and it will automatically do the `good`/`bad` for you**.
        -   This is super neat!

## Finishing up

A good workshop with just enough information to be productive using git.
I found `git bisect run` mind-blowing. Hopefully, I will not need to use it anytime soon, but If I do, I'm going to use the `run` option.
