package main

import (
	"os"
)

func main() {
	f, err := os.Open("somefile.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	//var r := io.Reader
	//r = f
}
