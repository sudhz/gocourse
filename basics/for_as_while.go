package basics

import "fmt"

func main() {
	// for loop as while loop with break
	sum := 0
	for {
		sum += 10
		fmt.Println("sum:", sum)
		if sum >= 50 {
			break
		}
	}

	// for loop as while loop with continue
	num := 1
	for num <= 10 {
		if num%2 == 0 {
			num++
			continue
		}
		fmt.Println("odd number:", num)
		num++
	}
}
