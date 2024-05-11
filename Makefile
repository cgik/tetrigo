dev:
	@go run main.go

test:
	@go test -v ./...

tidy:
	@go mod tidy

build:
	@go build -o build/tetrigo main.go

wasm:
	@GOOS=js GOARCH=wasm go build -o web/public/static/tetrigo.wasm cmd/wasm/main.go
	@cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" web/public/static/wasm_exec.js
