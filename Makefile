VERSION ?= $(shell git rev-parse --abbrev-ref HEAD)-$(shell git describe --long --always)
BUILD_TIME ?=`date +%FT%T%z`

all: build

compile:
	go build  -ldflags="-X main.BuildTime=${BUILD_TIME} -X main.Version=${VERSION}" -o newsy main.go

build: compile

run: build
	./newsy