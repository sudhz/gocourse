package intermediate

import (
	"errors"
	"fmt"
)

type customError struct {
	code    int
	message string
	err     error
}

// Implementing the Error method from the error interface
func (e *customError) Error() string {
	if e.err != nil {
		return fmt.Sprintf("error %d: %s: %v", e.code, e.message, e.err)
	}
	return fmt.Sprintf("error %d: %s", e.code, e.message)
}

func doSomething1() error {
	return &customError{
		code:    500,
		message: "something went wrong",
	}
}

func doSomething2() error {
	innerErr := doSomethingElse()
	if innerErr != nil {
		return &customError{
			code:    500,
			message: "something went wrong",
			err:     innerErr,
		}
	}
	return nil
}

func doSomethingElse() error {
	return errors.New("internal error")
}

func main() {
	err1 := doSomething1()
	if err1 != nil {
		fmt.Println(err1)
	} else {
		fmt.Println("operation completed successfully")
	}

	err2 := doSomething2()
	if err2 != nil {
		fmt.Println(err2)
	} else {
		fmt.Println("operation completed successfully")
	}
}
