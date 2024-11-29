EXECUTABLES = git go find pwd
K := $(foreach exec,$(EXECUTABLES),\
        $(if $(shell which $(exec)),some string,$(error "No $(exec) in PATH)))

ROOT_DIR:=$(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))

BINARY=ImapDownloader
VERSION=0.0.2
BUILD=$(shell date +"%Y-%m-%d_%H:%M:%S")
PLATFORMS=darwin linux windows
ARCHITECTURES=386 amd64

# Setup linker flags option for build that interoperate with variable names in src code
LDFLAGS=-ldflags "-w -s -X main.version=${VERSION} -X main.build=${BUILD}"

default: tidy build

all: clean tidy build_all install

build:
	go build ${LDFLAGS} -o ${BINARY} main.go

build_all:
	$(foreach GOOS, $(PLATFORMS),\
	$(foreach GOARCH, $(ARCHITECTURES), $(shell export GOOS=$(GOOS); export GOARCH=$(GOARCH); go build -v ${LDFLAGS} -o $(BINARY)-$(GOOS)-$(GOARCH))))

install:
	go install ${LDFLAGS}

tidy: 
	go mod tidy

# Remove only what we've created
clean:
	find ${ROOT_DIR} -name '${BINARY}[-?][a-zA-Z0-9]*[-?][a-zA-Z0-9]*' -delete
	rm -rf ${BINARY}

lint:
	golangci-lint run

.PHONY: check clean install build_all all
