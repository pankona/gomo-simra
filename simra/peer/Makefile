
all:
	@make --no-print-directory test
	@make --no-print-directory cover
	@make --no-print-directory metalinter

test:
	@go test ./...

cover:
	@go test -coverprofile=.cover.out ./...
	@go tool cover -html=.cover.out -o .cover.html

metalinter:
	@echo gometalinter is running...
	@gometalinter ./...
	@echo gometalinter finished!
