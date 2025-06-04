## Now

1. Init
2. Config
3. Hash-Objects
4. Cat-File
5. Add
6. Commit
7. Status
8. Log

## Future

- Man page
- Additional locations (system, global & worktree)
- Branches
- Remote

## Notes

### Plumbing

Git uses a cryptographic hash function called `SHA-1` to generate commit hashes.

All the data in a Git repository is stored directly in the (hidden) `.git` directory. That includes  all the commits, branches, tags, and other objects we'll learn about later. Git is made up of objects that are stored in the `.git/objects` directory. A commit is just a type of object.

- `tree`: git's way of storing a directory
- `blob`: git's way of storing a file

Git stores an entire snapshot of files on a per-commit level. Git compresses and packs files to store them more efficiently and deduplicates files that are the same across different commits.


