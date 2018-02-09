// +build android

package storage

/*
#include "storage_android.c"
*/
import "C"

import (
	"unsafe"

	"github.com/pankona/gomo-simra/simra/internal/jni"
	"github.com/pankona/gomo-simra/simra/simlog"
)

type storageAndroid struct{}

func NewStorage() Storager {
	return &storageAndroid{}
}

var path string

func (s *storageAndroid) DirectoryPath() string {
	if path != "" {
		return path
	}
	jni.RunOnJVM(
		func(vm, env, ctx uintptr) error {
			cpath := C.getFilesDir(C.uintptr_t(vm), C.uintptr_t(env), C.uintptr_t(ctx))
			if cpath == nil {
				simlog.Errorf("failed to get FilesDir!")
			}
			path = C.GoString(cpath)
			C.free(unsafe.Pointer(cpath)) // #nosec
			return nil
		})
	return path
}
