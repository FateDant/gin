package main

import (
	"gin/gorm"
	"gin/middle"
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
	router.Use(middle.MiddleWare, middle.MiddleWare2())
	//router.GET("/", func(context *gin.Context) {
	//	context.String(200, "hello gin", 123)
	//})

	router.GET("/user", Hello)

	//router.LoadHTMLGlob("views/**/*")
	router.LoadHTMLGlob("views/*")
	router.GET("/param/:id", Param) //必须传id
	//router.GET("/param/*id", Param) //非必传

	query := router.Group("/query")
	{
		query.GET("/query", study.GetQueryData)
		query.GET("/query_arr", study.GetQueryArrData)
		query.GET("/query_map", study.GetQueryMapData)
	}

	user := router.Group("/user")
	{
		user.POST("/user_add", study.UserAdd)
		user.POST("/user_add_bind", study.UserAddBind)
		user.POST("/file_upload", study.FileUpload)
		user.POST("/file_upload_batch", study.FileUploadBatch)
	}

	redirect := router.Group("/redirect")
	{
		redirect.GET("/redirect_a", study.RedirectA)
		redirect.GET("/redirect_b", study.RedirectB)
	}

	router.POST("/bind_form", study.BindForm)
	router.GET("/valid", study.Validate)

	//局部中间件 BasicAuth
	router.GET("/auth", gin.BasicAuth(gin.Accounts{
		"zs": "123456",
	}), middle.AuthTest)

	db := router.Group("/db")
	{
		db.GET("/connect", gorm.Connect)
		db.GET("/migrate", gorm.Migrate)
		db.GET("/create", gorm.CreateData)
		db.GET("/query", gorm.QueryData)
		db.GET("/update", gorm.UpdateData)
	}

	//router.Run(":9000")
	s := &http.Server{
		Addr:         ":9000",
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	s.ListenAndServe()
}
