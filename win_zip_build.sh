#!/bin/sh

go env -w GOOS=windows
cd conn-background || return
go build
mv conn-setup.exe ../
cd ..
zip -r usc-conn_linux.zip edge-plugin/ conn-setup.exe
rm -f conn-setup.exe
go env -w GOOS=linux