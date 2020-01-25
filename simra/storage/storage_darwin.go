// +build !android

package storage

import "path/filepath"

type storageLinux struct {
	dirpath string
}

// NewStorage returns an instance that implemensts Storager interface
func NewStorage() Storager {
	return &storageLinux{
		dirpath: filepath.Join("."),
	}
}

func (s *storageLinux) DirectoryPath() string {
	return s.dirpath
}
