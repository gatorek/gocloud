package main

import (
	"io"
)

type FileAcceess interface {
	  // takes ID from GET, returns file path
	  read(id string)                     string
	   add(id string, data io.ReadCloser) string
	update(id string, data io.ReadCloser) string
	delete(id string)                     string
}

func getFileAccessor(strategy string) FileAcceess {
	switch strategy {
		case "local":
			return new(localDirStrategy)
	}
	return nil
}
