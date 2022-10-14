package reflect_utils

//func TestWalk(t *testing.T) {
//
//	// case 001. struct
//	Walk(*test.NewFoo("bar1", "bar2"), func(v reflect.Value) bool {
//		t.Log(v.String())
//		return true
//	})
//	t.Log("--------------------------------------------------------------------------------------------")
//
//	// case 002. array
//	var array [2]int
//	array[0] = 10086
//	array[1] = 10010
//	Walk(array, func(v reflect.Value) bool {
//		t.Log(v.Int())
//		return true
//	})
//	t.Log("--------------------------------------------------------------------------------------------")
//
//	// case 003. slice
//	slice := []int{1, 2, 3, 4, 5}
//	Walk(slice, func(v reflect.Value) bool {
//		t.Log(v.Int())
//		return true
//	})
//	t.Log("--------------------------------------------------------------------------------------------")
//
//}
