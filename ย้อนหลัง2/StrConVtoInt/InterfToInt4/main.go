package main

import "fmt"

func foo(a interface{}) {
	fmt.Println(a.(int))
}
func main() {
	var a int = 10
	foo(a)
}
