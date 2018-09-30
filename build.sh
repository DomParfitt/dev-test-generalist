#!/bin/bash
build_path=$(pwd)

mkdir -p $build_path/bin

cd $build_path/bin
rm -rf *.exe

cd $build_path/src/cmd
go build main.go

mv main.exe $build_path/bin

cd $build_path