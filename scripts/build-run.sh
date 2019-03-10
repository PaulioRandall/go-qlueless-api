#!/bin/bash

app=qlueless-assembly-line-api

cd "${GOPATH}/src/github.com/PaulioRandall/${app}/scripts"
./build.sh

cd "${GOPATH}/bin"
./${app}