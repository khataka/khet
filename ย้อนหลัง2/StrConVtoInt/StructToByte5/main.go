package main

import (
	"bytes"
	"crypto/sha1"
	"encoding/binary"
	"fmt"
)

type myStruct struct {
	ID   [10]byte
	Data [10]byte
}

func main() {
	var bin_buf bytes.Buffer
	x := myStruct{"1", "Hello"}
	binary.Write(&bin_buf, binary.BigEndian, x)
	fmt.Printf("% x", sha1.Sum(bin_buf.Bytes()))
}
