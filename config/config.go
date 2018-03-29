package config

import (
	"github.com/astaxie/beego/config"
	"log"
	"encoding/json"
	"github.com/astaxie/beego/logs"
)

type Config struct {
	MySqlConf
	ProductConf
	LogConf
}

var (
	Conf Config
)

func init() {
	loadConfig("ini", "./config/config.conf")
}
func loadConfig(adapterName, filename string) {
	conf, err := config.NewConfig(adapterName, filename)
	if err != nil {
		panic(err)
	}
	Conf = Config{}
	//获取mysql的配置
	loadMySqlConf(conf)
	//加载项目的配置
	loadProductConf(conf)
    //加载Logger配置
	loadLogConf(conf)
	if b, e := json.Marshal(Conf); e == nil {
		log.Println(string(b))
	} else {
		panic(e)
	}

}

func loadProductConf(conf config.Configer) {
	Conf.ProductConf.RunHost = conf.String("product::host")
	Conf.ProductConf.RunPort = conf.String("product::port")
}

func loadMySqlConf(conf config.Configer) {
	Conf.MySqlConf.DbHost = conf.String("mysql::db_host")
	Conf.MySqlConf.DbPort = conf.String("mysql::db_port")
	Conf.MySqlConf.DbName = conf.String("mysql::db_name")
	Conf.MySqlConf.DbPwd = conf.String("mysql::db_pwd")
	Conf.MySqlConf.DbUser = conf.String("mysql::db_user")
}
func loadLogConf(conf config.Configer) {
	Conf.LogConf.FileName = conf.String("log::filename")
	level, err := conf.Int("log::level")
	if err != nil {
		Conf.LogConf.Level = logs.LevelDebug
	} else {
		Conf.LogConf.Level = level
	}
}
