package main

import (
	"fmt"
	"github.com/kataras/iris"
)

func main() {
	app := iris.Default()
	app.Get("/hello", func(ctx iris.Context) {
		ctx.WriteString("hello,world!")
	})
	err := app.Run(iris.Addr(":8082"))
	fmt.Println(err)
}
