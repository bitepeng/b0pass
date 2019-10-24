#!/usr/bin/env bash

# 进入工作目录
cd ../../

# 打包静态资源
echo y | gf pack public,template,config boot/resource.go -n=boot
# echo y | gf pack public,template,config resource.go -n=main