package basics

import "fmt"

func main() {
	// example for valid input
	process(10)

	// example for invalid input
	process(-3)
}

func process(input int) {
	defer fmt.Println("deferred 1")
	defer fmt.Println("deferred 2")

	if input < 0 {
		fmt.Println("before panic")
		panic("input must be a non negative number")
		fmt.Println("after panic")
	}

	fmt.Println("processing input:", input)
}
