package main

import "fmt"

func main() {
	n, e := fmt.Printf("hello, world", 123, 456, "\n")
	fmt.Printf("number of bite written :", n, "\n")
	fmt.Printf("write error encounter :", e, "\n")

}
