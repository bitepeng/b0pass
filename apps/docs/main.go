package app

import (
	"b0go/core/engine"
	"embed"
	"fmt"
	"html/template"
	"io/fs"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/logrusorgru/aurora"
)

type AppConfig struct {
	Live bool
}

var (
	app    *engine.AppConfig
	appId  = "docs"
	config = new(AppConfig)

	//go:embed ui/dist
	uiFS embed.FS

	//go:embed template
	tplFS embed.FS
)

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
	engine.Print(aurora.Green("App docs loaded"), aurora.BrightCyan(config))
	//template files
	if config.Live {
		engine.Gin.LoadHTMLGlob(app.UIDir + "/../../template/*")
	} else {
		t, _ := template.ParseFS(tplFS, "template/*")
		engine.Gin.SetHTMLTemplate(t)
	}
	// /docs/index.html 		=> 		/md/index.md
	// /docs/editor-index.html 	=> 		/md/editor/index.md
	// /?ext=doc				=>		显示侧边栏
	engine.GET(appId, "/:name", "{}", "md", md)
	engine.Gin.Static("/docs_root", "docs/")
}

// md 解析md文件为html
func md(c *gin.Context) {
	fmt.Println("=====md解析md文件为html=====")
	name := c.Param("name")
	ext := c.DefaultQuery("ext", "index")
	engine.Printf("docs/:name=%s", name)
	//file path
	name = strings.ReplaceAll(name, "-", "/")
	mdpath := strings.Replace(name, ".html", ".md", -1)
	engine.Printf("name =%s; ext=%s; mdfile =%s", name, ext, "md/"+mdpath)
	data, err := os.ReadFile("docs/" + mdpath)
	if err != nil {
		engine.Printf("%s", err)
	}
	//data str
	datastr := string(data)
	//title
	datas := strings.Split(strings.Trim(datastr, ""), "\n")
	title := strings.Trim(strings.ReplaceAll(datas[0], "#", ""), "")
	if title == "" {
		title = mdpath
	}
	c.HTML(http.StatusOK, "show.html", gin.H{
		"title": title,
		"ext":   ext,
		"data":  datastr,
	})
}
