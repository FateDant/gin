package main

import (
	_ "gin/data_source" //数据库连接
	"gin/gorm"
	_ "gin/logs_source" //日志配置
	"gin/middle"
	"gin/study"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
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
	//config目录  json格式的配置文件 不能写注释！

	router := gin.Default()
	//创建日志文件
	file, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(file, os.Stdout) //写入文件和输出到控制台

	//加密的盐 基于cookie
	//store := cookie.NewStore([]byte("hallen"))

	//size最大连接数 基于redis
	store, err := redis.NewStore(10, "tcp", "127.0.0.1:6379", "", []byte("hallen"))
	if err != nil {
		panic(err)
	}
	//使用session中间件
	router.Use(sessions.Sessions("gin_session", store))

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
		db.GET("/queryOne", gorm.QueryOneByOne)
		db.GET("/updateOne", gorm.UpdateOneByOne)
		db.GET("/deleteOne", gorm.DeleteOneByOne)
		db.GET("/createOneMore", gorm.CreateOneByMore)
		db.GET("/queryOneMore", gorm.QueryOneByMore)
		db.GET("/updateOneByMore", gorm.UpdateOneByMore)
		db.GET("/createManyToMany", gorm.CreateManyToMany)
		db.GET("/queryManyToMany", gorm.QueryManyToMany)
	}

	db_interface := router.Group("/db_interface")
	{
		db_interface.GET("/first", gorm.First)
		db_interface.GET("/where", gorm.Where)
		db_interface.GET("/create", gorm.Create)
		db_interface.GET("/update", gorm.Update)
		db_interface.GET("/not", gorm.Not)
		db_interface.GET("/group", gorm.Group)
		db_interface.GET("/join", gorm.Join)
		db_interface.GET("/firstInit", gorm.FirstInit)
	}

	db_errors := router.Group("/db_errors")
	{
		db_errors.GET("/errors", gorm.Errors)
	}

	db_sql := router.Group("/db_sql")
	{
		db_sql.GET("/raw", gorm.Raw)
		db_sql.GET("/exec", gorm.Exec)
	}

	log := router.Group("/log")
	{
		log.GET("test", study.LogTest)
	}

	session := router.Group("/session")
	{
		session.GET("test", study.SessionTest)
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
