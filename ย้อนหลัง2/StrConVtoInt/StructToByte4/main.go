package main

import (
	"fmt"
	"unsafe"
)

type A struct {
	Src     int32
	Dst     int32
	SrcPort uint16
	DstPort uint16
}

type ByteSliceA struct {
	Addr *A
	Len  int
	Cap  int
}

func main() {

	a := A{0x04030201, 0x08070605, 0x0A09, 0x0C0B}

	sb := &ByteSliceA{&a, 12, 12}

	var byteSlice []byte = *(*[]byte)(unsafe.Pointer(sb))

	fmt.Printf("%v\n", byteSlice)
}
