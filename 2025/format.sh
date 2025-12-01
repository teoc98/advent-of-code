#!/bin/sh

DIR=$(dirname "$0")
gofmt -l -s -w "$DIR"

