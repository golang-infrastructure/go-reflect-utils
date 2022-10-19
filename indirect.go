package reflect_utils

//func Indirect() bool {
//	return false
//}

//// Direct 如果是指针类型的，一直不断访问直到拿到非指针类型
//func Direct(t reflect.Type) reflect.Type {
//
//}

//func VisitIndirect(value reflect.Value) reflect.Value {
//	if value.Kind() == reflect.Ptr || value.Kind() == reflect.Interface {
//		return VisitIndirect(value.Elem())
//	} else {
//		return value
//	}
//}
