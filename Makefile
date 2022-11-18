BINARY_NAME=tesla-firmware-decrypt

build:
	go build -o "build/$(BINARY_NAME)" "./cmd/$(BINARY_NAME)/"

install:
	cp "build/$(BINARY_NAME)" "$$GOPATH/bin/$(BINARY_NAME)"