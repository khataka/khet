package Max

import "fmt"

func Max(num ...int) int {
	temp := num[0]
	for i := 0; i < len(num); i++ {

		if temp < num[i] {
			temp = num[i]
			fmt.Println("Max number =", i)
		}
	}
	return temp
}
func main() {}