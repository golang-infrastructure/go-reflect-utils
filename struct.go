package reflect_utils

//
//import (
//	"fmt"
//	"reflect"
//	"unsafe"
//)
//
//// GetStructPtrUnExportedField 获取struct上给定字段的值
//func GetStructPtrUnExportedField(source any, fieldName string) reflect.ItemValue {
//	v := reflect.ValueOf(source).Elem().FieldByName(fieldName)
//	if !v.IsValid() {
//		return v
//	}
//	return reflect.NewAt(v.Type(), unsafe.Pointer(v.UnsafeAddr())).Elem()
//}
//
//// SetStructPtrUnExportedStrField 设置struct上给定字段的值
//func SetStructPtrUnExportedStrField(source interface{}, fieldName string, fieldVal interface{}) error {
//	v := GetStructPtrUnExportedField(source, fieldName)
//	rv := reflect.ValueOf(fieldVal)
//	if v.Kind() != rv.Kind() {
//		return fmt.Errorf("invalid kind: expected kind %v, got kind: %v", v.Kind(), rv.Kind())
//	}
//	v.Set(rv)
//	return nil
//}
//
//// Merge 合并Struct
//func Merge(sourceStruct, destStruct any) error {
//	// TODO 2022-10-16 10:42:55
//
//	return nil
//}

//// StructToMap 把struct转为map类型，如果struct的field也是struct的话，会递归转换
//func StructToMap(value any) map[string]any {
//}

//func MapToStruct() {
//
//}
//
//// StructFieldNames 获取一个struct的所有字段名
//func StructFieldNames() {
//
//}
//
//// StructValues 获取一个struct的所有字段值
//func StructValues() {
//
//}
//
//// IsStruct 判断是否是struct，如果是指针的话就递归访问拿到真实的值
//func IsStruct() bool {
//
//}

// ------------------------------------------------- ------------------------------------------------------------------------

//func StructToBytes(structObject *any) []byte {
//	s := unsafe.Pointer(structObject)
//	sizeOfMyStruct := int(unsafe.Sizeof(structObject))
//
//	var x reflect.SliceHeader
//	x.Len = sizeOfMyStruct
//	x.Cap = sizeOfMyStruct
//	x.Data = uintptr(s)
//	return *(*[]byte)(unsafe.Pointer(&x))
//}
//
//func BytesToStruct() {
//
//}
