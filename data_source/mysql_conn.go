package data_source

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var Db *gorm.DB
var err error

func init() {
	mysqlConfig := LoadMysqlConf()
	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		mysqlConfig.UserName, mysqlConfig.Password, mysqlConfig.Host, mysqlConfig.Port, mysqlConfig.Database)

	Db, err = gorm.Open("mysql", dataSource)
	if err != nil {
		panic(err)
	}
	Db.LogMode(mysqlConfig.LogModel)
	//defer Db.Close()
	Db.DB().SetMaxOpenConns(100) //最大连接数
	Db.DB().SetMaxIdleConns(50)  //最大空闲数
}
