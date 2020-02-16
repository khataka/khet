package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var kpi interface{} = st
	var a []Animal
	err := json.Unmarshal([]byte(kpi.(string)), &a)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(a)
}

type Animal struct {
	Name  string
	Order string
}

var st = `[
    {"Name": "Platypus", "Order": "Monotremata"},
    {"Name": "Quoll",    "Order": "Dasyuromorphia"}
]`
