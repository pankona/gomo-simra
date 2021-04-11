# Makefile

DIRECTORIES   = $(dir $(wildcard $(CURDIR)/*/.))
BUILD_COMMAND = go
BUILD_TAGS    =

all: $(DIRECTORIES)

mobile:
	make BUILD_COMMAND=gomobile 

build: $(DIRECTORIES)
$(DIRECTORIES):
	@go version
	cd $@; go get && $(BUILD_COMMAND) build -tags=$(BUILD_TAGS)

.PHONY: all $(DIRECTORIES)
