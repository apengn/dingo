package model

import (
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"database/sql"
	"github.com/dingoblog/dingo/config"
)


var (
	Driver = "mysql"
)

var db *sql.DB

//初始化mysql数据库的连接
func initConnection() {
	var erro error
	db, erro = sql.Open(Driver, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true",
		config.Conf.MySqlConf.DbUser,
		config.Conf.MySqlConf.DbPwd,
		config.Conf.MySqlConf.DbHost,
		config.Conf.MySqlConf.DbPort,
		config.Conf.MySqlConf.DbName))
	if erro != nil {
		panic(erro)
	}
}

//创建表
func createTableIfNotExist() error {
	tx, erro := db.Begin()
	if erro != nil {
		tx.Rollback()
		return erro
	}
	for _, createTable := range CreateTables.createTableStr {
		if _, erro := tx.Exec(createTable); erro != nil {
			tx.Rollback()
			return erro
		}
	}
	tx.Commit()
	checkBlogSettings()
	return erro
}

func checkBlogSettings() {
	SetSettingIfNotExists("theme", "default", "blog")
	SetSettingIfNotExists("title", "My Blog", "blog")
	SetSettingIfNotExists("description", "Awesome blog created by Dingo.", "blog")
}
