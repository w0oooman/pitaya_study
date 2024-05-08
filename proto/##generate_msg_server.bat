echo off

@rem set COPYOPT= /D /Y
SET PROTO_MSG_PATH_GO=.\pb\go
SET PROTO_MSG_PATH_C#=.\pb\c#

echo 1. generate proto classes files from proto files...
dir /B ".\Common\" > Common.list
dir /B ".\Client\" > Client.list
dir /B ".\Server\" > Server.list

@rem for code generator we should generate some c# files
FOR /F "tokens=1 delims=, " %%i in (Common.list) do (
	.\protoc.exe .\Common\%%i --proto_path=.\Common --go_out=%PROTO_MSG_PATH_GO% --csharp_out=%PROTO_MSG_PATH_C#%
)
FOR /F "tokens=1 delims=, " %%i in (Client.list) do ( 
	.\protoc.exe .\Client\%%i --proto_path=.\Common --proto_path=.\Client --go_out=%PROTO_MSG_PATH_GO%
)
FOR /F "tokens=1 delims=, " %%i in (Server.list) do (
	.\protoc.exe .\Server\%%i --proto_path=.\Common --proto_path=.\Server --proto_path=.\Client --go_out=%PROTO_MSG_PATH_GO%
)

del *.list null

echo.
echo ### all tasks finished. ###
echo.

pause