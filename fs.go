package osutils

import (
	"encoding/json"
	"io"
	"io/fs"
	"os"
)

func (o *osutils) ReadDir(name string) ([]fs.DirEntry, error) {
	return os.ReadDir(name)
}

func (o *osutils) Stat(name string) (fs.FileInfo, error) {
	return os.Stat(name)
}

func (o *osutils) Open(name string) (*os.File, error) {
	return os.Open(name)
}

func (o *osutils) Create(name string) (*os.File, error) {
	return os.Create(name)
}

func (o *osutils) Mkdir(name string, mode os.FileMode) error {
	return os.Mkdir(name, mode)
}

func (o *osutils) MkdirAll(name string, mode os.FileMode) error {
	return os.Mkdir(name, mode)
}

func (o *osutils) Remove(name string) error {
	return os.Remove(name)
}

func (o *osutils) RemoveAll(path string) error {
	return os.RemoveAll(path)
}

func (o *osutils) Getenv(key string) string {
	return os.Getenv(key)
}

func (o *osutils) IoReadAll(r io.Reader) ([]byte, error) {
	return io.ReadAll(r)
}

func (o *osutils) JsonMarshal(v any) ([]byte, error) {
	return json.Marshal(v)
}

func (o *osutils) JsonUnmarshal(data []byte, v any) error {
	return json.Unmarshal(data, v)
}

func (o *osutils) JsonMarshalIndent(v any, prefix string, indent string) ([]byte, error) {
	return json.MarshalIndent(v, prefix, indent)
}
