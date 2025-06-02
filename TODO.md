## Now

### Init
- add an `init` command
- creates a `.berk` directory
- inside directory creates the following
    - hooks/
    - info/
    - objects/
    - refs/
    - config
    - description
    - HEAD

## Next

### Config
- on install (or first use?), create a `~/.berkconfig` file
- in this file store the username, useremail and default branch

### Status

A file can be in one of several states in a Git repository:
- `untracked`: Not being tracked by Git
- `staged`: Marked for unclusion in the next commit
- `committed`: Saved to the repository's history

The `git status` command shows you the current state of your repo. It will tell you which files are untracked, staged and committed.

```
On branch master

No commits yet

Untracked files:
  (use "git add <file>..." to include in what will be committed)
        contents.md

nothing added to commit but untracked files present (use "git add" to track)
```

```
On branch master

No commits yet

Changes to be committed:
  (use "git rm --cached <file>..." to unstage)
        new file:   contents.md
```

```
On branch master
nothing to commit, working tree clean
```

### Add

The `contents.md` file has been created, but as we saw it's untracked. We need to stage it (add it to the "index") with the `git add` command before committing it later.

Without staging, every file in the repository would be included in every commit, but that' often not what you want.

- Note: git add will autocomplete the file name...

### Commit 

After staging a file, we can commit it.

A commit is a snapshot of the repository at a given point in time. It's a way to save the state of the repository, and it's how Git keeps track of changes to the project. A commit comes with a message that describes the changes made in the commit.

### Log

A Git repo is a (potentially very long) list of commits, where each commit represents the full state of the repository at a given point in time. The `git log` command shows a history of the commits in a repository. This is what makes Git a version control system. You can see:

- Who made a commit
- When the commit was made
- What was changed

Each commit has a unique indentifier called a "commit hash". This is a long string of characters that uniquely identifies the commit. 

## CLI Commands

- `-v` or `--version` - returns version
- `-h` or `--help` - displays help
- `config` - Get and set repository or global options
- `init` - Create an empty Berk repository
- `add` - Add file contents to the index
- `commit` - Record changes to the repository
- `status` - Show the working tree status
- `diff` - Show changes between commits
- `log` - Shows a history of the commits in a repository

## Others

- Create a man page
