package repositories

import (
	"database/sql"

	"github.com/minecraft_bakend_web/config"
)

func GetConnectionDb() (*sql.DB, error) {
	configUser := config.LoadConfig()
	return sql.Open("mysql", configUser.SkinsDataBase.Dbuser+":"+configUser.SkinsDataBase.Dbpasswd+"@tcp("+configUser.SkinsDataBase.DbHost+")/"+configUser.SkinsDataBase.Dbname)
}
