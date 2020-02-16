package main

import (
	"fmt"
	"strconv"
)

func FloatToString(input_num float64) string {

	return strconv.FormatFloat(input_num, 'f', 6, 64)
}

func main() {
	fmt.Println(FloatToString(21312421.213123))
}
