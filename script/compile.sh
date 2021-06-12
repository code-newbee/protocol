if [ "$1" == "" ]; then
    echo "proto path empty"
    exit
fi
cd ../ && protoc -I ./ "$1"/geeker.proto --go_out=plugins=grpc:"$1"