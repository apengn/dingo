package log

import (
	"github.com/astaxie/beego/logs"
	"encoding/json"
	config2 "github.com/dingoblog/dingo/config"
	"fmt"
)

func init() {
	initLogger()
}

//初始化日志
func initLogger() {
	config := make(map[string]interface{})
	//设置日志文件
	config["filename"] = config2.Conf.FileName
	//设置日志等级
	config["level"] = config2.Conf.Level
	configStr, err := json.Marshal(config)
	if err != nil {
		panic(fmt.Sprintln("initLogger err", err.Error()))
		return
	}
	//把日志配置 设置到logger中
	err = logs.SetLogger(logs.AdapterFile, string(configStr))
	if err != nil {
		panic(fmt.Sprintln("SetLogger err", err.Error()))
		return
	}
}
