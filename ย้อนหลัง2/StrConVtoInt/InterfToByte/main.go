package main

import (
	"fmt"
)

func passInterface(v interface{}) {
	b, ok := v.(*[]byte)
	fmt.Println(ok)
	fmt.Println(b)
}

func main() {
	passInterface(&[]byte{0x00, 0x01, 0x02})
}
