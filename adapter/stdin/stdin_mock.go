package stdin

import (
	"io"
	"strconv"
)

type StdinMock struct {
	data []string
}

func NewStdinMock() *StdinMock {
	return &StdinMock{
		data: []string{},
	}
}

func (s *StdinMock) Write(content string) error {
	s.data = append(s.data, content)
	return nil
}

func (s *StdinMock) Read() (string, error) {
	if len(s.data) == 0 {
		return "", io.EOF
	}
	result := s.data[0]
	s.data = s.data[1:]
	return result, nil
}

func (s *StdinMock) Close() error {
	return nil
}

func (s *StdinMock) ReadInt() (int, error) {
	content, err := s.Read()
	if err != nil {
		return 0, err
	}
	if content == "" {
		return 0, io.EOF
	}
	result, err := strconv.Atoi(content)
	if err != nil {
		return 0, err
	}
	return result, nil
}
