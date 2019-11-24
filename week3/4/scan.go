package main

import "fmt"

func main() {
	fmt.Print("input : ")
	var name string
	var age int
	var height float32
	var weight float32
	n, err := fmt.Scan(&name, &age, &height, &weight)
	fmt.Println(name, age, weight, height)
	fmt.Println("number of argument", n)
	fmt.Println("Error", err)

}
