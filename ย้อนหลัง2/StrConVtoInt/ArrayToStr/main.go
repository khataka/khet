package main

import (
	"fmt"
	"reflect"
)

func main() {
	stringArray := []string{"Hello", "world", "!"}
	justString := fmt.Sprint(stringArray)

	fmt.Println("value \t=", justString, "\ntype \t=", reflect.TypeOf(justString))
}
