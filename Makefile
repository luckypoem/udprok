.PHONY: default server client clean all
export GOPATH:=$(shell pwd)

default: all

server:
	go install main/udprokd

client: simplejson
	go install main/udprok

all: client server

clean:
	go clean -i -r ./

simplejson: 
	go get github.com/bitly/go-simplejson