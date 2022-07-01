# sylph

![author](https://img.shields.io/badge/author-chariesgavin-blueviolet.svg)![last commit](https://img.shields.io/github/last-commit/guobinhit/sylph.svg)![issues](https://img.shields.io/github/issues/guobinhit/sylph.svg)![stars](https://img.shields.io/github/stars/guobinhit/sylph.svg)![forks](	https://img.shields.io/github/forks/guobinhit/sylph.svg)![license](https://img.shields.io/github/license/guobinhit/sylph.svg)

Sylph is the fairy of the wind. It is said that the breeze is the whisper of the fairy. Anyone with a pure heart will eventually become Sylph.

# INDEX
- [Overview](#overview)
- [Use Example](#use-example)

## Overview
You can use it to convert the base type to pointer conversion, or to convert the pointer of the base type to non pointer type.

## Use Example
Firstly, download this pkg,
```go
go get github.com/guobinhit/sylph
```
Secondly, use it:
```go
aIntPtr := pointer.Int(1120)
aInt := unpointer.IntOrDefault(1120)
```
Finally, good luck guys!