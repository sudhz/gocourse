package basics

import "fmt"

func main() {
	process()

	fmt.Println("returned from process")
}

func process() {
	defer func() {
		r := recover()
		// can also write -> if r:= recover(); r != nil
		if r != nil {
			fmt.Println("recovered:", r)
		}
	}()

	fmt.Println("start process")
	panic("something went wrong!")
	fmt.Println("end process")
}
