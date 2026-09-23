#!/usr/bin/env bash
set -euo pipefail

sdk_root=$(cd -- "$(dirname -- "${BASH_SOURCE[0]}")/.." && pwd)
test_folder=${1:?A version directory is required}
cd -- "$sdk_root/$test_folder"

mv api/words_api.go words_api.go
rm -rf api
mkdir api
mv words_api.go api/words_api.go
go mod edit -module github.com/aspose-words-cloud/aspose-words-cloud-go-testing
go mod tidy
bash "$sdk_root/Scripts/RunTestsStandalone.sh" "$test_folder"
