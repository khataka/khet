package main

import (
	"fmt"

	"github.com/fatih/structs"
)

type MyStructObject struct {
	Email string `json:"email_address"`
}

func main() {
	obj := &MyStructObject{Email: "test@test.com"}

	fmt.Println(StructToMap(obj))

	fmt.Println(structs.Map(obj))
}
