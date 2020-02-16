package main

import (
	"fmt"
)

type Before struct {
	m map[string]string
}

type After struct {
	Before
	s []string
}

func contrivedAfter(b interface{}) interface{} {
	return After{b.(Before), []string{"new value"}}
}

func main() {
	b := Before{map[string]string{"some": "value"}}
	a := contrivedAfter(b).(After)
	fmt.Println(a.m)
	fmt.Println(a.s)
}
