# 百灵快传（Win-Cli版）
电脑手机大文件上传工具

## 开始使用
1. 双击 b0pass.exe 执行
2. 启动后 主电脑网址 会自动打开（类似 http://192.168.1.x:8888）
3. 控制台（黑色框）在使用期间不要关闭

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
