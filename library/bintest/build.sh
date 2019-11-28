#!/usr/bin/env bash

# 设置项目名称
APPNAME="bintest"

# 设置GOPATH临时值
cd ../../../../
export GOPATH=${PWD}
export GO111MODULE=on

# 进入工作目录
echo ${GOPATH}
cd src/b0pass/library/${APPNAME}

# 打包可执行文件
##### mac os #####
CGO_ENABLED="0" GOARCH="amd64" GOOS="darwin" go build -mod=vendor -o ${GOPATH}/bin/${APPNAME}/${APPNAME}_mac_cli main.go

##### win32 os ##### -ldflags "-H windowsgui"
CGO_ENABLED="0" GOARCH="386" GOOS="windows" go build -mod=vendor -o ${GOPATH}/bin/${APPNAME}/${APPNAME}_win32_cli.exe  main.go

##### linux os #####
CGO_ENABLED="0" GOARCH="amd64" GOOS="linux" go build -mod=vendor -o ${GOPATH}/bin/${APPNAME}/${APPNAME}_linux_cli main.go

find ${GOPATH}/bin/${APPNAME}