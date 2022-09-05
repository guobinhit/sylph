package maths

import (
	"math/rand"
	"time"
)

var (
	intBitSize            = 32 << (^uint(0) >> 32 & 1) // size of an int in bits
	minInt                = -1 << (intBitSize - 1)     // minimum value of int type
	valuesIsEmptyPanicMsg = "values is empty"          // panic message of value is empty
)

// RangeRandomLCRC returns value in [left, right], LCRC means left close right close.
func RangeRandomLCRC(left, right int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(right-left+1) + left
}

// RangeRandomLCRO returns value in [left, right), LCRO means left close right open.
func RangeRandomLCRO(left, right int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(right-left) + left
}

// RangeRandomLORO returns value in (left, right), LORO means left open right open.
// Notice: because of manual +1, so the value of left+1 will be slightly more.
func RangeRandomLORO(left, right int) int {
	rand.Seed(time.Now().UnixNano())
	ans := rand.Intn(right-left) + left
	if ans == left {
		return ans + 1
	}
	return ans
}
