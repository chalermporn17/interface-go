package app

import (
	"interface/test/adapter/file"
	"interface/test/adapter/stdin"
	"testing"
)

func TestAdd_OK(t *testing.T) {
	input := stdin.NewStdinMock()
	backend := file.NewFileMock()
	input.Write("1")
	input.Write("2")
	input.Write("2")
	sum, err := Sum(input, backend)
	if err != nil {
		t.Fatal(err)
	}
	if sum != 5 {
		t.Fatal("sum should be 5")
	}
	input.Write("3")
	input.Write("4")
	sum, err = Sum(input, backend)
	if err != nil {
		t.Fatal(err)
	}
	if sum != 12 {
		t.Fatal("sum should be 12")
	}
}

func TestAdd_InvalidBackendContent(t *testing.T) {
	input := stdin.NewStdinMock()
	backend := file.NewFileMock()
	input.Write("1")
	input.Write("2")
	input.Write("2")
	backend.Write("invalid")
	_, err := Sum(input, backend)
	if err == nil {
		t.Fatal("should have failed")
	}
}

func TestAdd_InvalidInputContent(t *testing.T) {
	input := stdin.NewStdinMock()
	backend := file.NewFileMock()
	input.Write("1")
	input.Write("invalid")
	_, err := Sum(input, backend)
	if err == nil {
		t.Fatal("should have failed")
	}
}

func TestAdd_InvalidReadBackend(t *testing.T) {
	input := stdin.NewStdinMock()
	backend := file.NewFileMock()
	backend.SetFailToRead(true)
	_, err := Sum(input, backend)
	if err == nil {
		t.Fatal("should have failed")
	}
}

func TestAdd_InvalidWriteBackend(t *testing.T) {
	input := stdin.NewStdinMock()
	backend := file.NewFileMock()
	backend.SetFailToWrite(true)
	_, err := Sum(input, backend)
	if err == nil {
		t.Fatal("should have failed")
	}
}

func TestAdd_InvalidClearBackend(t *testing.T) {
	input := stdin.NewStdinMock()
	backend := file.NewFileMock()
	backend.SetFailToClear(true)
	_, err := Sum(input, backend)
	if err == nil {
		t.Fatal("should have failed")
	}
}
