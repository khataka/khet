package main

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type MyStruct struct {
	Name string `json:"name"`
}

func main() {
	testStruct := MyStruct{"hello world"}
	reqBodyBytes := new(bytes.Buffer)
	json.NewEncoder(reqBodyBytes).Encode(testStruct)

	fmt.Println(reqBodyBytes.Bytes())
	fmt.Println(string(reqBodyBytes.Bytes()))
}
