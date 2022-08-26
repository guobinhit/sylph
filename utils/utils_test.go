package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIf(t *testing.T) {
	is := assert.New(t)
	is.Equal(1, If(true, 1, 2))
	is.Equal("2", If(false, "1", "2"))
}

func TestJson(t *testing.T) {
	is := assert.New(t)
	is.Equal("1", Json(1))
	is.Equal(`"a"`, Json("a"))
	is.Equal("true", Json(true))
	is.Equal("", Json(nil))
	is.Equal(`{"key":"age","value":18}`, Json(struct {
		Key   string `json:"key"`
		Value int    `json:"value"`
	}{Key: "age", Value: 18}))
}

func TestDeepCopy(t *testing.T) {
	is := assert.New(t)
	sourceMap := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}
	targetMap := &map[string]int{}
	DeepCopy(sourceMap, targetMap)
	is.Equal(Json(targetMap), Json(sourceMap))

	type testStruct struct {
		Key   string `json:"key"`
		Value int    `json:"value"`
	}
	source := &testStruct{
		Key:   "abc",
		Value: 123,
	}
	target := &testStruct{}
	DeepCopy(source, target)
	is.Equal(Json(source), Json(target))
}
