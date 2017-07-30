#!/bin/bash

cd ./sample1    && go build && cd -
cd ./sample2    && go build && cd -
cd ./sample3    && go build && cd -
cd ./immortal   && go build && cd -
cd ./animation1 && go build && cd -
cd ./animation2 && go build && cd -
cd ./audio      && go build && cd -

