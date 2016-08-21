package intset

import (
	"reflect"
	"testing"
)

func TestLen(t *testing.T) {
	var x IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	if x.Len() != 3 {
		t.Error("TestLen() fails!")
	}
}

func TestRemove(t *testing.T) {
	var x IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	x.Add(42)

	if x.String() != "{1 9 42 144}" {
		t.Error("TestRemove() fails!")
	}

	x.Remove(42)

	if x.String() != "{1 9 144}" {
		t.Error("TestRemove() fails!")
	}
}

func TestClear(t *testing.T) {
	var x IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	x.Add(42)

	if x.String() != "{1 9 42 144}" {
		t.Error("TestClear() fails!")
	}

	x.Clear()

	if x.String() != "{}" {
		t.Error("TestClear() fails!")
	}
}

func TestCopy(t *testing.T) {
	var x IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	x.Add(99)
	if x.Len() != 4 || x.String() != "{1 9 99 144}" {
		t.Error("TestCopy() fails!")
	}

	z := x.Copy()
	x.Add(100)

	if x.Len() != 5 || x.String() != "{1 9 99 100 144}" {
		t.Error("TestCopy() fails!")
	}

	if z.Len() != 4 || z.String() != "{1 9 99 144}" {
		t.Error("TestCopy() fails!")
	}
}

func TestAddAll(t *testing.T) {
	var x IntSet
	x.AddAll(1, 2, 3)
	if x.Len() != 3 || x.String() != "{1 2 3}" {
		t.Error("TestAddAll() fails!")
	}

	x.AddAll(5, 15, 24, 4)
	if x.Len() != 7 || x.String() != "{1 2 3 4 5 15 24}" {
		t.Error("TestAddAll() fails!")
	}
}

func TestInsersectWith(t *testing.T) {
	var x, y IntSet
	x.AddAll(1, 2, 3, 4)
	y.AddAll(3, 4, 5)
	x.IntersectWith(&y)
	if x.Len() != 2 || x.String() != "{3 4}" {
		t.Error("TestInsersectWith() fails!")
	}
}

func TestDifferenceWith(t *testing.T) {
	var x, y IntSet
	x.AddAll(1, 2, 3, 4)
	y.AddAll(3, 4, 5)
	x.DifferenceWith(&y)
	if x.Len() != 2 || x.String() != "{1 2}" {
		t.Error("TestDifferenceWith() fails!")
	}
}

func TestSymmetricDifference(t *testing.T) {
	var x, y IntSet
	x.AddAll(1, 2, 3, 4)
	y.AddAll(3, 4, 5)
	x.SymmetricDifference(&y)
	if x.Len() != 3 || x.String() != "{1 2 5}" {
		t.Error("TestSymmetricDifference() fails!")
	}
}

func TestElems(t *testing.T) {
	var x IntSet
	x.AddAll(1, 2, 3, 4)
	xElems := x.Elems()
	ints := []int{1, 2, 3, 4}
	if !reflect.DeepEqual(xElems, ints) {
		t.Error("TestElems() fails!")
	}
}
