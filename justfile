#!/usr/bin/env just --justfile

run:
    go run .

build:
    go build .

update:
  go get -u
  go mod tidy -v