package main

import (
	"github.com/gin-gonic/gin"
)

func Hello(ctx *gin.Context) {
	//ctx.String(200, "hello gin", 456)
	name := "名称"
	ctx.HTML(200, "index.html", name)
}

func main() {
	router := gin.Default()
	//router := gin.New()
	//router.GET("/", func(context *gin.Context) {
	//	context.String(200, "hello gin", 123)
	//})

	router.GET("/user", Hello)

	//router.LoadHTMLGlob("views/**/*")
	router.LoadHTMLGlob("views/*")

	router.Run(":9000")
}
