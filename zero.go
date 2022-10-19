package reflect_utils

import "reflect"

// IsZero 是否是零值，如果是nil的也认为是零值
// 零值包括： 空指针、大小为0的Map、Chan、Slice、Array
func IsZero(value any) bool {
	if IsNil(value) {
		return true
	}
	reflectValue := reflect.ValueOf(value)
	switch reflectValue.Kind() {
	case reflect.Array, reflect.Map, reflect.Slice, reflect.String:
		return reflectValue.Len() == 0
	case reflect.Bool:
		return !reflectValue.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return reflectValue.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return reflectValue.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return reflectValue.Float() == 0
	case reflect.Interface, reflect.Ptr:
		if reflectValue.IsNil() {
			return true
		}
		return IsZero(reflectValue.Elem())
	case reflect.Func:
		return reflectValue.IsNil()
	case reflect.Invalid:
		return true
	}
	return false
}

// IsNotZero 是否是非零，只是IsZero的简单取反
func IsNotZero(v any) bool {
	return !IsZero(v)
}
