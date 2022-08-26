package maps

import (
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

func TestKeys(t *testing.T) {
	is := assert.New(t)

	var mInt64Str map[int]string
	intKeys := Keys(mInt64Str).([]int)
	is.Empty(intKeys)

	mInt64Str = map[int]string{}
	intKeys = Keys(mInt64Str).([]int)
	is.Empty(intKeys)

	mInt64Str = map[int]string{
		1: "1",
		2: "2",
		3: "3",
	}
	intKeys = Keys(mInt64Str).([]int)
	sort.Ints(intKeys)
	is.Equal([]int{1, 2, 3}, intKeys)

	mStrStruct := map[string]struct{}{
		"a": {},
		"b": {},
		"c": {},
	}
	strKeys := Keys(mStrStruct).([]string)
	sort.Strings(strKeys)
	is.Equal([]string{"a", "b", "c"}, strKeys)
}

func TestValues(t *testing.T) {
	is := assert.New(t)

	var mInt64Str map[int]string
	strValues := Values(mInt64Str).([]string)
	is.Empty(strValues)

	mInt64Str = map[int]string{}
	strValues = Values(mInt64Str).([]string)
	is.Empty(strValues)

	mInt64Str = map[int]string{
		1: "1",
		2: "2",
		3: "3",
	}
	strValues = Values(mInt64Str).([]string)
	sort.Strings(strValues)
	is.Equal([]string{"1", "2", "3"}, strValues)

	mStrStruct := map[string]struct{ ID int }{
		"3": {3},
		"2": {2},
		"1": {1},
	}
	models := Values(mStrStruct).([]struct{ ID int })
	modelIDs := make([]int, 0, len(models))
	for _, model := range models {
		modelIDs = append(modelIDs, model.ID)
	}
	sort.Ints(modelIDs)
	is.Equal([]int{1, 2, 3}, modelIDs)
}
