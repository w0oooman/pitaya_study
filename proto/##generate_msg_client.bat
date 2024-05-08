echo off

@rem set COPYOPT= /D /Y

SET PROTO_MSG_PATH=..\Client\Assets\Script\Application\Network\Msg\ProtoMsg
SET MSG_WRAP_PATH=..\Client\Assets\Script\Application\Network\Msg\MsgWrap

echo generate proto classes files from proto files...
dir /B ".\Common\" > Common.list
dir /B ".\Client\" > Client.list

@rem for code generator we should generate some c# files
FOR /F "tokens=1 delims=, " %%i in (Common.list) do call .\protoc.exe .\Common\%%i --proto_path=.\Common --csharp_out=%PROTO_MSG_PATH% 
FOR /F "tokens=1 delims=, " %%i in (Client.list) do call .\protoc.exe .\Client\%%i --proto_path=.\Common --proto_path=.\Client --csharp_out=%PROTO_MSG_PATH%

del *.list null

pause