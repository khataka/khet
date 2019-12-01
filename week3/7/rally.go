package main

import "fmt"

func main() {

	fmt.Print("input : ")
	var name string

	fmt.Scanln(&name)
	fmt.Println("name", name+name)

}
