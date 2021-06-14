# 变量定义
protoPath=""
protocol="gRpc"
lang="go"
workspace="/Users/ark/Code/Go/grpc/code-newbee/protocol"

if [ "$1" == "" ]
  then
    echo "proto path empty"
    exit
  else
    protoPath="$1"
fi

if [ "$2" != "" ]
  then
  protocol="$2"
fi

if [ "$3" != "" ]
  then
  lang="$3"
fi

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
            --java_out="$protoPath" \
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