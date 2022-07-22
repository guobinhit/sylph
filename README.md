# sylph

![author](https://img.shields.io/badge/author-chariesgavin-blueviolet.svg)[![Go](https://github.com/guobinhit/sylph/actions/workflows/go.yml/badge.svg)](https://github.com/guobinhit/sylph/actions/workflows/go.yml)![last
commit](https://img.shields.io/github/last-commit/guobinhit/sylph.svg)![issues](https://img.shields.io/github/issues/guobinhit/sylph.svg)![stars](https://img.shields.io/github/stars/guobinhit/sylph.svg)![forks](https://img.shields.io/github/forks/guobinhit/sylph.svg)![license](https://img.shields.io/github/license/guobinhit/sylph.svg)


Sylph is the fairy of the wind. It is said that the breeze is the whisper of the fairy. Anyone with a pure heart will
eventually become Sylph.

# Contributing

Contributions are very welcome!

If you see a problem that you'd like to see fixed, the best way to make it happen is to help out by submitting a pull request implementing it. Refer to the [CONTRIBUTING.md](../master/CONTRIBUTING.md) file for more details about the workflow.

You can also ask for problem-solving ideas and discuss in GitHub issues directly.

# INDEX

- [Overview](#overview)
- [Usage](#usage)

## Overview

In `common` package, provides some practical util, as below:

- `dates`: includes some practical date util, such as `GetTodayStart`.
  - `format`: format a `time` value to `string`, such as `GetYyyyMmDdHhMmSs`,
  - `parse`: parse a `string` value to `time`, such as `GetYyyyMmDdHhMmSs`.
- `maps`: includes some practical map util, such as `Keys`.
- `maths`: includes some practical math util, such as `RangeRandomLCRO`.
- `pointers`: convert base type to pointer, such as `Int`.
- `slices`: check base type slice contains target element or not, such as `StringContainsIgnoreCase`.
- `strings`: check string element equals or not, can ignore case, such as `EqualsIgnoreCase`.
- `unpointers`: convert pointer to base type, such as `Int64OrDefault`.
- `utils`: includes some practical util, such as `If`.

In `constant` package, provides some practical constant definitions, as below:

- `date_const`: includes some dates format constant, such as `YyyyMmDdHhMmSs`.
- `string_const`: includes some string constant, such as `EmptyString`.

## Usage

Firstly, download this pkg,

```go
go get github.com/guobinhit/sylph
```

Secondly, use it.

### dates

```go
import (
    "github.com/guobinhit/sylph/common/dates"
    "github.com/guobinhit/sylph/common/dates/format"
    "github.com/guobinhit/sylph/common/dates/parse"
)

// Get a specified time by add days, such as d is 2022-04-13 10:20:30 and days is 10, then aDate is 2022-04-23 10:20:30
aDate := dates.GetTimeAddDays(time.Now(), 10)

// Get a specified date format time string, such as d is 2022-04-13 10:20:30.999, then aString is "2022-04-23 10:20:30"
aString := format.GetYyyyMmDdHhMmSs(time.Now())

// Get a specified date format time string of china version, such as d is 2022-04-13 10:20:30.999, then aString is "2022年04月23日 10:20:30"
aString2 := format.GetCnOfYyyyMmDdHhMmSs(time.Now())

// Get a specified date format time, such as dStr is "2022-04-13 10:20:30", then aTime is 2022-04-23 10:20:30
aTime, err := parse.GetYyyyMmDdHhMmSs("2022-04-13 10:20:30")
```

### maps

```go
aKeySlice := maps.Keys(map[string]string{"a":1, "b":2})
```

### maths

```go
aRandInt := maths.RangeRandomLCRO(1, 10)
```

### pointers

```go
aIntPtr := pointers.Int(413)
```

### unpointers

```go
aInt := unpointers.IntOrDefault(pointer.Int(413), 0)
```

### slices

```go
aBool := slices.StringContainsIgnoreCase([]sring{"abc", "efg"}, "ABC")
```

### strings

```go
aBool := strings.EqualsIgnoreCase("abc", "ABC")
```

### utils

```go
aJsonString := utils.Json(struct{Value string}{Value: "sylph"})
```

Finally, good luck guys!
