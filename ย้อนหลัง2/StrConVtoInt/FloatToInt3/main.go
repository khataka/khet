package main

import (
	"fmt"
)

func main() {
	nf := []float64{-1.9999, -2.0001, -2.0, 0, 1.9999, 2.0001, 2.0}

	fmt.Printf("Round : ")
	for _, f := range nf {
		fmt.Printf("%d ", round(f))
	}
	fmt.Printf("\n")

	fmt.Printf("RoundD: ")
	for _, f := range nf {
		fmt.Printf("%d ", roundD(f))
	}
	fmt.Printf("\n")

	fmt.Printf("RoundU: ")
	for _, f := range nf {
		fmt.Printf("%d ", roundU(f))
	}
	fmt.Printf("\n")

}

func roundU(val float64) int {
	if val > 0 {
		return int(val + 1.0)
	}
	return int(val)
}

func roundD(val float64) int {
	if val < 0 {
		return int(val - 1.0)
	}
	return int(val)
}

func round(val float64) int {
	if val < 0 {
		return int(val - 0.5)
	}
	return int(val + 0.5)
}
