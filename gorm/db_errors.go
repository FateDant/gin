package gorm

import (
	"fmt"
	"gin/gorm/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func Errors(ctx *gin.Context) {
	db, err := gorm.Open("mysql", "root:123456@tcp(localhost:3306)/gin?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		panic(err)
	}
	defer db.Close()

	var users []models.User

	res := db.Where("name = ?", "JJJJ").Find(&users)
	fmt.Println(users)
	fmt.Println(res.RowsAffected) //获取查询结果条数

	if res.Error != nil {
		//发生错误
	}

	errs := res.GetErrors()

	for _, err := range errs {
		fmt.Println(err)
	}
}

func Transaction(ctx *gin.Context) {
	db, err := gorm.Open("mysql", "root:123456@tcp(localhost:3306)/gin?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	ct := db.Begin() // 开启事务
	res := ct.Commit()
	if res.Error != nil {
		ct.Rollback() //回滚
	}
}
