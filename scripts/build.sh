#!/bin/bash

app=qlueless-assembly-line-api

cd "${GOPATH}/src/github.com/PaulioRandall/${app}/cmd"
go build -o "${GOPATH}/bin/${app}" "./${app}.go"
