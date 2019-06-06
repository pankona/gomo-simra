
all:
	@make --no-print-directory cover

test:
	@go test ./...

cover:
	@go test -coverprofile=.cover.out ./...
	@go tool cover -html=.cover.out -o .cover.html
