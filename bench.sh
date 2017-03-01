#!/bin/sh
if [ ! $1 = "print" ];then
	go test ./... -test.bench ./... -benchtime 1s -benchmem > bench.out
else
	go test ./... -test.bench ./... -benchtime 1s -benchmem
fi