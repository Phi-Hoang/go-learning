package lodash

import (
	"fmt"
	"reflect"
)

// Reverse1 function to reverse elements in a slice/an array
func Reverse1(s []interface{}) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

//panic if s is not a slice
func reverseSlice(s interface{}) {
	size := reflect.ValueOf(s).Len()
	swap := reflect.Swapper(s)
	for i, j := 0, size-1; i < j; i, j = i+1, j-1 {
		swap(i, j)
	}
}

// Reverse function to reverse elements in a slice/an array
func Reverse(s interface{}) {
	switch v := reflect.ValueOf(s); v.Kind() {
	case reflect.String:
		runes := []rune(s.(string)) // The rune type is an alias for int32, and is used to emphasize than an integer represents a code point. A string is a sequence of bytes, not runes.
		reverseSlice(runes)
		s = string(runes)
		fmt.Println(s)
	case reflect.Slice:
		reverseSlice(s)
	default:
		fmt.Printf("unsupported kind %s", v.Kind())
	}
}
