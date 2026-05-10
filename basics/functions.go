package basics

import "fmt"

func main() {
	fmt.Println(add(1, 2))

	greet := func() {
		fmt.Println("Hello Anonymous Function")
	}

	greet()

	operation := add

	result := operation(3, 5)

	fmt.Println(result)

	// passing a function as an argument
	result = applyOperation(5, 3, add)
	fmt.Println("5 + 3 =", result)

	// returning and using a function
	multiplyBy2 := createMultiplier(2)
	fmt.Println("6 * 2 =", multiplyBy2(6))
}

func add(a, b int) int {
	return a + b
}

// function that takes a function as an argument
func applyOperation(x int, y int, operation func(int, int) int) int {
	return operation(x, y)
}

// function that returns a function
func createMultiplier(factor int) func(int) int {
	return func(x int) int {
		return x * factor
	}
}
