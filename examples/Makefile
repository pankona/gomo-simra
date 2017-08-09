# Makefile

DIRECTORIES = $(dir $(wildcard $(CURDIR)/*/.))
BUILD_COMMAND=go

all: $(DIRECTORIES)

mobile:
	make BUILD_COMMAND=gomobile 

build: $(DIRECTORIES)
$(DIRECTORIES):
	cd $@; $(BUILD_COMMAND) build

.PHONY: all $(DIRECTORIES)
