package main

import (
	"fmt"
	"strconv"
)

func main() {
	var v interface{}
	v = "4"

	i, errInt := strconv.ParseInt(v.(string), 10, 64)

	if errInt == nil {
		fmt.Printf("%d is a int", i)

	}
}
