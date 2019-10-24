package index

import "github.com/gogf/gf/frame/gmvc"

type Controller struct {
	gmvc.Controller
}

func (c *Controller) Index() {
	//c.Response.Write("Index")
	_ = c.View.Display("index.html")
}
