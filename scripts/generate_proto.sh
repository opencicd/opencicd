#!/bin/bash

sed -e "s:GOPATH:$GOPATH/src:g" -e "s:GOMOD:$GOPATH/pkg/mod:g" < prototool.json > .prototool.json

echo ls -GOPATH
ls $GOPATH

echo ls -GOPATH/pkg
ls $GOPATH/pkg

echo ls -GOPATH/pkg/mod
ls $GOPATH/pkg/mod

echo ls -GOPATH/pkg/mod/github.com/gogo/protobuf/gogoproto
ls $GOPATH/pkg/mod/github.com/gogo/protobuf/gogoproto

echo ls -GOPATH/src/github.com/gogo/protobuf/gogoproto
ls $GOPATH/src/github.com/gogo/protobuf/gogoproto

cat .prototool.json

prototool all --config-data "$(cat .prototool.json)"
