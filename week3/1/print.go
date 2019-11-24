package main

import "fmt"

func main() {
	n, e := fmt.Print("hello, world", 123, 456, "\n")
	fmt.Print("number of bite written :", n, "\n")
	fmt.Print("write error encounter :", e, "\n")

}
