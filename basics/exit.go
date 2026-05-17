package basics

import (
	"fmt"
	"os"
)

func main() {
	defer fmt.Println("deferred statement")

	fmt.Println("starting the main function")

	os.Exit(1)

	fmt.Println("end of main function")
}