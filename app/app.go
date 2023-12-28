package app

import (
	"interface/test/adapter"
	"io"
	"strconv"
)

func readInt(backend adapter.Reader) (int, error) {
	content, err := backend.Read()
	if err != nil {
		return 0, err
	}
	if content == "" {
		return 0, nil
	}
	result, err := strconv.Atoi(content)
	if err != nil {
		return 0, err
	}
	return result, nil
}

func Sum(input adapter.IntReadCloser, backend adapter.ReadWriteCloser) (int, error) {
	defer input.Close()
	defer backend.Close()
	sum, err := readInt(backend)
	if err != nil {
		return 0, err
	}
	newNumber, err := sumUntilEnd(input)
	if err != nil {
		return 0, err
	}
	sum += newNumber
	err = backend.Clear()
	if err != nil {
		return 0, err
	}
	err = backend.Write(strconv.Itoa(sum))
	if err != nil {
		return 0, err
	}
	return sum, nil
}

func sumUntilEnd(input adapter.IntReader) (int, error) {
	sum := 0
	for {
		number, err := input.ReadInt()
		if err != nil {
			if err == io.EOF {
				break
			}
			return 0, err
		}
		sum += number
	}
	return sum, nil
}
