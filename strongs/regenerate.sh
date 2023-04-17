#!/bin/bash -ex
# -*- compile-command: "./regenerate.sh"; -*-
go run ../cmd/strongs2go/main.go h*.html > hebrew.go
go run ../cmd/strongs2go/main.go g*.html > greek.go
