package log

import (
	"github.com/beego/bee/logger"
	"log"
	"github.com/astaxie/beego/logs"
	"encoding/json"
)

func init() {
	loadLogger()
}

func loadLogger() {

	config := make(map[string]string)

	config["filename"]=""
	config["level"]=""

	json.Marshal(config)
	logs.SetLogger(logs.AdapterFile, )
}
