for /f %%i in ('dir v* /b/a:d/od/t:c') do set LAST=%%i
echo {"AppSid":"%WordsAppSid%","AppKey":"%WordsAppKey%","BaseUrl":"%WordsBaseUrl%" } > %LAST%/config.json
docker run -v %cd%/%LAST%:c:/sdk -w="c:/sdk" --rm -t golang:1.14.0-windowsservercore-1809 go test ./tests/... -v