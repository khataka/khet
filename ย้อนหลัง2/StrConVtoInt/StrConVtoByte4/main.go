package main

import (
	"fmt"
)

func byte16PutString(s string) [16]byte {
	var a [16]byte
	if len(s) > 16 {
		copy(a[:], s)
	} else {
		copy(a[16-len(s):], s)
	}
	return a
}

func main() {
	var a [16]byte
	a = byte16PutString("abcaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")
	fmt.Printf("%v\n", a)
	var b [16]byte
	b = byte16PutString("abc")
	fmt.Printf("%v\n", b)
}
