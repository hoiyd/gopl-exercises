// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 359.

// Package equal provides a deep equivalence relation for arbitrary values.
package equal

import (
	"reflect"
	"unsafe"
)

const billion = 1000000000

func deepEqualNumber(x, y float64) bool {
	if x == y {
		return true
	}
	var diff float64
	if x > y {
		diff = x - y
	} else {
		diff = y - x
	}

	delta := diff * billion

	if x > y {
		return delta <= x
	} else {
		return delta <= y
	}

}

//!+
func equal(x, y reflect.Value, seen map[comparison]bool) bool {
	if !x.IsValid() || !y.IsValid() {
		return x.IsValid() == y.IsValid()
	}
	if x.Type() != y.Type() {
		return false
	}

	// ...cycle check omitted (shown later)...

	//!-
	//!+cyclecheck
	// cycle check
	if x.CanAddr() && y.CanAddr() {
		xptr := unsafe.Pointer(x.UnsafeAddr())
		yptr := unsafe.Pointer(y.UnsafeAddr())
		if xptr == yptr {
			return true // identical references
		}
		c := comparison{xptr, yptr, x.Type()}
		if seen[c] {
			return true // already seen
		}
		seen[c] = true
	}
	//!-cyclecheck
	//!+
	switch x.Kind() {
	case reflect.Bool:
		return x.Bool() == y.Bool()

	case reflect.String:
		return x.String() == y.String()

	// ...numeric cases omitted for brevity...

	//!-
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32,
		reflect.Int64:
		return deepEqualNumber(float64(x.Int()), float64(y.Int()))

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32,
		reflect.Uint64, reflect.Uintptr:
		return deepEqualNumber(float64(x.Uint()), float64(y.Uint()))

	case reflect.Float32, reflect.Float64:
		return deepEqualNumber(float64(x.Float()), float64(y.Float()))

	case reflect.Complex64, reflect.Complex128:
		realPart := deepEqualNumber(float64(real(x.Complex())), float64(real(y.Complex())))
		imagPart := deepEqualNumber(float64(imag(x.Complex())), float64(imag(y.Complex())))
		return realPart && imagPart
	//!+
	case reflect.Chan, reflect.UnsafePointer, reflect.Func:
		return x.Pointer() == y.Pointer()

	case reflect.Ptr, reflect.Interface:
		return equal(x.Elem(), y.Elem(), seen)

	case reflect.Array, reflect.Slice:
		if x.Len() != y.Len() {
			return false
		}
		for i := 0; i < x.Len(); i++ {
			if !equal(x.Index(i), y.Index(i), seen) {
				return false
			}
		}
		return true

	// ...struct and map cases omitted for brevity...
	//!-
	case reflect.Struct:
		for i, n := 0, x.NumField(); i < n; i++ {
			if !equal(x.Field(i), y.Field(i), seen) {
				return false
			}
		}
		return true

	case reflect.Map:
		if x.Len() != y.Len() {
			return false
		}
		for _, k := range x.MapKeys() {
			if !equal(x.MapIndex(k), y.MapIndex(k), seen) {
				return false
			}
		}
		return true
		//!+
	}
	panic("unreachable")
}

//!-

//!+comparison
// Equal reports whether x and y are deeply equal.
//!-comparison
//
// Map keys are always compared with ==, not deeply.
// (This matters for keys containing pointers or interfaces.)
//!+comparison
func Equal(x, y interface{}) bool {
	seen := make(map[comparison]bool)
	return equal(reflect.ValueOf(x), reflect.ValueOf(y), seen)
}

type comparison struct {
	x, y unsafe.Pointer
	t    reflect.Type
}

//!-comparison
