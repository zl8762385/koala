package koala

import (
	"fmt"
	"net/http"
)

// 定义koala 路由接口
type IRouter interface {
	Add(httpMethod, path string, router HandlerFunc)
	// Match(method , path string, c *Context) (string, string)
	Filter(rw http.ResponseWriter, req *http.Request) (error, bool)
	// 处理路由
	HandlerRouter(ctx *Context)
	// ByName(name string) string
	Test() string
}


type Router struct {
	koala *Koala
	routerName string
	method string
}

// 实例化路由
func NewRouter(k *Koala) IRouter {
	Ir := &Router{
		koala:k,
	}

	return Ir
}

// 添加路由节点
func (r *Router) Add(httpMethod, path string, router HandlerFunc) {

	fmt.Printf("Add %+v %+v \n", httpMethod, path)
	// 检查树是否为空
	if r.koala.trees == nil {
		r.koala.trees = make(map[string]*node)
	}

	// 找到大类GET POST PUT等
	root := r.koala.trees[httpMethod]
	if root == nil {
		// new 赋值 保存到节点上
		root = new(node)
		r.koala.trees[httpMethod] = root
	}

	root.addRoute(path, router)
}

/*
func (r *Router) Match(method, path string, c *Context) (string, string) {
	r.routerName = path
	r.method = method

	return method, path
}
*/

// 过滤请求，直接从http server层直接过滤
func(r *Router) Filter(rw http.ResponseWriter, req *http.Request) (error, bool) {
	if req.URL.RequestURI() == "/favicon.ico" {
		return nil, false
	}

	return nil, true
}

// 路由映射处理
func (r *Router) HandlerRouter(ctx *Context) {
	fmt.Println("路由映射处理 start")

	// 路由处理 执行对应函数
	if root := ctx.koala.trees[ctx.method]; root != nil {
		// ps /member/:name/:age
		if handlerFunc, ps, _ := root.getValue(ctx.routerName); handlerFunc != nil {

			ctx.Param = ps
			handlerFunc(ctx)

		} else if ctx.method != "CONNECT" {
			// 客户端没有使用connect隧道进行代理请求

			ctx.RW.Write([]byte("404"))
			// 永久重定向 使用GET方法请求
			fmt.Println("not connect", ctx.method)
		} else {
			http.NotFound(ctx.RW, ctx.Req)
		}
	}

	fmt.Println("路由映射处理 end")
}

func (r *Router) Test() string {
	return "xiaoliang11"
}
/*
func (r *Router) Router() *Router {

}
*/
