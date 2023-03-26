package gorm

import (
	"fmt"
	"gin/data_source"
	"gin/gorm/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func Raw(ctx *gin.Context) {
	//db, err := gorm.Open("mysql", "root:123456@tcp(localhost:3306)/gin?charset=utf8&parseTime=True&loc=Local")
	//if err != nil {
	//	panic(err)
	//}
	//defer db.Close()

	var users []models.User
	//db.Raw("select * from rename_user where name = 'JJJ'").Find(&users)
	data_source.Db.Raw("select * from rename_user where name = ?", "JJJ").Find(&users)
	fmt.Println(users)
}

func Exec(ctx *gin.Context) {
	db, err := gorm.Open("mysql", "root:123456@tcp(localhost:3306)/gin?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	//db.Exec("insert into rename_user(name,age,pic,phone) values (?,?,?,?)", "mingzi", "16", "xxxx", "1341234")

	db.Exec("update rename_user set name = ? where id = ?", "niuniuniu", "1")
}
