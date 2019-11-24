package main

import "fmt"

func main() {
	n, e := fmt.Println("hello, world", 123, 456, "\n")
	fmt.Println("number of bite written :", n, "\n")
	fmt.Println("write error encounter :", e, "\n")

}
