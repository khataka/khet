package main

import (
	"fmt"
)

func main() {
	f := []float64{1.9999, 2.0001, 2.0}

	fmt.Printf("RoundDown: %d %d %d \n", int(f[0]), int(f[1]), int(f[2]))
	fmt.Printf("RoundTo  : %d %d %d \n", int(f[0]+0.5), int(f[1]+0.5), int(f[2]+0.5))
	fmt.Printf("RoundUp  : %d %d %d \n", int(f[0]+1.0), int(f[1]+1.0), int(f[2]+1.0))
}
