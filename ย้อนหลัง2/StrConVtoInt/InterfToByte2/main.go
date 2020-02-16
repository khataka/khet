package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	u := map[string]interface{}{}
	u["name"] = "kish"
	u["age"] = 28
	u["work"] = "engine"

	u["hobbies"] = "art"

	b, err := json.Marshal(u)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))

	var a interface{}
	err = json.Unmarshal(b, &a)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(a)
}
