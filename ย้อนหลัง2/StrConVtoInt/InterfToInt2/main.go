package main

import "fmt"

func main() {
	var val interface{}
	val = 4

	var i int

	switch t := val.(type) {
	case int:
		fmt.Printf("%d == %T\n", t, t)
		i = t
	case int8:
		fmt.Printf("%d == %T\n", t, t)
		i = int(t)
	case int16:
		fmt.Printf("%d == %T\n", t, t)
		i = int(t)
	case int32:
		fmt.Printf("%d == %T\n", t, t)
		i = int(t)
	case int64:
		fmt.Printf("%d == %T\n", t, t)
		i = int(t)
	case bool:
		fmt.Printf("%t == %T\n", t, t)

	case float32:
		fmt.Printf("%g == %T\n", t, t)
		i = int(t)
	case float64:
		fmt.Printf("%f == %T\n", t, t)
		i = int(t)
	case uint8:
		fmt.Printf("%d == %T\n", t, t)
		i = int(t)
	case uint16:
		fmt.Printf("%d == %T\n", t, t)
		i = int(t)
	case uint32:
		fmt.Printf("%d == %T\n", t, t)
		i = int(t)
	case uint64:
		fmt.Printf("%d == %T\n", t, t)
		i = int(t)
	case string:
		fmt.Printf("%s == %T\n", t, t)

	default:

		fmt.Printf("%v == %T\n", t, t)
	}

	fmt.Printf("i == %d\n", i)
}
