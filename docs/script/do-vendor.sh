#!/usr/bin/env bash

# 设置GOPATH临时值
cd ../../../../
export GOPATH=${PWD}
export GOPROXY=https://goproxy.io
export GO111MODULE=on

# 进入工作目录
echo ${GOPATH}
cd src/b0pass


# 使用vendor依赖
go mod download
go mod vendor
go mod tidy
