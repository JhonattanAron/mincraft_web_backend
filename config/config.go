package config

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/minecraft_bakend_web/models"
)

const configFileName = "config.json"

func LoadConfig() models.Config {
	var config models.Config

	if _, err := os.Stat(configFileName); os.IsNotExist(err) {
		// El archivo no existe, crear con valores predeterminados
		config = models.Config{
			Ruta_essentials_player_data: "D:/test/Mohist Server/plugins/Essentials/userdata",
			Ruta_login_Segurity:         "D:/test/Mohist Server/plugins/LoginSecurity/LoginSecurity.db",
			SkinsDataBase: models.MysqlCredentials{
				Dbname:   "db_name",
				Dbuser:   "aron",
				Dbpasswd: "passwd",
				DbHost:   "127.0.0.1:3306",
			},
			Ruta_Luck_Perms: "D:/test/Mohist Server/plugins/LuckPerms/luckperms-sqlite.db",
		}
		SaveConfig(config)
	} else {
		// El archivo existe, leer y decodificar
		file, err := os.Open(configFileName)
		if err != nil {
			fmt.Println("Error abriendo archivo de configuración:", err)
			os.Exit(1)
		}
		defer file.Close()

		decoder := json.NewDecoder(file)
		err = decoder.Decode(&config)
		if err != nil {
			fmt.Println("Error decodificando archivo de configuración:", err)
			os.Exit(1)
		}
	}

	return config
}

func SaveConfig(config models.Config) {
	file, err := os.Create(configFileName)
	if err != nil {
		fmt.Println("Error creando archivo de configuración:", err)
		os.Exit(1)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(config)
	if err != nil {
		fmt.Println("Error codificando archivo de configuración:", err)
		os.Exit(1)
	}
}
