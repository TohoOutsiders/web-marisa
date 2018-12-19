@echo off
set Dist=dist_linux\

if exist %Dist% (
    rem dist exist
    rmdir /q /s %Dist%
    md %Dist%
) else (
    rem not exist
    md %Dist%
)

SET CGO_ENABLED=0
set GOARCH=amd64
set GOOS=linux

echo go build now ...
go build

rem @echo off
echo copy to dist
copy .\server %Dist%
del .\server
md %Dist%\Config
xcopy .\Config\*.* %Dist%\Config
del %Dist%\Config\.gitkeep

cls
echo build on linux ok!
ping -n 6 127.1 >nul
