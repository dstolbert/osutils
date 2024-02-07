package osutils

import (
	"encoding/csv"
	"io"
	"io/fs"
	"os"
	"os/exec"
	"syscall"
)

type osutils struct {
	/*
		generate::components
		skipTestFile::true
	*/
}

type Params struct{}

func (p *Params) Convert() *osutils {
	return &osutils{}
}

// Code below was generated by components. DO NOT EDIT.
// Component version: v0.1.0

type Osutils interface {
	CSVNewReader(r io.Reader) *csv.Reader
	CSVRead(csvReader *csv.Reader) ([]string, error)
	CommandStart(cmd *exec.Cmd) error
	CommandWait(cmd *exec.Cmd) error
	CommandOutput(cmd *exec.Cmd) ([]byte, error)
	SyscallKill(pid int, sig syscall.Signal) error
	StdoutPipe(cmd *exec.Cmd) (io.ReadCloser, error)
	Stat(name string) (fs.FileInfo, error)
	Open(name string) (*os.File, error)
	Create(name string) (*os.File, error)
	Mkdir(name string, mode os.FileMode) error
	MkdirAll(name string, mode os.FileMode) error
	Remove(name string) error
	RemoveAll(path string) error
	Getenv(key string) string
	IoReadAll(r io.Reader) ([]byte, error)
	JsonMarshal(v any) ([]byte, error)
	JsonUnmarshal(data []byte, v any) error
	JsonMarshalIndent(v any, prefix string, indent string) ([]byte, error)
}

func New(p Params) Osutils {
	return p.Convert()
}
