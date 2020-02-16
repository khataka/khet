package main

import (
	"encoding/json"
	"fmt"
)

type MyAddress struct {
	House  string
	School string
}
type Student struct {
	Id      int64
	Name    string
	Scores  float32
	Address MyAddress
	Labels  []string
}

func Test() {

	dict := make(map[string]interface{})
	dict["id"] = 201902181425
	dict["name"] = "jackytse"
	dict["scores"] = 123.456
	dict["address"] = map[string]string{"house": "my house", "school": "my school"}
	dict["labels"] = []string{"aries", "warmhearted", "frank"}

	jsonbody, err := json.Marshal(dict)
	if err != nil {

		fmt.Println(err)
		return
	}

	student := Student{}
	if err := json.Unmarshal(jsonbody, &student); err != nil {

		fmt.Println(err)
		return
	}

	fmt.Printf("%#v\n", student)
}

func main() {
	Test()
}
