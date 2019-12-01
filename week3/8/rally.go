package main

import "fmt"

func main() {

	fmt.Print("...")
	var text string

	fmt.Scanln(&text)
	fmt.Println("name :", text+" ", text)

}
