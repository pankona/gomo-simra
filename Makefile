
all:
	make test -C $(CURDIR)
	make lint -C $(CURDIR)
	make build-sample -C $(CURDIR)

build-sample:
	make -C $(CURDIR)/examples

build-sample-mobile-on-docker:
	docker run \
		-it \
		-v $(CURDIR):/app \
		-v gomo-simra-go-module-tmp:/go \
		-w /app \
		pankona/gomo-simra \
		make mobile -C /app/examples

test:
	go test ./...

lint:
	golangci-lint run --deadline=300s

docker-build:
	docker build . -t pankona/gomo-simra

docker-build-no-cache:
	docker build . --no-cache -t pankona/gomo-simra

circleci:
	circleci local execute

.PHONY: build-sample \
	build-sample-mobile \
	test \
	lint \
	docker-build \
	docker-build-no-cache \
	circleci

