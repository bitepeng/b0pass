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
APP="${GOPATH}/bin/${APPNAME}/${APPNAME}_OSX/${APPNAME}.app"
mkdir -p ${APP}/Contents/{MacOS,Resources}
##### mac os #####
#CGO_ENABLED="0" GOARCH="amd64" GOOS="darwin" go build -mod=vendor -o ${GOPATH}/bin/${APPNAME}/${APPNAME}_mac cli.go
CGO_ENABLED="0" GOARCH="amd64" GOOS="darwin" go build -mod=vendor -o ${APP}/Contents/MacOS/${APPNAME}_mac_ui main.go

cat > ${APP}/Contents/Info.plist << EOF
<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
	<key>CFBundleExecutable</key>
	<string>${APPNAME}_mac_ui</string>
	<key>CFBundleIconFile</key>
	<string>icon.icns</string>
	<key>CFBundleIdentifier</key>
	<string>com.b0cloud.${APPNAME}</string>
</dict>
</plist>
EOF
cp docs/icons/icon.icns ${APP}/Contents/Resources/icon.icns

##### win32 os ##### -ldflags "-H windowsgui"
CGO_ENABLED="0" GOARCH="386" GOOS="windows" go build -mod=vendor -o ${GOPATH}/bin/${APPNAME}/${APPNAME}_win32.exe main.go

##### linux os #####
# CGO_ENABLED="0" GOARCH="amd64" GOOS="linux" go build -mod=vendor -o ${GOPATH}/bin/${APPNAME}/${APPNAME}_linux cli.go

find ${GOPATH}/bin/${APPNAME}