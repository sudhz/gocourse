package basics

import (
	"fmt"
	"slices"
)

func main() {
	a := [5]int{1, 2, 3, 4, 5}
	slice1 := a[1:4]

	fmt.Println("slice1:", slice1)

	slice1 = append(slice1, 6, 7)
	fmt.Println("slice1:", slice1)

	sliceCopy := make([]int, len(slice1))
	copy(sliceCopy, slice1)

	fmt.Println("sliceCopy:", sliceCopy)

	for i, v := range slice1 {
		fmt.Printf("element at index %d is %d\n", i, v)
	}

	fmt.Println("element at index 3 of slice1:", slice1[3])

	if slices.Equal(slice1, sliceCopy) {
		fmt.Println("slice1 is equal to sliceCopy")
	}

	twoD := make([][]int, 3)
	for i := range 3 {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := range innerLen {
			twoD[i][j] = i + j
			fmt.Printf("adding value %d in outer slice at index %d, and in inner slice index of %d\n", i+j, i, j)
		}
	}

	fmt.Println("twoD:", twoD)

	slice2 := slice1[2:4]
	fmt.Println("slice2:", slice2)
}
