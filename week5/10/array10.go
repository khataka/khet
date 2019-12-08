package main

import "fmt"

func main() {
	alphabets := [3][4]string{{"a", "b", "c", "d"}, {"A", "B", "C", "D"}, {"G", "H", "I", "J"}}

	fmt.Println(alphabets[1][3])

	numbers := [3][4][4]int{
		{
			{0, 1, 2, 3},
			{0, 10, 20, 30},
			{0, 100, 200, 300},
		},
		{
			{0, 8, 9, 10},
			{0, 80, 90, 100},
			{0, 800, 900, 1000},
		},
		{
			{0, 4, 5, 6},
			{0, 40, 50, 60},
			{0, 400, 500, 600},
		},
	}
	fmt.Println(numbers[2][1][0])

}
