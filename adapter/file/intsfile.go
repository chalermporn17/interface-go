package file

import (
	"io"
	"strconv"
	"strings"
)

type IntsFile struct {
	*File
	data []int
}

func NewIntsFile(filename string) (*IntsFile, error) {
	file, err := NewFile(filename)
	if err != nil {
		return nil, err
	}
	result := &IntsFile{
		File: file,
		data: []int{},
	}
	err = result.initContent()
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (f *IntsFile) initContent() error {
	content, err := f.Read()
	if err != nil {
		return err
	}
	intsStr := strings.Split(content, "\n")
	for _, intStr := range intsStr {
		if intStr != "" {
			number, err := strconv.Atoi(intStr)
			if err != nil {
				return err
			}
			f.data = append(f.data, number)
		}
	}
	return nil
}

func (f *IntsFile) ReadInt() (int, error) {
	if len(f.data) == 0 {
		return 0, io.EOF
	}
	result := f.data[0]
	f.data = f.data[1:]
	return result, nil
}
