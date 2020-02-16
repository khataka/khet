package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

func main() {
	var f float64 = 12.666512
	var result []byte = float64ToByte(f)
	fmt.Printf("result:%+v", result)
}

func float64ToByte(f float64) []byte {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.BigEndian, f)
	if err != nil {
		fmt.Println("binary.Write failed:", err)
	}
	return buf.Bytes()
}
