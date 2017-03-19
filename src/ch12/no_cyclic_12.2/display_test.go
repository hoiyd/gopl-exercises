package display

import (
	"io"
	"net"
	"os"
	"reflect"
	"sync"
	"testing"
)

func Example_slice() {
	Display("slice", []*int{new(int), nil})
	// Output:
	// Display slice ([]*int):
	// (*slice[0]) = 0
	// slice[1] = nil
}

func Example_nilInterface() {
	var w io.Writer
	Display("w", w)
	// Output:
	// Display w (<nil>):
	// w = invalid
}

func Example_ptrToInterface() {
	var w io.Writer
	Display("&w", &w)
	// Output:
	// Display &w (*io.Writer):
	// (*&w) = nil
}

func Example_struct() {
	Display("x", struct{ x interface{} }{3})
	// Output:
	// Display x (struct { x interface {} }):
	// x.x.type = int
	// x.x.value = 3
}

func Example_ptrToInt() {
	var i int = 3
	Display("&i", &i)
	// Output:
	// Display &i (*int):
	// (*&i) = 3
}

func Example_interface() {
	var i interface{} = 3
	Display("i", i)
	// Output:
	// Display i (int):
	// i = 3
}

func Example_ptrToInterface2() {
	var i interface{} = 3
	Display("&i", &i)
	// Output:
	// Display &i (*interface {}):
	// (*&i).type = int
	// (*&i).value = 3
}

func Example_array() {
	Display("x", [1]interface{}{3})
	// Output:
	// Display x ([1]interface {}):
	// x[0].type = int
	// x[0].value = 3
}

func Example_movie() {
	//!+movie
	type Movie struct {
		Title, Subtitle string
		Year            int
		Color           bool
		Actor           map[string]string
		Oscars          []string
		Sequel          *string
	}
	//!-movie
	//!+strangelove
	strangelove := Movie{
		Title:    "Dr. Strangelove",
		Subtitle: "How I Learned to Stop Worrying and Love the Bomb",
		Year:     1964,
		Color:    false,
		Actor: map[string]string{
			"Dr. Strangelove":            "Peter Sellers",
			"Grp. Capt. Lionel Mandrake": "Peter Sellers",
			"Pres. Merkin Muffley":       "Peter Sellers",
			"Gen. Buck Turgidson":        "George C. Scott",
			"Brig. Gen. Jack D. Ripper":  "Sterling Hayden",
			`Maj. T.J. "King" Kong`:      "Slim Pickens",
		},

		Oscars: []string{
			"Best Actor (Nomin.)",
			"Best Adapted Screenplay (Nomin.)",
			"Best Director (Nomin.)",
			"Best Picture (Nomin.)",
		},
	}
	//!-strangelove
	Display("strangelove", strangelove)

	// We don't use an Output: comment since displaying
	// a map is nondeterministic.
	/*
	   //!+output
	   Display strangelove (display.Movie):
	   strangelove.Title = "Dr. Strangelove"
	   strangelove.Subtitle = "How I Learned to Stop Worrying and Love the Bomb"
	   strangelove.Year = 1964
	   strangelove.Color = false
	   strangelove.Actor["Gen. Buck Turgidson"] = "George C. Scott"
	   strangelove.Actor["Brig. Gen. Jack D. Ripper"] = "Sterling Hayden"
	   strangelove.Actor["Maj. T.J. \"King\" Kong"] = "Slim Pickens"
	   strangelove.Actor["Dr. Strangelove"] = "Peter Sellers"
	   strangelove.Actor["Grp. Capt. Lionel Mandrake"] = "Peter Sellers"
	   strangelove.Actor["Pres. Merkin Muffley"] = "Peter Sellers"
	   strangelove.Oscars[0] = "Best Actor (Nomin.)"
	   strangelove.Oscars[1] = "Best Adapted Screenplay (Nomin.)"
	   strangelove.Oscars[2] = "Best Director (Nomin.)"
	   strangelove.Oscars[3] = "Best Picture (Nomin.)"
	   strangelove.Sequel = nil
	   //!-output
	*/
}

// This test ensures that the program terminates without crashing.
func Test(t *testing.T) {
	// Some other values (YMMV)
	Display("os.Stderr", os.Stderr)
	// Output:
	// Display os.Stderr (*os.File):
	// (*(*os.Stderr).file).fd = 2
	// (*(*os.Stderr).file).name = "/dev/stderr"
	// (*(*os.Stderr).file).nepipe = 0

	var w io.Writer = os.Stderr
	Display("&w", &w)
	// Output:
	// Display &w (*io.Writer):
	// (*&w).type = *os.File
	// (*(*(*&w).value).file).fd = 2
	// (*(*(*&w).value).file).name = "/dev/stderr"
	// (*(*(*&w).value).file).nepipe = 0

	var locker sync.Locker = new(sync.Mutex)
	Display("(&locker)", &locker)
	// Output:
	// Display (&locker) (*sync.Locker):
	// (*(&locker)).type = *sync.Mutex
	// (*(*(&locker)).value).state = 0
	// (*(*(&locker)).value).sema = 0

	Display("locker", locker)
	// Output:
	// Display locker (*sync.Mutex):
	// (*locker).state = 0
	// (*locker).sema = 0
	// (*(&locker)) = nil

	locker = nil
	Display("(&locker)", &locker)
	// Output:
	// Display (&locker) (*sync.Locker):
	// (*(&locker)) = nil

	ips, _ := net.LookupHost("golang.org")
	Display("ips", ips)
	// Output:
	// Display ips ([]string):
	// ips[0] = "173.194.68.141"
	// ips[1] = "2607:f8b0:400d:c06::8d"

	// Even metarecursion!  (YMMV)
	Display("rV", reflect.ValueOf(os.Stderr))
	// Output:
	// Display rV (reflect.Value):
	// (*rV.typ).size = 8
	// (*rV.typ).ptrdata = 8
	// (*rV.typ).hash = 871609668
	// (*rV.typ)._ = 0
	// ...

	// a pointer that points to itself
	type P *P
	var p P
	p = &p
	Display("p", p)
	// Output:
	// Display p (display.P):
	// (*(*(*(*p)))) = main.P 0xc42002a020

	// a map that contains itself
	type M map[string]M
	m := make(M)
	m[""] = m
	// if false {
	Display("m", m)
	// Output:
	// Display m (display.M):
	// m[""][""][""][""] = main.M 0xc4200162d0

	// a slice that contains itself
	type S []S
	s := make(S, 1)
	s[0] = s
	Display("s", s)
	// Output:
	// Display s (display.S):
	// s[0][0][0][0] = main.S 0xc420012560

	// a linked list that eats its own tail
	type Cycle struct {
		Value int
		Tail  *Cycle
	}
	var c Cycle
	c = Cycle{42, &c}
	Display("c", c)
	// Output:
	// Display c (display.Cycle):
	// c.Value = 42
	// (*c.Tail).Value = 42
	// (*(*c.Tail).Tail) = main.Cycle value
}

func TestMapKeys(t *testing.T) {
	sm := map[struct{ x int }]int{
		{1}: 2,
		{2}: 3,
	}
	Display("sm", sm)
	// Output:
	// Display sm (map[struct { x int }]int):
	// sm[{x: 2}] = 3
	// sm[{x: 1}] = 2

	am := map[[3]int]int{
		{1, 2, 3}: 3,
		{2, 3, 4}: 4,
	}
	Display("am", am)
	// Output:
	// Display am (map[[3]int]int):
	// am[{1, 2, 3}] = 3
	// am[{2, 3, 4}] = 4

	nestedStructMapKey := map[struct {
		x int
		y struct{ z int }
	}]int{
		{y: struct{ z int }{z: 1}, x: 2}: 2,
		{y: struct{ z int }{z: 2}, x: 3}: 3,
	}
	Display("nestedStructMapKey", nestedStructMapKey)
	// Output:
	// Display nestedStructMapKey (map[struct { x int }]int):
	// nestedStructMapKey[{x: 2, y: {z: 1}}] = 2
	// nestedStructMapKey[{x: 3, y: {z: 2}}] = 3

	nestedArrayMapKey := map[[3][3]int]int{
		{{1, 2, 3}}: 3,
		{{2, 3, 4}}: 4,
	}
	Display("nestedArrayMapKey", nestedArrayMapKey)
	// Output:
	// Display nestedArrayMapKey (map[[3]int]int):
	// nestedArrayMapKey[{1, 2, 3}, {0, 0, 0}, {0, 0, 0}] = 3
	// nestedArrayMapKey[{2, 3, 4} ,{0, 0, 0}, {0, 0, 0}] = 4
}
