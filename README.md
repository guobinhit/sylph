# sylph

![author](https://img.shields.io/badge/author-chariesgavin-blueviolet.svg)![last
commit](https://img.shields.io/github/last-commit/guobinhit/sylph.svg)![issues](https://img.shields.io/github/issues/guobinhit/sylph.svg)![stars](https://img.shields.io/github/stars/guobinhit/sylph.svg)![forks](https://img.shields.io/github/forks/guobinhit/sylph.svg)![license](https://img.shields.io/github/license/guobinhit/sylph.svg)

Sylph is the fairy of the wind. It is said that the breeze is the whisper of the fairy. Anyone with a pure heart will
eventually become Sylph.

# INDEX

- [Overview](#overview)
- [Use Example](#use-example)

## Overview

In `common` package, provides some practical util, as below:

- `math`: includes some practical math util, such as `RangeRandomLCRO`.
- `pointer`: convert base type to pointer, such as `Int`.
- `slice`: check base type slice contains target element or not, such as `StringContainsIgnoreCase`.
- `string`: check string element equals or not, can ignore case, such as `EqualsIgnoreCase`.
- `unpointer`: convert pointer to base type, such as `Int64`.
- `util`: includes some practical util, such as `If`.

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
aRandInt := math.RangeRandomLCRO(1, 10)
aIntPtr := pointer.Int(413)
aInt := unpointer.IntOrDefault(413)
aBool := slice.StringContainsIgnoreCase([]sring{"abc", "efg"}, "ABC")
aBool2 := string.EqualsIgnoreCase("abc", "ABC")
```

Finally, good luck guys!