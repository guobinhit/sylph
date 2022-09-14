# diffs

![author](https://img.shields.io/badge/author-chariesgavin-blueviolet.svg)[![Go](https://github.com/guobinhit/sylph/actions/workflows/go.yml/badge.svg)](https://github.com/guobinhit/sylph/actions/workflows/go.yml)![last commit](https://img.shields.io/github/last-commit/guobinhit/sylph.svg)![issues](https://img.shields.io/github/issues/guobinhit/sylph.svg)![stars](https://img.shields.io/github/stars/guobinhit/sylph.svg)![forks](https://img.shields.io/github/forks/guobinhit/sylph.svg)![license](https://img.shields.io/github/license/guobinhit/sylph.svg)

# INDEX

- [JsonDiff](#jsondiff)
- [TextDiff](#textdiff)

# JsonDiff

The package of `jsondiff` used to compare two json difference.
There are three entities that need attention, as follows:

- `DiffType`, means compare type, or result type
  - `FullMatch`
  - `SupersetMatch`
  - `NoMatch`
  - `FirstArgIsInvalidJson`
  - `SecondArgIsInvalidJson`
  - `BothArgsAreInvalidJson`
- `Options`, can use to custom print result.
- `Tag`, a field type of `Options`.

In addition, there are three methods to facilitate use to build `options` param
- `DefaultJsonOptions`
- `DefaultConsoleOptions`
- `DefaultHtmlOptions`


```go
import (
    "github.com/guobinhit/sylph/diffs/jsondiff"
)

// Build compare param, convert to []byte type, call Compare method,
// return two result, diffType can ignore, diffStr is compare result.
firstParam := map[string]interface{}{}
secondParam := map[string]interface{}{}
opts := DefaultJsonOptions()
diffType, diffStr := jsondiff.Compare([]byte(utils.Json(firstParam)), []byte(utils.Json(secondParam)), opts)
```


# TextDiff

The package of `textdiff` used to compare two text difference.

```go
import (
    "github.com/guobinhit/sylph/diffs/textdiff"
)

// Build compare param, init DiffMatchPatch object, call Diff-style method,
// return []Diff, but Diff.Type is number, can call GetDiffPrettyStyle method get pretty result.
text1 := "Hello World"
text2 := "Hello Girl"
dmp := textdiff.New()
diffs := dmp.DiffMain(text1, text2, false)
fmt.Println(utils.Json(dmp.GetDiffPrettyStyle(diffs)))
```



