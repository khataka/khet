package main

import (
	"fmt"
	"reflect"
)

type User struct {
	name string
}

func main() {
	var user interface{}
	user = User{name: "Eugene"}

	switch v := user.(type) {
	case int:

		fmt.Printf("Integer: %v", v)
	case User:

		fmt.Printf("It's a user: %s\n", user.(User).name)
	}

	userType := reflect.TypeOf(user)
	fmt.Printf("%+v\n", userType)

	fmt.Printf("%T", user)

	userTypeAsString := fmt.Sprintf("%T", user)

	if userTypeAsString == "main.User" {
		fmt.Printf("\nIt's definitely a user")
	}
}
