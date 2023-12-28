package file

import "errors"

type FileMock struct {
	failToRead  bool
	failToWrite bool
	failToClear bool
	data        string
}

func NewFileMock() *FileMock {
	return &FileMock{
		data: "",
	}
}

func (f *FileMock) SetFailToRead(failToRead bool) {
	f.failToRead = failToRead
}

func (f *FileMock) SetFailToWrite(failToWrite bool) {
	f.failToWrite = failToWrite
}

func (f *FileMock) SetFailToClear(failToClear bool) {
	f.failToClear = failToClear
}

func (f *FileMock) Read() (string, error) {
	if f.failToRead {
		return "", errors.New("fail to read")
	}
	return f.data, nil
}

func (f *FileMock) Write(content string) error {
	if f.failToWrite {
		return errors.New("fail to write")
	}
	f.data = content
	return nil
}

func (f *FileMock) Clear() error {
	if f.failToClear {
		return errors.New("fail to clear")
	}
	f.data = ""
	return nil
}

func (f *FileMock) Close() error {
	return nil
}
