package file

import (
	"io"
	"os"
	"strings"
)

type File struct {
	file *os.File
}

func NewFile(filename string) (*File, error) {
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return nil, err
	}
	return &File{
		file: file,
	}, nil
}

func (f *File) Read() (string, error) {
	defer f.file.Seek(0, 0)
	f.file.Seek(0, 0)
	buffer := make([]byte, 100)
	var builder strings.Builder
	for {
		n, err := f.file.Read(buffer)
		if err != nil {
			if err == io.EOF {
				break
			}
			return "", err
		}
		builder.Write(buffer[:n])
	}
	return builder.String(), nil
}

func (f *File) Clear() error {
	return f.file.Truncate(0)
}

func (f *File) Write(content string) error {
	_, err := f.file.Write([]byte(content))
	return err
}

func (f *File) Close() error {
	return f.file.Close()
}
