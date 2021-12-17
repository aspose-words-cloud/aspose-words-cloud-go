cd %1
move api\words_api.go words_api.go
rd /s /q api
md api
move words_api.go api\words_api.go
go mod edit -module github.com/aspose-words-cloud/aspose-words-cloud-go-testing
go mod tidy
cd ..
c:/sdk/Scripts/RunTestsStandalone.bat %1