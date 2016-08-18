package main

import (
	"fmt"
)

func main() {
	defer func() {
		if p := recover(); p != nil {
			fmt.Println(p)
		}
	}()
	NoReturn()
}

func NoReturn() {
	panic("I'm a returned string.")
}
