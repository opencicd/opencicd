#!/bin/bash

sed "s:GOPATH:$GOPATH/src:g" prototool.json > .prototool.json
prototool all --config-data "$(cat .prototool.json)"