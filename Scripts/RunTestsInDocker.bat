for /f %%i in ('dir v* /b/a:d/od/t:c') do set LAST=%%i
echo {"AppSid":"%WordsAppSid%","AppKey":"%WordsAppKey%","BaseUrl":"%apiUrl%" } > %LAST%/config.json
docker run --isolation=hyperv -v %cd%:c:/sdk -w="c:/sdk/" --rm -t golang:1.14.0-windowsservercore-1809 c:/sdk/Scripts/RunTestsStandalone.bat