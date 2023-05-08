#!/usr/bin/env just --justfile


run:
    go run .

build:
    go build .

update:
  go get -u
  go mod tidy -v


push message="Update":
    git add .
    git commit -m {{message}}
    git push origin master

# release, e.g `just release v0.12`
release version:
    sed -i 's/Version = ".*"/Version = "{{version}}"/' ./internal/cons/cons.go
    git tag {{version}}
    git add .
    git commit -m "release {{version}}"
    git push
    git push --tags


