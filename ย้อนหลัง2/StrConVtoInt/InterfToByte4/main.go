package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	u := map[string]interface{}{}
	u["1"] = "one"
	b, err := json.Marshal(u)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))
}
