package main

import (
	"fmt"
	"io"
	"strings"
)

// errors
type AppError struct {
	Err    error
	Custom string
}

func (e *AppError) Error() string {
	fmt.Println(e.Custom)
	return e.Err.Error()
}

func m() error {
	return &AppError{
		Err:    fmt.Errorf("Custom Error"),
		Custom: "Value Custom Error",
	}
}

//reader interface

func main() {
	// Errors
	err := m()
	if err != nil {
		fmt.Println(err)
	}

	//IO reader
	r := strings.NewReader("Hello World")

	arr := make([]byte, 8)

	for {
		n, err := r.Read(arr)
		fmt.Printf("n = %d, err = %v, arr = %v\n", n, err, arr)
		fmt.Printf("arr n bytes: %q\n", arr[:n])
		if err == io.EOF {
			break
		}
	}

}
