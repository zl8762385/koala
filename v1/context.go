package koala

import (
	"net/http"
)

type IContext interface {
	Reset(rw http.ResponseWriter, req *http.Request,k *Koala)
}

// 上下文结构体
type Context struct {
	RW http.ResponseWriter
	Req *http.Request
	koala *Koala
	routerName string
	method string
	Param Params
}


//=========================================================================output
// 输出
func (c *Context) Text (value interface{}) {
	Text.Content(c.RW, value)
}

func (c *Context) Json (value interface{}) {
	Json.Content(c.RW, value)
}



// 重置上下文
func (c *Context) Reset(rw http.ResponseWriter, req *http.Request, k *Koala) {
	c.RW = rw
	c.Req = req
	c.koala = k
	c.routerName = req.URL.Path
	c.method = req.Method
}

// 执行句柄
// 首先处理中间件，然后处理路由句柄
func (c *Context) Next() {
	// http中间件处理
	c.middleware()

	// 路由映射处理
	c.koala.router().HandlerRouter(c)
}


// 执行HTTP中间件
func (c *Context) middleware() {
	for m := range c.koala.middleware{
		c.koala.middleware[m](c)
	}
}

