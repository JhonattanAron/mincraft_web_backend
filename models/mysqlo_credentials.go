package models

type MysqlCredentials struct {
	Dbuser   string `json:"dbuser"`
	Dbpasswd string `json:"dbpasswd"`
	Dbname   string `json:"dbname"`
	DbHost   string `json:"dbhost"`
}
