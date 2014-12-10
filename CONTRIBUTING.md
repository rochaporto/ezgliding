# Contributing to ezgliding

### Getting and building

A working [Go environment](http://golang.org/doc/code.html) is presumed.
```bash
go get github.com/rochaporto/ezgliding
cd $GOPATH/src/github.com/rochaporto/ezgliding
go install
go test -v ./...
```

### Code

Submit all contributions as pull requests.

Feel free to keep adding new commits during the discussion, but before finally
merging squash them into a single one, with a sensible commit log and a
reference to the issue being tackled (if any).

This should work (from within the branch):
```
git rebase master --interactive
pick ...
squash ...
squash ...

git push -f
```

### Merging a pull request

To try to keep the git history clean, after rebasing the branch with master as
above, merge into master as in:
```
git checkout master
git merge <branch-name>
git push
```

### Bugs, features, code review

GitHub for **everything**.

