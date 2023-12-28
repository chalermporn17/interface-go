package adapter

type Closer interface {
	Close() error
}

type Reader interface {
	Read() (string, error)
}

type ReadCloser interface {
	Reader
	Closer
}

type Writer interface {
	Clear() error
	Write(string) error
}

type WriteCloser interface {
	Writer
	Closer
}

type ReadWriter interface {
	Reader
	Writer
}

type ReadWriteCloser interface {
	ReadWriter
	Closer
}

type IntReader interface {
	Reader
	ReadInt() (int, error)
}

type IntReadCloser interface {
	IntReader
	Closer
}

type IntWriter interface {
	Writer
	WriteInt(int) error
}

type IntWriteCloser interface {
	IntWriter
	Closer
}

type IntReadWriter interface {
	IntReader
	IntWriter
}

type IntReadWriteCloser interface {
	IntReadWriter
	Closer
}
