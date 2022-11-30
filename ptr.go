package reflect_utils

import (
	"fmt"
	"reflect"
)

// GetUintPtrE 获取uintptr，不必担心panic的问题
func GetUintPtrE(value any) (uintptr, error) {
	reflectValue := reflect.ValueOf(value)
	switch reflectValue.Kind() {
	case reflect.Chan, reflect.Func, reflect.Map, reflect.Pointer, reflect.Slice, reflect.UnsafePointer:
		return reflectValue.Pointer(), nil
	default:
		return 0, fmt.Errorf("not support operation take pointer for type %#v", value)
	}
}

// GetUintPtr 获取uintptr，不必担心panic的问题
func GetUintPtr(value any) uintptr {
	ptr, _ := GetUintPtrE(value)
	return ptr
}


