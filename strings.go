package srutil

import (
	"reflect"
	"unsafe"
)

// DeduplicateArray : depulicate a slice
func DeduplicateArray(s []string) []string {
	res := NewSet(s...)
	return res.Slice()
}

// ContainsString : return first index of val in array, if val is not in array return -1
func ContainsString(array []string, val string) int {
	for idx, item := range array {
		if item == val {
			return idx
		}
	}
	return -1
}

func Str2Bytes(s string) (b []byte) {
	h1 := (*reflect.StringHeader)(unsafe.Pointer(&s))
	h2 := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	h2.Cap = h1.Len
	h2.Len = h1.Len
	h2.Data = h1.Data
	return
}

func Bytes2Str(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
