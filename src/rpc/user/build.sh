#!/usr/bin/env bash
RUN_NAME="src.user"

mkdir -p output/bin
cp script/* output/
chmod +x output/bootstrap.sh

if [ "$IS_SYSTEM_TEST_ENV" != "1" ]; then
    CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o output/bin/${RUN_NAME} .
else
    go test -c -covermode=set -o output/bin/${RUN_NAME} -coverpkg=./...
fi

