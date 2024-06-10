package repositories

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	"github.com/minecraft_bakend_web/config"
	"github.com/minecraft_bakend_web/models"
	"github.com/minecraft_bakend_web/utils"
)

func LoginUserHandler(user models.UserLoginModel) (string, string, error) {
	configUser := config.LoadConfig()
	db, err := sql.Open("sqlite3", configUser.Ruta_login_Segurity)
	if err != nil {
		return "250", "250", fmt.Errorf("error abriendo la base de datos: %v", err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT password, last_name FROM ls_players WHERE last_name = ?", user.User)
	if err != nil {
		return err.Error(), err.Error(), fmt.Errorf("error ejecutando la consulta: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var password, lastName string
		if err := rows.Scan(&password, &lastName); err != nil {
			return err.Error(), err.Error(), fmt.Errorf("error escaneando fila: %v", err)
		}

		if utils.CompareHash(password, user.Password) {
			uuidUser, err := FindUUIDByUsername(user.User)
			if err != nil {
				return err.Error(), err.Error(), nil
			}
			return "200", uuidUser, nil
		}
	}
	return "404", "404", nil
}
