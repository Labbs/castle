#!/bin/bash

# for d in $(ls -d */);
# do
#     cd $d
#     for f in $(ls *.proto);
#     do
#         protoc --proto_path=. --go_out=paths=source_relative:. $f
#         protoc-go-inject-tag -XXX_skip=yaml,xml,xorm -input=./$(echo $f | cut -d. -f1).pb.go
#     done
#     cd ..
# done

if [ ! -d "../gen" ]; then
    mkdir ../gen
fi

protoc --proto_path=../proto --go_out=paths=source_relative:../gen $(find ../proto -name "*.proto")