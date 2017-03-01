#!/bin/sh
go test -v ./... -test.bench ./... -benchtime 1s -benchmem