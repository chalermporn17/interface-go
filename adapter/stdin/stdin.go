package stdin

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

type Stdin struct {
	reader *bufio.Reader
}

func NewStdin() (*Stdin, error) {
	reader := bufio.NewReader(os.Stdin)
	return &Stdin{
		reader: reader,
	}, nil
}

func (s *Stdin) Read() (string, error) {
	content, err := s.reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	return strings.Trim(content, "\n"), nil
}

func (s *Stdin) Close() error {
	return nil
}

func (s *Stdin) ReadInt() (int, error) {
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
