package intermediate

import (
	"errors"
	"fmt"
	"math"
)

func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("math error: square root of negative number")
	}
	return math.Sqrt(x), nil
}

func process(data []byte) error {
	if len(data) == 0 {
		return errors.New("error: empty data")
	}
	// Process data
	return nil
}

func readConfig() error {
	return errors.New("config error")
}

func readData() error {
	err := readConfig()
	if err != nil {
		return fmt.Errorf("readData: %w", err)
	}
	return nil
}

type myError struct {
	message string
}

// We are implementing the built in error interface
// https://github.com/golang/go/blob/master/src/builtin/builtin.go
func (m *myError) Error() string {
	return fmt.Sprintf("error: %s", m.message)
}

func eprocess() error {
	return &myError{"custom error message"}
}

func main() {
	// Error handling
	res1, err1 := sqrt(16)
	if err1 != nil {
		fmt.Println(err1)
	} else {
		fmt.Println(res1)
	}

	res2, err2 := sqrt(-16)
	if err2 != nil {
		fmt.Println(err2)
	} else {
		fmt.Println(res2)
	}

	data := []byte{}

	// One way of error handling
	if err3 := process(data); err3 != nil {
		fmt.Println(err3)
	} else {
		fmt.Println("data processed successfully")
	}

	// Another way of error handling
	err4 := process(data)
	if err4 != nil {
		fmt.Println(err4)
	} else {
		fmt.Println("data processed successfully")
	}

	// Using the built in error interface
	err5 := eprocess()
	if err5 != nil {
		fmt.Println(err5)
	}

	// fmt.Errorf example
	err6 := readData()
	if err6 != nil {
		fmt.Println(err6)
	} else {
		fmt.Println("data read successfully")
	}
}
