set CURDIR=%~dp0
cd ..\example\proto
protoc --go_out=.. *.proto
cd %CURDIR%
python redis2go.py --proto_path=..\example\proto --go_out=..\example

pause