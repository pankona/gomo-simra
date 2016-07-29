# How to edit AndroidManifest.xml

this text introduces the way to edit AndroidManifest.xml
after creating apk by `gomobile build`

## why edit AndroidManifest.xml ?

for example:

* edit `android:screenOrientation` to lock the screen lotation
* set specified icons to apk
* change application name
* etc ...

## overview of the way

following procedures will be done:

* extract generated apk by apktool
* remove unnecessary files in it for signing
* edit AndroidManifest.xml
* recreate apk and sign it by jarsigner
* then it can be uploadable to Google Play.

## tools to use

at first, install following tools.

* apktool

apktool can be installed via brew if you use OSX.

`$ brew install apktool`

## recipe

### generate apk

at first, create apk using `gomobile build`

`$ gomobile build`

now I assume that the command generates "sample.apk"

### extract it to edit AndroidManifest.xml

then, use apktool to extract it.

`$ apktool d sample.apk`

directory `sample` will be generated.
you can find some AndroidManifest.xml in it.

### remove files in META-INF

remove all files in META-INF directory.
this is necessary for sign correctly later.

note that you can find some META-INF directory.
all of them are the target for remove. 

## remove unnecessary AndroidManifest.xml

there may be three AndroidManifest.xml files in sample directory.
remove them all except top directory's one.

## edit AndroidManifest.xml

then there's only one AndroidManifest.xml in sample directory.
edit it as you like.

## recreate apk

after editing, re-create apk from the directory by following command.

`$ apktool b -c sample -o new_sample.apk`

"sample" is directory name. "new_sample.apk" is the name of apk newly created.

## sign using jarsigner

currently new_sample.apk is not signed. it means it cannot neither install to device even using adb nor upload to Google Play.
it is necessary to sign to do them.

### create keystore

create keystore like following command.
by following command, ".keystore" is generated to current directory.

`$ keytool -genkey -v -keystore .keystore -alias sample -keyalg RSA -validity 10000`

### sign apk using generated keystore

sign apk by following command.

`$ jarsigner -verbose -tsa http://timestamp.digicert.com -keystore .keystore ./new_sample.apk sample`

by following command, confirming whether the sign went well.

`$ jarsigner -verify -verbose -certs ./new_sample.apk`

if you don't see any error message, it means the sign has been succeeded.

## after them

enjoy install apk to your device and upload the apk to google play.
