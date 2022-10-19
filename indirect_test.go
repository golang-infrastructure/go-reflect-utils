package reflect_utils

//import (
//	"reflect"
//	"testing"
//)
//
//func TestVisitIndirect(t *testing.T) {
//	var t1 any = "这是一个字符串"
//	for i := 0; i < 10; i++ {
//		t1 = &t1
//	}
//	value := reflect.ValueOf(t1)
//	indirect := VisitIndirect(value)
//	t.Log(indirect)
//	t.Log(indirect.String())
//}
