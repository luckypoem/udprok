.PHONY: default udprokd udproks udprokc clean all
export GOPATH:=$(shell pwd)

default: all

all: udprokd udproks udprokc

udprokd:
	go install main/udprokd

udproks: simplejson
	go install main/udproks

udprokc: simplejson
	go install main/udprokc

clean:
	go clean -i -r ./

simplejson: 
	go get github.com/bitly/go-simplejson