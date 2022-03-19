# Command overview

## Getting started

* [`chezmoi doctor`](/reference/commands/doctor/) checks for common problems. If you encounter something unexpected, run this first.

* [`chezmoi init`](/reference/commands/init/) creates chezmoi's source directory and a git repo on a new machine.

## Daily commands

* [`chezmoi add <file>`](/reference/commands/add/) adds `<file>`from your home directory to the source directory.

* [`chezmoi edit <file>`](/reference/commands/edit/) opens your editor with the file in the source directory that corresponds to `<file>`.

* [`chezmoi status`](/reference/commands/status/) gives a quick summary of what files would change if you ran `chezmoi apply`.

* [`chezmoi diff`](/reference/commands/diff/) shows the changes that `chezmoi apply` would make to your home directory.

* [`chezmoi apply`](/reference/commands/apply/) updates your dotfiles from the source directory.

* [`chezmoi edit --apply <file>`](/reference/commands/edit/) is like `chezmoi edit <file>` but also runs `chezmoi apply <file>` afterwards.

* [`chezmoi cd`](/reference/commands/cd/) opens a subshell in the source directory.

```mermaid
sequenceDiagram
    participant H as home directory
    participant W as working copy
    participant L as local repo
    participant R as remote repo
    H->>W: chezmoi add &lt;file&gt;
    W->>W: chezmoi edit &lt;file&gt;
    W-->>H: chezmoi status
    W-->>H: chezmoi diff
    W->>H: chezmoi apply
    W->>H: chezmoi edit --apply &lt;file&gt;
    H-->>W: chezmoi cd
```

## Using chezmoi across multiple machines

* [`chezmoi init <github-username>`](/reference/commands/init/) clones your dotfiles from GitHub into the source directory.

* [`chezmoi init --apply <github-username>`](/reference/commands/init/) clones your dotfiles from GitHub into the source directory and runs `chezmoi apply`.

* [`chezmoi update`](/reference/commands/update/) pulls the latest changes from your remote repo and runs `chezmoi apply`.

* Use normal git commands to add, commit, and push changes to your remote repo.

```mermaid
sequenceDiagram
    participant H as home directory
    participant W as working copy
    participant L as local repo
    participant R as remote repo
    R->>W: chezmoi init &lt;github-username&gt;
    R->>H: chezmoi init --apply &lt;github-username&gt;
    R->>H: chezmoi update &lt;github-username&gt;
    W->>L: git commit
    L->>R: git push
```

## Working with templates

* [`chezmoi data`](/reference/commands/data/) prints the available template data.

* [`chezmoi add --template <file>`](/reference/commands/add/) adds `<file>` as a template.

* [`chezmoi chattr +template <file>`](/reference/commands/chattr/) makes an existing file a template.

* [`chezmoi cat <file>`](/reference/commands/cat/) prints the target contents of `<file>`, without changing `<file>`.

* [`chezmoi execute-template`](/reference/commands/execute-template/) is useful for testing and debugging templates.