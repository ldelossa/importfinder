# ImportFinder

ImportFinder is a simple tool that can be ran from the root of any directory. The application will walk from the root recursively finding all files with .go extenstions.
Once inventoried we begin to use Go's parser library to extract the import information in a concurrent fashion. 

# Installing
feel free to acquire the source with go get
```
go get github.com/ldelossa/importfinder
```
and then go install within the /cmd/importfinder directory
```
❯ cd github.com/ldelossa/importfinder/cmd/importfinder
~/git/go/src/github.com/ldelossa/importfinder/cmd/importfinder

~/git/go/src/github.com/ldelossa/importfinder/cmd/importfinder master
❯ go install
```
the importfinder binary should now be located in your $GOPATH/bin directory.

# Usage

Simply build the application and run it at the root of a go directory

```
~/git/go/src/github.com/mdempsky
❯ $GOBIN/importfinder | jq .
{
  "gocode/client.go": [
    "\"flag\"",
    "\"fmt\"",
    "\"go/build\"",
    "\"io/ioutil\"",
    "\"log\"",
    "\"net/rpc\"",
    "\"os\"",
    "\"path/filepath\"",
    "\"strconv\"",
    "\"time\"",
    "\"runtime/debug\"",
    "\"github.com/mdempsky/gocode/internal/cache\"",
    "\"github.com/mdempsky/gocode/internal/suggest\""
  ],
  "gocode/gocode.go": [
    "\"flag\"",
    "\"fmt\"",
    "\"os\"",
    "\"path/filepath\""
  ],
  ...

```

