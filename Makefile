GOARCH=amd64
BINARY=disassemble8080

build: linux darwin windows

linux:
	GO111MODULE=on CGO_ENABLED=0 GOARCH=${GOARCH} GOOS=linux go build -ldflags="-s -w" -o bin/${BINARY}-linux-${GOARCH} ./cmd/disassemble8080/main.go

darwin:
	GO111MODULE=on CGO_ENABLED=0 GOARCH=${GOARCH} GOOS=darwin go build -ldflags="-s -w" -o bin/${BINARY}-darwin-${GOARCH} ./cmd/disassemble8080/main.go

windows:
	GO111MODULE=on CGO_ENABLED=0 GOARCH=${GOARCH} GOOS=windows go build -ldflags="-s -w" -o bin/${BINARY}-windows-${GOARCH}.exe ./cmd/disassemble8080/main.go

.PHONY: build