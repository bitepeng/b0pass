# 百灵快传（Linux版）
电脑手机大文件上传工具

## 开始使用
1. 给 start.sh 赋予可执行权限
2. ./start.sh 启动百灵快传，会在后台执行
3. 启动后 主电脑网址 可以在 nohup.out 中查看（类似 http://192.168.1.x:8888）

## 更多信息
官网 https://4bit.cn/p/b0pass

## 配置文件范例
```ini
[gateway]
ListenAddr = ":8888"
Domain=""

[pass]
Path = "files"
CodeReadonly = "123"    # 使用此code登录只有只读权限，留空则不启用
CodeReadwrite = "admin" # 使用此code登录拥有全部权限，留空则不启用
```