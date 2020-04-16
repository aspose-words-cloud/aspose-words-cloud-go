for /f %%i in ('dir v* /b/a:d/od/t:c') do set LAST=%%i
cd %LAST%
go get -u github.com/jstemmer/go-junit-report
go test ./tests/... -v -timeout 20m 2>&1 | go-junit-report > ../testReport.xml