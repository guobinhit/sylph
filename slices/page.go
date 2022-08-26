package slices

import (
	"fmt"
	"reflect"
)

// Page slice and returns the current page results and hasNext mark.
// Param of arr must be a slice and returns nil slice when: (pageNum - 1) * pageLimit > len(arr).
// Panics when:
//  1. arr type is unexpected type,
//  2. param: pageNum <= 0 or pageLimit <= 0.
func Page(arr interface{}, pageNum, pageLimit int) (interface{}, bool) {
	// Check param and type of arr
	if pageNum <= 0 {
		panic(fmt.Sprintf("pageNum invalid: %d", pageNum))
	}
	if pageLimit <= 0 {
		panic(fmt.Sprintf("pageLimit invalid: %d", pageLimit))
	}
	arrValue := reflect.ValueOf(arr)
	arrKind := arrValue.Kind()
	if arrKind != reflect.Slice {
		panic(fmt.Sprintf("arr not slice: kind=%s", arrKind))
	}
	arrLen := arrValue.Len()
	if arrLen == 0 {
		return reflect.Zero(arrValue.Type()).Interface(), false
	}
	offset := (pageNum - 1) * pageLimit
	if offset >= arrLen {
		return reflect.Zero(arrValue.Type()).Interface(), false
	}

	// Calculate page logic
	var hasNext bool
	offsetEnd := offset + pageLimit
	if offsetEnd < arrLen {
		hasNext = true
	} else if offsetEnd > arrLen {
		offsetEnd = arrLen
	}
	sliceLen := offsetEnd - offset
	retArr := reflect.MakeSlice(arrValue.Type(), sliceLen, sliceLen)
	reflect.Copy(retArr, arrValue.Slice(offset, offsetEnd))
	return retArr.Interface(), hasNext
}
