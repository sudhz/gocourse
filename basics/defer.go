package basics

import "fmt"

func main() {
	process(10)
}

func process(i int) {
	defer fmt.Println("deferred i value:", i)
	defer fmt.Println("first deferred statement executed")
	defer fmt.Println("second deferred statement executed")
	defer fmt.Println("third deferred statement executed")
	i++
	fmt.Println("normal execution statement")
	fmt.Println("value of i:", i)
}
