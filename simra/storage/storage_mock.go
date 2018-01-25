// +build !linux
// +build !darwin
// +build !windows
// +build !android

package storage

import "path/filepath"

type storageMock struct{}

func NewStorage() Storager {
	return &storageMock{}
}

func (s *storageAndroid) DirectoryPath() string {
	// not implemented
	return filepath.Join(".")
}
