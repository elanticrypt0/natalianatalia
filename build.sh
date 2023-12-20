#!/bin/bash

# binaries
# GOOS=windows GOARCH=amd64 go build -ldflags "-w -s"
GOOS=linux GOARCH=amd64 go build -ldflags "-w -s"

# dirs
rm -rf ./build
mkdir ./build
mkdir ./build/_db
mkdir ./build/logs
mkdir ./build/public

#copy files and dirs
# cp appconfig.toml ./build
cp -r ./config ./build/config
cp -r ./seeds ./build/seeds

#web user interface
cp -R wui/dist/* ./build/public

# add execution perms
chmod +x gasonline

# windows only
# mv gasonline.exe ./build
mv gasonline ./build