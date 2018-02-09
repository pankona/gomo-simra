#include <jni.h>
#include <stdlib.h>

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
