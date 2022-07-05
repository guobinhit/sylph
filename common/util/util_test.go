package util

import (
	"github.com/guobinhit/sylph/constant/date_const"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
	"time"
)

func TestIf(t *testing.T) {
	type args struct {
		b bool
		f interface{}
		s interface{}
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		{
			name: "If false",
			args: args{
				b: false,
				f: 1,
				s: 2,
			},
			want: 2,
		},
		{
			name: "If true",
			args: args{
				b: false,
				f: "hello",
				s: "world",
			},
			want: "world",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := If(tt.args.b, tt.args.f, tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("If() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEquals(t *testing.T) {
	type testDataStruct struct {
		Key   int
		Value string
	}

	is := assert.New(t)

	is.True(Equals(int(12), int(12)))
	is.False(Equals(int(12), int(21)))
	is.False(Equals(int(12), float32(12)))

	is.True(Equals(int8(12), int8(12)))
	is.False(Equals(int8(12), int8(21)))
	is.False(Equals(int8(12), float32(12)))

	is.True(Equals(int16(12), int16(12)))
	is.False(Equals(int16(12), int16(21)))
	is.False(Equals(int16(12), float32(12)))

	is.True(Equals(int32(12), int32(12)))
	is.False(Equals(int32(12), int32(21)))
	is.False(Equals(int32(12), float32(12)))

	is.True(Equals(int64(12), int64(12)))
	is.False(Equals(int64(12), int64(21)))
	is.False(Equals(int64(12), float32(12)))

	is.True(Equals(uint(12), uint(12)))
	is.False(Equals(uint(12), uint(21)))
	is.False(Equals(uint(12), float32(12)))

	is.True(Equals(uint8(12), uint8(12)))
	is.False(Equals(uint8(12), uint8(21)))
	is.False(Equals(uint8(12), float32(12)))

	is.True(Equals(uint16(12), uint16(12)))
	is.False(Equals(uint16(12), uint16(21)))
	is.False(Equals(uint16(12), float32(12)))

	is.True(Equals(uint32(12), uint32(12)))
	is.False(Equals(uint32(12), uint32(21)))
	is.False(Equals(uint32(12), float32(12)))

	is.True(Equals(uint64(12), uint64(12)))
	is.False(Equals(uint64(12), uint64(21)))
	is.False(Equals(uint64(12), float32(12)))

	is.True(Equals(float32(12), float32(12)))
	is.False(Equals(float32(12), float32(21)))
	is.False(Equals(float32(12), int64(12)))

	is.True(Equals(float64(12), float64(12)))
	is.False(Equals(float64(12), float64(21)))
	is.False(Equals(float64(12), int64(12)))

	is.True(Equals("zora", "zora"))
	is.False(Equals("zora", "orz"))
	is.False(Equals("123456", 123456))

	dateStr := "2022-04-13 00:00:00"
	t1, _ := time.Parse(date_const.YyyyMmDdHhMmSs, dateStr)
	t2, _ := time.Parse(date_const.YyyyMmDdHhMmSs, dateStr)
	t3, _ := time.Parse(date_const.YyyyMmDdHhMmSs, "2022-06-13 00:00:00")
	is.True(Equals(t1, t2))
	is.False(Equals(t1, t3))
	is.False(Equals(t1, 1234567890))

	d1 := testDataStruct{
		Key:   123456,
		Value: "1234567890",
	}
	d2 := testDataStruct{
		Key:   123456,
		Value: "1234567890",
	}
	d3 := testDataStruct{
		Key:   654321,
		Value: "1234567890",
	}
	is.True(Equals(d1, d2))
	is.False(Equals(d1, d3))
}
