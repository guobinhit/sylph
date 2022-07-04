package math

import (
	"fmt"
	"testing"
)

func BenchmarkRangeRandomLCRC(b *testing.B) {
	left := 1
	right := 10
	for i := 0; i < b.N; i++ {
		got := RangeRandomLCRC(left, right)
		if got < left || got > right {
			msg := fmt.Sprintf("RangeRandomLCRC [%d, %d] error, got: %d", left, right, got)
			fmt.Println(msg)
		}
		if got == left || got == right {
			msg := fmt.Sprintf("RangeRandomLCRC [%d, %d] meet boundary value, got: %d", left, right, got)
			fmt.Println(msg)
		}
	}
}

func BenchmarkRangeRandomLCRO(b *testing.B) {
	left := 1
	right := 10
	for i := 0; i < b.N; i++ {
		got := RangeRandomLCRO(left, right)
		if got < left || got > right {
			msg := fmt.Sprintf("RangeRandomLCRO [%d, %d) error, got: %d", left, right, got)
			fmt.Println(msg)
		}
		if got == left || got == right {
			msg := fmt.Sprintf("RangeRandomLCRO [%d, %d) meet boundary value, got: %d", left, right, got)
			fmt.Println(msg)
		}
	}
}

func BenchmarkRangeRandomLORO(b *testing.B) {
	left := 1
	right := 10
	for i := 0; i < b.N; i++ {
		got := RangeRandomLORO(left, right)
		if got <= left || got >= right {
			msg := fmt.Sprintf("RangeRandomLORO (%d, %d) error, got: %d", left, right, got)
			fmt.Println(msg)
		}
	}
}
