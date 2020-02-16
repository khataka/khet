package main

import (
	"fmt"
)

func weirdFunc(i int) interface{} {

	switch v := myInterface.(type) {
	case int:

		fmt.Printf("Integer: %v", v)
	case float64:

		fmt.Printf("Float64: %v", v)
	case string:

		fmt.Printf("String: %v", v)
	default:

		fmt.Printf("I don't know, ask stackoverflow.")
	}

}
