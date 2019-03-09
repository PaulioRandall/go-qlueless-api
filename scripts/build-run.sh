#!/bin/bash
@Echo off

cd "${GOPATH}/src/qlueless-assembly-line/scripts"
./build.sh

cd "${GOPATH}/src/qlueless-assembly-line/cmd"
./qlueless-assembly-line-api