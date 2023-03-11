package main

import (
	"gin/study"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
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
	router.POST("/file_upload", study.FileUpload)
	router.POST("/file_upload_batch", study.FileUploadBatch)
	router.GET("/redirect_a", study.RedirectA)
	router.GET("/redirect_b", study.RedirectB)
	router.POST("/bind_form", study.BindForm)
	router.GET("/valid", study.Validate)

	//router.Run(":9000")
	s := &http.Server{
		Addr:         ":9000",
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	s.ListenAndServe()
}
