FROM circleci/golang:1.11.1

ENV HOME     /home/circleci
WORKDIR      $HOME

ENV NDK      android-ndk-r15c
ENV NDK_ROOT $HOME/$NDK
ENV PATH     $GOPATH/bin:$PATH

RUN sudo apt-get update
RUN sudo apt-get install -y libegl1-mesa-dev
RUN sudo apt-get install -y libgles2-mesa-dev
RUN sudo apt-get install -y libx11-dev
RUN sudo apt-get install -y libasound2-dev
RUN wget https://dl.google.com/android/repository/$NDK-linux-x86_64.zip
RUN unzip -q $NDK-linux-x86_64.zip
RUN go get -u golang.org/x/mobile/cmd/gomobile
RUN go get -u github.com/golang/freetype/truetype
RUN go get -u github.com/boltdb/bolt
RUN go get -u github.com/hajimehoshi/oto
RUN go get -u github.com/hajimehoshi/go-mp3
RUN go get -u github.com/golang/lint/golint
RUN gomobile init --ndk $NDK_ROOT
RUN curl -L https://git.io/vp6lP | sh # install gometalinter
RUN mv ./bin/* $GOPATH/bin/.
