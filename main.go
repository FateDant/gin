package main

import (
	"gin/study"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Hello(ctx *gin.Context) {
	//ctx.String(200, "hello gin", 456)
	name := "名称"
	ctx.HTML(200, "index.html", name)
}

func Param(ctx *gin.Context) {
	id := ctx.Param("id")
	ctx.String(http.StatusOK, "hello,%s", id)
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
	router.GET("/param/:id", Param) //必须传id
	//router.GET("/param/*id", Param) //非必传

	router.GET("/query", study.GetQueryData)
	router.GET("/query_arr", study.GetQueryArrData)
	router.GET("/query_map", study.GetQueryMapData)
	router.POST("/user_add", study.UserAdd)
	router.POST("/user_add_bind", study.UserAddBind)

	router.Run(":9000")
}
