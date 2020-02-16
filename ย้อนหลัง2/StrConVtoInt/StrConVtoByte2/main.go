package main

import "fmt"

func main() {
	str := "abc"
	mySlice := []byte(str)
	fmt.Printf("%v -> '%s'", mySlice, mySlice)
}
