#!/bin/bash

# targetFile 编译后的输出文件名称
targetFile_win64="gotomongo.exe"
targetFile_linux="gotomongo"
# 目标编译包file
pkgfile="main.go"

goBuild() {
# 编译window平台执行文件
buildResult=`GOOS=windows GOARCH=amd64 go build go build -ldflags "-s -w" -o "${targetFile_win64}" "$pkgfile" 2>&1`
# 编译linux平台执行文件
buildResult=`GOOS=linux GOARCH=386 go build -ldflags "-s -w" -o "${targetFile_linux}" "$pkgfile" 2>&1`
# 压缩二进制文件，只适用于linux版本的二进制文件。需要安装upx工具。
# buildResult=`upx -9 $targetFile_linux `

if [ -z "$buildResult" ]; then
echo "success"
fi
echo $buildResult
}

goBuild
