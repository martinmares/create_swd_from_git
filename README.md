# create_swd_from_git

## TL;DR

`swd` is shortcut for `[s]oft[w]arová [d]ávka` ... sorry for this, but shortcut is in CZK lang ;-)

Small utility generate "directory tree" structure only changed files `from commit` → `to commit`.

```bash
$ ./binary/mac64/create_swd_from_git --help
Usage of binary/mac64/create_swd_from_git:
  -copyTo string
    	Copy to destination directory (default "./tmp")
  -fromCommit string
    	From commit... (default "...")
  -toCommit string
    	To commit... (default "...")
```

## Howto make

type `make all` in terminal.

output sample:

```bash
$ make all

CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o binary/win64/create_swd_from_git.exe -v
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o binary/linux64/create_swd_from_git -v
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o binary/mac64/create_swd_from_git -v
```
## Howto use

simple type `create_swd_from_git` (or `./create_swd_from_git`) with this parameters `fromCommit`, `toCommit`, `copyTo`:

```bash
$ binary/mac64/create_swd_from_git -fromCommit 6e5d405 -toCommit HEAD -copyTo swd
From commit: 6e5d405
To commit: HEAD
Copy files to: swd
  [000] binary → create_swd_from_git
  [001] binary → create_swd_from_git
  [002] binary → create_swd_from_git.exe
  [003]  → create_swd_from_git.go
```
And result is:

```bash
$ find ./swd

swd/
swd/binary
swd/binary/create_swd_from_git.exe
swd/binary/create_swd_from_git
swd/create_swd_from_git.go
```
