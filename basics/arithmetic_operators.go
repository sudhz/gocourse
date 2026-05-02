package basics

import (
	"fmt"
)

func main() {
	var a, b int = 10, 3
	var result int

	result = a + b
	fmt.Println("Addition:", result)

	result = a - b
	fmt.Println("Subtraction:", result)

	result = a * b
	fmt.Println("Multiplication:", result)

	result = a / b
	fmt.Println("Division", result)

	result = a % b
	fmt.Println("Remainder:", result)

	// Overflow with signed integers
	var maxInt int64 = 9223372036854775807 // math.MaxInt64
	fmt.Println("maxInt:", maxInt)

	maxInt = maxInt + 1
	fmt.Println("maxInt=maxInt+1:", maxInt)

	// Overflow with unsigned integers
	var uMaxInt uint64 = 18446744073709551615 // math.MaxUint64
	fmt.Println("uMaxInt:", uMaxInt)

	uMaxInt = uMaxInt + 1
	fmt.Println("uMaxInt=uMaxInt+1:", uMaxInt)

	// Underflow with floating point numbers
	var smallFloat float64 = 1.0e-323
	fmt.Println("smallFloat:", smallFloat)

	smallFloat = smallFloat / 1.7976931348623157e+308 // math.MaxFloat64
	fmt.Println("smallFloat=smallFloat/1.7976931348623157e+308:", smallFloat)
}
