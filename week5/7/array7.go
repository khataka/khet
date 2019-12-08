package main

import "fmt"

func main() {
	alphabets := [2][3]string{{"a", "b", "c"}, {"A", "B", "C"}}
	fmt.Println(alphabets)
	fmt.Println(alphabets[0][1])
	numbers := [2][3][2]int{
		{
			{1, 2},
			{10, 20},
			{100, 200},
		},
		{
			{8, 9},
			{80, 90},
			{800, 900},
		},
	}
	fmt.Println(numbers)
	fmt.Println(numbers[1][2][0])

}
