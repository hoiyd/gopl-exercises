package cycle

import (
	"reflect"
	"unsafe"
)

func hasCycle(x reflect.Value, seen map[cycleDetection]bool) bool {
	if x.CanAddr() {
		xptr := unsafe.Pointer(x.UnsafeAddr())

		c := cycleDetection{xptr, x.Type()}
		if seen[c] {
			return true // already seen
		}
		seen[c] = true
	}

	switch x.Kind() {
	case reflect.Ptr, reflect.Interface:
		return hasCycle(x.Elem(), seen)

	case reflect.Array, reflect.Slice:
		for i := 0; i < x.Len(); i++ {
			if hasCycle(x.Index(i), seen) {
				return true
			}
		}
		return false

	case reflect.Struct:
		for i, n := 0, x.NumField(); i < n; i++ {
			if hasCycle(x.Field(i), seen) {
				return true
			}
		}
		return false

	case reflect.Map:
		for _, k := range x.MapKeys() {
			// A cycle can probably lurk both in a key and a value.
			if hasCycle(x.MapIndex(k), seen) || hasCycle(k, seen) {
				return true
			}
		}
		return false

	default:
		return false
	}

	panic("unreachable")
}

func HasCycle(x interface{}) bool {
	seen := make(map[cycleDetection]bool)
	return hasCycle(reflect.ValueOf(x), seen)
}

type cycleDetection struct {
	x unsafe.Pointer
	t reflect.Type
}
