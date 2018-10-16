.PHONY: all clean

export CGO_ENABLED=0
export GOOS=linux
export GOARCH=amd64

BINARY=riot
VERSION=0.1.0
BUILD=$(shell git rev-parse --short HEAD)
LDFLAGS=-ldflags "-X main.Version=$(VERSION) -X main.Build=$(BUILD)"

all:
	go build -o $(BINARY) $(LDFLAGS)

fmt:
	find . -name '*.go' -not -wholename './vendor/*' | while read -r file; do gofmt -w -s "$$file"; goimports -w "$$file"; done

docker:
	docker build -t "ishantanu16/$(BINARY):$(VERSION)" \
		--build-arg build=$(BUILD) --build-arg version=$(VERSION) \
		-f Dockerfile .

clean:
	-rm $(BINARY)
