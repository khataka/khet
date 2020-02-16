package main

import (
	"fmt"
	"time"
)

func main() {

	for true {
		fmt.Println("Infinite Loop")
		time.Sleep(time.Second)
	}
}
