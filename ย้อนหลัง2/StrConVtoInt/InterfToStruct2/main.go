package main

import (
	"fmt"
)

type Before struct {
	m map[string]string
}

func contrivedAfter(b interface{}) interface{} {
	return struct {
		Before
		s []string
	}{b.(Before), []string{"new value"}}
}

func main() {
	b := Before{map[string]string{"some": "value"}}
	a := contrivedAfter(b)
	fmt.Println(a)
}
