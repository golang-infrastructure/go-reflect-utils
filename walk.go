package reflect_util

//func Walk(v any, visitor func(key, value reflect.Value) bool) {
//	reflectValue := reflect.ValueOf(v)
//	reflectType := reflect.TypeOf(v)
//	switch reflectType.Kind() {
//	case reflect.Struct:
//		for i := 0; i < reflectType.NumField(); i++ {
//			v := reflectValue.Field(i)
//			if !visitor(v, nil) {
//				return
//			}
//		}
//	case reflect.Slice, reflect.Array:
//		for i := 0; i < reflectValue.Len(); i++ {
//			v := reflectValue.Index(i)
//			if !visitor(v, nil) {
//				return
//			}
//		}
//	case reflect.Map:
//		keys := reflectValue.MapKeys()
//		for i := range keys {
//			// TODO
//		}
//	default:
//		visitor(reflectValue, nil)
//	}
//}
