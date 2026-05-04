package basics

import "fmt"

func main() {
	// Arrays syntax
	var numbers [5]int
	fmt.Println(numbers)

	numbers[4] = 20
	fmt.Println(numbers)

	numbers[0] = 9
	fmt.Println(numbers)

	fruits := [4]string{"Apple", "Banana", "Orange", "Grapes"}
	fmt.Println("fruits array:", fruits)

	fmt.Println("third element:", fruits[2])

	// Underscore is blank identifier, used to store unused variables

	for i := 0; i < len(numbers); i++ {
		fmt.Println("element at index,", i, ":", numbers[i])
	}

	for _, v := range numbers {
		fmt.Printf("Value: %d\n", v)
	}

	a, _ := someFunction()
	fmt.Println(a)

	// Comparing arrays
	array1 := [3]int{1, 2, 3}
	array2 := [3]int{10, 2, 3}

	fmt.Println("array1 is equal to array2:", array1 == array2)

	// Matrix
	matrix := [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	fmt.Println(matrix)

	// Copying arrays
	originalArray := [3]int{1, 2, 3}
	var copiedArray *[3]int

	copiedArray = &originalArray
	copiedArray[0] = 100

	fmt.Println("originalArray:", originalArray)
	fmt.Println("copiedArray:", copiedArray)
}

func someFunction() (int, int) {
	return 1, 2
}
