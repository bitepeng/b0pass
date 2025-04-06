# B0Boot-GO
> b0boot-go -- 基于Go语言的极简Web开发脚手架


## 技术特性
- 极致的微内核插件化应用架构
- 模块应用的前后端代码物理隔离
- 减少非必要代码包装和依赖
- 提供应用引擎和模块注册网关
- 统一配置、统一日志、中间件支持


## 快速开始

开发应用模块，如 app1 模块，一般步骤如下：
- 模块目录：强烈建议使用约定的目录结构开发，事半功倍。
- 入口文件：按框架结构开发本模块的main.go
- 其他文件：也可以按golang规则开发其他go文件，按需开发ui界面
- 注册应用：按“应用模块注册”将本模块注册到主入口
- 启动应用：进入main目录，启动main.go

### 1. 基本规则
#### 1.1 建议目录结构
```
- apps              # 应用模块根目录
    - app1          # 应用1根目录
        - ui        # 应用1UI资源目录
        - main.go   # 应用1主程序
    - docs          # 应用2根目录
        - ui        # 应用2UI资源目录
        - main.go   # 应用2主程序
    ...             # 其他应用
- core              # 内核目录
    - engine        # 引擎目录
    - gateway       # 网关目录
    - tools         # 自定义工具库
- main              # 主程序目录
    - config.ini    # 配置文件
    - main.go       # 入口主程序
...
```
#### 1.2 路由规则
<img src="/app/docs/img/devdoc.png" width="500">
<p>内置DevDoc：http://127.0.0.1:8899/dev/doc</p>

```
一、网页内容/由docs应用提供/Markdown自动解析
- /							=> 首页 跳转到 /docs/index.html
- /docs/%mdfile.html%	  	=> markdown文档 /docs/%mdfile.md%
- /docs_root/%static%	  	=> 静态文件 /docs_root/%static%

二、网关GateWay相关/开发文档
- /dev/doc				 	=> api doc 文档页面
- /dev/api				 	=> api json 后端接口

三、具体应用相关路由
- /app/%appid%			 	=> 具体应用首页
- /app/%appid%/%act%	   	=> 具体应用api json
```
路由规则定义在 "core/gateway/main.go" 中定义


### 2. 具体模块开发


#### apps/app1/main.go
```go
package app

import (
	"b0go/core/engine"
	"embed"
	"io/fs"

	"github.com/gin-gonic/gin"
)

// 定义该应用配置项
type AppConfig struct {
	Name string
}

// 定义应用全局变量
var (
	app    *engine.AppConfig
	appId  = "app1"
	config = new(AppConfig)

	//go:embed ui/dist
	uiFS embed.FS
)

// init()在main.go中被调用
func init() {
	uiDist, _ := fs.Sub(uiFS, "ui/dist")
	app = &engine.AppConfig{
		Name:   appId,
		Type:   engine.APP_APP,
		Config: config,
		UIFS:   uiDist,
		Run:    run,
	}
	engine.AppInstall(app)
}

func run() {
	engine.GET(appId, "/ping", "{}", "ping", ping)
	engine.POST(appId, "/ping", "{}", "ping", ping)

}

func ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": appId + " pong " + config.Name,
	})
}

```

### 3. 应用模块注册
#### main/config.ini
```ini
[gateway]
ListenAddr = ":8899"
Live = true

[app1]
Name = "b0go"

[docs]
Live = true
```
注意：
- 此配置文件至少要配置[gateway]节点，其他为注册应用的配置
- 每个注册应用（main.go的import定义）都必须在此文件有一个节点，即使没有配置项

#### main/main.go
```go
package main

import (
	_ "b0go/apps/app1"      //注册app1
	_ "b0go/apps/docs"      //注册docs

	"b0go/core/engine"
	_ "b0go/core/gateway"   //默认必须注册
)

func main() {
	engine.Run("config.ini")
	select {}
}
```

### 4. 启动或打包
#### 启动项目
```shell
go run main/main.go
```
访问 http://127.0.0.1:8081
#### 打包项目
```shell
go build main/main.go
```
部署到服务器上

## 技术选型
- golang
- gin
- gorm