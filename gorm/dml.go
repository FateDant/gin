package gorm

import (
	"fmt"
	"gin/gorm/models"
	"gin/gorm/models/relate"
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

func QueryOneByOne(ctx *gin.Context) {
	db, err := gorm.Open("mysql", "root:123456@tcp(localhost:3306)/gin?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close() //关闭空闲连接

	//gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
	//	return "pre_" + defaultTableName + "_suf"
	//}
	//master := relate.UserMaster{
	//	Id:   2,
	//	Name: "iuuiu",
	//	Age:  45,
	//	Addr: "wre",
	//}

	//slave := relate.UserSlave{
	//	Id:    2,
	//	Pic:   "zxc",
	//	Phone: "12321312",
	//}

	//db.Model(&slave).Association("UserMaster").Append(&master)

	var oneByone relate.UserSlave

	//第一种
	db.Debug().First(&oneByone, 2) //debug 打印sql语句
	fmt.Println(oneByone)
	err = db.Model(&oneByone).Association("UserMaster").Error //UserMaster 这个是模型里面的字段名，但是字段名要和关联模型的名称保持一直
	db.Model(&oneByone).Association("UserMaster").Find(&oneByone.UserMaster)
	fmt.Println(err)
	fmt.Println(oneByone)

	//第二种
	db.Debug().Preload("UserMaster").First(&oneByone, 2) //这个更简单
	fmt.Println(oneByone)
}

func UpdateOneByOne(ctx *gin.Context) {
	db, err := gorm.Open("mysql", "root:123456@tcp(localhost:3306)/gin?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close() //关闭空闲连接

	var oneByOne relate.UserSlave
	db.Preload("UserMaster").First(&oneByOne, 2)
	db.Model(&oneByOne.UserMaster).Update(relate.UserMaster{Name: "uuuu", Age: 55, Addr: "iiii"})
	fmt.Println(oneByOne)
}

func DeleteOneByOne(ctx *gin.Context) {
	db, err := gorm.Open("mysql", "root:123456@tcp(localhost:3306)/gin?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close() //关闭空闲连接

	var oneByOne relate.UserSlave
	db.Preload("UserMaster").First(&oneByOne, 2)
	db.Delete(&oneByOne.UserMaster)
}

func CreateOneByMore(ctx *gin.Context) {
	db, err := gorm.Open("mysql", "root:123456@tcp(localhost:3306)/gin?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close() //关闭空闲连接

	one := relate.One{
		Name: "ttt",
		Model: gorm.Model{
			ID: 1,
		},
		Mores: []relate.More{
			{
				Model: gorm.Model{
					ID: 1,
				},
				Hobbies: "qq",
			},
			{
				Model: gorm.Model{
					ID: 2,
				},
				Hobbies: "qq",
			},
			{Hobbies: "EEE"},
		},
	}
	db.Create(&one)
}

func QueryOneByMore(ctx *gin.Context) {
	db, err := gorm.Open("mysql", "root:123456@tcp(localhost:3306)/gin?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close() //关闭空闲连接

	var oneByMore relate.One
	db.Preload("Mores").Find(&oneByMore, 1)
	fmt.Println(&oneByMore)
}

func UpdateOneByMore(ctx *gin.Context) {
	db, err := gorm.Open("mysql", "root:123456@tcp(localhost:3306)/gin?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close() //关闭空闲连接

	var oneByMore relate.One
	db.Preload("Mores").Find(&oneByMore, 1)
	//需要加条件： 都知道条件了为什么不直接更新，关联的意义何在？
	db.Model(&oneByMore.Mores).Where("id = ?", 1).Update("hobbies", "uuuu")
}

func CreateManyToMany(ctx *gin.Context) {
	db, err := gorm.Open("mysql", "root:123456@tcp(localhost:3306)/gin?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close() //关闭空闲连接

	ManyOne := relate.ManyOne{
		Id:   1,
		Name: "wer",
		ManyTwo: []relate.ManyTwo{
			{
				Id:  1,
				Age: 15,
			},
			{
				Id:  2,
				Age: 17,
			},
		},
	}
	db.Create(ManyOne)
}

func QueryManyToMany(ctx *gin.Context) {
	db, err := gorm.Open("mysql", "root:123456@tcp(localhost:3306)/gin?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close() //关闭空闲连接

	var ManyOne relate.ManyOne
	db.Preload("ManyTwo").Find(&ManyOne, 1)
	fmt.Println(ManyOne)
}
