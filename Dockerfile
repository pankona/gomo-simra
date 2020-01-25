FROM circleci/golang:1.13.6

ENV HOME /home/circleci
WORKDIR $HOME

ENV PATH $GOPATH/bin:$PATH

RUN wget https://dl.google.com/dl/android/studio/ide-zips/3.5.3.0/android-studio-ide-191.6010548-linux.tar.gz
RUN tar zxf android-studio-ide-191.6010548-linux.tar.gz

ENV JAVA_HOME $HOME/android-studio/jre
ENV PATH $PATH:$JAVA_HOME/bin

RUN sudo apt-get update
RUN sudo apt-get install -y libegl1-mesa-dev
RUN sudo apt-get install -y libgles2-mesa-dev
RUN sudo apt-get install -y libx11-dev
RUN sudo apt-get install -y libasound2-dev
RUN go get -u golang.org/x/mobile/cmd/gomobile
RUN curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh -s -- -b $(go env GOPATH)/bin v1.23.1

# install android sdk
ENV ANDROID_SDK_VERSION=4333796
RUN wget https://dl.google.com/android/repository/sdk-tools-linux-$ANDROID_SDK_VERSION.zip
RUN unzip -q sdk-tools-linux-$ANDROID_SDK_VERSION.zip
RUN yes | $HOME/tools/bin/sdkmanager "ndk-bundle"

# configure environment variables
ENV ANDROID_HOME $HOME
ENV NDK_PATH $HOME/ndk-bundle
RUN gomobile init
