package errors

import (
	"fmt"
	"strconv"
	"time"
)

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {

	return fmt.Sprintf("at %v, %s", e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

func MainError() {
	if err := run(); err != nil {
		fmt.Println(err)
	}

	_, err := strconv.Atoi("abs")
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
}
