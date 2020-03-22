package main

import (
	"fmt"
	"koala/v1"
)

func main() {
	app := koala.New()
	app.Add("GET", "/profile/xiaoliang", func(ctx *koala.Context) {
		fmt.Println("xiaoliang")
		ctx.Text("profile.xiaoliang")
	})
	app.Run(":8080")
}