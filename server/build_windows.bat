@echo off
set Dist=dist_windows\

if exist %Dist% (
    rem dist exist
    rmdir /q /s %Dist%
    md %Dist%
) else (
    rem not exist
    md %Dist%
)

echo go build now ...
go build

@echo off
echo copy to dist
copy .\server.exe %Dist%
del .\server.exe
md %Dist%\Config
xcopy .\Config\*.* %Dist%\Config

cls
echo build on windows ok!
ping -n 6 127.1 >nul