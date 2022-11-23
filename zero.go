package reflect_utils

import (
	"math"
	"reflect"
)

// IsZero 是否是零值，如果是nil的也认为是零值
// 零值包括： 空指针、大小为0的Map、Chan、Slice、Array
func IsZero(value any) bool {
	if IsNil(value) {
		return true
	}
	reflectValue := reflect.ValueOf(value)
	switch reflectValue.Kind() {
	case reflect.Bool:
		return !reflectValue.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return reflectValue.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return reflectValue.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return math.Float64bits(reflectValue.Float()) == 0
	case reflect.Complex64, reflect.Complex128:
		c := reflectValue.Complex()
		return math.Float64bits(real(c)) == 0 && math.Float64bits(imag(c)) == 0
	case reflect.Array:
		for i := 0; i < reflectValue.Len(); i++ {
			if !reflectValue.Index(i).IsZero() {
				return false
			}
		}
		return true
	case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Pointer, reflect.Slice, reflect.UnsafePointer:
		return reflectValue.IsNil()
	case reflect.String:
		return reflectValue.Len() == 0
	case reflect.Struct:
		for i := 0; i < reflectValue.NumField(); i++ {
			if !reflectValue.Field(i).IsZero() {
				return false
			}
		}
		return true
	default:
		return false
	}
}

// IsNotZero 是否是非零，只是IsZero的简单取反
func IsNotZero(v any) bool {
	return !IsZero(v)
}
