package reflect_util

import (
	"reflect"
)

// LenType 安全的计算Type的长度，如果类型不合适会返回0而不是panic
func LenType(t reflect.Type) int {
	if t.Kind() != reflect.Array {
		return 0
	}
	return t.Len()
}

// LenValue 安全的计算Value的长度，如果类型不合适会返回0而不是panic
func LenValue(v reflect.Value) int {
	switch v.Kind() {
	case reflect.Slice, reflect.Array, reflect.Chan, reflect.Map, reflect.String, reflect.Ptr:
		return v.Len()
	default:
		return 0
	}
}
