#!/bin/sh


go env -w GOOS=linux
cd conn-background || return
go build
mv conn-setup ../
cd ..
zip -r usc-conn_linux.zip edge-plugin/ conn-setup
rm -f conn-setup

