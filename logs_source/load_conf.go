package logs_source

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type LogConfig struct {
	LogDir   string `json:"log_dir"`
	LogLevel string `json:"log_level"`
}

func LoadLogConfig() *LogConfig {
	logConf := LogConfig{}
	file, err := os.Open("config/log_conf.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	byteDatas, err1 := ioutil.ReadAll(file)
	if err1 != nil {
		panic(err1)
	}
	err2 := json.Unmarshal(byteDatas, &logConf)
	if err2 != nil {
		panic(err2)
	}
	return &logConf
}
