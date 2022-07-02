# sylph

![author](https://img.shields.io/badge/author-chariesgavin-blueviolet.svg)![last
commit](https://img.shields.io/github/last-commit/guobinhit/sylph.svg)![issues](https://img.shields.io/github/issues/guobinhit/sylph.svg)![stars](https://img.shields.io/github/stars/guobinhit/sylph.svg)![forks](    https://img.shields.io/github/forks/guobinhit/sylph.svg)![license](https://img.shields.io/github/license/guobinhit/sylph.svg)

Sylph is the fairy of the wind. It is said that the breeze is the whisper of the fairy. Anyone with a pure heart will
eventually become Sylph.

# INDEX

- [Overview](#overview)
- [Use Example](#use-example)

## Overview

In `common` package, includes to some util, as below:

- `pointer`: convert base type to pointer.
- `unpointer`: convert pointer to base type.
- `slice`: check base type slice contains target element or not.
- `string`: check string element equals or not, can ignore case.

## Use Example

Firstly, download this pkg,

```go
go get github.com/guobinhit/sylph
```

Secondly, use it:

```go
aIntPtr := pointer.Int(1120)
aInt := unpointer.IntOrDefault(1120)
aBool := slice.StringContainsIgnoreCase([]sring{"abc", "efg"}, "ABC")
aBool2 := string.EqualsIgnoreCase("abc", "ABC")
```

Finally, good luck guys!