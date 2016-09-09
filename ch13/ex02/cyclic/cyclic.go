package cyclic

import (
	"reflect"
	"unsafe"
)

// IsCyclic return true if i has cyclic structure.
func IsCyclic(i interface{}) bool {
	var seen []unsafe.Pointer
	return isCyclic(reflect.ValueOf(i), seen)
}

func isCyclic(v reflect.Value, seen []unsafe.Pointer) bool {
	if v.CanAddr() {
		ptr := unsafe.Pointer(v.UnsafeAddr())
		for _, p := range seen {
			if p == ptr {
				return true
			}
		}
		seen = append(seen, ptr)
	}

	if v.Kind() == reflect.Slice || v.Kind() == reflect.Map {
		ptr := unsafe.Pointer(v.Pointer())
		for _, p := range seen {
			if p == ptr {
				return true
			}
		}
		seen = append(seen, ptr)
	}

	switch v.Kind() {
	case reflect.Ptr, reflect.Interface:
		if isCyclic(v.Elem(), seen) {
			return true
		}

	case reflect.Array, reflect.Slice:
		for i := 0; i < v.Len(); i++ {
			if isCyclic(v.Index(i), seen) {
				return true
			}
		}

	case reflect.Map:
		for _, key := range v.MapKeys() {
			if isCyclic(v.MapIndex(key), seen) {
				return true
			}
		}

	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			if isCyclic(v.Field(i), seen) {
				return true
			}
		}
	}

	return false
}
