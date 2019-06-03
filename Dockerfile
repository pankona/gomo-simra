FROM circleci/golang:1.12.5

ENV HOME /home/circleci
WORKDIR $HOME

ENV PATH $GOPATH/bin:$PATH

RUN sudo apt-get update
RUN sudo apt-get install -y libegl1-mesa-dev
RUN sudo apt-get install -y libgles2-mesa-dev
RUN sudo apt-get install -y libx11-dev
RUN sudo apt-get install -y libasound2-dev
RUN go get -u golang.org/x/mobile/cmd/gomobile
RUN go get -u github.com/golang/freetype/truetype
RUN go get -u github.com/boltdb/bolt
RUN go get -u github.com/hajimehoshi/oto
RUN go get -u github.com/hajimehoshi/go-mp3
RUN go get -u github.com/golang/lint/golint
RUN curl -L https://git.io/vp6lP | sh # install gometalinter
RUN mv ./bin/* $GOPATH/bin/.

# install android sdk
ENV ANDROID_SDK_VERSION=4333796
RUN wget https://dl.google.com/android/repository/sdk-tools-linux-$ANDROID_SDK_VERSION.zip
RUN unzip -q sdk-tools-linux-$ANDROID_SDK_VERSION.zip
RUN sudo apt-get install -y openjdk-8-jdk
RUN yes | $HOME/tools/bin/sdkmanager "ndk-bundle"

# configure environment variables
ENV ANDROID_HOME $HOME
ENV NDK_PATH $HOME/ndk-bundle
RUN gomobile init
