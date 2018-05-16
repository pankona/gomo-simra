
all:
	make -C $(CURDIR)/simra

misspell:
	@find . -type f -name '*.go' | grep -v vendor/ | xargs misspell -w -error

gometalinter:
	@gometalinter.v2 --vendor ./...