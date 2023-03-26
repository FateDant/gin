package data_source

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type MysqlConf struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	UserName string `json:"user_name"`
	Password string `json:"password"`
	Database string `json:"database"`
	LogModel bool   `json:"log_model"`
}

func LoadMysqlConf() *MysqlConf {
	mysql_conf := MysqlConf{}
	file, err := os.Open("config/mysql.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	byteData, err1 := ioutil.ReadAll(file)
	if err1 != nil {
		panic(err1)
	}

	err2 := json.Unmarshal(byteData, &mysql_conf)
	if err2 != nil {
		panic(err2)
	}
	return &mysql_conf
}
