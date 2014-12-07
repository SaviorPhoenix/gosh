#!/bin/bash

DIRS=(./cmd ./env)
for i in ${DIRS[@]}; do
    go test -cover -coverprofile=coverage.out $i
    go tool cover -html=coverage.out
done
