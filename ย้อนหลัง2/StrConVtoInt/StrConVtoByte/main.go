package main

import "fmt"

func main() {
	s := "abc"
	var a [20]byte
	copy(a[:], s)
	fmt.Println("s:", []byte(s), "a:", a)
}
