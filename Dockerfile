FROM circleci/golang:1.16.2

ENV HOME /home/circleci
WORKDIR $HOME

ENV PATH $GOPATH/bin:$PATH

RUN wget https://dl.google.com/dl/android/studio/ide-zips/4.1.3.0/android-studio-ide-201.7199119-linux.tar.gz
RUN tar zxf android-studio-ide-201.7199119-linux.tar.gz

ENV JAVA_HOME $HOME/android-studio/jre
ENV PATH $PATH:$JAVA_HOME/bin

RUN sudo apt-get update
RUN sudo apt-get install -y libegl1-mesa-dev
RUN sudo apt-get install -y libgles2-mesa-dev
RUN sudo apt-get install -y libx11-dev
RUN sudo apt-get install -y libasound2-dev
RUN curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh -s -- -b $(go env GOPATH)/bin v1.39.0

RUN wget https://dl.google.com/android/repository/android-ndk-r21e-linux-x86_64.zip
RUN unzip -q android-ndk-r21e-linux-x86_64.zip
ENV ANDROID_NDK_HOME ${HOME}/android-ndk-r21e

ENV GO111MODULE auto
RUN go get golang.org/x/mobile/cmd/gomobile
RUN gomobile init
