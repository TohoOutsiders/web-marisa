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

SET CGO_ENABLED=0

echo go build now ...
go build -o server.exe

rem @echo off
echo copy to dist
copy .\server.exe %Dist%
del .\server.exe
md %Dist%\Config
xcopy .\Config\*.* %Dist%\Config
del %Dist%\Config\.gitkeep

cls
echo build on windows ok!
ping -n 6 127.1 >nul
