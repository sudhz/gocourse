package basics

import "fmt"

var lastName = "Makwana"

func main() {
	lastName := "AnotherOne"

	fmt.Println(lastName)

	printName()
}

func printName() {
	firstName := "Sudhanshu"
	fmt.Println(firstName)
}