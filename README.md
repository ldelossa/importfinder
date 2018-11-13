# ImportFinder

ImportFinder is a simple tool that can be ran from the root of any directory. The application will walk from the root recursively finding all files with .go extenstions.
Once inventoried we begin to use Go's parser library to extract the import information in a concurrent fashion. 

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

