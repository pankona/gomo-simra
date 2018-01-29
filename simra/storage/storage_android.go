// +build android

package storage

/*
#cgo LDFLAGS: -landroid -llog
#include <jni.h>
#include <stdlib.h>
#include <android/log.h>

static const char *
getFilesDir(uintptr_t java_vm, uintptr_t jni_env, uintptr_t jni_ctx) {
	JNIEnv* env = (JNIEnv*)jni_env;
	jobject ctx = (jobject)jni_ctx;

	jclass context = (*env)->FindClass(env, "android/content/Context");
	if(context == NULL) {
		return NULL;
	}
	jmethodID getFilesDir = (*env)->GetMethodID(env, context, "getFilesDir", "()Ljava/io/File;");
	if(getFilesDir == NULL){
		return NULL;
	}

	jobject f = (*env)->CallObjectMethod(env, ctx, getFilesDir);
	if (f == NULL) {
		return NULL;
	}

	jclass file = (*env)->FindClass(env, "java/io/File");
	if (file == NULL) {
		return NULL;
	}

	jmethodID getAbsolutePath = (*env)->GetMethodID(env, file, "getAbsolutePath", "()Ljava/lang/String;");
	if (getAbsolutePath == NULL) {
		return NULL;
	}

	jstring path = (jstring)(*env)->CallObjectMethod(env, f, getAbsolutePath);
	return (*env)->GetStringUTFChars(env, path, 0);
}
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
			)}
			path = C.GoString(cpath)
			C.free(unsafe.Pointer(cpath))
			return nil
		})
	return path
}
