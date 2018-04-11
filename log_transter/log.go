package log_tranter

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"encoding/json"
	"errors"
)

func convertLogLevel(level string) int {
	switch level {
	case "debug":
		return logs.LevelDebug
	case "warn":
		return logs.LevelWarn
	case "info":
		return logs.LevelInfo
	case "trace":
		return logs.LevelTrace
	default:
		return logs.LevelDebug
	}
}

func initLogger() error {

	config := make(map[string]interface{})
	//设置日志文件
	config["filename"] = LogConfig.LogPath
	//设置日志等级
	config["level"] = convertLogLevel(LogConfig.LogLevel)
	configStr, err := json.Marshal(config)
	if err != nil {
		return errors.New(fmt.Sprintln("json parse fail", err.Error()))
	}
	//把日志配置 设置到logger中
	err = logs.SetLogger(logs.AdapterFile, string(configStr))
	if err != nil {
		return errors.New(fmt.Sprintln("SetLogger fail", err.Error()))
	}
	return nil
}
