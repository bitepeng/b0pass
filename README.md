# 百灵快传（B0Pass）

基于Go的高性能“手机电脑超大文件传输神器”、“局域网共享文件服务器”，只需一个文件双击开启。

## 1. 主要功能
### 1.1 功能截图


<img src="https://raw.githubusercontent.com/bitepeng/b0pass/master/docs/images/s4.jpg" height=400/>
    <p>上传（可以选超大文件）</p>

<table width="100%">
<tr>
<td>
    <img src="https://raw.githubusercontent.com/bitepeng/b0pass/master/docs/images/s1.jpg" height=400/>
    <p>主页（文件共享页）</p>
</td>
<td>
    <img src="https://raw.githubusercontent.com/bitepeng/b0pass/master/docs/images/s2.jpg" height=400/>
    <p>手机扫码，或获取链接地址</p>
</td>
<td>
    <img src="https://raw.githubusercontent.com/bitepeng/b0pass/master/docs/images/s3.png" height=400/>
    <p>上传（上传页面）</p>
</td>
</tr>
<tr>
<td>
    <img src="https://raw.githubusercontent.com/bitepeng/b0pass/master/docs/images/s5.jpg" height=400/>
    <p>上传（上传过程页面）</p>
</td>
<td>
    <img src="https://raw.githubusercontent.com/bitepeng/b0pass/master/docs/images/s6.jpg" height=400/>
    <p>可点击在线浏览或下载</p>
</td>
<td>
    <img src="https://raw.githubusercontent.com/bitepeng/b0pass/master/docs/images/s15.jpg" height=400/>
    <p>主页（管理文件）可点击删除</p>
</td>
</tr>
</table>


## 2. 下载使用
### 2.1 Mac OS
- [b0pass_mac.dmg](https://github.com/bitepeng/b0pass/blob/master/docs/release/v0.1/b0pass_mac.dmg)

### 2.2 Windows
- [b0pass_win.exe](https://github.com/bitepeng/b0pass/blob/master/docs/release/v0.1/b0pass_wn32.exe)

## 3. 源码编译
```
# 下载代码，推荐使用go mod模式管理依赖
git clone https://github.com/bitepeng/b0pass.git

# 配置Goland支持go mod，更新依赖
cd docs/script && chomd +x ./do-vendor && ./do-vendor

# 编译运行开发版本
cd docs/script && chomd +x build-develop.sh && build-develop.sh
```
