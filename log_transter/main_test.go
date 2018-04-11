package log_tranter

import (
	"testing"
	"fmt"
	"github.com/astaxie/beego/logs"
)

func Test_es(t *testing.T) {

	err := initConfig("ini", "./config.conf")
	if err != nil {
		panic(err)
		return
	}
	fmt.Println(LogConfig)
	err = initLogger()
	if err != nil {
		panic(err)
		return
	}
	logs.Debug("init logger success")


	err = initKafka()
	if err != nil {
		logs.Error("init kafka fail", err)
		return
	}
	logs.Debug("init kafka success")
	err = initEs()
	if err != nil {
		logs.Error("init es fail", err)
		return
	}
	//
	err = run()
	if err != nil {
		logs.Error("run fail", err)
		return
	}
	logs.Warn("warning:", "exit")
}
