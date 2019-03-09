#!/bin/bash

app=qlueless-assembly-line-api

cd "${GOPATH}/src/${app}/cmd"
go build -o "${GOPATH}/bin/${app}" "./${app}.go"
