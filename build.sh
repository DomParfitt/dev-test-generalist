#!/bin/bash
build_path=$(pwd)

cd $build_path/bin
rm -rf *.exe

cd $build_path/src/cmd
go build main.go

mv main.exe $build_path/bin