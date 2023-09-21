# 百灵快传（B0Pass）

LAN large file transfer tool。

基于Go语言的高性能 “手机电脑超大文件传输神器”、“局域网共享文件服务器”。

只需一个文件（exe）双击开启。

## 1. 主要功能

### 1.1 功能描述

- [x] 局域网文件共享服务器
- [x] 简单的单个可执行文件
- [x] 共享文件界面（在同一局域网或WIFI下，传输超大文件）
- [x] 二维码扫码界面（支持手机传输，支持其它电脑输入网址）
- [x] 共享文件在线管理界面（主电脑打开、图片浏览器、重命名、删除等）
- [x] 更简洁高效的操作界面
- [x] 使用自研的 <a href="//github.com/bitepeng/b0boot-go" target="_blank">B0Boot-Go</a> 框架重构代码，更简洁、更模块化
- [x] 文件上传界面支持多次选择（PC端支持拖拽上传）
- [x] 大文件上传过程分片处理（上传更丝滑，不卡顿）
- [x] 支持路径有空格或中文的情况
- [x] 支持Windows、Linux、MacOS操作系统
- [x] 支持端口（port）自定义配置
- [x] 支持域名（domain）自定义配置
- [x] 支持安卓手机APK应用（Uniapp开发的手机端）
- [ ] 支持安全代码（code）自定义配置（增强安全性控制）
- [ ] 支持对文件多选和全选操作
- [ ] 支持对上传的压缩包在线解压
- [ ] 支持上传整个文件夹
- [ ] 自动检查更新版本

### 1.2 PC操作截图

<table width="100%">
<tr>
<td width="33%">
    <img src="https://gitee.com/b0cloud/b0pass/raw/master/zdoc/_images/pc/01.png" width="100%"/>
    <p>主界面（功能说明）</p>
</td>
<td width="33%">
    <img src="https://gitee.com/b0cloud/b0pass/raw/master/zdoc/_images/pc/02.png" width="100%"/>
    <p>主界面（图文模式、文件菜单）</p>
</td>
<td width="33%">
    <img src="https://gitee.com/b0cloud/b0pass/raw/master/zdoc/_images/pc/03.png" width="100%"/>
    <p>主界面（列表模式）</p>
</td>
</tr>
<tr>
<td width="33%">
    <img src="https://gitee.com/b0cloud/b0pass/raw/master/zdoc/_images/pc/04-1.png" width="100%"/>
    <p>手机扫码（到主界面）</p>
</td>
<td width="33%">
    <img src="https://gitee.com/b0cloud/b0pass/raw/master/zdoc/_images/pc/04-2.png" width="100%"/>
    <p>手机扫码（到某个文件）</p>
</td>
<td width="33%">
    <img src="https://gitee.com/b0cloud/b0pass/raw/master/zdoc/_images/pc/05.png" width="100%"/>
    <p>大文件上传（选择文件）</p>
</td>
</tr>
<tr>
<td width="33%">
    <img src="https://gitee.com/b0cloud/b0pass/raw/master/zdoc/_images/pc/06.png" width="100%"/>
    <p>大文件上传（上传完成）</p>
</td>
<td width="33%">
    <img src="https://gitee.com/b0cloud/b0pass/raw/master/zdoc/_images/pc/07-1.png" width="100%"/>
    <p>图片浏览器</p>
</td>
<td width="33%">
    <img src="https://gitee.com/b0cloud/b0pass/raw/master/zdoc/_images/pc/07-2.png" width="100%"/>
    <p>Html文件 Web服务</p>
</td>
</tr>
</table>

### 1.3 手机操作截图

<table width="100%">
<tr>
<td width="33%">
    <img src="https://gitee.com/b0cloud/b0pass/raw/master/zdoc/_images/phone/01.jpg" width="100%"/>
    <p>主界面</p>
</td>
<td width="33%">
    <img src="https://gitee.com/b0cloud/b0pass/raw/master/zdoc/_images/phone/02.jpg" width="100%"/>
    <p>上传文件</p>
</td>
<td width="33%">
    <img src="https://gitee.com/b0cloud/b0pass/raw/master/zdoc/_images/phone/03.jpg" width="100%"/>
    <p>文件操作</p>
</td>
</tr>
</table>


## 2. 发行版下载使用

### 下载前请阅读
-  已编译好***Windows、MacOS、Linux***等平台的可执行文件
-  只需下载到电脑，双击开启即可使用
-  注意1：建议不要安装到“C:\Program Files (x86)”等系统目录，否则需要右键“以管理员身份运行”
-  注意2：如果要自定义端口等配置，请修改`config.ini`文件
```
[gateway]
ListenAddr = ":8888"      # 配置IP和端口
Domain = "test.com:8888"  # 可选配置，配置访问域名（若使用80端口或nginx代理，可忽略端口）
[pass]
Path = "files"            # 文件管理根目录
```

### 最新版下载地址
- https://4bit.cn/p/b0pass    (项目官网，直接下载)

## 3. 代码仓库
- https://github.com/bitepeng/b0pass   GitHub（主库）   欢迎star支持
- https://gitee.com/b0cloud/b0pass     GitEE（国内同步） 欢迎star支持


## 4. 使用场景
- ***手机电脑共享文件***

    电脑上双击执行 -> 手机扫码 -> 手机、电脑文件可以互传。

- ***电脑之间共享文件***

    电脑A上双击执行 -> 电脑B上浏览器输入A的地址 -> 电脑A、电脑B文件可以互传。

- ***虚拟机和电脑之间共享文件***

    电脑上双击执行 -> 虚拟机上浏览器输入电脑的地址 -> 虚拟机、电脑文件可以互传。

- ***更多使用场景***

    也可以用作“家庭影音中心”、“办公室文件共享”、“产品原型服务器”等。走局域网HTTP协议，Windows、MacOS、虚拟机、安卓、iPhone等都可以跨平台共享文件。

## 5. 源码编译
```
# 下载代码
git clone https://github.com/bitepeng/b0pass.git

# 使用go mod更新依赖
go mod tidy

# 开发运行
cd main && go run ./main.go

# 编译运行开发版本
cd main && ./build.bat
```
