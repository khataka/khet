package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

func main() {
	var f float32 = 1.0
	var result []byte = float64ToByte(f)
	fmt.Printf("result:%#v", result)
}
func float64ToByte(f float32) []byte {
	var buf bytes.Buffer
	err := binary.Write(&buf, binary.LittleEndian, f)
	if err != nil {
		fmt.Println("binary.Write failed:", err)
	}
	return buf.Bytes()
}
