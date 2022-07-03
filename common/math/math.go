package math

import (
	"math/rand"
	"time"
)

func RangeRandom(left, right int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(right-left) + left
}
