package reflect_utils
//
//import "reflect"
//
//// 解析tag
//
//const DefaultTagPattern = "ItemName:key=value,"
//
//// StructFieldTag 表示解析到的一个Struct的Field的标签信息
//type StructFieldTag struct {
//	FieldType       reflect.Type
//	FieldValue      reflect.Value
//	FieldTagItemMap map[string]TagItem
//}
//
//type TagItem struct {
//	ItemName           string
//	ItemValue          string
//	SubItemKeyValueMap map[string]string
//}
//
//func (x *TagItem) Contains(name string) bool {
//	_, exists := x.SubItemKeyValueMap[name]
//	return exists
//}
//
//func (x *TagItem) Get(name string) string {
//
//}
//
//func (x *TagItem) GetAsBool(name string) bool{
//
//}
//
//func (x *TagItem) GetAsInt() {
//
//}
//
//func (x *TagItem) GetAsFloat64() {
//
//}
//
//func (x *TagItem) GetAsFloat32() {
//
//}
//
//func (x *TagItem) GetOrDefault() {
//
//}
//
//func (x *TagItem) Lookup() {
//
//}
//
//// ParseTag 解析Struct字段上的TAG
//func ParseStructTag(structValue any) {
//	reflectType := reflect.TypeOf(structValue)
//	reflectType.Field(0).Tag
//	reflectValue := reflect.ValueOf(structValue)
//
//}
//
//func ParseTag(tagStr string) {
//
//}
