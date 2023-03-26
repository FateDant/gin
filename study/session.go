package study

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func SessionTest(ctx *gin.Context) {
	session := sessions.Default(ctx)
	session.Set("name", "adsfas")
	name := session.Get("name")
	fmt.Println(name)

	//session.Delete("name")
	//name2 := session.Get("name")
	//fmt.Println(name2)

	//删除所有
	//session.Clear()
}
