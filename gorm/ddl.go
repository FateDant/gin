package gorm

import (
	"gin/gorm/models"
	"gin/gorm/models/relate"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"net/http"
)

func Connect(ctx *gin.Context) {
	//用户名：密码@tcp(ip:port)/数据库？charset=utf8&parseTime=True&loc=Local
	db, err := gorm.Open("mysql", "root:123456@tcp(localhost:3306)/gin?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close() //关闭空闲连接

	//建表
	db.CreateTable(&models.User{}) //users
	//指定表名
	db.Table("user").CreateTable(&models.User{})

	//删除表
	//db.DropTable("users")
	db.DropTable(&models.User{})

	//表是否存在
	b := db.HasTable("user")
	ctx.JSON(http.StatusOK, gin.H{"b": b})

	b = db.HasTable(&models.User{})
	ctx.JSON(http.StatusOK, gin.H{"b": b})
}

func Migrate(ctx *gin.Context) {
	db, err := gorm.Open("mysql", "root:123456@tcp(localhost:3306)/gin?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	//统一加前缀或后缀
	//gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
	//	return "pre_" + defaultTableName + "_suf"
	//}

	//自动迁移，仅仅会创建表，添加列和索引，不会改变现有列的结构或删除列
	//db.AutoMigrate(&models.User{}, &models.GormModel{}, &models.TagModel{}, &relate.UserMaster{}, &relate.UserSlave{})
	db.AutoMigrate(&relate.UserMaster{}, &relate.UserSlave{}, &relate.One{}, &relate.More{}, &relate.ManyOne{}, &relate.ManyTwo{})
}
