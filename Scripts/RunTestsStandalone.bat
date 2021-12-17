cd %1
go get -u github.com/jstemmer/go-junit-report
go test ./tests/... -v -timeout 1h 2>&1 | go-junit-report > ../testReport.xml