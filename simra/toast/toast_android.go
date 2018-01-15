// +build android

package toast

/*
#include <jni.h>
#include <stdlib.h>

static jclass android_app_Activity;
static jclass android_os_Bundle;
static jclass android_widget_Toast;

static int
showToast(uintptr_t java_vm, uintptr_t* jni_env, jobject thiz, jobject text) {
    JavaVM* vm  = (JavaVM*)java_vm;
	JNIEnv* env = (JNIEnv*)jni_env;
	int n = (*env)->CallStaticIntMethod(
		env, android_media_AudioTrack,
		(*env)->GetStaticMethodID(env, android_media_AudioTrack, "makeText", "(III)I"),
		sampleRate, channel, encoding);

	jclass toastClass = (*env)->FindClass(env, "android/widget/Toast");
	if(toastClass == NULL) {
		return 1;
	}

	jmethodID makeText = (*env)->GetStaticMethodID(
			env, toastClass, "makeText", "(Landroid/content/Context;Ljava/lang/CharSequence;I)Landroid/widget/Toast;");
	if(makeText == NULL) {
		return 1;
	}

	jobject toastObject = (*env)->CallStaticObjectMethod(env, toastClass, makeText, thiz, text, 0); // 0 = Toast.LENGTH_SHORT
	if(toastObject == NULL) {
		return 1;
	}

	jmethodID show = (*env)->GetMethodID(env, toastClass, "show", "()V");
	if(show == NULL) {
		return 1;
	}

	(*env)->CallVoidMethod(env, toastObject, show);
	return 0;
}
*/
import "C"

type t struct{}

func NewToaster() Toaster {
	return &t{}
}

func (t *t) Show(text string) error {
	jni.RunOnJVM(
		func(vm, env, ctx uintptr) error {
			C.showToast(C.uintptr_t(vm), C.uintptr_t(env), C.uintptr_t(ctx), C.GoString(text))
			return nil
		})
	return nil
}
