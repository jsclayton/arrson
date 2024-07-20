# First recipe is the default
help:
  @just --list

build:
  CGO_ENABLED=0 go build -o bin/arrson ./cmd/arrson/main.go