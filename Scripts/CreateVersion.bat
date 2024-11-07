:copy files version folder
xcopy dev %1 /E /Y /I
xcopy *.md %1/ /Y
: fix files
powershell -command "& { (Get-Content %1\README.md).Replace('v2007', '%1').Replace('2007', '%1') | Set-Content %1\README.md }"
powershell -command "& { (Get-Content %1\go.mod).Replace('/dev', '/%1') | Set-Content %1\go.mod }"
powershell -command "& { Get-ChildItem "%1\tests\*.go" | ForEach-Object -Process {(Get-Content $_).Replace('/dev/', '/%1/') | Set-Content $_ } }"
powershell -command "& { Get-ChildItem "%1\examples\*" | ForEach-Object -Process {(Get-Content $_).Replace('/dev/', '/%1/') | Set-Content $_ } }"
powershell -command "& { Get-ChildItem "%1\api\*.go" | ForEach-Object -Process {(Get-Content $_).Replace('/dev/', '/%1/') | Set-Content $_ } }"

: commit if any files are changed
git add .
git diff-index --quiet HEAD --
if %errorlevel% == 1 (
    git commit -m 'Prepare to release %1'
    git push https://%2:%3@git.auckland.dynabic.com/words-cloud/words-cloud-dotnet.git master
)

