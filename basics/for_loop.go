package basics

import "fmt"

func main() {
	// Simple iteration over a range
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}

	// Iterate over collection
	numbers := []int{1, 2, 3, 4, 5, 6}
	for index, value := range numbers {
		fmt.Printf("index: %v, value: %v\n", index, value)
	}

	// Demonstration of continue and break
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println("odd number", i)
		if i == 5 {
			break
		}
	}

	// Multiple for loops to print asterisk pattern
	rows := 5
	for i := 1; i <= rows; i++ {
		for j := 1; j <= rows-i; j++ {
			fmt.Print(" ")
		}
		for range 2*i - 1 {
			fmt.Print("*")
		}
		fmt.Println()
	}

	// Latest syntax
	for i := range 10 {
		i++
		fmt.Println(i)
	}
	fmt.Println("we have a lift off")
}
