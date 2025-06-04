package cmd

/*
currentLimits:
- unknown

checks:
- .berk/config file exists

actions:
- returns current status of untracked, staged & committed files (examples below)


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

currentLimits:
-

checks:
-

actions:
-
*/

func commandStatus() error {
	return nil
}
