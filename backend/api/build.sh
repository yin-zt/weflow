#!/usr/bin/env bash
RUN_NAME="flow"

mkdir -p output/bin output/conf
cp script/* output/
cp -r conf/* output/conf
chmod +x output/bootstrap.sh
go build -o output/bin/${RUN_NAME}