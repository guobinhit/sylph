# sylph

![author](https://img.shields.io/badge/author-chariesgavin-blueviolet.svg)![last
commit](https://img.shields.io/github/last-commit/guobinhit/sylph.svg)![issues](https://img.shields.io/github/issues/guobinhit/sylph.svg)![stars](https://img.shields.io/github/stars/guobinhit/sylph.svg)![forks](https://img.shields.io/github/forks/guobinhit/sylph.svg)![license](https://img.shields.io/github/license/guobinhit/sylph.svg)

Sylph is the fairy of the wind. It is said that the breeze is the whisper of the fairy. Anyone with a pure heart will
eventually become Sylph.

# Contributing

Contributions are very welcome!

If you see a problem that you'd like to see fixed, the best way to make it happen is to help out by submitting a pull request implementing it. Refer to the [CONTRIBUTING.md](../master/CONTRIBUTING.md) file for more details about the workflow.

You can also ask for problem-solving ideas and discuss in GitHub issues directly.

# INDEX

- [Overview](#overview)
- [Use Example](#use-example)

## Overview

In `common` package, provides some practical util, as below:

- `dates`: includes some practical date util, such as `GetTodayStart`.
- `maps`: includes some practical map util, such as `Keys`.
- `maths`: includes some practical math util, such as `RangeRandomLCRO`.
- `pointers`: convert base type to pointer, such as `Int`.
- `slices`: check base type slice contains target element or not, such as `StringContainsIgnoreCase`.
- `strings`: check string element equals or not, can ignore case, such as `EqualsIgnoreCase`.
- `unpointers`: convert pointer to base type, such as `Int64OrDefault`.
- `utils`: includes some practical util, such as `If`.

In `constant` package, provides some practical constant definitions, as below:

- `date_const`: includes some dates format constant, such as `EnOfYyyyMmDdHhMmSs`.
- `string_const`: includes some string constant, such as `EmptyString`.

## Use Example

Firstly, download this pkg,

```go
go get github.com/guobinhit/sylph
```

Secondly, use it:

```go
aDate := dates.GetTimeAddDays(time.Now(), 10)
aKeySlice := maps.Keys(map[string]string{"a":1, "b":2})
aRandInt := maths.RangeRandomLCRO(1, 10)
aIntPtr := pointers.Int(413)
aInt := unpointers.IntOrDefault(pointer.Int(413), 0)
aBool := slices.StringContainsIgnoreCase([]sring{"abc", "efg"}, "ABC")
aBool2 := strings.EqualsIgnoreCase("abc", "ABC")
```

Finally, good luck guys!
