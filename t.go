package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	p := sync.Pool{
		New: func() interface{} {
			return "xiaoliang"
		},
	}

	runtime.GOMAXPROCS(4)

	a := p.Get().(string)
	fmt.Println(a)
	p.Put("5")

	b :=p.Get().(string)
	//runtime.GC()

	fmt.Printf("%+v", b)
}
