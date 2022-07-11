package unpointers

import (
	"github.com/guobinhit/sylph/common/pointers"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInt(t *testing.T) {
	is := assert.New(t)
	is.Equal(123, Int(pointers.Int(123)))
	is.Equal(0, Int(nil))
}

func TestInt8(t *testing.T) {
	is := assert.New(t)
	is.Equal(int8(123), Int8(pointers.Int8(123)))
	is.Equal(int8(0), Int8(nil))
}

func TestInt16(t *testing.T) {
	is := assert.New(t)
	is.Equal(int16(123), Int16(pointers.Int16(123)))
	is.Equal(int16(0), Int16(nil))
}

func TestInt32(t *testing.T) {
	is := assert.New(t)
	is.Equal(int32(123), Int32(pointers.Int32(123)))
	is.Equal(int32(0), Int32(nil))
}

func TestInt64(t *testing.T) {
	is := assert.New(t)
	is.Equal(int64(123), Int64(pointers.Int64(123)))
	is.Equal(int64(0), Int64(nil))
}

func TestFloat32(t *testing.T) {
	is := assert.New(t)
	is.Equal(float32(123), Float32(pointers.Float32(float32(123))))
	is.Equal(float32(0), Float32(nil))
}

func TestFloat64(t *testing.T) {
	is := assert.New(t)
	is.Equal(float64(123), Float64(pointers.Float64(float64(123))))
	is.Equal(float64(0), Float64(nil))
}

func TestBool(t *testing.T) {
	is := assert.New(t)
	is.Equal(true, Bool(pointers.Bool(true)))
	is.Equal(false, Bool(nil))
}

func TestString(t *testing.T) {
	is := assert.New(t)
	is.Equal("zora", String(pointers.String("zora")))
	is.Equal("", String(nil))
}

func TestIntOrDefault(t *testing.T) {
	is := assert.New(t)
	is.Equal(123, IntOrDefault(pointers.Int(123), 0))
	is.Equal(123, IntOrDefault(nil, 123))
}

func TestInt8OrDefault(t *testing.T) {
	is := assert.New(t)
	is.Equal(int8(123), Int8OrDefault(pointers.Int8(123), 0))
	is.Equal(int8(123), Int8OrDefault(nil, int8(123)))
}

func TestInt16OrDefault(t *testing.T) {
	is := assert.New(t)
	is.Equal(int16(123), Int16OrDefault(pointers.Int16(123), 0))
	is.Equal(int16(123), Int16OrDefault(nil, int16(123)))
}

func TestInt32OrDefault(t *testing.T) {
	is := assert.New(t)
	is.Equal(int32(123), Int32OrDefault(pointers.Int32(123), 0))
	is.Equal(int32(123), Int32OrDefault(nil, int32(123)))
}

func TestInt64OrDefault(t *testing.T) {
	is := assert.New(t)
	is.Equal(int64(123), Int64OrDefault(pointers.Int64(123), 0))
	is.Equal(int64(123), Int64OrDefault(nil, int64(123)))
}

func TestFloat32OrDefault(t *testing.T) {
	is := assert.New(t)
	is.Equal(float32(123), Float32OrDefault(pointers.Float32(float32(123)), 0))
	is.Equal(float32(123), Float32OrDefault(nil, float32(123)))
}

func TestFloat64OrDefault(t *testing.T) {
	is := assert.New(t)
	is.Equal(float64(123), Float64OrDefault(pointers.Float64(float64(123)), 0))
	is.Equal(float64(123), Float64OrDefault(nil, float64(123)))
}

func TestBoolOrDefault(t *testing.T) {
	is := assert.New(t)
	is.Equal(false, BoolOrDefault(pointers.Bool(false), true))
	is.Equal(true, BoolOrDefault(nil, true))
}

func TestStringOrDefault(t *testing.T) {
	is := assert.New(t)
	is.Equal("zora", StringOrDefault(pointers.String("zora"), ""))
	is.Equal("zora", StringOrDefault(nil, "zora"))
}
