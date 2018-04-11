package model

import (
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"github.com/dingoblog/dingo/config"
	"github.com/astaxie/beego/logs"
	"database/sql"
)

var (
	Driver = "mysql"
)

var db *sql.DB

//初始化mysql数据库的连接
func initConnection() {
	var err error
	//db, err = gorm.Open(Driver, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true",
	//	config.Conf.MySqlConf.DbUser,
	//	config.Conf.MySqlConf.DbPwd,
	//	config.Conf.MySqlConf.DbHost,
	//	config.Conf.MySqlConf.DbPort,
	//	config.Conf.MySqlConf.DbName))
	db, err = sql.Open(Driver, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true",
		config.Conf.MySqlConf.DbUser,
		config.Conf.MySqlConf.DbPwd,
		config.Conf.MySqlConf.DbHost,
		config.Conf.MySqlConf.DbPort,
		config.Conf.MySqlConf.DbName))
	if err != nil {
		logs.Error(err)
		panic(err)
	}
	logs.Info("open mysql success")
}

//创建表
func createTableIfNotExist() {
	tx, err := db.Begin()
	if err != nil {
		tx.Rollback()
	}
	for _, createTable := range CreateTables.createTableStr {
		_, err = tx.Exec(createTable)
		if err != nil {
			tx.Rollback()
			return
		}
	}
	tx.Commit()
	checkBlogSettings()
}

func checkBlogSettings() {
	SetSettingIfNotExists("theme", "default", "blog")
	SetSettingIfNotExists("title", "My Blog", "blog")
	SetSettingIfNotExists("description", "Awesome blog created by Dingo.", "blog")
}
