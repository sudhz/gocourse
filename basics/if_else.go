package main

import "fmt"

func main() {
	num := 18

	fmt.Println("num:", num)

	if num%2 == 0 {
		if num%3 == 0 {
			fmt.Println("Number is divisible by both 2 and 3")
		} else {
			fmt.Println("Number is divisible by 2 but not 3")
		}
	} else {
		fmt.Println("Number is not divisble by 2")
	}
}
