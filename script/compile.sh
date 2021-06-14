#!/bin/bash

# 变量定义
protoPath=""
protocol="gRpc"
lang="go"
workspace="/Users/ark/Code/Go/grpc/code-newbee/protocol"

# p protoPath
# t protocol http, grpc
# l lang java go
while getopts "p:t:l:" opt; do
  case $opt in
      p)
        if [ "$OPTARG" == "" ];
          then
            echo "proto path empty"
            exit
          else
            protoPath="$OPTARG"
        fi
        ;;
      t)
        protocol="$OPTARG"
        ;;
      l)
        lang="$OPTARG"
        ;;
      *)
        echo "-$opt not recognized"
        exit
        ;;
  esac
done

cd "$workspace" || exit

case "$lang" in
  "go")
  protoc \
            -I ./ \
            -I /usr/local/include \
            -I "$GOPATH"/src \
            --go_out=plugins=grpc:"$protoPath" \
            "$protoPath"/"$protoPath".proto \
  ;;

  "java")
    protoc \
            -I ./ \
            -I /usr/local/include \
            -I "$GOPATH"/src \
            --java_out=plugins=grpc:"$protoPath" \
            "$protoPath"/"$protoPath".proto \
  ;;
esac

case "$protocol" in
  "http")
  # add gRpc-gateway pb if its http svr
  protoc \
          -I ./ \
          -I /usr/local/include \
          -I "$GOPATH"/src \
          --grpc-gateway_out=logtostderr=true:"$protoPath" \
          "$protoPath"/"$protoPath".proto
  ;;
esac