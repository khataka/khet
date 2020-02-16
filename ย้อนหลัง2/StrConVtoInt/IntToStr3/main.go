package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	i := strconv.Itoa(123)
	fmt.Println(i)
	fmt.Println(reflect.TypeOf(i))
}
