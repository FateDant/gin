package middle

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func MiddleWare(ctx *gin.Context) {
	fmt.Println("自定义中间件")
}

func MiddleWare2() gin.HandlerFunc {
	return func(context *gin.Context) {
		fmt.Println("自定义中间件2")
	}
}

var map_data = map[string]interface{}{
	"zs": gin.H{"age": 18, "addr": "aaa"},
	"ls": "uuuu",
}

func AuthTest(ctx *gin.Context) {
	userName := ctx.Query("user_name")
	i, ok := map_data[userName]
	if ok {
		ctx.JSON(http.StatusOK, gin.H{"user": userName, "data": i})
	} else {
		ctx.JSON(http.StatusNotFound, gin.H{"user": userName, "data": nil})
	}
}
