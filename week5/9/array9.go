package main

import "fmt"

func main() {
	alphabets := [3][4]string{{"a", "b", "c", "d"}, {"A", "B", "C", "D"}, {"G", "H", "I", "J"}}
	fmt.Println(alphabets)
	fmt.Println(alphabets[1][3])

}
