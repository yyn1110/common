package common

import (
	"fmt"
	"reflect"
)

// only support original type and its array or map or slice
func Compare(a, b interface{}, sequence bool) bool {
	if reflect.TypeOf(a) != reflect.TypeOf(b) {
		// fmt.Println("not the same type can not compare", reflect.TypeOf(a), reflect.TypeOf(b))
		return false
	}
	// compare nil
	if a == nil || b == nil {
		if a == nil && b == nil {
			return true
		}
		return false
	}
	return compare(reflect.ValueOf(a), reflect.ValueOf(b), sequence)
}

// TODO all int derived type can compare with each other
func compare(x, y reflect.Value, sequence bool) bool {
	if x.Kind() != y.Kind() {
		// fmt.Println("not the same kind can not compare", x.Kind(), y.Kind())
		return false
	}
	switch x.Kind() {
	// basic duck type
	case reflect.Bool:
		return x.Bool() == y.Bool()
	case reflect.String:
		return x.String() == y.String()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return x.Int() == y.Int()
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return x.Uint() == y.Uint()
	// container type
	case reflect.Struct:
		if x.NumField() != y.NumField() {
			return false
		}
		n := x.NumField()
		for i := 0; i < n; i++ {
			if !compare(x.Field(i), y.Field(i), sequence) {
				return false
			}
		}
		return true
	case reflect.Array, reflect.Slice:
		if x.Len() != y.Len() {
			return false
		}
		n := x.Len()
		if sequence {
			for i := 0; i < n; i++ {
				if !compare(x.Index(i), y.Index(i), true) {
					return false
				}
			}
		} else {
			for i := 0; i < n; i++ {
				find := false
				for j := 0; j < n; j++ {
					if compare(x.Index(i), y.Index(j), false) {
						find = true
						break
					}
				}
				if !find {
					return false
				}
			}
		}
		return true
	case reflect.Map:
		if x.Len() != y.Len() {
			return false
		}
		keys := x.MapKeys()
		for _, k := range keys {
			v1 := x.MapIndex(k)
			v2 := y.MapIndex(k)
			if !v2.IsValid() || !compare(v1, v2, sequence) {
				return false
			}
		}
		return true
	default:
		panic(fmt.Sprintf("this type[%s] not supported", x.Kind()))
	}
}
