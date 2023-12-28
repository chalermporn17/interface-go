package main

import (
	"fmt"
	"interface/test/adapter"
	"interface/test/adapter/file"
	"interface/test/adapter/stdin"
	"interface/test/app"
	"os"
)

func selectInput() (adapter.IntReadCloser, error) {
	inputName := os.Getenv("INPUT_NAME")
	var input adapter.IntReadCloser
	var err error

	switch inputName {
	case "stdin":
		input, err = stdin.NewStdin()
	case "file":
		input, err = file.NewIntsFile("input.txt")
	default:
		return nil, fmt.Errorf("invalid input name: %s", inputName)
	}
	if err != nil {
		return nil, err
	}
	return input, nil
}

func selectBackend() (adapter.ReadWriteCloser, error) {
	backendName := os.Getenv("BACKEND_NAME")
	var backend adapter.ReadWriteCloser
	var err error

	switch backendName {
	case "file":
		backend, err = file.NewFile("data.txt")
	default:
		return nil, fmt.Errorf("invalid backend name: %s", backendName)
	}
	if err != nil {
		return nil, err
	}
	return backend, nil
}

func main() {
	input, err := selectInput()
	if err != nil {
		panic(err)
	}
	backend, err := selectBackend()
	if err != nil {
		panic(err)
	}
	result, err := app.Sum(input, backend)
	if err != nil {
		panic(err)
	}
	fmt.Println("result is", result)
}
