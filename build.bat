@echo off
set build_path=%cd%

if not exist %build_path%\bin mkdir %build_path%\bin

cd %build_path%\bin
del *.exe

cd %build_path%\src\cmd
go build main.go

move main.exe %build_path%\bin

cd %build_path%