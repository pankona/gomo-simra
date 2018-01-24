// +build linux
// +build darwin
// +build windows

package storage

import "path/filepath"

type storageLinux struct {
	dirpath string
}

func NewStorage() Storager {
	return &storageLinux{
		dirpath: filepath.Join("."),
	}
}

func (s *storageLinux) DirectoryPath() string {
	return s.dirpath
}
