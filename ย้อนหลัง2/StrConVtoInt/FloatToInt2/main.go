package main

import (
	"fmt"
	"strconv"
)

func main() {
	floats := []float64{1.9999, 2.0001, 2.0}
	for _, f := range floats {
		t := int(f)
		s := fmt.Sprintf("%.0f", f)
		if i, err := strconv.Atoi(s); err == nil {
			fmt.Println(f, t, i)
		} else {
			fmt.Println(f, t, err)
		}
	}
}
