// +build android

package storage

import "os"

type storageAndroid struct {
	dirpath string
}

func NewStorage() Storager {
	return &storageAndroid{
		dirpath: os.Getenv("TMPDIR"),
	}
}

func (s *storageAndroid) DirectoryPath() string {
	s.dirpath
}
