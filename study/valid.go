package study

import (
	"fmt"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Article
// valid : beego验证器
type Article struct {
	Id      int    `form:"_"`
	Title   string `form:"title" valid:"Required"`
	Content string `form:"content" valid:"Min(5)"`
	Desc    string `form:"desc"`
}

func Validate(ctx *gin.Context) {
	var article Article
	err := ctx.ShouldBind(&article) //绑定参数
	fmt.Println(article)
	fmt.Println(err)

	// 自定义提示语
	var MessageTmp = map[string]string{
		"Required": "不能为空",
		"Min":      "最小值：%d",
	}
	validation.SetDefaultMessage(MessageTmp)

	vaild := validation.Validation{}
	bool, err := vaild.Valid(&article) //开始验证

	//设置中文key值
	keyMap := map[string]interface{}{
		"Title.Required.": "标题",
		"Content.Min.":    "内容",
	}

	if !bool {
		//循环错误信息
		for _, err2 := range vaild.Errors {
			fmt.Println(err2.Key)
			fmt.Println(err2.Message)
			key := keyMap[err2.Key]
			ctx.String(http.StatusOK, key.(string)+err2.Message)
			//ctx.String(http.StatusOK, err2.Key)
		}
	}
	//ctx.String(http.StatusOK, "success")
}
