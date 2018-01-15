package main

import (
	"io"
	"os"
	"path/filepath"
)

type localDirStrategy struct {
}

func (s *localDirStrategy) read(id string) string {
	dir := config("dir")
	return filepath.Join(dir, id)
}

func (s *localDirStrategy) add(id string, data io.ReadCloser) string {
	path      := config("dir")
	fullname  := filepath.Join(path, id)
	file, err := os.Create(fullname)
	if err != nil {
		return "error"
	}
	io.Copy(file, data)
	if r := recover(); r != nil {
		return "error"
	}
	return "ok"
}

func (s *localDirStrategy) update(id string, data io.ReadCloser) string {
	return ""
}

func (s *localDirStrategy) delete(id string) string {
	return ""
}
