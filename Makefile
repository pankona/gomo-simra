
all:
	make -C $(CURDIR)/simra

misspell:
	@find . -type f -name '*.go' | grep -v vendor/ | xargs misspell -w -error

lint:
	golangci-lint run --deadline=300s

docker-build:
	docker build . -t pankona/gomo-simra

circleci:
	circleci local execute

.PHONY: docker-build
