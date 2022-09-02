@echo off
windres main.rc -o main.syso
go generate
go build -ldflags "-H windowsgui -w -s" -o json2go.exe
