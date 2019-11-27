#!/usr/bin/env bash

# 设置项目名称
APPNAME="b0pass"

# 设置GOPATH临时值
cd ../../../../
export GOPATH=${PWD}
export GO111MODULE=on

# 进入工作目录
echo ${GOPATH}
cd src/${APPNAME}

# 打包可执行文件
##### mac os #####
CGO_ENABLED="0" GOARCH="amd64" GOOS="darwin" go build -mod=vendor -o ${GOPATH}/bin/${APPNAME}/${APPNAME}_mac_cli cli.go

##### win32 os ##### -ldflags "-H windowsgui"
# CGO_ENABLED="0" GOARCH="386" GOOS="windows" go build -mod=vendor -o ${GOPATH}/bin/${APPNAME}/${APPNAME}_win32_cli.exe  cli.go

##### linux os #####
CGO_ENABLED="0" GOARCH="amd64" GOOS="linux" go build -mod=vendor -o ${GOPATH}/bin/${APPNAME}/${APPNAME}_linux_cli cli.go

find ${GOPATH}/bin/${APPNAME}