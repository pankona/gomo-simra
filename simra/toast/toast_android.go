// +build android

package toast

/*
#include <jni.h>
#include <stdlib.h>

static jclass android_app_Activity;
static jclass android_os_Bundle;
static jclass android_widget_Toast;

static int
showToast(uintptr_t java_vm, uintptr_t jni_env, uintptr_t jni_ctx, char* text) {
	JNIEnv* env = (JNIEnv*)jni_env;
	jobject ctx = (jobject)jni_ctx;

	jclass toast = (*env)->FindClass(env, "android/widget/Toast");
	if(toast == NULL) {
		return 1;
	}
	jmethodID methodMakeText = (*env)->GetStaticMethodID(env, toast, "makeText",
		"(Landroid/content/Context;Ljava/lang/CharSequence;I)Landroid/widget/Toast;");
	if(methodMakeText == NULL){
		return 2;
	}

	jobject toastobj = (*env)->CallStaticObjectMethod(env, toast, methodMakeText,
		ctx, (jstring)"hoge", 0);
	if (toastobj == NULL) {
		return 3;
	}
	jmethodID methodShow = (*env)->GetMethodID(env, toast, "show", "()V");
	if (methodShow == NULL) {
		return 4;
	}
	(*env)->CallVoidMethod(env, toastobj, methodShow);

	//jobject toastObject = (*env)->CallStaticObjectMethod(env, toast, methodMakeText,
	//	ctx, "hoge", 0); // 0 = Toast.LENGTH_SHORT
	//if(toast == NULL) {
	//	return 3;
	//}

	//jmethodID show = (*env)->GetMethodID(env, toastClass, "show", "()V");
	//if(show == NULL) {
	//	return 4;
	//}

	//(*env)->CallVoidMethod(env, toastObject, show);
	return 0;
}
*/
import "C"

import (
	"fmt"

	"github.com/pankona/gomo-simra/simra/jni"
)

type t struct{}

func NewToaster() Toaster {
	return &t{}
}

func (t *t) Show(text string) error {
	jni.RunOnJVM(
		func(vm, env, ctx uintptr) error {
			ret := C.showToast(C.uintptr_t(vm), C.uintptr_t(env), C.uintptr_t(ctx), C.CString(text))
			if ret != 0 {
				fmt.Printf("error!! ret = %d\n", ret)
			}
			return nil
		})
	return nil
}
