package config

type MySqlConf struct {
	DbHost string `json:"db_host"`
	DbPort string `json:"db_port"`
	DbUser string `json:"db_user"`
	DbPwd  string `json:"db_pwd"`
	DbName string `json:"db_name"`
}