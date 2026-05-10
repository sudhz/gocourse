package basics

import "fmt"

func main() {
	message := "Hello World"

	for i, v := range message {
		fmt.Printf("index: %d, rune: %c\n", i, v)
	}
}
