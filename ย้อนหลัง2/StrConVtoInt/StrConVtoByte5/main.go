package main

import (
	"bytes"
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	a := "hello"
	b := []byte("hello")

	p := *(*[]byte)(unsafe.Pointer((*reflect.SliceHeader)(unsafe.Pointer(&a))))

	fmt.Println(bytes.Compare(b, p))
}
