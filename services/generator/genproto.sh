#!/usr/bin/env bash

set -e

PATH=$PATH:$GOPATH/bin
protodir=../pb

protoc --go_out=plugins=grpc:genproto -I $protodir $protodir/crds.proto