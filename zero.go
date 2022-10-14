package reflect_utils

import "reflect"

// IsZero 是否是零值
// 零值包括： 空指针、大小为0的Map、Chan、Slice、Array
func IsZero(value any) bool {
	if IsNil(value) {
		return true
	}
	reflectValue := reflect.ValueOf(value)
	switch reflectValue.Type().Kind() {
	case reflect.Map, reflect.Chan, reflect.Slice, reflect.Array:
		return reflectValue.Len() == 0
	case reflect.Ptr:
		return reflectValue.IsZero()
	default:
		return false
	}
}

func IsNotZero(v any) bool {

	return !IsZero(v)
}
