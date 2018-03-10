package model

import (
	_ "github.com/go-sql-driver/mysql"
	"io/ioutil"
	"os"
	"fmt"
	"path/filepath"
	"database/sql"
	"encoding/json"
)

type DbConfig struct {
	DbHost string `json:"db_host"`
	DbPort int    `json:"db_port"`
	DbUser string `json:"db_user"`
	DbPwd  string `json:"db_pwd"`
	DbName string `json:"db_name"`
}

var (
	Driver = "mysql"
)

var dbConfig DbConfig
var db *sql.DB

func initConnection() {
	initDbConfig()
	var erro error
	db, erro = sql.Open(Driver, fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=true",
		dbConfig.DbUser, dbConfig.DbPwd, dbConfig.DbHost, dbConfig.DbPort, dbConfig.DbName))
	if erro != nil {
		panic(erro)
	}
}

func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}
func initDbConfig() {
	workroot, erro := os.Getwd()
	if erro != nil {
		panic(erro)
	}
	dbFilePath := filepath.Join(workroot, "./config/config.json")

	if fileExists(dbFilePath) {

	}
	data, erro := ioutil.ReadFile(dbFilePath)
	if erro != nil {
		panic(erro)
	}
	erro = json.Unmarshal([]byte(string(data)), &dbConfig)
	if erro != nil {
		panic(erro.Error())
	}
}

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
