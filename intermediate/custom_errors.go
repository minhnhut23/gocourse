package intermediate

import (
	"errors"
	"fmt"
)

func main() {

	err := doSomething()
	if err != nil {
		fmt.Print(err)
		return
	}

}

type customError struct {
	code    int
	message string
	err     error
}

// Error return the error message. Implementing Error() method of error
func (e *customError) Error() string {
	return fmt.Sprintf("Error %d: %s! %v\n ", e.code, e.message, e.err)
}

// Function that return a custom error
// func doSomething() error {
// 	return &customError{
// 		code:    500,
// 		message: "Something went wrong",
// 	}
// }

func doSomething() error {
	er := doSomethingElse()
	if er != nil {
		return &customError{
			code:    500,
			message: "Something went wrong",
			err:     er,
		}
	}
	return nil
}

func doSomethingElse() error {
	return errors.New("internal error")
}
