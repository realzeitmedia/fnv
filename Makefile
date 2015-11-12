.PHONY: all test install

all: test install

test:
	go test

install:
	go install
