#!/bin/bash

# TODO: check if it's being run from the root directory

go build -o uploader -i ./cmd/go-ride-fare-estimation/main.go

./go-ride-fare-estimation