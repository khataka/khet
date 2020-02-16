package main

import "fmt"

func weirdFunc(i int) interface{} {
	if i == 0 {
		return "zero"
	}
	return i
}
func main() {
	var i = 5
	var w = weirdFunc(5)

	if tmp, ok := w.(int); ok {
		i += tmp
	}

	fmt.Println("i =", i)
}
