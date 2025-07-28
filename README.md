# Code Retreat Skeleton Code

This repository contains skeleton code that Code Retreat Participants can use to get started.

## Go

```shell
echo "{\"PleasePath\":\"$PWD/pleasew\"}" > puku.json

# Update BUILD files
./pleasew puku fmt //go/...

# Run the binary
./pleasew run //go/cmd/conway

# Run the tests
./pleasew test //go/...

# Install third party packages
./pleasew build //third_party/go:toolchain
GO="$PWD/$(./pleasew query outputs //third_party/go:toolchain)/bin/go"
(cd third_party/go && $GO get github.com/stretchr/testify@latest) && ./pleasew puku sync -w
```

## Python

```shell
# Run the binary
./pleasew run //python:main

# Run the tests
./pleasew test //python/...
```
