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

### Docker

To build the docker images, do the following:
```shell
# go build
# sudo docker build -t rochaporto/ezgliding .
# sudo docker build -t rochaporto/ezgliding-web - < Dockerfile-web
# sudo docker images
rochaporto/ezgliding-web   latest              9c034c6d954b        5 seconds ago        9.045 MB
rochaporto/ezgliding       latest              6c09d51e6020        About a minute ago   9.045 MB
scratch                    latest              511136ea3c5a        19 months ago        0 B
```

The first build creates the base image containing the docker binary.

The second build creates the ezgliding web service container image, including
an entrypoint running `ezgliding web`.

### Bugs, features, code review

GitHub for **everything**.
