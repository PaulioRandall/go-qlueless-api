#!/bin/bash

app=qlueless-assembly-line-api

cd "${GOPATH}/src/${app}/scripts"
./build.sh

cd "${GOPATH}/bin"
./${app}