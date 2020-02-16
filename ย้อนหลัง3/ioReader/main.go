package main

import (
	"fmt"
	"io"
	"os"
)

var _ io.Reader = (*os.File)(nil)

func main() {
	fmt.Println("Hello, playground")
}
