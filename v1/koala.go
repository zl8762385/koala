package koala

import (
	"fmt"
	"koala/v1/utils"
	"net/http"
	"sync"
)

// 基础信息
const (
	Version       = "0.0.1"
	FrameworkName = "koala"
	Anthor        = "xiaoliang"
)

// 接口类型
type Middleware interface {}

// 上下文处理函数
type HandlerFunc func(*Context)

// 定义 koala引擎结构
type Koala struct {
	// 调试模式
	debug bool

	// 版本号
	version string

	// 节点树
	trees map[string]*node

	// 临时对象池
	pool sync.Pool

	// 单例注入
	di DIer

	// 中间件
	middleware []HandlerFunc

}

// 实例化Koala
func New() *Koala {

	engine := &Koala{
		version: Version,
	}

	// 初始化临时对象池
	engine.pool = sync.Pool{
		New: func() interface{} {
			return engine.httpContext()
		},
	}

	// 中间件初始化
	engine.middleware = make([]HandlerFunc, 0)

	// 依赖注入
	engine.InjectDIer(NewDi())
	//// 路由
	engine.Inject("router", NewRouter(engine))

	return engine
}
//////////////////////////////////////////////////////////////// 内部

// init初始化
func init() {
	utils.DebugPrintf("%s","启动中")
}

// http请求上下文
func (k *Koala) httpContext() *Context {
	return &Context{koala:k}
}

func (k *Koala) router() IRouter {
	router := k.di.Invoke("router").(IRouter)
	return router
}

// run http server
func (k *Koala) run (address string) {
	utils.DebugPrintf("监听端口%s", address)
	utils.DebugPrintf("%s", "服务启动成功")
	http.ListenAndServe(address, k)
}

//////////////////////////////////////////////////////////////// 开放
// 注册中间件
func(k *Koala) Use(m ...Middleware) {
	for i := range m {
		if m[i] != nil {
			k.middleware = append(k.middleware, warpMiddleware(m[i]))
		}
	}
	//fmt.Printf("%+V", k.middleware)
}

// 处理
func warpMiddleware(m Middleware) HandlerFunc {
	switch m:= m.(type) {
	case HandlerFunc:
		return m
	case func (*Context):
		return m
	default:
		fmt.Printf("%+V", m)
		panic("没找到相关中间件")
	}
}


// 依赖注入开放接口
func (k *Koala) InjectDIer(di DIer) {
	k.di = di
}

// 注入
func (k *Koala) Inject(name string, in interface{}) {
	switch name {
	case "router":
		if _, ok := in.(IRouter); !ok {
			panic("Di router必须实现 interface koala.IRouter")
		}
	}
	k.di.Inject(name ,in)
}

// 调用
func (k *Koala) Invoke(name string) interface{} {
	return k.di.Invoke(name)
}

// 添加路由
func (k *Koala) Add(httpMethod, path string, router HandlerFunc) {
	k.router().Add(httpMethod,path,router)
}

// 静态文件 服务
func (k *Koala) ServeFiles() {

}

// 实现net/http 需要的servehttp服务
func (k *Koala) ServeHTTP(rw http.ResponseWriter, req *http.Request) {

	// 取一条对象 默认ctx.rw是空，这里在最开始New时已经初始化完成
	ctx := k.pool.Get().(*Context)
	ctx.Reset(rw,req,k)

	// 过滤相关文件
	if _, ok := k.router().Filter(rw, req); !ok {
		return
	}

	// 执行相关操作
	ctx.Next()

	k.pool.Put(ctx)
}

// 运行器
func (k *Koala) Run(address string) {
	k.run(address)
}
