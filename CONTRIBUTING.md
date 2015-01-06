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

### Logging

In the command line tool, enable using (v is 0 to 20):
```
./ezgliding airfield-get -alsologtostderr -v=10
```

In tests, compile the test binary and run with similar flags (on a single package):
```
cd fusiontables
go test -c
./fusiontables.test -alsologtostderr -v=20 -test.v -test.run=GetAirfield
```

### Bugs, features, code review

GitHub for **everything**.

