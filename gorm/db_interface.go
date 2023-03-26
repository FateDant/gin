package gorm

import (
	"fmt"
	"gin/gorm/models"
	"gin/gorm/models/relate"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// First 查询第一个 last 查询最后一个 take查询一条数据 find 可以查询多条
func First(ctx *gin.Context) {
	db, err := gorm.Open("mysql", "root:123456@tcp(localhost:3306)/gin?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close() //关闭空闲连接

	var user models.User
	db.First(&user, 1)
	fmt.Println(user)

	var user1 models.User //必须得重新初始化一个变量模型
	res := db.Where("name = ?", "wer").First(&user1)
	fmt.Println(res.RowsAffected)
	fmt.Println(user1)

	var user2 models.User
	user3 := models.User{
		Name:  "asf",
		Age:   "981",
		Pic:   "adf",
		Phone: "1234",
	}
	db.Debug().FirstOrCreate(&user2, user3)

	var user4 models.User
	db.Last(&user4)
	fmt.Println(user4)
}

func Where(ctx *gin.Context) {
	db, err := gorm.Open("mysql", "root:123456@tcp(localhost:3306)/gin?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close() //关闭空闲连接

	var user models.User
	db.Debug().Where("name = ?", "wer").Find(&user)
	fmt.Println(user)

	var user2 models.User
	db.Debug().Select("name,age").Where("name = ?", "wer").Find(&user2)
	fmt.Println(user2)
}

func Create(ctx *gin.Context) {
	db, err := gorm.Open("mysql", "root:123456@tcp(localhost:3306)/gin?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close() //关闭空闲连接
	//var user = models.User{
	//	Name:  "uuuu",
	//	Age:   "78",
	//	Pic:   "werw",
	//	Phone: "uiqwerqwr",
	//}
	//db.Create(&user)

	//gorm 不支持批量插入
	//user2 := []models.User{
	//	{
	//		Name:  "TTT",
	//		Age:   "78",
	//		Pic:   "adsf",
	//		Phone: "cbvb",
	//	},
	//	{
	//		Name:  "JJJ",
	//		Age:   "23",
	//		Pic:   "asdf",
	//		Phone: "er",
	//	},
	//}
	//db.Create(&user2)

	//user3 := models.User{
	//	Name:  "JJJ",
	//	Age:   "23",
	//	Pic:   "asdf",
	//	Phone: "er",
	//}
	//db.Save(&user3)

	user4 := models.User{}
	db.Where("name = ?", "cccc").First(&user4)
	fmt.Println(user4)
	user4.Name = "dddd"
	//db.Save(&user4)
	db.Create(&user4)
}

func Update(ctx *gin.Context) {
	db, err := gorm.Open("mysql", "root:123456@tcp(localhost:3306)/gin?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close() //关闭空闲连接

	user := models.User{}
	db.Where("name = ?", "cccc").First(&user)
	fmt.Println(user)
	db.Model(&user).Update("age", "23")

	user2 := models.User{}
	//db.Debug().Where("name = ?", "JJJ").First(&user2).Update("age", "uuu") //只能更新一条
	//db.Debug().Where("name = ?", "JJJ").Find(&user2).Update("age", "iiii") //只能更新一条
	//db.Debug().Where("name = ?", "JJJ").Find(&user2).Update(models.User{Age: "iioo"}) //只能更新一条
	//db.Debug().Where("name = ?", "JJJ").Find(&user2).Update(map[string]interface{}{"Age": "wer"}) //只能更新一条
	//db.Model(&user2).Where("name = ?", "JJJ").Updates(models.User{Age: "zxc", Pic: "789"}) //批量更新
	fmt.Println(user2)

	user3 := models.User{}
	db.Model(&user3).Where("name = ?", "asf").Delete(&user3) //可以批量删除
	fmt.Println(user3)
}

func Not(ctx *gin.Context) {
	db, err := gorm.Open("mysql", "root:123456@tcp(localhost:3306)/gin?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close() //关闭空闲连接

	var user []models.User //得是切片才能接收多条数据
	db.Not("name = ?", "JJJ").Find(&user)
	fmt.Println(user)

	var user2 []models.User
	db.Debug().Where("name = ?", "JJJ").Or("name = ?", "cccc").Find(&user2)
	fmt.Println(user2)

	var user3 []models.User
	db.Debug().Where("name LIKE ?", "%J%").Order("id desc").Find(&user3)
	fmt.Println(user3)
}

func Group(ctx *gin.Context) {
	db, err := gorm.Open("mysql", "root:123456@tcp(localhost:3306)/gin?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close() //关闭空闲连接

	//定义一个要被赋值的结构体
	type GroupData struct {
		Name  string
		Count int
		Age   string
	}

	var user []models.User
	var groupData []GroupData
	db.Select("name,count(*) as count,age").Group("name").Having("age = ?", "zxc").Find(&user).Scan(&groupData)
	//scan 查出来的值扫描到结构体上！！ 查出来的数据不能放在user结构体里面，所以需要定义一个新的结构体来进行赋值
	// having 过滤的条件，必须先在select里面查询出来
	fmt.Println(user)
	fmt.Println(groupData)
}

func Join(ctx *gin.Context) {
	db, err := gorm.Open("mysql", "root:123456@tcp(localhost:3306)/gin?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close() //关闭空闲连接

	type MasterAndSlaves struct {
		Id    int
		Name  string
		Age   int
		Addr  string
		Pic   string
		Phone string
	}

	var masterAndSlaves []MasterAndSlaves
	var userMaster relate.UserMaster
	//db.Joins("left join user_slaves on user_masters.id = user_slaves.id").Find(&userMaster).Scan(&masterAndSlaves)
	db.Debug().Select("*").Joins("left join user_slaves on user_masters.id = user_slaves.id").Find(&userMaster).Scan(&masterAndSlaves)
	fmt.Println(masterAndSlaves) //所有字段必须要加select （*）
	fmt.Println(userMaster)
}

func FirstInit(ctx *gin.Context) {
	db, err := gorm.Open("mysql", "root:123456@tcp(localhost:3306)/gin?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close() //关闭空闲连接
	//全局打印所有的sql信息
	db.LogMode(true)

	//firstInit 如果查询到就返回查询结果，否则就初始化结构体。
	//{0 10   }
	var user models.User
	db.FirstOrInit(&user, models.User{Name: "10"})
	fmt.Println(user)

	//Attr初始化 、Assign替换
}
