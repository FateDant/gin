package study

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
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
