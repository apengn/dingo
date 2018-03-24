package config

import (
	"github.com/astaxie/beego/config"
	"log"
	"encoding/json"
)

type Config struct {
	MySqlConf
	ProductConf
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

	loadProductConf(conf)

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
