package models

type Config struct {
	Ruta_essentials_player_data string           `json:"path_essentials_player_data"`
	Ruta_login_Segurity         string           `json:"path_login_Segurity_sqlite"`
	SkinsDataBase               MysqlCredentials `json:"skins_mysql_data_base"`
	Ruta_Luck_Perms             string           `json:"path_luck_perms_sqlite"`
}
