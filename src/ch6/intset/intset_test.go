package intset

import (
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
