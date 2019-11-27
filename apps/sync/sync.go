package sync

import (
	"github.com/gogf/gf/container/gmap"
	"github.com/gogf/gf/container/gset"
	"github.com/gogf/gf/frame/gmvc"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/glog"
	"log"
)


type Controller struct {
	gmvc.Controller
	ws *ghttp.WebSocket
}

var (
	users = gmap.New()
	names = gset.NewStrSet()
)

// Index 触发页面
// /sync/
func (c *Controller) Index() {
	if !c.Session.Contains("clientId") {
		_ = c.Session.Set("clientId", c.Session.Id())
	}
	_ = c.View.Display("sync.html")
}

// WebSocket 接口
// /sync/web-socket
func (c *Controller) WebSocket() {

	// 初始化WebSocket请求
	if ws, err := c.Request.WebSocket(); err == nil {
		c.ws = ws
	} else {
		glog.Error(err)
		return
	}

	// 初始化时设置用户信息
	clientId := c.Session.GetString("clientId")
	if clientId=="" {
		_ = c.Session.Set("clientId", c.Session.Id())
	}
	users.Set(c.ws, clientId)
	names.Add(clientId)

	for {
		// 阻塞读取WS数据
		msgType, msg, err  := c.ws.ReadMessage()
		if err != nil {
			users.Remove(c.ws)
			names.Remove(clientId)
			break
		}

		// 群发同步所有端
		glog.Cat("sync").Println("[sync] ",clientId,msg)
		_ = c.writeUsers()
		if msg != nil {
			msgs:= "{" +
					"\"clientId\":\""+clientId+"\"," +
					"\"msg\":\""+string(msg)+"\""+
				"}"
			_ = c.writeGroup(msgType,msgs)
		}
	}
}

// 群发消息
func (c *Controller) writeGroup(msgType int,msg string) error {
	msgs:= []byte(msg)
	users.RLockFunc(func(m map[interface{}]interface{}) {
		for user := range m {
			_ = user.(*ghttp.WebSocket).WriteMessage(msgType, []byte(msgs))
		}
	})

	return nil
}

// 向客户端返回用户列表
func (c *Controller) writeUsers() error {
	nameStr := ""
	names.Iterator(func(v string) bool {
		if nameStr =="" {
			nameStr+=v
		}else{
			nameStr+=","+v
		}
		return true
	})
	log.Println(nameStr)

	msgs:= "{" +
		"\"clientId\":\"0\"," +
		"\"msg\":\""+nameStr+"\"" +
		"}"
	if err := c.writeGroup(1,msgs); err != nil {
		return err
	}
	return nil
}