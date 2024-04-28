dev:
	@go run main.go

test:
	@go test -v ./...

tidy:
	@go mod tidy

build:
	@go build -o bin/tetrigo main.go

wasm:
	@GOOS=js GOARCH=wasm go build -o web/public/static/tetrigo.wasm cmd/wasm/main.go
