package main

import (
	"fmt"
	"koala/v1"
)

// 上下文中间件测试
func middloger() koala.HandlerFunc {
	return func(ctx *koala.Context) {

		fmt.Printf("%+v", ctx.Req)
		fmt.Println("middloger 中间件")
		// fmt.Printf("%+V 进入到中间件了\n", ctx)
	}
}


func main() {
	app := koala.New()

	/*
	app.Use(middloger())

	app.Use(func(ctx *koala.Context) {
		fmt.Println("自定义中间件")
	})

	app.Use(func(ctx *koala.Context) {
		fmt.Println("自定义中间件111")
	})
	*/




	app.Add("GET", "/member/:id", func(ctx *koala.Context) {

		type ss struct {
			Name string `json:"name"`
			Age int `json:"age"`
		}

		test := ss{"xiaoliang", 32}
		ctx.Json(test)
	})

	app.Run(":8080")
}