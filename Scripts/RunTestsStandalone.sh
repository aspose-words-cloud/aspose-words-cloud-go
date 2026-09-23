#!/usr/bin/env bash
set -euo pipefail

sdk_root=$(cd -- "$(dirname -- "${BASH_SOURCE[0]}")/.." && pwd)
cd -- "$sdk_root/${1:?A test directory is required}"

go get -u github.com/jstemmer/go-junit-report
# Go 1.18 no longer installs executables as part of go get.
go install github.com/jstemmer/go-junit-report
go test ./tests/... -v -timeout 1h 2>&1 | go-junit-report > "$sdk_root/testReport.xml"
