package study

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// User
// form标签支持 get 和 post
type User struct {
	Name string `form:"name" json:"name" uri:"name"`
	Age  int    `form:"age" json:"age" uri:"age"`
	Addr string `form:"addr" json:"addr" uri:"addr"`
}

func BindForm(ctx *gin.Context) {
	var user User
	err := ctx.ShouldBind(&user)
	ctx.JSON(http.StatusOK, gin.H{"code": 200, "msg": "success", "err": err, "user": user})

	if err != nil {
		ctx.String(http.StatusOK, "success")
	} else {
		ctx.String(http.StatusInternalServerError, "error")
	}
}
