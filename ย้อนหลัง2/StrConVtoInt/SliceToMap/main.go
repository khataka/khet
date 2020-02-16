package main

import "fmt"

func main() {
	elements := []string{"abc", "def", "fgi", "adi"}
	elementMap := make(map[string]string)
	for i := 0; i < len(elements); i += 2 {
		elementMap[elements[i]] = elements[i+1]
	}
	fmt.Println(elementMap)
}
