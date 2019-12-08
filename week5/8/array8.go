package main

import "fmt"

func main() {

	numbers := [3][4][3]int{
		{
			{1, 2, 3},
			{10, 20, 30},
			{100, 200, 300},
		},
		{
			{8, 9, 10},
			{80, 90, 100},
			{800, 900, 1000},
		},
		{
			{4, 5, 6},
			{40, 50, 60},
			{400, 500, 600},
		},
	}
	fmt.Println(numbers)
	fmt.Println(numbers[1][3][2])

}
