# 百灵快传（B0Pass）

LAN large file transfer tool。

基于Go语言的高性能 “手机电脑超大文件传输神器”、“局域网共享文件服务器”。

只需一个文件（exe）双击开启。

## 1. 主要功能

### 1.1 功能描述

- [x] 文件共享服务器
- [x] 简单的单个可执行文件
- [x] 共享文件界面（只要在同一局域网或WIFI下，可以传输超大文件）
- [x] 上传文件界面（支持点选和拖拽）
- [x] 二维码扫码界面（支持手机传输，支持其它电脑输入网址）
- [x] 共享文件在线管理界面（可删除）
- [x] 开发linux可部署版本
- [x] 支持命令行端口自定义
- [ ] 端口如果被使用，可以自行开启其它端口 
- [ ] 使用WebSocket实时通知文件变更
- [ ] 更简洁高效的操作界面
- [ ] 自动检查更新版本

### 1.2 功能截图

<table width="100%">
<tr>
<td width="33%">
    <img src="https://gitee.com/b0cloud/b0pass/raw/master/docs/images/s1.jpg" width="100%"/>
    <p>主页（文件共享页）</p>
</td>
<td width="33%">
    <img src="https://gitee.com/b0cloud/b0pass/raw/master/docs/images/s2.jpg" width="100%"/>
    <p>手机扫码，或获取链接地址</p>
</td>
<td width="33%">
    <img src="https://gitee.com/b0cloud/b0pass/raw/master/docs/images/s3.png" width="100%"/>
    <p>上传（上传页面）</p>
</td>
</tr>
<tr>
<td width="33%">
    <img src="https://gitee.com/b0cloud/b0pass/raw/master/docs/images/s5.jpg" width="100%"/>
    <p>上传（上传过程页面）</p>
</td>
<td width="33%">
    <img src="https://gitee.com/b0cloud/b0pass/raw/master/docs/images/s6.jpg" width="100%"/>
    <p>可点击在线浏览或下载</p>
</td>
<td width="33%">
    <img src="https://gitee.com/b0cloud/b0pass/raw/master/docs/images/s15.jpg" width="100%"/>
    <p>主页（管理文件）可点击删除</p>
</td>
</tr>
</table>

<img src="https://gitee.com/b0cloud/b0pass/raw/master/docs/images/s4.jpg" width=100%/>
    <p>上传超大文件</p>
    
![linux/amd64版本发行版](https://images.gitee.com/uploads/images/2019/1029/165512_c182287e_77462.png "b0pass_linux_cli.png")

linux/amd64版本发行版


## 2. 发行版下载使用

### 下载前请阅读
-  已编译好***Mac OS、Windows、Linux amd64***等平台的可执行文件（exe）
-  只需下载到电脑，双击开启即可使用
-  为了流畅使用UI界面，最好先安装了***谷歌浏览器***
-  如果要自定义端口，可以在命令行附加“ ****-p=1234**** ”自定义服务端口为1234

### 最新版下载地址
- https://github.com/bitepeng/b0pass/releases （直接下载，无需注册）
- https://gitee.com/b0cloud/b0pass/releases    <-国内推荐（虽需注册，但下载速度快）

## 3. 代码仓库
- https://github.com/bitepeng/b0pass   GitHub（主库）   欢迎star支持
- https://gitee.com/b0cloud/b0pass     GitEE（国内同步） 欢迎star支持


## 4. 使用场景
- ***手机电脑共享文件***

    电脑上双击执行 --> 手机扫码 --> 手机上的大文件传到电脑、或者电脑传文件到手机。

- ***电脑之间共享文件***

    电脑A上双击执行 --> 电脑B上浏览器输入A的地址 --> 电脑A上的大文件传到电脑B、或者电脑B传文件到电脑A。

- ***虚拟机和电脑之间共享文件***

    电脑上双击执行 --> 虚拟机上浏览器输入电脑的地址 --> 虚拟机上的大文件传到电脑、或者电脑传文件到虚拟机。

- ***更多使用场景***

    也可以用作“家庭影音中心”、“办公室文件共享”、“产品原型服务器”等。总之走局域网的HTTP协议，和是不是iPhone、iOS、安卓、虚拟机等都没有关系，跨平台共享文件。

## 5. 源码编译
```
# 下载代码，推荐使用go mod模式管理依赖
git clone https://github.com/bitepeng/b0pass.git

# 配置Goland支持go mod，更新依赖
cd docs/script && chomd +x ./do-vendor && ./do-vendor

# 编译运行开发版本
cd docs/script && chomd +x build-develop.sh && build-develop.sh
```
