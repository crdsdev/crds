#!/usr/bin/env bash

set -e

mkdir -p proto
cp -r ../pb/* ./proto
# protoc -I=. proto/crds.proto --js_out=import_style=commonjs:. --grpc-web_out=import_style=commonjs,mode=grpcwebtext:.