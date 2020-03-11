package koala

import (
	"sync"
)

// 注入依赖
type DIer interface {
	// 注入  interface
	Inject(name string, in interface{})
	// 调用 interface
	Invoke(name string) interface{}

	////////////////////目测下面这个是鸡肋，等用到时候再开放吧
	// 注入 func
	InjectFun(name string, in DIHandler)
	// 调用 func
	InvokeFun(name string) DIHandler
}

// DI数据结构
type DI struct {
	store map[string]interface{}
	mutex sync.RWMutex
}

// di 函数类型
type DIHandler func() interface{}

func NewDi() DIer {
	di := new(DI)
	di.store = make(map[string]interface{})
	return di
}

// 注入
// @param in interface{} 只能特定注入
func (d *DI) Inject (name string, in interface{})  {
	d.mutex.Lock()
	d.store[name] = in
	d.mutex.Unlock()
}

// 调用
// 使用时需要进行断言才可使用
func (d *DI) Invoke (name string) interface{} {
	d.mutex.RLock()
	value := d.store[name]
	d.mutex.RUnlock()

	return value
}

// 注入 func
func (d *DI) InjectFun (name string, in DIHandler)  {
	d.mutex.Lock()
	d.store[name] = in
	d.mutex.Unlock()
}

// 调用 func
func (d *DI) InvokeFun (name string) DIHandler{
	d.mutex.RLock()
	value := d.store[name]
	d.mutex.RUnlock()

	return value.(DIHandler)
}

