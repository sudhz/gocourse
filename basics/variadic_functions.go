package basics

import "fmt"

func main() {
	sequence, total := sum(1, 10, 20, 30, 40, 50, 60)
	fmt.Println("sequence:", sequence, "total:", total)

	sequence2, total2 := sum(2, 40, 36, 40, 50, 60)
	fmt.Println("sequence2:", sequence2, "total:", total2)

	numbers := []int{1, 2, 3, 4, 5, 9}
	sequence3, total3 := sum(3, numbers...)
	fmt.Println("sequence:", sequence3, "total:", total3)
}

func sum(sequence int, nums ...int) (int, int) {
	total := 0
	for _, num := range nums {
		total += num
	}
	return sequence, total
}
