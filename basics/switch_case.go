package basics

import "fmt"

func main() {
	// Multiple cases in a switch
	day := "Monday"
	switch day {
	case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
		fmt.Println("It's a weekday")
	case "Saturday", "Sunday":
		fmt.Println("It's a weekend")
	default:
		fmt.Println("Invalid day")
	}

	// Fallthrough in switch case
	num := 2
	switch {
	case num > 1:
		fmt.Println("Number is greater than 1")
		fallthrough
	case num == 2:
		fmt.Println("Number is 2")
	default:
		fmt.Println("Not two")
	}

	// Type switch
	checkType(10)
	checkType(3.14)
	checkType("Hello")
	checkType(true)
}

func checkType(x any) {
	switch x.(type) {
	case int:
		fmt.Println("It's an integer")
	case float64:
		fmt.Println("It's a float")
	case string:
		fmt.Println("It's a string")
	default:
		fmt.Println("Unknown type")
	}
}