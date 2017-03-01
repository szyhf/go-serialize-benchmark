#!/bin/sh
go test ./... -test.bench ./... -benchtime 1s -benchmem >> bench.out