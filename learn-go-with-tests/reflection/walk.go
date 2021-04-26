package reflection

import "reflect"

// walk through struct,map,slice,array,pointer,string
func Walk(x interface{}, f func(string)) {
	value := getValue(x)

	walkValue := func(val reflect.Value) {
		Walk(val.Interface(), f)
	}

	switch value.Kind() {
	case reflect.String:
		f(value.String())
	case reflect.Struct:
		for i := 0; i < value.NumField(); i++ {
			walkValue(value.Field(i))
		}
	case reflect.Slice, reflect.Array:
		for i := 0; i < value.Len(); i++ {
			walkValue(value.Index(i))
		}
	case reflect.Map:
		for _, key := range value.MapKeys() {
			walkValue(value.MapIndex(key))
		}
	}
}

// get value from the pointer
func getValue(x interface{}) reflect.Value {
	val := reflect.ValueOf(x)

	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	return val
}
