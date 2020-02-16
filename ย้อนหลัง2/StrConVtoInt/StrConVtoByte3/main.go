package main

import (
	"fmt"
	"strings"
)

func main() {

	stringSlice := []string{"hello", "world"}

	stringByte := strings.Join(stringSlice, " ")

	fmt.Println([]byte(stringByte))

	fmt.Println(string([]byte(stringByte)))
}
