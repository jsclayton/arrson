# First recipe is the default
help:
  @just --list

build:
  CGO_ENABLED=0 go build -o bin/arr ./cmd/arr/main.go