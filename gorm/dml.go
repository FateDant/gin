package gorm

import (
	"gin/gorm/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"net/http"
)

func CreateData(ctx *gin.Context) {
	db, err := gorm.Open("mysql", "root:123456@tcp(localhost:3306)/gin?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close() //关闭空闲连接

	db.Create(&models.User{Name: "张三", Age: "18", Pic: "照片", Phone: "1234567890"})
}

func QueryData(ctx *gin.Context) {
	db, err := gorm.Open("mysql", "root:123456@tcp(localhost:3306)/gin?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close() //关闭空闲连接

	var user models.User
	//db.First(&user, 1)              //默认id
	db.First(&user, "name=?", "张三") //指定字段
	ctx.JSON(http.StatusOK, gin.H{"user": user})
}

func UpdateData(ctx *gin.Context) {
	db, err := gorm.Open("mysql", "root:123456@tcp(localhost:3306)/gin?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close() //关闭空闲连接

	var user models.User
	db.First(&user, 1)
	db.Model(&user).Update("age", "25")
	db.Model(&user).Update("pic", "zs-xxxx")
}

func DeleteData(*gin.Context) {
	db, err := gorm.Open("mysql", "root:123456@tcp(localhost:3306)/gin?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close() //关闭空闲连接

	var user models.User
	db.First(&user, 2)
	db.Delete(&user)
}
