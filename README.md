# sylph

![author](https://img.shields.io/badge/author-chariesgavin-blueviolet.svg)[![Go](https://github.com/guobinhit/sylph/actions/workflows/go.yml/badge.svg)](https://github.com/guobinhit/sylph/actions/workflows/go.yml)![last commit](https://img.shields.io/github/last-commit/guobinhit/sylph.svg)![issues](https://img.shields.io/github/issues/guobinhit/sylph.svg)![stars](https://img.shields.io/github/stars/guobinhit/sylph.svg)![forks](https://img.shields.io/github/forks/guobinhit/sylph.svg)![license](https://img.shields.io/github/license/guobinhit/sylph.svg)

> Sylph is the fairy of the wind. It is said that the breeze is the whisper of the fairy. 
> 
> Anyone with a pure heart will eventually become Sylph.

# Contributing

Contributions are very welcome!

If you see a problem that you'd like to see fixed, the best way to make it happen is to help out by submitting a pull request implementing it. Refer to the [CONTRIBUTING.md](../master/CONTRIBUTING.md) file for more details about the workflow.

You can also ask for problem-solving ideas and discuss in GitHub issues directly.

# INDEX

- [Overview](#overview)
- [Usage](#usage)

# Overview

In specified package, provides some practical util, as below:

- `dates`: includes some practical date util, such as `GetTodayStart`.
  - `format`: format a `time` value to `string`, such as `GetYyyyMmDdHhMmSs`,
  - `parse`: parse a `string` value to `time`, such as `GetYyyyMmDdHhMmSs`.
- `diffs`: includes some diff util, as below:
  - `jsondiff`: diff json and return difference item.
  - `textdiff`: diff text and return difference item.
- `maps`: includes some practical map util, such as `Keys`.
- `maths`: includes some practical math util, such as `RangeRandomLCRO`.
- `pointers`: convert base type to pointer, such as `Int`.
- `slices`: check base type slice contains target element or not, such as `ContainStringsIgnoreCase`.
- `strings`: check string element equals or not, can ignore case, such as `EqualsIgnoreCase`.
- `unpointers`: convert pointer to base type, such as `Int64OrDefault`.
- `utils`: includes some practical util, such as `If`.

In `constant` package, provides some practical constant definitions, as below:

- `date_const`: includes some dates format constant, such as `YyyyMmDdHhMmSs`.
- `string_const`: includes some string constant, such as `EmptyString`.

# Usage

Firstly, download this pkg,

```go
go get github.com/guobinhit/sylph
```

Secondly, use it.

## dates

```go
import (
    "github.com/guobinhit/sylph/dates"
    "github.com/guobinhit/sylph/dates/format"
    "github.com/guobinhit/sylph/dates/parse"
)

// Get a specified time by add days, such as d is 2022-04-13 10:20:30 and days is 10,
// then aDate is 2022-04-23 10:20:30
aDate := dates.GetTimeAddDays(time.Now(), 10)

// Get a specified date format time string of common version, such as d is 2022-04-13 10:20:30,
// then aString is "2022-04-23 10:20:30"
aString := format.GetYyyyMmDdHhMmSs(time.Now())

// Get a specified date format time string of china version, such as d is 2022-04-13 10:20:30,
// then aString2 is "2022年04月23日 10:20:30", like GetCnOfYyyyMmDdHhMmSs method, 
// GetEnOfYyyyMmDdHhMmSs return "2022/04/23 10:20:30" and GetCptOfYyyyMmDdHhMmSs return "20220423102030".
aString2 := format.GetCnOfYyyyMmDdHhMmSs(time.Now())

// Get a specified date format time, such as dStr is "2022-04-13 10:20:30",
// then aTime is 2022-04-23 10:20:30
aTime, err := parse.GetYyyyMmDdHhMmSs("2022-04-13 10:20:30")
```

## diffs

```go
import (
    "github.com/guobinhit/sylph/diffs/jsondiff"
    "github.com/guobinhit/sylph/diffs/textdiff"
)

// Build compare param, convert to []byte type, call Compare method,
// return two result, diffType can ignore, diffStr is compare result.
firstParam := map[string]interface{}{}
secondParam := map[string]interface{}{}
opts := DefaultJsonOptions()
diffType, diffStr := jsondiff.Compare([]byte(utils.Json(firstParam)), []byte(utils.Json(secondParam)), opts)

// Build compare param, init DiffMatchPatch object, call Diff-style method,
// return []Diff, but Diff.Type is number, can call GetDiffPrettyStyle method get pretty result.
text1 := "Hello World"
text2 := "Hello Girl"
dmp := textdiff.New()
diffs := dmp.DiffMain(text1, text2, false)
fmt.Println(utils.Json(dmp.GetDiffPrettyStyle(diffs)))
```

## maps

```go
import (
    "github.com/guobinhit/sylph/maps"
)

// Get a slice from map, element of slice is key of map, such as map is {"a":1, "b":2},
// then aKeySlice is {"a", "b"}, we can use assert convert to actual type.
aKeySlice := maps.Keys(map[string]string{"a":1, "b":2})
aActualKeySlice := aKeySlice.([]string)

// Get a slice from map, element of slice is value of map, such as map is {"a":1, "b":2},
// then aValueSlice is {1, 2}, we can use assert convert to actual type.
aValueSlice := maps.Values(map[string]string{"a":1, "b":2})
aActualValueSlice := aValueSlice.([]int)
```

## maths

```go
import (
    "github.com/guobinhit/sylph/maths"
)

// Get a random int value, LCRC means left close right close, left is 1, right is 10
// then min value of aRandIntOfLCRC is 1, max value of aRandIntOfLCRC is 10
aRandIntOfLCRC := maths.RangeRandomLCRC(1, 10)

// Get a random int value, LCRO means left close right open, left is 1, right is 10
// then min value of aRandIntOfLCRO is 1, max value of aRandIntOfLCRO is 9
aRandIntOfLCRO := maths.RangeRandomLCRO(1, 10)

// Get a random int value, LORO means left open right open, left is 1, right is 10
// then min value of aRandIntOfLORO is 2, max value of aRandIntOfLORO is 9
aRandIntOfLORO := maths.RangeRandomLORO(1, 10)
```

## pointers

```go
import (
    "github.com/guobinhit/sylph/pointers"
)

// Get a pointer type of int, 
// supports int, int8, int16, int32, int64, float32, float64 and string.
aIntPtr := pointers.Int(413)
```

## unpointers

```go
import (
    "github.com/guobinhit/sylph/unpointers"
)

// Get a base type value from pointer, 
// supports int, int8, int16, int32, int64, float32, float64 and string.
aInt := unpointers.Int(pointer.Int(413))

// Get a base type value or default value from pointer, if pointer is nil, then return default value,
// supports int, int8, int16, int32, int64, float32, float64 and string.
aIntOrDefault := unpointers.IntOrDefault(pointer.Int(413), 0)
```

## slices

```go
import (
    "github.com/guobinhit/sylph/slices"
)

// Get a distinct slice from param slice, such as param s is {"a","b", "c", "a"},
// then aDistinctSlice is {"a","b", "c"}.
aDistinctSlice := slices.DistinctStrings([]string{"a","b", "c", "a"})

// Get a bool value, if slice param contains specified element e return true, else return false,
// such as s is {"a","b", "c"}, e is "c", then aContainBool true.
// Like ContainStrings method, ContainStringsIgnoreCase method can ignore case.
aContainBool := slices.ContainStrings([]string{"a","b", "c"}, "c")

// Get a filter slice, applies a fn to each element of s, return a slices of make fn true,
// such as fn is func(v string) bool { return len(v) > 3} and s is {"abc", "zora"}, then aFilterSlice is {"zora"}
// Like ContainStrings method, AllStrings and AnyStrings method can provide similar functions,
// supports int, int8, int16, int32, int64, float32, float64 and string.
aFilterSlice := slices.FilterStrings([]string{"abc", "zora"}, func(v string) bool { return len(v) > 3})

// Paging slice and returns the current page results and hasNext mark,
// such as arr is {"1", "2", "3","4", "5"}, pageNum is 2 and pageLimit is 2,
// then aPageSlice is {"3", "4"}, hasNext is true. We can use assert convert to actual type.
aPageSlice, hasNext := slices.Page([]string{"1", "2", "3","4", "5"}, 2, 2)
aActualPageSlice := aPageSlice.([]string)
```

## strings

```go
import (
    "github.com/guobinhit/sylph/strings"
)

// Get a bool value, if s1 is equals to s2 return true, else return false,
// such as s1 is "abc", s2 is "ABC", then aBool is false.
aBool := strings.Equals("abc", "ABC")

// Get a bool value, if s1 is equals to s2 (ignore case) return true, else return false,
// such as s1 is "abc", s2 is "ABC", then aBool2 is true.
aBool2 := strings.EqualsIgnoreCase("abc", "ABC")
```

## utils

```go
import (
    "github.com/guobinhit/sylph/utils"
)

// If function similar to ternary operators，if b is ture return t, else return f.
aValue := utils.If(b, t, f)

// Returns a json format string.
aJsonString := utils.Json(struct{Value string `json:"value"`}{Value: "sylph"})

// DeepCopy source and target must have the same structure, and target must be a pointer.
emptyStruct := struct{Value string `json:"value"`}{}
utils.DeepCopy(struct{Value string `json:"value"`}{Value: "sylph"}, &emptyStruct)
```

## constant

```go
import (
    "github.com/guobinhit/sylph/constant/date_const"
    "github.com/guobinhit/sylph/constant/string_const"
)

// Define dates format constant, use it directly.
aYyyyMmDdHhMmSsFormat := date_const.YyyyMmDdHhMmSs

// Define common strings type constant, use it directly.
aEnglishComma := string_const.EnglishComma
```

Finally, good luck guys!
