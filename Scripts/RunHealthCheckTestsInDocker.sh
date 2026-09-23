#!/usr/bin/env bash
set -euo pipefail

sdk_root=$(cd -- "$(dirname -- "${BASH_SOURCE[0]}")/.." && pwd)
test_folder=${1:?A version directory is required}

docker run --rm \
    --user "$(id -u):$(id -g)" \
    --volume "$sdk_root:/sdk" --workdir /sdk \
    --env ClientId --env ClientSecret --env apiUrl \
    --env HOME=/tmp --env GOPATH=/tmp/go --env GOCACHE=/tmp/go-build \
    golang:1.18.10 bash -c '
        set -euo pipefail
        export PATH="/tmp/go/bin:$PATH"
        echo "{\"ClientId\":\"${ClientId:-}\",\"ClientSecret\":\"${ClientSecret:-}\",\"BaseUrl\":\"${apiUrl:-}\" }" > config.json
        bash /sdk/Scripts/RunHealthCheckTestsStandalone.sh "$1"
    ' bash "$test_folder"
