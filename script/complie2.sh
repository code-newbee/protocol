#!/bin/bash

protoPath=geeker
protoc \
          -I ./ \
          -I /usr/local/include \
          -I "$GOPATH"/src \
          --go_out=plugins=grpc:"$protoPath" \
          "$protoPath"/"$protoPath".proto