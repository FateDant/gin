package study

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

func GetQueryData(ctx *gin.Context) {
	id := ctx.Query("id")
	name := ctx.DefaultQuery("name", "ming")
	ctx.String(http.StatusOK, "word,%s,%s", id, name)
}

// GetQueryArrData
// http://127.0.0.1:9000/query_arr?id=12,322,43
func GetQueryArrData(ctx *gin.Context) {
	ids := ctx.QueryArray("id")
	ctx.String(http.StatusOK, "word,%s", ids)
}

// GetQueryMapData
// http://127.0.0.1:9000/query_map?user[name]=1&user[age]=2
func GetQueryMapData(ctx *gin.Context) {
	maps := ctx.QueryMap("user")
	ctx.String(http.StatusOK, "word,%s", maps)
}

// UserAdd
//
//	curl -X POST \
//	 http://127.0.0.1:9000/user_add \
//	 -H 'Postman-Token: 863e0a7b-442b-480c-8e61-71559c1fb101' \
//	 -H 'cache-control: no-cache' \
//	 -H 'content-type: multipart/form-data; boundary=----WebKitFormBoundary7MA4YWxkTrZu0gW' \
//	 -F name=1 \
//	 -F pwd=2 \
//	 -F love=22 \
//	 -F love=33 \
//	 -F 'user[addr]=aaaaa' \
//	 -F 'user[phone]=99999'
func UserAdd(ctx *gin.Context) {
	name := ctx.PostForm("name")
	pwd := ctx.PostForm("pwd")
	age := ctx.DefaultPostForm("age", "18")
	love := ctx.PostFormArray("love")
	user := ctx.PostFormMap("user")
	ctx.String(http.StatusOK, "word,%s,%s,%s,%s,%s", name, pwd, age, love, user)
	ctx.JSON(http.StatusOK, gin.H{"code": 200, "msg": "success"})
}

type UserInfo struct {
	Id   int    `form:"id"`
	Name string `form:"name"`
	Age  int    `form:"age"`
	Addr string `form:"addr"`
}

func UserAddBind(ctx *gin.Context) {
	var user_info UserInfo
	err := ctx.ShouldBind(&user_info)
	fmt.Println(err)
	ctx.JSON(http.StatusOK, gin.H{"code": 200, "msg": "success", "err": err, "user_info": user_info})
}

func FileUpload(ctx *gin.Context) {
	file, err := ctx.FormFile("file")
	timeUnixInt := time.Now().Unix()
	timeUnixStr := strconv.FormatInt(timeUnixInt, 10) //需要将int转成string类型 不能通过string(int) 转换
	dst := "upload/" + timeUnixStr + file.Filename
	ctx.SaveUploadedFile(file, dst)
	//fmt.Println(file.Filename)
	ctx.JSON(http.StatusOK, gin.H{"code": 200, "msg": "success", "err": err, "file": file.Filename})
}

func FileUploadBatch(ctx *gin.Context) {
	form, err := ctx.MultipartForm()
	files := form.File["file"]

	for _, file := range files {
		timeUnixInt := time.Now().Unix()
		timeUnixStr := strconv.FormatInt(timeUnixInt, 10) //需要将int转成string类型 不能通过string(int) 转换
		dst := "upload/" + timeUnixStr + file.Filename
		ctx.SaveUploadedFile(file, dst)
	}
	ctx.JSON(http.StatusOK, gin.H{"code": 200, "msg": "success", "err": err, "file": files})
}

func RedirectA(ctx *gin.Context) {
	fmt.Println("这是A")
	ctx.Redirect(http.StatusFound, "/redirect_b")
}

func RedirectB(ctx *gin.Context) {
	fmt.Println("这是B")
	ctx.String(http.StatusOK, "这是B")
}
