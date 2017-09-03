#!/bin/sh
# for go below 1.9
# go test -v ./... -test.bench ./... -benchtime 1s -benchmem
# for go 1.9
go test -v ./... -test.bench . -benchtime 1s -benchmem