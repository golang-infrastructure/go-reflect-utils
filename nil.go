package reflect_utils

import "reflect"

// IsNil 判断参数是否为nil，当只有type但是value为nil的时候会认为是nil
func IsNil(v any) bool {
	if v == nil {
		return true
	}
	reflectValue := reflect.ValueOf(v)
	kind := reflectValue.Kind()
	if kind >= reflect.Chan && kind <= reflect.Slice && reflectValue.IsNil() {
		return true
	}
	if reflectValue.Kind() == reflect.Ptr {
		return reflectValue.IsNil()
	}
	return false
}

// IsNotNil 判断参数是否不为nil
func IsNotNil(v any) bool {
	return !IsNil(v)
}
